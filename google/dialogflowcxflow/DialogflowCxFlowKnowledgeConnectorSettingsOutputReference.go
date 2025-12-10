// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/dialogflowcxflow/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DialogflowCxFlowKnowledgeConnectorSettingsOutputReference interface {
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
	DataStoreConnections() DialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsList
	DataStoreConnectionsInput() interface{}
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *DialogflowCxFlowKnowledgeConnectorSettings
	SetInternalValue(val *DialogflowCxFlowKnowledgeConnectorSettings)
	TargetFlow() *string
	SetTargetFlow(val *string)
	TargetFlowInput() *string
	TargetPage() *string
	SetTargetPage(val *string)
	TargetPageInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TriggerFulfillment() DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentOutputReference
	TriggerFulfillmentInput() *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillment
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutDataStoreConnections(value interface{})
	PutTriggerFulfillment(value *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillment)
	ResetDataStoreConnections()
	ResetEnabled()
	ResetTargetFlow()
	ResetTargetPage()
	ResetTriggerFulfillment()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DialogflowCxFlowKnowledgeConnectorSettingsOutputReference
type jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference) DataStoreConnections() DialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsList {
	var returns DialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsList
	_jsii_.Get(
		j,
		"dataStoreConnections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference) DataStoreConnectionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataStoreConnectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference) InternalValue() *DialogflowCxFlowKnowledgeConnectorSettings {
	var returns *DialogflowCxFlowKnowledgeConnectorSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference) TargetFlow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetFlow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference) TargetFlowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetFlowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference) TargetPage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetPage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference) TargetPageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetPageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference) TriggerFulfillment() DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentOutputReference {
	var returns DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillmentOutputReference
	_jsii_.Get(
		j,
		"triggerFulfillment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference) TriggerFulfillmentInput() *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillment {
	var returns *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillment
	_jsii_.Get(
		j,
		"triggerFulfillmentInput",
		&returns,
	)
	return returns
}


func NewDialogflowCxFlowKnowledgeConnectorSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DialogflowCxFlowKnowledgeConnectorSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowCxFlowKnowledgeConnectorSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxFlow.DialogflowCxFlowKnowledgeConnectorSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDialogflowCxFlowKnowledgeConnectorSettingsOutputReference_Override(d DialogflowCxFlowKnowledgeConnectorSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxFlow.DialogflowCxFlowKnowledgeConnectorSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference)SetInternalValue(val *DialogflowCxFlowKnowledgeConnectorSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference)SetTargetFlow(val *string) {
	if err := j.validateSetTargetFlowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetFlow",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference)SetTargetPage(val *string) {
	if err := j.validateSetTargetPageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetPage",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference) PutDataStoreConnections(value interface{}) {
	if err := d.validatePutDataStoreConnectionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDataStoreConnections",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference) PutTriggerFulfillment(value *DialogflowCxFlowKnowledgeConnectorSettingsTriggerFulfillment) {
	if err := d.validatePutTriggerFulfillmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTriggerFulfillment",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference) ResetDataStoreConnections() {
	_jsii_.InvokeVoid(
		d,
		"resetDataStoreConnections",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference) ResetTargetFlow() {
	_jsii_.InvokeVoid(
		d,
		"resetTargetFlow",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference) ResetTargetPage() {
	_jsii_.InvokeVoid(
		d,
		"resetTargetPage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference) ResetTriggerFulfillment() {
	_jsii_.InvokeVoid(
		d,
		"resetTriggerFulfillment",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxFlowKnowledgeConnectorSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

