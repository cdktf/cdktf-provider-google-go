// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lookerinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v11/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v11/lookerinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LookerInstanceDenyMaintenancePeriodOutputReference interface {
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
	EndDate() LookerInstanceDenyMaintenancePeriodEndDateOutputReference
	EndDateInput() *LookerInstanceDenyMaintenancePeriodEndDate
	// Experimental.
	Fqn() *string
	InternalValue() *LookerInstanceDenyMaintenancePeriod
	SetInternalValue(val *LookerInstanceDenyMaintenancePeriod)
	StartDate() LookerInstanceDenyMaintenancePeriodStartDateOutputReference
	StartDateInput() *LookerInstanceDenyMaintenancePeriodStartDate
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Time() LookerInstanceDenyMaintenancePeriodTimeOutputReference
	TimeInput() *LookerInstanceDenyMaintenancePeriodTime
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
	PutEndDate(value *LookerInstanceDenyMaintenancePeriodEndDate)
	PutStartDate(value *LookerInstanceDenyMaintenancePeriodStartDate)
	PutTime(value *LookerInstanceDenyMaintenancePeriodTime)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LookerInstanceDenyMaintenancePeriodOutputReference
type jsiiProxy_LookerInstanceDenyMaintenancePeriodOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LookerInstanceDenyMaintenancePeriodOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstanceDenyMaintenancePeriodOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstanceDenyMaintenancePeriodOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstanceDenyMaintenancePeriodOutputReference) EndDate() LookerInstanceDenyMaintenancePeriodEndDateOutputReference {
	var returns LookerInstanceDenyMaintenancePeriodEndDateOutputReference
	_jsii_.Get(
		j,
		"endDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstanceDenyMaintenancePeriodOutputReference) EndDateInput() *LookerInstanceDenyMaintenancePeriodEndDate {
	var returns *LookerInstanceDenyMaintenancePeriodEndDate
	_jsii_.Get(
		j,
		"endDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstanceDenyMaintenancePeriodOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstanceDenyMaintenancePeriodOutputReference) InternalValue() *LookerInstanceDenyMaintenancePeriod {
	var returns *LookerInstanceDenyMaintenancePeriod
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstanceDenyMaintenancePeriodOutputReference) StartDate() LookerInstanceDenyMaintenancePeriodStartDateOutputReference {
	var returns LookerInstanceDenyMaintenancePeriodStartDateOutputReference
	_jsii_.Get(
		j,
		"startDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstanceDenyMaintenancePeriodOutputReference) StartDateInput() *LookerInstanceDenyMaintenancePeriodStartDate {
	var returns *LookerInstanceDenyMaintenancePeriodStartDate
	_jsii_.Get(
		j,
		"startDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstanceDenyMaintenancePeriodOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstanceDenyMaintenancePeriodOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstanceDenyMaintenancePeriodOutputReference) Time() LookerInstanceDenyMaintenancePeriodTimeOutputReference {
	var returns LookerInstanceDenyMaintenancePeriodTimeOutputReference
	_jsii_.Get(
		j,
		"time",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstanceDenyMaintenancePeriodOutputReference) TimeInput() *LookerInstanceDenyMaintenancePeriodTime {
	var returns *LookerInstanceDenyMaintenancePeriodTime
	_jsii_.Get(
		j,
		"timeInput",
		&returns,
	)
	return returns
}


func NewLookerInstanceDenyMaintenancePeriodOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LookerInstanceDenyMaintenancePeriodOutputReference {
	_init_.Initialize()

	if err := validateNewLookerInstanceDenyMaintenancePeriodOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LookerInstanceDenyMaintenancePeriodOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.lookerInstance.LookerInstanceDenyMaintenancePeriodOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLookerInstanceDenyMaintenancePeriodOutputReference_Override(l LookerInstanceDenyMaintenancePeriodOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.lookerInstance.LookerInstanceDenyMaintenancePeriodOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LookerInstanceDenyMaintenancePeriodOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LookerInstanceDenyMaintenancePeriodOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LookerInstanceDenyMaintenancePeriodOutputReference)SetInternalValue(val *LookerInstanceDenyMaintenancePeriod) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LookerInstanceDenyMaintenancePeriodOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LookerInstanceDenyMaintenancePeriodOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LookerInstanceDenyMaintenancePeriodOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookerInstanceDenyMaintenancePeriodOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookerInstanceDenyMaintenancePeriodOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookerInstanceDenyMaintenancePeriodOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookerInstanceDenyMaintenancePeriodOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookerInstanceDenyMaintenancePeriodOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookerInstanceDenyMaintenancePeriodOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookerInstanceDenyMaintenancePeriodOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookerInstanceDenyMaintenancePeriodOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookerInstanceDenyMaintenancePeriodOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookerInstanceDenyMaintenancePeriodOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookerInstanceDenyMaintenancePeriodOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookerInstanceDenyMaintenancePeriodOutputReference) PutEndDate(value *LookerInstanceDenyMaintenancePeriodEndDate) {
	if err := l.validatePutEndDateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putEndDate",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LookerInstanceDenyMaintenancePeriodOutputReference) PutStartDate(value *LookerInstanceDenyMaintenancePeriodStartDate) {
	if err := l.validatePutStartDateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putStartDate",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LookerInstanceDenyMaintenancePeriodOutputReference) PutTime(value *LookerInstanceDenyMaintenancePeriodTime) {
	if err := l.validatePutTimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTime",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LookerInstanceDenyMaintenancePeriodOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookerInstanceDenyMaintenancePeriodOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

