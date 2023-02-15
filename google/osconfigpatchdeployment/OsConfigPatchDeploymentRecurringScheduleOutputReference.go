package osconfigpatchdeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v5/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v5/osconfigpatchdeployment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OsConfigPatchDeploymentRecurringScheduleOutputReference interface {
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
	EndTime() *string
	SetEndTime(val *string)
	EndTimeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *OsConfigPatchDeploymentRecurringSchedule
	SetInternalValue(val *OsConfigPatchDeploymentRecurringSchedule)
	LastExecuteTime() *string
	Monthly() OsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference
	MonthlyInput() *OsConfigPatchDeploymentRecurringScheduleMonthly
	NextExecuteTime() *string
	StartTime() *string
	SetStartTime(val *string)
	StartTimeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeOfDay() OsConfigPatchDeploymentRecurringScheduleTimeOfDayOutputReference
	TimeOfDayInput() *OsConfigPatchDeploymentRecurringScheduleTimeOfDay
	TimeZone() OsConfigPatchDeploymentRecurringScheduleTimeZoneOutputReference
	TimeZoneInput() *OsConfigPatchDeploymentRecurringScheduleTimeZone
	Weekly() OsConfigPatchDeploymentRecurringScheduleWeeklyOutputReference
	WeeklyInput() *OsConfigPatchDeploymentRecurringScheduleWeekly
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
	PutMonthly(value *OsConfigPatchDeploymentRecurringScheduleMonthly)
	PutTimeOfDay(value *OsConfigPatchDeploymentRecurringScheduleTimeOfDay)
	PutTimeZone(value *OsConfigPatchDeploymentRecurringScheduleTimeZone)
	PutWeekly(value *OsConfigPatchDeploymentRecurringScheduleWeekly)
	ResetEndTime()
	ResetMonthly()
	ResetStartTime()
	ResetWeekly()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OsConfigPatchDeploymentRecurringScheduleOutputReference
type jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference) EndTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference) EndTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference) InternalValue() *OsConfigPatchDeploymentRecurringSchedule {
	var returns *OsConfigPatchDeploymentRecurringSchedule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference) LastExecuteTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastExecuteTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference) Monthly() OsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference {
	var returns OsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference
	_jsii_.Get(
		j,
		"monthly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference) MonthlyInput() *OsConfigPatchDeploymentRecurringScheduleMonthly {
	var returns *OsConfigPatchDeploymentRecurringScheduleMonthly
	_jsii_.Get(
		j,
		"monthlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference) NextExecuteTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextExecuteTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference) StartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference) StartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference) TimeOfDay() OsConfigPatchDeploymentRecurringScheduleTimeOfDayOutputReference {
	var returns OsConfigPatchDeploymentRecurringScheduleTimeOfDayOutputReference
	_jsii_.Get(
		j,
		"timeOfDay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference) TimeOfDayInput() *OsConfigPatchDeploymentRecurringScheduleTimeOfDay {
	var returns *OsConfigPatchDeploymentRecurringScheduleTimeOfDay
	_jsii_.Get(
		j,
		"timeOfDayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference) TimeZone() OsConfigPatchDeploymentRecurringScheduleTimeZoneOutputReference {
	var returns OsConfigPatchDeploymentRecurringScheduleTimeZoneOutputReference
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference) TimeZoneInput() *OsConfigPatchDeploymentRecurringScheduleTimeZone {
	var returns *OsConfigPatchDeploymentRecurringScheduleTimeZone
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference) Weekly() OsConfigPatchDeploymentRecurringScheduleWeeklyOutputReference {
	var returns OsConfigPatchDeploymentRecurringScheduleWeeklyOutputReference
	_jsii_.Get(
		j,
		"weekly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference) WeeklyInput() *OsConfigPatchDeploymentRecurringScheduleWeekly {
	var returns *OsConfigPatchDeploymentRecurringScheduleWeekly
	_jsii_.Get(
		j,
		"weeklyInput",
		&returns,
	)
	return returns
}


func NewOsConfigPatchDeploymentRecurringScheduleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OsConfigPatchDeploymentRecurringScheduleOutputReference {
	_init_.Initialize()

	if err := validateNewOsConfigPatchDeploymentRecurringScheduleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.osConfigPatchDeployment.OsConfigPatchDeploymentRecurringScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOsConfigPatchDeploymentRecurringScheduleOutputReference_Override(o OsConfigPatchDeploymentRecurringScheduleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.osConfigPatchDeployment.OsConfigPatchDeploymentRecurringScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference)SetEndTime(val *string) {
	if err := j.validateSetEndTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endTime",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference)SetInternalValue(val *OsConfigPatchDeploymentRecurringSchedule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference)SetStartTime(val *string) {
	if err := j.validateSetStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startTime",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference) PutMonthly(value *OsConfigPatchDeploymentRecurringScheduleMonthly) {
	if err := o.validatePutMonthlyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putMonthly",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference) PutTimeOfDay(value *OsConfigPatchDeploymentRecurringScheduleTimeOfDay) {
	if err := o.validatePutTimeOfDayParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putTimeOfDay",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference) PutTimeZone(value *OsConfigPatchDeploymentRecurringScheduleTimeZone) {
	if err := o.validatePutTimeZoneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putTimeZone",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference) PutWeekly(value *OsConfigPatchDeploymentRecurringScheduleWeekly) {
	if err := o.validatePutWeeklyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putWeekly",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference) ResetEndTime() {
	_jsii_.InvokeVoid(
		o,
		"resetEndTime",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference) ResetMonthly() {
	_jsii_.InvokeVoid(
		o,
		"resetMonthly",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference) ResetStartTime() {
	_jsii_.InvokeVoid(
		o,
		"resetStartTime",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference) ResetWeekly() {
	_jsii_.InvokeVoid(
		o,
		"resetWeekly",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

