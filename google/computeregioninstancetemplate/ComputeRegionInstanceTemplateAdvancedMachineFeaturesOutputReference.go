// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregioninstancetemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/computeregioninstancetemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference interface {
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
	EnableNestedVirtualization() interface{}
	SetEnableNestedVirtualization(val interface{})
	EnableNestedVirtualizationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *ComputeRegionInstanceTemplateAdvancedMachineFeatures
	SetInternalValue(val *ComputeRegionInstanceTemplateAdvancedMachineFeatures)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ThreadsPerCore() *float64
	SetThreadsPerCore(val *float64)
	ThreadsPerCoreInput() *float64
	TurboMode() *string
	SetTurboMode(val *string)
	TurboModeInput() *string
	VisibleCoreCount() *float64
	SetVisibleCoreCount(val *float64)
	VisibleCoreCountInput() *float64
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
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetEnableNestedVirtualization()
	ResetThreadsPerCore()
	ResetTurboMode()
	ResetVisibleCoreCount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference
type jsiiProxy_ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference) EnableNestedVirtualization() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNestedVirtualization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference) EnableNestedVirtualizationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNestedVirtualizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference) InternalValue() *ComputeRegionInstanceTemplateAdvancedMachineFeatures {
	var returns *ComputeRegionInstanceTemplateAdvancedMachineFeatures
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference) ThreadsPerCore() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadsPerCore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference) ThreadsPerCoreInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadsPerCoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference) TurboMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"turboMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference) TurboModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"turboModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference) VisibleCoreCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"visibleCoreCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference) VisibleCoreCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"visibleCoreCountInput",
		&returns,
	)
	return returns
}


func NewComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference {
	_init_.Initialize()

	if err := validateNewComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionInstanceTemplate.ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference_Override(c ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionInstanceTemplate.ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference)SetEnableNestedVirtualization(val interface{}) {
	if err := j.validateSetEnableNestedVirtualizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableNestedVirtualization",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference)SetInternalValue(val *ComputeRegionInstanceTemplateAdvancedMachineFeatures) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference)SetThreadsPerCore(val *float64) {
	if err := j.validateSetThreadsPerCoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threadsPerCore",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference)SetTurboMode(val *string) {
	if err := j.validateSetTurboModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"turboMode",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference)SetVisibleCoreCount(val *float64) {
	if err := j.validateSetVisibleCoreCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visibleCoreCount",
		val,
	)
}

func (c *jsiiProxy_ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference) ResetEnableNestedVirtualization() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableNestedVirtualization",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference) ResetThreadsPerCore() {
	_jsii_.InvokeVoid(
		c,
		"resetThreadsPerCore",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference) ResetTurboMode() {
	_jsii_.InvokeVoid(
		c,
		"resetTurboMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference) ResetVisibleCoreCount() {
	_jsii_.InvokeVoid(
		c,
		"resetVisibleCoreCount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeRegionInstanceTemplateAdvancedMachineFeaturesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

