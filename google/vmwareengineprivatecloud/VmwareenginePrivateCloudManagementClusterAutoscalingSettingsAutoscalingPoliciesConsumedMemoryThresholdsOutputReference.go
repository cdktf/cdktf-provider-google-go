// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vmwareengineprivatecloud

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/vmwareengineprivatecloud/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference interface {
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
	InternalValue() *VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholds
	SetInternalValue(val *VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholds)
	ScaleIn() *float64
	SetScaleIn(val *float64)
	ScaleInInput() *float64
	ScaleOut() *float64
	SetScaleOut(val *float64)
	ScaleOutInput() *float64
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference
type jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference) InternalValue() *VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholds {
	var returns *VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholds
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference) ScaleIn() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference) ScaleInInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleInInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference) ScaleOut() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleOut",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference) ScaleOutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleOutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference {
	_init_.Initialize()

	if err := validateNewVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.vmwareenginePrivateCloud.VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference_Override(v VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.vmwareenginePrivateCloud.VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference)SetInternalValue(val *VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholds) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference)SetScaleIn(val *float64) {
	if err := j.validateSetScaleInParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleIn",
		val,
	)
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference)SetScaleOut(val *float64) {
	if err := j.validateSetScaleOutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleOut",
		val,
	)
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesConsumedMemoryThresholdsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

