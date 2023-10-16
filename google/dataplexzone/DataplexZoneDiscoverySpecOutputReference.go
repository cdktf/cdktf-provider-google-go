// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataplexzone

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v11/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v11/dataplexzone/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataplexZoneDiscoverySpecOutputReference interface {
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
	CsvOptions() DataplexZoneDiscoverySpecCsvOptionsOutputReference
	CsvOptionsInput() *DataplexZoneDiscoverySpecCsvOptions
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	ExcludePatterns() *[]*string
	SetExcludePatterns(val *[]*string)
	ExcludePatternsInput() *[]*string
	// Experimental.
	Fqn() *string
	IncludePatterns() *[]*string
	SetIncludePatterns(val *[]*string)
	IncludePatternsInput() *[]*string
	InternalValue() *DataplexZoneDiscoverySpec
	SetInternalValue(val *DataplexZoneDiscoverySpec)
	JsonOptions() DataplexZoneDiscoverySpecJsonOptionsOutputReference
	JsonOptionsInput() *DataplexZoneDiscoverySpecJsonOptions
	Schedule() *string
	SetSchedule(val *string)
	ScheduleInput() *string
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
	PutCsvOptions(value *DataplexZoneDiscoverySpecCsvOptions)
	PutJsonOptions(value *DataplexZoneDiscoverySpecJsonOptions)
	ResetCsvOptions()
	ResetExcludePatterns()
	ResetIncludePatterns()
	ResetJsonOptions()
	ResetSchedule()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataplexZoneDiscoverySpecOutputReference
type jsiiProxy_DataplexZoneDiscoverySpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataplexZoneDiscoverySpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexZoneDiscoverySpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexZoneDiscoverySpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexZoneDiscoverySpecOutputReference) CsvOptions() DataplexZoneDiscoverySpecCsvOptionsOutputReference {
	var returns DataplexZoneDiscoverySpecCsvOptionsOutputReference
	_jsii_.Get(
		j,
		"csvOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexZoneDiscoverySpecOutputReference) CsvOptionsInput() *DataplexZoneDiscoverySpecCsvOptions {
	var returns *DataplexZoneDiscoverySpecCsvOptions
	_jsii_.Get(
		j,
		"csvOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexZoneDiscoverySpecOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexZoneDiscoverySpecOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexZoneDiscoverySpecOutputReference) ExcludePatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludePatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexZoneDiscoverySpecOutputReference) ExcludePatternsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludePatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexZoneDiscoverySpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexZoneDiscoverySpecOutputReference) IncludePatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includePatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexZoneDiscoverySpecOutputReference) IncludePatternsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includePatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexZoneDiscoverySpecOutputReference) InternalValue() *DataplexZoneDiscoverySpec {
	var returns *DataplexZoneDiscoverySpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexZoneDiscoverySpecOutputReference) JsonOptions() DataplexZoneDiscoverySpecJsonOptionsOutputReference {
	var returns DataplexZoneDiscoverySpecJsonOptionsOutputReference
	_jsii_.Get(
		j,
		"jsonOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexZoneDiscoverySpecOutputReference) JsonOptionsInput() *DataplexZoneDiscoverySpecJsonOptions {
	var returns *DataplexZoneDiscoverySpecJsonOptions
	_jsii_.Get(
		j,
		"jsonOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexZoneDiscoverySpecOutputReference) Schedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexZoneDiscoverySpecOutputReference) ScheduleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexZoneDiscoverySpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexZoneDiscoverySpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataplexZoneDiscoverySpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataplexZoneDiscoverySpecOutputReference {
	_init_.Initialize()

	if err := validateNewDataplexZoneDiscoverySpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataplexZoneDiscoverySpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataplexZone.DataplexZoneDiscoverySpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataplexZoneDiscoverySpecOutputReference_Override(d DataplexZoneDiscoverySpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataplexZone.DataplexZoneDiscoverySpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataplexZoneDiscoverySpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataplexZoneDiscoverySpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataplexZoneDiscoverySpecOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_DataplexZoneDiscoverySpecOutputReference)SetExcludePatterns(val *[]*string) {
	if err := j.validateSetExcludePatternsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludePatterns",
		val,
	)
}

func (j *jsiiProxy_DataplexZoneDiscoverySpecOutputReference)SetIncludePatterns(val *[]*string) {
	if err := j.validateSetIncludePatternsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includePatterns",
		val,
	)
}

func (j *jsiiProxy_DataplexZoneDiscoverySpecOutputReference)SetInternalValue(val *DataplexZoneDiscoverySpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataplexZoneDiscoverySpecOutputReference)SetSchedule(val *string) {
	if err := j.validateSetScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schedule",
		val,
	)
}

func (j *jsiiProxy_DataplexZoneDiscoverySpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataplexZoneDiscoverySpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataplexZoneDiscoverySpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataplexZoneDiscoverySpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataplexZoneDiscoverySpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataplexZoneDiscoverySpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataplexZoneDiscoverySpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataplexZoneDiscoverySpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataplexZoneDiscoverySpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataplexZoneDiscoverySpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataplexZoneDiscoverySpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataplexZoneDiscoverySpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataplexZoneDiscoverySpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataplexZoneDiscoverySpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataplexZoneDiscoverySpecOutputReference) PutCsvOptions(value *DataplexZoneDiscoverySpecCsvOptions) {
	if err := d.validatePutCsvOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCsvOptions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataplexZoneDiscoverySpecOutputReference) PutJsonOptions(value *DataplexZoneDiscoverySpecJsonOptions) {
	if err := d.validatePutJsonOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putJsonOptions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataplexZoneDiscoverySpecOutputReference) ResetCsvOptions() {
	_jsii_.InvokeVoid(
		d,
		"resetCsvOptions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexZoneDiscoverySpecOutputReference) ResetExcludePatterns() {
	_jsii_.InvokeVoid(
		d,
		"resetExcludePatterns",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexZoneDiscoverySpecOutputReference) ResetIncludePatterns() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludePatterns",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexZoneDiscoverySpecOutputReference) ResetJsonOptions() {
	_jsii_.InvokeVoid(
		d,
		"resetJsonOptions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexZoneDiscoverySpecOutputReference) ResetSchedule() {
	_jsii_.InvokeVoid(
		d,
		"resetSchedule",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexZoneDiscoverySpecOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataplexZoneDiscoverySpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

