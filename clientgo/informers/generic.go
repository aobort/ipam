// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

// Code generated by informer-gen. DO NOT EDIT.

package informers

import (
	"fmt"

	v1alpha1 "github.com/ironcore-dev/ipam/api/ipam/v1alpha1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=ipam.ironcore.dev, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("ips"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Ipam().V1alpha1().IPs().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("networks"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Ipam().V1alpha1().Networks().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("networkcounters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Ipam().V1alpha1().NetworkCounters().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("subnets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Ipam().V1alpha1().Subnets().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
