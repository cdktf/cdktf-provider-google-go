// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxwebhook

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/dialogflowcxwebhook/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference interface {
	cdktf.ComplexObject
	AllowedCaCerts() *[]*string
	SetAllowedCaCerts(val *[]*string)
	AllowedCaCertsInput() *[]*string
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
	HttpMethod() *string
	SetHttpMethod(val *string)
	HttpMethodInput() *string
	InternalValue() *DialogflowCxWebhookServiceDirectoryGenericWebService
	SetInternalValue(val *DialogflowCxWebhookServiceDirectoryGenericWebService)
	OauthConfig() DialogflowCxWebhookServiceDirectoryGenericWebServiceOauthConfigOutputReference
	OauthConfigInput() *DialogflowCxWebhookServiceDirectoryGenericWebServiceOauthConfig
	ParameterMapping() *map[string]*string
	SetParameterMapping(val *map[string]*string)
	ParameterMappingInput() *map[string]*string
	RequestBody() *string
	SetRequestBody(val *string)
	RequestBodyInput() *string
	RequestHeaders() *map[string]*string
	SetRequestHeaders(val *map[string]*string)
	RequestHeadersInput() *map[string]*string
	SecretVersionForUsernamePassword() *string
	SetSecretVersionForUsernamePassword(val *string)
	SecretVersionForUsernamePasswordInput() *string
	SecretVersionsForRequestHeaders() DialogflowCxWebhookServiceDirectoryGenericWebServiceSecretVersionsForRequestHeadersList
	SecretVersionsForRequestHeadersInput() interface{}
	ServiceAgentAuth() *string
	SetServiceAgentAuth(val *string)
	ServiceAgentAuthInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Uri() *string
	SetUri(val *string)
	UriInput() *string
	WebhookType() *string
	SetWebhookType(val *string)
	WebhookTypeInput() *string
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
	PutOauthConfig(value *DialogflowCxWebhookServiceDirectoryGenericWebServiceOauthConfig)
	PutSecretVersionsForRequestHeaders(value interface{})
	ResetAllowedCaCerts()
	ResetHttpMethod()
	ResetOauthConfig()
	ResetParameterMapping()
	ResetRequestBody()
	ResetRequestHeaders()
	ResetSecretVersionForUsernamePassword()
	ResetSecretVersionsForRequestHeaders()
	ResetServiceAgentAuth()
	ResetWebhookType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference
type jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) AllowedCaCerts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedCaCerts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) AllowedCaCertsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedCaCertsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) HttpMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) HttpMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) InternalValue() *DialogflowCxWebhookServiceDirectoryGenericWebService {
	var returns *DialogflowCxWebhookServiceDirectoryGenericWebService
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) OauthConfig() DialogflowCxWebhookServiceDirectoryGenericWebServiceOauthConfigOutputReference {
	var returns DialogflowCxWebhookServiceDirectoryGenericWebServiceOauthConfigOutputReference
	_jsii_.Get(
		j,
		"oauthConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) OauthConfigInput() *DialogflowCxWebhookServiceDirectoryGenericWebServiceOauthConfig {
	var returns *DialogflowCxWebhookServiceDirectoryGenericWebServiceOauthConfig
	_jsii_.Get(
		j,
		"oauthConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) ParameterMapping() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameterMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) ParameterMappingInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameterMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) RequestBody() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) RequestBodyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) RequestHeaders() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) RequestHeadersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) SecretVersionForUsernamePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretVersionForUsernamePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) SecretVersionForUsernamePasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretVersionForUsernamePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) SecretVersionsForRequestHeaders() DialogflowCxWebhookServiceDirectoryGenericWebServiceSecretVersionsForRequestHeadersList {
	var returns DialogflowCxWebhookServiceDirectoryGenericWebServiceSecretVersionsForRequestHeadersList
	_jsii_.Get(
		j,
		"secretVersionsForRequestHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) SecretVersionsForRequestHeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretVersionsForRequestHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) ServiceAgentAuth() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAgentAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) ServiceAgentAuthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAgentAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) WebhookType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhookType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) WebhookTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhookTypeInput",
		&returns,
	)
	return returns
}


func NewDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxWebhook.DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference_Override(d DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxWebhook.DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference)SetAllowedCaCerts(val *[]*string) {
	if err := j.validateSetAllowedCaCertsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedCaCerts",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference)SetHttpMethod(val *string) {
	if err := j.validateSetHttpMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpMethod",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference)SetInternalValue(val *DialogflowCxWebhookServiceDirectoryGenericWebService) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference)SetParameterMapping(val *map[string]*string) {
	if err := j.validateSetParameterMappingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameterMapping",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference)SetRequestBody(val *string) {
	if err := j.validateSetRequestBodyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestBody",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference)SetRequestHeaders(val *map[string]*string) {
	if err := j.validateSetRequestHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestHeaders",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference)SetSecretVersionForUsernamePassword(val *string) {
	if err := j.validateSetSecretVersionForUsernamePasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretVersionForUsernamePassword",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference)SetServiceAgentAuth(val *string) {
	if err := j.validateSetServiceAgentAuthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAgentAuth",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference)SetUri(val *string) {
	if err := j.validateSetUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uri",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference)SetWebhookType(val *string) {
	if err := j.validateSetWebhookTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webhookType",
		val,
	)
}

func (d *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) PutOauthConfig(value *DialogflowCxWebhookServiceDirectoryGenericWebServiceOauthConfig) {
	if err := d.validatePutOauthConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOauthConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) PutSecretVersionsForRequestHeaders(value interface{}) {
	if err := d.validatePutSecretVersionsForRequestHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSecretVersionsForRequestHeaders",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) ResetAllowedCaCerts() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowedCaCerts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) ResetHttpMethod() {
	_jsii_.InvokeVoid(
		d,
		"resetHttpMethod",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) ResetOauthConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetOauthConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) ResetParameterMapping() {
	_jsii_.InvokeVoid(
		d,
		"resetParameterMapping",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) ResetRequestBody() {
	_jsii_.InvokeVoid(
		d,
		"resetRequestBody",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) ResetRequestHeaders() {
	_jsii_.InvokeVoid(
		d,
		"resetRequestHeaders",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) ResetSecretVersionForUsernamePassword() {
	_jsii_.InvokeVoid(
		d,
		"resetSecretVersionForUsernamePassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) ResetSecretVersionsForRequestHeaders() {
	_jsii_.InvokeVoid(
		d,
		"resetSecretVersionsForRequestHeaders",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) ResetServiceAgentAuth() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceAgentAuth",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) ResetWebhookType() {
	_jsii_.InvokeVoid(
		d,
		"resetWebhookType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

