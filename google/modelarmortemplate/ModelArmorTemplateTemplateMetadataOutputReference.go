// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelarmortemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/modelarmortemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ModelArmorTemplateTemplateMetadataOutputReference interface {
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
	CustomLlmResponseSafetyErrorCode() *float64
	SetCustomLlmResponseSafetyErrorCode(val *float64)
	CustomLlmResponseSafetyErrorCodeInput() *float64
	CustomLlmResponseSafetyErrorMessage() *string
	SetCustomLlmResponseSafetyErrorMessage(val *string)
	CustomLlmResponseSafetyErrorMessageInput() *string
	CustomPromptSafetyErrorCode() *float64
	SetCustomPromptSafetyErrorCode(val *float64)
	CustomPromptSafetyErrorCodeInput() *float64
	CustomPromptSafetyErrorMessage() *string
	SetCustomPromptSafetyErrorMessage(val *string)
	CustomPromptSafetyErrorMessageInput() *string
	EnforcementType() *string
	SetEnforcementType(val *string)
	EnforcementTypeInput() *string
	// Experimental.
	Fqn() *string
	IgnorePartialInvocationFailures() interface{}
	SetIgnorePartialInvocationFailures(val interface{})
	IgnorePartialInvocationFailuresInput() interface{}
	InternalValue() *ModelArmorTemplateTemplateMetadata
	SetInternalValue(val *ModelArmorTemplateTemplateMetadata)
	LogSanitizeOperations() interface{}
	SetLogSanitizeOperations(val interface{})
	LogSanitizeOperationsInput() interface{}
	LogTemplateOperations() interface{}
	SetLogTemplateOperations(val interface{})
	LogTemplateOperationsInput() interface{}
	MultiLanguageDetection() ModelArmorTemplateTemplateMetadataMultiLanguageDetectionOutputReference
	MultiLanguageDetectionInput() *ModelArmorTemplateTemplateMetadataMultiLanguageDetection
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
	PutMultiLanguageDetection(value *ModelArmorTemplateTemplateMetadataMultiLanguageDetection)
	ResetCustomLlmResponseSafetyErrorCode()
	ResetCustomLlmResponseSafetyErrorMessage()
	ResetCustomPromptSafetyErrorCode()
	ResetCustomPromptSafetyErrorMessage()
	ResetEnforcementType()
	ResetIgnorePartialInvocationFailures()
	ResetLogSanitizeOperations()
	ResetLogTemplateOperations()
	ResetMultiLanguageDetection()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ModelArmorTemplateTemplateMetadataOutputReference
type jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) CustomLlmResponseSafetyErrorCode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"customLlmResponseSafetyErrorCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) CustomLlmResponseSafetyErrorCodeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"customLlmResponseSafetyErrorCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) CustomLlmResponseSafetyErrorMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customLlmResponseSafetyErrorMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) CustomLlmResponseSafetyErrorMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customLlmResponseSafetyErrorMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) CustomPromptSafetyErrorCode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"customPromptSafetyErrorCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) CustomPromptSafetyErrorCodeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"customPromptSafetyErrorCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) CustomPromptSafetyErrorMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customPromptSafetyErrorMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) CustomPromptSafetyErrorMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customPromptSafetyErrorMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) EnforcementType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enforcementType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) EnforcementTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enforcementTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) IgnorePartialInvocationFailures() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignorePartialInvocationFailures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) IgnorePartialInvocationFailuresInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignorePartialInvocationFailuresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) InternalValue() *ModelArmorTemplateTemplateMetadata {
	var returns *ModelArmorTemplateTemplateMetadata
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) LogSanitizeOperations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logSanitizeOperations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) LogSanitizeOperationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logSanitizeOperationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) LogTemplateOperations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logTemplateOperations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) LogTemplateOperationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logTemplateOperationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) MultiLanguageDetection() ModelArmorTemplateTemplateMetadataMultiLanguageDetectionOutputReference {
	var returns ModelArmorTemplateTemplateMetadataMultiLanguageDetectionOutputReference
	_jsii_.Get(
		j,
		"multiLanguageDetection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) MultiLanguageDetectionInput() *ModelArmorTemplateTemplateMetadataMultiLanguageDetection {
	var returns *ModelArmorTemplateTemplateMetadataMultiLanguageDetection
	_jsii_.Get(
		j,
		"multiLanguageDetectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewModelArmorTemplateTemplateMetadataOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ModelArmorTemplateTemplateMetadataOutputReference {
	_init_.Initialize()

	if err := validateNewModelArmorTemplateTemplateMetadataOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.modelArmorTemplate.ModelArmorTemplateTemplateMetadataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewModelArmorTemplateTemplateMetadataOutputReference_Override(m ModelArmorTemplateTemplateMetadataOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.modelArmorTemplate.ModelArmorTemplateTemplateMetadataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference)SetCustomLlmResponseSafetyErrorCode(val *float64) {
	if err := j.validateSetCustomLlmResponseSafetyErrorCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customLlmResponseSafetyErrorCode",
		val,
	)
}

func (j *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference)SetCustomLlmResponseSafetyErrorMessage(val *string) {
	if err := j.validateSetCustomLlmResponseSafetyErrorMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customLlmResponseSafetyErrorMessage",
		val,
	)
}

func (j *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference)SetCustomPromptSafetyErrorCode(val *float64) {
	if err := j.validateSetCustomPromptSafetyErrorCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customPromptSafetyErrorCode",
		val,
	)
}

