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
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

// KuberhealthyJobsGetter has a method to return a KuberhealthyJobInterface.
// A group's client should implement this interface.
type KuberhealthyJobsGetter interface {
	KuberhealthyJobs(namespace string) KuberhealthyJobInterface
}

// KuberhealthyJobInterface has methods to work with KuberhealthyJob resources.
type KuberhealthyJobInterface interface {
	Create(*KuberhealthyJob) (KuberhealthyJob, error)
	Update(*KuberhealthyJob) (KuberhealthyJob, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (KuberhealthyJob, error)
	List(opts metav1.ListOptions) (KuberhealthyJobList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result KuberhealthyJob, err error)
}

// kuberhealthyJobs implements KuberhealthyJobInterface
type kuberhealthyJobs struct {
	client rest.Interface
	ns     string
}

// newKuberhealthyJobs returns a KuberhealthyJobs
func newKuberhealthyJobs(c *KHJobV1Client, namespace string) *kuberhealthyJobs {
	return &kuberhealthyJobs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the kuberhealthyJob, and returns the corresponding kuberhealthyJob object, and an error if there is any.
func (c *kuberhealthyJobs) Get(name string, options metav1.GetOptions) (result KuberhealthyJob, err error) {
	result = KuberhealthyJob{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("khjobs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(&result)
	return
}

// List takes label and field selectors, and returns the list of KuberhealthyJobs that match those selectors.
func (c *kuberhealthyJobs) List(opts metav1.ListOptions) (result KuberhealthyJobList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = KuberhealthyJobList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("khjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(&result)
	return
}

// Watch returns a watch.Interface that watches the requested kuberhealthyJobs.
func (c *kuberhealthyJobs) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("khjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a kuberhealthyJob and creates it.  Returns the server's representation of the kuberhealthyJob, and an error, if there is any.
func (c *kuberhealthyJobs) Create(kuberhealthyJob *KuberhealthyJob) (result KuberhealthyJob, err error) {
	result = KuberhealthyJob{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("khjobs").
		Body(kuberhealthyJob).
		Do().
		Into(&result)
	return
}

// Update takes the representation of a kuberhealthyJob and updates it. Returns the server's representation of the kuberhealthyJob, and an error, if there is any.
func (c *kuberhealthyJobs) Update(kuberhealthyJob *KuberhealthyJob) (result KuberhealthyJob, err error) {
	result = KuberhealthyJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("khjobs").
		Name(kuberhealthyJob.Name).
		Body(kuberhealthyJob).
		Do().
		Into(&result)
	return
}

// Delete takes name of the kuberhealthyJob and deletes it. Returns an error if one occurs.
func (c *kuberhealthyJobs) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("khjobs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *kuberhealthyJobs) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("khjobs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched kuberhealthyJob.
func (c *kuberhealthyJobs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result KuberhealthyJob, err error) {
	result = KuberhealthyJob{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("khjobs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(&result)
	return
}
