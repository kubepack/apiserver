/*
Copyright 2018 The Kubepack Authors.

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
package fake

import (
	v1alpha1 "github.com/kubepack/packserver/apis/apps/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeManifests implements ManifestInterface
type FakeManifests struct {
	Fake *FakeAppsV1alpha1
	ns   string
}

var manifestsResource = schema.GroupVersionResource{Group: "apps.kubepack.com", Version: "v1alpha1", Resource: "manifests"}

var manifestsKind = schema.GroupVersionKind{Group: "apps.kubepack.com", Version: "v1alpha1", Kind: "Manifest"}

// Get takes name of the manifest, and returns the corresponding manifest object, and an error if there is any.
func (c *FakeManifests) Get(name string, options v1.GetOptions) (result *v1alpha1.Manifest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(manifestsResource, c.ns, name), &v1alpha1.Manifest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Manifest), err
}

// List takes label and field selectors, and returns the list of Manifests that match those selectors.
func (c *FakeManifests) List(opts v1.ListOptions) (result *v1alpha1.ManifestList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(manifestsResource, manifestsKind, c.ns, opts), &v1alpha1.ManifestList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ManifestList{}
	for _, item := range obj.(*v1alpha1.ManifestList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested manifests.
func (c *FakeManifests) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(manifestsResource, c.ns, opts))

}

// Create takes the representation of a manifest and creates it.  Returns the server's representation of the manifest, and an error, if there is any.
func (c *FakeManifests) Create(manifest *v1alpha1.Manifest) (result *v1alpha1.Manifest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(manifestsResource, c.ns, manifest), &v1alpha1.Manifest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Manifest), err
}

// Update takes the representation of a manifest and updates it. Returns the server's representation of the manifest, and an error, if there is any.
func (c *FakeManifests) Update(manifest *v1alpha1.Manifest) (result *v1alpha1.Manifest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(manifestsResource, c.ns, manifest), &v1alpha1.Manifest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Manifest), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeManifests) UpdateStatus(manifest *v1alpha1.Manifest) (*v1alpha1.Manifest, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(manifestsResource, "status", c.ns, manifest), &v1alpha1.Manifest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Manifest), err
}

// Delete takes name of the manifest and deletes it. Returns an error if one occurs.
func (c *FakeManifests) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(manifestsResource, c.ns, name), &v1alpha1.Manifest{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeManifests) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(manifestsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ManifestList{})
	return err
}

// Patch applies the patch and returns the patched manifest.
func (c *FakeManifests) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Manifest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(manifestsResource, c.ns, name, data, subresources...), &v1alpha1.Manifest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Manifest), err
}
