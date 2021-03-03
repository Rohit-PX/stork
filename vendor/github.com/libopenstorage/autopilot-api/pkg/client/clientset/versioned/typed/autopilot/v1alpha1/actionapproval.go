/*
Copyright 2019 Openstorage.org

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
	v1alpha1 "github.com/libopenstorage/autopilot-api/pkg/apis/autopilot/v1alpha1"
	scheme "github.com/libopenstorage/autopilot-api/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	"time"
)

// ActionApprovalsGetter has a method to return a ActionApprovalInterface.
// A group's client should implement this interface.
type ActionApprovalsGetter interface {
	ActionApprovals(namespace string) ActionApprovalInterface
}

// ActionApprovalInterface has methods to work with ActionApproval resources.
type ActionApprovalInterface interface {
	Create(*v1alpha1.ActionApproval) (*v1alpha1.ActionApproval, error)
	Update(*v1alpha1.ActionApproval) (*v1alpha1.ActionApproval, error)
	UpdateStatus(*v1alpha1.ActionApproval) (*v1alpha1.ActionApproval, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ActionApproval, error)
	List(opts v1.ListOptions) (*v1alpha1.ActionApprovalList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ActionApproval, err error)
	ActionApprovalExpansion
}

// actionApprovals implements ActionApprovalInterface
type actionApprovals struct {
	client rest.Interface
	ns     string
}

// newActionApprovals returns a ActionApprovals
func newActionApprovals(c *AutopilotV1alpha1Client, namespace string) *actionApprovals {
	return &actionApprovals{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the actionApproval, and returns the corresponding actionApproval object, and an error if there is any.
func (c *actionApprovals) Get(name string, options v1.GetOptions) (result *v1alpha1.ActionApproval, err error) {
	result = &v1alpha1.ActionApproval{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("actionapprovals").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ActionApprovals that match those selectors.
func (c *actionApprovals) List(opts v1.ListOptions) (result *v1alpha1.ActionApprovalList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ActionApprovalList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("actionapprovals").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(context.TODO()).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested actionApprovals.
func (c *actionApprovals) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("actionapprovals").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(context.TODO())
}

// Create takes the representation of a actionApproval and creates it.  Returns the server's representation of the actionApproval, and an error, if there is any.
func (c *actionApprovals) Create(actionApproval *v1alpha1.ActionApproval) (result *v1alpha1.ActionApproval, err error) {
	result = &v1alpha1.ActionApproval{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("actionapprovals").
		Body(actionApproval).
		Do(context.TODO()).
		Into(result)
	return
}

// Update takes the representation of a actionApproval and updates it. Returns the server's representation of the actionApproval, and an error, if there is any.
func (c *actionApprovals) Update(actionApproval *v1alpha1.ActionApproval) (result *v1alpha1.ActionApproval, err error) {
	result = &v1alpha1.ActionApproval{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("actionapprovals").
		Name(actionApproval.Name).
		Body(actionApproval).
		Do(context.TODO()).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *actionApprovals) UpdateStatus(actionApproval *v1alpha1.ActionApproval) (result *v1alpha1.ActionApproval, err error) {
	result = &v1alpha1.ActionApproval{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("actionapprovals").
		Name(actionApproval.Name).
		SubResource("status").
		Body(actionApproval).
		Do(context.TODO()).
		Into(result)
	return
}

// Delete takes name of the actionApproval and deletes it. Returns an error if one occurs.
func (c *actionApprovals) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("actionapprovals").
		Name(name).
		Body(options).
		Do(context.TODO()).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *actionApprovals) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("actionapprovals").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do(context.TODO()).
		Error()
}

// Patch applies the patch and returns the patched actionApproval.
func (c *actionApprovals) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ActionApproval, err error) {
	result = &v1alpha1.ActionApproval{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("actionapprovals").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(context.TODO()).
		Into(result)
	return
}
