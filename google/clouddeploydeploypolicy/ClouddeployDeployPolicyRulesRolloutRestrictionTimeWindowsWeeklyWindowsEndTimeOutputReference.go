// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clouddeploydeploypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/clouddeploydeploypolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference interface {
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
	Hours() *float64
	SetHours(val *float64)
	HoursInput() *float64
	InternalValue() *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTime
	SetInternalValue(val *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTime)
	Minutes() *float64
	SetMinutes(val *float64)
	MinutesInput() *float64
	Nanos() *float64
	SetNanos(val *float64)
	NanosInput() *float64
	Seconds() *float64
	SetSeconds(val *float64)
	SecondsInput() *float64
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
	ResetHours()
	ResetMinutes()
	ResetNanos()
	ResetSeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference
type jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference) Hours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference) HoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference) InternalValue() *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTime {
	var returns *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTime
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference) Minutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference) MinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference) Nanos() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nanos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference) NanosInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nanosInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference) Seconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"seconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference) SecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"secondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference {
	_init_.Initialize()

	if err := validateNewClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.clouddeployDeployPolicy.ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference_Override(c ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.clouddeployDeployPolicy.ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference)SetHours(val *float64) {
	if err := j.validateSetHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hours",
		val,
	)
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference)SetInternalValue(val *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTime) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference)SetMinutes(val *float64) {
	if err := j.validateSetMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minutes",
		val,
	)
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference)SetNanos(val *float64) {
	if err := j.validateSetNanosParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nanos",
		val,
	)
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference)SetSeconds(val *float64) {
	if err := j.validateSetSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"seconds",
		val,
	)
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference) ResetHours() {
	_jsii_.InvokeVoid(
		c,
		"resetHours",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference) ResetMinutes() {
	_jsii_.InvokeVoid(
		c,
		"resetMinutes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference) ResetNanos() {
	_jsii_.InvokeVoid(
		c,
		"resetNanos",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference) ResetSeconds() {
	_jsii_.InvokeVoid(
		c,
		"resetSeconds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

