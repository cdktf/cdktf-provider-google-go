// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxtestcase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/dialogflowcxtestcase/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference interface {
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
	EnableSentimentAnalysis() interface{}
	SetEnableSentimentAnalysis(val interface{})
	EnableSentimentAnalysisInput() interface{}
	// Experimental.
	Fqn() *string
	InjectedParameters() *string
	SetInjectedParameters(val *string)
	InjectedParametersInput() *string
	Input() DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference
	InputInput() *DialogflowCxTestCaseTestCaseConversationTurnsUserInputInput
	InternalValue() *DialogflowCxTestCaseTestCaseConversationTurnsUserInput
	SetInternalValue(val *DialogflowCxTestCaseTestCaseConversationTurnsUserInput)
	IsWebhookEnabled() interface{}
	SetIsWebhookEnabled(val interface{})
	IsWebhookEnabledInput() interface{}
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
	PutInput(value *DialogflowCxTestCaseTestCaseConversationTurnsUserInputInput)
	ResetEnableSentimentAnalysis()
	ResetInjectedParameters()
	ResetInput()
	ResetIsWebhookEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference
type jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference) EnableSentimentAnalysis() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSentimentAnalysis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference) EnableSentimentAnalysisInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableSentimentAnalysisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference) InjectedParameters() *string {
	var returns *string
	_jsii_.Get(
		j,
		"injectedParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference) InjectedParametersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"injectedParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference) Input() DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference {
	var returns DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference
	_jsii_.Get(
		j,
		"input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference) InputInput() *DialogflowCxTestCaseTestCaseConversationTurnsUserInputInput {
	var returns *DialogflowCxTestCaseTestCaseConversationTurnsUserInputInput
	_jsii_.Get(
		j,
		"inputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference) InternalValue() *DialogflowCxTestCaseTestCaseConversationTurnsUserInput {
	var returns *DialogflowCxTestCaseTestCaseConversationTurnsUserInput
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference) IsWebhookEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isWebhookEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference) IsWebhookEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isWebhookEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxTestCase.DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference_Override(d DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxTestCase.DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference)SetEnableSentimentAnalysis(val interface{}) {
	if err := j.validateSetEnableSentimentAnalysisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableSentimentAnalysis",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference)SetInjectedParameters(val *string) {
	if err := j.validateSetInjectedParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"injectedParameters",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference)SetInternalValue(val *DialogflowCxTestCaseTestCaseConversationTurnsUserInput) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference)SetIsWebhookEnabled(val interface{}) {
	if err := j.validateSetIsWebhookEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isWebhookEnabled",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference) PutInput(value *DialogflowCxTestCaseTestCaseConversationTurnsUserInputInput) {
	if err := d.validatePutInputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putInput",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference) ResetEnableSentimentAnalysis() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableSentimentAnalysis",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference) ResetInjectedParameters() {
	_jsii_.InvokeVoid(
		d,
		"resetInjectedParameters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference) ResetInput() {
	_jsii_.InvokeVoid(
		d,
		"resetInput",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference) ResetIsWebhookEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetIsWebhookEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

