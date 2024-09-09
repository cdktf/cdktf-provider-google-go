// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataplexdatascan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/dataplexdatascan/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference interface {
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
	InternalValue() *DataplexDatascanDataQualitySpecRulesRangeExpectation
	SetInternalValue(val *DataplexDatascanDataQualitySpecRulesRangeExpectation)
	MaxValue() *string
	SetMaxValue(val *string)
	MaxValueInput() *string
	MinValue() *string
	SetMinValue(val *string)
	MinValueInput() *string
	StrictMaxEnabled() interface{}
	SetStrictMaxEnabled(val interface{})
	StrictMaxEnabledInput() interface{}
	StrictMinEnabled() interface{}
	SetStrictMinEnabled(val interface{})
	StrictMinEnabledInput() interface{}
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
	ResetMaxValue()
	ResetMinValue()
	ResetStrictMaxEnabled()
	ResetStrictMinEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference
type jsiiProxy_DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) InternalValue() *DataplexDatascanDataQualitySpecRulesRangeExpectation {
	var returns *DataplexDatascanDataQualitySpecRulesRangeExpectation
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) MaxValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) MaxValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) MinValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) MinValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) StrictMaxEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strictMaxEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) StrictMaxEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strictMaxEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) StrictMinEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strictMinEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) StrictMinEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strictMinEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference {
	_init_.Initialize()

	if err := validateNewDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataplexDatascan.DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference_Override(d DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataplexDatascan.DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference)SetInternalValue(val *DataplexDatascanDataQualitySpecRulesRangeExpectation) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference)SetMaxValue(val *string) {
	if err := j.validateSetMaxValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxValue",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference)SetMinValue(val *string) {
	if err := j.validateSetMinValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minValue",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference)SetStrictMaxEnabled(val interface{}) {
	if err := j.validateSetStrictMaxEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"strictMaxEnabled",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference)SetStrictMinEnabled(val interface{}) {
	if err := j.validateSetStrictMinEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"strictMinEnabled",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) ResetMaxValue() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxValue",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) ResetMinValue() {
	_jsii_.InvokeVoid(
		d,
		"resetMinValue",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) ResetStrictMaxEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetStrictMaxEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) ResetStrictMinEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetStrictMinEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataplexDatascanDataQualitySpecRulesRangeExpectationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

