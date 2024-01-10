// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identityplatformconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v13/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v13/identityplatformconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference interface {
	cdktf.ComplexObject
	AccessToken() interface{}
	SetAccessToken(val interface{})
	AccessTokenInput() interface{}
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
	IdToken() interface{}
	SetIdToken(val interface{})
	IdTokenInput() interface{}
	InternalValue() *IdentityPlatformConfigBlockingFunctionsForwardInboundCredentials
	SetInternalValue(val *IdentityPlatformConfigBlockingFunctionsForwardInboundCredentials)
	RefreshToken() interface{}
	SetRefreshToken(val interface{})
	RefreshTokenInput() interface{}
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
	ResetAccessToken()
	ResetIdToken()
	ResetRefreshToken()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference
type jsiiProxy_IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) AccessToken() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) AccessTokenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) IdToken() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"idToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) IdTokenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"idTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) InternalValue() *IdentityPlatformConfigBlockingFunctionsForwardInboundCredentials {
	var returns *IdentityPlatformConfigBlockingFunctionsForwardInboundCredentials
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) RefreshToken() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"refreshToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) RefreshTokenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"refreshTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference {
	_init_.Initialize()

	if err := validateNewIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.identityPlatformConfig.IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference_Override(i IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.identityPlatformConfig.IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference)SetAccessToken(val interface{}) {
	if err := j.validateSetAccessTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessToken",
		val,
	)
}

func (j *jsiiProxy_IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference)SetIdToken(val interface{}) {
	if err := j.validateSetIdTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idToken",
		val,
	)
}

func (j *jsiiProxy_IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference)SetInternalValue(val *IdentityPlatformConfigBlockingFunctionsForwardInboundCredentials) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference)SetRefreshToken(val interface{}) {
	if err := j.validateSetRefreshTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"refreshToken",
		val,
	)
}

func (j *jsiiProxy_IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) ResetAccessToken() {
	_jsii_.InvokeVoid(
		i,
		"resetAccessToken",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) ResetIdToken() {
	_jsii_.InvokeVoid(
		i,
		"resetIdToken",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) ResetRefreshToken() {
	_jsii_.InvokeVoid(
		i,
		"resetRefreshToken",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (i *jsiiProxy_IdentityPlatformConfigBlockingFunctionsForwardInboundCredentialsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

