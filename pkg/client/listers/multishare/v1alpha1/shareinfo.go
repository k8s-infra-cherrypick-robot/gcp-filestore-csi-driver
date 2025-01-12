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
	v1alpha1 "sigs.k8s.io/gcp-filestore-csi-driver/pkg/apis/multishare/v1alpha1"
)

// ShareInfoLister helps list ShareInfos.
// All objects returned here must be treated as read-only.
type ShareInfoLister interface {
	// List lists all ShareInfos in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ShareInfo, err error)
	// Get retrieves the ShareInfo from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ShareInfo, error)
	ShareInfoListerExpansion
}

// shareInfoLister implements the ShareInfoLister interface.
type shareInfoLister struct {
	indexer cache.Indexer
}

// NewShareInfoLister returns a new ShareInfoLister.
func NewShareInfoLister(indexer cache.Indexer) ShareInfoLister {
	return &shareInfoLister{indexer: indexer}
}

// List lists all ShareInfos in the indexer.
func (s *shareInfoLister) List(selector labels.Selector) (ret []*v1alpha1.ShareInfo, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ShareInfo))
	})
	return ret, err
}

// Get retrieves the ShareInfo from the index for a given name.
func (s *shareInfoLister) Get(name string) (*v1alpha1.ShareInfo, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("shareinfo"), name)
	}
	return obj.(*v1alpha1.ShareInfo), nil
}
