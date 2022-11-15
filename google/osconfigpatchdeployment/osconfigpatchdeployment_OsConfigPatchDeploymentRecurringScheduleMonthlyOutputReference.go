package osconfigpatchdeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v4/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v4/osconfigpatchdeployment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference interface {
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
	InternalValue() *OsConfigPatchDeploymentRecurringScheduleMonthly
	SetInternalValue(val *OsConfigPatchDeploymentRecurringScheduleMonthly)
	MonthDay() *float64
	SetMonthDay(val *float64)
	MonthDayInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WeekDayOfMonth() OsConfigPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthOutputReference
	WeekDayOfMonthInput() *OsConfigPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth
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
	PutWeekDayOfMonth(value *OsConfigPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth)
	ResetMonthDay()
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

// The jsii proxy struct for OsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference
type jsiiProxy_OsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference) InternalValue() *OsConfigPatchDeploymentRecurringScheduleMonthly {
	var returns *OsConfigPatchDeploymentRecurringScheduleMonthly
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference) MonthDay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monthDay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference) MonthDayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monthDayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference) WeekDayOfMonth() OsConfigPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthOutputReference {
	var returns OsConfigPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthOutputReference
	_jsii_.Get(
		j,
		"weekDayOfMonth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference) WeekDayOfMonthInput() *OsConfigPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth {
	var returns *OsConfigPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth
	_jsii_.Get(
		j,
		"weekDayOfMonthInput",
		&returns,
	)
	return returns
}


func NewOsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference {
	_init_.Initialize()

	if err := validateNewOsConfigPatchDeploymentRecurringScheduleMonthlyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.osConfigPatchDeployment.OsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference_Override(o OsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.osConfigPatchDeployment.OsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference)SetInternalValue(val *OsConfigPatchDeploymentRecurringScheduleMonthly) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference)SetMonthDay(val *float64) {
	if err := j.validateSetMonthDayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monthDay",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference) PutWeekDayOfMonth(value *OsConfigPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth) {
	if err := o.validatePutWeekDayOfMonthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putWeekDayOfMonth",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference) ResetMonthDay() {
	_jsii_.InvokeVoid(
		o,
		"resetMonthDay",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference) ResetWeekDayOfMonth() {
	_jsii_.InvokeVoid(
		o,
		"resetWeekDayOfMonth",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OsConfigPatchDeploymentRecurringScheduleMonthlyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

