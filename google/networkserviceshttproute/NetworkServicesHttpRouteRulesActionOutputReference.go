// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkserviceshttproute

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/networkserviceshttproute/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkServicesHttpRouteRulesActionOutputReference interface {
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
	CorsPolicy() NetworkServicesHttpRouteRulesActionCorsPolicyOutputReference
	CorsPolicyInput() *NetworkServicesHttpRouteRulesActionCorsPolicy
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Destinations() NetworkServicesHttpRouteRulesActionDestinationsList
	DestinationsInput() interface{}
	FaultInjectionPolicy() NetworkServicesHttpRouteRulesActionFaultInjectionPolicyOutputReference
	FaultInjectionPolicyInput() *NetworkServicesHttpRouteRulesActionFaultInjectionPolicy
	// Experimental.
	Fqn() *string
	InternalValue() *NetworkServicesHttpRouteRulesAction
	SetInternalValue(val *NetworkServicesHttpRouteRulesAction)
	Redirect() NetworkServicesHttpRouteRulesActionRedirectOutputReference
	RedirectInput() *NetworkServicesHttpRouteRulesActionRedirect
	RequestHeaderModifier() NetworkServicesHttpRouteRulesActionRequestHeaderModifierOutputReference
	RequestHeaderModifierInput() *NetworkServicesHttpRouteRulesActionRequestHeaderModifier
	RequestMirrorPolicy() NetworkServicesHttpRouteRulesActionRequestMirrorPolicyOutputReference
	RequestMirrorPolicyInput() *NetworkServicesHttpRouteRulesActionRequestMirrorPolicy
	ResponseHeaderModifier() NetworkServicesHttpRouteRulesActionResponseHeaderModifierOutputReference
	ResponseHeaderModifierInput() *NetworkServicesHttpRouteRulesActionResponseHeaderModifier
	RetryPolicy() NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference
	RetryPolicyInput() *NetworkServicesHttpRouteRulesActionRetryPolicy
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timeout() *string
	SetTimeout(val *string)
	TimeoutInput() *string
	UrlRewrite() NetworkServicesHttpRouteRulesActionUrlRewriteOutputReference
	UrlRewriteInput() *NetworkServicesHttpRouteRulesActionUrlRewrite
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
	PutCorsPolicy(value *NetworkServicesHttpRouteRulesActionCorsPolicy)
	PutDestinations(value interface{})
	PutFaultInjectionPolicy(value *NetworkServicesHttpRouteRulesActionFaultInjectionPolicy)
	PutRedirect(value *NetworkServicesHttpRouteRulesActionRedirect)
	PutRequestHeaderModifier(value *NetworkServicesHttpRouteRulesActionRequestHeaderModifier)
	PutRequestMirrorPolicy(value *NetworkServicesHttpRouteRulesActionRequestMirrorPolicy)
	PutResponseHeaderModifier(value *NetworkServicesHttpRouteRulesActionResponseHeaderModifier)
	PutRetryPolicy(value *NetworkServicesHttpRouteRulesActionRetryPolicy)
	PutUrlRewrite(value *NetworkServicesHttpRouteRulesActionUrlRewrite)
	ResetCorsPolicy()
	ResetDestinations()
	ResetFaultInjectionPolicy()
	ResetRedirect()
	ResetRequestHeaderModifier()
	ResetRequestMirrorPolicy()
	ResetResponseHeaderModifier()
	ResetRetryPolicy()
	ResetTimeout()
	ResetUrlRewrite()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkServicesHttpRouteRulesActionOutputReference
type jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) CorsPolicy() NetworkServicesHttpRouteRulesActionCorsPolicyOutputReference {
	var returns NetworkServicesHttpRouteRulesActionCorsPolicyOutputReference
	_jsii_.Get(
		j,
		"corsPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) CorsPolicyInput() *NetworkServicesHttpRouteRulesActionCorsPolicy {
	var returns *NetworkServicesHttpRouteRulesActionCorsPolicy
	_jsii_.Get(
		j,
		"corsPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) Destinations() NetworkServicesHttpRouteRulesActionDestinationsList {
	var returns NetworkServicesHttpRouteRulesActionDestinationsList
	_jsii_.Get(
		j,
		"destinations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) DestinationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) FaultInjectionPolicy() NetworkServicesHttpRouteRulesActionFaultInjectionPolicyOutputReference {
	var returns NetworkServicesHttpRouteRulesActionFaultInjectionPolicyOutputReference
	_jsii_.Get(
		j,
		"faultInjectionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) FaultInjectionPolicyInput() *NetworkServicesHttpRouteRulesActionFaultInjectionPolicy {
	var returns *NetworkServicesHttpRouteRulesActionFaultInjectionPolicy
	_jsii_.Get(
		j,
		"faultInjectionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) InternalValue() *NetworkServicesHttpRouteRulesAction {
	var returns *NetworkServicesHttpRouteRulesAction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) Redirect() NetworkServicesHttpRouteRulesActionRedirectOutputReference {
	var returns NetworkServicesHttpRouteRulesActionRedirectOutputReference
	_jsii_.Get(
		j,
		"redirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) RedirectInput() *NetworkServicesHttpRouteRulesActionRedirect {
	var returns *NetworkServicesHttpRouteRulesActionRedirect
	_jsii_.Get(
		j,
		"redirectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) RequestHeaderModifier() NetworkServicesHttpRouteRulesActionRequestHeaderModifierOutputReference {
	var returns NetworkServicesHttpRouteRulesActionRequestHeaderModifierOutputReference
	_jsii_.Get(
		j,
		"requestHeaderModifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) RequestHeaderModifierInput() *NetworkServicesHttpRouteRulesActionRequestHeaderModifier {
	var returns *NetworkServicesHttpRouteRulesActionRequestHeaderModifier
	_jsii_.Get(
		j,
		"requestHeaderModifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) RequestMirrorPolicy() NetworkServicesHttpRouteRulesActionRequestMirrorPolicyOutputReference {
	var returns NetworkServicesHttpRouteRulesActionRequestMirrorPolicyOutputReference
	_jsii_.Get(
		j,
		"requestMirrorPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) RequestMirrorPolicyInput() *NetworkServicesHttpRouteRulesActionRequestMirrorPolicy {
	var returns *NetworkServicesHttpRouteRulesActionRequestMirrorPolicy
	_jsii_.Get(
		j,
		"requestMirrorPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) ResponseHeaderModifier() NetworkServicesHttpRouteRulesActionResponseHeaderModifierOutputReference {
	var returns NetworkServicesHttpRouteRulesActionResponseHeaderModifierOutputReference
	_jsii_.Get(
		j,
		"responseHeaderModifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) ResponseHeaderModifierInput() *NetworkServicesHttpRouteRulesActionResponseHeaderModifier {
	var returns *NetworkServicesHttpRouteRulesActionResponseHeaderModifier
	_jsii_.Get(
		j,
		"responseHeaderModifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) RetryPolicy() NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference {
	var returns NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference
	_jsii_.Get(
		j,
		"retryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) RetryPolicyInput() *NetworkServicesHttpRouteRulesActionRetryPolicy {
	var returns *NetworkServicesHttpRouteRulesActionRetryPolicy
	_jsii_.Get(
		j,
		"retryPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) Timeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) TimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) UrlRewrite() NetworkServicesHttpRouteRulesActionUrlRewriteOutputReference {
	var returns NetworkServicesHttpRouteRulesActionUrlRewriteOutputReference
	_jsii_.Get(
		j,
		"urlRewrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) UrlRewriteInput() *NetworkServicesHttpRouteRulesActionUrlRewrite {
	var returns *NetworkServicesHttpRouteRulesActionUrlRewrite
	_jsii_.Get(
		j,
		"urlRewriteInput",
		&returns,
	)
	return returns
}


func NewNetworkServicesHttpRouteRulesActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NetworkServicesHttpRouteRulesActionOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkServicesHttpRouteRulesActionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.networkServicesHttpRoute.NetworkServicesHttpRouteRulesActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetworkServicesHttpRouteRulesActionOutputReference_Override(n NetworkServicesHttpRouteRulesActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.networkServicesHttpRoute.NetworkServicesHttpRouteRulesActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference)SetInternalValue(val *NetworkServicesHttpRouteRulesAction) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference)SetTimeout(val *string) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) PutCorsPolicy(value *NetworkServicesHttpRouteRulesActionCorsPolicy) {
	if err := n.validatePutCorsPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putCorsPolicy",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) PutDestinations(value interface{}) {
	if err := n.validatePutDestinationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putDestinations",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) PutFaultInjectionPolicy(value *NetworkServicesHttpRouteRulesActionFaultInjectionPolicy) {
	if err := n.validatePutFaultInjectionPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putFaultInjectionPolicy",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) PutRedirect(value *NetworkServicesHttpRouteRulesActionRedirect) {
	if err := n.validatePutRedirectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putRedirect",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) PutRequestHeaderModifier(value *NetworkServicesHttpRouteRulesActionRequestHeaderModifier) {
	if err := n.validatePutRequestHeaderModifierParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putRequestHeaderModifier",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) PutRequestMirrorPolicy(value *NetworkServicesHttpRouteRulesActionRequestMirrorPolicy) {
	if err := n.validatePutRequestMirrorPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putRequestMirrorPolicy",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) PutResponseHeaderModifier(value *NetworkServicesHttpRouteRulesActionResponseHeaderModifier) {
	if err := n.validatePutResponseHeaderModifierParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putResponseHeaderModifier",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) PutRetryPolicy(value *NetworkServicesHttpRouteRulesActionRetryPolicy) {
	if err := n.validatePutRetryPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putRetryPolicy",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) PutUrlRewrite(value *NetworkServicesHttpRouteRulesActionUrlRewrite) {
	if err := n.validatePutUrlRewriteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putUrlRewrite",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) ResetCorsPolicy() {
	_jsii_.InvokeVoid(
		n,
		"resetCorsPolicy",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) ResetDestinations() {
	_jsii_.InvokeVoid(
		n,
		"resetDestinations",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) ResetFaultInjectionPolicy() {
	_jsii_.InvokeVoid(
		n,
		"resetFaultInjectionPolicy",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) ResetRedirect() {
	_jsii_.InvokeVoid(
		n,
		"resetRedirect",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) ResetRequestHeaderModifier() {
	_jsii_.InvokeVoid(
		n,
		"resetRequestHeaderModifier",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) ResetRequestMirrorPolicy() {
	_jsii_.InvokeVoid(
		n,
		"resetRequestMirrorPolicy",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) ResetResponseHeaderModifier() {
	_jsii_.InvokeVoid(
		n,
		"resetResponseHeaderModifier",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) ResetRetryPolicy() {
	_jsii_.InvokeVoid(
		n,
		"resetRetryPolicy",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) ResetTimeout() {
	_jsii_.InvokeVoid(
		n,
		"resetTimeout",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) ResetUrlRewrite() {
	_jsii_.InvokeVoid(
		n,
		"resetUrlRewrite",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := n.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

