// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iamworkforcepoolprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/iamworkforcepoolprovider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference interface {
	cdktf.ComplexObject
	AttributesType() *string
	SetAttributesType(val *string)
	AttributesTypeInput() *string
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecret() IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference
	ClientSecretInput() *IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecret
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
	InternalValue() *IamWorkforcePoolProviderExtraAttributesOauth2Client
	SetInternalValue(val *IamWorkforcePoolProviderExtraAttributesOauth2Client)
	IssuerUri() *string
	SetIssuerUri(val *string)
	IssuerUriInput() *string
	QueryParameters() IamWorkforcePoolProviderExtraAttributesOauth2ClientQueryParametersOutputReference
	QueryParametersInput() *IamWorkforcePoolProviderExtraAttributesOauth2ClientQueryParameters
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutClientSecret(value *IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecret)
	PutQueryParameters(value *IamWorkforcePoolProviderExtraAttributesOauth2ClientQueryParameters)
	ResetQueryParameters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference
type jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) AttributesType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributesType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) AttributesTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributesTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) ClientSecret() IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference {
	var returns IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) ClientSecretInput() *IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecret {
	var returns *IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecret
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) InternalValue() *IamWorkforcePoolProviderExtraAttributesOauth2Client {
	var returns *IamWorkforcePoolProviderExtraAttributesOauth2Client
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) IssuerUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) IssuerUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) QueryParameters() IamWorkforcePoolProviderExtraAttributesOauth2ClientQueryParametersOutputReference {
	var returns IamWorkforcePoolProviderExtraAttributesOauth2ClientQueryParametersOutputReference
	_jsii_.Get(
		j,
		"queryParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) QueryParametersInput() *IamWorkforcePoolProviderExtraAttributesOauth2ClientQueryParameters {
	var returns *IamWorkforcePoolProviderExtraAttributesOauth2ClientQueryParameters
	_jsii_.Get(
		j,
		"queryParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference {
	_init_.Initialize()

	if err := validateNewIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.iamWorkforcePoolProvider.IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference_Override(i IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.iamWorkforcePoolProvider.IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference)SetAttributesType(val *string) {
	if err := j.validateSetAttributesTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attributesType",
		val,
	)
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference)SetInternalValue(val *IamWorkforcePoolProviderExtraAttributesOauth2Client) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference)SetIssuerUri(val *string) {
	if err := j.validateSetIssuerUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuerUri",
		val,
	)
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) PutClientSecret(value *IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecret) {
	if err := i.validatePutClientSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putClientSecret",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) PutQueryParameters(value *IamWorkforcePoolProviderExtraAttributesOauth2ClientQueryParameters) {
	if err := i.validatePutQueryParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putQueryParameters",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) ResetQueryParameters() {
	_jsii_.InvokeVoid(
		i,
		"resetQueryParameters",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := i.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