func (j *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference)SetCustomPromptSafetyErrorMessage(val *string) {
	if err := j.validateSetCustomPromptSafetyErrorMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customPromptSafetyErrorMessage",
		val,
	)
}

func (j *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference)SetEnforcementType(val *string) {
	if err := j.validateSetEnforcementTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforcementType",
		val,
	)
}

func (j *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference)SetIgnorePartialInvocationFailures(val interface{}) {
	if err := j.validateSetIgnorePartialInvocationFailuresParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignorePartialInvocationFailures",
		val,
	)
}

func (j *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference)SetInternalValue(val *ModelArmorTemplateTemplateMetadata) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference)SetLogSanitizeOperations(val interface{}) {
	if err := j.validateSetLogSanitizeOperationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logSanitizeOperations",
		val,
	)
}

func (j *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference)SetLogTemplateOperations(val interface{}) {
	if err := j.validateSetLogTemplateOperationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logTemplateOperations",
		val,
	)
}

func (j *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) PutMultiLanguageDetection(value *ModelArmorTemplateTemplateMetadataMultiLanguageDetection) {
	if err := m.validatePutMultiLanguageDetectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putMultiLanguageDetection",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) ResetCustomLlmResponseSafetyErrorCode() {
	_jsii_.InvokeVoid(
		m,
		"resetCustomLlmResponseSafetyErrorCode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) ResetCustomLlmResponseSafetyErrorMessage() {
	_jsii_.InvokeVoid(
		m,
		"resetCustomLlmResponseSafetyErrorMessage",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) ResetCustomPromptSafetyErrorCode() {
	_jsii_.InvokeVoid(
		m,
		"resetCustomPromptSafetyErrorCode",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) ResetCustomPromptSafetyErrorMessage() {
	_jsii_.InvokeVoid(
		m,
		"resetCustomPromptSafetyErrorMessage",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) ResetEnforcementType() {
	_jsii_.InvokeVoid(
		m,
		"resetEnforcementType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) ResetIgnorePartialInvocationFailures() {
	_jsii_.InvokeVoid(
		m,
		"resetIgnorePartialInvocationFailures",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) ResetLogSanitizeOperations() {
	_jsii_.InvokeVoid(
		m,
		"resetLogSanitizeOperations",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) ResetLogTemplateOperations() {
	_jsii_.InvokeVoid(
		m,
		"resetLogTemplateOperations",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) ResetMultiLanguageDetection() {
	_jsii_.InvokeVoid(
		m,
		"resetMultiLanguageDetection",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorTemplateTemplateMetadataOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

