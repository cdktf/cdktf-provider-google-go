// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vmwareengineprivatecloud

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/vmwareengineprivatecloud/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference interface {
	cdktf.ComplexObject
	AutoscalingPolicies() VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesList
	AutoscalingPoliciesInput() interface{}
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
	CoolDownPeriod() *string
	SetCoolDownPeriod(val *string)
	CoolDownPeriodInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *VmwareenginePrivateCloudManagementClusterAutoscalingSettings
	SetInternalValue(val *VmwareenginePrivateCloudManagementClusterAutoscalingSettings)
	MaxClusterNodeCount() *float64
	SetMaxClusterNodeCount(val *float64)
	MaxClusterNodeCountInput() *float64
	MinClusterNodeCount() *float64
	SetMinClusterNodeCount(val *float64)
	MinClusterNodeCountInput() *float64
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
	PutAutoscalingPolicies(value interface{})
	ResetCoolDownPeriod()
	ResetMaxClusterNodeCount()
	ResetMinClusterNodeCount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference
type jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) AutoscalingPolicies() VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesList {
	var returns VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesList
	_jsii_.Get(
		j,
		"autoscalingPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) AutoscalingPoliciesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoscalingPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) CoolDownPeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coolDownPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) CoolDownPeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coolDownPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) InternalValue() *VmwareenginePrivateCloudManagementClusterAutoscalingSettings {
	var returns *VmwareenginePrivateCloudManagementClusterAutoscalingSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) MaxClusterNodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxClusterNodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) MaxClusterNodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxClusterNodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) MinClusterNodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minClusterNodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) MinClusterNodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minClusterNodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.vmwareenginePrivateCloud.VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference_Override(v VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.vmwareenginePrivateCloud.VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference)SetCoolDownPeriod(val *string) {
	if err := j.validateSetCoolDownPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"coolDownPeriod",
		val,
	)
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference)SetInternalValue(val *VmwareenginePrivateCloudManagementClusterAutoscalingSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference)SetMaxClusterNodeCount(val *float64) {
	if err := j.validateSetMaxClusterNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxClusterNodeCount",
		val,
	)
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference)SetMinClusterNodeCount(val *float64) {
	if err := j.validateSetMinClusterNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minClusterNodeCount",
		val,
	)
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) PutAutoscalingPolicies(value interface{}) {
	if err := v.validatePutAutoscalingPoliciesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putAutoscalingPolicies",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) ResetCoolDownPeriod() {
	_jsii_.InvokeVoid(
		v,
		"resetCoolDownPeriod",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) ResetMaxClusterNodeCount() {
	_jsii_.InvokeVoid(
		v,
		"resetMaxClusterNodeCount",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) ResetMinClusterNodeCount() {
	_jsii_.InvokeVoid(
		v,
		"resetMinClusterNodeCount",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VmwareenginePrivateCloudManagementClusterAutoscalingSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

