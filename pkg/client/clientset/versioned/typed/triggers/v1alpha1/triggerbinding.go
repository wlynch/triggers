/*
Copyright 2018 The Knative Authors

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
	"time"

	v1alpha1 "github.com/tektoncd/triggers/pkg/apis/triggers/v1alpha1"
	scheme "github.com/tektoncd/triggers/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TriggerBindingsGetter has a method to return a TriggerBindingInterface.
// A group's client should implement this interface.
type TriggerBindingsGetter interface {
	TriggerBindings(namespace string) TriggerBindingInterface
}

// TriggerBindingInterface has methods to work with TriggerBinding resources.
type TriggerBindingInterface interface {
	Create(*v1alpha1.TriggerBinding) (*v1alpha1.TriggerBinding, error)
	Update(*v1alpha1.TriggerBinding) (*v1alpha1.TriggerBinding, error)
	UpdateStatus(*v1alpha1.TriggerBinding) (*v1alpha1.TriggerBinding, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.TriggerBinding, error)
	List(opts v1.ListOptions) (*v1alpha1.TriggerBindingList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.TriggerBinding, err error)
	TriggerBindingExpansion
}

// triggerBindings implements TriggerBindingInterface
type triggerBindings struct {
	client rest.Interface
	ns     string
}

// newTriggerBindings returns a TriggerBindings
func newTriggerBindings(c *TriggersV1alpha1Client, namespace string) *triggerBindings {
	return &triggerBindings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the triggerBinding, and returns the corresponding triggerBinding object, and an error if there is any.
func (c *triggerBindings) Get(name string, options v1.GetOptions) (result *v1alpha1.TriggerBinding, err error) {
	result = &v1alpha1.TriggerBinding{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("triggerbindings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TriggerBindings that match those selectors.
func (c *triggerBindings) List(opts v1.ListOptions) (result *v1alpha1.TriggerBindingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.TriggerBindingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("triggerbindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested triggerBindings.
func (c *triggerBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("triggerbindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a triggerBinding and creates it.  Returns the server's representation of the triggerBinding, and an error, if there is any.
func (c *triggerBindings) Create(triggerBinding *v1alpha1.TriggerBinding) (result *v1alpha1.TriggerBinding, err error) {
	result = &v1alpha1.TriggerBinding{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("triggerbindings").
		Body(triggerBinding).
		Do().
		Into(result)
	return
}

// Update takes the representation of a triggerBinding and updates it. Returns the server's representation of the triggerBinding, and an error, if there is any.
func (c *triggerBindings) Update(triggerBinding *v1alpha1.TriggerBinding) (result *v1alpha1.TriggerBinding, err error) {
	result = &v1alpha1.TriggerBinding{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("triggerbindings").
		Name(triggerBinding.Name).
		Body(triggerBinding).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *triggerBindings) UpdateStatus(triggerBinding *v1alpha1.TriggerBinding) (result *v1alpha1.TriggerBinding, err error) {
	result = &v1alpha1.TriggerBinding{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("triggerbindings").
		Name(triggerBinding.Name).
		SubResource("status").
		Body(triggerBinding).
		Do().
		Into(result)
	return
}

// Delete takes name of the triggerBinding and deletes it. Returns an error if one occurs.
func (c *triggerBindings) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("triggerbindings").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *triggerBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("triggerbindings").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched triggerBinding.
func (c *triggerBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.TriggerBinding, err error) {
	result = &v1alpha1.TriggerBinding{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("triggerbindings").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
