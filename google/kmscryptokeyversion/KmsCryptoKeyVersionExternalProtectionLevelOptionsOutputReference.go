// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kmscryptokeyversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/kmscryptokeyversion/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KmsCryptoKeyVersionExternalProtectionLevelOptionsOutputReference interface {
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
	EkmConnectionKeyPath() *string
	SetEkmConnectionKeyPath(val *string)
	EkmConnectionKeyPathInput() *string
	ExternalKeyUri() *string
	SetExternalKeyUri(val *string)
	ExternalKeyUriInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *KmsCryptoKeyVersionExternalProtectionLevelOptions
	SetInternalValue(val *KmsCryptoKeyVersionExternalProtectionLevelOptions)
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
	ResetEkmConnectionKeyPath()
	ResetExternalKeyUri()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KmsCryptoKeyVersionExternalProtectionLevelOptionsOutputReference
type jsiiProxy_KmsCryptoKeyVersionExternalProtectionLevelOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KmsCryptoKeyVersionExternalProtectionLevelOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKeyVersionExternalProtectionLevelOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKeyVersionExternalProtectionLevelOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKeyVersionExternalProtectionLevelOptionsOutputReference) EkmConnectionKeyPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ekmConnectionKeyPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKeyVersionExternalProtectionLevelOptionsOutputReference) EkmConnectionKeyPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ekmConnectionKeyPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKeyVersionExternalProtectionLevelOptionsOutputReference) ExternalKeyUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalKeyUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKeyVersionExternalProtectionLevelOptionsOutputReference) ExternalKeyUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalKeyUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKeyVersionExternalProtectionLevelOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKeyVersionExternalProtectionLevelOptionsOutputReference) InternalValue() *KmsCryptoKeyVersionExternalProtectionLevelOptions {
	var returns *KmsCryptoKeyVersionExternalProtectionLevelOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKeyVersionExternalProtectionLevelOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmsCryptoKeyVersionExternalProtectionLevelOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKmsCryptoKeyVersionExternalProtectionLevelOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KmsCryptoKeyVersionExternalProtectionLevelOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewKmsCryptoKeyVersionExternalProtectionLevelOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KmsCryptoKeyVersionExternalProtectionLevelOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.kmsCryptoKeyVersion.KmsCryptoKeyVersionExternalProtectionLevelOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKmsCryptoKeyVersionExternalProtectionLevelOptionsOutputReference_Override(k KmsCryptoKeyVersionExternalProtectionLevelOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.kmsCryptoKeyVersion.KmsCryptoKeyVersionExternalProtectionLevelOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KmsCryptoKeyVersionExternalProtectionLevelOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KmsCryptoKeyVersionExternalProtectionLevelOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KmsCryptoKeyVersionExternalProtectionLevelOptionsOutputReference)SetEkmConnectionKeyPath(val *string) {
	if err := j.validateSetEkmConnectionKeyPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ekmConnectionKeyPath",
		val,
	)
}

func (j *jsiiProxy_KmsCryptoKeyVersionExternalProtectionLevelOptionsOutputReference)SetExternalKeyUri(val *string) {
	if err := j.validateSetExternalKeyUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalKeyUri",
		val,
	)
}

func (j *jsiiProxy_KmsCryptoKeyVersionExternalProtectionLevelOptionsOutputReference)SetInternalValue(val *KmsCryptoKeyVersionExternalProtectionLevelOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KmsCryptoKeyVersionExternalProtectionLevelOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KmsCryptoKeyVersionExternalProtectionLevelOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KmsCryptoKeyVersionExternalProtectionLevelOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsCryptoKeyVersionExternalProtectionLevelOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsCryptoKeyVersionExternalProtectionLevelOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsCryptoKeyVersionExternalProtectionLevelOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsCryptoKeyVersionExternalProtectionLevelOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsCryptoKeyVersionExternalProtectionLevelOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsCryptoKeyVersionExternalProtectionLevelOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsCryptoKeyVersionExternalProtectionLevelOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsCryptoKeyVersionExternalProtectionLevelOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsCryptoKeyVersionExternalProtectionLevelOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsCryptoKeyVersionExternalProtectionLevelOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsCryptoKeyVersionExternalProtectionLevelOptionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsCryptoKeyVersionExternalProtectionLevelOptionsOutputReference) ResetEkmConnectionKeyPath() {
	_jsii_.InvokeVoid(
		k,
		"resetEkmConnectionKeyPath",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmsCryptoKeyVersionExternalProtectionLevelOptionsOutputReference) ResetExternalKeyUri() {
	_jsii_.InvokeVoid(
		k,
		"resetExternalKeyUri",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmsCryptoKeyVersionExternalProtectionLevelOptionsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := k.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmsCryptoKeyVersionExternalProtectionLevelOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

