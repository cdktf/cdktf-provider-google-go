// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sccmanagementorganizationsecurityhealthanalyticscustommodule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/sccmanagementorganizationsecurityhealthanalyticscustommodule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference interface {
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
	CustomOutput() SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference
	CustomOutputInput() *SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutput
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfig
	SetInternalValue(val *SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfig)
	Predicate() SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigPredicateOutputReference
	PredicateInput() *SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigPredicate
	Recommendation() *string
	SetRecommendation(val *string)
	RecommendationInput() *string
	ResourceSelector() SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigResourceSelectorOutputReference
	ResourceSelectorInput() *SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigResourceSelector
	Severity() *string
	SetSeverity(val *string)
	SeverityInput() *string
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
	PutCustomOutput(value *SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutput)
	PutPredicate(value *SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigPredicate)
	PutResourceSelector(value *SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigResourceSelector)
	ResetCustomOutput()
	ResetDescription()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference
type jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) CustomOutput() SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference {
	var returns SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference
	_jsii_.Get(
		j,
		"customOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) CustomOutputInput() *SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutput {
	var returns *SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutput
	_jsii_.Get(
		j,
		"customOutputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) InternalValue() *SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfig {
	var returns *SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) Predicate() SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigPredicateOutputReference {
	var returns SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigPredicateOutputReference
	_jsii_.Get(
		j,
		"predicate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) PredicateInput() *SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigPredicate {
	var returns *SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigPredicate
	_jsii_.Get(
		j,
		"predicateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) Recommendation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recommendation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) RecommendationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recommendationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) ResourceSelector() SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigResourceSelectorOutputReference {
	var returns SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigResourceSelectorOutputReference
	_jsii_.Get(
		j,
		"resourceSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) ResourceSelectorInput() *SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigResourceSelector {
	var returns *SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigResourceSelector
	_jsii_.Get(
		j,
		"resourceSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) Severity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"severity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) SeverityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"severityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference {
	_init_.Initialize()

	if err := validateNewSccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.sccManagementOrganizationSecurityHealthAnalyticsCustomModule.SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference_Override(s SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.sccManagementOrganizationSecurityHealthAnalyticsCustomModule.SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference)SetInternalValue(val *SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference)SetRecommendation(val *string) {
	if err := j.validateSetRecommendationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recommendation",
		val,
	)
}

func (j *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference)SetSeverity(val *string) {
	if err := j.validateSetSeverityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"severity",
		val,
	)
}

func (j *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) PutCustomOutput(value *SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutput) {
	if err := s.validatePutCustomOutputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCustomOutput",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) PutPredicate(value *SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigPredicate) {
	if err := s.validatePutPredicateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPredicate",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) PutResourceSelector(value *SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigResourceSelector) {
	if err := s.validatePutResourceSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putResourceSelector",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) ResetCustomOutput() {
	_jsii_.InvokeVoid(
		s,
		"resetCustomOutput",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		s,
		"resetDescription",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

