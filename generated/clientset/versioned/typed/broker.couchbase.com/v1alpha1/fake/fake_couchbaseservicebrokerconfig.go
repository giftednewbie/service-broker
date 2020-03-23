/* Copyright (C) Couchbase, Inc 2020 - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/couchbase/service-broker/pkg/apis/broker.couchbase.com/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCouchbaseServiceBrokerConfigs implements CouchbaseServiceBrokerConfigInterface
type FakeCouchbaseServiceBrokerConfigs struct {
	Fake *FakeBrokerV1alpha1
	ns   string
}

var couchbaseservicebrokerconfigsResource = schema.GroupVersionResource{Group: "broker.couchbase.com", Version: "v1alpha1", Resource: "couchbaseservicebrokerconfigs"}

var couchbaseservicebrokerconfigsKind = schema.GroupVersionKind{Group: "broker.couchbase.com", Version: "v1alpha1", Kind: "CouchbaseServiceBrokerConfig"}

// Get takes name of the couchbaseServiceBrokerConfig, and returns the corresponding couchbaseServiceBrokerConfig object, and an error if there is any.
func (c *FakeCouchbaseServiceBrokerConfigs) Get(name string, options v1.GetOptions) (result *v1alpha1.CouchbaseServiceBrokerConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(couchbaseservicebrokerconfigsResource, c.ns, name), &v1alpha1.CouchbaseServiceBrokerConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CouchbaseServiceBrokerConfig), err
}

// List takes label and field selectors, and returns the list of CouchbaseServiceBrokerConfigs that match those selectors.
func (c *FakeCouchbaseServiceBrokerConfigs) List(opts v1.ListOptions) (result *v1alpha1.CouchbaseServiceBrokerConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(couchbaseservicebrokerconfigsResource, couchbaseservicebrokerconfigsKind, c.ns, opts), &v1alpha1.CouchbaseServiceBrokerConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CouchbaseServiceBrokerConfigList{ListMeta: obj.(*v1alpha1.CouchbaseServiceBrokerConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.CouchbaseServiceBrokerConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested couchbaseServiceBrokerConfigs.
func (c *FakeCouchbaseServiceBrokerConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(couchbaseservicebrokerconfigsResource, c.ns, opts))

}

// Create takes the representation of a couchbaseServiceBrokerConfig and creates it.  Returns the server's representation of the couchbaseServiceBrokerConfig, and an error, if there is any.
func (c *FakeCouchbaseServiceBrokerConfigs) Create(couchbaseServiceBrokerConfig *v1alpha1.CouchbaseServiceBrokerConfig) (result *v1alpha1.CouchbaseServiceBrokerConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(couchbaseservicebrokerconfigsResource, c.ns, couchbaseServiceBrokerConfig), &v1alpha1.CouchbaseServiceBrokerConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CouchbaseServiceBrokerConfig), err
}

// Update takes the representation of a couchbaseServiceBrokerConfig and updates it. Returns the server's representation of the couchbaseServiceBrokerConfig, and an error, if there is any.
func (c *FakeCouchbaseServiceBrokerConfigs) Update(couchbaseServiceBrokerConfig *v1alpha1.CouchbaseServiceBrokerConfig) (result *v1alpha1.CouchbaseServiceBrokerConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(couchbaseservicebrokerconfigsResource, c.ns, couchbaseServiceBrokerConfig), &v1alpha1.CouchbaseServiceBrokerConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CouchbaseServiceBrokerConfig), err
}

// Delete takes name of the couchbaseServiceBrokerConfig and deletes it. Returns an error if one occurs.
func (c *FakeCouchbaseServiceBrokerConfigs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(couchbaseservicebrokerconfigsResource, c.ns, name), &v1alpha1.CouchbaseServiceBrokerConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCouchbaseServiceBrokerConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(couchbaseservicebrokerconfigsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.CouchbaseServiceBrokerConfigList{})
	return err
}

// Patch applies the patch and returns the patched couchbaseServiceBrokerConfig.
func (c *FakeCouchbaseServiceBrokerConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CouchbaseServiceBrokerConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(couchbaseservicebrokerconfigsResource, c.ns, name, pt, data, subresources...), &v1alpha1.CouchbaseServiceBrokerConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CouchbaseServiceBrokerConfig), err
}
