// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/containercluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference interface {
	cdktf.ComplexObject
	AllowedUnsafeSysctls() *[]*string
	SetAllowedUnsafeSysctls(val *[]*string)
	AllowedUnsafeSysctlsInput() *[]*string
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
	ContainerLogMaxFiles() *float64
	SetContainerLogMaxFiles(val *float64)
	ContainerLogMaxFilesInput() *float64
	ContainerLogMaxSize() *string
	SetContainerLogMaxSize(val *string)
	ContainerLogMaxSizeInput() *string
	CpuCfsQuota() interface{}
	SetCpuCfsQuota(val interface{})
	CpuCfsQuotaInput() interface{}
	CpuCfsQuotaPeriod() *string
	SetCpuCfsQuotaPeriod(val *string)
	CpuCfsQuotaPeriodInput() *string
	CpuManagerPolicy() *string
	SetCpuManagerPolicy(val *string)
	CpuManagerPolicyInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	ImageGcHighThresholdPercent() *float64
	SetImageGcHighThresholdPercent(val *float64)
	ImageGcHighThresholdPercentInput() *float64
	ImageGcLowThresholdPercent() *float64
	SetImageGcLowThresholdPercent(val *float64)
	ImageGcLowThresholdPercentInput() *float64
	ImageMaximumGcAge() *string
	SetImageMaximumGcAge(val *string)
	ImageMaximumGcAgeInput() *string
	ImageMinimumGcAge() *string
	SetImageMinimumGcAge(val *string)
	ImageMinimumGcAgeInput() *string
	InsecureKubeletReadonlyPortEnabled() *string
	SetInsecureKubeletReadonlyPortEnabled(val *string)
	InsecureKubeletReadonlyPortEnabledInput() *string
	InternalValue() *ContainerClusterNodePoolNodeConfigKubeletConfig
	SetInternalValue(val *ContainerClusterNodePoolNodeConfigKubeletConfig)
	PodPidsLimit() *float64
	SetPodPidsLimit(val *float64)
	PodPidsLimitInput() *float64
	SingleProcessOomKill() interface{}
	SetSingleProcessOomKill(val interface{})
	SingleProcessOomKillInput() interface{}
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
	ResetAllowedUnsafeSysctls()
	ResetContainerLogMaxFiles()
	ResetContainerLogMaxSize()
	ResetCpuCfsQuota()
	ResetCpuCfsQuotaPeriod()
	ResetCpuManagerPolicy()
	ResetImageGcHighThresholdPercent()
	ResetImageGcLowThresholdPercent()
	ResetImageMaximumGcAge()
	ResetImageMinimumGcAge()
	ResetInsecureKubeletReadonlyPortEnabled()
	ResetPodPidsLimit()
	ResetSingleProcessOomKill()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference
type jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) AllowedUnsafeSysctls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedUnsafeSysctls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) AllowedUnsafeSysctlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedUnsafeSysctlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) ContainerLogMaxFiles() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerLogMaxFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) ContainerLogMaxFilesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerLogMaxFilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) ContainerLogMaxSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerLogMaxSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) ContainerLogMaxSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerLogMaxSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) CpuCfsQuota() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuCfsQuota",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) CpuCfsQuotaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuCfsQuotaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) CpuCfsQuotaPeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuCfsQuotaPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) CpuCfsQuotaPeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuCfsQuotaPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) CpuManagerPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuManagerPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) CpuManagerPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuManagerPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) ImageGcHighThresholdPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"imageGcHighThresholdPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) ImageGcHighThresholdPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"imageGcHighThresholdPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) ImageGcLowThresholdPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"imageGcLowThresholdPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) ImageGcLowThresholdPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"imageGcLowThresholdPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) ImageMaximumGcAge() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageMaximumGcAge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) ImageMaximumGcAgeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageMaximumGcAgeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) ImageMinimumGcAge() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageMinimumGcAge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) ImageMinimumGcAgeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageMinimumGcAgeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) InsecureKubeletReadonlyPortEnabled() *string {
	var returns *string
	_jsii_.Get(
		j,
		"insecureKubeletReadonlyPortEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) InsecureKubeletReadonlyPortEnabledInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"insecureKubeletReadonlyPortEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) InternalValue() *ContainerClusterNodePoolNodeConfigKubeletConfig {
	var returns *ContainerClusterNodePoolNodeConfigKubeletConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) PodPidsLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"podPidsLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) PodPidsLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"podPidsLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) SingleProcessOomKill() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"singleProcessOomKill",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) SingleProcessOomKillInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"singleProcessOomKillInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewContainerClusterNodePoolNodeConfigKubeletConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference {
	_init_.Initialize()

	if err := validateNewContainerClusterNodePoolNodeConfigKubeletConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.containerCluster.ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerClusterNodePoolNodeConfigKubeletConfigOutputReference_Override(c ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.containerCluster.ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference)SetAllowedUnsafeSysctls(val *[]*string) {
	if err := j.validateSetAllowedUnsafeSysctlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedUnsafeSysctls",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference)SetContainerLogMaxFiles(val *float64) {
	if err := j.validateSetContainerLogMaxFilesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerLogMaxFiles",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference)SetContainerLogMaxSize(val *string) {
	if err := j.validateSetContainerLogMaxSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerLogMaxSize",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference)SetCpuCfsQuota(val interface{}) {
	if err := j.validateSetCpuCfsQuotaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuCfsQuota",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference)SetCpuCfsQuotaPeriod(val *string) {
	if err := j.validateSetCpuCfsQuotaPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuCfsQuotaPeriod",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference)SetCpuManagerPolicy(val *string) {
	if err := j.validateSetCpuManagerPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuManagerPolicy",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference)SetImageGcHighThresholdPercent(val *float64) {
	if err := j.validateSetImageGcHighThresholdPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageGcHighThresholdPercent",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference)SetImageGcLowThresholdPercent(val *float64) {
	if err := j.validateSetImageGcLowThresholdPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageGcLowThresholdPercent",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference)SetImageMaximumGcAge(val *string) {
	if err := j.validateSetImageMaximumGcAgeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageMaximumGcAge",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference)SetImageMinimumGcAge(val *string) {
	if err := j.validateSetImageMinimumGcAgeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageMinimumGcAge",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference)SetInsecureKubeletReadonlyPortEnabled(val *string) {
	if err := j.validateSetInsecureKubeletReadonlyPortEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insecureKubeletReadonlyPortEnabled",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference)SetInternalValue(val *ContainerClusterNodePoolNodeConfigKubeletConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference)SetPodPidsLimit(val *float64) {
	if err := j.validateSetPodPidsLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"podPidsLimit",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference)SetSingleProcessOomKill(val interface{}) {
	if err := j.validateSetSingleProcessOomKillParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"singleProcessOomKill",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) ResetAllowedUnsafeSysctls() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowedUnsafeSysctls",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) ResetContainerLogMaxFiles() {
	_jsii_.InvokeVoid(
		c,
		"resetContainerLogMaxFiles",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) ResetContainerLogMaxSize() {
	_jsii_.InvokeVoid(
		c,
		"resetContainerLogMaxSize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) ResetCpuCfsQuota() {
	_jsii_.InvokeVoid(
		c,
		"resetCpuCfsQuota",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) ResetCpuCfsQuotaPeriod() {
	_jsii_.InvokeVoid(
		c,
		"resetCpuCfsQuotaPeriod",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) ResetCpuManagerPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetCpuManagerPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) ResetImageGcHighThresholdPercent() {
	_jsii_.InvokeVoid(
		c,
		"resetImageGcHighThresholdPercent",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) ResetImageGcLowThresholdPercent() {
	_jsii_.InvokeVoid(
		c,
		"resetImageGcLowThresholdPercent",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) ResetImageMaximumGcAge() {
	_jsii_.InvokeVoid(
		c,
		"resetImageMaximumGcAge",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) ResetImageMinimumGcAge() {
	_jsii_.InvokeVoid(
		c,
		"resetImageMinimumGcAge",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) ResetInsecureKubeletReadonlyPortEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetInsecureKubeletReadonlyPortEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) ResetPodPidsLimit() {
	_jsii_.InvokeVoid(
		c,
		"resetPodPidsLimit",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) ResetSingleProcessOomKill() {
	_jsii_.InvokeVoid(
		c,
		"resetSingleProcessOomKill",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigKubeletConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

