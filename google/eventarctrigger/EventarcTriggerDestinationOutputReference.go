// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventarctrigger

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v13/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v13/eventarctrigger/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EventarcTriggerDestinationOutputReference interface {
	cdktf.ComplexObject
	CloudFunction() *string
	CloudRunService() EventarcTriggerDestinationCloudRunServiceOutputReference
	CloudRunServiceInput() *EventarcTriggerDestinationCloudRunService
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
	Gke() EventarcTriggerDestinationGkeOutputReference
	GkeInput() *EventarcTriggerDestinationGke
	HttpEndpoint() EventarcTriggerDestinationHttpEndpointOutputReference
	HttpEndpointInput() *EventarcTriggerDestinationHttpEndpoint
	InternalValue() *EventarcTriggerDestination
	SetInternalValue(val *EventarcTriggerDestination)
	NetworkConfig() EventarcTriggerDestinationNetworkConfigOutputReference
	NetworkConfigInput() *EventarcTriggerDestinationNetworkConfig
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutCloudRunService(value *EventarcTriggerDestinationCloudRunService)
	PutGke(value *EventarcTriggerDestinationGke)
	PutHttpEndpoint(value *EventarcTriggerDestinationHttpEndpoint)
	PutNetworkConfig(value *EventarcTriggerDestinationNetworkConfig)
	ResetCloudRunService()
	ResetGke()
	ResetHttpEndpoint()
	ResetNetworkConfig()
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

// The jsii proxy struct for EventarcTriggerDestinationOutputReference
type jsiiProxy_EventarcTriggerDestinationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventarcTriggerDestinationOutputReference) CloudFunction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcTriggerDestinationOutputReference) CloudRunService() EventarcTriggerDestinationCloudRunServiceOutputReference {
	var returns EventarcTriggerDestinationCloudRunServiceOutputReference
	_jsii_.Get(
		j,
		"cloudRunService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcTriggerDestinationOutputReference) CloudRunServiceInput() *EventarcTriggerDestinationCloudRunService {
	var returns *EventarcTriggerDestinationCloudRunService
	_jsii_.Get(
		j,
		"cloudRunServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcTriggerDestinationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcTriggerDestinationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcTriggerDestinationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcTriggerDestinationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcTriggerDestinationOutputReference) Gke() EventarcTriggerDestinationGkeOutputReference {
	var returns EventarcTriggerDestinationGkeOutputReference
	_jsii_.Get(
		j,
		"gke",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcTriggerDestinationOutputReference) GkeInput() *EventarcTriggerDestinationGke {
	var returns *EventarcTriggerDestinationGke
	_jsii_.Get(
		j,
		"gkeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcTriggerDestinationOutputReference) HttpEndpoint() EventarcTriggerDestinationHttpEndpointOutputReference {
	var returns EventarcTriggerDestinationHttpEndpointOutputReference
	_jsii_.Get(
		j,
		"httpEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcTriggerDestinationOutputReference) HttpEndpointInput() *EventarcTriggerDestinationHttpEndpoint {
	var returns *EventarcTriggerDestinationHttpEndpoint
	_jsii_.Get(
		j,
		"httpEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcTriggerDestinationOutputReference) InternalValue() *EventarcTriggerDestination {
	var returns *EventarcTriggerDestination
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcTriggerDestinationOutputReference) NetworkConfig() EventarcTriggerDestinationNetworkConfigOutputReference {
	var returns EventarcTriggerDestinationNetworkConfigOutputReference
	_jsii_.Get(
		j,
		"networkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcTriggerDestinationOutputReference) NetworkConfigInput() *EventarcTriggerDestinationNetworkConfig {
	var returns *EventarcTriggerDestinationNetworkConfig
	_jsii_.Get(
		j,
		"networkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcTriggerDestinationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcTriggerDestinationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcTriggerDestinationOutputReference) Workflow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workflow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcTriggerDestinationOutputReference) WorkflowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workflowInput",
		&returns,
	)
	return returns
}


func NewEventarcTriggerDestinationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EventarcTriggerDestinationOutputReference {
	_init_.Initialize()

	if err := validateNewEventarcTriggerDestinationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EventarcTriggerDestinationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.eventarcTrigger.EventarcTriggerDestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEventarcTriggerDestinationOutputReference_Override(e EventarcTriggerDestinationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.eventarcTrigger.EventarcTriggerDestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EventarcTriggerDestinationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventarcTriggerDestinationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventarcTriggerDestinationOutputReference)SetInternalValue(val *EventarcTriggerDestination) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventarcTriggerDestinationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventarcTriggerDestinationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EventarcTriggerDestinationOutputReference)SetWorkflow(val *string) {
	if err := j.validateSetWorkflowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workflow",
		val,
	)
}

func (e *jsiiProxy_EventarcTriggerDestinationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventarcTriggerDestinationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EventarcTriggerDestinationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EventarcTriggerDestinationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EventarcTriggerDestinationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EventarcTriggerDestinationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EventarcTriggerDestinationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EventarcTriggerDestinationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EventarcTriggerDestinationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EventarcTriggerDestinationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EventarcTriggerDestinationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventarcTriggerDestinationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EventarcTriggerDestinationOutputReference) PutCloudRunService(value *EventarcTriggerDestinationCloudRunService) {
	if err := e.validatePutCloudRunServiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putCloudRunService",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventarcTriggerDestinationOutputReference) PutGke(value *EventarcTriggerDestinationGke) {
	if err := e.validatePutGkeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putGke",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventarcTriggerDestinationOutputReference) PutHttpEndpoint(value *EventarcTriggerDestinationHttpEndpoint) {
	if err := e.validatePutHttpEndpointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putHttpEndpoint",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventarcTriggerDestinationOutputReference) PutNetworkConfig(value *EventarcTriggerDestinationNetworkConfig) {
	if err := e.validatePutNetworkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNetworkConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventarcTriggerDestinationOutputReference) ResetCloudRunService() {
	_jsii_.InvokeVoid(
		e,
		"resetCloudRunService",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventarcTriggerDestinationOutputReference) ResetGke() {
	_jsii_.InvokeVoid(
		e,
		"resetGke",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventarcTriggerDestinationOutputReference) ResetHttpEndpoint() {
	_jsii_.InvokeVoid(
		e,
		"resetHttpEndpoint",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventarcTriggerDestinationOutputReference) ResetNetworkConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetNetworkConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventarcTriggerDestinationOutputReference) ResetWorkflow() {
	_jsii_.InvokeVoid(
		e,
		"resetWorkflow",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventarcTriggerDestinationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_EventarcTriggerDestinationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

