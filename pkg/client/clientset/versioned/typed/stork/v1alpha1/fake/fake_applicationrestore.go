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

package fake

import (
	v1alpha1 "github.com/libopenstorage/stork/pkg/apis/stork/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeApplicationRestores implements ApplicationRestoreInterface
type FakeApplicationRestores struct {
	Fake *FakeStorkV1alpha1
	ns   string
}

var applicationrestoresResource = schema.GroupVersionResource{Group: "stork.libopenstorage.org", Version: "v1alpha1", Resource: "applicationrestores"}

var applicationrestoresKind = schema.GroupVersionKind{Group: "stork.libopenstorage.org", Version: "v1alpha1", Kind: "ApplicationRestore"}

// Get takes name of the applicationRestore, and returns the corresponding applicationRestore object, and an error if there is any.
func (c *FakeApplicationRestores) Get(name string, options v1.GetOptions) (result *v1alpha1.ApplicationRestore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(applicationrestoresResource, c.ns, name), &v1alpha1.ApplicationRestore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApplicationRestore), err
}

// List takes label and field selectors, and returns the list of ApplicationRestores that match those selectors.
func (c *FakeApplicationRestores) List(opts v1.ListOptions) (result *v1alpha1.ApplicationRestoreList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(applicationrestoresResource, applicationrestoresKind, c.ns, opts), &v1alpha1.ApplicationRestoreList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ApplicationRestoreList{ListMeta: obj.(*v1alpha1.ApplicationRestoreList).ListMeta}
	for _, item := range obj.(*v1alpha1.ApplicationRestoreList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested applicationRestores.
func (c *FakeApplicationRestores) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(applicationrestoresResource, c.ns, opts))

}

// Create takes the representation of a applicationRestore and creates it.  Returns the server's representation of the applicationRestore, and an error, if there is any.
func (c *FakeApplicationRestores) Create(applicationRestore *v1alpha1.ApplicationRestore) (result *v1alpha1.ApplicationRestore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(applicationrestoresResource, c.ns, applicationRestore), &v1alpha1.ApplicationRestore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApplicationRestore), err
}

// Update takes the representation of a applicationRestore and updates it. Returns the server's representation of the applicationRestore, and an error, if there is any.
func (c *FakeApplicationRestores) Update(applicationRestore *v1alpha1.ApplicationRestore) (result *v1alpha1.ApplicationRestore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(applicationrestoresResource, c.ns, applicationRestore), &v1alpha1.ApplicationRestore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApplicationRestore), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApplicationRestores) UpdateStatus(applicationRestore *v1alpha1.ApplicationRestore) (*v1alpha1.ApplicationRestore, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(applicationrestoresResource, "status", c.ns, applicationRestore), &v1alpha1.ApplicationRestore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApplicationRestore), err
}

// Delete takes name of the applicationRestore and deletes it. Returns an error if one occurs.
func (c *FakeApplicationRestores) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(applicationrestoresResource, c.ns, name), &v1alpha1.ApplicationRestore{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApplicationRestores) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(applicationrestoresResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApplicationRestoreList{})
	return err
}

// Patch applies the patch and returns the patched applicationRestore.
func (c *FakeApplicationRestores) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApplicationRestore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(applicationrestoresResource, c.ns, name, pt, data, subresources...), &v1alpha1.ApplicationRestore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApplicationRestore), err
}
