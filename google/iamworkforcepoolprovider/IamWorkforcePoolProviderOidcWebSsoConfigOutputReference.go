// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iamworkforcepoolprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v9/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v9/iamworkforcepoolprovider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IamWorkforcePoolProviderOidcWebSsoConfigOutputReference interface {
	cdktf.ComplexObject
	AdditionalScopes() *[]*string
	SetAdditionalScopes(val *[]*string)
	AdditionalScopesInput() *[]*string
	AssertionClaimsBehavior() *string
	SetAssertionClaimsBehavior(val *string)
	AssertionClaimsBehaviorInput() *string
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
	InternalValue() *IamWorkforcePoolProviderOidcWebSsoConfig
	SetInternalValue(val *IamWorkforcePoolProviderOidcWebSsoConfig)
	ResponseType() *string
	SetResponseType(val *string)
	ResponseTypeInput() *string
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
	ResetAdditionalScopes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IamWorkforcePoolProviderOidcWebSsoConfigOutputReference
type jsiiProxy_IamWorkforcePoolProviderOidcWebSsoConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcWebSsoConfigOutputReference) AdditionalScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcWebSsoConfigOutputReference) AdditionalScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"additionalScopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcWebSsoConfigOutputReference) AssertionClaimsBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assertionClaimsBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcWebSsoConfigOutputReference) AssertionClaimsBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assertionClaimsBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcWebSsoConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcWebSsoConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcWebSsoConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcWebSsoConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcWebSsoConfigOutputReference) InternalValue() *IamWorkforcePoolProviderOidcWebSsoConfig {
	var returns *IamWorkforcePoolProviderOidcWebSsoConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcWebSsoConfigOutputReference) ResponseType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcWebSsoConfigOutputReference) ResponseTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcWebSsoConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcWebSsoConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIamWorkforcePoolProviderOidcWebSsoConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IamWorkforcePoolProviderOidcWebSsoConfigOutputReference {
	_init_.Initialize()

	if err := validateNewIamWorkforcePoolProviderOidcWebSsoConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IamWorkforcePoolProviderOidcWebSsoConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.iamWorkforcePoolProvider.IamWorkforcePoolProviderOidcWebSsoConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIamWorkforcePoolProviderOidcWebSsoConfigOutputReference_Override(i IamWorkforcePoolProviderOidcWebSsoConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.iamWorkforcePoolProvider.IamWorkforcePoolProviderOidcWebSsoConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcWebSsoConfigOutputReference)SetAdditionalScopes(val *[]*string) {
	if err := j.validateSetAdditionalScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalScopes",
		val,
	)
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcWebSsoConfigOutputReference)SetAssertionClaimsBehavior(val *string) {
	if err := j.validateSetAssertionClaimsBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assertionClaimsBehavior",
		val,
	)
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcWebSsoConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcWebSsoConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcWebSsoConfigOutputReference)SetInternalValue(val *IamWorkforcePoolProviderOidcWebSsoConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcWebSsoConfigOutputReference)SetResponseType(val *string) {
	if err := j.validateSetResponseTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseType",
		val,
	)
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcWebSsoConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IamWorkforcePoolProviderOidcWebSsoConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IamWorkforcePoolProviderOidcWebSsoConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkforcePoolProviderOidcWebSsoConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IamWorkforcePoolProviderOidcWebSsoConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IamWorkforcePoolProviderOidcWebSsoConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IamWorkforcePoolProviderOidcWebSsoConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IamWorkforcePoolProviderOidcWebSsoConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IamWorkforcePoolProviderOidcWebSsoConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IamWorkforcePoolProviderOidcWebSsoConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IamWorkforcePoolProviderOidcWebSsoConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IamWorkforcePoolProviderOidcWebSsoConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IamWorkforcePoolProviderOidcWebSsoConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkforcePoolProviderOidcWebSsoConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IamWorkforcePoolProviderOidcWebSsoConfigOutputReference) ResetAdditionalScopes() {
	_jsii_.InvokeVoid(
		i,
		"resetAdditionalScopes",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamWorkforcePoolProviderOidcWebSsoConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (i *jsiiProxy_IamWorkforcePoolProviderOidcWebSsoConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

