// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package backupdrbackupplan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/backupdrbackupplan/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BackupDrBackupPlanBackupRulesStandardScheduleOutputReference interface {
	cdktf.ComplexObject
	BackupWindow() BackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference
	BackupWindowInput() *BackupDrBackupPlanBackupRulesStandardScheduleBackupWindow
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
	SetDaysOfMonth(val *[]*float64)
	DaysOfMonthInput() *[]*float64
	DaysOfWeek() *[]*string
	SetDaysOfWeek(val *[]*string)
	DaysOfWeekInput() *[]*string
	// Experimental.
	Fqn() *string
	HourlyFrequency() *float64
	SetHourlyFrequency(val *float64)
	HourlyFrequencyInput() *float64
	InternalValue() *BackupDrBackupPlanBackupRulesStandardSchedule
	SetInternalValue(val *BackupDrBackupPlanBackupRulesStandardSchedule)
	Months() *[]*string
	SetMonths(val *[]*string)
	MonthsInput() *[]*string
	RecurrenceType() *string
	SetRecurrenceType(val *string)
	RecurrenceTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeZone() *string
	SetTimeZone(val *string)
	TimeZoneInput() *string
	WeekDayOfMonth() BackupDrBackupPlanBackupRulesStandardScheduleWeekDayOfMonthOutputReference
	WeekDayOfMonthInput() *BackupDrBackupPlanBackupRulesStandardScheduleWeekDayOfMonth
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
	PutBackupWindow(value *BackupDrBackupPlanBackupRulesStandardScheduleBackupWindow)
	PutWeekDayOfMonth(value *BackupDrBackupPlanBackupRulesStandardScheduleWeekDayOfMonth)
	ResetBackupWindow()
	ResetDaysOfMonth()
	ResetDaysOfWeek()
	ResetHourlyFrequency()
	ResetMonths()
	ResetWeekDayOfMonth()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BackupDrBackupPlanBackupRulesStandardScheduleOutputReference
type jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference) BackupWindow() BackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference {
	var returns BackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference
	_jsii_.Get(
		j,
		"backupWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference) BackupWindowInput() *BackupDrBackupPlanBackupRulesStandardScheduleBackupWindow {
	var returns *BackupDrBackupPlanBackupRulesStandardScheduleBackupWindow
	_jsii_.Get(
		j,
		"backupWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference) DaysOfMonth() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"daysOfMonth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference) DaysOfMonthInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"daysOfMonthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference) DaysOfWeek() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"daysOfWeek",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference) DaysOfWeekInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"daysOfWeekInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference) HourlyFrequency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hourlyFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference) HourlyFrequencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hourlyFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference) InternalValue() *BackupDrBackupPlanBackupRulesStandardSchedule {
	var returns *BackupDrBackupPlanBackupRulesStandardSchedule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference) Months() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"months",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference) MonthsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"monthsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference) RecurrenceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recurrenceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference) RecurrenceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recurrenceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference) TimeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference) WeekDayOfMonth() BackupDrBackupPlanBackupRulesStandardScheduleWeekDayOfMonthOutputReference {
	var returns BackupDrBackupPlanBackupRulesStandardScheduleWeekDayOfMonthOutputReference
	_jsii_.Get(
		j,
		"weekDayOfMonth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference) WeekDayOfMonthInput() *BackupDrBackupPlanBackupRulesStandardScheduleWeekDayOfMonth {
	var returns *BackupDrBackupPlanBackupRulesStandardScheduleWeekDayOfMonth
	_jsii_.Get(
		j,
		"weekDayOfMonthInput",
		&returns,
	)
	return returns
}


func NewBackupDrBackupPlanBackupRulesStandardScheduleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BackupDrBackupPlanBackupRulesStandardScheduleOutputReference {
	_init_.Initialize()

	if err := validateNewBackupDrBackupPlanBackupRulesStandardScheduleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.backupDrBackupPlan.BackupDrBackupPlanBackupRulesStandardScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBackupDrBackupPlanBackupRulesStandardScheduleOutputReference_Override(b BackupDrBackupPlanBackupRulesStandardScheduleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.backupDrBackupPlan.BackupDrBackupPlanBackupRulesStandardScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference)SetDaysOfMonth(val *[]*float64) {
	if err := j.validateSetDaysOfMonthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"daysOfMonth",
		val,
	)
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference)SetDaysOfWeek(val *[]*string) {
	if err := j.validateSetDaysOfWeekParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"daysOfWeek",
		val,
	)
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference)SetHourlyFrequency(val *float64) {
	if err := j.validateSetHourlyFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hourlyFrequency",
		val,
	)
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference)SetInternalValue(val *BackupDrBackupPlanBackupRulesStandardSchedule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference)SetMonths(val *[]*string) {
	if err := j.validateSetMonthsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"months",
		val,
	)
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference)SetRecurrenceType(val *string) {
	if err := j.validateSetRecurrenceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recurrenceType",
		val,
	)
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference)SetTimeZone(val *string) {
	if err := j.validateSetTimeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeZone",
		val,
	)
}

func (b *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference) PutBackupWindow(value *BackupDrBackupPlanBackupRulesStandardScheduleBackupWindow) {
	if err := b.validatePutBackupWindowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putBackupWindow",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference) PutWeekDayOfMonth(value *BackupDrBackupPlanBackupRulesStandardScheduleWeekDayOfMonth) {
	if err := b.validatePutWeekDayOfMonthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putWeekDayOfMonth",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference) ResetBackupWindow() {
	_jsii_.InvokeVoid(
		b,
		"resetBackupWindow",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference) ResetDaysOfMonth() {
	_jsii_.InvokeVoid(
		b,
		"resetDaysOfMonth",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference) ResetDaysOfWeek() {
	_jsii_.InvokeVoid(
		b,
		"resetDaysOfWeek",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference) ResetHourlyFrequency() {
	_jsii_.InvokeVoid(
		b,
		"resetHourlyFrequency",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference) ResetMonths() {
	_jsii_.InvokeVoid(
		b,
		"resetMonths",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference) ResetWeekDayOfMonth() {
	_jsii_.InvokeVoid(
		b,
		"resetWeekDayOfMonth",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := b.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

