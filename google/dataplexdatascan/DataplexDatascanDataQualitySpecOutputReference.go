// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataplexdatascan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v12/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v12/dataplexdatascan/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataplexDatascanDataQualitySpecOutputReference interface {
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
	InternalValue() *DataplexDatascanDataQualitySpec
	SetInternalValue(val *DataplexDatascanDataQualitySpec)
	PostScanActions() DataplexDatascanDataQualitySpecPostScanActionsOutputReference
	PostScanActionsInput() *DataplexDatascanDataQualitySpecPostScanActions
	RowFilter() *string
	SetRowFilter(val *string)
	RowFilterInput() *string
	Rules() DataplexDatascanDataQualitySpecRulesList
	RulesInput() interface{}
	SamplingPercent() *float64
	SetSamplingPercent(val *float64)
	SamplingPercentInput() *float64
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
	PutPostScanActions(value *DataplexDatascanDataQualitySpecPostScanActions)
	PutRules(value interface{})
	ResetPostScanActions()
	ResetRowFilter()
	ResetRules()
	ResetSamplingPercent()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataplexDatascanDataQualitySpecOutputReference
type jsiiProxy_DataplexDatascanDataQualitySpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecOutputReference) InternalValue() *DataplexDatascanDataQualitySpec {
	var returns *DataplexDatascanDataQualitySpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecOutputReference) PostScanActions() DataplexDatascanDataQualitySpecPostScanActionsOutputReference {
	var returns DataplexDatascanDataQualitySpecPostScanActionsOutputReference
	_jsii_.Get(
		j,
		"postScanActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecOutputReference) PostScanActionsInput() *DataplexDatascanDataQualitySpecPostScanActions {
	var returns *DataplexDatascanDataQualitySpecPostScanActions
	_jsii_.Get(
		j,
		"postScanActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecOutputReference) RowFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rowFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecOutputReference) RowFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rowFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecOutputReference) Rules() DataplexDatascanDataQualitySpecRulesList {
	var returns DataplexDatascanDataQualitySpecRulesList
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecOutputReference) RulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecOutputReference) SamplingPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"samplingPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecOutputReference) SamplingPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"samplingPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataplexDatascanDataQualitySpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataplexDatascanDataQualitySpecOutputReference {
	_init_.Initialize()

	if err := validateNewDataplexDatascanDataQualitySpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataplexDatascanDataQualitySpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataplexDatascan.DataplexDatascanDataQualitySpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataplexDatascanDataQualitySpecOutputReference_Override(d DataplexDatascanDataQualitySpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataplexDatascan.DataplexDatascanDataQualitySpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecOutputReference)SetInternalValue(val *DataplexDatascanDataQualitySpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecOutputReference)SetRowFilter(val *string) {
	if err := j.validateSetRowFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rowFilter",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecOutputReference)SetSamplingPercent(val *float64) {
	if err := j.validateSetSamplingPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"samplingPercent",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanDataQualitySpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataplexDatascanDataQualitySpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataplexDatascanDataQualitySpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataplexDatascanDataQualitySpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataplexDatascanDataQualitySpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataplexDatascanDataQualitySpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataplexDatascanDataQualitySpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataplexDatascanDataQualitySpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataplexDatascanDataQualitySpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataplexDatascanDataQualitySpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataplexDatascanDataQualitySpecOutputReference) PutPostScanActions(value *DataplexDatascanDataQualitySpecPostScanActions) {
	if err := d.validatePutPostScanActionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPostScanActions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecOutputReference) PutRules(value interface{}) {
	if err := d.validatePutRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRules",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecOutputReference) ResetPostScanActions() {
	_jsii_.InvokeVoid(
		d,
		"resetPostScanActions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecOutputReference) ResetRowFilter() {
	_jsii_.InvokeVoid(
		d,
		"resetRowFilter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecOutputReference) ResetRules() {
	_jsii_.InvokeVoid(
		d,
		"resetRules",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecOutputReference) ResetSamplingPercent() {
	_jsii_.InvokeVoid(
		d,
		"resetSamplingPercent",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexDatascanDataQualitySpecOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataplexDatascanDataQualitySpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

