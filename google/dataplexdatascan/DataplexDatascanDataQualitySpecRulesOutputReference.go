// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataplexdatascan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/dataplexdatascan/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataplexDatascanDataQualitySpecRulesOutputReference interface {
	cdktf.ComplexObject
	Column() *string
	SetColumn(val *string)
	ColumnInput() *string
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Dimension() *string
	SetDimension(val *string)
	DimensionInput() *string
	// Experimental.
	Fqn() *string
	IgnoreNull() interface{}
	SetIgnoreNull(val interface{})
	IgnoreNullInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	NonNullExpectation() DataplexDatascanDataQualitySpecRulesNonNullExpectationOutputReference
	NonNullExpectationInput() *DataplexDatascanDataQualitySpecRulesNonNullExpectation
	RangeExpectation() DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference
	RangeExpectationInput() *DataplexDatascanDataQualitySpecRulesRangeExpectation
	RegexExpectation() DataplexDatascanDataQualitySpecRulesRegexExpectationOutputReference
	RegexExpectationInput() *DataplexDatascanDataQualitySpecRulesRegexExpectation
	RowConditionExpectation() DataplexDatascanDataQualitySpecRulesRowConditionExpectationOutputReference
	RowConditionExpectationInput() *DataplexDatascanDataQualitySpecRulesRowConditionExpectation
	SetExpectation() DataplexDatascanDataQualitySpecRulesSetExpectationOutputReference
	SetExpectationInput() *DataplexDatascanDataQualitySpecRulesSetExpectation
	SqlAssertion() DataplexDatascanDataQualitySpecRulesSqlAssertionOutputReference
	SqlAssertionInput() *DataplexDatascanDataQualitySpecRulesSqlAssertion
	StatisticRangeExpectation() DataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference
	StatisticRangeExpectationInput() *DataplexDatascanDataQualitySpecRulesStatisticRangeExpectation
	TableConditionExpectation() DataplexDatascanDataQualitySpecRulesTableConditionExpectationOutputReference
	TableConditionExpectationInput() *DataplexDatascanDataQualitySpecRulesTableConditionExpectation
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Threshold() *float64
	SetThreshold(val *float64)
	ThresholdInput() *float64
	UniquenessExpectation() DataplexDatascanDataQualitySpecRulesUniquenessExpectationOutputReference
	UniquenessExpectationInput() *DataplexDatascanDataQualitySpecRulesUniquenessExpectation
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
	PutNonNullExpectation(value *DataplexDatascanDataQualitySpecRulesNonNullExpectation)
	PutRangeExpectation(value *DataplexDatascanDataQualitySpecRulesRangeExpectation)
	PutRegexExpectation(value *DataplexDatascanDataQualitySpecRulesRegexExpectation)
	PutRowConditionExpectation(value *DataplexDatascanDataQualitySpecRulesRowConditionExpectation)
	PutSetExpectation(value *DataplexDatascanDataQualitySpecRulesSetExpectation)
	PutSqlAssertion(value *DataplexDatascanDataQualitySpecRulesSqlAssertion)
	PutStatisticRangeExpectation(value *DataplexDatascanDataQualitySpecRulesStatisticRangeExpectation)
	PutTableConditionExpectation(value *DataplexDatascanDataQualitySpecRulesTableConditionExpectation)
	PutUniquenessExpectation(value *DataplexDatascanDataQualitySpecRulesUniquenessExpectation)
	ResetColumn()
	ResetDescription()
	ResetIgnoreNull()
	ResetName()
	ResetNonNullExpectation()
	ResetRangeExpectation()
	ResetRegexExpectation()
	ResetRowConditionExpectation()
	ResetSetExpectation()
	ResetSqlAssertion()
	ResetStatisticRangeExpectation()
	ResetTableConditionExpectation()
	ResetThreshold()
	ResetUniquenessExpectation()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataplexDatascanDataQualitySpecRulesOutputReference
type jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) Column() *string {
	var returns *string
	_jsii_.Get(
		j,
		"column",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) ColumnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"columnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) Dimension() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dimension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) DimensionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dimensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) IgnoreNull() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreNull",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) IgnoreNullInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreNullInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) NonNullExpectation() DataplexDatascanDataQualitySpecRulesNonNullExpectationOutputReference {
	var returns DataplexDatascanDataQualitySpecRulesNonNullExpectationOutputReference
	_jsii_.Get(
		j,
		"nonNullExpectation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) NonNullExpectationInput() *DataplexDatascanDataQualitySpecRulesNonNullExpectation {
	var returns *DataplexDatascanDataQualitySpecRulesNonNullExpectation
	_jsii_.Get(
		j,
		"nonNullExpectationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) RangeExpectation() DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference {
	var returns DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference
	_jsii_.Get(
		j,
		"rangeExpectation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) RangeExpectationInput() *DataplexDatascanDataQualitySpecRulesRangeExpectation {
	var returns *DataplexDatascanDataQualitySpecRulesRangeExpectation
	_jsii_.Get(
		j,
		"rangeExpectationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) RegexExpectation() DataplexDatascanDataQualitySpecRulesRegexExpectationOutputReference {
	var returns DataplexDatascanDataQualitySpecRulesRegexExpectationOutputReference
	_jsii_.Get(
		j,
		"regexExpectation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) RegexExpectationInput() *DataplexDatascanDataQualitySpecRulesRegexExpectation {
	var returns *DataplexDatascanDataQualitySpecRulesRegexExpectation
	_jsii_.Get(
		j,
		"regexExpectationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) RowConditionExpectation() DataplexDatascanDataQualitySpecRulesRowConditionExpectationOutputReference {
	var returns DataplexDatascanDataQualitySpecRulesRowConditionExpectationOutputReference
	_jsii_.Get(
		j,
		"rowConditionExpectation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) RowConditionExpectationInput() *DataplexDatascanDataQualitySpecRulesRowConditionExpectation {
	var returns *DataplexDatascanDataQualitySpecRulesRowConditionExpectation
	_jsii_.Get(
		j,
		"rowConditionExpectationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) SetExpectation() DataplexDatascanDataQualitySpecRulesSetExpectationOutputReference {
	var returns DataplexDatascanDataQualitySpecRulesSetExpectationOutputReference
	_jsii_.Get(
		j,
		"setExpectation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) SetExpectationInput() *DataplexDatascanDataQualitySpecRulesSetExpectation {
	var returns *DataplexDatascanDataQualitySpecRulesSetExpectation
	_jsii_.Get(
		j,
		"setExpectationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) SqlAssertion() DataplexDatascanDataQualitySpecRulesSqlAssertionOutputReference {
	var returns DataplexDatascanDataQualitySpecRulesSqlAssertionOutputReference
	_jsii_.Get(
		j,
		"sqlAssertion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) SqlAssertionInput() *DataplexDatascanDataQualitySpecRulesSqlAssertion {
	var returns *DataplexDatascanDataQualitySpecRulesSqlAssertion
	_jsii_.Get(
		j,
		"sqlAssertionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) StatisticRangeExpectation() DataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference {
	var returns DataplexDatascanDataQualitySpecRulesStatisticRangeExpectationOutputReference
	_jsii_.Get(
		j,
		"statisticRangeExpectation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) StatisticRangeExpectationInput() *DataplexDatascanDataQualitySpecRulesStatisticRangeExpectation {
	var returns *DataplexDatascanDataQualitySpecRulesStatisticRangeExpectation
	_jsii_.Get(
		j,
		"statisticRangeExpectationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) TableConditionExpectation() DataplexDatascanDataQualitySpecRulesTableConditionExpectationOutputReference {
	var returns DataplexDatascanDataQualitySpecRulesTableConditionExpectationOutputReference
	_jsii_.Get(
		j,
		"tableConditionExpectation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) TableConditionExpectationInput() *DataplexDatascanDataQualitySpecRulesTableConditionExpectation {
	var returns *DataplexDatascanDataQualitySpecRulesTableConditionExpectation
	_jsii_.Get(
		j,
		"tableConditionExpectationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) Threshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) ThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"thresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) UniquenessExpectation() DataplexDatascanDataQualitySpecRulesUniquenessExpectationOutputReference {
	var returns DataplexDatascanDataQualitySpecRulesUniquenessExpectationOutputReference
	_jsii_.Get(
		j,
		"uniquenessExpectation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) UniquenessExpectationInput() *DataplexDatascanDataQualitySpecRulesUniquenessExpectation {
	var returns *DataplexDatascanDataQualitySpecRulesUniquenessExpectation
	_jsii_.Get(
		j,
		"uniquenessExpectationInput",
		&returns,
	)
	return returns
}


func NewDataplexDatascanDataQualitySpecRulesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataplexDatascanDataQualitySpecRulesOutputReference {
	_init_.Initialize()

	if err := validateNewDataplexDatascanDataQualitySpecRulesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataplexDatascan.DataplexDatascanDataQualitySpecRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataplexDatascanDataQualitySpecRulesOutputReference_Override(d DataplexDatascanDataQualitySpecRulesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataplexDatascan.DataplexDatascanDataQualitySpecRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference)SetColumn(val *string) {
	if err := j.validateSetColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"column",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference)SetDimension(val *string) {
	if err := j.validateSetDimensionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dimension",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference)SetIgnoreNull(val interface{}) {
	if err := j.validateSetIgnoreNullParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreNull",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference)SetThreshold(val *float64) {
	if err := j.validateSetThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threshold",
		val,
	)
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) PutNonNullExpectation(value *DataplexDatascanDataQualitySpecRulesNonNullExpectation) {
	if err := d.validatePutNonNullExpectationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNonNullExpectation",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) PutRangeExpectation(value *DataplexDatascanDataQualitySpecRulesRangeExpectation) {
	if err := d.validatePutRangeExpectationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRangeExpectation",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) PutRegexExpectation(value *DataplexDatascanDataQualitySpecRulesRegexExpectation) {
	if err := d.validatePutRegexExpectationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRegexExpectation",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) PutRowConditionExpectation(value *DataplexDatascanDataQualitySpecRulesRowConditionExpectation) {
	if err := d.validatePutRowConditionExpectationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRowConditionExpectation",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) PutSetExpectation(value *DataplexDatascanDataQualitySpecRulesSetExpectation) {
	if err := d.validatePutSetExpectationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSetExpectation",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) PutSqlAssertion(value *DataplexDatascanDataQualitySpecRulesSqlAssertion) {
	if err := d.validatePutSqlAssertionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSqlAssertion",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) PutStatisticRangeExpectation(value *DataplexDatascanDataQualitySpecRulesStatisticRangeExpectation) {
	if err := d.validatePutStatisticRangeExpectationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putStatisticRangeExpectation",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) PutTableConditionExpectation(value *DataplexDatascanDataQualitySpecRulesTableConditionExpectation) {
	if err := d.validatePutTableConditionExpectationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTableConditionExpectation",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) PutUniquenessExpectation(value *DataplexDatascanDataQualitySpecRulesUniquenessExpectation) {
	if err := d.validatePutUniquenessExpectationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putUniquenessExpectation",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) ResetColumn() {
	_jsii_.InvokeVoid(
		d,
		"resetColumn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) ResetIgnoreNull() {
	_jsii_.InvokeVoid(
		d,
		"resetIgnoreNull",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		d,
		"resetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) ResetNonNullExpectation() {
	_jsii_.InvokeVoid(
		d,
		"resetNonNullExpectation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) ResetRangeExpectation() {
	_jsii_.InvokeVoid(
		d,
		"resetRangeExpectation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) ResetRegexExpectation() {
	_jsii_.InvokeVoid(
		d,
		"resetRegexExpectation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) ResetRowConditionExpectation() {
	_jsii_.InvokeVoid(
		d,
		"resetRowConditionExpectation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) ResetSetExpectation() {
	_jsii_.InvokeVoid(
		d,
		"resetSetExpectation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) ResetSqlAssertion() {
	_jsii_.InvokeVoid(
		d,
		"resetSqlAssertion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) ResetStatisticRangeExpectation() {
	_jsii_.InvokeVoid(
		d,
		"resetStatisticRangeExpectation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) ResetTableConditionExpectation() {
	_jsii_.InvokeVoid(
		d,
		"resetTableConditionExpectation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) ResetThreshold() {
	_jsii_.InvokeVoid(
		d,
		"resetThreshold",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) ResetUniquenessExpectation() {
	_jsii_.InvokeVoid(
		d,
		"resetUniquenessExpectation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

