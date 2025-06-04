// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package backupdrbackupplan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/backupdrbackupplan/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference interface {
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
	EndHourOfDay() *float64
	SetEndHourOfDay(val *float64)
	EndHourOfDayInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *BackupDrBackupPlanBackupRulesStandardScheduleBackupWindow
	SetInternalValue(val *BackupDrBackupPlanBackupRulesStandardScheduleBackupWindow)
	StartHourOfDay() *float64
	SetStartHourOfDay(val *float64)
	StartHourOfDayInput() *float64
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
	ResetEndHourOfDay()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference
type jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference) EndHourOfDay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"endHourOfDay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference) EndHourOfDayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"endHourOfDayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference) InternalValue() *BackupDrBackupPlanBackupRulesStandardScheduleBackupWindow {
	var returns *BackupDrBackupPlanBackupRulesStandardScheduleBackupWindow
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference) StartHourOfDay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startHourOfDay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference) StartHourOfDayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startHourOfDayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference {
	_init_.Initialize()

	if err := validateNewBackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.backupDrBackupPlan.BackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference_Override(b BackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.backupDrBackupPlan.BackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference)SetEndHourOfDay(val *float64) {
	if err := j.validateSetEndHourOfDayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endHourOfDay",
		val,
	)
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference)SetInternalValue(val *BackupDrBackupPlanBackupRulesStandardScheduleBackupWindow) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference)SetStartHourOfDay(val *float64) {
	if err := j.validateSetStartHourOfDayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startHourOfDay",
		val,
	)
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference) ResetEndHourOfDay() {
	_jsii_.InvokeVoid(
		b,
		"resetEndHourOfDay",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (b *jsiiProxy_BackupDrBackupPlanBackupRulesStandardScheduleBackupWindowOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

