// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagoogledataplexdataqualityrules

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/datagoogledataplexdataqualityrules/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleDataplexDataQualityRulesRulesOutputReference interface {
	cdktf.ComplexObject
	Column() *string
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
	Dimension() *string
	// Experimental.
	Fqn() *string
	IgnoreNull() cdktf.IResolvable
	InternalValue() *DataGoogleDataplexDataQualityRulesRules
	SetInternalValue(val *DataGoogleDataplexDataQualityRulesRules)
	Name() *string
	NonNullExpectation() DataGoogleDataplexDataQualityRulesRulesNonNullExpectationList
	RangeExpectation() DataGoogleDataplexDataQualityRulesRulesRangeExpectationList
	RegexExpectation() DataGoogleDataplexDataQualityRulesRulesRegexExpectationList
	RowConditionExpectation() DataGoogleDataplexDataQualityRulesRulesRowConditionExpectationList
	SetExpectation() DataGoogleDataplexDataQualityRulesRulesSetExpectationList
	SqlAssertion() DataGoogleDataplexDataQualityRulesRulesSqlAssertionList
	StatisticRangeExpectation() DataGoogleDataplexDataQualityRulesRulesStatisticRangeExpectationList
	Suspended() cdktf.IResolvable
	TableConditionExpectation() DataGoogleDataplexDataQualityRulesRulesTableConditionExpectationList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Threshold() *float64
	UniquenessExpectation() DataGoogleDataplexDataQualityRulesRulesUniquenessExpectationList
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataGoogleDataplexDataQualityRulesRulesOutputReference
type jsiiProxy_DataGoogleDataplexDataQualityRulesRulesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataGoogleDataplexDataQualityRulesRulesOutputReference) Column() *string {
	var returns *string
	_jsii_.Get(
		j,
		"column",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleDataplexDataQualityRulesRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleDataplexDataQualityRulesRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleDataplexDataQualityRulesRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleDataplexDataQualityRulesRulesOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleDataplexDataQualityRulesRulesOutputReference) Dimension() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dimension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleDataplexDataQualityRulesRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleDataplexDataQualityRulesRulesOutputReference) IgnoreNull() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"ignoreNull",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleDataplexDataQualityRulesRulesOutputReference) InternalValue() *DataGoogleDataplexDataQualityRulesRules {
	var returns *DataGoogleDataplexDataQualityRulesRules
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleDataplexDataQualityRulesRulesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleDataplexDataQualityRulesRulesOutputReference) NonNullExpectation() DataGoogleDataplexDataQualityRulesRulesNonNullExpectationList {
	var returns DataGoogleDataplexDataQualityRulesRulesNonNullExpectationList
	_jsii_.Get(
		j,
		"nonNullExpectation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleDataplexDataQualityRulesRulesOutputReference) RangeExpectation() DataGoogleDataplexDataQualityRulesRulesRangeExpectationList {
	var returns DataGoogleDataplexDataQualityRulesRulesRangeExpectationList
	_jsii_.Get(
		j,
		"rangeExpectation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleDataplexDataQualityRulesRulesOutputReference) RegexExpectation() DataGoogleDataplexDataQualityRulesRulesRegexExpectationList {
	var returns DataGoogleDataplexDataQualityRulesRulesRegexExpectationList
	_jsii_.Get(
		j,
		"regexExpectation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleDataplexDataQualityRulesRulesOutputReference) RowConditionExpectation() DataGoogleDataplexDataQualityRulesRulesRowConditionExpectationList {
	var returns DataGoogleDataplexDataQualityRulesRulesRowConditionExpectationList
	_jsii_.Get(
		j,
		"rowConditionExpectation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleDataplexDataQualityRulesRulesOutputReference) SetExpectation() DataGoogleDataplexDataQualityRulesRulesSetExpectationList {
	var returns DataGoogleDataplexDataQualityRulesRulesSetExpectationList
	_jsii_.Get(
		j,
		"setExpectation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleDataplexDataQualityRulesRulesOutputReference) SqlAssertion() DataGoogleDataplexDataQualityRulesRulesSqlAssertionList {
	var returns DataGoogleDataplexDataQualityRulesRulesSqlAssertionList
	_jsii_.Get(
		j,
		"sqlAssertion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleDataplexDataQualityRulesRulesOutputReference) StatisticRangeExpectation() DataGoogleDataplexDataQualityRulesRulesStatisticRangeExpectationList {
	var returns DataGoogleDataplexDataQualityRulesRulesStatisticRangeExpectationList
	_jsii_.Get(
		j,
		"statisticRangeExpectation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleDataplexDataQualityRulesRulesOutputReference) Suspended() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"suspended",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleDataplexDataQualityRulesRulesOutputReference) TableConditionExpectation() DataGoogleDataplexDataQualityRulesRulesTableConditionExpectationList {
	var returns DataGoogleDataplexDataQualityRulesRulesTableConditionExpectationList
	_jsii_.Get(
		j,
		"tableConditionExpectation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleDataplexDataQualityRulesRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleDataplexDataQualityRulesRulesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleDataplexDataQualityRulesRulesOutputReference) Threshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleDataplexDataQualityRulesRulesOutputReference) UniquenessExpectation() DataGoogleDataplexDataQualityRulesRulesUniquenessExpectationList {
	var returns DataGoogleDataplexDataQualityRulesRulesUniquenessExpectationList
	_jsii_.Get(
		j,
		"uniquenessExpectation",
		&returns,
	)
	return returns
}


func NewDataGoogleDataplexDataQualityRulesRulesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataGoogleDataplexDataQualityRulesRulesOutputReference {
	_init_.Initialize()

	if err := validateNewDataGoogleDataplexDataQualityRulesRulesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGoogleDataplexDataQualityRulesRulesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleDataplexDataQualityRules.DataGoogleDataplexDataQualityRulesRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataGoogleDataplexDataQualityRulesRulesOutputReference_Override(d DataGoogleDataplexDataQualityRulesRulesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleDataplexDataQualityRules.DataGoogleDataplexDataQualityRulesRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataGoogleDataplexDataQualityRulesRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataGoogleDataplexDataQualityRulesRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataGoogleDataplexDataQualityRulesRulesOutputReference)SetInternalValue(val *DataGoogleDataplexDataQualityRulesRules) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataGoogleDataplexDataQualityRulesRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataGoogleDataplexDataQualityRulesRulesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataGoogleDataplexDataQualityRulesRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleDataplexDataQualityRulesRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataGoogleDataplexDataQualityRulesRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGoogleDataplexDataQualityRulesRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataGoogleDataplexDataQualityRulesRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataGoogleDataplexDataQualityRulesRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataGoogleDataplexDataQualityRulesRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataGoogleDataplexDataQualityRulesRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataGoogleDataplexDataQualityRulesRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataGoogleDataplexDataQualityRulesRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataGoogleDataplexDataQualityRulesRulesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleDataplexDataQualityRulesRulesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGoogleDataplexDataQualityRulesRulesOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataGoogleDataplexDataQualityRulesRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

