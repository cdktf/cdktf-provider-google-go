// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventarcpipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/eventarcpipeline/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EventarcPipelineDestinationsOutputReference interface {
	cdktf.ComplexObject
	AuthenticationConfig() EventarcPipelineDestinationsAuthenticationConfigOutputReference
	AuthenticationConfigInput() *EventarcPipelineDestinationsAuthenticationConfig
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
	// Experimental.
	Fqn() *string
	HttpEndpoint() EventarcPipelineDestinationsHttpEndpointOutputReference
	HttpEndpointInput() *EventarcPipelineDestinationsHttpEndpoint
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MessageBus() *string
	SetMessageBus(val *string)
	MessageBusInput() *string
	NetworkConfig() EventarcPipelineDestinationsNetworkConfigOutputReference
	NetworkConfigInput() *EventarcPipelineDestinationsNetworkConfig
	OutputPayloadFormat() EventarcPipelineDestinationsOutputPayloadFormatOutputReference
	OutputPayloadFormatInput() *EventarcPipelineDestinationsOutputPayloadFormat
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Topic() *string
	SetTopic(val *string)
	TopicInput() *string
	Workflow() *string
	SetWorkflow(val *string)
	WorkflowInput() *string
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
	PutAuthenticationConfig(value *EventarcPipelineDestinationsAuthenticationConfig)
	PutHttpEndpoint(value *EventarcPipelineDestinationsHttpEndpoint)
	PutNetworkConfig(value *EventarcPipelineDestinationsNetworkConfig)
	PutOutputPayloadFormat(value *EventarcPipelineDestinationsOutputPayloadFormat)
	ResetAuthenticationConfig()
	ResetHttpEndpoint()
	ResetMessageBus()
	ResetNetworkConfig()
	ResetOutputPayloadFormat()
	ResetTopic()
	ResetWorkflow()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventarcPipelineDestinationsOutputReference
