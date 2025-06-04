// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxpage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/dialogflowcxpage/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference interface {
	cdktf.ComplexObject
	Channel() *string
	SetChannel(val *string)
	ChannelInput() *string
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
	ConversationSuccess() DialogflowCxPageEventHandlersTriggerFulfillmentMessagesConversationSuccessOutputReference
	ConversationSuccessInput() *DialogflowCxPageEventHandlersTriggerFulfillmentMessagesConversationSuccess
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LiveAgentHandoff() DialogflowCxPageEventHandlersTriggerFulfillmentMessagesLiveAgentHandoffOutputReference
	LiveAgentHandoffInput() *DialogflowCxPageEventHandlersTriggerFulfillmentMessagesLiveAgentHandoff
	OutputAudioText() DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference
	OutputAudioTextInput() *DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioText
	Payload() *string
	SetPayload(val *string)
	PayloadInput() *string
	PlayAudio() DialogflowCxPageEventHandlersTriggerFulfillmentMessagesPlayAudioOutputReference
	PlayAudioInput() *DialogflowCxPageEventHandlersTriggerFulfillmentMessagesPlayAudio
	TelephonyTransferCall() DialogflowCxPageEventHandlersTriggerFulfillmentMessagesTelephonyTransferCallOutputReference
	TelephonyTransferCallInput() *DialogflowCxPageEventHandlersTriggerFulfillmentMessagesTelephonyTransferCall
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Text() DialogflowCxPageEventHandlersTriggerFulfillmentMessagesTextOutputReference
	TextInput() *DialogflowCxPageEventHandlersTriggerFulfillmentMessagesText
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
	PutConversationSuccess(value *DialogflowCxPageEventHandlersTriggerFulfillmentMessagesConversationSuccess)
	PutLiveAgentHandoff(value *DialogflowCxPageEventHandlersTriggerFulfillmentMessagesLiveAgentHandoff)
	PutOutputAudioText(value *DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioText)
	PutPlayAudio(value *DialogflowCxPageEventHandlersTriggerFulfillmentMessagesPlayAudio)
	PutTelephonyTransferCall(value *DialogflowCxPageEventHandlersTriggerFulfillmentMessagesTelephonyTransferCall)
	PutText(value *DialogflowCxPageEventHandlersTriggerFulfillmentMessagesText)
	ResetChannel()
	ResetConversationSuccess()
	ResetLiveAgentHandoff()
	ResetOutputAudioText()
	ResetPayload()
	ResetPlayAudio()
	ResetTelephonyTransferCall()
	ResetText()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference
type jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) Channel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) ChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) ConversationSuccess() DialogflowCxPageEventHandlersTriggerFulfillmentMessagesConversationSuccessOutputReference {
	var returns DialogflowCxPageEventHandlersTriggerFulfillmentMessagesConversationSuccessOutputReference
	_jsii_.Get(
		j,
		"conversationSuccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) ConversationSuccessInput() *DialogflowCxPageEventHandlersTriggerFulfillmentMessagesConversationSuccess {
	var returns *DialogflowCxPageEventHandlersTriggerFulfillmentMessagesConversationSuccess
	_jsii_.Get(
		j,
		"conversationSuccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) LiveAgentHandoff() DialogflowCxPageEventHandlersTriggerFulfillmentMessagesLiveAgentHandoffOutputReference {
	var returns DialogflowCxPageEventHandlersTriggerFulfillmentMessagesLiveAgentHandoffOutputReference
	_jsii_.Get(
		j,
		"liveAgentHandoff",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) LiveAgentHandoffInput() *DialogflowCxPageEventHandlersTriggerFulfillmentMessagesLiveAgentHandoff {
	var returns *DialogflowCxPageEventHandlersTriggerFulfillmentMessagesLiveAgentHandoff
	_jsii_.Get(
		j,
		"liveAgentHandoffInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) OutputAudioText() DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference {
	var returns DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference
	_jsii_.Get(
		j,
		"outputAudioText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) OutputAudioTextInput() *DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioText {
	var returns *DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioText
	_jsii_.Get(
		j,
		"outputAudioTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) Payload() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) PayloadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payloadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) PlayAudio() DialogflowCxPageEventHandlersTriggerFulfillmentMessagesPlayAudioOutputReference {
	var returns DialogflowCxPageEventHandlersTriggerFulfillmentMessagesPlayAudioOutputReference
	_jsii_.Get(
		j,
		"playAudio",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) PlayAudioInput() *DialogflowCxPageEventHandlersTriggerFulfillmentMessagesPlayAudio {
	var returns *DialogflowCxPageEventHandlersTriggerFulfillmentMessagesPlayAudio
	_jsii_.Get(
		j,
		"playAudioInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) TelephonyTransferCall() DialogflowCxPageEventHandlersTriggerFulfillmentMessagesTelephonyTransferCallOutputReference {
	var returns DialogflowCxPageEventHandlersTriggerFulfillmentMessagesTelephonyTransferCallOutputReference
	_jsii_.Get(
		j,
		"telephonyTransferCall",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) TelephonyTransferCallInput() *DialogflowCxPageEventHandlersTriggerFulfillmentMessagesTelephonyTransferCall {
	var returns *DialogflowCxPageEventHandlersTriggerFulfillmentMessagesTelephonyTransferCall
	_jsii_.Get(
		j,
		"telephonyTransferCallInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) Text() DialogflowCxPageEventHandlersTriggerFulfillmentMessagesTextOutputReference {
	var returns DialogflowCxPageEventHandlersTriggerFulfillmentMessagesTextOutputReference
	_jsii_.Get(
		j,
		"text",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) TextInput() *DialogflowCxPageEventHandlersTriggerFulfillmentMessagesText {
	var returns *DialogflowCxPageEventHandlersTriggerFulfillmentMessagesText
	_jsii_.Get(
		j,
		"textInput",
		&returns,
	)
	return returns
}


func NewDialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxPage.DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference_Override(d DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxPage.DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference)SetChannel(val *string) {
	if err := j.validateSetChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"channel",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference)SetPayload(val *string) {
	if err := j.validateSetPayloadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"payload",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) PutConversationSuccess(value *DialogflowCxPageEventHandlersTriggerFulfillmentMessagesConversationSuccess) {
	if err := d.validatePutConversationSuccessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putConversationSuccess",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) PutLiveAgentHandoff(value *DialogflowCxPageEventHandlersTriggerFulfillmentMessagesLiveAgentHandoff) {
	if err := d.validatePutLiveAgentHandoffParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLiveAgentHandoff",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) PutOutputAudioText(value *DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioText) {
	if err := d.validatePutOutputAudioTextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOutputAudioText",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) PutPlayAudio(value *DialogflowCxPageEventHandlersTriggerFulfillmentMessagesPlayAudio) {
	if err := d.validatePutPlayAudioParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPlayAudio",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) PutTelephonyTransferCall(value *DialogflowCxPageEventHandlersTriggerFulfillmentMessagesTelephonyTransferCall) {
	if err := d.validatePutTelephonyTransferCallParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTelephonyTransferCall",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) PutText(value *DialogflowCxPageEventHandlersTriggerFulfillmentMessagesText) {
	if err := d.validatePutTextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putText",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) ResetChannel() {
	_jsii_.InvokeVoid(
		d,
		"resetChannel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) ResetConversationSuccess() {
	_jsii_.InvokeVoid(
		d,
		"resetConversationSuccess",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) ResetLiveAgentHandoff() {
	_jsii_.InvokeVoid(
		d,
		"resetLiveAgentHandoff",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) ResetOutputAudioText() {
	_jsii_.InvokeVoid(
		d,
		"resetOutputAudioText",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) ResetPayload() {
	_jsii_.InvokeVoid(
		d,
		"resetPayload",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) ResetPlayAudio() {
	_jsii_.InvokeVoid(
		d,
		"resetPlayAudio",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) ResetTelephonyTransferCall() {
	_jsii_.InvokeVoid(
		d,
		"resetTelephonyTransferCall",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) ResetText() {
	_jsii_.InvokeVoid(
		d,
		"resetText",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

