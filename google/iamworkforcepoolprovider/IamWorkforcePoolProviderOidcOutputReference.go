// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iamworkforcepoolprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/iamworkforcepoolprovider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IamWorkforcePoolProviderOidcOutputReference interface {
	cdktf.ComplexObject
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecret() IamWorkforcePoolProviderOidcClientSecretOutputReference
	ClientSecretInput() *IamWorkforcePoolProviderOidcClientSecret
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
	InternalValue() *IamWorkforcePoolProviderOidc
	SetInternalValue(val *IamWorkforcePoolProviderOidc)
	IssuerUri() *string
	SetIssuerUri(val *string)
	IssuerUriInput() *string
	JwksJson() *string
	SetJwksJson(val *string)
	JwksJsonInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WebSsoConfig() IamWorkforcePoolProviderOidcWebSsoConfigOutputReference
	WebSsoConfigInput() *IamWorkforcePoolProviderOidcWebSsoConfig
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
	PutClientSecret(value *IamWorkforcePoolProviderOidcClientSecret)
	PutWebSsoConfig(value *IamWorkforcePoolProviderOidcWebSsoConfig)
	ResetClientSecret()
	ResetJwksJson()
	ResetWebSsoConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IamWorkforcePoolProviderOidcOutputReference
type jsiiProxy_IamWorkforcePoolProviderOidcOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcOutputReference) ClientSecret() IamWorkforcePoolProviderOidcClientSecretOutputReference {
	var returns IamWorkforcePoolProviderOidcClientSecretOutputReference
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcOutputReference) ClientSecretInput() *IamWorkforcePoolProviderOidcClientSecret {
	var returns *IamWorkforcePoolProviderOidcClientSecret
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcOutputReference) InternalValue() *IamWorkforcePoolProviderOidc {
	var returns *IamWorkforcePoolProviderOidc
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcOutputReference) IssuerUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcOutputReference) IssuerUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcOutputReference) JwksJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwksJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcOutputReference) JwksJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwksJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcOutputReference) WebSsoConfig() IamWorkforcePoolProviderOidcWebSsoConfigOutputReference {
	var returns IamWorkforcePoolProviderOidcWebSsoConfigOutputReference
	_jsii_.Get(
		j,
		"webSsoConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcOutputReference) WebSsoConfigInput() *IamWorkforcePoolProviderOidcWebSsoConfig {
	var returns *IamWorkforcePoolProviderOidcWebSsoConfig
	_jsii_.Get(
		j,
		"webSsoConfigInput",
		&returns,
	)
	return returns
}


func NewIamWorkforcePoolProviderOidcOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IamWorkforcePoolProviderOidcOutputReference {
	_init_.Initialize()

	if err := validateNewIamWorkforcePoolProviderOidcOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IamWorkforcePoolProviderOidcOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.iamWorkforcePoolProvider.IamWorkforcePoolProviderOidcOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIamWorkforcePoolProviderOidcOutputReference_Override(i IamWorkforcePoolProviderOidcOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.iamWorkforcePoolProvider.IamWorkforcePoolProviderOidcOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcOutputReference)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcOutputReference)SetInternalValue(val *IamWorkforcePoolProviderOidc) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcOutputReference)SetIssuerUri(val *string) {
	if err := j.validateSetIssuerUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuerUri",
		val,
	)
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcOutputReference)SetJwksJson(val *string) {
	if err := j.validateSetJwksJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jwksJson",
		val,
	)
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IamWorkforcePoolProviderOidcOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkforcePoolProviderOidcOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IamWorkforcePoolProviderOidcOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IamWorkforcePoolProviderOidcOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IamWorkforcePoolProviderOidcOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IamWorkforcePoolProviderOidcOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IamWorkforcePoolProviderOidcOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IamWorkforcePoolProviderOidcOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IamWorkforcePoolProviderOidcOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IamWorkforcePoolProviderOidcOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IamWorkforcePoolProviderOidcOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkforcePoolProviderOidcOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IamWorkforcePoolProviderOidcOutputReference) PutClientSecret(value *IamWorkforcePoolProviderOidcClientSecret) {
	if err := i.validatePutClientSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putClientSecret",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IamWorkforcePoolProviderOidcOutputReference) PutWebSsoConfig(value *IamWorkforcePoolProviderOidcWebSsoConfig) {
	if err := i.validatePutWebSsoConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putWebSsoConfig",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IamWorkforcePoolProviderOidcOutputReference) ResetClientSecret() {
	_jsii_.InvokeVoid(
		i,
		"resetClientSecret",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamWorkforcePoolProviderOidcOutputReference) ResetJwksJson() {
	_jsii_.InvokeVoid(
		i,
		"resetJwksJson",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamWorkforcePoolProviderOidcOutputReference) ResetWebSsoConfig() {
	_jsii_.InvokeVoid(
		i,
		"resetWebSsoConfig",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamWorkforcePoolProviderOidcOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (i *jsiiProxy_IamWorkforcePoolProviderOidcOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

