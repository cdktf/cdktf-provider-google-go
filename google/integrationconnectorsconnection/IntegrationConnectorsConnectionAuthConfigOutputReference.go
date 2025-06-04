// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationconnectorsconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/integrationconnectorsconnection/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IntegrationConnectorsConnectionAuthConfigOutputReference interface {
	cdktf.ComplexObject
	AdditionalVariable() IntegrationConnectorsConnectionAuthConfigAdditionalVariableList
	AdditionalVariableInput() interface{}
	AuthKey() *string
	SetAuthKey(val *string)
	AuthKeyInput() *string
	AuthType() *string
	SetAuthType(val *string)
	AuthTypeInput() *string
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
	InternalValue() *IntegrationConnectorsConnectionAuthConfig
	SetInternalValue(val *IntegrationConnectorsConnectionAuthConfig)
	Oauth2AuthCodeFlow() IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference
	Oauth2AuthCodeFlowInput() *IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlow
	Oauth2ClientCredentials() IntegrationConnectorsConnectionAuthConfigOauth2ClientCredentialsOutputReference
	Oauth2ClientCredentialsInput() *IntegrationConnectorsConnectionAuthConfigOauth2ClientCredentials
	Oauth2JwtBearer() IntegrationConnectorsConnectionAuthConfigOauth2JwtBearerOutputReference
	Oauth2JwtBearerInput() *IntegrationConnectorsConnectionAuthConfigOauth2JwtBearer
	SshPublicKey() IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference
	SshPublicKeyInput() *IntegrationConnectorsConnectionAuthConfigSshPublicKey
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserPassword() IntegrationConnectorsConnectionAuthConfigUserPasswordOutputReference
	UserPasswordInput() *IntegrationConnectorsConnectionAuthConfigUserPassword
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
	PutAdditionalVariable(value interface{})
	PutOauth2AuthCodeFlow(value *IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlow)
	PutOauth2ClientCredentials(value *IntegrationConnectorsConnectionAuthConfigOauth2ClientCredentials)
	PutOauth2JwtBearer(value *IntegrationConnectorsConnectionAuthConfigOauth2JwtBearer)
	PutSshPublicKey(value *IntegrationConnectorsConnectionAuthConfigSshPublicKey)
	PutUserPassword(value *IntegrationConnectorsConnectionAuthConfigUserPassword)
	ResetAdditionalVariable()
	ResetAuthKey()
	ResetOauth2AuthCodeFlow()
	ResetOauth2ClientCredentials()
	ResetOauth2JwtBearer()
	ResetSshPublicKey()
	ResetUserPassword()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IntegrationConnectorsConnectionAuthConfigOutputReference
type jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) AdditionalVariable() IntegrationConnectorsConnectionAuthConfigAdditionalVariableList {
	var returns IntegrationConnectorsConnectionAuthConfigAdditionalVariableList
	_jsii_.Get(
		j,
		"additionalVariable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) AdditionalVariableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalVariableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) AuthKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) AuthKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) AuthType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) AuthTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) InternalValue() *IntegrationConnectorsConnectionAuthConfig {
	var returns *IntegrationConnectorsConnectionAuthConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) Oauth2AuthCodeFlow() IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference {
	var returns IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference
	_jsii_.Get(
		j,
		"oauth2AuthCodeFlow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) Oauth2AuthCodeFlowInput() *IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlow {
	var returns *IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlow
	_jsii_.Get(
		j,
		"oauth2AuthCodeFlowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) Oauth2ClientCredentials() IntegrationConnectorsConnectionAuthConfigOauth2ClientCredentialsOutputReference {
	var returns IntegrationConnectorsConnectionAuthConfigOauth2ClientCredentialsOutputReference
	_jsii_.Get(
		j,
		"oauth2ClientCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) Oauth2ClientCredentialsInput() *IntegrationConnectorsConnectionAuthConfigOauth2ClientCredentials {
	var returns *IntegrationConnectorsConnectionAuthConfigOauth2ClientCredentials
	_jsii_.Get(
		j,
		"oauth2ClientCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) Oauth2JwtBearer() IntegrationConnectorsConnectionAuthConfigOauth2JwtBearerOutputReference {
	var returns IntegrationConnectorsConnectionAuthConfigOauth2JwtBearerOutputReference
	_jsii_.Get(
		j,
		"oauth2JwtBearer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) Oauth2JwtBearerInput() *IntegrationConnectorsConnectionAuthConfigOauth2JwtBearer {
	var returns *IntegrationConnectorsConnectionAuthConfigOauth2JwtBearer
	_jsii_.Get(
		j,
		"oauth2JwtBearerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) SshPublicKey() IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference {
	var returns IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference
	_jsii_.Get(
		j,
		"sshPublicKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) SshPublicKeyInput() *IntegrationConnectorsConnectionAuthConfigSshPublicKey {
	var returns *IntegrationConnectorsConnectionAuthConfigSshPublicKey
	_jsii_.Get(
		j,
		"sshPublicKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) UserPassword() IntegrationConnectorsConnectionAuthConfigUserPasswordOutputReference {
	var returns IntegrationConnectorsConnectionAuthConfigUserPasswordOutputReference
	_jsii_.Get(
		j,
		"userPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) UserPasswordInput() *IntegrationConnectorsConnectionAuthConfigUserPassword {
	var returns *IntegrationConnectorsConnectionAuthConfigUserPassword
	_jsii_.Get(
		j,
		"userPasswordInput",
		&returns,
	)
	return returns
}


func NewIntegrationConnectorsConnectionAuthConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IntegrationConnectorsConnectionAuthConfigOutputReference {
	_init_.Initialize()

	if err := validateNewIntegrationConnectorsConnectionAuthConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.integrationConnectorsConnection.IntegrationConnectorsConnectionAuthConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIntegrationConnectorsConnectionAuthConfigOutputReference_Override(i IntegrationConnectorsConnectionAuthConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.integrationConnectorsConnection.IntegrationConnectorsConnectionAuthConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference)SetAuthKey(val *string) {
	if err := j.validateSetAuthKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authKey",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference)SetAuthType(val *string) {
	if err := j.validateSetAuthTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authType",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference)SetInternalValue(val *IntegrationConnectorsConnectionAuthConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) PutAdditionalVariable(value interface{}) {
	if err := i.validatePutAdditionalVariableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAdditionalVariable",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) PutOauth2AuthCodeFlow(value *IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlow) {
	if err := i.validatePutOauth2AuthCodeFlowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putOauth2AuthCodeFlow",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) PutOauth2ClientCredentials(value *IntegrationConnectorsConnectionAuthConfigOauth2ClientCredentials) {
	if err := i.validatePutOauth2ClientCredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putOauth2ClientCredentials",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) PutOauth2JwtBearer(value *IntegrationConnectorsConnectionAuthConfigOauth2JwtBearer) {
	if err := i.validatePutOauth2JwtBearerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putOauth2JwtBearer",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) PutSshPublicKey(value *IntegrationConnectorsConnectionAuthConfigSshPublicKey) {
	if err := i.validatePutSshPublicKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSshPublicKey",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) PutUserPassword(value *IntegrationConnectorsConnectionAuthConfigUserPassword) {
	if err := i.validatePutUserPasswordParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putUserPassword",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) ResetAdditionalVariable() {
	_jsii_.InvokeVoid(
		i,
		"resetAdditionalVariable",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) ResetAuthKey() {
	_jsii_.InvokeVoid(
		i,
		"resetAuthKey",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) ResetOauth2AuthCodeFlow() {
	_jsii_.InvokeVoid(
		i,
		"resetOauth2AuthCodeFlow",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) ResetOauth2ClientCredentials() {
	_jsii_.InvokeVoid(
		i,
		"resetOauth2ClientCredentials",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) ResetOauth2JwtBearer() {
	_jsii_.InvokeVoid(
		i,
		"resetOauth2JwtBearer",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) ResetSshPublicKey() {
	_jsii_.InvokeVoid(
		i,
		"resetSshPublicKey",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) ResetUserPassword() {
	_jsii_.InvokeVoid(
		i,
		"resetUserPassword",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

