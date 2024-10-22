// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transcoderjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/transcoderjob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TranscoderJobConfigEncryptionsOutputReference interface {
	cdktf.ComplexObject
	Aes128() TranscoderJobConfigEncryptionsAes128OutputReference
	Aes128Input() *TranscoderJobConfigEncryptionsAes128
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
	DrmSystems() TranscoderJobConfigEncryptionsDrmSystemsOutputReference
	DrmSystemsInput() *TranscoderJobConfigEncryptionsDrmSystems
	// Experimental.
	Fqn() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MpegCenc() TranscoderJobConfigEncryptionsMpegCencOutputReference
	MpegCencInput() *TranscoderJobConfigEncryptionsMpegCenc
	SampleAes() TranscoderJobConfigEncryptionsSampleAesOutputReference
	SampleAesInput() *TranscoderJobConfigEncryptionsSampleAes
	SecretManagerKeySource() TranscoderJobConfigEncryptionsSecretManagerKeySourceOutputReference
	SecretManagerKeySourceInput() *TranscoderJobConfigEncryptionsSecretManagerKeySource
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
	PutAes128(value *TranscoderJobConfigEncryptionsAes128)
	PutDrmSystems(value *TranscoderJobConfigEncryptionsDrmSystems)
	PutMpegCenc(value *TranscoderJobConfigEncryptionsMpegCenc)
	PutSampleAes(value *TranscoderJobConfigEncryptionsSampleAes)
	PutSecretManagerKeySource(value *TranscoderJobConfigEncryptionsSecretManagerKeySource)
	ResetAes128()
	ResetDrmSystems()
	ResetMpegCenc()
	ResetSampleAes()
	ResetSecretManagerKeySource()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TranscoderJobConfigEncryptionsOutputReference
type jsiiProxy_TranscoderJobConfigEncryptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference) Aes128() TranscoderJobConfigEncryptionsAes128OutputReference {
	var returns TranscoderJobConfigEncryptionsAes128OutputReference
	_jsii_.Get(
		j,
		"aes128",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference) Aes128Input() *TranscoderJobConfigEncryptionsAes128 {
	var returns *TranscoderJobConfigEncryptionsAes128
	_jsii_.Get(
		j,
		"aes128Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference) DrmSystems() TranscoderJobConfigEncryptionsDrmSystemsOutputReference {
	var returns TranscoderJobConfigEncryptionsDrmSystemsOutputReference
	_jsii_.Get(
		j,
		"drmSystems",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference) DrmSystemsInput() *TranscoderJobConfigEncryptionsDrmSystems {
	var returns *TranscoderJobConfigEncryptionsDrmSystems
	_jsii_.Get(
		j,
		"drmSystemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference) MpegCenc() TranscoderJobConfigEncryptionsMpegCencOutputReference {
	var returns TranscoderJobConfigEncryptionsMpegCencOutputReference
	_jsii_.Get(
		j,
		"mpegCenc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference) MpegCencInput() *TranscoderJobConfigEncryptionsMpegCenc {
	var returns *TranscoderJobConfigEncryptionsMpegCenc
	_jsii_.Get(
		j,
		"mpegCencInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference) SampleAes() TranscoderJobConfigEncryptionsSampleAesOutputReference {
	var returns TranscoderJobConfigEncryptionsSampleAesOutputReference
	_jsii_.Get(
		j,
		"sampleAes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference) SampleAesInput() *TranscoderJobConfigEncryptionsSampleAes {
	var returns *TranscoderJobConfigEncryptionsSampleAes
	_jsii_.Get(
		j,
		"sampleAesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference) SecretManagerKeySource() TranscoderJobConfigEncryptionsSecretManagerKeySourceOutputReference {
	var returns TranscoderJobConfigEncryptionsSecretManagerKeySourceOutputReference
	_jsii_.Get(
		j,
		"secretManagerKeySource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference) SecretManagerKeySourceInput() *TranscoderJobConfigEncryptionsSecretManagerKeySource {
	var returns *TranscoderJobConfigEncryptionsSecretManagerKeySource
	_jsii_.Get(
		j,
		"secretManagerKeySourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewTranscoderJobConfigEncryptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TranscoderJobConfigEncryptionsOutputReference {
	_init_.Initialize()

	if err := validateNewTranscoderJobConfigEncryptionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TranscoderJobConfigEncryptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.transcoderJob.TranscoderJobConfigEncryptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewTranscoderJobConfigEncryptionsOutputReference_Override(t TranscoderJobConfigEncryptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.transcoderJob.TranscoderJobConfigEncryptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (t *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference) PutAes128(value *TranscoderJobConfigEncryptionsAes128) {
	if err := t.validatePutAes128Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAes128",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference) PutDrmSystems(value *TranscoderJobConfigEncryptionsDrmSystems) {
	if err := t.validatePutDrmSystemsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDrmSystems",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference) PutMpegCenc(value *TranscoderJobConfigEncryptionsMpegCenc) {
	if err := t.validatePutMpegCencParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMpegCenc",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference) PutSampleAes(value *TranscoderJobConfigEncryptionsSampleAes) {
	if err := t.validatePutSampleAesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSampleAes",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference) PutSecretManagerKeySource(value *TranscoderJobConfigEncryptionsSecretManagerKeySource) {
	if err := t.validatePutSecretManagerKeySourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSecretManagerKeySource",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference) ResetAes128() {
	_jsii_.InvokeVoid(
		t,
		"resetAes128",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference) ResetDrmSystems() {
	_jsii_.InvokeVoid(
		t,
		"resetDrmSystems",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference) ResetMpegCenc() {
	_jsii_.InvokeVoid(
		t,
		"resetMpegCenc",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference) ResetSampleAes() {
	_jsii_.InvokeVoid(
		t,
		"resetSampleAes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference) ResetSecretManagerKeySource() {
	_jsii_.InvokeVoid(
		t,
		"resetSecretManagerKeySource",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := t.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobConfigEncryptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

