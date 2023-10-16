// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeresourcepolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v11/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v11/computeresourcepolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeResourcePolicySnapshotSchedulePolicyOutputReference interface {
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
	InternalValue() *ComputeResourcePolicySnapshotSchedulePolicy
	SetInternalValue(val *ComputeResourcePolicySnapshotSchedulePolicy)
	RetentionPolicy() ComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference
	RetentionPolicyInput() *ComputeResourcePolicySnapshotSchedulePolicyRetentionPolicy
	Schedule() ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference
	ScheduleInput() *ComputeResourcePolicySnapshotSchedulePolicySchedule
	SnapshotProperties() ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference
	SnapshotPropertiesInput() *ComputeResourcePolicySnapshotSchedulePolicySnapshotProperties
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
	PutRetentionPolicy(value *ComputeResourcePolicySnapshotSchedulePolicyRetentionPolicy)
	PutSchedule(value *ComputeResourcePolicySnapshotSchedulePolicySchedule)
	PutSnapshotProperties(value *ComputeResourcePolicySnapshotSchedulePolicySnapshotProperties)
	ResetRetentionPolicy()
	ResetSnapshotProperties()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeResourcePolicySnapshotSchedulePolicyOutputReference
type jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyOutputReference) InternalValue() *ComputeResourcePolicySnapshotSchedulePolicy {
	var returns *ComputeResourcePolicySnapshotSchedulePolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyOutputReference) RetentionPolicy() ComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference {
	var returns ComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference
	_jsii_.Get(
		j,
		"retentionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyOutputReference) RetentionPolicyInput() *ComputeResourcePolicySnapshotSchedulePolicyRetentionPolicy {
	var returns *ComputeResourcePolicySnapshotSchedulePolicyRetentionPolicy
	_jsii_.Get(
		j,
		"retentionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyOutputReference) Schedule() ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference {
	var returns ComputeResourcePolicySnapshotSchedulePolicyScheduleOutputReference
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyOutputReference) ScheduleInput() *ComputeResourcePolicySnapshotSchedulePolicySchedule {
	var returns *ComputeResourcePolicySnapshotSchedulePolicySchedule
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyOutputReference) SnapshotProperties() ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference {
	var returns ComputeResourcePolicySnapshotSchedulePolicySnapshotPropertiesOutputReference
	_jsii_.Get(
		j,
		"snapshotProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyOutputReference) SnapshotPropertiesInput() *ComputeResourcePolicySnapshotSchedulePolicySnapshotProperties {
	var returns *ComputeResourcePolicySnapshotSchedulePolicySnapshotProperties
	_jsii_.Get(
		j,
		"snapshotPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeResourcePolicySnapshotSchedulePolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeResourcePolicySnapshotSchedulePolicyOutputReference {
	_init_.Initialize()

	if err := validateNewComputeResourcePolicySnapshotSchedulePolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeResourcePolicy.ComputeResourcePolicySnapshotSchedulePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeResourcePolicySnapshotSchedulePolicyOutputReference_Override(c ComputeResourcePolicySnapshotSchedulePolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeResourcePolicy.ComputeResourcePolicySnapshotSchedulePolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyOutputReference)SetInternalValue(val *ComputeResourcePolicySnapshotSchedulePolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyOutputReference) PutRetentionPolicy(value *ComputeResourcePolicySnapshotSchedulePolicyRetentionPolicy) {
	if err := c.validatePutRetentionPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRetentionPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyOutputReference) PutSchedule(value *ComputeResourcePolicySnapshotSchedulePolicySchedule) {
	if err := c.validatePutScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSchedule",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyOutputReference) PutSnapshotProperties(value *ComputeResourcePolicySnapshotSchedulePolicySnapshotProperties) {
	if err := c.validatePutSnapshotPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSnapshotProperties",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyOutputReference) ResetRetentionPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetRetentionPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyOutputReference) ResetSnapshotProperties() {
	_jsii_.InvokeVoid(
		c,
		"resetSnapshotProperties",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeResourcePolicySnapshotSchedulePolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

