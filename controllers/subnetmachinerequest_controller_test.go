package controllers

import (
	"context"
	machinerequestv1alpha1 "github.com/onmetal/k8s-machine-requests/api/v1alpha1"
	"github.com/onmetal/k8s-subnet-machine-request/api/v1alpha1"
	subnetv1alpha1 "github.com/onmetal/k8s-subnet/api/v1alpha1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("SubnetMachineRequest controller", func() {
	Context("SubnetMachineRequest controller test", func() {
		It("Should allocate free IP", func() {
			machine := &machinerequestv1alpha1.MachineRequest{
				TypeMeta: metav1.TypeMeta{
					APIVersion: ApiVersion,
					Kind:       "MachineRequest",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "machinerequest1",
					Namespace: Namespace,
				},
				Spec: machinerequestv1alpha1.MachineRequestSpec{
					Name: "machinerequest1",
				},
			}
			By("Expecting Machine Request Create Successful")
			ctx := context.Background()
			Expect(k8sClient.Create(ctx, machine)).Should(Succeed())

			subnet := &subnetv1alpha1.Subnet{
				TypeMeta: metav1.TypeMeta{
					APIVersion: ApiVersion,
					Kind:       "Subnet",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "subnet1",
					Namespace: Namespace,
				},
				Spec: subnetv1alpha1.SubnetSpec{
					Type: "ipv4",
					CIDR: "10.12.34.0/24",
				},
			}
			By("Expecting Subnet 1 Create Successful")
			Expect(k8sClient.Create(ctx, subnet)).Should(Succeed())

			subnet2 := &subnetv1alpha1.Subnet{
				TypeMeta: metav1.TypeMeta{
					APIVersion: ApiVersion,
					Kind:       "Subnet",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "subnet2",
					Namespace: Namespace,
				},
				Spec: subnetv1alpha1.SubnetSpec{
					Type:        "ipv4",
					CIDR:        "10.12.34.0/26",
					PartitionID: "subnet1",
				},
			}
			By("Expecting Subnet 2 Create Successful")
			Expect(k8sClient.Create(ctx, subnet2)).Should(Succeed())

			subnet3 := &subnetv1alpha1.Subnet{
				TypeMeta: metav1.TypeMeta{
					APIVersion: ApiVersion,
					Kind:       "Subnet",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "subnet3",
					Namespace: Namespace,
				},
				Spec: subnetv1alpha1.SubnetSpec{
					Type:        "ipv4",
					CIDR:        "10.12.34.128/25",
					PartitionID: "subnet1",
				},
			}
			By("Expecting Subnet 3 Create Successful")
			Expect(k8sClient.Create(ctx, subnet3)).Should(Succeed())

			subnetMachineRequest := &v1alpha1.SubnetMachineRequest{
				TypeMeta: metav1.TypeMeta{
					APIVersion: ApiVersion,
					Kind:       "SubnetMachineRequest",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "subnet3",
					Namespace: Namespace,
				},
				Spec: v1alpha1.SubnetMachineRequestSpec{
					Subnet:         "subnet1",
					MachineRequest: "machinerequest1",
				},
			}
			By("Expecting SubnetMachineRequest Create Successful")
			Expect(k8sClient.Create(ctx, subnetMachineRequest)).Should(Succeed())
		})
	})
})
