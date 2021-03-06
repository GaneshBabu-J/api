// Copyright (c) 2020 Tigera, Inc. All rights reserved.

// Code generated by client-gen. DO NOT EDIT.

package v3

import (
	"time"

	v3 "github.com/tigera/api/pkg/apis/projectcalico/v3"
	scheme "github.com/tigera/api/pkg/client/clientset_generated/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// LicenseKeysGetter has a method to return a LicenseKeyInterface.
// A group's client should implement this interface.
type LicenseKeysGetter interface {
	LicenseKeys(namespace string) LicenseKeyInterface
}

// LicenseKeyInterface has methods to work with LicenseKey resources.
type LicenseKeyInterface interface {
	Create(*v3.LicenseKey) (*v3.LicenseKey, error)
	Update(*v3.LicenseKey) (*v3.LicenseKey, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v3.LicenseKey, error)
	List(opts v1.ListOptions) (*v3.LicenseKeyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v3.LicenseKey, err error)
	LicenseKeyExpansion
}

// licenseKeys implements LicenseKeyInterface
type licenseKeys struct {
	client rest.Interface
	ns     string
}

// newLicenseKeys returns a LicenseKeys
func newLicenseKeys(c *ProjectcalicoV3Client, namespace string) *licenseKeys {
	return &licenseKeys{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the licenseKey, and returns the corresponding licenseKey object, and an error if there is any.
func (c *licenseKeys) Get(name string, options v1.GetOptions) (result *v3.LicenseKey, err error) {
	result = &v3.LicenseKey{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("licensekeys").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of LicenseKeys that match those selectors.
func (c *licenseKeys) List(opts v1.ListOptions) (result *v3.LicenseKeyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v3.LicenseKeyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("licensekeys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested licenseKeys.
func (c *licenseKeys) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("licensekeys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a licenseKey and creates it.  Returns the server's representation of the licenseKey, and an error, if there is any.
func (c *licenseKeys) Create(licenseKey *v3.LicenseKey) (result *v3.LicenseKey, err error) {
	result = &v3.LicenseKey{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("licensekeys").
		Body(licenseKey).
		Do().
		Into(result)
	return
}

// Update takes the representation of a licenseKey and updates it. Returns the server's representation of the licenseKey, and an error, if there is any.
func (c *licenseKeys) Update(licenseKey *v3.LicenseKey) (result *v3.LicenseKey, err error) {
	result = &v3.LicenseKey{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("licensekeys").
		Name(licenseKey.Name).
		Body(licenseKey).
		Do().
		Into(result)
	return
}

// Delete takes name of the licenseKey and deletes it. Returns an error if one occurs.
func (c *licenseKeys) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("licensekeys").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *licenseKeys) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("licensekeys").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched licenseKey.
func (c *licenseKeys) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v3.LicenseKey, err error) {
	result = &v3.LicenseKey{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("licensekeys").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
