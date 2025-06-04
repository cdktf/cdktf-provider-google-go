// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationsauthconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/integrationsauthconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IntegrationsAuthConfigDecryptedCredentialOutputReference interface {
	cdktf.ComplexObject
	AuthToken() IntegrationsAuthConfigDecryptedCredentialAuthTokenOutputReference
	AuthTokenInput() *IntegrationsAuthConfigDecryptedCredentialAuthToken
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
	CredentialType() *string
	SetCredentialType(val *string)
	CredentialTypeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *IntegrationsAuthConfigDecryptedCredential
	SetInternalValue(val *IntegrationsAuthConfigDecryptedCredential)
	Jwt() IntegrationsAuthConfigDecryptedCredentialJwtOutputReference
	JwtInput() *IntegrationsAuthConfigDecryptedCredentialJwt
	Oauth2AuthorizationCode() IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeOutputReference
	Oauth2AuthorizationCodeInput() *IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCode
	Oauth2ClientCredentials() IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsOutputReference
	Oauth2ClientCredentialsInput() *IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentials
	OidcToken() IntegrationsAuthConfigDecryptedCredentialOidcTokenOutputReference
	OidcTokenInput() *IntegrationsAuthConfigDecryptedCredentialOidcToken
	ServiceAccountCredentials() IntegrationsAuthConfigDecryptedCredentialServiceAccountCredentialsOutputReference
	ServiceAccountCredentialsInput() *IntegrationsAuthConfigDecryptedCredentialServiceAccountCredentials
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UsernameAndPassword() IntegrationsAuthConfigDecryptedCredentialUsernameAndPasswordOutputReference
	UsernameAndPasswordInput() *IntegrationsAuthConfigDecryptedCredentialUsernameAndPassword
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
	PutAuthToken(value *IntegrationsAuthConfigDecryptedCredentialAuthToken)
	PutJwt(value *IntegrationsAuthConfigDecryptedCredentialJwt)
	PutOauth2AuthorizationCode(value *IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCode)
	PutOauth2ClientCredentials(value *IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentials)
	PutOidcToken(value *IntegrationsAuthConfigDecryptedCredentialOidcToken)
	PutServiceAccountCredentials(value *IntegrationsAuthConfigDecryptedCredentialServiceAccountCredentials)
	PutUsernameAndPassword(value *IntegrationsAuthConfigDecryptedCredentialUsernameAndPassword)
	ResetAuthToken()
	ResetJwt()
	ResetOauth2AuthorizationCode()
	ResetOauth2ClientCredentials()
	ResetOidcToken()
	ResetServiceAccountCredentials()
	ResetUsernameAndPassword()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IntegrationsAuthConfigDecryptedCredentialOutputReference
type jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) AuthToken() IntegrationsAuthConfigDecryptedCredentialAuthTokenOutputReference {
	var returns IntegrationsAuthConfigDecryptedCredentialAuthTokenOutputReference
	_jsii_.Get(
		j,
		"authToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) AuthTokenInput() *IntegrationsAuthConfigDecryptedCredentialAuthToken {
	var returns *IntegrationsAuthConfigDecryptedCredentialAuthToken
	_jsii_.Get(
		j,
		"authTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) CredentialType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) CredentialTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) InternalValue() *IntegrationsAuthConfigDecryptedCredential {
	var returns *IntegrationsAuthConfigDecryptedCredential
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) Jwt() IntegrationsAuthConfigDecryptedCredentialJwtOutputReference {
	var returns IntegrationsAuthConfigDecryptedCredentialJwtOutputReference
	_jsii_.Get(
		j,
		"jwt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) JwtInput() *IntegrationsAuthConfigDecryptedCredentialJwt {
	var returns *IntegrationsAuthConfigDecryptedCredentialJwt
	_jsii_.Get(
		j,
		"jwtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) Oauth2AuthorizationCode() IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeOutputReference {
	var returns IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCodeOutputReference
	_jsii_.Get(
		j,
		"oauth2AuthorizationCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) Oauth2AuthorizationCodeInput() *IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCode {
	var returns *IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCode
	_jsii_.Get(
		j,
		"oauth2AuthorizationCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) Oauth2ClientCredentials() IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsOutputReference {
	var returns IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentialsOutputReference
	_jsii_.Get(
		j,
		"oauth2ClientCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) Oauth2ClientCredentialsInput() *IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentials {
	var returns *IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentials
	_jsii_.Get(
		j,
		"oauth2ClientCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) OidcToken() IntegrationsAuthConfigDecryptedCredentialOidcTokenOutputReference {
	var returns IntegrationsAuthConfigDecryptedCredentialOidcTokenOutputReference
	_jsii_.Get(
		j,
		"oidcToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) OidcTokenInput() *IntegrationsAuthConfigDecryptedCredentialOidcToken {
	var returns *IntegrationsAuthConfigDecryptedCredentialOidcToken
	_jsii_.Get(
		j,
		"oidcTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) ServiceAccountCredentials() IntegrationsAuthConfigDecryptedCredentialServiceAccountCredentialsOutputReference {
	var returns IntegrationsAuthConfigDecryptedCredentialServiceAccountCredentialsOutputReference
	_jsii_.Get(
		j,
		"serviceAccountCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) ServiceAccountCredentialsInput() *IntegrationsAuthConfigDecryptedCredentialServiceAccountCredentials {
	var returns *IntegrationsAuthConfigDecryptedCredentialServiceAccountCredentials
	_jsii_.Get(
		j,
		"serviceAccountCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) UsernameAndPassword() IntegrationsAuthConfigDecryptedCredentialUsernameAndPasswordOutputReference {
	var returns IntegrationsAuthConfigDecryptedCredentialUsernameAndPasswordOutputReference
	_jsii_.Get(
		j,
		"usernameAndPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) UsernameAndPasswordInput() *IntegrationsAuthConfigDecryptedCredentialUsernameAndPassword {
	var returns *IntegrationsAuthConfigDecryptedCredentialUsernameAndPassword
	_jsii_.Get(
		j,
		"usernameAndPasswordInput",
		&returns,
	)
	return returns
}


func NewIntegrationsAuthConfigDecryptedCredentialOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IntegrationsAuthConfigDecryptedCredentialOutputReference {
	_init_.Initialize()

	if err := validateNewIntegrationsAuthConfigDecryptedCredentialOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.integrationsAuthConfig.IntegrationsAuthConfigDecryptedCredentialOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIntegrationsAuthConfigDecryptedCredentialOutputReference_Override(i IntegrationsAuthConfigDecryptedCredentialOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.integrationsAuthConfig.IntegrationsAuthConfigDecryptedCredentialOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference)SetCredentialType(val *string) {
	if err := j.validateSetCredentialTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"credentialType",
		val,
	)
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference)SetInternalValue(val *IntegrationsAuthConfigDecryptedCredential) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) PutAuthToken(value *IntegrationsAuthConfigDecryptedCredentialAuthToken) {
	if err := i.validatePutAuthTokenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAuthToken",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) PutJwt(value *IntegrationsAuthConfigDecryptedCredentialJwt) {
	if err := i.validatePutJwtParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putJwt",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) PutOauth2AuthorizationCode(value *IntegrationsAuthConfigDecryptedCredentialOauth2AuthorizationCode) {
	if err := i.validatePutOauth2AuthorizationCodeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putOauth2AuthorizationCode",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) PutOauth2ClientCredentials(value *IntegrationsAuthConfigDecryptedCredentialOauth2ClientCredentials) {
	if err := i.validatePutOauth2ClientCredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putOauth2ClientCredentials",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) PutOidcToken(value *IntegrationsAuthConfigDecryptedCredentialOidcToken) {
	if err := i.validatePutOidcTokenParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putOidcToken",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) PutServiceAccountCredentials(value *IntegrationsAuthConfigDecryptedCredentialServiceAccountCredentials) {
	if err := i.validatePutServiceAccountCredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putServiceAccountCredentials",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) PutUsernameAndPassword(value *IntegrationsAuthConfigDecryptedCredentialUsernameAndPassword) {
	if err := i.validatePutUsernameAndPasswordParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putUsernameAndPassword",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) ResetAuthToken() {
	_jsii_.InvokeVoid(
		i,
		"resetAuthToken",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) ResetJwt() {
	_jsii_.InvokeVoid(
		i,
		"resetJwt",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) ResetOauth2AuthorizationCode() {
	_jsii_.InvokeVoid(
		i,
		"resetOauth2AuthorizationCode",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) ResetOauth2ClientCredentials() {
	_jsii_.InvokeVoid(
		i,
		"resetOauth2ClientCredentials",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) ResetOidcToken() {
	_jsii_.InvokeVoid(
		i,
		"resetOidcToken",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) ResetServiceAccountCredentials() {
	_jsii_.InvokeVoid(
		i,
		"resetServiceAccountCredentials",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) ResetUsernameAndPassword() {
	_jsii_.InvokeVoid(
		i,
		"resetUsernameAndPassword",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := i.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationsAuthConfigDecryptedCredentialOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

