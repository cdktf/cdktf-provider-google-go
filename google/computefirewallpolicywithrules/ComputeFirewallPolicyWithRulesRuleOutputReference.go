// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computefirewallpolicywithrules

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/computefirewallpolicywithrules/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeFirewallPolicyWithRulesRuleOutputReference interface {
	cdktf.ComplexObject
	Action() *string
	SetAction(val *string)
	ActionInput() *string
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Direction() *string
	SetDirection(val *string)
	DirectionInput() *string
	Disabled() interface{}
	SetDisabled(val interface{})
	DisabledInput() interface{}
	EnableLogging() interface{}
	SetEnableLogging(val interface{})
	EnableLoggingInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Match() ComputeFirewallPolicyWithRulesRuleMatchOutputReference
	MatchInput() *ComputeFirewallPolicyWithRulesRuleMatch
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
	RuleName() *string
	SetRuleName(val *string)
	RuleNameInput() *string
	SecurityProfileGroup() *string
	SetSecurityProfileGroup(val *string)
	SecurityProfileGroupInput() *string
	TargetResources() *[]*string
	SetTargetResources(val *[]*string)
	TargetResourcesInput() *[]*string
	TargetServiceAccounts() *[]*string
	SetTargetServiceAccounts(val *[]*string)
	TargetServiceAccountsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TlsInspect() interface{}
	SetTlsInspect(val interface{})
	TlsInspectInput() interface{}
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
	PutMatch(value *ComputeFirewallPolicyWithRulesRuleMatch)
	ResetDescription()
	ResetDirection()
	ResetDisabled()
	ResetEnableLogging()
	ResetRuleName()
	ResetSecurityProfileGroup()
	ResetTargetResources()
	ResetTargetServiceAccounts()
	ResetTlsInspect()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeFirewallPolicyWithRulesRuleOutputReference
type jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) Direction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"direction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) DirectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) Disabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) DisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) EnableLogging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) EnableLoggingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLoggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) Match() ComputeFirewallPolicyWithRulesRuleMatchOutputReference {
	var returns ComputeFirewallPolicyWithRulesRuleMatchOutputReference
	_jsii_.Get(
		j,
		"match",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) MatchInput() *ComputeFirewallPolicyWithRulesRuleMatch {
	var returns *ComputeFirewallPolicyWithRulesRuleMatch
	_jsii_.Get(
		j,
		"matchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) RuleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) RuleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) SecurityProfileGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProfileGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) SecurityProfileGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityProfileGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) TargetResources() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) TargetResourcesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) TargetServiceAccounts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetServiceAccounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) TargetServiceAccountsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetServiceAccountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) TlsInspect() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsInspect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) TlsInspectInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsInspectInput",
		&returns,
	)
	return returns
}


func NewComputeFirewallPolicyWithRulesRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ComputeFirewallPolicyWithRulesRuleOutputReference {
	_init_.Initialize()

	if err := validateNewComputeFirewallPolicyWithRulesRuleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeFirewallPolicyWithRules.ComputeFirewallPolicyWithRulesRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewComputeFirewallPolicyWithRulesRuleOutputReference_Override(c ComputeFirewallPolicyWithRulesRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeFirewallPolicyWithRules.ComputeFirewallPolicyWithRulesRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference)SetAction(val *string) {
	if err := j.validateSetActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference)SetDirection(val *string) {
	if err := j.validateSetDirectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"direction",
		val,
	)
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference)SetDisabled(val interface{}) {
	if err := j.validateSetDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disabled",
		val,
	)
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference)SetEnableLogging(val interface{}) {
	if err := j.validateSetEnableLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableLogging",
		val,
	)
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference)SetRuleName(val *string) {
	if err := j.validateSetRuleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleName",
		val,
	)
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference)SetSecurityProfileGroup(val *string) {
	if err := j.validateSetSecurityProfileGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityProfileGroup",
		val,
	)
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference)SetTargetResources(val *[]*string) {
	if err := j.validateSetTargetResourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetResources",
		val,
	)
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference)SetTargetServiceAccounts(val *[]*string) {
	if err := j.validateSetTargetServiceAccountsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetServiceAccounts",
		val,
	)
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference)SetTlsInspect(val interface{}) {
	if err := j.validateSetTlsInspectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsInspect",
		val,
	)
}

func (c *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) PutMatch(value *ComputeFirewallPolicyWithRulesRuleMatch) {
	if err := c.validatePutMatchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMatch",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) ResetDirection() {
	_jsii_.InvokeVoid(
		c,
		"resetDirection",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) ResetDisabled() {
	_jsii_.InvokeVoid(
		c,
		"resetDisabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) ResetEnableLogging() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableLogging",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) ResetRuleName() {
	_jsii_.InvokeVoid(
		c,
		"resetRuleName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) ResetSecurityProfileGroup() {
	_jsii_.InvokeVoid(
		c,
		"resetSecurityProfileGroup",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) ResetTargetResources() {
	_jsii_.InvokeVoid(
		c,
		"resetTargetResources",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) ResetTargetServiceAccounts() {
	_jsii_.InvokeVoid(
		c,
		"resetTargetServiceAccounts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) ResetTlsInspect() {
	_jsii_.InvokeVoid(
		c,
		"resetTlsInspect",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeFirewallPolicyWithRulesRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

