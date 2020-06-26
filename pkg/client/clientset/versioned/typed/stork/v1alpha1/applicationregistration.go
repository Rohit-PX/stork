/*
Copyright 2018 Openstorage.org

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

	v1alpha1 "github.com/libopenstorage/stork/pkg/apis/stork/v1alpha1"
	scheme "github.com/libopenstorage/stork/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ApplicationRegistrationsGetter has a method to return a ApplicationRegistrationInterface.
// A group's client should implement this interface.
type ApplicationRegistrationsGetter interface {
	ApplicationRegistrations() ApplicationRegistrationInterface
}

// ApplicationRegistrationInterface has methods to work with ApplicationRegistration resources.
type ApplicationRegistrationInterface interface {
	Create(*v1alpha1.ApplicationRegistration) (*v1alpha1.ApplicationRegistration, error)
	Update(*v1alpha1.ApplicationRegistration) (*v1alpha1.ApplicationRegistration, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ApplicationRegistration, error)
	List(opts v1.ListOptions) (*v1alpha1.ApplicationRegistrationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApplicationRegistration, err error)
	ApplicationRegistrationExpansion
}

// applicationRegistrations implements ApplicationRegistrationInterface
type applicationRegistrations struct {
	client rest.Interface
}

// newApplicationRegistrations returns a ApplicationRegistrations
func newApplicationRegistrations(c *StorkV1alpha1Client) *applicationRegistrations {
	return &applicationRegistrations{
		client: c.RESTClient(),
	}
}

// Get takes name of the applicationRegistration, and returns the corresponding applicationRegistration object, and an error if there is any.
func (c *applicationRegistrations) Get(name string, options v1.GetOptions) (result *v1alpha1.ApplicationRegistration, err error) {
	result = &v1alpha1.ApplicationRegistration{}
	err = c.client.Get().
		Resource("applicationregistrations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ApplicationRegistrations that match those selectors.
func (c *applicationRegistrations) List(opts v1.ListOptions) (result *v1alpha1.ApplicationRegistrationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ApplicationRegistrationList{}
	err = c.client.Get().
		Resource("applicationregistrations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested applicationRegistrations.
func (c *applicationRegistrations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("applicationregistrations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a applicationRegistration and creates it.  Returns the server's representation of the applicationRegistration, and an error, if there is any.
func (c *applicationRegistrations) Create(applicationRegistration *v1alpha1.ApplicationRegistration) (result *v1alpha1.ApplicationRegistration, err error) {
	result = &v1alpha1.ApplicationRegistration{}
	err = c.client.Post().
		Resource("applicationregistrations").
		Body(applicationRegistration).
		Do().
		Into(result)
	return
}

// Update takes the representation of a applicationRegistration and updates it. Returns the server's representation of the applicationRegistration, and an error, if there is any.
func (c *applicationRegistrations) Update(applicationRegistration *v1alpha1.ApplicationRegistration) (result *v1alpha1.ApplicationRegistration, err error) {
	result = &v1alpha1.ApplicationRegistration{}
	err = c.client.Put().
		Resource("applicationregistrations").
		Name(applicationRegistration.Name).
		Body(applicationRegistration).
		Do().
		Into(result)
	return
}

// Delete takes name of the applicationRegistration and deletes it. Returns an error if one occurs.
func (c *applicationRegistrations) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("applicationregistrations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *applicationRegistrations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("applicationregistrations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched applicationRegistration.
func (c *applicationRegistrations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApplicationRegistration, err error) {
	result = &v1alpha1.ApplicationRegistration{}
	err = c.client.Patch(pt).
		Resource("applicationregistrations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
