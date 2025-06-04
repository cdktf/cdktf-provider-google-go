// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transcoderjobtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/transcoderjobtemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TranscoderJobTemplateConfigMuxStreamsOutputReference interface {
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
	SegmentSettings() TranscoderJobTemplateConfigMuxStreamsSegmentSettingsOutputReference
	SegmentSettingsInput() *TranscoderJobTemplateConfigMuxStreamsSegmentSettings
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
	PutSegmentSettings(value *TranscoderJobTemplateConfigMuxStreamsSegmentSettings)
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

// The jsii proxy struct for TranscoderJobTemplateConfigMuxStreamsOutputReference
type jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference) Container() *string {
	var returns *string
	_jsii_.Get(
		j,
		"container",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference) ContainerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference) ElementaryStreams() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"elementaryStreams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference) ElementaryStreamsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"elementaryStreamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference) EncryptionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference) EncryptionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference) FileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference) FileNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference) SegmentSettings() TranscoderJobTemplateConfigMuxStreamsSegmentSettingsOutputReference {
	var returns TranscoderJobTemplateConfigMuxStreamsSegmentSettingsOutputReference
	_jsii_.Get(
		j,
		"segmentSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference) SegmentSettingsInput() *TranscoderJobTemplateConfigMuxStreamsSegmentSettings {
	var returns *TranscoderJobTemplateConfigMuxStreamsSegmentSettings
	_jsii_.Get(
		j,
		"segmentSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewTranscoderJobTemplateConfigMuxStreamsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TranscoderJobTemplateConfigMuxStreamsOutputReference {
	_init_.Initialize()

	if err := validateNewTranscoderJobTemplateConfigMuxStreamsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.transcoderJobTemplate.TranscoderJobTemplateConfigMuxStreamsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewTranscoderJobTemplateConfigMuxStreamsOutputReference_Override(t TranscoderJobTemplateConfigMuxStreamsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.transcoderJobTemplate.TranscoderJobTemplateConfigMuxStreamsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference)SetContainer(val *string) {
	if err := j.validateSetContainerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"container",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference)SetElementaryStreams(val *[]*string) {
	if err := j.validateSetElementaryStreamsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elementaryStreams",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference)SetEncryptionId(val *string) {
	if err := j.validateSetEncryptionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionId",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference)SetFileName(val *string) {
	if err := j.validateSetFileNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileName",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference)SetKey(val *string) {
	if err := j.validateSetKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (t *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (t *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference) PutSegmentSettings(value *TranscoderJobTemplateConfigMuxStreamsSegmentSettings) {
	if err := t.validatePutSegmentSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSegmentSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference) ResetContainer() {
	_jsii_.InvokeVoid(
		t,
		"resetContainer",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference) ResetElementaryStreams() {
	_jsii_.InvokeVoid(
		t,
		"resetElementaryStreams",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference) ResetEncryptionId() {
	_jsii_.InvokeVoid(
		t,
		"resetEncryptionId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference) ResetFileName() {
	_jsii_.InvokeVoid(
		t,
		"resetFileName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference) ResetKey() {
	_jsii_.InvokeVoid(
		t,
		"resetKey",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference) ResetSegmentSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetSegmentSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (t *jsiiProxy_TranscoderJobTemplateConfigMuxStreamsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

