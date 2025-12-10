// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowconversationprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/dialogflowconversationprofile/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DialogflowConversationProfileTtsConfigOutputReference interface {
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
	EffectsProfileId() *[]*string
	SetEffectsProfileId(val *[]*string)
	EffectsProfileIdInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *DialogflowConversationProfileTtsConfig
	SetInternalValue(val *DialogflowConversationProfileTtsConfig)
	Pitch() *float64
	SetPitch(val *float64)
	PitchInput() *float64
	SpeakingRate() *float64
	SetSpeakingRate(val *float64)
	SpeakingRateInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Voice() DialogflowConversationProfileTtsConfigVoiceOutputReference
	VoiceInput() *DialogflowConversationProfileTtsConfigVoice
	VolumeGainDb() *float64
	SetVolumeGainDb(val *float64)
	VolumeGainDbInput() *float64
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
	PutVoice(value *DialogflowConversationProfileTtsConfigVoice)
	ResetEffectsProfileId()
	ResetPitch()
	ResetSpeakingRate()
	ResetVoice()
	ResetVolumeGainDb()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DialogflowConversationProfileTtsConfigOutputReference
type jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference) EffectsProfileId() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"effectsProfileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference) EffectsProfileIdInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"effectsProfileIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference) InternalValue() *DialogflowConversationProfileTtsConfig {
	var returns *DialogflowConversationProfileTtsConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference) Pitch() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pitch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference) PitchInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pitchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference) SpeakingRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"speakingRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference) SpeakingRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"speakingRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference) Voice() DialogflowConversationProfileTtsConfigVoiceOutputReference {
	var returns DialogflowConversationProfileTtsConfigVoiceOutputReference
	_jsii_.Get(
		j,
		"voice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference) VoiceInput() *DialogflowConversationProfileTtsConfigVoice {
	var returns *DialogflowConversationProfileTtsConfigVoice
	_jsii_.Get(
		j,
		"voiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference) VolumeGainDb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeGainDb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference) VolumeGainDbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeGainDbInput",
		&returns,
	)
	return returns
}


func NewDialogflowConversationProfileTtsConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DialogflowConversationProfileTtsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowConversationProfileTtsConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileTtsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDialogflowConversationProfileTtsConfigOutputReference_Override(d DialogflowConversationProfileTtsConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileTtsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference)SetEffectsProfileId(val *[]*string) {
	if err := j.validateSetEffectsProfileIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"effectsProfileId",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference)SetInternalValue(val *DialogflowConversationProfileTtsConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference)SetPitch(val *float64) {
	if err := j.validateSetPitchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pitch",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference)SetSpeakingRate(val *float64) {
	if err := j.validateSetSpeakingRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"speakingRate",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference)SetVolumeGainDb(val *float64) {
	if err := j.validateSetVolumeGainDbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeGainDb",
		val,
	)
}

func (d *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference) PutVoice(value *DialogflowConversationProfileTtsConfigVoice) {
	if err := d.validatePutVoiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putVoice",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference) ResetEffectsProfileId() {
	_jsii_.InvokeVoid(
		d,
		"resetEffectsProfileId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference) ResetPitch() {
	_jsii_.InvokeVoid(
		d,
		"resetPitch",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference) ResetSpeakingRate() {
	_jsii_.InvokeVoid(
		d,
		"resetSpeakingRate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference) ResetVoice() {
	_jsii_.InvokeVoid(
		d,
		"resetVoice",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference) ResetVolumeGainDb() {
	_jsii_.InvokeVoid(
		d,
		"resetVolumeGainDb",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowConversationProfileTtsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

