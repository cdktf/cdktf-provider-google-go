// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computesecuritypolicyrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/computesecuritypolicyrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RequestCookie() ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookieAList
	RequestCookieInput() interface{}
	RequestHeader() ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionRequestHeaderAList
	RequestHeaderInput() interface{}
	RequestQueryParam() ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParamAList
	RequestQueryParamInput() interface{}
	RequestUri() ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionRequestUriAList
	RequestUriInput() interface{}
	TargetRuleIds() *[]*string
	SetTargetRuleIds(val *[]*string)
	TargetRuleIdsInput() *[]*string
	TargetRuleSet() *string
	SetTargetRuleSet(val *string)
	TargetRuleSetInput() *string
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
	PutRequestCookie(value interface{})
	PutRequestHeader(value interface{})
	PutRequestQueryParam(value interface{})
	PutRequestUri(value interface{})
	ResetRequestCookie()
	ResetRequestHeader()
	ResetRequestQueryParam()
	ResetRequestUri()
	ResetTargetRuleIds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference
type jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference) RequestCookie() ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookieAList {
	var returns ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionRequestCookieAList
	_jsii_.Get(
		j,
		"requestCookie",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference) RequestCookieInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestCookieInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference) RequestHeader() ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionRequestHeaderAList {
	var returns ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionRequestHeaderAList
	_jsii_.Get(
		j,
		"requestHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference) RequestHeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference) RequestQueryParam() ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParamAList {
	var returns ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionRequestQueryParamAList
	_jsii_.Get(
		j,
		"requestQueryParam",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference) RequestQueryParamInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestQueryParamInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference) RequestUri() ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionRequestUriAList {
	var returns ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionRequestUriAList
	_jsii_.Get(
		j,
		"requestUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference) RequestUriInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference) TargetRuleIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetRuleIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference) TargetRuleIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetRuleIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference) TargetRuleSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetRuleSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference) TargetRuleSetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetRuleSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference {
	_init_.Initialize()

	if err := validateNewComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeSecurityPolicyRule.ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference_Override(c ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeSecurityPolicyRule.ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference)SetTargetRuleIds(val *[]*string) {
	if err := j.validateSetTargetRuleIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetRuleIds",
		val,
	)
}

func (j *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference)SetTargetRuleSet(val *string) {
	if err := j.validateSetTargetRuleSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetRuleSet",
		val,
	)
}

func (j *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference) PutRequestCookie(value interface{}) {
	if err := c.validatePutRequestCookieParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRequestCookie",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference) PutRequestHeader(value interface{}) {
	if err := c.validatePutRequestHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRequestHeader",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference) PutRequestQueryParam(value interface{}) {
	if err := c.validatePutRequestQueryParamParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRequestQueryParam",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference) PutRequestUri(value interface{}) {
	if err := c.validatePutRequestUriParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRequestUri",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference) ResetRequestCookie() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestCookie",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference) ResetRequestHeader() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestHeader",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference) ResetRequestQueryParam() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestQueryParam",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference) ResetRequestUri() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestUri",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference) ResetTargetRuleIds() {
	_jsii_.InvokeVoid(
		c,
		"resetTargetRuleIds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeSecurityPolicyRulePreconfiguredWafConfigExclusionAOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

