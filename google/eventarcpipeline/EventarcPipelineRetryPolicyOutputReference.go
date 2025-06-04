// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventarcpipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/eventarcpipeline/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EventarcPipelineRetryPolicyOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *EventarcPipelineRetryPolicy
	SetInternalValue(val *EventarcPipelineRetryPolicy)
	MaxAttempts() *float64
	SetMaxAttempts(val *float64)
	MaxAttemptsInput() *float64
	MaxRetryDelay() *string
	SetMaxRetryDelay(val *string)
	MaxRetryDelayInput() *string
	MinRetryDelay() *string
	SetMinRetryDelay(val *string)
	MinRetryDelayInput() *string
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
	ResetMaxAttempts()
	ResetMaxRetryDelay()
	ResetMinRetryDelay()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventarcPipelineRetryPolicyOutputReference
type jsiiProxy_EventarcPipelineRetryPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventarcPipelineRetryPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcPipelineRetryPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcPipelineRetryPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcPipelineRetryPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcPipelineRetryPolicyOutputReference) InternalValue() *EventarcPipelineRetryPolicy {
	var returns *EventarcPipelineRetryPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcPipelineRetryPolicyOutputReference) MaxAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcPipelineRetryPolicyOutputReference) MaxAttemptsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAttemptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcPipelineRetryPolicyOutputReference) MaxRetryDelay() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxRetryDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcPipelineRetryPolicyOutputReference) MaxRetryDelayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxRetryDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcPipelineRetryPolicyOutputReference) MinRetryDelay() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minRetryDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcPipelineRetryPolicyOutputReference) MinRetryDelayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minRetryDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcPipelineRetryPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventarcPipelineRetryPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEventarcPipelineRetryPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EventarcPipelineRetryPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewEventarcPipelineRetryPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EventarcPipelineRetryPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.eventarcPipeline.EventarcPipelineRetryPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEventarcPipelineRetryPolicyOutputReference_Override(e EventarcPipelineRetryPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.eventarcPipeline.EventarcPipelineRetryPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EventarcPipelineRetryPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventarcPipelineRetryPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventarcPipelineRetryPolicyOutputReference)SetInternalValue(val *EventarcPipelineRetryPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventarcPipelineRetryPolicyOutputReference)SetMaxAttempts(val *float64) {
	if err := j.validateSetMaxAttemptsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxAttempts",
		val,
	)
}

func (j *jsiiProxy_EventarcPipelineRetryPolicyOutputReference)SetMaxRetryDelay(val *string) {
	if err := j.validateSetMaxRetryDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRetryDelay",
		val,
	)
}

func (j *jsiiProxy_EventarcPipelineRetryPolicyOutputReference)SetMinRetryDelay(val *string) {
	if err := j.validateSetMinRetryDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minRetryDelay",
		val,
	)
}

func (j *jsiiProxy_EventarcPipelineRetryPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventarcPipelineRetryPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EventarcPipelineRetryPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventarcPipelineRetryPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EventarcPipelineRetryPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EventarcPipelineRetryPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EventarcPipelineRetryPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EventarcPipelineRetryPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EventarcPipelineRetryPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EventarcPipelineRetryPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EventarcPipelineRetryPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EventarcPipelineRetryPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EventarcPipelineRetryPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventarcPipelineRetryPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EventarcPipelineRetryPolicyOutputReference) ResetMaxAttempts() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxAttempts",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventarcPipelineRetryPolicyOutputReference) ResetMaxRetryDelay() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxRetryDelay",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventarcPipelineRetryPolicyOutputReference) ResetMinRetryDelay() {
	_jsii_.InvokeVoid(
		e,
		"resetMinRetryDelay",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventarcPipelineRetryPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_EventarcPipelineRetryPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

