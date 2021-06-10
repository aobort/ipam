package controllers

import (
	"context"
	"github.com/onmetal/ipam/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("IP controller", func() {
	const (
		timeout   = time.Second * 5
		interval  = time.Millisecond * 250
		Namespace = "default"
	)

	cidrMustParse := func(cidrString string) *v1alpha1.CIDR {
		cidr, err := v1alpha1.CIDRFromString(cidrString)
		if err != nil {
			panic(err)
		}
		return cidr
	}

	Context("IP controller test", func() {
		FIt("Should get IP assigned", func() {
			ctx := context.Background()
			subnet := &v1alpha1.Subnet{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "subnet1",
					Namespace: Namespace,
				},
				Spec: v1alpha1.SubnetSpec{
					CIDR:              *cidrMustParse("10.0.0.0/30"),
					ParentSubnetName:  "subnet1",
					Regions:           []string{"euw"},
					AvailabilityZones: []string{"a"},
				},
				Status: v1alpha1.SubnetStatus{
					Type: v1alpha1.CIPv4SubnetType,
				},
			}
			By("Expecting Subnet Create Successful")

			Expect(k8sClient.Create(ctx, subnet)).Should(Succeed())

			createdSubnet := v1alpha1.Subnet{}
			namespacedName := types.NamespacedName{
				Name:      "subnet1",
				Namespace: Namespace,
			}

			By("Subnet CIDR is reserved in Network")
			Eventually(func() bool {
				err := k8sClient.Get(ctx, namespacedName, &createdSubnet)
				if err != nil {
					return false
				}
				if createdSubnet.Status.State != v1alpha1.CFinishedSubnetState {
					return false
				}
				return true
			}, timeout, interval).Should(BeTrue())

			ip := &v1alpha1.Ip{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "ip1",
					Namespace: Namespace,
				},
				Spec: v1alpha1.IpSpec{
					Subnet: "subnet1",
					IP:     "10.0.0.1",
				},
			}
			By("Expecting Ip Create Successful")
			Expect(k8sClient.Create(ctx, ip)).Should(Succeed())

			Eventually(func() bool {
				err := k8sClient.Get(ctx, namespacedName, &createdSubnet)
				if err != nil {
					return false
				}
				return len(createdSubnet.Status.Vacant) == 2
			}, timeout, interval).Should(BeTrue())

			Expect(k8sClient.Delete(ctx, ip)).Should(Succeed())

			Eventually(func() bool {
				err := k8sClient.Get(ctx, namespacedName, &createdSubnet)
				if err != nil {
					return false
				}
				return len(createdSubnet.Status.Vacant) == 1
			}, timeout, interval).Should(BeTrue())
		})
	})
})
