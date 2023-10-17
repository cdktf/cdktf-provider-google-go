// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxtestcase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v12/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v12/dialogflowcxtestcase/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DialogflowCxTestCaseLastTestResultConversationTurnsVirtualAgentOutputTriggeredIntentList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) DialogflowCxTestCaseLastTestResultConversationTurnsVirtualAgentOutputTriggeredIntentOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DialogflowCxTestCaseLastTestResultConversationTurnsVirtualAgentOutputTriggeredIntentList
type jsiiProxy_DialogflowCxTestCaseLastTestResultConversationTurnsVirtualAgentOutputTriggeredIntentList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_DialogflowCxTestCaseLastTestResultConversationTurnsVirtualAgentOutputTriggeredIntentList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxTestCaseLastTestResultConversationTurnsVirtualAgentOutputTriggeredIntentList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxTestCaseLastTestResultConversationTurnsVirtualAgentOutputTriggeredIntentList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxTestCaseLastTestResultConversationTurnsVirtualAgentOutputTriggeredIntentList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxTestCaseLastTestResultConversationTurnsVirtualAgentOutputTriggeredIntentList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDialogflowCxTestCaseLastTestResultConversationTurnsVirtualAgentOutputTriggeredIntentList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DialogflowCxTestCaseLastTestResultConversationTurnsVirtualAgentOutputTriggeredIntentList {
	_init_.Initialize()

	if err := validateNewDialogflowCxTestCaseLastTestResultConversationTurnsVirtualAgentOutputTriggeredIntentListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowCxTestCaseLastTestResultConversationTurnsVirtualAgentOutputTriggeredIntentList{}

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxTestCase.DialogflowCxTestCaseLastTestResultConversationTurnsVirtualAgentOutputTriggeredIntentList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDialogflowCxTestCaseLastTestResultConversationTurnsVirtualAgentOutputTriggeredIntentList_Override(d DialogflowCxTestCaseLastTestResultConversationTurnsVirtualAgentOutputTriggeredIntentList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxTestCase.DialogflowCxTestCaseLastTestResultConversationTurnsVirtualAgentOutputTriggeredIntentList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DialogflowCxTestCaseLastTestResultConversationTurnsVirtualAgentOutputTriggeredIntentList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxTestCaseLastTestResultConversationTurnsVirtualAgentOutputTriggeredIntentList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxTestCaseLastTestResultConversationTurnsVirtualAgentOutputTriggeredIntentList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DialogflowCxTestCaseLastTestResultConversationTurnsVirtualAgentOutputTriggeredIntentList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxTestCaseLastTestResultConversationTurnsVirtualAgentOutputTriggeredIntentList) Get(index *float64) DialogflowCxTestCaseLastTestResultConversationTurnsVirtualAgentOutputTriggeredIntentOutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DialogflowCxTestCaseLastTestResultConversationTurnsVirtualAgentOutputTriggeredIntentOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxTestCaseLastTestResultConversationTurnsVirtualAgentOutputTriggeredIntentList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DialogflowCxTestCaseLastTestResultConversationTurnsVirtualAgentOutputTriggeredIntentList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

