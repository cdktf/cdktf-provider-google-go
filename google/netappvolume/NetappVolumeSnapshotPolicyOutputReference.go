// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package netappvolume

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/netappvolume/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetappVolumeSnapshotPolicyOutputReference interface {
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
	DailySchedule() NetappVolumeSnapshotPolicyDailyScheduleOutputReference
	DailyScheduleInput() *NetappVolumeSnapshotPolicyDailySchedule
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	HourlySchedule() NetappVolumeSnapshotPolicyHourlyScheduleOutputReference
	HourlyScheduleInput() *NetappVolumeSnapshotPolicyHourlySchedule
	InternalValue() *NetappVolumeSnapshotPolicy
	SetInternalValue(val *NetappVolumeSnapshotPolicy)
	MonthlySchedule() NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference
	MonthlyScheduleInput() *NetappVolumeSnapshotPolicyMonthlySchedule
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WeeklySchedule() NetappVolumeSnapshotPolicyWeeklyScheduleOutputReference
	WeeklyScheduleInput() *NetappVolumeSnapshotPolicyWeeklySchedule
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutDailySchedule(value *NetappVolumeSnapshotPolicyDailySchedule)
	PutHourlySchedule(value *NetappVolumeSnapshotPolicyHourlySchedule)
	PutMonthlySchedule(value *NetappVolumeSnapshotPolicyMonthlySchedule)
	PutWeeklySchedule(value *NetappVolumeSnapshotPolicyWeeklySchedule)
	ResetDailySchedule()
	ResetEnabled()
	ResetHourlySchedule()
	ResetMonthlySchedule()
	ResetWeeklySchedule()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetappVolumeSnapshotPolicyOutputReference
type jsiiProxy_NetappVolumeSnapshotPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference) DailySchedule() NetappVolumeSnapshotPolicyDailyScheduleOutputReference {
	var returns NetappVolumeSnapshotPolicyDailyScheduleOutputReference
	_jsii_.Get(
		j,
		"dailySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference) DailyScheduleInput() *NetappVolumeSnapshotPolicyDailySchedule {
	var returns *NetappVolumeSnapshotPolicyDailySchedule
	_jsii_.Get(
		j,
		"dailyScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference) HourlySchedule() NetappVolumeSnapshotPolicyHourlyScheduleOutputReference {
	var returns NetappVolumeSnapshotPolicyHourlyScheduleOutputReference
	_jsii_.Get(
		j,
		"hourlySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference) HourlyScheduleInput() *NetappVolumeSnapshotPolicyHourlySchedule {
	var returns *NetappVolumeSnapshotPolicyHourlySchedule
	_jsii_.Get(
		j,
		"hourlyScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference) InternalValue() *NetappVolumeSnapshotPolicy {
	var returns *NetappVolumeSnapshotPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference) MonthlySchedule() NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference {
	var returns NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference
	_jsii_.Get(
		j,
		"monthlySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference) MonthlyScheduleInput() *NetappVolumeSnapshotPolicyMonthlySchedule {
	var returns *NetappVolumeSnapshotPolicyMonthlySchedule
	_jsii_.Get(
		j,
		"monthlyScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference) WeeklySchedule() NetappVolumeSnapshotPolicyWeeklyScheduleOutputReference {
	var returns NetappVolumeSnapshotPolicyWeeklyScheduleOutputReference
	_jsii_.Get(
		j,
		"weeklySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference) WeeklyScheduleInput() *NetappVolumeSnapshotPolicyWeeklySchedule {
	var returns *NetappVolumeSnapshotPolicyWeeklySchedule
	_jsii_.Get(
		j,
		"weeklyScheduleInput",
		&returns,
	)
	return returns
}


func NewNetappVolumeSnapshotPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NetappVolumeSnapshotPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewNetappVolumeSnapshotPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetappVolumeSnapshotPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.netappVolume.NetappVolumeSnapshotPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetappVolumeSnapshotPolicyOutputReference_Override(n NetappVolumeSnapshotPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.netappVolume.NetappVolumeSnapshotPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference)SetInternalValue(val *NetappVolumeSnapshotPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference) PutDailySchedule(value *NetappVolumeSnapshotPolicyDailySchedule) {
	if err := n.validatePutDailyScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putDailySchedule",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference) PutHourlySchedule(value *NetappVolumeSnapshotPolicyHourlySchedule) {
	if err := n.validatePutHourlyScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putHourlySchedule",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference) PutMonthlySchedule(value *NetappVolumeSnapshotPolicyMonthlySchedule) {
	if err := n.validatePutMonthlyScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putMonthlySchedule",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference) PutWeeklySchedule(value *NetappVolumeSnapshotPolicyWeeklySchedule) {
	if err := n.validatePutWeeklyScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putWeeklySchedule",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference) ResetDailySchedule() {
	_jsii_.InvokeVoid(
		n,
		"resetDailySchedule",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference) ResetHourlySchedule() {
	_jsii_.InvokeVoid(
		n,
		"resetHourlySchedule",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference) ResetMonthlySchedule() {
	_jsii_.InvokeVoid(
		n,
		"resetMonthlySchedule",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference) ResetWeeklySchedule() {
	_jsii_.InvokeVoid(
		n,
		"resetWeeklySchedule",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := n.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeSnapshotPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

