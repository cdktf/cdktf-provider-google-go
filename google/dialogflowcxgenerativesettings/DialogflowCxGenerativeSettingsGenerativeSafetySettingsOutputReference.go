// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxgenerativesettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/dialogflowcxgenerativesettings/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference interface {
	cdktf.ComplexObject
	BannedPhrases() DialogflowCxGenerativeSettingsGenerativeSafetySettingsBannedPhrasesList
	BannedPhrasesInput() interface{}
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
	DefaultBannedPhraseMatchStrategy() *string
	SetDefaultBannedPhraseMatchStrategy(val *string)
	DefaultBannedPhraseMatchStrategyInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DialogflowCxGenerativeSettingsGenerativeSafetySettings
	SetInternalValue(val *DialogflowCxGenerativeSettingsGenerativeSafetySettings)
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
	PutBannedPhrases(value interface{})
	ResetBannedPhrases()
	ResetDefaultBannedPhraseMatchStrategy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference
type jsiiProxy_DialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) BannedPhrases() DialogflowCxGenerativeSettingsGenerativeSafetySettingsBannedPhrasesList {
	var returns DialogflowCxGenerativeSettingsGenerativeSafetySettingsBannedPhrasesList
	_jsii_.Get(
		j,
		"bannedPhrases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) BannedPhrasesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bannedPhrasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) DefaultBannedPhraseMatchStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultBannedPhraseMatchStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) DefaultBannedPhraseMatchStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultBannedPhraseMatchStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) InternalValue() *DialogflowCxGenerativeSettingsGenerativeSafetySettings {
	var returns *DialogflowCxGenerativeSettingsGenerativeSafetySettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxGenerativeSettings.DialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference_Override(d DialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxGenerativeSettings.DialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference)SetDefaultBannedPhraseMatchStrategy(val *string) {
	if err := j.validateSetDefaultBannedPhraseMatchStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultBannedPhraseMatchStrategy",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference)SetInternalValue(val *DialogflowCxGenerativeSettingsGenerativeSafetySettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) PutBannedPhrases(value interface{}) {
	if err := d.validatePutBannedPhrasesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putBannedPhrases",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) ResetBannedPhrases() {
	_jsii_.InvokeVoid(
		d,
		"resetBannedPhrases",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) ResetDefaultBannedPhraseMatchStrategy() {
	_jsii_.InvokeVoid(
		d,
		"resetDefaultBannedPhraseMatchStrategy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DialogflowCxGenerativeSettingsGenerativeSafetySettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

