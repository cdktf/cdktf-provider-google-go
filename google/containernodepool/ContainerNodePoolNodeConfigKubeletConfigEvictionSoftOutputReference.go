// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containernodepool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/containernodepool/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	ImagefsAvailable() *string
	SetImagefsAvailable(val *string)
	ImagefsAvailableInput() *string
	ImagefsInodesFree() *string
	SetImagefsInodesFree(val *string)
	ImagefsInodesFreeInput() *string
	InternalValue() *ContainerNodePoolNodeConfigKubeletConfigEvictionSoft
	SetInternalValue(val *ContainerNodePoolNodeConfigKubeletConfigEvictionSoft)
	MemoryAvailable() *string
	SetMemoryAvailable(val *string)
	MemoryAvailableInput() *string
	NodefsAvailable() *string
	SetNodefsAvailable(val *string)
	NodefsAvailableInput() *string
	NodefsInodesFree() *string
	SetNodefsInodesFree(val *string)
	NodefsInodesFreeInput() *string
	PidAvailable() *string
	SetPidAvailable(val *string)
	PidAvailableInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	ResetImagefsAvailable()
	ResetImagefsInodesFree()
	ResetMemoryAvailable()
	ResetNodefsAvailable()
	ResetNodefsInodesFree()
	ResetPidAvailable()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference
type jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) ImagefsAvailable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagefsAvailable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) ImagefsAvailableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagefsAvailableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) ImagefsInodesFree() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagefsInodesFree",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) ImagefsInodesFreeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imagefsInodesFreeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) InternalValue() *ContainerNodePoolNodeConfigKubeletConfigEvictionSoft {
	var returns *ContainerNodePoolNodeConfigKubeletConfigEvictionSoft
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) MemoryAvailable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memoryAvailable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) MemoryAvailableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"memoryAvailableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) NodefsAvailable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodefsAvailable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) NodefsAvailableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodefsAvailableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) NodefsInodesFree() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodefsInodesFree",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) NodefsInodesFreeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodefsInodesFreeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) PidAvailable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pidAvailable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) PidAvailableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pidAvailableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference {
	_init_.Initialize()

	if err := validateNewContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.containerNodePool.ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference_Override(c ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.containerNodePool.ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference)SetImagefsAvailable(val *string) {
	if err := j.validateSetImagefsAvailableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imagefsAvailable",
		val,
	)
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference)SetImagefsInodesFree(val *string) {
	if err := j.validateSetImagefsInodesFreeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imagefsInodesFree",
		val,
	)
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference)SetInternalValue(val *ContainerNodePoolNodeConfigKubeletConfigEvictionSoft) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference)SetMemoryAvailable(val *string) {
	if err := j.validateSetMemoryAvailableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryAvailable",
		val,
	)
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference)SetNodefsAvailable(val *string) {
	if err := j.validateSetNodefsAvailableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodefsAvailable",
		val,
	)
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference)SetNodefsInodesFree(val *string) {
	if err := j.validateSetNodefsInodesFreeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodefsInodesFree",
		val,
	)
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference)SetPidAvailable(val *string) {
	if err := j.validateSetPidAvailableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pidAvailable",
		val,
	)
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) ResetImagefsAvailable() {
	_jsii_.InvokeVoid(
		c,
		"resetImagefsAvailable",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) ResetImagefsInodesFree() {
	_jsii_.InvokeVoid(
		c,
		"resetImagefsInodesFree",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) ResetMemoryAvailable() {
	_jsii_.InvokeVoid(
		c,
		"resetMemoryAvailable",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) ResetNodefsAvailable() {
	_jsii_.InvokeVoid(
		c,
		"resetNodefsAvailable",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) ResetNodefsInodesFree() {
	_jsii_.InvokeVoid(
		c,
		"resetNodefsInodesFree",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) ResetPidAvailable() {
	_jsii_.InvokeVoid(
		c,
		"resetPidAvailable",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigKubeletConfigEvictionSoftOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

