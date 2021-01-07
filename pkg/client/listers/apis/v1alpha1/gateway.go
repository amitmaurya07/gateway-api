/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "sigs.k8s.io/service-apis/apis/v1alpha1"
)

// GatewayLister helps list Gateways.
// All objects returned here must be treated as read-only.
type GatewayLister interface {
	// List lists all Gateways in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Gateway, err error)
	// Gateways returns an object that can list and get Gateways.
	Gateways(namespace string) GatewayNamespaceLister
	GatewayListerExpansion
}

// gatewayLister implements the GatewayLister interface.
type gatewayLister struct {
	indexer cache.Indexer
}

// NewGatewayLister returns a new GatewayLister.
func NewGatewayLister(indexer cache.Indexer) GatewayLister {
	return &gatewayLister{indexer: indexer}
}

// List lists all Gateways in the indexer.
func (s *gatewayLister) List(selector labels.Selector) (ret []*v1alpha1.Gateway, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Gateway))
	})
	return ret, err
}

// Gateways returns an object that can list and get Gateways.
func (s *gatewayLister) Gateways(namespace string) GatewayNamespaceLister {
	return gatewayNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// GatewayNamespaceLister helps list and get Gateways.
// All objects returned here must be treated as read-only.
type GatewayNamespaceLister interface {
	// List lists all Gateways in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Gateway, err error)
	// Get retrieves the Gateway from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Gateway, error)
	GatewayNamespaceListerExpansion
}

// gatewayNamespaceLister implements the GatewayNamespaceLister
// interface.
type gatewayNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Gateways in the indexer for a given namespace.
func (s gatewayNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Gateway, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Gateway))
	})
	return ret, err
}

// Get retrieves the Gateway from the indexer for a given namespace and name.
func (s gatewayNamespaceLister) Get(name string) (*v1alpha1.Gateway, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("gateway"), name)
	}
	return obj.(*v1alpha1.Gateway), nil
}
