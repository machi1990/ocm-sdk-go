/*
Copyright (c) 2019 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/uhc-sdk-go/pkg/client/clustersmgmt/v1

// ClusterCredentialsKind is the name of the type used to represent objects
// of type 'cluster_credentials'.
const ClusterCredentialsKind = "ClusterCredentials"

// ClusterCredentialsLinkKind is the name of the type used to represent links
// to objects of type 'cluster_credentials'.
const ClusterCredentialsLinkKind = "ClusterCredentialsLink"

// ClusterCredentialsNilKind is the name of the type used to nil references
// to objects of type 'cluster_credentials'.
const ClusterCredentialsNilKind = "ClusterCredentialsNil"

// ClusterCredentials represents the values of the 'cluster_credentials' type.
//
// Credentials of the a cluster.
type ClusterCredentials struct {
	id         *string
	href       *string
	link       bool
	kubeconfig *string
	ssh        *Sshcredentials
	admin      *AdminCredentials
}

// Kind returns the name of the type of the object.
func (o *ClusterCredentials) Kind() string {
	if o == nil {
		return ClusterCredentialsNilKind
	}
	if o.link {
		return ClusterCredentialsLinkKind
	}
	return ClusterCredentialsKind
}

// ID returns the identifier of the object.
func (o *ClusterCredentials) ID() string {
	if o != nil && o.id != nil {
		return *o.id
	}
	return ""
}

// GetID returns the identifier of the object and a flag indicating if the
// identifier has a value.
func (o *ClusterCredentials) GetID() (value string, ok bool) {
	ok = o != nil && o.id != nil
	if ok {
		value = *o.id
	}
	return
}

// Link returns true iif this is a link.
func (o *ClusterCredentials) Link() bool {
	return o != nil && o.link
}

// HREF returns the link to the object.
func (o *ClusterCredentials) HREF() string {
	if o != nil && o.href != nil {
		return *o.href
	}
	return ""
}

// GetHREF returns the link of the object and a flag indicating if the
// link has a value.
func (o *ClusterCredentials) GetHREF() (value string, ok bool) {
	ok = o != nil && o.href != nil
	if ok {
		value = *o.href
	}
	return
}

// Kubeconfig returns the value of the 'kubeconfig' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Administrator _kubeconfig_ file for the cluster.
func (o *ClusterCredentials) Kubeconfig() string {
	if o != nil && o.kubeconfig != nil {
		return *o.kubeconfig
	}
	return ""
}

// GetKubeconfig returns the value of the 'kubeconfig' attribute and
// a flag indicating if the attribute has a value.
//
// Administrator _kubeconfig_ file for the cluster.
func (o *ClusterCredentials) GetKubeconfig() (value string, ok bool) {
	ok = o != nil && o.kubeconfig != nil
	if ok {
		value = *o.kubeconfig
	}
	return
}

// SSH returns the value of the 'SSH' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// SSH key pair of the cluster.
func (o *ClusterCredentials) SSH() *Sshcredentials {
	if o == nil {
		return nil
	}
	return o.ssh
}

// GetSSH returns the value of the 'SSH' attribute and
// a flag indicating if the attribute has a value.
//
// SSH key pair of the cluster.
func (o *ClusterCredentials) GetSSH() (value *Sshcredentials, ok bool) {
	ok = o != nil && o.ssh != nil
	if ok {
		value = o.ssh
	}
	return
}

// Admin returns the value of the 'admin' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Temporary administrator credentials generated during the installation
// of the cluster.
func (o *ClusterCredentials) Admin() *AdminCredentials {
	if o == nil {
		return nil
	}
	return o.admin
}

// GetAdmin returns the value of the 'admin' attribute and
// a flag indicating if the attribute has a value.
//
// Temporary administrator credentials generated during the installation
// of the cluster.
func (o *ClusterCredentials) GetAdmin() (value *AdminCredentials, ok bool) {
	ok = o != nil && o.admin != nil
	if ok {
		value = o.admin
	}
	return
}

// ClusterCredentialsList is a list of values of the 'cluster_credentials' type.
type ClusterCredentialsList struct {
	items []*ClusterCredentials
}

// Len returns the length of the list.
func (l *ClusterCredentialsList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Slice returns an slice containing the items of the list. The returned slice is a
// copy of the one used internally, so it can be modified without affecting the
// internal representation.
//
// If you don't need to modify the returned slice consider using the Each or Range
// functions, as they don't need to allocate a new slice.
func (l *ClusterCredentialsList) Slice() []*ClusterCredentials {
	var slice []*ClusterCredentials
	if l == nil {
		slice = make([]*ClusterCredentials, 0)
	} else {
		slice = make([]*ClusterCredentials, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *ClusterCredentialsList) Each(f func(item *ClusterCredentials) bool) {
	if l == nil {
		return
	}
	for _, item := range l.items {
		if !f(item) {
			break
		}
	}
}

// Range runs the given function for each index and item of the list, in order. If
// the function returns false the iteration stops, otherwise it continues till all
// the elements of the list have been processed.
func (l *ClusterCredentialsList) Range(f func(index int, item *ClusterCredentials) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}