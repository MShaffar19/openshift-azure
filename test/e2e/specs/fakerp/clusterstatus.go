package fakerp

import (
	"context"
	"encoding/json"
	"os"
	"sort"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	v1 "k8s.io/api/core/v1"

	"github.com/openshift/openshift-azure/test/clients/azure"
)

var _ = Describe("Control Plane Pods Status E2E tests [EveryPR]", func() {
	It("should allow an SRE to fetch the status of control plane pods", func() {
		By("Using the OSA admin client to fetch the raw cluster status")
		b, err := azure.RPClient.OpenShiftManagedClustersAdmin.GetControlPlanePods(context.Background(), os.Getenv("RESOURCEGROUP"), os.Getenv("RESOURCEGROUP"))
		Expect(err).NotTo(HaveOccurred())
		Expect(b).NotTo(BeNil())

		By("Unmarshalling the returned raw cluster status to a Pod slice")
		var podArray []v1.Pod
		err = json.Unmarshal(b, &podArray)
		Expect(err).NotTo(HaveOccurred())

		By("Constructing a mapping of namespaces to pods")
		podnames := make(map[string][]string)
		for _, pod := range podArray {
			list := append(podnames[pod.Namespace], pod.Name)
			sort.Strings(list)
			podnames[pod.Namespace] = list
		}

		By("Verifying that certain namespaces contain certain expected pods")
		namespace := "kube-system"
		Expect(sort.SearchStrings(podnames[namespace], "controllers-master-000000")).NotTo(Equal(len(podnames[namespace])))
		Expect(sort.SearchStrings(podnames[namespace], "controllers-master-000001")).NotTo(Equal(len(podnames[namespace])))
		Expect(sort.SearchStrings(podnames[namespace], "controllers-master-000002")).NotTo(Equal(len(podnames[namespace])))
		Expect(sort.SearchStrings(podnames[namespace], "master-api-master-000000")).NotTo(Equal(len(podnames[namespace])))
		Expect(sort.SearchStrings(podnames[namespace], "master-api-master-000001")).NotTo(Equal(len(podnames[namespace])))
		Expect(sort.SearchStrings(podnames[namespace], "master-api-master-000002")).NotTo(Equal(len(podnames[namespace])))
		Expect(sort.SearchStrings(podnames[namespace], "master-etcd-master-000000")).NotTo(Equal(len(podnames[namespace])))
		Expect(sort.SearchStrings(podnames[namespace], "master-etcd-master-000001")).NotTo(Equal(len(podnames[namespace])))
		Expect(sort.SearchStrings(podnames[namespace], "master-etcd-master-000002")).NotTo(Equal(len(podnames[namespace])))
		Expect(sort.SearchStrings(podnames[namespace], "sync-")).NotTo(Equal(len(podnames[namespace])))
	})
})
