// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alloydbcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v11/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v11/alloydbcluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlloydbClusterAutomatedBackupPolicyOutputReference interface {
	cdktf.ComplexObject
	BackupWindow() *string
	SetBackupWindow(val *string)
	BackupWindowInput() *string
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	EncryptionConfig() AlloydbClusterAutomatedBackupPolicyEncryptionConfigOutputReference
	EncryptionConfigInput() *AlloydbClusterAutomatedBackupPolicyEncryptionConfig
	// Experimental.
	Fqn() *string
	InternalValue() *AlloydbClusterAutomatedBackupPolicy
	SetInternalValue(val *AlloydbClusterAutomatedBackupPolicy)
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	QuantityBasedRetention() AlloydbClusterAutomatedBackupPolicyQuantityBasedRetentionOutputReference
	QuantityBasedRetentionInput() *AlloydbClusterAutomatedBackupPolicyQuantityBasedRetention
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeBasedRetention() AlloydbClusterAutomatedBackupPolicyTimeBasedRetentionOutputReference
	TimeBasedRetentionInput() *AlloydbClusterAutomatedBackupPolicyTimeBasedRetention
	WeeklySchedule() AlloydbClusterAutomatedBackupPolicyWeeklyScheduleOutputReference
	WeeklyScheduleInput() *AlloydbClusterAutomatedBackupPolicyWeeklySchedule
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
	PutEncryptionConfig(value *AlloydbClusterAutomatedBackupPolicyEncryptionConfig)
	PutQuantityBasedRetention(value *AlloydbClusterAutomatedBackupPolicyQuantityBasedRetention)
	PutTimeBasedRetention(value *AlloydbClusterAutomatedBackupPolicyTimeBasedRetention)
	PutWeeklySchedule(value *AlloydbClusterAutomatedBackupPolicyWeeklySchedule)
	ResetBackupWindow()
	ResetEnabled()
	ResetEncryptionConfig()
	ResetLabels()
	ResetLocation()
	ResetQuantityBasedRetention()
	ResetTimeBasedRetention()
	ResetWeeklySchedule()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AlloydbClusterAutomatedBackupPolicyOutputReference
type jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) BackupWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) BackupWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) EncryptionConfig() AlloydbClusterAutomatedBackupPolicyEncryptionConfigOutputReference {
	var returns AlloydbClusterAutomatedBackupPolicyEncryptionConfigOutputReference
	_jsii_.Get(
		j,
		"encryptionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) EncryptionConfigInput() *AlloydbClusterAutomatedBackupPolicyEncryptionConfig {
	var returns *AlloydbClusterAutomatedBackupPolicyEncryptionConfig
	_jsii_.Get(
		j,
		"encryptionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) InternalValue() *AlloydbClusterAutomatedBackupPolicy {
	var returns *AlloydbClusterAutomatedBackupPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) QuantityBasedRetention() AlloydbClusterAutomatedBackupPolicyQuantityBasedRetentionOutputReference {
	var returns AlloydbClusterAutomatedBackupPolicyQuantityBasedRetentionOutputReference
	_jsii_.Get(
		j,
		"quantityBasedRetention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) QuantityBasedRetentionInput() *AlloydbClusterAutomatedBackupPolicyQuantityBasedRetention {
	var returns *AlloydbClusterAutomatedBackupPolicyQuantityBasedRetention
	_jsii_.Get(
		j,
		"quantityBasedRetentionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) TimeBasedRetention() AlloydbClusterAutomatedBackupPolicyTimeBasedRetentionOutputReference {
	var returns AlloydbClusterAutomatedBackupPolicyTimeBasedRetentionOutputReference
	_jsii_.Get(
		j,
		"timeBasedRetention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) TimeBasedRetentionInput() *AlloydbClusterAutomatedBackupPolicyTimeBasedRetention {
	var returns *AlloydbClusterAutomatedBackupPolicyTimeBasedRetention
	_jsii_.Get(
		j,
		"timeBasedRetentionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) WeeklySchedule() AlloydbClusterAutomatedBackupPolicyWeeklyScheduleOutputReference {
	var returns AlloydbClusterAutomatedBackupPolicyWeeklyScheduleOutputReference
	_jsii_.Get(
		j,
		"weeklySchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) WeeklyScheduleInput() *AlloydbClusterAutomatedBackupPolicyWeeklySchedule {
	var returns *AlloydbClusterAutomatedBackupPolicyWeeklySchedule
	_jsii_.Get(
		j,
		"weeklyScheduleInput",
		&returns,
	)
	return returns
}


func NewAlloydbClusterAutomatedBackupPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AlloydbClusterAutomatedBackupPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewAlloydbClusterAutomatedBackupPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.alloydbCluster.AlloydbClusterAutomatedBackupPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAlloydbClusterAutomatedBackupPolicyOutputReference_Override(a AlloydbClusterAutomatedBackupPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.alloydbCluster.AlloydbClusterAutomatedBackupPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference)SetBackupWindow(val *string) {
	if err := j.validateSetBackupWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupWindow",
		val,
	)
}

func (j *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference)SetInternalValue(val *AlloydbClusterAutomatedBackupPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) PutEncryptionConfig(value *AlloydbClusterAutomatedBackupPolicyEncryptionConfig) {
	if err := a.validatePutEncryptionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEncryptionConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) PutQuantityBasedRetention(value *AlloydbClusterAutomatedBackupPolicyQuantityBasedRetention) {
	if err := a.validatePutQuantityBasedRetentionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putQuantityBasedRetention",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) PutTimeBasedRetention(value *AlloydbClusterAutomatedBackupPolicyTimeBasedRetention) {
	if err := a.validatePutTimeBasedRetentionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeBasedRetention",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) PutWeeklySchedule(value *AlloydbClusterAutomatedBackupPolicyWeeklySchedule) {
	if err := a.validatePutWeeklyScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWeeklySchedule",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) ResetBackupWindow() {
	_jsii_.InvokeVoid(
		a,
		"resetBackupWindow",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) ResetEncryptionConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetEncryptionConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		a,
		"resetLabels",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) ResetLocation() {
	_jsii_.InvokeVoid(
		a,
		"resetLocation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) ResetQuantityBasedRetention() {
	_jsii_.InvokeVoid(
		a,
		"resetQuantityBasedRetention",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) ResetTimeBasedRetention() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeBasedRetention",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) ResetWeeklySchedule() {
	_jsii_.InvokeVoid(
		a,
		"resetWeeklySchedule",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlloydbClusterAutomatedBackupPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

