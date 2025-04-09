// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clouddeployautomation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/clouddeployautomation/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference interface {
	cdktf.ComplexObject
	Attempts() *string
	SetAttempts(val *string)
	AttemptsInput() *string
	BackoffMode() *string
	SetBackoffMode(val *string)
	BackoffModeInput() *string
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
	InternalValue() *ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetry
	SetInternalValue(val *ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetry)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Wait() *string
	SetWait(val *string)
	WaitInput() *string
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
	ResetBackoffMode()
	ResetWait()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference
type jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) Attempts() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) AttemptsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attemptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) BackoffMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backoffMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) BackoffModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backoffModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) InternalValue() *ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetry {
	var returns *ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetry
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) Wait() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wait",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) WaitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"waitInput",
		&returns,
	)
	return returns
}


func NewClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference {
	_init_.Initialize()

	if err := validateNewClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.clouddeployAutomation.ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference_Override(c ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.clouddeployAutomation.ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference)SetAttempts(val *string) {
	if err := j.validateSetAttemptsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attempts",
		val,
	)
}

func (j *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference)SetBackoffMode(val *string) {
	if err := j.validateSetBackoffModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backoffMode",
		val,
	)
}

func (j *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference)SetInternalValue(val *ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetry) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference)SetWait(val *string) {
	if err := j.validateSetWaitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wait",
		val,
	)
}

func (c *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) ResetBackoffMode() {
	_jsii_.InvokeVoid(
		c,
		"resetBackoffMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) ResetWait() {
	_jsii_.InvokeVoid(
		c,
		"resetWait",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

