// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clouddeploydeploypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/clouddeploydeploypolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference interface {
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
	EndDate() ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsEndDateOutputReference
	EndDateInput() *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsEndDate
	EndTime() ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsEndTimeOutputReference
	EndTimeInput() *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsEndTime
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	StartDate() ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsStartDateOutputReference
	StartDateInput() *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsStartDate
	StartTime() ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsStartTimeOutputReference
	StartTimeInput() *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsStartTime
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
	PutEndDate(value *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsEndDate)
	PutEndTime(value *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsEndTime)
	PutStartDate(value *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsStartDate)
	PutStartTime(value *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsStartTime)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference
type jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) EndDate() ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsEndDateOutputReference {
	var returns ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsEndDateOutputReference
	_jsii_.Get(
		j,
		"endDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) EndDateInput() *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsEndDate {
	var returns *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsEndDate
	_jsii_.Get(
		j,
		"endDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) EndTime() ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsEndTimeOutputReference {
	var returns ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsEndTimeOutputReference
	_jsii_.Get(
		j,
		"endTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) EndTimeInput() *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsEndTime {
	var returns *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsEndTime
	_jsii_.Get(
		j,
		"endTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) StartDate() ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsStartDateOutputReference {
	var returns ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsStartDateOutputReference
	_jsii_.Get(
		j,
		"startDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) StartDateInput() *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsStartDate {
	var returns *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsStartDate
	_jsii_.Get(
		j,
		"startDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) StartTime() ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsStartTimeOutputReference {
	var returns ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsStartTimeOutputReference
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) StartTimeInput() *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsStartTime {
	var returns *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsStartTime
	_jsii_.Get(
		j,
		"startTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference {
	_init_.Initialize()

	if err := validateNewClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.clouddeployDeployPolicy.ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference_Override(c ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.clouddeployDeployPolicy.ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) PutEndDate(value *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsEndDate) {
	if err := c.validatePutEndDateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEndDate",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) PutEndTime(value *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsEndTime) {
	if err := c.validatePutEndTimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEndTime",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) PutStartDate(value *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsStartDate) {
	if err := c.validatePutStartDateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putStartDate",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) PutStartTime(value *ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsStartTime) {
	if err := c.validatePutStartTimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putStartTime",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ClouddeployDeployPolicyRulesRolloutRestrictionTimeWindowsOneTimeWindowsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

