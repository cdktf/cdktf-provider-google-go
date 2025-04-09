// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clouddeploydeploypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/clouddeploydeploypolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference interface {
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
	DaysOfWeek() *[]*string
	SetDaysOfWeek(val *[]*string)
	DaysOfWeekInput() *[]*string
	EndTime() ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference
	EndTimeInput() *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTime
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	StartTime() ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsStartTimeOutputReference
	StartTimeInput() *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsStartTime
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
	PutEndTime(value *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTime)
	PutStartTime(value *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsStartTime)
	ResetDaysOfWeek()
	ResetEndTime()
	ResetStartTime()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference
type jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference) DaysOfWeek() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"daysOfWeek",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference) DaysOfWeekInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"daysOfWeekInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference) EndTime() ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference {
	var returns ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTimeOutputReference
	_jsii_.Get(
		j,
		"endTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference) EndTimeInput() *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTime {
	var returns *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTime
	_jsii_.Get(
		j,
		"endTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference) StartTime() ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsStartTimeOutputReference {
	var returns ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsStartTimeOutputReference
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference) StartTimeInput() *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsStartTime {
	var returns *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsStartTime
	_jsii_.Get(
		j,
		"startTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference {
	_init_.Initialize()

	if err := validateNewClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.clouddeployDeployPolicy.ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference_Override(c ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.clouddeployDeployPolicy.ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference)SetDaysOfWeek(val *[]*string) {
	if err := j.validateSetDaysOfWeekParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"daysOfWeek",
		val,
	)
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference) PutEndTime(value *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsEndTime) {
	if err := c.validatePutEndTimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEndTime",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference) PutStartTime(value *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsStartTime) {
	if err := c.validatePutStartTimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putStartTime",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference) ResetDaysOfWeek() {
	_jsii_.InvokeVoid(
		c,
		"resetDaysOfWeek",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference) ResetEndTime() {
	_jsii_.InvokeVoid(
		c,
		"resetEndTime",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference) ResetStartTime() {
	_jsii_.InvokeVoid(
		c,
		"resetStartTime",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsWeeklyWindowsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

