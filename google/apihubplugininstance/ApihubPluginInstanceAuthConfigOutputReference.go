// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apihubplugininstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/apihubplugininstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApihubPluginInstanceAuthConfigOutputReference interface {
	cdktf.ComplexObject
	ApiKeyConfig() ApihubPluginInstanceAuthConfigApiKeyConfigOutputReference
	ApiKeyConfigInput() *ApihubPluginInstanceAuthConfigApiKeyConfig
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
	GoogleServiceAccountConfig() ApihubPluginInstanceAuthConfigGoogleServiceAccountConfigOutputReference
	GoogleServiceAccountConfigInput() *ApihubPluginInstanceAuthConfigGoogleServiceAccountConfig
	InternalValue() *ApihubPluginInstanceAuthConfig
	SetInternalValue(val *ApihubPluginInstanceAuthConfig)
	Oauth2ClientCredentialsConfig() ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigOutputReference
	Oauth2ClientCredentialsConfigInput() *ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfig
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserPasswordConfig() ApihubPluginInstanceAuthConfigUserPasswordConfigOutputReference
	UserPasswordConfigInput() *ApihubPluginInstanceAuthConfigUserPasswordConfig
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
	PutApiKeyConfig(value *ApihubPluginInstanceAuthConfigApiKeyConfig)
	PutGoogleServiceAccountConfig(value *ApihubPluginInstanceAuthConfigGoogleServiceAccountConfig)
	PutOauth2ClientCredentialsConfig(value *ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfig)
	PutUserPasswordConfig(value *ApihubPluginInstanceAuthConfigUserPasswordConfig)
	ResetApiKeyConfig()
	ResetGoogleServiceAccountConfig()
	ResetOauth2ClientCredentialsConfig()
	ResetUserPasswordConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApihubPluginInstanceAuthConfigOutputReference
type jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference) ApiKeyConfig() ApihubPluginInstanceAuthConfigApiKeyConfigOutputReference {
	var returns ApihubPluginInstanceAuthConfigApiKeyConfigOutputReference
	_jsii_.Get(
		j,
		"apiKeyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference) ApiKeyConfigInput() *ApihubPluginInstanceAuthConfigApiKeyConfig {
	var returns *ApihubPluginInstanceAuthConfigApiKeyConfig
	_jsii_.Get(
		j,
		"apiKeyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference) AuthType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference) AuthTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference) GoogleServiceAccountConfig() ApihubPluginInstanceAuthConfigGoogleServiceAccountConfigOutputReference {
	var returns ApihubPluginInstanceAuthConfigGoogleServiceAccountConfigOutputReference
	_jsii_.Get(
		j,
		"googleServiceAccountConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference) GoogleServiceAccountConfigInput() *ApihubPluginInstanceAuthConfigGoogleServiceAccountConfig {
	var returns *ApihubPluginInstanceAuthConfigGoogleServiceAccountConfig
	_jsii_.Get(
		j,
		"googleServiceAccountConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference) InternalValue() *ApihubPluginInstanceAuthConfig {
	var returns *ApihubPluginInstanceAuthConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference) Oauth2ClientCredentialsConfig() ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigOutputReference {
	var returns ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfigOutputReference
	_jsii_.Get(
		j,
		"oauth2ClientCredentialsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference) Oauth2ClientCredentialsConfigInput() *ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfig {
	var returns *ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfig
	_jsii_.Get(
		j,
		"oauth2ClientCredentialsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference) UserPasswordConfig() ApihubPluginInstanceAuthConfigUserPasswordConfigOutputReference {
	var returns ApihubPluginInstanceAuthConfigUserPasswordConfigOutputReference
	_jsii_.Get(
		j,
		"userPasswordConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference) UserPasswordConfigInput() *ApihubPluginInstanceAuthConfigUserPasswordConfig {
	var returns *ApihubPluginInstanceAuthConfigUserPasswordConfig
	_jsii_.Get(
		j,
		"userPasswordConfigInput",
		&returns,
	)
	return returns
}


func NewApihubPluginInstanceAuthConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApihubPluginInstanceAuthConfigOutputReference {
	_init_.Initialize()

	if err := validateNewApihubPluginInstanceAuthConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.apihubPluginInstance.ApihubPluginInstanceAuthConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApihubPluginInstanceAuthConfigOutputReference_Override(a ApihubPluginInstanceAuthConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.apihubPluginInstance.ApihubPluginInstanceAuthConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference)SetAuthType(val *string) {
	if err := j.validateSetAuthTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authType",
		val,
	)
}

func (j *jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference)SetInternalValue(val *ApihubPluginInstanceAuthConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference) PutApiKeyConfig(value *ApihubPluginInstanceAuthConfigApiKeyConfig) {
	if err := a.validatePutApiKeyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putApiKeyConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference) PutGoogleServiceAccountConfig(value *ApihubPluginInstanceAuthConfigGoogleServiceAccountConfig) {
	if err := a.validatePutGoogleServiceAccountConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGoogleServiceAccountConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference) PutOauth2ClientCredentialsConfig(value *ApihubPluginInstanceAuthConfigOauth2ClientCredentialsConfig) {
	if err := a.validatePutOauth2ClientCredentialsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOauth2ClientCredentialsConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference) PutUserPasswordConfig(value *ApihubPluginInstanceAuthConfigUserPasswordConfig) {
	if err := a.validatePutUserPasswordConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUserPasswordConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference) ResetApiKeyConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetApiKeyConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference) ResetGoogleServiceAccountConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetGoogleServiceAccountConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference) ResetOauth2ClientCredentialsConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetOauth2ClientCredentialsConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference) ResetUserPasswordConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetUserPasswordConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApihubPluginInstanceAuthConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

