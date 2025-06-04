// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package migrationcenterpreferenceset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/migrationcenterpreferenceset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference interface {
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
	ComputeEnginePreferences() MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference
	ComputeEnginePreferencesInput() *MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferences
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *MigrationCenterPreferenceSetVirtualMachinePreferences
	SetInternalValue(val *MigrationCenterPreferenceSetVirtualMachinePreferences)
	RegionPreferences() MigrationCenterPreferenceSetVirtualMachinePreferencesRegionPreferencesOutputReference
	RegionPreferencesInput() *MigrationCenterPreferenceSetVirtualMachinePreferencesRegionPreferences
	SizingOptimizationStrategy() *string
	SetSizingOptimizationStrategy(val *string)
	SizingOptimizationStrategyInput() *string
	SoleTenancyPreferences() MigrationCenterPreferenceSetVirtualMachinePreferencesSoleTenancyPreferencesOutputReference
	SoleTenancyPreferencesInput() *MigrationCenterPreferenceSetVirtualMachinePreferencesSoleTenancyPreferences
	TargetProduct() *string
	SetTargetProduct(val *string)
	TargetProductInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VmwareEnginePreferences() MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference
	VmwareEnginePreferencesInput() *MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferences
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
	PutComputeEnginePreferences(value *MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferences)
	PutRegionPreferences(value *MigrationCenterPreferenceSetVirtualMachinePreferencesRegionPreferences)
	PutSoleTenancyPreferences(value *MigrationCenterPreferenceSetVirtualMachinePreferencesSoleTenancyPreferences)
	PutVmwareEnginePreferences(value *MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferences)
	ResetCommitmentPlan()
	ResetComputeEnginePreferences()
	ResetRegionPreferences()
	ResetSizingOptimizationStrategy()
	ResetSoleTenancyPreferences()
	ResetTargetProduct()
	ResetVmwareEnginePreferences()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference
type jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) CommitmentPlan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitmentPlan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) CommitmentPlanInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitmentPlanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) ComputeEnginePreferences() MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference {
	var returns MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferencesOutputReference
	_jsii_.Get(
		j,
		"computeEnginePreferences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) ComputeEnginePreferencesInput() *MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferences {
	var returns *MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferences
	_jsii_.Get(
		j,
		"computeEnginePreferencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) InternalValue() *MigrationCenterPreferenceSetVirtualMachinePreferences {
	var returns *MigrationCenterPreferenceSetVirtualMachinePreferences
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) RegionPreferences() MigrationCenterPreferenceSetVirtualMachinePreferencesRegionPreferencesOutputReference {
	var returns MigrationCenterPreferenceSetVirtualMachinePreferencesRegionPreferencesOutputReference
	_jsii_.Get(
		j,
		"regionPreferences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) RegionPreferencesInput() *MigrationCenterPreferenceSetVirtualMachinePreferencesRegionPreferences {
	var returns *MigrationCenterPreferenceSetVirtualMachinePreferencesRegionPreferences
	_jsii_.Get(
		j,
		"regionPreferencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) SizingOptimizationStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizingOptimizationStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) SizingOptimizationStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizingOptimizationStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) SoleTenancyPreferences() MigrationCenterPreferenceSetVirtualMachinePreferencesSoleTenancyPreferencesOutputReference {
	var returns MigrationCenterPreferenceSetVirtualMachinePreferencesSoleTenancyPreferencesOutputReference
	_jsii_.Get(
		j,
		"soleTenancyPreferences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) SoleTenancyPreferencesInput() *MigrationCenterPreferenceSetVirtualMachinePreferencesSoleTenancyPreferences {
	var returns *MigrationCenterPreferenceSetVirtualMachinePreferencesSoleTenancyPreferences
	_jsii_.Get(
		j,
		"soleTenancyPreferencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) TargetProduct() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetProduct",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) TargetProductInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetProductInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) VmwareEnginePreferences() MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference {
	var returns MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferencesOutputReference
	_jsii_.Get(
		j,
		"vmwareEnginePreferences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) VmwareEnginePreferencesInput() *MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferences {
	var returns *MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferences
	_jsii_.Get(
		j,
		"vmwareEnginePreferencesInput",
		&returns,
	)
	return returns
}


func NewMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference {
	_init_.Initialize()

	if err := validateNewMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.migrationCenterPreferenceSet.MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference_Override(m MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.migrationCenterPreferenceSet.MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference)SetCommitmentPlan(val *string) {
	if err := j.validateSetCommitmentPlanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commitmentPlan",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference)SetInternalValue(val *MigrationCenterPreferenceSetVirtualMachinePreferences) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference)SetSizingOptimizationStrategy(val *string) {
	if err := j.validateSetSizingOptimizationStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sizingOptimizationStrategy",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference)SetTargetProduct(val *string) {
	if err := j.validateSetTargetProductParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetProduct",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) PutComputeEnginePreferences(value *MigrationCenterPreferenceSetVirtualMachinePreferencesComputeEnginePreferences) {
	if err := m.validatePutComputeEnginePreferencesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putComputeEnginePreferences",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) PutRegionPreferences(value *MigrationCenterPreferenceSetVirtualMachinePreferencesRegionPreferences) {
	if err := m.validatePutRegionPreferencesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putRegionPreferences",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) PutSoleTenancyPreferences(value *MigrationCenterPreferenceSetVirtualMachinePreferencesSoleTenancyPreferences) {
	if err := m.validatePutSoleTenancyPreferencesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putSoleTenancyPreferences",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) PutVmwareEnginePreferences(value *MigrationCenterPreferenceSetVirtualMachinePreferencesVmwareEnginePreferences) {
	if err := m.validatePutVmwareEnginePreferencesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putVmwareEnginePreferences",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) ResetCommitmentPlan() {
	_jsii_.InvokeVoid(
		m,
		"resetCommitmentPlan",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) ResetComputeEnginePreferences() {
	_jsii_.InvokeVoid(
		m,
		"resetComputeEnginePreferences",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) ResetRegionPreferences() {
	_jsii_.InvokeVoid(
		m,
		"resetRegionPreferences",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) ResetSizingOptimizationStrategy() {
	_jsii_.InvokeVoid(
		m,
		"resetSizingOptimizationStrategy",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) ResetSoleTenancyPreferences() {
	_jsii_.InvokeVoid(
		m,
		"resetSoleTenancyPreferences",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) ResetTargetProduct() {
	_jsii_.InvokeVoid(
		m,
		"resetTargetProduct",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) ResetVmwareEnginePreferences() {
	_jsii_.InvokeVoid(
		m,
		"resetVmwareEnginePreferences",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MigrationCenterPreferenceSetVirtualMachinePreferencesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

