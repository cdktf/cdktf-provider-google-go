// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clouddeployautomation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/clouddeployautomation/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ClouddeployAutomationRulesOutputReference interface {
	cdktf.ComplexObject
	AdvanceRolloutRule() ClouddeployAutomationRulesAdvanceRolloutRuleOutputReference
	AdvanceRolloutRuleInput() *ClouddeployAutomationRulesAdvanceRolloutRule
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PromoteReleaseRule() ClouddeployAutomationRulesPromoteReleaseRuleOutputReference
	PromoteReleaseRuleInput() *ClouddeployAutomationRulesPromoteReleaseRule
	RepairRolloutRule() ClouddeployAutomationRulesRepairRolloutRuleOutputReference
	RepairRolloutRuleInput() *ClouddeployAutomationRulesRepairRolloutRule
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimedPromoteReleaseRule() ClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference
	TimedPromoteReleaseRuleInput() *ClouddeployAutomationRulesTimedPromoteReleaseRule
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
	PutAdvanceRolloutRule(value *ClouddeployAutomationRulesAdvanceRolloutRule)
	PutPromoteReleaseRule(value *ClouddeployAutomationRulesPromoteReleaseRule)
	PutRepairRolloutRule(value *ClouddeployAutomationRulesRepairRolloutRule)
	PutTimedPromoteReleaseRule(value *ClouddeployAutomationRulesTimedPromoteReleaseRule)
	ResetAdvanceRolloutRule()
	ResetPromoteReleaseRule()
	ResetRepairRolloutRule()
	ResetTimedPromoteReleaseRule()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ClouddeployAutomationRulesOutputReference
type jsiiProxy_ClouddeployAutomationRulesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ClouddeployAutomationRulesOutputReference) AdvanceRolloutRule() ClouddeployAutomationRulesAdvanceRolloutRuleOutputReference {
	var returns ClouddeployAutomationRulesAdvanceRolloutRuleOutputReference
	_jsii_.Get(
		j,
		"advanceRolloutRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesOutputReference) AdvanceRolloutRuleInput() *ClouddeployAutomationRulesAdvanceRolloutRule {
	var returns *ClouddeployAutomationRulesAdvanceRolloutRule
	_jsii_.Get(
		j,
		"advanceRolloutRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesOutputReference) PromoteReleaseRule() ClouddeployAutomationRulesPromoteReleaseRuleOutputReference {
	var returns ClouddeployAutomationRulesPromoteReleaseRuleOutputReference
	_jsii_.Get(
		j,
		"promoteReleaseRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesOutputReference) PromoteReleaseRuleInput() *ClouddeployAutomationRulesPromoteReleaseRule {
	var returns *ClouddeployAutomationRulesPromoteReleaseRule
	_jsii_.Get(
		j,
		"promoteReleaseRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesOutputReference) RepairRolloutRule() ClouddeployAutomationRulesRepairRolloutRuleOutputReference {
	var returns ClouddeployAutomationRulesRepairRolloutRuleOutputReference
	_jsii_.Get(
		j,
		"repairRolloutRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesOutputReference) RepairRolloutRuleInput() *ClouddeployAutomationRulesRepairRolloutRule {
	var returns *ClouddeployAutomationRulesRepairRolloutRule
	_jsii_.Get(
		j,
		"repairRolloutRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesOutputReference) TimedPromoteReleaseRule() ClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference {
	var returns ClouddeployAutomationRulesTimedPromoteReleaseRuleOutputReference
	_jsii_.Get(
		j,
		"timedPromoteReleaseRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployAutomationRulesOutputReference) TimedPromoteReleaseRuleInput() *ClouddeployAutomationRulesTimedPromoteReleaseRule {
	var returns *ClouddeployAutomationRulesTimedPromoteReleaseRule
	_jsii_.Get(
		j,
		"timedPromoteReleaseRuleInput",
		&returns,
	)
	return returns
}


func NewClouddeployAutomationRulesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ClouddeployAutomationRulesOutputReference {
	_init_.Initialize()

	if err := validateNewClouddeployAutomationRulesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ClouddeployAutomationRulesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.clouddeployAutomation.ClouddeployAutomationRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewClouddeployAutomationRulesOutputReference_Override(c ClouddeployAutomationRulesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.clouddeployAutomation.ClouddeployAutomationRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ClouddeployAutomationRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ClouddeployAutomationRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ClouddeployAutomationRulesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ClouddeployAutomationRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ClouddeployAutomationRulesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ClouddeployAutomationRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClouddeployAutomationRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ClouddeployAutomationRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ClouddeployAutomationRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ClouddeployAutomationRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ClouddeployAutomationRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ClouddeployAutomationRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ClouddeployAutomationRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ClouddeployAutomationRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ClouddeployAutomationRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ClouddeployAutomationRulesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClouddeployAutomationRulesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ClouddeployAutomationRulesOutputReference) PutAdvanceRolloutRule(value *ClouddeployAutomationRulesAdvanceRolloutRule) {
	if err := c.validatePutAdvanceRolloutRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAdvanceRolloutRule",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ClouddeployAutomationRulesOutputReference) PutPromoteReleaseRule(value *ClouddeployAutomationRulesPromoteReleaseRule) {
	if err := c.validatePutPromoteReleaseRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPromoteReleaseRule",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ClouddeployAutomationRulesOutputReference) PutRepairRolloutRule(value *ClouddeployAutomationRulesRepairRolloutRule) {
	if err := c.validatePutRepairRolloutRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRepairRolloutRule",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ClouddeployAutomationRulesOutputReference) PutTimedPromoteReleaseRule(value *ClouddeployAutomationRulesTimedPromoteReleaseRule) {
	if err := c.validatePutTimedPromoteReleaseRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimedPromoteReleaseRule",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ClouddeployAutomationRulesOutputReference) ResetAdvanceRolloutRule() {
	_jsii_.InvokeVoid(
		c,
		"resetAdvanceRolloutRule",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClouddeployAutomationRulesOutputReference) ResetPromoteReleaseRule() {
	_jsii_.InvokeVoid(
		c,
		"resetPromoteReleaseRule",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClouddeployAutomationRulesOutputReference) ResetRepairRolloutRule() {
	_jsii_.InvokeVoid(
		c,
		"resetRepairRolloutRule",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClouddeployAutomationRulesOutputReference) ResetTimedPromoteReleaseRule() {
	_jsii_.InvokeVoid(
		c,
		"resetTimedPromoteReleaseRule",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClouddeployAutomationRulesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ClouddeployAutomationRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

