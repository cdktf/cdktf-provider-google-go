// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudtasksqueue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/cloudtasksqueue/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudTasksQueueHttpTargetOutputReference interface {
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
	HeaderOverrides() CloudTasksQueueHttpTargetHeaderOverridesList
	HeaderOverridesInput() interface{}
	HttpMethod() *string
	SetHttpMethod(val *string)
	HttpMethodInput() *string
	InternalValue() *CloudTasksQueueHttpTarget
	SetInternalValue(val *CloudTasksQueueHttpTarget)
	OauthToken() CloudTasksQueueHttpTargetOauthTokenOutputReference
	OauthTokenInput() *CloudTasksQueueHttpTargetOauthToken
	OidcToken() CloudTasksQueueHttpTargetOidcTokenOutputReference
	OidcTokenInput() *CloudTasksQueueHttpTargetOidcToken
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UriOverride() CloudTasksQueueHttpTargetUriOverrideOutputReference
	UriOverrideInput() *CloudTasksQueueHttpTargetUriOverride
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
	PutHeaderOverrides(value interface{})
	PutOauthToken(value *CloudTasksQueueHttpTargetOauthToken)
	PutOidcToken(value *CloudTasksQueueHttpTargetOidcToken)
	PutUriOverride(value *CloudTasksQueueHttpTargetUriOverride)
	ResetHeaderOverrides()
	ResetHttpMethod()
	ResetOauthToken()
	ResetOidcToken()
	ResetUriOverride()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudTasksQueueHttpTargetOutputReference
type jsiiProxy_CloudTasksQueueHttpTargetOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetOutputReference) HeaderOverrides() CloudTasksQueueHttpTargetHeaderOverridesList {
	var returns CloudTasksQueueHttpTargetHeaderOverridesList
	_jsii_.Get(
		j,
		"headerOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetOutputReference) HeaderOverridesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headerOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetOutputReference) HttpMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetOutputReference) HttpMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetOutputReference) InternalValue() *CloudTasksQueueHttpTarget {
	var returns *CloudTasksQueueHttpTarget
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetOutputReference) OauthToken() CloudTasksQueueHttpTargetOauthTokenOutputReference {
	var returns CloudTasksQueueHttpTargetOauthTokenOutputReference
	_jsii_.Get(
		j,
		"oauthToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetOutputReference) OauthTokenInput() *CloudTasksQueueHttpTargetOauthToken {
	var returns *CloudTasksQueueHttpTargetOauthToken
	_jsii_.Get(
		j,
		"oauthTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetOutputReference) OidcToken() CloudTasksQueueHttpTargetOidcTokenOutputReference {
	var returns CloudTasksQueueHttpTargetOidcTokenOutputReference
	_jsii_.Get(
		j,
		"oidcToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetOutputReference) OidcTokenInput() *CloudTasksQueueHttpTargetOidcToken {
	var returns *CloudTasksQueueHttpTargetOidcToken
	_jsii_.Get(
		j,
		"oidcTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetOutputReference) UriOverride() CloudTasksQueueHttpTargetUriOverrideOutputReference {
	var returns CloudTasksQueueHttpTargetUriOverrideOutputReference
	_jsii_.Get(
		j,
		"uriOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetOutputReference) UriOverrideInput() *CloudTasksQueueHttpTargetUriOverride {
	var returns *CloudTasksQueueHttpTargetUriOverride
	_jsii_.Get(
		j,
		"uriOverrideInput",
		&returns,
	)
	return returns
}


func NewCloudTasksQueueHttpTargetOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CloudTasksQueueHttpTargetOutputReference {
	_init_.Initialize()

	if err := validateNewCloudTasksQueueHttpTargetOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudTasksQueueHttpTargetOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cloudTasksQueue.CloudTasksQueueHttpTargetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudTasksQueueHttpTargetOutputReference_Override(c CloudTasksQueueHttpTargetOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cloudTasksQueue.CloudTasksQueueHttpTargetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetOutputReference)SetHttpMethod(val *string) {
	if err := j.validateSetHttpMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpMethod",
		val,
	)
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetOutputReference)SetInternalValue(val *CloudTasksQueueHttpTarget) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudTasksQueueHttpTargetOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudTasksQueueHttpTargetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudTasksQueueHttpTargetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudTasksQueueHttpTargetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudTasksQueueHttpTargetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudTasksQueueHttpTargetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudTasksQueueHttpTargetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudTasksQueueHttpTargetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudTasksQueueHttpTargetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudTasksQueueHttpTargetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudTasksQueueHttpTargetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudTasksQueueHttpTargetOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudTasksQueueHttpTargetOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudTasksQueueHttpTargetOutputReference) PutHeaderOverrides(value interface{}) {
	if err := c.validatePutHeaderOverridesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHeaderOverrides",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudTasksQueueHttpTargetOutputReference) PutOauthToken(value *CloudTasksQueueHttpTargetOauthToken) {
	if err := c.validatePutOauthTokenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOauthToken",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudTasksQueueHttpTargetOutputReference) PutOidcToken(value *CloudTasksQueueHttpTargetOidcToken) {
	if err := c.validatePutOidcTokenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOidcToken",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudTasksQueueHttpTargetOutputReference) PutUriOverride(value *CloudTasksQueueHttpTargetUriOverride) {
	if err := c.validatePutUriOverrideParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putUriOverride",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudTasksQueueHttpTargetOutputReference) ResetHeaderOverrides() {
	_jsii_.InvokeVoid(
		c,
		"resetHeaderOverrides",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudTasksQueueHttpTargetOutputReference) ResetHttpMethod() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpMethod",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudTasksQueueHttpTargetOutputReference) ResetOauthToken() {
	_jsii_.InvokeVoid(
		c,
		"resetOauthToken",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudTasksQueueHttpTargetOutputReference) ResetOidcToken() {
	_jsii_.InvokeVoid(
		c,
		"resetOidcToken",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudTasksQueueHttpTargetOutputReference) ResetUriOverride() {
	_jsii_.InvokeVoid(
		c,
		"resetUriOverride",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudTasksQueueHttpTargetOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CloudTasksQueueHttpTargetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