type jsiiProxy_EventarcPipelineDestinationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventarcPipelineDestinationsOutputReference) AuthenticationConfig() EventarcPipelineDestinationsAuthenticationConfigOutputReference {
	var returns EventarcPipelineDestinationsAuthenticationConfigOutputReference
	_jsii_.Get(
		j,
		"authenticationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcPipelineDestinationsOutputReference) AuthenticationConfigInput() *EventarcPipelineDestinationsAuthenticationConfig {
	var returns *EventarcPipelineDestinationsAuthenticationConfig
	_jsii_.Get(
		j,
		"authenticationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcPipelineDestinationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcPipelineDestinationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcPipelineDestinationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcPipelineDestinationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcPipelineDestinationsOutputReference) HttpEndpoint() EventarcPipelineDestinationsHttpEndpointOutputReference {
	var returns EventarcPipelineDestinationsHttpEndpointOutputReference
	_jsii_.Get(
		j,
		"httpEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcPipelineDestinationsOutputReference) HttpEndpointInput() *EventarcPipelineDestinationsHttpEndpoint {
	var returns *EventarcPipelineDestinationsHttpEndpoint
	_jsii_.Get(
		j,
		"httpEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcPipelineDestinationsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcPipelineDestinationsOutputReference) MessageBus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageBus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcPipelineDestinationsOutputReference) MessageBusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageBusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcPipelineDestinationsOutputReference) NetworkConfig() EventarcPipelineDestinationsNetworkConfigOutputReference {
	var returns EventarcPipelineDestinationsNetworkConfigOutputReference
	_jsii_.Get(
		j,
		"networkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcPipelineDestinationsOutputReference) NetworkConfigInput() *EventarcPipelineDestinationsNetworkConfig {
	var returns *EventarcPipelineDestinationsNetworkConfig
	_jsii_.Get(
		j,
		"networkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcPipelineDestinationsOutputReference) OutputPayloadFormat() EventarcPipelineDestinationsOutputPayloadFormatOutputReference {
	var returns EventarcPipelineDestinationsOutputPayloadFormatOutputReference
	_jsii_.Get(
		j,
		"outputPayloadFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcPipelineDestinationsOutputReference) OutputPayloadFormatInput() *EventarcPipelineDestinationsOutputPayloadFormat {
	var returns *EventarcPipelineDestinationsOutputPayloadFormat
	_jsii_.Get(
		j,
		"outputPayloadFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcPipelineDestinationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcPipelineDestinationsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcPipelineDestinationsOutputReference) Topic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcPipelineDestinationsOutputReference) TopicInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcPipelineDestinationsOutputReference) Workflow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workflow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcPipelineDestinationsOutputReference) WorkflowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workflowInput",
		&returns,
	)
	return returns
}


func NewEventarcPipelineDestinationsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EventarcPipelineDestinationsOutputReference {
	_init_.Initialize()

	if err := validateNewEventarcPipelineDestinationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_EventarcPipelineDestinationsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.eventarcPipeline.EventarcPipelineDestinationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEventarcPipelineDestinationsOutputReference_Override(e EventarcPipelineDestinationsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.eventarcPipeline.EventarcPipelineDestinationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EventarcPipelineDestinationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventarcPipelineDestinationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventarcPipelineDestinationsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventarcPipelineDestinationsOutputReference)SetMessageBus(val *string) {
	if err := j.validateSetMessageBusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messageBus",
		val,
	)
}

func (j *jsiiProxy_EventarcPipelineDestinationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventarcPipelineDestinationsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventarcPipelineDestinationsOutputReference)SetTopic(val *string) {
	if err := j.validateSetTopicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topic",
		val,
	)
}

func (j *jsiiProxy_EventarcPipelineDestinationsOutputReference)SetWorkflow(val *string) {
	if err := j.validateSetWorkflowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workflow",
		val,
	)
}

func (e *jsiiProxy_EventarcPipelineDestinationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventarcPipelineDestinationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventarcPipelineDestinationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventarcPipelineDestinationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventarcPipelineDestinationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventarcPipelineDestinationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventarcPipelineDestinationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventarcPipelineDestinationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventarcPipelineDestinationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventarcPipelineDestinationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventarcPipelineDestinationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventarcPipelineDestinationsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventarcPipelineDestinationsOutputReference) PutAuthenticationConfig(value *EventarcPipelineDestinationsAuthenticationConfig) {
	if err := e.validatePutAuthenticationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAuthenticationConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventarcPipelineDestinationsOutputReference) PutHttpEndpoint(value *EventarcPipelineDestinationsHttpEndpoint) {
	if err := e.validatePutHttpEndpointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putHttpEndpoint",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventarcPipelineDestinationsOutputReference) PutNetworkConfig(value *EventarcPipelineDestinationsNetworkConfig) {
	if err := e.validatePutNetworkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNetworkConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventarcPipelineDestinationsOutputReference) PutOutputPayloadFormat(value *EventarcPipelineDestinationsOutputPayloadFormat) {
	if err := e.validatePutOutputPayloadFormatParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putOutputPayloadFormat",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventarcPipelineDestinationsOutputReference) ResetAuthenticationConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetAuthenticationConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventarcPipelineDestinationsOutputReference) ResetHttpEndpoint() {
	_jsii_.InvokeVoid(
		e,
		"resetHttpEndpoint",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventarcPipelineDestinationsOutputReference) ResetMessageBus() {
	_jsii_.InvokeVoid(
		e,
		"resetMessageBus",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventarcPipelineDestinationsOutputReference) ResetNetworkConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetNetworkConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventarcPipelineDestinationsOutputReference) ResetOutputPayloadFormat() {
	_jsii_.InvokeVoid(
		e,
		"resetOutputPayloadFormat",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventarcPipelineDestinationsOutputReference) ResetTopic() {
	_jsii_.InvokeVoid(
		e,
		"resetTopic",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventarcPipelineDestinationsOutputReference) ResetWorkflow() {
	_jsii_.InvokeVoid(
		e,
		"resetWorkflow",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventarcPipelineDestinationsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventarcPipelineDestinationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

