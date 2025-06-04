// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/dialogflowcxflow/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference interface {
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
	ConversationSuccess() DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesConversationSuccessOutputReference
	ConversationSuccessInput() *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesConversationSuccess
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EndInteraction() DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesEndInteractionList
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KnowledgeInfoCard() DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesKnowledgeInfoCardOutputReference
	KnowledgeInfoCardInput() *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesKnowledgeInfoCard
	LiveAgentHandoff() DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesLiveAgentHandoffOutputReference
	LiveAgentHandoffInput() *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesLiveAgentHandoff
	MixedAudio() DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesMixedAudioList
	OutputAudioText() DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputAudioTextOutputReference
	OutputAudioTextInput() *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputAudioText
	Payload() *string
	SetPayload(val *string)
	PayloadInput() *string
	PlayAudio() DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesPlayAudioOutputReference
	PlayAudioInput() *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesPlayAudio
	TelephonyTransferCall() DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesTelephonyTransferCallOutputReference
	TelephonyTransferCallInput() *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesTelephonyTransferCall
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Text() DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesTextOutputReference
	TextInput() *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesText
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
	PutConversationSuccess(value *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesConversationSuccess)
	PutKnowledgeInfoCard(value *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesKnowledgeInfoCard)
	PutLiveAgentHandoff(value *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesLiveAgentHandoff)
	PutOutputAudioText(value *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputAudioText)
	PutPlayAudio(value *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesPlayAudio)
	PutTelephonyTransferCall(value *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesTelephonyTransferCall)
	PutText(value *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesText)
	ResetChannel()
	ResetConversationSuccess()
	ResetKnowledgeInfoCard()
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

// The jsii proxy struct for DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference
type jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) Channel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ConversationSuccess() DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesConversationSuccessOutputReference {
	var returns DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesConversationSuccessOutputReference
	_jsii_.Get(
		j,
		"conversationSuccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ConversationSuccessInput() *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesConversationSuccess {
	var returns *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesConversationSuccess
	_jsii_.Get(
		j,
		"conversationSuccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) EndInteraction() DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesEndInteractionList {
	var returns DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesEndInteractionList
	_jsii_.Get(
		j,
		"endInteraction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) KnowledgeInfoCard() DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesKnowledgeInfoCardOutputReference {
	var returns DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesKnowledgeInfoCardOutputReference
	_jsii_.Get(
		j,
		"knowledgeInfoCard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) KnowledgeInfoCardInput() *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesKnowledgeInfoCard {
	var returns *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesKnowledgeInfoCard
	_jsii_.Get(
		j,
		"knowledgeInfoCardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) LiveAgentHandoff() DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesLiveAgentHandoffOutputReference {
	var returns DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesLiveAgentHandoffOutputReference
	_jsii_.Get(
		j,
		"liveAgentHandoff",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) LiveAgentHandoffInput() *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesLiveAgentHandoff {
	var returns *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesLiveAgentHandoff
	_jsii_.Get(
		j,
		"liveAgentHandoffInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) MixedAudio() DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesMixedAudioList {
	var returns DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesMixedAudioList
	_jsii_.Get(
		j,
		"mixedAudio",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) OutputAudioText() DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputAudioTextOutputReference {
	var returns DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputAudioTextOutputReference
	_jsii_.Get(
		j,
		"outputAudioText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) OutputAudioTextInput() *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputAudioText {
	var returns *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputAudioText
	_jsii_.Get(
		j,
		"outputAudioTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) Payload() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) PayloadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payloadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) PlayAudio() DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesPlayAudioOutputReference {
	var returns DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesPlayAudioOutputReference
	_jsii_.Get(
		j,
		"playAudio",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) PlayAudioInput() *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesPlayAudio {
	var returns *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesPlayAudio
	_jsii_.Get(
		j,
		"playAudioInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) TelephonyTransferCall() DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesTelephonyTransferCallOutputReference {
	var returns DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesTelephonyTransferCallOutputReference
	_jsii_.Get(
		j,
		"telephonyTransferCall",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) TelephonyTransferCallInput() *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesTelephonyTransferCall {
	var returns *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesTelephonyTransferCall
	_jsii_.Get(
		j,
		"telephonyTransferCallInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) Text() DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesTextOutputReference {
	var returns DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesTextOutputReference
	_jsii_.Get(
		j,
		"text",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) TextInput() *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesText {
	var returns *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesText
	_jsii_.Get(
		j,
		"textInput",
		&returns,
	)
	return returns
}


func NewDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxFlow.DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference_Override(d DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxFlow.DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference)SetChannel(val *string) {
	if err := j.validateSetChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"channel",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference)SetPayload(val *string) {
	if err := j.validateSetPayloadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"payload",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) PutConversationSuccess(value *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesConversationSuccess) {
	if err := d.validatePutConversationSuccessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putConversationSuccess",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) PutKnowledgeInfoCard(value *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesKnowledgeInfoCard) {
	if err := d.validatePutKnowledgeInfoCardParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putKnowledgeInfoCard",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) PutLiveAgentHandoff(value *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesLiveAgentHandoff) {
	if err := d.validatePutLiveAgentHandoffParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLiveAgentHandoff",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) PutOutputAudioText(value *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputAudioText) {
	if err := d.validatePutOutputAudioTextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOutputAudioText",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) PutPlayAudio(value *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesPlayAudio) {
	if err := d.validatePutPlayAudioParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPlayAudio",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) PutTelephonyTransferCall(value *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesTelephonyTransferCall) {
	if err := d.validatePutTelephonyTransferCallParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTelephonyTransferCall",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) PutText(value *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesText) {
	if err := d.validatePutTextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putText",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ResetChannel() {
	_jsii_.InvokeVoid(
		d,
		"resetChannel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ResetConversationSuccess() {
	_jsii_.InvokeVoid(
		d,
		"resetConversationSuccess",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ResetKnowledgeInfoCard() {
	_jsii_.InvokeVoid(
		d,
		"resetKnowledgeInfoCard",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ResetLiveAgentHandoff() {
	_jsii_.InvokeVoid(
		d,
		"resetLiveAgentHandoff",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ResetOutputAudioText() {
	_jsii_.InvokeVoid(
		d,
		"resetOutputAudioText",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ResetPayload() {
	_jsii_.InvokeVoid(
		d,
		"resetPayload",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ResetPlayAudio() {
	_jsii_.InvokeVoid(
		d,
		"resetPlayAudio",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ResetTelephonyTransferCall() {
	_jsii_.InvokeVoid(
		d,
		"resetTelephonyTransferCall",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ResetText() {
	_jsii_.InvokeVoid(
		d,
		"resetText",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

