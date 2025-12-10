// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxwebhook

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/dialogflowcxwebhook/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DialogflowCxWebhookGenericWebServiceOutputReference interface {
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
	InternalValue() *DialogflowCxWebhookGenericWebService
	SetInternalValue(val *DialogflowCxWebhookGenericWebService)
	OauthConfig() DialogflowCxWebhookGenericWebServiceOauthConfigOutputReference
	OauthConfigInput() *DialogflowCxWebhookGenericWebServiceOauthConfig
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
	SecretVersionsForRequestHeaders() DialogflowCxWebhookGenericWebServiceSecretVersionsForRequestHeadersList
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutOauthConfig(value *DialogflowCxWebhookGenericWebServiceOauthConfig)
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
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DialogflowCxWebhookGenericWebServiceOutputReference
type jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) AllowedCaCerts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedCaCerts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) AllowedCaCertsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedCaCertsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) HttpMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) HttpMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) InternalValue() *DialogflowCxWebhookGenericWebService {
	var returns *DialogflowCxWebhookGenericWebService
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) OauthConfig() DialogflowCxWebhookGenericWebServiceOauthConfigOutputReference {
	var returns DialogflowCxWebhookGenericWebServiceOauthConfigOutputReference
	_jsii_.Get(
		j,
		"oauthConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) OauthConfigInput() *DialogflowCxWebhookGenericWebServiceOauthConfig {
	var returns *DialogflowCxWebhookGenericWebServiceOauthConfig
	_jsii_.Get(
		j,
		"oauthConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) ParameterMapping() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameterMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) ParameterMappingInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameterMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) RequestBody() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) RequestBodyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) RequestHeaders() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) RequestHeadersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requestHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) SecretVersionForUsernamePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretVersionForUsernamePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) SecretVersionForUsernamePasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretVersionForUsernamePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) SecretVersionsForRequestHeaders() DialogflowCxWebhookGenericWebServiceSecretVersionsForRequestHeadersList {
	var returns DialogflowCxWebhookGenericWebServiceSecretVersionsForRequestHeadersList
	_jsii_.Get(
		j,
		"secretVersionsForRequestHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) SecretVersionsForRequestHeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretVersionsForRequestHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) ServiceAgentAuth() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAgentAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) ServiceAgentAuthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAgentAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) WebhookType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhookType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) WebhookTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webhookTypeInput",
		&returns,
	)
	return returns
}


func NewDialogflowCxWebhookGenericWebServiceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DialogflowCxWebhookGenericWebServiceOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowCxWebhookGenericWebServiceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxWebhook.DialogflowCxWebhookGenericWebServiceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDialogflowCxWebhookGenericWebServiceOutputReference_Override(d DialogflowCxWebhookGenericWebServiceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxWebhook.DialogflowCxWebhookGenericWebServiceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference)SetAllowedCaCerts(val *[]*string) {
	if err := j.validateSetAllowedCaCertsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedCaCerts",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference)SetHttpMethod(val *string) {
	if err := j.validateSetHttpMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpMethod",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference)SetInternalValue(val *DialogflowCxWebhookGenericWebService) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference)SetParameterMapping(val *map[string]*string) {
	if err := j.validateSetParameterMappingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameterMapping",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference)SetRequestBody(val *string) {
	if err := j.validateSetRequestBodyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestBody",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference)SetRequestHeaders(val *map[string]*string) {
	if err := j.validateSetRequestHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestHeaders",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference)SetSecretVersionForUsernamePassword(val *string) {
	if err := j.validateSetSecretVersionForUsernamePasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretVersionForUsernamePassword",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference)SetServiceAgentAuth(val *string) {
	if err := j.validateSetServiceAgentAuthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAgentAuth",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference)SetUri(val *string) {
	if err := j.validateSetUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uri",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference)SetWebhookType(val *string) {
	if err := j.validateSetWebhookTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webhookType",
		val,
	)
}

func (d *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) PutOauthConfig(value *DialogflowCxWebhookGenericWebServiceOauthConfig) {
	if err := d.validatePutOauthConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOauthConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) PutSecretVersionsForRequestHeaders(value interface{}) {
	if err := d.validatePutSecretVersionsForRequestHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSecretVersionsForRequestHeaders",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) ResetAllowedCaCerts() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowedCaCerts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) ResetHttpMethod() {
	_jsii_.InvokeVoid(
		d,
		"resetHttpMethod",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) ResetOauthConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetOauthConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) ResetParameterMapping() {
	_jsii_.InvokeVoid(
		d,
		"resetParameterMapping",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) ResetRequestBody() {
	_jsii_.InvokeVoid(
		d,
		"resetRequestBody",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) ResetRequestHeaders() {
	_jsii_.InvokeVoid(
		d,
		"resetRequestHeaders",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) ResetSecretVersionForUsernamePassword() {
	_jsii_.InvokeVoid(
		d,
		"resetSecretVersionForUsernamePassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) ResetSecretVersionsForRequestHeaders() {
	_jsii_.InvokeVoid(
		d,
		"resetSecretVersionsForRequestHeaders",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) ResetServiceAgentAuth() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceAgentAuth",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) ResetWebhookType() {
	_jsii_.InvokeVoid(
		d,
		"resetWebhookType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxWebhookGenericWebServiceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

