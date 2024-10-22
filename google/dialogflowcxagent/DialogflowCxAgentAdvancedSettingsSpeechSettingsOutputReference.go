// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/dialogflowcxagent/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference interface {
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
	EndpointerSensitivity() *float64
	SetEndpointerSensitivity(val *float64)
	EndpointerSensitivityInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *DialogflowCxAgentAdvancedSettingsSpeechSettings
	SetInternalValue(val *DialogflowCxAgentAdvancedSettingsSpeechSettings)
	Models() *map[string]*string
	SetModels(val *map[string]*string)
	ModelsInput() *map[string]*string
	NoSpeechTimeout() *string
	SetNoSpeechTimeout(val *string)
	NoSpeechTimeoutInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UseTimeoutBasedEndpointing() interface{}
	SetUseTimeoutBasedEndpointing(val interface{})
	UseTimeoutBasedEndpointingInput() interface{}
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
	ResetEndpointerSensitivity()
	ResetModels()
	ResetNoSpeechTimeout()
	ResetUseTimeoutBasedEndpointing()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference
type jsiiProxy_DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference) EndpointerSensitivity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"endpointerSensitivity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference) EndpointerSensitivityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"endpointerSensitivityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference) InternalValue() *DialogflowCxAgentAdvancedSettingsSpeechSettings {
	var returns *DialogflowCxAgentAdvancedSettingsSpeechSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference) Models() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"models",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference) ModelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"modelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference) NoSpeechTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noSpeechTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference) NoSpeechTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noSpeechTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference) UseTimeoutBasedEndpointing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useTimeoutBasedEndpointing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference) UseTimeoutBasedEndpointingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useTimeoutBasedEndpointingInput",
		&returns,
	)
	return returns
}


func NewDialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxAgent.DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference_Override(d DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxAgent.DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference)SetEndpointerSensitivity(val *float64) {
	if err := j.validateSetEndpointerSensitivityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointerSensitivity",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference)SetInternalValue(val *DialogflowCxAgentAdvancedSettingsSpeechSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference)SetModels(val *map[string]*string) {
	if err := j.validateSetModelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"models",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference)SetNoSpeechTimeout(val *string) {
	if err := j.validateSetNoSpeechTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noSpeechTimeout",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference)SetUseTimeoutBasedEndpointing(val interface{}) {
	if err := j.validateSetUseTimeoutBasedEndpointingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useTimeoutBasedEndpointing",
		val,
	)
}

func (d *jsiiProxy_DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference) ResetEndpointerSensitivity() {
	_jsii_.InvokeVoid(
		d,
		"resetEndpointerSensitivity",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference) ResetModels() {
	_jsii_.InvokeVoid(
		d,
		"resetModels",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference) ResetNoSpeechTimeout() {
	_jsii_.InvokeVoid(
		d,
		"resetNoSpeechTimeout",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference) ResetUseTimeoutBasedEndpointing() {
	_jsii_.InvokeVoid(
		d,
		"resetUseTimeoutBasedEndpointing",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DialogflowCxAgentAdvancedSettingsSpeechSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

