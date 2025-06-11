// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeinterconnectattachmentgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/computeinterconnectattachmentgroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeInterconnectAttachmentGroupLogicalStructureRegionsMetrosFacilitiesList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) ComputeInterconnectAttachmentGroupLogicalStructureRegionsMetrosFacilitiesOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeInterconnectAttachmentGroupLogicalStructureRegionsMetrosFacilitiesList
type jsiiProxy_ComputeInterconnectAttachmentGroupLogicalStructureRegionsMetrosFacilitiesList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_ComputeInterconnectAttachmentGroupLogicalStructureRegionsMetrosFacilitiesList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachmentGroupLogicalStructureRegionsMetrosFacilitiesList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachmentGroupLogicalStructureRegionsMetrosFacilitiesList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachmentGroupLogicalStructureRegionsMetrosFacilitiesList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachmentGroupLogicalStructureRegionsMetrosFacilitiesList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewComputeInterconnectAttachmentGroupLogicalStructureRegionsMetrosFacilitiesList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ComputeInterconnectAttachmentGroupLogicalStructureRegionsMetrosFacilitiesList {
	_init_.Initialize()

	if err := validateNewComputeInterconnectAttachmentGroupLogicalStructureRegionsMetrosFacilitiesListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeInterconnectAttachmentGroupLogicalStructureRegionsMetrosFacilitiesList{}

	_jsii_.Create(
		"@cdktf/provider-google.computeInterconnectAttachmentGroup.ComputeInterconnectAttachmentGroupLogicalStructureRegionsMetrosFacilitiesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewComputeInterconnectAttachmentGroupLogicalStructureRegionsMetrosFacilitiesList_Override(c ComputeInterconnectAttachmentGroupLogicalStructureRegionsMetrosFacilitiesList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeInterconnectAttachmentGroup.ComputeInterconnectAttachmentGroupLogicalStructureRegionsMetrosFacilitiesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		c,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachmentGroupLogicalStructureRegionsMetrosFacilitiesList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachmentGroupLogicalStructureRegionsMetrosFacilitiesList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachmentGroupLogicalStructureRegionsMetrosFacilitiesList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (c *jsiiProxy_ComputeInterconnectAttachmentGroupLogicalStructureRegionsMetrosFacilitiesList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := c.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		c,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInterconnectAttachmentGroupLogicalStructureRegionsMetrosFacilitiesList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInterconnectAttachmentGroupLogicalStructureRegionsMetrosFacilitiesList) Get(index *float64) ComputeInterconnectAttachmentGroupLogicalStructureRegionsMetrosFacilitiesOutputReference {
	if err := c.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns ComputeInterconnectAttachmentGroupLogicalStructureRegionsMetrosFacilitiesOutputReference

	_jsii_.Invoke(
		c,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInterconnectAttachmentGroupLogicalStructureRegionsMetrosFacilitiesList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInterconnectAttachmentGroupLogicalStructureRegionsMetrosFacilitiesList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

