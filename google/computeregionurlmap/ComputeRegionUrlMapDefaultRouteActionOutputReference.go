package computeregionurlmap

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v7/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v7/computeregionurlmap/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeRegionUrlMapDefaultRouteActionOutputReference interface {
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
	CorsPolicy() ComputeRegionUrlMapDefaultRouteActionCorsPolicyOutputReference
	CorsPolicyInput() *ComputeRegionUrlMapDefaultRouteActionCorsPolicy
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	FaultInjectionPolicy() ComputeRegionUrlMapDefaultRouteActionFaultInjectionPolicyOutputReference
	FaultInjectionPolicyInput() *ComputeRegionUrlMapDefaultRouteActionFaultInjectionPolicy
	// Experimental.
	Fqn() *string
	InternalValue() *ComputeRegionUrlMapDefaultRouteAction
	SetInternalValue(val *ComputeRegionUrlMapDefaultRouteAction)
	RequestMirrorPolicy() ComputeRegionUrlMapDefaultRouteActionRequestMirrorPolicyOutputReference
	RequestMirrorPolicyInput() *ComputeRegionUrlMapDefaultRouteActionRequestMirrorPolicy
	RetryPolicy() ComputeRegionUrlMapDefaultRouteActionRetryPolicyOutputReference
	RetryPolicyInput() *ComputeRegionUrlMapDefaultRouteActionRetryPolicy
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timeout() ComputeRegionUrlMapDefaultRouteActionTimeoutOutputReference
	TimeoutInput() *ComputeRegionUrlMapDefaultRouteActionTimeout
	UrlRewrite() ComputeRegionUrlMapDefaultRouteActionUrlRewriteOutputReference
	UrlRewriteInput() *ComputeRegionUrlMapDefaultRouteActionUrlRewrite
	WeightedBackendServices() ComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesList
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
	PutCorsPolicy(value *ComputeRegionUrlMapDefaultRouteActionCorsPolicy)
	PutFaultInjectionPolicy(value *ComputeRegionUrlMapDefaultRouteActionFaultInjectionPolicy)
	PutRequestMirrorPolicy(value *ComputeRegionUrlMapDefaultRouteActionRequestMirrorPolicy)
	PutRetryPolicy(value *ComputeRegionUrlMapDefaultRouteActionRetryPolicy)
	PutTimeout(value *ComputeRegionUrlMapDefaultRouteActionTimeout)
	PutUrlRewrite(value *ComputeRegionUrlMapDefaultRouteActionUrlRewrite)
	PutWeightedBackendServices(value interface{})
	ResetCorsPolicy()
	ResetFaultInjectionPolicy()
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

// The jsii proxy struct for ComputeRegionUrlMapDefaultRouteActionOutputReference
type jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) CorsPolicy() ComputeRegionUrlMapDefaultRouteActionCorsPolicyOutputReference {
	var returns ComputeRegionUrlMapDefaultRouteActionCorsPolicyOutputReference
	_jsii_.Get(
		j,
		"corsPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) CorsPolicyInput() *ComputeRegionUrlMapDefaultRouteActionCorsPolicy {
	var returns *ComputeRegionUrlMapDefaultRouteActionCorsPolicy
	_jsii_.Get(
		j,
		"corsPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) FaultInjectionPolicy() ComputeRegionUrlMapDefaultRouteActionFaultInjectionPolicyOutputReference {
	var returns ComputeRegionUrlMapDefaultRouteActionFaultInjectionPolicyOutputReference
	_jsii_.Get(
		j,
		"faultInjectionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) FaultInjectionPolicyInput() *ComputeRegionUrlMapDefaultRouteActionFaultInjectionPolicy {
	var returns *ComputeRegionUrlMapDefaultRouteActionFaultInjectionPolicy
	_jsii_.Get(
		j,
		"faultInjectionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) InternalValue() *ComputeRegionUrlMapDefaultRouteAction {
	var returns *ComputeRegionUrlMapDefaultRouteAction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) RequestMirrorPolicy() ComputeRegionUrlMapDefaultRouteActionRequestMirrorPolicyOutputReference {
	var returns ComputeRegionUrlMapDefaultRouteActionRequestMirrorPolicyOutputReference
	_jsii_.Get(
		j,
		"requestMirrorPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) RequestMirrorPolicyInput() *ComputeRegionUrlMapDefaultRouteActionRequestMirrorPolicy {
	var returns *ComputeRegionUrlMapDefaultRouteActionRequestMirrorPolicy
	_jsii_.Get(
		j,
		"requestMirrorPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) RetryPolicy() ComputeRegionUrlMapDefaultRouteActionRetryPolicyOutputReference {
	var returns ComputeRegionUrlMapDefaultRouteActionRetryPolicyOutputReference
	_jsii_.Get(
		j,
		"retryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) RetryPolicyInput() *ComputeRegionUrlMapDefaultRouteActionRetryPolicy {
	var returns *ComputeRegionUrlMapDefaultRouteActionRetryPolicy
	_jsii_.Get(
		j,
		"retryPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) Timeout() ComputeRegionUrlMapDefaultRouteActionTimeoutOutputReference {
	var returns ComputeRegionUrlMapDefaultRouteActionTimeoutOutputReference
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) TimeoutInput() *ComputeRegionUrlMapDefaultRouteActionTimeout {
	var returns *ComputeRegionUrlMapDefaultRouteActionTimeout
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) UrlRewrite() ComputeRegionUrlMapDefaultRouteActionUrlRewriteOutputReference {
	var returns ComputeRegionUrlMapDefaultRouteActionUrlRewriteOutputReference
	_jsii_.Get(
		j,
		"urlRewrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) UrlRewriteInput() *ComputeRegionUrlMapDefaultRouteActionUrlRewrite {
	var returns *ComputeRegionUrlMapDefaultRouteActionUrlRewrite
	_jsii_.Get(
		j,
		"urlRewriteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) WeightedBackendServices() ComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesList {
	var returns ComputeRegionUrlMapDefaultRouteActionWeightedBackendServicesList
	_jsii_.Get(
		j,
		"weightedBackendServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) WeightedBackendServicesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"weightedBackendServicesInput",
		&returns,
	)
	return returns
}


func NewComputeRegionUrlMapDefaultRouteActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeRegionUrlMapDefaultRouteActionOutputReference {
	_init_.Initialize()

	if err := validateNewComputeRegionUrlMapDefaultRouteActionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapDefaultRouteActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeRegionUrlMapDefaultRouteActionOutputReference_Override(c ComputeRegionUrlMapDefaultRouteActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapDefaultRouteActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference)SetInternalValue(val *ComputeRegionUrlMapDefaultRouteAction) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) PutCorsPolicy(value *ComputeRegionUrlMapDefaultRouteActionCorsPolicy) {
	if err := c.validatePutCorsPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCorsPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) PutFaultInjectionPolicy(value *ComputeRegionUrlMapDefaultRouteActionFaultInjectionPolicy) {
	if err := c.validatePutFaultInjectionPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFaultInjectionPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) PutRequestMirrorPolicy(value *ComputeRegionUrlMapDefaultRouteActionRequestMirrorPolicy) {
	if err := c.validatePutRequestMirrorPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRequestMirrorPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) PutRetryPolicy(value *ComputeRegionUrlMapDefaultRouteActionRetryPolicy) {
	if err := c.validatePutRetryPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRetryPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) PutTimeout(value *ComputeRegionUrlMapDefaultRouteActionTimeout) {
	if err := c.validatePutTimeoutParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeout",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) PutUrlRewrite(value *ComputeRegionUrlMapDefaultRouteActionUrlRewrite) {
	if err := c.validatePutUrlRewriteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putUrlRewrite",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) PutWeightedBackendServices(value interface{}) {
	if err := c.validatePutWeightedBackendServicesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putWeightedBackendServices",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) ResetCorsPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetCorsPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) ResetFaultInjectionPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetFaultInjectionPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) ResetRequestMirrorPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestMirrorPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) ResetRetryPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetRetryPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) ResetTimeout() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeout",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) ResetUrlRewrite() {
	_jsii_.InvokeVoid(
		c,
		"resetUrlRewrite",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) ResetWeightedBackendServices() {
	_jsii_.InvokeVoid(
		c,
		"resetWeightedBackendServices",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeRegionUrlMapDefaultRouteActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

