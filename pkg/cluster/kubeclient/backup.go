package kubeclient

import (
	"context"
	"fmt"
	"time"

	"github.com/Azure/go-autorest/autorest/to"
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/openshift/openshift-azure/pkg/api/validate"
	"github.com/openshift/openshift-azure/pkg/util/ready"
	"github.com/openshift/openshift-azure/pkg/util/wait"
)

const (
	etcdBackupNamespace = "openshift-etcd"
)

func (u *kubeclient) BackupCluster(ctx context.Context, backupName string) error {
	u.log.Infof("running an etcd backup")
	if len(backupName) > 255 || !validate.RxRfc1123.MatchString(backupName) {
		return fmt.Errorf("invalid backup name: %s", backupName)
	}
	cronjob, err := u.client.BatchV1beta1().CronJobs(etcdBackupNamespace).Get("etcd-backup", metav1.GetOptions{})
	if err != nil {
		return err
	}
	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      backupName,
			Namespace: etcdBackupNamespace,
		},
		Spec: cronjob.Spec.JobTemplate.Spec,
	}
	job.Spec.BackoffLimit = to.Int32Ptr(0)
	job.Spec.Template.Spec.Containers[0].Args = []string{fmt.Sprintf("-blobname=%s", backupName), "save"}
	backup, err := u.client.BatchV1().Jobs(etcdBackupNamespace).Create(job)
	if err != nil {
		return err
	}
	defer func() {
		err = u.client.BatchV1().Jobs(etcdBackupNamespace).Delete(backupName, &metav1.DeleteOptions{})
		if err != nil {
			u.log.Infof("failed to delete job: %s", backupName)
		}
	}()
	err = wait.PollImmediateUntil(2*time.Second, func() (bool, error) {
		return ready.BatchIsReady(u.client.BatchV1().Jobs(backup.Namespace), backup.Name)()
	}, ctx.Done())
	if err != nil {
		return err
	}
	// TODO: verify that the backup blob was created/exists and has bytes
	if backup.Status.Failed > 0 {
		return fmt.Errorf("backup pod failed: %s", backupName)
	}
	return nil
}
