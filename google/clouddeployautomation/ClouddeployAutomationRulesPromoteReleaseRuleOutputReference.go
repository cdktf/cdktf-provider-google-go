// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clouddeployautomation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/clouddeployautomation/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ClouddeployAutomationRulesPromoteReleaseRuleOutputReference interface {
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
	DestinationPhase() *string
	SetDestinationPhase(val *string)
	DestinationPhaseInput() *string
	DestinationTargetId() *string
	SetDestinationTargetId(val *string)
	DestinationTargetIdInput() *string
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() *ClouddeployAutomationRulesPromoteReleaseRule
	SetInternalValue(val *ClouddeployAutomationRulesPromoteReleaseRule)
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
	ResetDestinationPhase()
	ResetDestinationTargetId()
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

// The jsii proxy struct for ClouddeployAutomationRulesPromoteReleaseRuleOutputReference
type jsiiProxy_ClouddeployAutomationRulesPromoteReleaseRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ClouddeployAutomationRulesPromoteReleaseRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesPromoteReleaseRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesPromoteReleaseRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesPromoteReleaseRuleOutputReference) DestinationPhase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPhase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesPromoteReleaseRuleOutputReference) DestinationPhaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPhaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesPromoteReleaseRuleOutputReference) DestinationTargetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationTargetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesPromoteReleaseRuleOutputReference) DestinationTargetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationTargetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesPromoteReleaseRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesPromoteReleaseRuleOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesPromoteReleaseRuleOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesPromoteReleaseRuleOutputReference) InternalValue() *ClouddeployAutomationRulesPromoteReleaseRule {
	var returns *ClouddeployAutomationRulesPromoteReleaseRule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesPromoteReleaseRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesPromoteReleaseRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesPromoteReleaseRuleOutputReference) Wait() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wait",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesPromoteReleaseRuleOutputReference) WaitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"waitInput",
		&returns,
	)
	return returns
}


func NewClouddeployAutomationRulesPromoteReleaseRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ClouddeployAutomationRulesPromoteReleaseRuleOutputReference {
	_init_.Initialize()

	if err := validateNewClouddeployAutomationRulesPromoteReleaseRuleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ClouddeployAutomationRulesPromoteReleaseRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.clouddeployAutomation.ClouddeployAutomationRulesPromoteReleaseRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewClouddeployAutomationRulesPromoteReleaseRuleOutputReference_Override(c ClouddeployAutomationRulesPromoteReleaseRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.clouddeployAutomation.ClouddeployAutomationRulesPromoteReleaseRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ClouddeployAutomationRulesPromoteReleaseRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ClouddeployAutomationRulesPromoteReleaseRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ClouddeployAutomationRulesPromoteReleaseRuleOutputReference)SetDestinationPhase(val *string) {
	if err := j.validateSetDestinationPhaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationPhase",
		val,
	)
}

func (j *jsiiProxy_ClouddeployAutomationRulesPromoteReleaseRuleOutputReference)SetDestinationTargetId(val *string) {
	if err := j.validateSetDestinationTargetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationTargetId",
		val,
	)
}

func (j *jsiiProxy_ClouddeployAutomationRulesPromoteReleaseRuleOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ClouddeployAutomationRulesPromoteReleaseRuleOutputReference)SetInternalValue(val *ClouddeployAutomationRulesPromoteReleaseRule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ClouddeployAutomationRulesPromoteReleaseRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ClouddeployAutomationRulesPromoteReleaseRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ClouddeployAutomationRulesPromoteReleaseRuleOutputReference)SetWait(val *string) {
	if err := j.validateSetWaitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wait",
		val,
	)
}

func (c *jsiiProxy_ClouddeployAutomationRulesPromoteReleaseRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClouddeployAutomationRulesPromoteReleaseRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ClouddeployAutomationRulesPromoteReleaseRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ClouddeployAutomationRulesPromoteReleaseRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ClouddeployAutomationRulesPromoteReleaseRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ClouddeployAutomationRulesPromoteReleaseRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ClouddeployAutomationRulesPromoteReleaseRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ClouddeployAutomationRulesPromoteReleaseRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ClouddeployAutomationRulesPromoteReleaseRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ClouddeployAutomationRulesPromoteReleaseRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ClouddeployAutomationRulesPromoteReleaseRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClouddeployAutomationRulesPromoteReleaseRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ClouddeployAutomationRulesPromoteReleaseRuleOutputReference) ResetDestinationPhase() {
	_jsii_.InvokeVoid(
		c,
		"resetDestinationPhase",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClouddeployAutomationRulesPromoteReleaseRuleOutputReference) ResetDestinationTargetId() {
	_jsii_.InvokeVoid(
		c,
		"resetDestinationTargetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClouddeployAutomationRulesPromoteReleaseRuleOutputReference) ResetWait() {
	_jsii_.InvokeVoid(
		c,
		"resetWait",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClouddeployAutomationRulesPromoteReleaseRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ClouddeployAutomationRulesPromoteReleaseRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

