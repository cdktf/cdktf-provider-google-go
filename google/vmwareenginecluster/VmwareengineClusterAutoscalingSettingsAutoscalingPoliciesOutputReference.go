// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vmwareenginecluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/vmwareenginecluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference interface {
	cdktf.ComplexObject
	AutoscalePolicyId() *string
	SetAutoscalePolicyId(val *string)
	AutoscalePolicyIdInput() *string
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
	ConsumedMemoryThresholds() VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference
	ConsumedMemoryThresholdsInput() *VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholds
	CpuThresholds() VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesCpuThresholdsOutputReference
	CpuThresholdsInput() *VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesCpuThresholds
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NodeTypeId() *string
	SetNodeTypeId(val *string)
	NodeTypeIdInput() *string
	ScaleOutSize() *float64
	SetScaleOutSize(val *float64)
	ScaleOutSizeInput() *float64
	StorageThresholds() VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesStorageThresholdsOutputReference
	StorageThresholdsInput() *VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesStorageThresholds
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
	PutConsumedMemoryThresholds(value *VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholds)
	PutCpuThresholds(value *VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesCpuThresholds)
	PutStorageThresholds(value *VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesStorageThresholds)
	ResetConsumedMemoryThresholds()
	ResetCpuThresholds()
	ResetStorageThresholds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference
type jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) AutoscalePolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoscalePolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) AutoscalePolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoscalePolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) ConsumedMemoryThresholds() VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference {
	var returns VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference
	_jsii_.Get(
		j,
		"consumedMemoryThresholds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) ConsumedMemoryThresholdsInput() *VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholds {
	var returns *VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholds
	_jsii_.Get(
		j,
		"consumedMemoryThresholdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) CpuThresholds() VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesCpuThresholdsOutputReference {
	var returns VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesCpuThresholdsOutputReference
	_jsii_.Get(
		j,
		"cpuThresholds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) CpuThresholdsInput() *VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesCpuThresholds {
	var returns *VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesCpuThresholds
	_jsii_.Get(
		j,
		"cpuThresholdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) NodeTypeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTypeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) NodeTypeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTypeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) ScaleOutSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleOutSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) ScaleOutSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleOutSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) StorageThresholds() VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesStorageThresholdsOutputReference {
	var returns VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesStorageThresholdsOutputReference
	_jsii_.Get(
		j,
		"storageThresholds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) StorageThresholdsInput() *VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesStorageThresholds {
	var returns *VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesStorageThresholds
	_jsii_.Get(
		j,
		"storageThresholdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference {
	_init_.Initialize()

	if err := validateNewVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.vmwareengineCluster.VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewVmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference_Override(v VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.vmwareengineCluster.VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		v,
	)
}

func (j *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference)SetAutoscalePolicyId(val *string) {
	if err := j.validateSetAutoscalePolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoscalePolicyId",
		val,
	)
}

func (j *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference)SetNodeTypeId(val *string) {
	if err := j.validateSetNodeTypeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeTypeId",
		val,
	)
}

func (j *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference)SetScaleOutSize(val *float64) {
	if err := j.validateSetScaleOutSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleOutSize",
		val,
	)
}

func (j *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) PutConsumedMemoryThresholds(value *VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholds) {
	if err := v.validatePutConsumedMemoryThresholdsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putConsumedMemoryThresholds",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) PutCpuThresholds(value *VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesCpuThresholds) {
	if err := v.validatePutCpuThresholdsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putCpuThresholds",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) PutStorageThresholds(value *VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesStorageThresholds) {
	if err := v.validatePutStorageThresholdsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putStorageThresholds",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) ResetConsumedMemoryThresholds() {
	_jsii_.InvokeVoid(
		v,
		"resetConsumedMemoryThresholds",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) ResetCpuThresholds() {
	_jsii_.InvokeVoid(
		v,
		"resetCpuThresholds",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) ResetStorageThresholds() {
	_jsii_.InvokeVoid(
		v,
		"resetStorageThresholds",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := v.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareengineClusterAutoscalingSettingsAutoscalingPoliciesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

