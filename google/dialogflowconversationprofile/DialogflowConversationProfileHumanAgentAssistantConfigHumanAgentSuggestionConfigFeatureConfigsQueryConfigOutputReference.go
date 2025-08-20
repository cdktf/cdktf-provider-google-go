// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dialogflowconversationprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/dialogflowconversationprofile/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference interface {
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
	ConfidenceThreshold() *float64
	SetConfidenceThreshold(val *float64)
	ConfidenceThresholdInput() *float64
	ContextFilterSettings() DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigContextFilterSettingsOutputReference
	ContextFilterSettingsInput() *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigContextFilterSettings
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DialogflowQuerySource() DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigDialogflowQuerySourceOutputReference
	DialogflowQuerySourceInput() *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigDialogflowQuerySource
	// Experimental.
	Fqn() *string
	InternalValue() *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfig
	SetInternalValue(val *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfig)
	MaxResults() *float64
	SetMaxResults(val *float64)
	MaxResultsInput() *float64
	Sections() DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigSectionsOutputReference
	SectionsInput() *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigSections
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
	PutContextFilterSettings(value *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigContextFilterSettings)
	PutDialogflowQuerySource(value *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigDialogflowQuerySource)
	PutSections(value *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigSections)
	ResetConfidenceThreshold()
	ResetContextFilterSettings()
	ResetDialogflowQuerySource()
	ResetMaxResults()
	ResetSections()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference
type jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference) ConfidenceThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"confidenceThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference) ConfidenceThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"confidenceThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference) ContextFilterSettings() DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigContextFilterSettingsOutputReference {
	var returns DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigContextFilterSettingsOutputReference
	_jsii_.Get(
		j,
		"contextFilterSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference) ContextFilterSettingsInput() *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigContextFilterSettings {
	var returns *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigContextFilterSettings
	_jsii_.Get(
		j,
		"contextFilterSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference) DialogflowQuerySource() DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigDialogflowQuerySourceOutputReference {
	var returns DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigDialogflowQuerySourceOutputReference
	_jsii_.Get(
		j,
		"dialogflowQuerySource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference) DialogflowQuerySourceInput() *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigDialogflowQuerySource {
	var returns *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigDialogflowQuerySource
	_jsii_.Get(
		j,
		"dialogflowQuerySourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference) InternalValue() *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfig {
	var returns *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference) MaxResults() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxResults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference) MaxResultsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxResultsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference) Sections() DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigSectionsOutputReference {
	var returns DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigSectionsOutputReference
	_jsii_.Get(
		j,
		"sections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference) SectionsInput() *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigSections {
	var returns *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigSections
	_jsii_.Get(
		j,
		"sectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference_Override(d DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowConversationProfile.DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference)SetConfidenceThreshold(val *float64) {
	if err := j.validateSetConfidenceThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"confidenceThreshold",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference)SetInternalValue(val *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference)SetMaxResults(val *float64) {
	if err := j.validateSetMaxResultsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxResults",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference) PutContextFilterSettings(value *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigContextFilterSettings) {
	if err := d.validatePutContextFilterSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putContextFilterSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference) PutDialogflowQuerySource(value *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigDialogflowQuerySource) {
	if err := d.validatePutDialogflowQuerySourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDialogflowQuerySource",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference) PutSections(value *DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigSections) {
	if err := d.validatePutSectionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSections",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference) ResetConfidenceThreshold() {
	_jsii_.InvokeVoid(
		d,
		"resetConfidenceThreshold",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference) ResetContextFilterSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetContextFilterSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference) ResetDialogflowQuerySource() {
	_jsii_.InvokeVoid(
		d,
		"resetDialogflowQuerySource",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference) ResetMaxResults() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxResults",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference) ResetSections() {
	_jsii_.InvokeVoid(
		d,
		"resetSections",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DialogflowConversationProfileHumanAgentAssistantConfigHumanAgentSuggestionConfigFeatureConfigsQueryConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

