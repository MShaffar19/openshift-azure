package cluster

//go:generate go get github.com/golang/mock/mockgen
//go:generate mockgen -destination=../util/mocks/mock_$GOPACKAGE/types.go github.com/openshift/openshift-azure/pkg/$GOPACKAGE Upgrader
//go:generate gofmt -s -l -w ../util/mocks/mock_$GOPACKAGE/types.go
//go:generate go get golang.org/x/tools/cmd/goimports
//go:generate goimports -local=github.com/openshift/openshift-azure -e -w ../util/mocks/mock_$GOPACKAGE/types.go

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/openshift/openshift-azure/pkg/api"
	"github.com/openshift/openshift-azure/pkg/arm"
	"github.com/openshift/openshift-azure/pkg/cluster/kubeclient"
	"github.com/openshift/openshift-azure/pkg/cluster/scaler"
	"github.com/openshift/openshift-azure/pkg/cluster/updateblob"
	"github.com/openshift/openshift-azure/pkg/startup"
	"github.com/openshift/openshift-azure/pkg/util/azureclient"
	"github.com/openshift/openshift-azure/pkg/util/azureclient/storage"
	"github.com/openshift/openshift-azure/pkg/util/enrich"
)

// here follow well known container and blob names
const (
	ConfigContainerName     = "config"
	SyncBlobName            = "sync"
	MasterStartupBlobName   = "master-startup"
	WorkerStartupBlobName   = "worker-startup"
	EtcdBackupContainerName = "etcd"
)

// Upgrader is the public interface to the upgrade module used by the plugin.
type Upgrader interface {
	CreateOrUpdateConfigStorageAccount(ctx context.Context, cs *api.OpenShiftManagedCluster) error
	EnrichCertificatesFromVault(ctx context.Context, cs *api.OpenShiftManagedCluster) error
	EnrichStorageAccountKeys(ctx context.Context, cs *api.OpenShiftManagedCluster) error
	InitializeUpdateBlob(cs *api.OpenShiftManagedCluster, suffix string) error
	WaitForHealthzStatusOk(ctx context.Context, cs *api.OpenShiftManagedCluster) error
	HealthCheck(ctx context.Context, cs *api.OpenShiftManagedCluster) *api.PluginError
	SortedAgentPoolProfilesForRole(cs *api.OpenShiftManagedCluster, role api.AgentPoolProfileRole) []api.AgentPoolProfile
	WaitForNodesInAgentPoolProfile(ctx context.Context, cs *api.OpenShiftManagedCluster, app *api.AgentPoolProfile, suffix string) error
	UpdateMasterAgentPool(ctx context.Context, cs *api.OpenShiftManagedCluster, app *api.AgentPoolProfile) *api.PluginError
	UpdateWorkerAgentPool(ctx context.Context, cs *api.OpenShiftManagedCluster, app *api.AgentPoolProfile, suffix string) *api.PluginError
	CreateOrUpdateSyncPod(ctx context.Context, cs *api.OpenShiftManagedCluster) error
	EtcdBlobExists(ctx context.Context, blobName string) error
	EtcdRestoreDeleteMasterScaleSet(ctx context.Context, cs *api.OpenShiftManagedCluster) *api.PluginError
	EtcdRestoreDeleteMasterScaleSetHashes(ctx context.Context, cs *api.OpenShiftManagedCluster) *api.PluginError
	ResetUpdateBlob(cs *api.OpenShiftManagedCluster) error
	Reimage(ctx context.Context, cs *api.OpenShiftManagedCluster, scaleset, instanceID string) error
	ListVMHostnames(ctx context.Context, cs *api.OpenShiftManagedCluster) ([]string, error)
	RunCommand(ctx context.Context, cs *api.OpenShiftManagedCluster, scaleset, instanceID, command string) error
	WriteStartupBlobs(cs *api.OpenShiftManagedCluster) error
	GenerateARM(ctx context.Context, cs *api.OpenShiftManagedCluster, backupBlob string, isUpdate bool, suffix string) (map[string]interface{}, error)

	kubeclient.Kubeclient
}

type simpleUpgrader struct {
	kubeclient.Kubeclient

	testConfig        api.TestConfig
	accountsClient    azureclient.AccountsClient
	storageClient     storage.Client
	updateBlobService updateblob.BlobService
	vmc               azureclient.VirtualMachineScaleSetVMsClient
	ssc               azureclient.VirtualMachineScaleSetsClient
	kvc               azureclient.KeyVaultClient
	log               *logrus.Entry
	scalerFactory     scaler.Factory
	hasher            Hasher
	arm               arm.Interface
}

var _ Upgrader = &simpleUpgrader{}

// NewSimpleUpgrader creates a new upgrader instance
func NewSimpleUpgrader(ctx context.Context, log *logrus.Entry, cs *api.OpenShiftManagedCluster, initializeStorageClients, disableKeepAlives bool, testConfig api.TestConfig) (Upgrader, error) {
	authorizer, err := azureclient.GetAuthorizerFromContext(ctx, api.ContextKeyClientAuthorizer)
	if err != nil {
		return nil, err
	}

	vaultauthorizer, err := azureclient.GetAuthorizerFromContext(ctx, api.ContextKeyVaultClientAuthorizer)
	if err != nil {
		return nil, err
	}

	kubeclient, err := kubeclient.NewKubeclient(log, cs.Config.AdminKubeconfig, disableKeepAlives)
	if err != nil {
		return nil, err
	}

	arm, err := arm.New(ctx, log, cs, testConfig)
	if err != nil {
		return nil, err
	}

	u := &simpleUpgrader{
		Kubeclient: kubeclient,

		testConfig:     testConfig,
		accountsClient: azureclient.NewAccountsClient(ctx, cs.Properties.AzProfile.SubscriptionID, authorizer),
		vmc:            azureclient.NewVirtualMachineScaleSetVMsClient(ctx, cs.Properties.AzProfile.SubscriptionID, authorizer),
		ssc:            azureclient.NewVirtualMachineScaleSetsClient(ctx, cs.Properties.AzProfile.SubscriptionID, authorizer),
		kvc:            azureclient.NewKeyVaultClient(ctx, vaultauthorizer),
		log:            log,
		scalerFactory:  scaler.NewFactory(),
		hasher: &hasher{
			log:            log,
			testConfig:     testConfig,
			startupFactory: startup.New,
			arm:            arm,
		},
		arm: arm,
	}

	if initializeStorageClients {
		err = u.initializeStorageClients(ctx, cs)
		if err != nil {
			return nil, err
		}
	}

	return u, nil
}

func (u *simpleUpgrader) EnrichCertificatesFromVault(ctx context.Context, cs *api.OpenShiftManagedCluster) error {
	return enrich.CertificatesFromVault(ctx, u.kvc, cs)
}

func (u *simpleUpgrader) EnrichStorageAccountKeys(ctx context.Context, cs *api.OpenShiftManagedCluster) error {
	return enrich.StorageAccountKeys(ctx, u.accountsClient, cs)
}

func (u *simpleUpgrader) GenerateARM(ctx context.Context, cs *api.OpenShiftManagedCluster, backupBlob string, isUpdate bool, suffix string) (map[string]interface{}, error) {
	err := enrich.SASURIs(u.storageClient, cs)
	if err != nil {
		return nil, err
	}

	return u.arm.Generate(ctx, backupBlob, isUpdate, suffix)
}
