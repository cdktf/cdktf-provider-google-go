// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sccmanagementprojectsecurityhealthanalyticscustommodule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/sccmanagementprojectsecurityhealthanalyticscustommodule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference interface {
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
	CustomOutput() SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference
	CustomOutputInput() *SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutput
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfig
	SetInternalValue(val *SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfig)
	Predicate() SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigPredicateOutputReference
	PredicateInput() *SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigPredicate
	Recommendation() *string
	SetRecommendation(val *string)
	RecommendationInput() *string
	ResourceSelector() SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigResourceSelectorOutputReference
	ResourceSelectorInput() *SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigResourceSelector
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
	PutCustomOutput(value *SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutput)
	PutPredicate(value *SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigPredicate)
	PutResourceSelector(value *SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigResourceSelector)
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

// The jsii proxy struct for SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference
type jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) CustomOutput() SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference {
	var returns SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutputOutputReference
	_jsii_.Get(
		j,
		"customOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) CustomOutputInput() *SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutput {
	var returns *SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutput
	_jsii_.Get(
		j,
		"customOutputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) InternalValue() *SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfig {
	var returns *SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) Predicate() SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigPredicateOutputReference {
	var returns SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigPredicateOutputReference
	_jsii_.Get(
		j,
		"predicate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) PredicateInput() *SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigPredicate {
	var returns *SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigPredicate
	_jsii_.Get(
		j,
		"predicateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) Recommendation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recommendation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) RecommendationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recommendationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) ResourceSelector() SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigResourceSelectorOutputReference {
	var returns SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigResourceSelectorOutputReference
	_jsii_.Get(
		j,
		"resourceSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) ResourceSelectorInput() *SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigResourceSelector {
	var returns *SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigResourceSelector
	_jsii_.Get(
		j,
		"resourceSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) Severity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"severity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) SeverityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"severityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference {
	_init_.Initialize()

	if err := validateNewSccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.sccManagementProjectSecurityHealthAnalyticsCustomModule.SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference_Override(s SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.sccManagementProjectSecurityHealthAnalyticsCustomModule.SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference)SetInternalValue(val *SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference)SetRecommendation(val *string) {
	if err := j.validateSetRecommendationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recommendation",
		val,
	)
}

func (j *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference)SetSeverity(val *string) {
	if err := j.validateSetSeverityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"severity",
		val,
	)
}

func (j *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) PutCustomOutput(value *SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigCustomOutput) {
	if err := s.validatePutCustomOutputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCustomOutput",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) PutPredicate(value *SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigPredicate) {
	if err := s.validatePutPredicateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPredicate",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) PutResourceSelector(value *SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigResourceSelector) {
	if err := s.validatePutResourceSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putResourceSelector",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) ResetCustomOutput() {
	_jsii_.InvokeVoid(
		s,
		"resetCustomOutput",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		s,
		"resetDescription",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

