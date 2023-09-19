// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxtestcase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v9/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v9/dialogflowcxtestcase/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference interface {
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
	Dtmf() DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputDtmfOutputReference
	DtmfInput() *DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputDtmf
	Event() DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputEventOutputReference
	EventInput() *DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputEvent
	// Experimental.
	Fqn() *string
	InternalValue() *DialogflowCxTestCaseTestCaseConversationTurnsUserInputInput
	SetInternalValue(val *DialogflowCxTestCaseTestCaseConversationTurnsUserInputInput)
	LanguageCode() *string
	SetLanguageCode(val *string)
	LanguageCodeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Text() DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputTextOutputReference
	TextInput() *DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputText
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
	PutDtmf(value *DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputDtmf)
	PutEvent(value *DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputEvent)
	PutText(value *DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputText)
	ResetDtmf()
	ResetEvent()
	ResetLanguageCode()
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

// The jsii proxy struct for DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference
type jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) Dtmf() DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputDtmfOutputReference {
	var returns DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputDtmfOutputReference
	_jsii_.Get(
		j,
		"dtmf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) DtmfInput() *DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputDtmf {
	var returns *DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputDtmf
	_jsii_.Get(
		j,
		"dtmfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) Event() DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputEventOutputReference {
	var returns DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputEventOutputReference
	_jsii_.Get(
		j,
		"event",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) EventInput() *DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputEvent {
	var returns *DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputEvent
	_jsii_.Get(
		j,
		"eventInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) InternalValue() *DialogflowCxTestCaseTestCaseConversationTurnsUserInputInput {
	var returns *DialogflowCxTestCaseTestCaseConversationTurnsUserInputInput
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) LanguageCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) LanguageCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"languageCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) Text() DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputTextOutputReference {
	var returns DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputTextOutputReference
	_jsii_.Get(
		j,
		"text",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) TextInput() *DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputText {
	var returns *DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputText
	_jsii_.Get(
		j,
		"textInput",
		&returns,
	)
	return returns
}


func NewDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxTestCase.DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference_Override(d DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxTestCase.DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference)SetInternalValue(val *DialogflowCxTestCaseTestCaseConversationTurnsUserInputInput) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference)SetLanguageCode(val *string) {
	if err := j.validateSetLanguageCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"languageCode",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) PutDtmf(value *DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputDtmf) {
	if err := d.validatePutDtmfParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDtmf",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) PutEvent(value *DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputEvent) {
	if err := d.validatePutEventParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEvent",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) PutText(value *DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputText) {
	if err := d.validatePutTextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putText",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) ResetDtmf() {
	_jsii_.InvokeVoid(
		d,
		"resetDtmf",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) ResetEvent() {
	_jsii_.InvokeVoid(
		d,
		"resetEvent",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) ResetLanguageCode() {
	_jsii_.InvokeVoid(
		d,
		"resetLanguageCode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) ResetText() {
	_jsii_.InvokeVoid(
		d,
		"resetText",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsUserInputInputOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

