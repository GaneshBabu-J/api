// Copyright (c) 2020 Tigera, Inc. All rights reserved.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	projectcalico "github.com/tigera/api/pkg/apis/projectcalico"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGlobalReportTypes implements GlobalReportTypeInterface
type FakeGlobalReportTypes struct {
	Fake *FakeProjectcalico
	ns   string
}

var globalreporttypesResource = schema.GroupVersionResource{Group: "projectcalico.org", Version: "", Resource: "globalreporttypes"}

var globalreporttypesKind = schema.GroupVersionKind{Group: "projectcalico.org", Version: "", Kind: "GlobalReportType"}

// Get takes name of the globalReportType, and returns the corresponding globalReportType object, and an error if there is any.
func (c *FakeGlobalReportTypes) Get(name string, options v1.GetOptions) (result *projectcalico.GlobalReportType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(globalreporttypesResource, c.ns, name), &projectcalico.GlobalReportType{})

	if obj == nil {
		return nil, err
	}
	return obj.(*projectcalico.GlobalReportType), err
}

// List takes label and field selectors, and returns the list of GlobalReportTypes that match those selectors.
func (c *FakeGlobalReportTypes) List(opts v1.ListOptions) (result *projectcalico.GlobalReportTypeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(globalreporttypesResource, globalreporttypesKind, c.ns, opts), &projectcalico.GlobalReportTypeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &projectcalico.GlobalReportTypeList{ListMeta: obj.(*projectcalico.GlobalReportTypeList).ListMeta}
	for _, item := range obj.(*projectcalico.GlobalReportTypeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested globalReportTypes.
func (c *FakeGlobalReportTypes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(globalreporttypesResource, c.ns, opts))

}

// Create takes the representation of a globalReportType and creates it.  Returns the server's representation of the globalReportType, and an error, if there is any.
func (c *FakeGlobalReportTypes) Create(globalReportType *projectcalico.GlobalReportType) (result *projectcalico.GlobalReportType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(globalreporttypesResource, c.ns, globalReportType), &projectcalico.GlobalReportType{})

	if obj == nil {
		return nil, err
	}
	return obj.(*projectcalico.GlobalReportType), err
}

// Update takes the representation of a globalReportType and updates it. Returns the server's representation of the globalReportType, and an error, if there is any.
func (c *FakeGlobalReportTypes) Update(globalReportType *projectcalico.GlobalReportType) (result *projectcalico.GlobalReportType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(globalreporttypesResource, c.ns, globalReportType), &projectcalico.GlobalReportType{})

	if obj == nil {
		return nil, err
	}
	return obj.(*projectcalico.GlobalReportType), err
}

// Delete takes name of the globalReportType and deletes it. Returns an error if one occurs.
func (c *FakeGlobalReportTypes) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(globalreporttypesResource, c.ns, name), &projectcalico.GlobalReportType{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGlobalReportTypes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(globalreporttypesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &projectcalico.GlobalReportTypeList{})
	return err
}

// Patch applies the patch and returns the patched globalReportType.
func (c *FakeGlobalReportTypes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *projectcalico.GlobalReportType, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(globalreporttypesResource, c.ns, name, pt, data, subresources...), &projectcalico.GlobalReportType{})

	if obj == nil {
		return nil, err
	}
	return obj.(*projectcalico.GlobalReportType), err
}
