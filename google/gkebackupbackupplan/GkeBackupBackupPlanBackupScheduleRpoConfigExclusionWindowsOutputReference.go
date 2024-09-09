// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkebackupbackupplan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/gkebackupbackupplan/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference interface {
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
	Daily() interface{}
	SetDaily(val interface{})
	DailyInput() interface{}
	DaysOfWeek() GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsDaysOfWeekOutputReference
	DaysOfWeekInput() *GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsDaysOfWeek
	Duration() *string
	SetDuration(val *string)
	DurationInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SingleOccurrenceDate() GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsSingleOccurrenceDateOutputReference
	SingleOccurrenceDateInput() *GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsSingleOccurrenceDate
	StartTime() GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsStartTimeOutputReference
	StartTimeInput() *GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsStartTime
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
	PutDaysOfWeek(value *GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsDaysOfWeek)
	PutSingleOccurrenceDate(value *GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsSingleOccurrenceDate)
	PutStartTime(value *GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsStartTime)
	ResetDaily()
	ResetDaysOfWeek()
	ResetSingleOccurrenceDate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference
type jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) Daily() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"daily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) DailyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dailyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) DaysOfWeek() GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsDaysOfWeekOutputReference {
	var returns GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsDaysOfWeekOutputReference
	_jsii_.Get(
		j,
		"daysOfWeek",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) DaysOfWeekInput() *GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsDaysOfWeek {
	var returns *GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsDaysOfWeek
	_jsii_.Get(
		j,
		"daysOfWeekInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) Duration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"duration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) DurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"durationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) SingleOccurrenceDate() GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsSingleOccurrenceDateOutputReference {
	var returns GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsSingleOccurrenceDateOutputReference
	_jsii_.Get(
		j,
		"singleOccurrenceDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) SingleOccurrenceDateInput() *GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsSingleOccurrenceDate {
	var returns *GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsSingleOccurrenceDate
	_jsii_.Get(
		j,
		"singleOccurrenceDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) StartTime() GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsStartTimeOutputReference {
	var returns GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsStartTimeOutputReference
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) StartTimeInput() *GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsStartTime {
	var returns *GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsStartTime
	_jsii_.Get(
		j,
		"startTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference {
	_init_.Initialize()

	if err := validateNewGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.gkeBackupBackupPlan.GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference_Override(g GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.gkeBackupBackupPlan.GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference)SetDaily(val interface{}) {
	if err := j.validateSetDailyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"daily",
		val,
	)
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference)SetDuration(val *string) {
	if err := j.validateSetDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"duration",
		val,
	)
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) PutDaysOfWeek(value *GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsDaysOfWeek) {
	if err := g.validatePutDaysOfWeekParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDaysOfWeek",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) PutSingleOccurrenceDate(value *GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsSingleOccurrenceDate) {
	if err := g.validatePutSingleOccurrenceDateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSingleOccurrenceDate",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) PutStartTime(value *GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsStartTime) {
	if err := g.validatePutStartTimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStartTime",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) ResetDaily() {
	_jsii_.InvokeVoid(
		g,
		"resetDaily",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) ResetDaysOfWeek() {
	_jsii_.InvokeVoid(
		g,
		"resetDaysOfWeek",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) ResetSingleOccurrenceDate() {
	_jsii_.InvokeVoid(
		g,
		"resetSingleOccurrenceDate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

