// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containernodepool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v10/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v10/containernodepool/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference interface {
	cdktf.ComplexObject
	BatchNodeCount() *float64
	SetBatchNodeCount(val *float64)
	BatchNodeCountInput() *float64
	BatchPercentage() *float64
	SetBatchPercentage(val *float64)
	BatchPercentageInput() *float64
	BatchSoakDuration() *string
	SetBatchSoakDuration(val *string)
	BatchSoakDurationInput() *string
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
	InternalValue() *ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicy
	SetInternalValue(val *ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicy)
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
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetBatchNodeCount()
	ResetBatchPercentage()
	ResetBatchSoakDuration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference
type jsiiProxy_ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) BatchNodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchNodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) BatchNodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchNodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) BatchPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) BatchPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) BatchSoakDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"batchSoakDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) BatchSoakDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"batchSoakDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) InternalValue() *ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicy {
	var returns *ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.containerNodePool.ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference_Override(c ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.containerNodePool.ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference)SetBatchNodeCount(val *float64) {
	if err := j.validateSetBatchNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchNodeCount",
		val,
	)
}

func (j *jsiiProxy_ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference)SetBatchPercentage(val *float64) {
	if err := j.validateSetBatchPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchPercentage",
		val,
	)
}

func (j *jsiiProxy_ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference)SetBatchSoakDuration(val *string) {
	if err := j.validateSetBatchSoakDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchSoakDuration",
		val,
	)
}

func (j *jsiiProxy_ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference)SetInternalValue(val *ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) ResetBatchNodeCount() {
	_jsii_.InvokeVoid(
		c,
		"resetBatchNodeCount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) ResetBatchPercentage() {
	_jsii_.InvokeVoid(
		c,
		"resetBatchPercentage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) ResetBatchSoakDuration() {
	_jsii_.InvokeVoid(
		c,
		"resetBatchSoakDuration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ContainerNodePoolUpgradeSettingsBlueGreenSettingsStandardRolloutPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

