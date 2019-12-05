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

package v1

import (
	v1 "tkestack.io/cron-hpa/pkg/apis/cronhpacontroller/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CronHPALister helps list CronHPAs.
type CronHPALister interface {
	// List lists all CronHPAs in the indexer.
	List(selector labels.Selector) (ret []*v1.CronHPA, err error)
	// CronHPAs returns an object that can list and get CronHPAs.
	CronHPAs(namespace string) CronHPANamespaceLister
	CronHPAListerExpansion
}

// cronHPALister implements the CronHPALister interface.
type cronHPALister struct {
	indexer cache.Indexer
}

// NewCronHPALister returns a new CronHPALister.
func NewCronHPALister(indexer cache.Indexer) CronHPALister {
	return &cronHPALister{indexer: indexer}
}

// List lists all CronHPAs in the indexer.
func (s *cronHPALister) List(selector labels.Selector) (ret []*v1.CronHPA, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.CronHPA))
	})
	return ret, err
}

// CronHPAs returns an object that can list and get CronHPAs.
func (s *cronHPALister) CronHPAs(namespace string) CronHPANamespaceLister {
	return cronHPANamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CronHPANamespaceLister helps list and get CronHPAs.
type CronHPANamespaceLister interface {
	// List lists all CronHPAs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.CronHPA, err error)
	// Get retrieves the CronHPA from the indexer for a given namespace and name.
	Get(name string) (*v1.CronHPA, error)
	CronHPANamespaceListerExpansion
}

// cronHPANamespaceLister implements the CronHPANamespaceLister
// interface.
type cronHPANamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CronHPAs in the indexer for a given namespace.
func (s cronHPANamespaceLister) List(selector labels.Selector) (ret []*v1.CronHPA, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.CronHPA))
	})
	return ret, err
}

// Get retrieves the CronHPA from the indexer for a given namespace and name.
func (s cronHPANamespaceLister) Get(name string) (*v1.CronHPA, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("cronhpa"), name)
	}
	return obj.(*v1.CronHPA), nil
}
