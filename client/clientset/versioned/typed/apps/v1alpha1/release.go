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
package v1alpha1

import (
	v1alpha1 "github.com/kubepack/packserver/apis/apps/v1alpha1"
	scheme "github.com/kubepack/packserver/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ReleasesGetter has a method to return a ReleaseInterface.
// A group's client should implement this interface.
type ReleasesGetter interface {
	Releases(namespace string) ReleaseInterface
}

// ReleaseInterface has methods to work with Release resources.
type ReleaseInterface interface {
	Create(*v1alpha1.Release) (*v1alpha1.Release, error)
	Update(*v1alpha1.Release) (*v1alpha1.Release, error)
	UpdateStatus(*v1alpha1.Release) (*v1alpha1.Release, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Release, error)
	List(opts v1.ListOptions) (*v1alpha1.ReleaseList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Release, err error)
	ReleaseExpansion
}

// releases implements ReleaseInterface
type releases struct {
	client rest.Interface
	ns     string
}

// newReleases returns a Releases
func newReleases(c *AppsV1alpha1Client, namespace string) *releases {
	return &releases{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the release, and returns the corresponding release object, and an error if there is any.
func (c *releases) Get(name string, options v1.GetOptions) (result *v1alpha1.Release, err error) {
	result = &v1alpha1.Release{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("releases").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Releases that match those selectors.
func (c *releases) List(opts v1.ListOptions) (result *v1alpha1.ReleaseList, err error) {
	result = &v1alpha1.ReleaseList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("releases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested releases.
func (c *releases) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("releases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a release and creates it.  Returns the server's representation of the release, and an error, if there is any.
func (c *releases) Create(release *v1alpha1.Release) (result *v1alpha1.Release, err error) {
	result = &v1alpha1.Release{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("releases").
		Body(release).
		Do().
		Into(result)
	return
}

// Update takes the representation of a release and updates it. Returns the server's representation of the release, and an error, if there is any.
func (c *releases) Update(release *v1alpha1.Release) (result *v1alpha1.Release, err error) {
	result = &v1alpha1.Release{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("releases").
		Name(release.Name).
		Body(release).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *releases) UpdateStatus(release *v1alpha1.Release) (result *v1alpha1.Release, err error) {
	result = &v1alpha1.Release{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("releases").
		Name(release.Name).
		SubResource("status").
		Body(release).
		Do().
		Into(result)
	return
}

// Delete takes name of the release and deletes it. Returns an error if one occurs.
func (c *releases) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("releases").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *releases) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("releases").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched release.
func (c *releases) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Release, err error) {
	result = &v1alpha1.Release{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("releases").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
