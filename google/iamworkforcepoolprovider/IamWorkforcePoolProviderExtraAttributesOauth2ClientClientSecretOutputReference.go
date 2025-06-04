// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iamworkforcepoolprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/iamworkforcepoolprovider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference interface {
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
	InternalValue() *IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecret
	SetInternalValue(val *IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecret)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Value() IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretValueOutputReference
	ValueInput() *IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretValue
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
	PutValue(value *IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretValue)
	ResetValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference
type jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference) InternalValue() *IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecret {
	var returns *IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecret
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference) Value() IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretValueOutputReference {
	var returns IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretValueOutputReference
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference) ValueInput() *IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretValue {
	var returns *IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretValue
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewIamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference {
	_init_.Initialize()

	if err := validateNewIamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.iamWorkforcePoolProvider.IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference_Override(i IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.iamWorkforcePoolProvider.IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference)SetInternalValue(val *IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecret) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference) PutValue(value *IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretValue) {
	if err := i.validatePutValueParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putValue",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference) ResetValue() {
	_jsii_.InvokeVoid(
		i,
		"resetValue",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (i *jsiiProxy_IamWorkforcePoolProviderExtraAttributesOauth2ClientClientSecretOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

