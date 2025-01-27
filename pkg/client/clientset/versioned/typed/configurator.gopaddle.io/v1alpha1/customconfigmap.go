/*
Copyright 2021.

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

	v1alpha1 "github.com/gopaddle-io/configurator/apis/configurator.gopaddle.io/v1alpha1"
	scheme "github.com/gopaddle-io/configurator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CustomConfigMapsGetter has a method to return a CustomConfigMapInterface.
// A group's client should implement this interface.
type CustomConfigMapsGetter interface {
	CustomConfigMaps(namespace string) CustomConfigMapInterface
}

// CustomConfigMapInterface has methods to work with CustomConfigMap resources.
type CustomConfigMapInterface interface {
	Create(ctx context.Context, customConfigMap *v1alpha1.CustomConfigMap, opts v1.CreateOptions) (*v1alpha1.CustomConfigMap, error)
	Update(ctx context.Context, customConfigMap *v1alpha1.CustomConfigMap, opts v1.UpdateOptions) (*v1alpha1.CustomConfigMap, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.CustomConfigMap, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.CustomConfigMapList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CustomConfigMap, err error)
	CustomConfigMapExpansion
}

// customConfigMaps implements CustomConfigMapInterface
type customConfigMaps struct {
	client rest.Interface
	ns     string
}

// newCustomConfigMaps returns a CustomConfigMaps
func newCustomConfigMaps(c *ConfiguratorV1alpha1Client, namespace string) *customConfigMaps {
	return &customConfigMaps{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the customConfigMap, and returns the corresponding customConfigMap object, and an error if there is any.
func (c *customConfigMaps) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CustomConfigMap, err error) {
	result = &v1alpha1.CustomConfigMap{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("customconfigmaps").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CustomConfigMaps that match those selectors.
func (c *customConfigMaps) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CustomConfigMapList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CustomConfigMapList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("customconfigmaps").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested customConfigMaps.
func (c *customConfigMaps) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("customconfigmaps").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a customConfigMap and creates it.  Returns the server's representation of the customConfigMap, and an error, if there is any.
func (c *customConfigMaps) Create(ctx context.Context, customConfigMap *v1alpha1.CustomConfigMap, opts v1.CreateOptions) (result *v1alpha1.CustomConfigMap, err error) {
	result = &v1alpha1.CustomConfigMap{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("customconfigmaps").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(customConfigMap).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a customConfigMap and updates it. Returns the server's representation of the customConfigMap, and an error, if there is any.
func (c *customConfigMaps) Update(ctx context.Context, customConfigMap *v1alpha1.CustomConfigMap, opts v1.UpdateOptions) (result *v1alpha1.CustomConfigMap, err error) {
	result = &v1alpha1.CustomConfigMap{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("customconfigmaps").
		Name(customConfigMap.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(customConfigMap).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the customConfigMap and deletes it. Returns an error if one occurs.
func (c *customConfigMaps) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("customconfigmaps").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *customConfigMaps) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("customconfigmaps").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched customConfigMap.
func (c *customConfigMaps) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CustomConfigMap, err error) {
	result = &v1alpha1.CustomConfigMap{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("customconfigmaps").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
