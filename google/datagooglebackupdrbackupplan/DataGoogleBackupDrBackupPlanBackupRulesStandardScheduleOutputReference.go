// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagooglebackupdrbackupplan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/datagooglebackupdrbackupplan/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference interface {
	cdktf.ComplexObject
	BackupWindow() DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleBackupWindowList
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
	DaysOfMonth() *[]*float64
	DaysOfWeek() *[]*string
	// Experimental.
	Fqn() *string
	HourlyFrequency() *float64
	InternalValue() *DataGoogleBackupDrBackupPlanBackupRulesStandardSchedule
	SetInternalValue(val *DataGoogleBackupDrBackupPlanBackupRulesStandardSchedule)
	Months() *[]*string
	RecurrenceType() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeZone() *string
	WeekDayOfMonth() DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleWeekDayOfMonthList
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference
type jsiiProxy_DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) BackupWindow() DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleBackupWindowList {
	var returns DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleBackupWindowList
	_jsii_.Get(
		j,
		"backupWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) DaysOfMonth() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"daysOfMonth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) DaysOfWeek() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"daysOfWeek",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) HourlyFrequency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hourlyFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) InternalValue() *DataGoogleBackupDrBackupPlanBackupRulesStandardSchedule {
	var returns *DataGoogleBackupDrBackupPlanBackupRulesStandardSchedule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) Months() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"months",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) RecurrenceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recurrenceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) WeekDayOfMonth() DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleWeekDayOfMonthList {
	var returns DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleWeekDayOfMonthList
	_jsii_.Get(
		j,
		"weekDayOfMonth",
		&returns,
	)
	return returns
}


func NewDataGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference {
	_init_.Initialize()

	if err := validateNewDataGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleBackupDrBackupPlan.DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference_Override(d DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleBackupDrBackupPlan.DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference)SetInternalValue(val *DataGoogleBackupDrBackupPlanBackupRulesStandardSchedule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataGoogleBackupDrBackupPlanBackupRulesStandardScheduleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

