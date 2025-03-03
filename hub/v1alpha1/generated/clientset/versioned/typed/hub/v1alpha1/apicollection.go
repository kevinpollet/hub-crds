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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/traefik/hub-crds/hub/v1alpha1"
	scheme "github.com/traefik/hub-crds/hub/v1alpha1/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// APICollectionsGetter has a method to return a APICollectionInterface.
// A group's client should implement this interface.
type APICollectionsGetter interface {
	APICollections() APICollectionInterface
}

// APICollectionInterface has methods to work with APICollection resources.
type APICollectionInterface interface {
	Create(ctx context.Context, aPICollection *v1alpha1.APICollection, opts v1.CreateOptions) (*v1alpha1.APICollection, error)
	Update(ctx context.Context, aPICollection *v1alpha1.APICollection, opts v1.UpdateOptions) (*v1alpha1.APICollection, error)
	UpdateStatus(ctx context.Context, aPICollection *v1alpha1.APICollection, opts v1.UpdateOptions) (*v1alpha1.APICollection, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.APICollection, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.APICollectionList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.APICollection, err error)
	APICollectionExpansion
}

// aPICollections implements APICollectionInterface
type aPICollections struct {
	client rest.Interface
}

// newAPICollections returns a APICollections
func newAPICollections(c *HubV1alpha1Client) *aPICollections {
	return &aPICollections{
		client: c.RESTClient(),
	}
}

// Get takes name of the aPICollection, and returns the corresponding aPICollection object, and an error if there is any.
func (c *aPICollections) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.APICollection, err error) {
	result = &v1alpha1.APICollection{}
	err = c.client.Get().
		Resource("apicollections").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of APICollections that match those selectors.
func (c *aPICollections) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.APICollectionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.APICollectionList{}
	err = c.client.Get().
		Resource("apicollections").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested aPICollections.
func (c *aPICollections) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("apicollections").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a aPICollection and creates it.  Returns the server's representation of the aPICollection, and an error, if there is any.
func (c *aPICollections) Create(ctx context.Context, aPICollection *v1alpha1.APICollection, opts v1.CreateOptions) (result *v1alpha1.APICollection, err error) {
	result = &v1alpha1.APICollection{}
	err = c.client.Post().
		Resource("apicollections").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aPICollection).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a aPICollection and updates it. Returns the server's representation of the aPICollection, and an error, if there is any.
func (c *aPICollections) Update(ctx context.Context, aPICollection *v1alpha1.APICollection, opts v1.UpdateOptions) (result *v1alpha1.APICollection, err error) {
	result = &v1alpha1.APICollection{}
	err = c.client.Put().
		Resource("apicollections").
		Name(aPICollection.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aPICollection).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *aPICollections) UpdateStatus(ctx context.Context, aPICollection *v1alpha1.APICollection, opts v1.UpdateOptions) (result *v1alpha1.APICollection, err error) {
	result = &v1alpha1.APICollection{}
	err = c.client.Put().
		Resource("apicollections").
		Name(aPICollection.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aPICollection).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the aPICollection and deletes it. Returns an error if one occurs.
func (c *aPICollections) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("apicollections").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *aPICollections) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("apicollections").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched aPICollection.
func (c *aPICollections) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.APICollection, err error) {
	result = &v1alpha1.APICollection{}
	err = c.client.Patch(pt).
		Resource("apicollections").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
