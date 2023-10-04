// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventionjobtrigger

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v10/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v10/datalosspreventionjobtrigger/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference interface {
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
	ExcludedFields() DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFieldsList
	ExcludedFieldsInput() interface{}
	// Experimental.
	Fqn() *string
	IdentifyingFields() DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFieldsList
	IdentifyingFieldsInput() interface{}
	IncludedFields() DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFieldsList
	IncludedFieldsInput() interface{}
	InternalValue() *DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptions
	SetInternalValue(val *DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptions)
	RowsLimit() *float64
	SetRowsLimit(val *float64)
	RowsLimitInput() *float64
	RowsLimitPercent() *float64
	SetRowsLimitPercent(val *float64)
	RowsLimitPercentInput() *float64
	SampleMethod() *string
	SetSampleMethod(val *string)
	SampleMethodInput() *string
	TableReference() DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsTableReferenceOutputReference
	TableReferenceInput() *DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference
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
	PutExcludedFields(value interface{})
	PutIdentifyingFields(value interface{})
	PutIncludedFields(value interface{})
	PutTableReference(value *DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference)
	ResetExcludedFields()
	ResetIdentifyingFields()
	ResetIncludedFields()
	ResetRowsLimit()
	ResetRowsLimitPercent()
	ResetSampleMethod()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference
type jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) ExcludedFields() DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFieldsList {
	var returns DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsExcludedFieldsList
	_jsii_.Get(
		j,
		"excludedFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) ExcludedFieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludedFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) IdentifyingFields() DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFieldsList {
	var returns DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsIdentifyingFieldsList
	_jsii_.Get(
		j,
		"identifyingFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) IdentifyingFieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"identifyingFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) IncludedFields() DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFieldsList {
	var returns DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsIncludedFieldsList
	_jsii_.Get(
		j,
		"includedFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) IncludedFieldsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includedFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) InternalValue() *DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptions {
	var returns *DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) RowsLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rowsLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) RowsLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rowsLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) RowsLimitPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rowsLimitPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) RowsLimitPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rowsLimitPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) SampleMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sampleMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) SampleMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sampleMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) TableReference() DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsTableReferenceOutputReference {
	var returns DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsTableReferenceOutputReference
	_jsii_.Get(
		j,
		"tableReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) TableReferenceInput() *DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference {
	var returns *DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference
	_jsii_.Get(
		j,
		"tableReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataLossPreventionJobTrigger.DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference_Override(d DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataLossPreventionJobTrigger.DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference)SetInternalValue(val *DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference)SetRowsLimit(val *float64) {
	if err := j.validateSetRowsLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rowsLimit",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference)SetRowsLimitPercent(val *float64) {
	if err := j.validateSetRowsLimitPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rowsLimitPercent",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference)SetSampleMethod(val *string) {
	if err := j.validateSetSampleMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sampleMethod",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) PutExcludedFields(value interface{}) {
	if err := d.validatePutExcludedFieldsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putExcludedFields",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) PutIdentifyingFields(value interface{}) {
	if err := d.validatePutIdentifyingFieldsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putIdentifyingFields",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) PutIncludedFields(value interface{}) {
	if err := d.validatePutIncludedFieldsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putIncludedFields",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) PutTableReference(value *DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsTableReference) {
	if err := d.validatePutTableReferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTableReference",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) ResetExcludedFields() {
	_jsii_.InvokeVoid(
		d,
		"resetExcludedFields",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) ResetIdentifyingFields() {
	_jsii_.InvokeVoid(
		d,
		"resetIdentifyingFields",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) ResetIncludedFields() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludedFields",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) ResetRowsLimit() {
	_jsii_.InvokeVoid(
		d,
		"resetRowsLimit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) ResetRowsLimitPercent() {
	_jsii_.InvokeVoid(
		d,
		"resetRowsLimitPercent",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) ResetSampleMethod() {
	_jsii_.InvokeVoid(
		d,
		"resetSampleMethod",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

