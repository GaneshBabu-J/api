// Copyright (c) 2020 Tigera, Inc. All rights reserved.

// Code generated by client-gen. DO NOT EDIT.

package internalversion

import (
	"time"

	projectcalico "github.com/tigera/api/pkg/apis/projectcalico"
	scheme "github.com/tigera/api/pkg/client/clientset_generated/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// GlobalAlertTemplatesGetter has a method to return a GlobalAlertTemplateInterface.
// A group's client should implement this interface.
type GlobalAlertTemplatesGetter interface {
	GlobalAlertTemplates() GlobalAlertTemplateInterface
}

// GlobalAlertTemplateInterface has methods to work with GlobalAlertTemplate resources.
type GlobalAlertTemplateInterface interface {
	Create(*projectcalico.GlobalAlertTemplate) (*projectcalico.GlobalAlertTemplate, error)
	Update(*projectcalico.GlobalAlertTemplate) (*projectcalico.GlobalAlertTemplate, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*projectcalico.GlobalAlertTemplate, error)
	List(opts v1.ListOptions) (*projectcalico.GlobalAlertTemplateList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *projectcalico.GlobalAlertTemplate, err error)
	GlobalAlertTemplateExpansion
}

// globalAlertTemplates implements GlobalAlertTemplateInterface
type globalAlertTemplates struct {
	client rest.Interface
}

// newGlobalAlertTemplates returns a GlobalAlertTemplates
func newGlobalAlertTemplates(c *ProjectcalicoClient) *globalAlertTemplates {
	return &globalAlertTemplates{
		client: c.RESTClient(),
	}
}

// Get takes name of the globalAlertTemplate, and returns the corresponding globalAlertTemplate object, and an error if there is any.
func (c *globalAlertTemplates) Get(name string, options v1.GetOptions) (result *projectcalico.GlobalAlertTemplate, err error) {
	result = &projectcalico.GlobalAlertTemplate{}
	err = c.client.Get().
		Resource("globalalerttemplates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GlobalAlertTemplates that match those selectors.
func (c *globalAlertTemplates) List(opts v1.ListOptions) (result *projectcalico.GlobalAlertTemplateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &projectcalico.GlobalAlertTemplateList{}
	err = c.client.Get().
		Resource("globalalerttemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested globalAlertTemplates.
func (c *globalAlertTemplates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("globalalerttemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a globalAlertTemplate and creates it.  Returns the server's representation of the globalAlertTemplate, and an error, if there is any.
func (c *globalAlertTemplates) Create(globalAlertTemplate *projectcalico.GlobalAlertTemplate) (result *projectcalico.GlobalAlertTemplate, err error) {
	result = &projectcalico.GlobalAlertTemplate{}
	err = c.client.Post().
		Resource("globalalerttemplates").
		Body(globalAlertTemplate).
		Do().
		Into(result)
	return
}

// Update takes the representation of a globalAlertTemplate and updates it. Returns the server's representation of the globalAlertTemplate, and an error, if there is any.
func (c *globalAlertTemplates) Update(globalAlertTemplate *projectcalico.GlobalAlertTemplate) (result *projectcalico.GlobalAlertTemplate, err error) {
	result = &projectcalico.GlobalAlertTemplate{}
	err = c.client.Put().
		Resource("globalalerttemplates").
		Name(globalAlertTemplate.Name).
		Body(globalAlertTemplate).
		Do().
		Into(result)
	return
}

// Delete takes name of the globalAlertTemplate and deletes it. Returns an error if one occurs.
func (c *globalAlertTemplates) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("globalalerttemplates").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *globalAlertTemplates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("globalalerttemplates").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched globalAlertTemplate.
func (c *globalAlertTemplates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *projectcalico.GlobalAlertTemplate, err error) {
	result = &projectcalico.GlobalAlertTemplate{}
	err = c.client.Patch(pt).
		Resource("globalalerttemplates").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}