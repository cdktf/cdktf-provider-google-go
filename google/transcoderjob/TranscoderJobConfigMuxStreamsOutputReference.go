// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transcoderjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/transcoderjob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TranscoderJobConfigMuxStreamsOutputReference interface {
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
	Container() *string
	SetContainer(val *string)
	ContainerInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ElementaryStreams() *[]*string
	SetElementaryStreams(val *[]*string)
	ElementaryStreamsInput() *[]*string
	EncryptionId() *string
	SetEncryptionId(val *string)
	EncryptionIdInput() *string
	FileName() *string
	SetFileName(val *string)
	FileNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	SegmentSettings() TranscoderJobConfigMuxStreamsSegmentSettingsOutputReference
	SegmentSettingsInput() *TranscoderJobConfigMuxStreamsSegmentSettings
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
	PutSegmentSettings(value *TranscoderJobConfigMuxStreamsSegmentSettings)
	ResetContainer()
	ResetElementaryStreams()
	ResetEncryptionId()
	ResetFileName()
	ResetKey()
	ResetSegmentSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TranscoderJobConfigMuxStreamsOutputReference
type jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference) Container() *string {
	var returns *string
	_jsii_.Get(
		j,
		"container",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference) ContainerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference) ElementaryStreams() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"elementaryStreams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference) ElementaryStreamsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"elementaryStreamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference) EncryptionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference) EncryptionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference) FileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference) FileNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference) SegmentSettings() TranscoderJobConfigMuxStreamsSegmentSettingsOutputReference {
	var returns TranscoderJobConfigMuxStreamsSegmentSettingsOutputReference
	_jsii_.Get(
		j,
		"segmentSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference) SegmentSettingsInput() *TranscoderJobConfigMuxStreamsSegmentSettings {
	var returns *TranscoderJobConfigMuxStreamsSegmentSettings
	_jsii_.Get(
		j,
		"segmentSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewTranscoderJobConfigMuxStreamsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TranscoderJobConfigMuxStreamsOutputReference {
	_init_.Initialize()

	if err := validateNewTranscoderJobConfigMuxStreamsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.transcoderJob.TranscoderJobConfigMuxStreamsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewTranscoderJobConfigMuxStreamsOutputReference_Override(t TranscoderJobConfigMuxStreamsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.transcoderJob.TranscoderJobConfigMuxStreamsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference)SetContainer(val *string) {
	if err := j.validateSetContainerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"container",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference)SetElementaryStreams(val *[]*string) {
	if err := j.validateSetElementaryStreamsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elementaryStreams",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference)SetEncryptionId(val *string) {
	if err := j.validateSetEncryptionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionId",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference)SetFileName(val *string) {
	if err := j.validateSetFileNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileName",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference)SetKey(val *string) {
	if err := j.validateSetKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (t *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (t *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference) PutSegmentSettings(value *TranscoderJobConfigMuxStreamsSegmentSettings) {
	if err := t.validatePutSegmentSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSegmentSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference) ResetContainer() {
	_jsii_.InvokeVoid(
		t,
		"resetContainer",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference) ResetElementaryStreams() {
	_jsii_.InvokeVoid(
		t,
		"resetElementaryStreams",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference) ResetEncryptionId() {
	_jsii_.InvokeVoid(
		t,
		"resetEncryptionId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference) ResetFileName() {
	_jsii_.InvokeVoid(
		t,
		"resetFileName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference) ResetKey() {
	_jsii_.InvokeVoid(
		t,
		"resetKey",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference) ResetSegmentSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetSegmentSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (t *jsiiProxy_TranscoderJobConfigMuxStreamsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

