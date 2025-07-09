// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package contactcenterinsightsanalysisrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/contactcenterinsightsanalysisrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference interface {
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
	InternalValue() *ContactCenterInsightsAnalysisRuleAnnotatorSelector
	SetInternalValue(val *ContactCenterInsightsAnalysisRuleAnnotatorSelector)
	IssueModels() *[]*string
	SetIssueModels(val *[]*string)
	IssueModelsInput() *[]*string
	PhraseMatchers() *[]*string
	SetPhraseMatchers(val *[]*string)
	PhraseMatchersInput() *[]*string
	QaConfig() ContactCenterInsightsAnalysisRuleAnnotatorSelectorQaConfigOutputReference
	QaConfigInput() *ContactCenterInsightsAnalysisRuleAnnotatorSelectorQaConfig
	RunEntityAnnotator() interface{}
	SetRunEntityAnnotator(val interface{})
	RunEntityAnnotatorInput() interface{}
	RunIntentAnnotator() interface{}
	SetRunIntentAnnotator(val interface{})
	RunIntentAnnotatorInput() interface{}
	RunInterruptionAnnotator() interface{}
	SetRunInterruptionAnnotator(val interface{})
	RunInterruptionAnnotatorInput() interface{}
	RunIssueModelAnnotator() interface{}
	SetRunIssueModelAnnotator(val interface{})
	RunIssueModelAnnotatorInput() interface{}
	RunPhraseMatcherAnnotator() interface{}
	SetRunPhraseMatcherAnnotator(val interface{})
	RunPhraseMatcherAnnotatorInput() interface{}
	RunQaAnnotator() interface{}
	SetRunQaAnnotator(val interface{})
	RunQaAnnotatorInput() interface{}
	RunSentimentAnnotator() interface{}
	SetRunSentimentAnnotator(val interface{})
	RunSentimentAnnotatorInput() interface{}
	RunSilenceAnnotator() interface{}
	SetRunSilenceAnnotator(val interface{})
	RunSilenceAnnotatorInput() interface{}
	RunSummarizationAnnotator() interface{}
	SetRunSummarizationAnnotator(val interface{})
	RunSummarizationAnnotatorInput() interface{}
	SummarizationConfig() ContactCenterInsightsAnalysisRuleAnnotatorSelectorSummarizationConfigOutputReference
	SummarizationConfigInput() *ContactCenterInsightsAnalysisRuleAnnotatorSelectorSummarizationConfig
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
	PutQaConfig(value *ContactCenterInsightsAnalysisRuleAnnotatorSelectorQaConfig)
	PutSummarizationConfig(value *ContactCenterInsightsAnalysisRuleAnnotatorSelectorSummarizationConfig)
	ResetIssueModels()
	ResetPhraseMatchers()
	ResetQaConfig()
	ResetRunEntityAnnotator()
	ResetRunIntentAnnotator()
	ResetRunInterruptionAnnotator()
	ResetRunIssueModelAnnotator()
	ResetRunPhraseMatcherAnnotator()
	ResetRunQaAnnotator()
	ResetRunSentimentAnnotator()
	ResetRunSilenceAnnotator()
	ResetRunSummarizationAnnotator()
	ResetSummarizationConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference
type jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) InternalValue() *ContactCenterInsightsAnalysisRuleAnnotatorSelector {
	var returns *ContactCenterInsightsAnalysisRuleAnnotatorSelector
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) IssueModels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"issueModels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) IssueModelsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"issueModelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) PhraseMatchers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"phraseMatchers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) PhraseMatchersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"phraseMatchersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) QaConfig() ContactCenterInsightsAnalysisRuleAnnotatorSelectorQaConfigOutputReference {
	var returns ContactCenterInsightsAnalysisRuleAnnotatorSelectorQaConfigOutputReference
	_jsii_.Get(
		j,
		"qaConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) QaConfigInput() *ContactCenterInsightsAnalysisRuleAnnotatorSelectorQaConfig {
	var returns *ContactCenterInsightsAnalysisRuleAnnotatorSelectorQaConfig
	_jsii_.Get(
		j,
		"qaConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) RunEntityAnnotator() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runEntityAnnotator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) RunEntityAnnotatorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runEntityAnnotatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) RunIntentAnnotator() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runIntentAnnotator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) RunIntentAnnotatorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runIntentAnnotatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) RunInterruptionAnnotator() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runInterruptionAnnotator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) RunInterruptionAnnotatorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runInterruptionAnnotatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) RunIssueModelAnnotator() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runIssueModelAnnotator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) RunIssueModelAnnotatorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runIssueModelAnnotatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) RunPhraseMatcherAnnotator() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runPhraseMatcherAnnotator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) RunPhraseMatcherAnnotatorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runPhraseMatcherAnnotatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) RunQaAnnotator() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runQaAnnotator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) RunQaAnnotatorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runQaAnnotatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) RunSentimentAnnotator() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runSentimentAnnotator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) RunSentimentAnnotatorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runSentimentAnnotatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) RunSilenceAnnotator() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runSilenceAnnotator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) RunSilenceAnnotatorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runSilenceAnnotatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) RunSummarizationAnnotator() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runSummarizationAnnotator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) RunSummarizationAnnotatorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runSummarizationAnnotatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) SummarizationConfig() ContactCenterInsightsAnalysisRuleAnnotatorSelectorSummarizationConfigOutputReference {
	var returns ContactCenterInsightsAnalysisRuleAnnotatorSelectorSummarizationConfigOutputReference
	_jsii_.Get(
		j,
		"summarizationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) SummarizationConfigInput() *ContactCenterInsightsAnalysisRuleAnnotatorSelectorSummarizationConfig {
	var returns *ContactCenterInsightsAnalysisRuleAnnotatorSelectorSummarizationConfig
	_jsii_.Get(
		j,
		"summarizationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference {
	_init_.Initialize()

	if err := validateNewContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.contactCenterInsightsAnalysisRule.ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference_Override(c ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.contactCenterInsightsAnalysisRule.ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference)SetInternalValue(val *ContactCenterInsightsAnalysisRuleAnnotatorSelector) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference)SetIssueModels(val *[]*string) {
	if err := j.validateSetIssueModelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issueModels",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference)SetPhraseMatchers(val *[]*string) {
	if err := j.validateSetPhraseMatchersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"phraseMatchers",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference)SetRunEntityAnnotator(val interface{}) {
	if err := j.validateSetRunEntityAnnotatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runEntityAnnotator",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference)SetRunIntentAnnotator(val interface{}) {
	if err := j.validateSetRunIntentAnnotatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runIntentAnnotator",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference)SetRunInterruptionAnnotator(val interface{}) {
	if err := j.validateSetRunInterruptionAnnotatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runInterruptionAnnotator",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference)SetRunIssueModelAnnotator(val interface{}) {
	if err := j.validateSetRunIssueModelAnnotatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runIssueModelAnnotator",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference)SetRunPhraseMatcherAnnotator(val interface{}) {
	if err := j.validateSetRunPhraseMatcherAnnotatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runPhraseMatcherAnnotator",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference)SetRunQaAnnotator(val interface{}) {
	if err := j.validateSetRunQaAnnotatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runQaAnnotator",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference)SetRunSentimentAnnotator(val interface{}) {
	if err := j.validateSetRunSentimentAnnotatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runSentimentAnnotator",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference)SetRunSilenceAnnotator(val interface{}) {
	if err := j.validateSetRunSilenceAnnotatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runSilenceAnnotator",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference)SetRunSummarizationAnnotator(val interface{}) {
	if err := j.validateSetRunSummarizationAnnotatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runSummarizationAnnotator",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) PutQaConfig(value *ContactCenterInsightsAnalysisRuleAnnotatorSelectorQaConfig) {
	if err := c.validatePutQaConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putQaConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) PutSummarizationConfig(value *ContactCenterInsightsAnalysisRuleAnnotatorSelectorSummarizationConfig) {
	if err := c.validatePutSummarizationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSummarizationConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) ResetIssueModels() {
	_jsii_.InvokeVoid(
		c,
		"resetIssueModels",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) ResetPhraseMatchers() {
	_jsii_.InvokeVoid(
		c,
		"resetPhraseMatchers",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) ResetQaConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetQaConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) ResetRunEntityAnnotator() {
	_jsii_.InvokeVoid(
		c,
		"resetRunEntityAnnotator",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) ResetRunIntentAnnotator() {
	_jsii_.InvokeVoid(
		c,
		"resetRunIntentAnnotator",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) ResetRunInterruptionAnnotator() {
	_jsii_.InvokeVoid(
		c,
		"resetRunInterruptionAnnotator",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) ResetRunIssueModelAnnotator() {
	_jsii_.InvokeVoid(
		c,
		"resetRunIssueModelAnnotator",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) ResetRunPhraseMatcherAnnotator() {
	_jsii_.InvokeVoid(
		c,
		"resetRunPhraseMatcherAnnotator",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) ResetRunQaAnnotator() {
	_jsii_.InvokeVoid(
		c,
		"resetRunQaAnnotator",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) ResetRunSentimentAnnotator() {
	_jsii_.InvokeVoid(
		c,
		"resetRunSentimentAnnotator",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) ResetRunSilenceAnnotator() {
	_jsii_.InvokeVoid(
		c,
		"resetRunSilenceAnnotator",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) ResetRunSummarizationAnnotator() {
	_jsii_.InvokeVoid(
		c,
		"resetRunSummarizationAnnotator",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) ResetSummarizationConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetSummarizationConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContactCenterInsightsAnalysisRuleAnnotatorSelectorOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

