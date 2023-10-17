// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxsecuritysettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v12/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v12/dialogflowcxsecuritysettings/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DialogflowCxSecuritySettingsAudioExportSettingsOutputReference interface {
	cdktf.ComplexObject
	AudioExportPattern() *string
	SetAudioExportPattern(val *string)
	AudioExportPatternInput() *string
	AudioFormat() *string
	SetAudioFormat(val *string)
	AudioFormatInput() *string
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
	EnableAudioRedaction() interface{}
	SetEnableAudioRedaction(val interface{})
	EnableAudioRedactionInput() interface{}
	// Experimental.
	Fqn() *string
	GcsBucket() *string
	SetGcsBucket(val *string)
	GcsBucketInput() *string
	InternalValue() *DialogflowCxSecuritySettingsAudioExportSettings
	SetInternalValue(val *DialogflowCxSecuritySettingsAudioExportSettings)
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
	ResetAudioExportPattern()
	ResetAudioFormat()
	ResetEnableAudioRedaction()
	ResetGcsBucket()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DialogflowCxSecuritySettingsAudioExportSettingsOutputReference
type jsiiProxy_DialogflowCxSecuritySettingsAudioExportSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DialogflowCxSecuritySettingsAudioExportSettingsOutputReference) AudioExportPattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioExportPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettingsAudioExportSettingsOutputReference) AudioExportPatternInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioExportPatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettingsAudioExportSettingsOutputReference) AudioFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettingsAudioExportSettingsOutputReference) AudioFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettingsAudioExportSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettingsAudioExportSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettingsAudioExportSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettingsAudioExportSettingsOutputReference) EnableAudioRedaction() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAudioRedaction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettingsAudioExportSettingsOutputReference) EnableAudioRedactionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAudioRedactionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettingsAudioExportSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettingsAudioExportSettingsOutputReference) GcsBucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcsBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettingsAudioExportSettingsOutputReference) GcsBucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcsBucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettingsAudioExportSettingsOutputReference) InternalValue() *DialogflowCxSecuritySettingsAudioExportSettings {
	var returns *DialogflowCxSecuritySettingsAudioExportSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettingsAudioExportSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxSecuritySettingsAudioExportSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDialogflowCxSecuritySettingsAudioExportSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DialogflowCxSecuritySettingsAudioExportSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowCxSecuritySettingsAudioExportSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowCxSecuritySettingsAudioExportSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxSecuritySettings.DialogflowCxSecuritySettingsAudioExportSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDialogflowCxSecuritySettingsAudioExportSettingsOutputReference_Override(d DialogflowCxSecuritySettingsAudioExportSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxSecuritySettings.DialogflowCxSecuritySettingsAudioExportSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DialogflowCxSecuritySettingsAudioExportSettingsOutputReference)SetAudioExportPattern(val *string) {
	if err := j.validateSetAudioExportPatternParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audioExportPattern",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxSecuritySettingsAudioExportSettingsOutputReference)SetAudioFormat(val *string) {
	if err := j.validateSetAudioFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audioFormat",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxSecuritySettingsAudioExportSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxSecuritySettingsAudioExportSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxSecuritySettingsAudioExportSettingsOutputReference)SetEnableAudioRedaction(val interface{}) {
	if err := j.validateSetEnableAudioRedactionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAudioRedaction",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxSecuritySettingsAudioExportSettingsOutputReference)SetGcsBucket(val *string) {
	if err := j.validateSetGcsBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gcsBucket",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxSecuritySettingsAudioExportSettingsOutputReference)SetInternalValue(val *DialogflowCxSecuritySettingsAudioExportSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxSecuritySettingsAudioExportSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxSecuritySettingsAudioExportSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DialogflowCxSecuritySettingsAudioExportSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxSecuritySettingsAudioExportSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DialogflowCxSecuritySettingsAudioExportSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowCxSecuritySettingsAudioExportSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DialogflowCxSecuritySettingsAudioExportSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DialogflowCxSecuritySettingsAudioExportSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DialogflowCxSecuritySettingsAudioExportSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DialogflowCxSecuritySettingsAudioExportSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DialogflowCxSecuritySettingsAudioExportSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DialogflowCxSecuritySettingsAudioExportSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DialogflowCxSecuritySettingsAudioExportSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxSecuritySettingsAudioExportSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowCxSecuritySettingsAudioExportSettingsOutputReference) ResetAudioExportPattern() {
	_jsii_.InvokeVoid(
		d,
		"resetAudioExportPattern",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxSecuritySettingsAudioExportSettingsOutputReference) ResetAudioFormat() {
	_jsii_.InvokeVoid(
		d,
		"resetAudioFormat",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxSecuritySettingsAudioExportSettingsOutputReference) ResetEnableAudioRedaction() {
	_jsii_.InvokeVoid(
		d,
		"resetEnableAudioRedaction",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxSecuritySettingsAudioExportSettingsOutputReference) ResetGcsBucket() {
	_jsii_.InvokeVoid(
		d,
		"resetGcsBucket",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxSecuritySettingsAudioExportSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DialogflowCxSecuritySettingsAudioExportSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

