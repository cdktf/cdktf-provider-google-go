// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oracledatabasecloudexadatainfrastructure

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/oracledatabasecloudexadatainfrastructure/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference interface {
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
	CustomActionTimeoutMins() *float64
	SetCustomActionTimeoutMins(val *float64)
	CustomActionTimeoutMinsInput() *float64
	DaysOfWeek() *[]*string
	SetDaysOfWeek(val *[]*string)
	DaysOfWeekInput() *[]*string
	// Experimental.
	Fqn() *string
	HoursOfDay() *[]*float64
	SetHoursOfDay(val *[]*float64)
	HoursOfDayInput() *[]*float64
	InternalValue() *OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindow
	SetInternalValue(val *OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindow)
	IsCustomActionTimeoutEnabled() interface{}
	SetIsCustomActionTimeoutEnabled(val interface{})
	IsCustomActionTimeoutEnabledInput() interface{}
	LeadTimeWeek() *float64
	SetLeadTimeWeek(val *float64)
	LeadTimeWeekInput() *float64
	Months() *[]*string
	SetMonths(val *[]*string)
	MonthsInput() *[]*string
	PatchingMode() *string
	SetPatchingMode(val *string)
	PatchingModeInput() *string
	Preference() *string
	SetPreference(val *string)
	PreferenceInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WeeksOfMonth() *[]*float64
	SetWeeksOfMonth(val *[]*float64)
	WeeksOfMonthInput() *[]*float64
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
	ResetCustomActionTimeoutMins()
	ResetDaysOfWeek()
	ResetHoursOfDay()
	ResetIsCustomActionTimeoutEnabled()
	ResetLeadTimeWeek()
	ResetMonths()
	ResetPatchingMode()
	ResetPreference()
	ResetWeeksOfMonth()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference
type jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) CustomActionTimeoutMins() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"customActionTimeoutMins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) CustomActionTimeoutMinsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"customActionTimeoutMinsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) DaysOfWeek() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"daysOfWeek",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) DaysOfWeekInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"daysOfWeekInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) HoursOfDay() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"hoursOfDay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) HoursOfDayInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"hoursOfDayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) InternalValue() *OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindow {
	var returns *OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindow
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) IsCustomActionTimeoutEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isCustomActionTimeoutEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) IsCustomActionTimeoutEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isCustomActionTimeoutEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) LeadTimeWeek() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"leadTimeWeek",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) LeadTimeWeekInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"leadTimeWeekInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) Months() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"months",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) MonthsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"monthsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) PatchingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patchingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) PatchingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patchingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) Preference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) PreferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) WeeksOfMonth() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"weeksOfMonth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) WeeksOfMonthInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"weeksOfMonthInput",
		&returns,
	)
	return returns
}


func NewOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference {
	_init_.Initialize()

	if err := validateNewOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.oracleDatabaseCloudExadataInfrastructure.OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference_Override(o OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.oracleDatabaseCloudExadataInfrastructure.OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference)SetCustomActionTimeoutMins(val *float64) {
	if err := j.validateSetCustomActionTimeoutMinsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customActionTimeoutMins",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference)SetDaysOfWeek(val *[]*string) {
	if err := j.validateSetDaysOfWeekParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"daysOfWeek",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference)SetHoursOfDay(val *[]*float64) {
	if err := j.validateSetHoursOfDayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hoursOfDay",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference)SetInternalValue(val *OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindow) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference)SetIsCustomActionTimeoutEnabled(val interface{}) {
	if err := j.validateSetIsCustomActionTimeoutEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isCustomActionTimeoutEnabled",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference)SetLeadTimeWeek(val *float64) {
	if err := j.validateSetLeadTimeWeekParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"leadTimeWeek",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference)SetMonths(val *[]*string) {
	if err := j.validateSetMonthsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"months",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference)SetPatchingMode(val *string) {
	if err := j.validateSetPatchingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"patchingMode",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference)SetPreference(val *string) {
	if err := j.validateSetPreferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preference",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference)SetWeeksOfMonth(val *[]*float64) {
	if err := j.validateSetWeeksOfMonthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weeksOfMonth",
		val,
	)
}

func (o *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) ResetCustomActionTimeoutMins() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomActionTimeoutMins",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) ResetDaysOfWeek() {
	_jsii_.InvokeVoid(
		o,
		"resetDaysOfWeek",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) ResetHoursOfDay() {
	_jsii_.InvokeVoid(
		o,
		"resetHoursOfDay",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) ResetIsCustomActionTimeoutEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetIsCustomActionTimeoutEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) ResetLeadTimeWeek() {
	_jsii_.InvokeVoid(
		o,
		"resetLeadTimeWeek",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) ResetMonths() {
	_jsii_.InvokeVoid(
		o,
		"resetMonths",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) ResetPatchingMode() {
	_jsii_.InvokeVoid(
		o,
		"resetPatchingMode",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) ResetPreference() {
	_jsii_.InvokeVoid(
		o,
		"resetPreference",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) ResetWeeksOfMonth() {
	_jsii_.InvokeVoid(
		o,
		"resetWeeksOfMonth",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindowOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

