// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computebackendbucket

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/computebackendbucket/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeBackendBucketCdnPolicyOutputReference interface {
	cdktf.ComplexObject
	BypassCacheOnRequestHeaders() ComputeBackendBucketCdnPolicyBypassCacheOnRequestHeadersList
	BypassCacheOnRequestHeadersInput() interface{}
	CacheKeyPolicy() ComputeBackendBucketCdnPolicyCacheKeyPolicyOutputReference
	CacheKeyPolicyInput() *ComputeBackendBucketCdnPolicyCacheKeyPolicy
	CacheMode() *string
	SetCacheMode(val *string)
	CacheModeInput() *string
	ClientTtl() *float64
	SetClientTtl(val *float64)
	ClientTtlInput() *float64
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
	DefaultTtl() *float64
	SetDefaultTtl(val *float64)
	DefaultTtlInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *ComputeBackendBucketCdnPolicy
	SetInternalValue(val *ComputeBackendBucketCdnPolicy)
	MaxTtl() *float64
	SetMaxTtl(val *float64)
	MaxTtlInput() *float64
	NegativeCaching() interface{}
	SetNegativeCaching(val interface{})
	NegativeCachingInput() interface{}
	NegativeCachingPolicy() ComputeBackendBucketCdnPolicyNegativeCachingPolicyList
	NegativeCachingPolicyInput() interface{}
	RequestCoalescing() interface{}
	SetRequestCoalescing(val interface{})
	RequestCoalescingInput() interface{}
	ServeWhileStale() *float64
	SetServeWhileStale(val *float64)
	ServeWhileStaleInput() *float64
	SignedUrlCacheMaxAgeSec() *float64
	SetSignedUrlCacheMaxAgeSec(val *float64)
	SignedUrlCacheMaxAgeSecInput() *float64
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
	PutBypassCacheOnRequestHeaders(value interface{})
	PutCacheKeyPolicy(value *ComputeBackendBucketCdnPolicyCacheKeyPolicy)
	PutNegativeCachingPolicy(value interface{})
	ResetBypassCacheOnRequestHeaders()
	ResetCacheKeyPolicy()
	ResetCacheMode()
	ResetClientTtl()
	ResetDefaultTtl()
	ResetMaxTtl()
	ResetNegativeCaching()
	ResetNegativeCachingPolicy()
	ResetRequestCoalescing()
	ResetServeWhileStale()
	ResetSignedUrlCacheMaxAgeSec()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeBackendBucketCdnPolicyOutputReference
type jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) BypassCacheOnRequestHeaders() ComputeBackendBucketCdnPolicyBypassCacheOnRequestHeadersList {
	var returns ComputeBackendBucketCdnPolicyBypassCacheOnRequestHeadersList
	_jsii_.Get(
		j,
		"bypassCacheOnRequestHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) BypassCacheOnRequestHeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bypassCacheOnRequestHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) CacheKeyPolicy() ComputeBackendBucketCdnPolicyCacheKeyPolicyOutputReference {
	var returns ComputeBackendBucketCdnPolicyCacheKeyPolicyOutputReference
	_jsii_.Get(
		j,
		"cacheKeyPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) CacheKeyPolicyInput() *ComputeBackendBucketCdnPolicyCacheKeyPolicy {
	var returns *ComputeBackendBucketCdnPolicyCacheKeyPolicy
	_jsii_.Get(
		j,
		"cacheKeyPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) CacheMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) CacheModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) ClientTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) ClientTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"clientTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) DefaultTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) DefaultTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) InternalValue() *ComputeBackendBucketCdnPolicy {
	var returns *ComputeBackendBucketCdnPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) MaxTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) MaxTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) NegativeCaching() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negativeCaching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) NegativeCachingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negativeCachingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) NegativeCachingPolicy() ComputeBackendBucketCdnPolicyNegativeCachingPolicyList {
	var returns ComputeBackendBucketCdnPolicyNegativeCachingPolicyList
	_jsii_.Get(
		j,
		"negativeCachingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) NegativeCachingPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negativeCachingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) RequestCoalescing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestCoalescing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) RequestCoalescingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestCoalescingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) ServeWhileStale() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"serveWhileStale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) ServeWhileStaleInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"serveWhileStaleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) SignedUrlCacheMaxAgeSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"signedUrlCacheMaxAgeSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) SignedUrlCacheMaxAgeSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"signedUrlCacheMaxAgeSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeBackendBucketCdnPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeBackendBucketCdnPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewComputeBackendBucketCdnPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeBackendBucket.ComputeBackendBucketCdnPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeBackendBucketCdnPolicyOutputReference_Override(c ComputeBackendBucketCdnPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeBackendBucket.ComputeBackendBucketCdnPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference)SetCacheMode(val *string) {
	if err := j.validateSetCacheModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheMode",
		val,
	)
}

func (j *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference)SetClientTtl(val *float64) {
	if err := j.validateSetClientTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientTtl",
		val,
	)
}

func (j *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference)SetDefaultTtl(val *float64) {
	if err := j.validateSetDefaultTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultTtl",
		val,
	)
}

func (j *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference)SetInternalValue(val *ComputeBackendBucketCdnPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference)SetMaxTtl(val *float64) {
	if err := j.validateSetMaxTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxTtl",
		val,
	)
}

func (j *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference)SetNegativeCaching(val interface{}) {
	if err := j.validateSetNegativeCachingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"negativeCaching",
		val,
	)
}

func (j *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference)SetRequestCoalescing(val interface{}) {
	if err := j.validateSetRequestCoalescingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestCoalescing",
		val,
	)
}

func (j *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference)SetServeWhileStale(val *float64) {
	if err := j.validateSetServeWhileStaleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serveWhileStale",
		val,
	)
}

func (j *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference)SetSignedUrlCacheMaxAgeSec(val *float64) {
	if err := j.validateSetSignedUrlCacheMaxAgeSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signedUrlCacheMaxAgeSec",
		val,
	)
}

func (j *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) PutBypassCacheOnRequestHeaders(value interface{}) {
	if err := c.validatePutBypassCacheOnRequestHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBypassCacheOnRequestHeaders",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) PutCacheKeyPolicy(value *ComputeBackendBucketCdnPolicyCacheKeyPolicy) {
	if err := c.validatePutCacheKeyPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCacheKeyPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) PutNegativeCachingPolicy(value interface{}) {
	if err := c.validatePutNegativeCachingPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNegativeCachingPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) ResetBypassCacheOnRequestHeaders() {
	_jsii_.InvokeVoid(
		c,
		"resetBypassCacheOnRequestHeaders",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) ResetCacheKeyPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetCacheKeyPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) ResetCacheMode() {
	_jsii_.InvokeVoid(
		c,
		"resetCacheMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) ResetClientTtl() {
	_jsii_.InvokeVoid(
		c,
		"resetClientTtl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) ResetDefaultTtl() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultTtl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) ResetMaxTtl() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxTtl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) ResetNegativeCaching() {
	_jsii_.InvokeVoid(
		c,
		"resetNegativeCaching",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) ResetNegativeCachingPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetNegativeCachingPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) ResetRequestCoalescing() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestCoalescing",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) ResetServeWhileStale() {
	_jsii_.InvokeVoid(
		c,
		"resetServeWhileStale",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) ResetSignedUrlCacheMaxAgeSec() {
	_jsii_.InvokeVoid(
		c,
		"resetSignedUrlCacheMaxAgeSec",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeBackendBucketCdnPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

