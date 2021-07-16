/*
 Copyright 2020 The Knative Authors

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

package v1

import (
	"context"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

// KuberhealthyStatesGetter has a method to return a KuberhealthyStateInterface.
// A group's client should implement this interface.
type KuberhealthyStatesGetter interface {
	KuberhealthyStates(namespace string) KuberhealthyStateInterface
}

// KuberhealthyStateInterface has methods to work with KuberhealthyState resources.
type KuberhealthyStateInterface interface {
	Create(*KuberhealthyState) (KuberhealthyState, error)
	Update(*KuberhealthyState) (KuberhealthyState, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (KuberhealthyState, error)
	List(opts metav1.ListOptions) (KuberhealthyStateList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result KuberhealthyState, err error)
}

// kuberhealthyStates implements KuberhealthyStateInterface
type kuberhealthyStates struct {
	client rest.Interface
	ns     string
}

// newKuberhealthyStates returns a KuberhealthyStates
func newKuberhealthyStates(c *KHStateV1Client, namespace string) *kuberhealthyStates {
	return &kuberhealthyStates{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the kuberhealthyState, and returns the corresponding kuberhealthyState object, and an error if there is any.
func (c *kuberhealthyStates) Get(name string, options metav1.GetOptions) (result KuberhealthyState, err error) {
	result = KuberhealthyState{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("khstates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(&result)
	return
}

// List takes label and field selectors, and returns the list of KuberhealthyStates that match those selectors.
func (c *kuberhealthyStates) List(opts metav1.ListOptions) (result KuberhealthyStateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = KuberhealthyStateList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("khstates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(context.TODO()).
		Into(&result)
	return
}

// Watch returns a watch.Interface that watches the requested kuberhealthyStates.
func (c *kuberhealthyStates) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("khstates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(context.TODO())
}

// Create takes the representation of a kuberhealthyState and creates it.  Returns the server's representation of the kuberhealthyState, and an error, if there is any.
func (c *kuberhealthyStates) Create(kuberhealthyState *KuberhealthyState) (result KuberhealthyState, err error) {
	result = KuberhealthyState{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("khstates").
		Body(kuberhealthyState).
		Do(context.TODO()).
		Into(&result)
	return
}

// Update takes the representation of a kuberhealthyState and updates it. Returns the server's representation of the kuberhealthyState, and an error, if there is any.
func (c *kuberhealthyStates) Update(kuberhealthyState *KuberhealthyState) (result KuberhealthyState, err error) {
	result = KuberhealthyState{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("khstates").
		Name(kuberhealthyState.Name).
		Body(kuberhealthyState).
		Do(context.TODO()).
		Into(&result)
	return
}

// Delete takes name of the kuberhealthyState and deletes it. Returns an error if one occurs.
func (c *kuberhealthyStates) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("khstates").
		Name(name).
		Body(options).
		Do(context.TODO()).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *kuberhealthyStates) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("khstates").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do(context.TODO()).
		Error()
}

// Patch applies the patch and returns the patched kuberhealthyState.
func (c *kuberhealthyStates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result KuberhealthyState, err error) {
	result = KuberhealthyState{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("khstates").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(context.TODO()).
		Into(&result)
	return
}
