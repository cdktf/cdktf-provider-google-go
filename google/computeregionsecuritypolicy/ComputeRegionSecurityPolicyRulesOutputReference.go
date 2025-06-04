// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregionsecuritypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/computeregionsecuritypolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeRegionSecurityPolicyRulesOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Match() ComputeRegionSecurityPolicyRulesMatchOutputReference
	MatchInput() *ComputeRegionSecurityPolicyRulesMatch
	NetworkMatch() ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference
	NetworkMatchInput() *ComputeRegionSecurityPolicyRulesNetworkMatch
	PreconfiguredWafConfig() ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigOutputReference
	PreconfiguredWafConfigInput() *ComputeRegionSecurityPolicyRulesPreconfiguredWafConfig
	Preview() interface{}
	SetPreview(val interface{})
	PreviewInput() interface{}
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
	RateLimitOptions() ComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference
	RateLimitOptionsInput() *ComputeRegionSecurityPolicyRulesRateLimitOptions
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
	PutMatch(value *ComputeRegionSecurityPolicyRulesMatch)
	PutNetworkMatch(value *ComputeRegionSecurityPolicyRulesNetworkMatch)
	PutPreconfiguredWafConfig(value *ComputeRegionSecurityPolicyRulesPreconfiguredWafConfig)
	PutRateLimitOptions(value *ComputeRegionSecurityPolicyRulesRateLimitOptions)
	ResetDescription()
	ResetMatch()
	ResetNetworkMatch()
	ResetPreconfiguredWafConfig()
	ResetPreview()
	ResetRateLimitOptions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeRegionSecurityPolicyRulesOutputReference
type jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) Match() ComputeRegionSecurityPolicyRulesMatchOutputReference {
	var returns ComputeRegionSecurityPolicyRulesMatchOutputReference
	_jsii_.Get(
		j,
		"match",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) MatchInput() *ComputeRegionSecurityPolicyRulesMatch {
	var returns *ComputeRegionSecurityPolicyRulesMatch
	_jsii_.Get(
		j,
		"matchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) NetworkMatch() ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference {
	var returns ComputeRegionSecurityPolicyRulesNetworkMatchOutputReference
	_jsii_.Get(
		j,
		"networkMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) NetworkMatchInput() *ComputeRegionSecurityPolicyRulesNetworkMatch {
	var returns *ComputeRegionSecurityPolicyRulesNetworkMatch
	_jsii_.Get(
		j,
		"networkMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) PreconfiguredWafConfig() ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigOutputReference {
	var returns ComputeRegionSecurityPolicyRulesPreconfiguredWafConfigOutputReference
	_jsii_.Get(
		j,
		"preconfiguredWafConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) PreconfiguredWafConfigInput() *ComputeRegionSecurityPolicyRulesPreconfiguredWafConfig {
	var returns *ComputeRegionSecurityPolicyRulesPreconfiguredWafConfig
	_jsii_.Get(
		j,
		"preconfiguredWafConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) Preview() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preview",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) PreviewInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"previewInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) RateLimitOptions() ComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference {
	var returns ComputeRegionSecurityPolicyRulesRateLimitOptionsOutputReference
	_jsii_.Get(
		j,
		"rateLimitOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) RateLimitOptionsInput() *ComputeRegionSecurityPolicyRulesRateLimitOptions {
	var returns *ComputeRegionSecurityPolicyRulesRateLimitOptions
	_jsii_.Get(
		j,
		"rateLimitOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeRegionSecurityPolicyRulesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ComputeRegionSecurityPolicyRulesOutputReference {
	_init_.Initialize()

	if err := validateNewComputeRegionSecurityPolicyRulesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionSecurityPolicy.ComputeRegionSecurityPolicyRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewComputeRegionSecurityPolicyRulesOutputReference_Override(c ComputeRegionSecurityPolicyRulesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionSecurityPolicy.ComputeRegionSecurityPolicyRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference)SetAction(val *string) {
	if err := j.validateSetActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference)SetPreview(val interface{}) {
	if err := j.validateSetPreviewParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preview",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) PutMatch(value *ComputeRegionSecurityPolicyRulesMatch) {
	if err := c.validatePutMatchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMatch",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) PutNetworkMatch(value *ComputeRegionSecurityPolicyRulesNetworkMatch) {
	if err := c.validatePutNetworkMatchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNetworkMatch",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) PutPreconfiguredWafConfig(value *ComputeRegionSecurityPolicyRulesPreconfiguredWafConfig) {
	if err := c.validatePutPreconfiguredWafConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPreconfiguredWafConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) PutRateLimitOptions(value *ComputeRegionSecurityPolicyRulesRateLimitOptions) {
	if err := c.validatePutRateLimitOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRateLimitOptions",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) ResetMatch() {
	_jsii_.InvokeVoid(
		c,
		"resetMatch",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) ResetNetworkMatch() {
	_jsii_.InvokeVoid(
		c,
		"resetNetworkMatch",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) ResetPreconfiguredWafConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetPreconfiguredWafConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) ResetPreview() {
	_jsii_.InvokeVoid(
		c,
		"resetPreview",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) ResetRateLimitOptions() {
	_jsii_.InvokeVoid(
		c,
		"resetRateLimitOptions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeRegionSecurityPolicyRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

