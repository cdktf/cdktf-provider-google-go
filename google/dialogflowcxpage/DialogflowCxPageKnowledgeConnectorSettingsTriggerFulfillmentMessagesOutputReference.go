// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxpage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/dialogflowcxpage/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference interface {
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
	ConversationSuccess() DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesConversationSuccessOutputReference
	ConversationSuccessInput() *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesConversationSuccess
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EndInteraction() DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesEndInteractionList
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KnowledgeInfoCard() DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesKnowledgeInfoCardOutputReference
	KnowledgeInfoCardInput() *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesKnowledgeInfoCard
	LiveAgentHandoff() DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesLiveAgentHandoffOutputReference
	LiveAgentHandoffInput() *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesLiveAgentHandoff
	MixedAudio() DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesMixedAudioList
	OutputAudioText() DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputAudioTextOutputReference
	OutputAudioTextInput() *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputAudioText
	Payload() *string
	SetPayload(val *string)
	PayloadInput() *string
	PlayAudio() DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesPlayAudioOutputReference
	PlayAudioInput() *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesPlayAudio
	TelephonyTransferCall() DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesTelephonyTransferCallOutputReference
	TelephonyTransferCallInput() *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesTelephonyTransferCall
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Text() DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesTextOutputReference
	TextInput() *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesText
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
	PutConversationSuccess(value *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesConversationSuccess)
	PutKnowledgeInfoCard(value *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesKnowledgeInfoCard)
	PutLiveAgentHandoff(value *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesLiveAgentHandoff)
	PutOutputAudioText(value *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputAudioText)
	PutPlayAudio(value *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesPlayAudio)
	PutTelephonyTransferCall(value *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesTelephonyTransferCall)
	PutText(value *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesText)
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

// The jsii proxy struct for DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference
type jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) Channel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ConversationSuccess() DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesConversationSuccessOutputReference {
	var returns DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesConversationSuccessOutputReference
	_jsii_.Get(
		j,
		"conversationSuccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ConversationSuccessInput() *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesConversationSuccess {
	var returns *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesConversationSuccess
	_jsii_.Get(
		j,
		"conversationSuccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) EndInteraction() DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesEndInteractionList {
	var returns DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesEndInteractionList
	_jsii_.Get(
		j,
		"endInteraction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) KnowledgeInfoCard() DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesKnowledgeInfoCardOutputReference {
	var returns DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesKnowledgeInfoCardOutputReference
	_jsii_.Get(
		j,
		"knowledgeInfoCard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) KnowledgeInfoCardInput() *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesKnowledgeInfoCard {
	var returns *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesKnowledgeInfoCard
	_jsii_.Get(
		j,
		"knowledgeInfoCardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) LiveAgentHandoff() DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesLiveAgentHandoffOutputReference {
	var returns DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesLiveAgentHandoffOutputReference
	_jsii_.Get(
		j,
		"liveAgentHandoff",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) LiveAgentHandoffInput() *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesLiveAgentHandoff {
	var returns *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesLiveAgentHandoff
	_jsii_.Get(
		j,
		"liveAgentHandoffInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) MixedAudio() DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesMixedAudioList {
	var returns DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesMixedAudioList
	_jsii_.Get(
		j,
		"mixedAudio",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) OutputAudioText() DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputAudioTextOutputReference {
	var returns DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputAudioTextOutputReference
	_jsii_.Get(
		j,
		"outputAudioText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) OutputAudioTextInput() *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputAudioText {
	var returns *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputAudioText
	_jsii_.Get(
		j,
		"outputAudioTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) Payload() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) PayloadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payloadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) PlayAudio() DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesPlayAudioOutputReference {
	var returns DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesPlayAudioOutputReference
	_jsii_.Get(
		j,
		"playAudio",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) PlayAudioInput() *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesPlayAudio {
	var returns *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesPlayAudio
	_jsii_.Get(
		j,
		"playAudioInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) TelephonyTransferCall() DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesTelephonyTransferCallOutputReference {
	var returns DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesTelephonyTransferCallOutputReference
	_jsii_.Get(
		j,
		"telephonyTransferCall",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) TelephonyTransferCallInput() *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesTelephonyTransferCall {
	var returns *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesTelephonyTransferCall
	_jsii_.Get(
		j,
		"telephonyTransferCallInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) Text() DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesTextOutputReference {
	var returns DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesTextOutputReference
	_jsii_.Get(
		j,
		"text",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) TextInput() *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesText {
	var returns *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesText
	_jsii_.Get(
		j,
		"textInput",
		&returns,
	)
	return returns
}


func NewDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxPage.DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference_Override(d DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxPage.DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference)SetChannel(val *string) {
	if err := j.validateSetChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"channel",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference)SetPayload(val *string) {
	if err := j.validateSetPayloadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"payload",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) PutConversationSuccess(value *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesConversationSuccess) {
	if err := d.validatePutConversationSuccessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putConversationSuccess",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) PutKnowledgeInfoCard(value *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesKnowledgeInfoCard) {
	if err := d.validatePutKnowledgeInfoCardParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putKnowledgeInfoCard",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) PutLiveAgentHandoff(value *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesLiveAgentHandoff) {
	if err := d.validatePutLiveAgentHandoffParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLiveAgentHandoff",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) PutOutputAudioText(value *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputAudioText) {
	if err := d.validatePutOutputAudioTextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOutputAudioText",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) PutPlayAudio(value *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesPlayAudio) {
	if err := d.validatePutPlayAudioParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPlayAudio",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) PutTelephonyTransferCall(value *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesTelephonyTransferCall) {
	if err := d.validatePutTelephonyTransferCallParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTelephonyTransferCall",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) PutText(value *DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesText) {
	if err := d.validatePutTextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putText",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ResetChannel() {
	_jsii_.InvokeVoid(
		d,
		"resetChannel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ResetConversationSuccess() {
	_jsii_.InvokeVoid(
		d,
		"resetConversationSuccess",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ResetKnowledgeInfoCard() {
	_jsii_.InvokeVoid(
		d,
		"resetKnowledgeInfoCard",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ResetLiveAgentHandoff() {
	_jsii_.InvokeVoid(
		d,
		"resetLiveAgentHandoff",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ResetOutputAudioText() {
	_jsii_.InvokeVoid(
		d,
		"resetOutputAudioText",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ResetPayload() {
	_jsii_.InvokeVoid(
		d,
		"resetPayload",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ResetPlayAudio() {
	_jsii_.InvokeVoid(
		d,
		"resetPlayAudio",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ResetTelephonyTransferCall() {
	_jsii_.InvokeVoid(
		d,
		"resetTelephonyTransferCall",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ResetText() {
	_jsii_.InvokeVoid(
		d,
		"resetText",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DialogflowCxPageKnowledgeConnectorSettingsTriggerFulfillmentMessagesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

