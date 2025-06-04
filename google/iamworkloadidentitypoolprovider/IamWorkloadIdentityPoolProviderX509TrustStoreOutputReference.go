// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iamworkloadidentitypoolprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/iamworkloadidentitypoolprovider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IamWorkloadIdentityPoolProviderX509TrustStoreOutputReference interface {
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
	IntermediateCas() IamWorkloadIdentityPoolProviderX509TrustStoreIntermediateCasList
	IntermediateCasInput() interface{}
	InternalValue() *IamWorkloadIdentityPoolProviderX509TrustStore
	SetInternalValue(val *IamWorkloadIdentityPoolProviderX509TrustStore)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TrustAnchors() IamWorkloadIdentityPoolProviderX509TrustStoreTrustAnchorsList
	TrustAnchorsInput() interface{}
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
	PutIntermediateCas(value interface{})
	PutTrustAnchors(value interface{})
	ResetIntermediateCas()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IamWorkloadIdentityPoolProviderX509TrustStoreOutputReference
type jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) IntermediateCas() IamWorkloadIdentityPoolProviderX509TrustStoreIntermediateCasList {
	var returns IamWorkloadIdentityPoolProviderX509TrustStoreIntermediateCasList
	_jsii_.Get(
		j,
		"intermediateCas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) IntermediateCasInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"intermediateCasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) InternalValue() *IamWorkloadIdentityPoolProviderX509TrustStore {
	var returns *IamWorkloadIdentityPoolProviderX509TrustStore
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) TrustAnchors() IamWorkloadIdentityPoolProviderX509TrustStoreTrustAnchorsList {
	var returns IamWorkloadIdentityPoolProviderX509TrustStoreTrustAnchorsList
	_jsii_.Get(
		j,
		"trustAnchors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) TrustAnchorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trustAnchorsInput",
		&returns,
	)
	return returns
}


func NewIamWorkloadIdentityPoolProviderX509TrustStoreOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IamWorkloadIdentityPoolProviderX509TrustStoreOutputReference {
	_init_.Initialize()

	if err := validateNewIamWorkloadIdentityPoolProviderX509TrustStoreOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.iamWorkloadIdentityPoolProvider.IamWorkloadIdentityPoolProviderX509TrustStoreOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIamWorkloadIdentityPoolProviderX509TrustStoreOutputReference_Override(i IamWorkloadIdentityPoolProviderX509TrustStoreOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.iamWorkloadIdentityPoolProvider.IamWorkloadIdentityPoolProviderX509TrustStoreOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreOutputReference)SetInternalValue(val *IamWorkloadIdentityPoolProviderX509TrustStore) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) PutIntermediateCas(value interface{}) {
	if err := i.validatePutIntermediateCasParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putIntermediateCas",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) PutTrustAnchors(value interface{}) {
	if err := i.validatePutTrustAnchorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTrustAnchors",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) ResetIntermediateCas() {
	_jsii_.InvokeVoid(
		i,
		"resetIntermediateCas",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (i *jsiiProxy_IamWorkloadIdentityPoolProviderX509TrustStoreOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

