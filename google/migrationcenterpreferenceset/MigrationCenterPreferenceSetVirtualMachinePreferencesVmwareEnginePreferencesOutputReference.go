// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package migrationcenterpreferenceset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/migrationcenterpreferenceset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference interface {
	cdktf.ComplexObject
	CommitmentPlan() *string
	SetCommitmentPlan(val *string)
	CommitmentPlanInput() *string
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
	CpuOvercommitRatio() *float64
	SetCpuOvercommitRatio(val *float64)
	CpuOvercommitRatioInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferences
	SetInternalValue(val *MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferences)
	MemoryOvercommitRatio() *float64
	SetMemoryOvercommitRatio(val *float64)
	MemoryOvercommitRatioInput() *float64
	StorageDeduplicationCompressionRatio() *float64
	SetStorageDeduplicationCompressionRatio(val *float64)
	StorageDeduplicationCompressionRatioInput() *float64
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
	ResetCommitmentPlan()
	ResetCpuOvercommitRatio()
	ResetMemoryOvercommitRatio()
	ResetStorageDeduplicationCompressionRatio()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference
type jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference) CommitmentPlan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitmentPlan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference) CommitmentPlanInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitmentPlanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference) CpuOvercommitRatio() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuOvercommitRatio",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference) CpuOvercommitRatioInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuOvercommitRatioInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference) InternalValue() *MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferences {
	var returns *MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferences
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference) MemoryOvercommitRatio() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryOvercommitRatio",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference) MemoryOvercommitRatioInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryOvercommitRatioInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference) StorageDeduplicationCompressionRatio() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageDeduplicationCompressionRatio",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference) StorageDeduplicationCompressionRatioInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageDeduplicationCompressionRatioInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference {
	_init_.Initialize()

	if err := validateNewMigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.migrationCenterPreferenceSet.MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference_Override(m MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.migrationCenterPreferenceSet.MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference)SetCommitmentPlan(val *string) {
	if err := j.validateSetCommitmentPlanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commitmentPlan",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference)SetCpuOvercommitRatio(val *float64) {
	if err := j.validateSetCpuOvercommitRatioParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuOvercommitRatio",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference)SetInternalValue(val *MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferences) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference)SetMemoryOvercommitRatio(val *float64) {
	if err := j.validateSetMemoryOvercommitRatioParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryOvercommitRatio",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference)SetStorageDeduplicationCompressionRatio(val *float64) {
	if err := j.validateSetStorageDeduplicationCompressionRatioParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageDeduplicationCompressionRatio",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference) ResetCommitmentPlan() {
	_jsii_.InvokeVoid(
		m,
		"resetCommitmentPlan",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference) ResetCpuOvercommitRatio() {
	_jsii_.InvokeVoid(
		m,
		"resetCpuOvercommitRatio",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference) ResetMemoryOvercommitRatio() {
	_jsii_.InvokeVoid(
		m,
		"resetMemoryOvercommitRatio",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference) ResetStorageDeduplicationCompressionRatio() {
	_jsii_.InvokeVoid(
		m,
		"resetStorageDeduplicationCompressionRatio",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

