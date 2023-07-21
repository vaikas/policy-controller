// Copyright 2022 The Sigstore Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "github.com/sigstore/policy-controller/pkg/apis/policy/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterImagePolicies implements ClusterImagePolicyInterface
type FakeClusterImagePolicies struct {
	Fake *FakePolicyV1beta1
}

var clusterimagepoliciesResource = v1beta1.SchemeGroupVersion.WithResource("clusterimagepolicies")

var clusterimagepoliciesKind = v1beta1.SchemeGroupVersion.WithKind("ClusterImagePolicy")

// Get takes name of the clusterImagePolicy, and returns the corresponding clusterImagePolicy object, and an error if there is any.
func (c *FakeClusterImagePolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ClusterImagePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterimagepoliciesResource, name), &v1beta1.ClusterImagePolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ClusterImagePolicy), err
}

// List takes label and field selectors, and returns the list of ClusterImagePolicies that match those selectors.
func (c *FakeClusterImagePolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ClusterImagePolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterimagepoliciesResource, clusterimagepoliciesKind, opts), &v1beta1.ClusterImagePolicyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ClusterImagePolicyList{ListMeta: obj.(*v1beta1.ClusterImagePolicyList).ListMeta}
	for _, item := range obj.(*v1beta1.ClusterImagePolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterImagePolicies.
func (c *FakeClusterImagePolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterimagepoliciesResource, opts))
}

// Create takes the representation of a clusterImagePolicy and creates it.  Returns the server's representation of the clusterImagePolicy, and an error, if there is any.
func (c *FakeClusterImagePolicies) Create(ctx context.Context, clusterImagePolicy *v1beta1.ClusterImagePolicy, opts v1.CreateOptions) (result *v1beta1.ClusterImagePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterimagepoliciesResource, clusterImagePolicy), &v1beta1.ClusterImagePolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ClusterImagePolicy), err
}

// Update takes the representation of a clusterImagePolicy and updates it. Returns the server's representation of the clusterImagePolicy, and an error, if there is any.
func (c *FakeClusterImagePolicies) Update(ctx context.Context, clusterImagePolicy *v1beta1.ClusterImagePolicy, opts v1.UpdateOptions) (result *v1beta1.ClusterImagePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterimagepoliciesResource, clusterImagePolicy), &v1beta1.ClusterImagePolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ClusterImagePolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterImagePolicies) UpdateStatus(ctx context.Context, clusterImagePolicy *v1beta1.ClusterImagePolicy, opts v1.UpdateOptions) (*v1beta1.ClusterImagePolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clusterimagepoliciesResource, "status", clusterImagePolicy), &v1beta1.ClusterImagePolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ClusterImagePolicy), err
}

// Delete takes name of the clusterImagePolicy and deletes it. Returns an error if one occurs.
func (c *FakeClusterImagePolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(clusterimagepoliciesResource, name, opts), &v1beta1.ClusterImagePolicy{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterImagePolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterimagepoliciesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.ClusterImagePolicyList{})
	return err
}

// Patch applies the patch and returns the patched clusterImagePolicy.
func (c *FakeClusterImagePolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ClusterImagePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterimagepoliciesResource, name, pt, data, subresources...), &v1beta1.ClusterImagePolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ClusterImagePolicy), err
}
