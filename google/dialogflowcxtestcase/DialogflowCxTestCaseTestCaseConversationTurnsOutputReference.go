// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxtestcase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v9/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v9/dialogflowcxtestcase/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DialogflowCxTestCaseTestCaseConversationTurnsOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserInput() DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference
	UserInputInput() *DialogflowCxTestCaseTestCaseConversationTurnsUserInput
	VirtualAgentOutput() DialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference
	VirtualAgentOutputInput() *DialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutput
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
	PutUserInput(value *DialogflowCxTestCaseTestCaseConversationTurnsUserInput)
	PutVirtualAgentOutput(value *DialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutput)
	ResetUserInput()
	ResetVirtualAgentOutput()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DialogflowCxTestCaseTestCaseConversationTurnsOutputReference
type jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsOutputReference) UserInput() DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference {
	var returns DialogflowCxTestCaseTestCaseConversationTurnsUserInputOutputReference
	_jsii_.Get(
		j,
		"userInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsOutputReference) UserInputInput() *DialogflowCxTestCaseTestCaseConversationTurnsUserInput {
	var returns *DialogflowCxTestCaseTestCaseConversationTurnsUserInput
	_jsii_.Get(
		j,
		"userInputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsOutputReference) VirtualAgentOutput() DialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference {
	var returns DialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutputOutputReference
	_jsii_.Get(
		j,
		"virtualAgentOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsOutputReference) VirtualAgentOutputInput() *DialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutput {
	var returns *DialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutput
	_jsii_.Get(
		j,
		"virtualAgentOutputInput",
		&returns,
	)
	return returns
}


func NewDialogflowCxTestCaseTestCaseConversationTurnsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DialogflowCxTestCaseTestCaseConversationTurnsOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowCxTestCaseTestCaseConversationTurnsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxTestCase.DialogflowCxTestCaseTestCaseConversationTurnsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDialogflowCxTestCaseTestCaseConversationTurnsOutputReference_Override(d DialogflowCxTestCaseTestCaseConversationTurnsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxTestCase.DialogflowCxTestCaseTestCaseConversationTurnsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsOutputReference) PutUserInput(value *DialogflowCxTestCaseTestCaseConversationTurnsUserInput) {
	if err := d.validatePutUserInputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putUserInput",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsOutputReference) PutVirtualAgentOutput(value *DialogflowCxTestCaseTestCaseConversationTurnsVirtualAgentOutput) {
	if err := d.validatePutVirtualAgentOutputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putVirtualAgentOutput",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsOutputReference) ResetUserInput() {
	_jsii_.InvokeVoid(
		d,
		"resetUserInput",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsOutputReference) ResetVirtualAgentOutput() {
	_jsii_.InvokeVoid(
		d,
		"resetVirtualAgentOutput",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DialogflowCxTestCaseTestCaseConversationTurnsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

