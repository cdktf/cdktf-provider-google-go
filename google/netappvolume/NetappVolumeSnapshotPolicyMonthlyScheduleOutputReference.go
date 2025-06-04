// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package netappvolume

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/netappvolume/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference interface {
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
	DaysOfMonth() *string
	SetDaysOfMonth(val *string)
	DaysOfMonthInput() *string
	// Experimental.
	Fqn() *string
	Hour() *float64
	SetHour(val *float64)
	HourInput() *float64
	InternalValue() *NetappVolumeSnapshotPolicyMonthlySchedule
	SetInternalValue(val *NetappVolumeSnapshotPolicyMonthlySchedule)
	Minute() *float64
	SetMinute(val *float64)
	MinuteInput() *float64
	SnapshotsToKeep() *float64
	SetSnapshotsToKeep(val *float64)
	SnapshotsToKeepInput() *float64
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
	ResetDaysOfMonth()
	ResetHour()
	ResetMinute()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference
type jsiiProxy_NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference) DaysOfMonth() *string {
	var returns *string
	_jsii_.Get(
		j,
		"daysOfMonth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference) DaysOfMonthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"daysOfMonthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference) Hour() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hour",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference) HourInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hourInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference) InternalValue() *NetappVolumeSnapshotPolicyMonthlySchedule {
	var returns *NetappVolumeSnapshotPolicyMonthlySchedule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference) Minute() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference) MinuteInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minuteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference) SnapshotsToKeep() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"snapshotsToKeep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference) SnapshotsToKeepInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"snapshotsToKeepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNetappVolumeSnapshotPolicyMonthlyScheduleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference {
	_init_.Initialize()

	if err := validateNewNetappVolumeSnapshotPolicyMonthlyScheduleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.netappVolume.NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetappVolumeSnapshotPolicyMonthlyScheduleOutputReference_Override(n NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.netappVolume.NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference)SetDaysOfMonth(val *string) {
	if err := j.validateSetDaysOfMonthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"daysOfMonth",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference)SetHour(val *float64) {
	if err := j.validateSetHourParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hour",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference)SetInternalValue(val *NetappVolumeSnapshotPolicyMonthlySchedule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference)SetMinute(val *float64) {
	if err := j.validateSetMinuteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minute",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference)SetSnapshotsToKeep(val *float64) {
	if err := j.validateSetSnapshotsToKeepParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotsToKeep",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference) ResetDaysOfMonth() {
	_jsii_.InvokeVoid(
		n,
		"resetDaysOfMonth",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference) ResetHour() {
	_jsii_.InvokeVoid(
		n,
		"resetHour",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference) ResetMinute() {
	_jsii_.InvokeVoid(
		n,
		"resetMinute",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := n.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeSnapshotPolicyMonthlyScheduleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

