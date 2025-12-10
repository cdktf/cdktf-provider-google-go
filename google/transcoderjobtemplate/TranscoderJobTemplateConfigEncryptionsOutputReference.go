// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transcoderjobtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/transcoderjobtemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TranscoderJobTemplateConfigEncryptionsOutputReference interface {
	cdktf.ComplexObject
	Aes128() TranscoderJobTemplateConfigEncryptionsAes128OutputReference
	Aes128Input() *TranscoderJobTemplateConfigEncryptionsAes128
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
	DrmSystems() TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference
	DrmSystemsInput() *TranscoderJobTemplateConfigEncryptionsDrmSystems
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MpegCenc() TranscoderJobTemplateConfigEncryptionsMpegCencOutputReference
	MpegCencInput() *TranscoderJobTemplateConfigEncryptionsMpegCenc
	SampleAes() TranscoderJobTemplateConfigEncryptionsSampleAesOutputReference
	SampleAesInput() *TranscoderJobTemplateConfigEncryptionsSampleAes
	SecretManagerKeySource() TranscoderJobTemplateConfigEncryptionsSecretManagerKeySourceOutputReference
	SecretManagerKeySourceInput() *TranscoderJobTemplateConfigEncryptionsSecretManagerKeySource
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
	PutAes128(value *TranscoderJobTemplateConfigEncryptionsAes128)
	PutDrmSystems(value *TranscoderJobTemplateConfigEncryptionsDrmSystems)
	PutMpegCenc(value *TranscoderJobTemplateConfigEncryptionsMpegCenc)
	PutSampleAes(value *TranscoderJobTemplateConfigEncryptionsSampleAes)
	PutSecretManagerKeySource(value *TranscoderJobTemplateConfigEncryptionsSecretManagerKeySource)
	ResetAes128()
	ResetDrmSystems()
	ResetMpegCenc()
	ResetSampleAes()
	ResetSecretManagerKeySource()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TranscoderJobTemplateConfigEncryptionsOutputReference
type jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference) Aes128() TranscoderJobTemplateConfigEncryptionsAes128OutputReference {
	var returns TranscoderJobTemplateConfigEncryptionsAes128OutputReference
	_jsii_.Get(
		j,
		"aes128",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference) Aes128Input() *TranscoderJobTemplateConfigEncryptionsAes128 {
	var returns *TranscoderJobTemplateConfigEncryptionsAes128
	_jsii_.Get(
		j,
		"aes128Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference) DrmSystems() TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference {
	var returns TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference
	_jsii_.Get(
		j,
		"drmSystems",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference) DrmSystemsInput() *TranscoderJobTemplateConfigEncryptionsDrmSystems {
	var returns *TranscoderJobTemplateConfigEncryptionsDrmSystems
	_jsii_.Get(
		j,
		"drmSystemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference) MpegCenc() TranscoderJobTemplateConfigEncryptionsMpegCencOutputReference {
	var returns TranscoderJobTemplateConfigEncryptionsMpegCencOutputReference
	_jsii_.Get(
		j,
		"mpegCenc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference) MpegCencInput() *TranscoderJobTemplateConfigEncryptionsMpegCenc {
	var returns *TranscoderJobTemplateConfigEncryptionsMpegCenc
	_jsii_.Get(
		j,
		"mpegCencInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference) SampleAes() TranscoderJobTemplateConfigEncryptionsSampleAesOutputReference {
	var returns TranscoderJobTemplateConfigEncryptionsSampleAesOutputReference
	_jsii_.Get(
		j,
		"sampleAes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference) SampleAesInput() *TranscoderJobTemplateConfigEncryptionsSampleAes {
	var returns *TranscoderJobTemplateConfigEncryptionsSampleAes
	_jsii_.Get(
		j,
		"sampleAesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference) SecretManagerKeySource() TranscoderJobTemplateConfigEncryptionsSecretManagerKeySourceOutputReference {
	var returns TranscoderJobTemplateConfigEncryptionsSecretManagerKeySourceOutputReference
	_jsii_.Get(
		j,
		"secretManagerKeySource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference) SecretManagerKeySourceInput() *TranscoderJobTemplateConfigEncryptionsSecretManagerKeySource {
	var returns *TranscoderJobTemplateConfigEncryptionsSecretManagerKeySource
	_jsii_.Get(
		j,
		"secretManagerKeySourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewTranscoderJobTemplateConfigEncryptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TranscoderJobTemplateConfigEncryptionsOutputReference {
	_init_.Initialize()

	if err := validateNewTranscoderJobTemplateConfigEncryptionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.transcoderJobTemplate.TranscoderJobTemplateConfigEncryptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewTranscoderJobTemplateConfigEncryptionsOutputReference_Override(t TranscoderJobTemplateConfigEncryptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.transcoderJobTemplate.TranscoderJobTemplateConfigEncryptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference) PutAes128(value *TranscoderJobTemplateConfigEncryptionsAes128) {
	if err := t.validatePutAes128Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAes128",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference) PutDrmSystems(value *TranscoderJobTemplateConfigEncryptionsDrmSystems) {
	if err := t.validatePutDrmSystemsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDrmSystems",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference) PutMpegCenc(value *TranscoderJobTemplateConfigEncryptionsMpegCenc) {
	if err := t.validatePutMpegCencParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMpegCenc",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference) PutSampleAes(value *TranscoderJobTemplateConfigEncryptionsSampleAes) {
	if err := t.validatePutSampleAesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSampleAes",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference) PutSecretManagerKeySource(value *TranscoderJobTemplateConfigEncryptionsSecretManagerKeySource) {
	if err := t.validatePutSecretManagerKeySourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSecretManagerKeySource",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference) ResetAes128() {
	_jsii_.InvokeVoid(
		t,
		"resetAes128",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference) ResetDrmSystems() {
	_jsii_.InvokeVoid(
		t,
		"resetDrmSystems",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference) ResetMpegCenc() {
	_jsii_.InvokeVoid(
		t,
		"resetMpegCenc",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference) ResetSampleAes() {
	_jsii_.InvokeVoid(
		t,
		"resetSampleAes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference) ResetSecretManagerKeySource() {
	_jsii_.InvokeVoid(
		t,
		"resetSecretManagerKeySource",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := t.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

