// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/dialogflowcxflow/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DialogflowCxFlowAdvancedSettingsOutputReference interface {
	cdktf.ComplexObject
	AudioExportGcsDestination() DialogflowCxFlowAdvancedSettingsAudioExportGcsDestinationOutputReference
	AudioExportGcsDestinationInput() *DialogflowCxFlowAdvancedSettingsAudioExportGcsDestination
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
	DtmfSettings() DialogflowCxFlowAdvancedSettingsDtmfSettingsOutputReference
	DtmfSettingsInput() *DialogflowCxFlowAdvancedSettingsDtmfSettings
	// Experimental.
	Fqn() *string
	InternalValue() *DialogflowCxFlowAdvancedSettings
	SetInternalValue(val *DialogflowCxFlowAdvancedSettings)
	LoggingSettings() DialogflowCxFlowAdvancedSettingsLoggingSettingsOutputReference
	LoggingSettingsInput() *DialogflowCxFlowAdvancedSettingsLoggingSettings
	SpeechSettings() DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference
	SpeechSettingsInput() *DialogflowCxFlowAdvancedSettingsSpeechSettings
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
	PutAudioExportGcsDestination(value *DialogflowCxFlowAdvancedSettingsAudioExportGcsDestination)
	PutDtmfSettings(value *DialogflowCxFlowAdvancedSettingsDtmfSettings)
	PutLoggingSettings(value *DialogflowCxFlowAdvancedSettingsLoggingSettings)
	PutSpeechSettings(value *DialogflowCxFlowAdvancedSettingsSpeechSettings)
	ResetAudioExportGcsDestination()
	ResetDtmfSettings()
	ResetLoggingSettings()
	ResetSpeechSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DialogflowCxFlowAdvancedSettingsOutputReference
type jsiiProxy_DialogflowCxFlowAdvancedSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DialogflowCxFlowAdvancedSettingsOutputReference) AudioExportGcsDestination() DialogflowCxFlowAdvancedSettingsAudioExportGcsDestinationOutputReference {
	var returns DialogflowCxFlowAdvancedSettingsAudioExportGcsDestinationOutputReference
	_jsii_.Get(
		j,
		"audioExportGcsDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowAdvancedSettingsOutputReference) AudioExportGcsDestinationInput() *DialogflowCxFlowAdvancedSettingsAudioExportGcsDestination {
	var returns *DialogflowCxFlowAdvancedSettingsAudioExportGcsDestination
	_jsii_.Get(
		j,
		"audioExportGcsDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowAdvancedSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowAdvancedSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowAdvancedSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowAdvancedSettingsOutputReference) DtmfSettings() DialogflowCxFlowAdvancedSettingsDtmfSettingsOutputReference {
	var returns DialogflowCxFlowAdvancedSettingsDtmfSettingsOutputReference
	_jsii_.Get(
		j,
		"dtmfSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowAdvancedSettingsOutputReference) DtmfSettingsInput() *DialogflowCxFlowAdvancedSettingsDtmfSettings {
	var returns *DialogflowCxFlowAdvancedSettingsDtmfSettings
	_jsii_.Get(
		j,
		"dtmfSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowAdvancedSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowAdvancedSettingsOutputReference) InternalValue() *DialogflowCxFlowAdvancedSettings {
	var returns *DialogflowCxFlowAdvancedSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowAdvancedSettingsOutputReference) LoggingSettings() DialogflowCxFlowAdvancedSettingsLoggingSettingsOutputReference {
	var returns DialogflowCxFlowAdvancedSettingsLoggingSettingsOutputReference
	_jsii_.Get(
		j,
		"loggingSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowAdvancedSettingsOutputReference) LoggingSettingsInput() *DialogflowCxFlowAdvancedSettingsLoggingSettings {
	var returns *DialogflowCxFlowAdvancedSettingsLoggingSettings
	_jsii_.Get(
		j,
		"loggingSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowAdvancedSettingsOutputReference) SpeechSettings() DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference {
	var returns DialogflowCxFlowAdvancedSettingsSpeechSettingsOutputReference
	_jsii_.Get(
		j,
		"speechSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowAdvancedSettingsOutputReference) SpeechSettingsInput() *DialogflowCxFlowAdvancedSettingsSpeechSettings {
	var returns *DialogflowCxFlowAdvancedSettingsSpeechSettings
	_jsii_.Get(
		j,
		"speechSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowAdvancedSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowAdvancedSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDialogflowCxFlowAdvancedSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DialogflowCxFlowAdvancedSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowCxFlowAdvancedSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowCxFlowAdvancedSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxFlow.DialogflowCxFlowAdvancedSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDialogflowCxFlowAdvancedSettingsOutputReference_Override(d DialogflowCxFlowAdvancedSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxFlow.DialogflowCxFlowAdvancedSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DialogflowCxFlowAdvancedSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxFlowAdvancedSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxFlowAdvancedSettingsOutputReference)SetInternalValue(val *DialogflowCxFlowAdvancedSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxFlowAdvancedSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxFlowAdvancedSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DialogflowCxFlowAdvancedSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxFlowAdvancedSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxFlowAdvancedSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxFlowAdvancedSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxFlowAdvancedSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxFlowAdvancedSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxFlowAdvancedSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxFlowAdvancedSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxFlowAdvancedSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxFlowAdvancedSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxFlowAdvancedSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxFlowAdvancedSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxFlowAdvancedSettingsOutputReference) PutAudioExportGcsDestination(value *DialogflowCxFlowAdvancedSettingsAudioExportGcsDestination) {
	if err := d.validatePutAudioExportGcsDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAudioExportGcsDestination",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxFlowAdvancedSettingsOutputReference) PutDtmfSettings(value *DialogflowCxFlowAdvancedSettingsDtmfSettings) {
	if err := d.validatePutDtmfSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDtmfSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxFlowAdvancedSettingsOutputReference) PutLoggingSettings(value *DialogflowCxFlowAdvancedSettingsLoggingSettings) {
	if err := d.validatePutLoggingSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLoggingSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxFlowAdvancedSettingsOutputReference) PutSpeechSettings(value *DialogflowCxFlowAdvancedSettingsSpeechSettings) {
	if err := d.validatePutSpeechSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSpeechSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxFlowAdvancedSettingsOutputReference) ResetAudioExportGcsDestination() {
	_jsii_.InvokeVoid(
		d,
		"resetAudioExportGcsDestination",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxFlowAdvancedSettingsOutputReference) ResetDtmfSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetDtmfSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxFlowAdvancedSettingsOutputReference) ResetLoggingSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetLoggingSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxFlowAdvancedSettingsOutputReference) ResetSpeechSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetSpeechSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxFlowAdvancedSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxFlowAdvancedSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

