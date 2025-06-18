// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregionurlmap

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/computeregionurlmap/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference interface {
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
	CorsPolicy() ComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference
	CorsPolicyInput() *ComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicy
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	FaultInjectionPolicy() ComputeRegionUrlMapPathMatcherDefaultRouteActionFaultInjectionPolicyOutputReference
	FaultInjectionPolicyInput() *ComputeRegionUrlMapPathMatcherDefaultRouteActionFaultInjectionPolicy
	// Experimental.
	Fqn() *string
	InternalValue() *ComputeRegionUrlMapPathMatcherDefaultRouteAction
	SetInternalValue(val *ComputeRegionUrlMapPathMatcherDefaultRouteAction)
	MaxStreamDuration() ComputeRegionUrlMapPathMatcherDefaultRouteActionMaxStreamDurationOutputReference
	MaxStreamDurationInput() *ComputeRegionUrlMapPathMatcherDefaultRouteActionMaxStreamDuration
	RequestMirrorPolicy() ComputeRegionUrlMapPathMatcherDefaultRouteActionRequestMirrorPolicyOutputReference
	RequestMirrorPolicyInput() *ComputeRegionUrlMapPathMatcherDefaultRouteActionRequestMirrorPolicy
	RetryPolicy() ComputeRegionUrlMapPathMatcherDefaultRouteActionRetryPolicyOutputReference
	RetryPolicyInput() *ComputeRegionUrlMapPathMatcherDefaultRouteActionRetryPolicy
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timeout() ComputeRegionUrlMapPathMatcherDefaultRouteActionTimeoutOutputReference
	TimeoutInput() *ComputeRegionUrlMapPathMatcherDefaultRouteActionTimeout
	UrlRewrite() ComputeRegionUrlMapPathMatcherDefaultRouteActionUrlRewriteOutputReference
	UrlRewriteInput() *ComputeRegionUrlMapPathMatcherDefaultRouteActionUrlRewrite
	WeightedBackendServices() ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesList
	WeightedBackendServicesInput() interface{}
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
	PutCorsPolicy(value *ComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicy)
	PutFaultInjectionPolicy(value *ComputeRegionUrlMapPathMatcherDefaultRouteActionFaultInjectionPolicy)
	PutMaxStreamDuration(value *ComputeRegionUrlMapPathMatcherDefaultRouteActionMaxStreamDuration)
	PutRequestMirrorPolicy(value *ComputeRegionUrlMapPathMatcherDefaultRouteActionRequestMirrorPolicy)
	PutRetryPolicy(value *ComputeRegionUrlMapPathMatcherDefaultRouteActionRetryPolicy)
	PutTimeout(value *ComputeRegionUrlMapPathMatcherDefaultRouteActionTimeout)
	PutUrlRewrite(value *ComputeRegionUrlMapPathMatcherDefaultRouteActionUrlRewrite)
	PutWeightedBackendServices(value interface{})
	ResetCorsPolicy()
	ResetFaultInjectionPolicy()
	ResetMaxStreamDuration()
	ResetRequestMirrorPolicy()
	ResetRetryPolicy()
	ResetTimeout()
	ResetUrlRewrite()
	ResetWeightedBackendServices()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference
type jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) CorsPolicy() ComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference {
	var returns ComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicyOutputReference
	_jsii_.Get(
		j,
		"corsPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) CorsPolicyInput() *ComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicy {
	var returns *ComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicy
	_jsii_.Get(
		j,
		"corsPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) FaultInjectionPolicy() ComputeRegionUrlMapPathMatcherDefaultRouteActionFaultInjectionPolicyOutputReference {
	var returns ComputeRegionUrlMapPathMatcherDefaultRouteActionFaultInjectionPolicyOutputReference
	_jsii_.Get(
		j,
		"faultInjectionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) FaultInjectionPolicyInput() *ComputeRegionUrlMapPathMatcherDefaultRouteActionFaultInjectionPolicy {
	var returns *ComputeRegionUrlMapPathMatcherDefaultRouteActionFaultInjectionPolicy
	_jsii_.Get(
		j,
		"faultInjectionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) InternalValue() *ComputeRegionUrlMapPathMatcherDefaultRouteAction {
	var returns *ComputeRegionUrlMapPathMatcherDefaultRouteAction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) MaxStreamDuration() ComputeRegionUrlMapPathMatcherDefaultRouteActionMaxStreamDurationOutputReference {
	var returns ComputeRegionUrlMapPathMatcherDefaultRouteActionMaxStreamDurationOutputReference
	_jsii_.Get(
		j,
		"maxStreamDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) MaxStreamDurationInput() *ComputeRegionUrlMapPathMatcherDefaultRouteActionMaxStreamDuration {
	var returns *ComputeRegionUrlMapPathMatcherDefaultRouteActionMaxStreamDuration
	_jsii_.Get(
		j,
		"maxStreamDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) RequestMirrorPolicy() ComputeRegionUrlMapPathMatcherDefaultRouteActionRequestMirrorPolicyOutputReference {
	var returns ComputeRegionUrlMapPathMatcherDefaultRouteActionRequestMirrorPolicyOutputReference
	_jsii_.Get(
		j,
		"requestMirrorPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) RequestMirrorPolicyInput() *ComputeRegionUrlMapPathMatcherDefaultRouteActionRequestMirrorPolicy {
	var returns *ComputeRegionUrlMapPathMatcherDefaultRouteActionRequestMirrorPolicy
	_jsii_.Get(
		j,
		"requestMirrorPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) RetryPolicy() ComputeRegionUrlMapPathMatcherDefaultRouteActionRetryPolicyOutputReference {
	var returns ComputeRegionUrlMapPathMatcherDefaultRouteActionRetryPolicyOutputReference
	_jsii_.Get(
		j,
		"retryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) RetryPolicyInput() *ComputeRegionUrlMapPathMatcherDefaultRouteActionRetryPolicy {
	var returns *ComputeRegionUrlMapPathMatcherDefaultRouteActionRetryPolicy
	_jsii_.Get(
		j,
		"retryPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) Timeout() ComputeRegionUrlMapPathMatcherDefaultRouteActionTimeoutOutputReference {
	var returns ComputeRegionUrlMapPathMatcherDefaultRouteActionTimeoutOutputReference
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) TimeoutInput() *ComputeRegionUrlMapPathMatcherDefaultRouteActionTimeout {
	var returns *ComputeRegionUrlMapPathMatcherDefaultRouteActionTimeout
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) UrlRewrite() ComputeRegionUrlMapPathMatcherDefaultRouteActionUrlRewriteOutputReference {
	var returns ComputeRegionUrlMapPathMatcherDefaultRouteActionUrlRewriteOutputReference
	_jsii_.Get(
		j,
		"urlRewrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) UrlRewriteInput() *ComputeRegionUrlMapPathMatcherDefaultRouteActionUrlRewrite {
	var returns *ComputeRegionUrlMapPathMatcherDefaultRouteActionUrlRewrite
	_jsii_.Get(
		j,
		"urlRewriteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) WeightedBackendServices() ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesList {
	var returns ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesList
	_jsii_.Get(
		j,
		"weightedBackendServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) WeightedBackendServicesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"weightedBackendServicesInput",
		&returns,
	)
	return returns
}


func NewComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference {
	_init_.Initialize()

	if err := validateNewComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference_Override(c ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference)SetInternalValue(val *ComputeRegionUrlMapPathMatcherDefaultRouteAction) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) PutCorsPolicy(value *ComputeRegionUrlMapPathMatcherDefaultRouteActionCorsPolicy) {
	if err := c.validatePutCorsPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCorsPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) PutFaultInjectionPolicy(value *ComputeRegionUrlMapPathMatcherDefaultRouteActionFaultInjectionPolicy) {
	if err := c.validatePutFaultInjectionPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFaultInjectionPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) PutMaxStreamDuration(value *ComputeRegionUrlMapPathMatcherDefaultRouteActionMaxStreamDuration) {
	if err := c.validatePutMaxStreamDurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMaxStreamDuration",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) PutRequestMirrorPolicy(value *ComputeRegionUrlMapPathMatcherDefaultRouteActionRequestMirrorPolicy) {
	if err := c.validatePutRequestMirrorPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRequestMirrorPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) PutRetryPolicy(value *ComputeRegionUrlMapPathMatcherDefaultRouteActionRetryPolicy) {
	if err := c.validatePutRetryPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRetryPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) PutTimeout(value *ComputeRegionUrlMapPathMatcherDefaultRouteActionTimeout) {
	if err := c.validatePutTimeoutParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeout",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) PutUrlRewrite(value *ComputeRegionUrlMapPathMatcherDefaultRouteActionUrlRewrite) {
	if err := c.validatePutUrlRewriteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putUrlRewrite",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) PutWeightedBackendServices(value interface{}) {
	if err := c.validatePutWeightedBackendServicesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putWeightedBackendServices",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) ResetCorsPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetCorsPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) ResetFaultInjectionPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetFaultInjectionPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) ResetMaxStreamDuration() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxStreamDuration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) ResetRequestMirrorPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestMirrorPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) ResetRetryPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetRetryPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) ResetTimeout() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeout",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) ResetUrlRewrite() {
	_jsii_.InvokeVoid(
		c,
		"resetUrlRewrite",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) ResetWeightedBackendServices() {
	_jsii_.InvokeVoid(
		c,
		"resetWeightedBackendServices",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

