// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregionautoscaler

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v10/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v10/computeregionautoscaler/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeRegionAutoscalerAutoscalingPolicyOutputReference interface {
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
	CooldownPeriod() *float64
	SetCooldownPeriod(val *float64)
	CooldownPeriodInput() *float64
	CpuUtilization() ComputeRegionAutoscalerAutoscalingPolicyCpuUtilizationOutputReference
	CpuUtilizationInput() *ComputeRegionAutoscalerAutoscalingPolicyCpuUtilization
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *ComputeRegionAutoscalerAutoscalingPolicy
	SetInternalValue(val *ComputeRegionAutoscalerAutoscalingPolicy)
	LoadBalancingUtilization() ComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationOutputReference
	LoadBalancingUtilizationInput() *ComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilization
	MaxReplicas() *float64
	SetMaxReplicas(val *float64)
	MaxReplicasInput() *float64
	Metric() ComputeRegionAutoscalerAutoscalingPolicyMetricList
	MetricInput() interface{}
	MinReplicas() *float64
	SetMinReplicas(val *float64)
	MinReplicasInput() *float64
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	ScaleInControl() ComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference
	ScaleInControlInput() *ComputeRegionAutoscalerAutoscalingPolicyScaleInControl
	ScalingSchedules() ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesList
	ScalingSchedulesInput() interface{}
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
	PutCpuUtilization(value *ComputeRegionAutoscalerAutoscalingPolicyCpuUtilization)
	PutLoadBalancingUtilization(value *ComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilization)
	PutMetric(value interface{})
	PutScaleInControl(value *ComputeRegionAutoscalerAutoscalingPolicyScaleInControl)
	PutScalingSchedules(value interface{})
	ResetCooldownPeriod()
	ResetCpuUtilization()
	ResetLoadBalancingUtilization()
	ResetMetric()
	ResetMode()
	ResetScaleInControl()
	ResetScalingSchedules()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeRegionAutoscalerAutoscalingPolicyOutputReference
type jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) CooldownPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cooldownPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) CooldownPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cooldownPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) CpuUtilization() ComputeRegionAutoscalerAutoscalingPolicyCpuUtilizationOutputReference {
	var returns ComputeRegionAutoscalerAutoscalingPolicyCpuUtilizationOutputReference
	_jsii_.Get(
		j,
		"cpuUtilization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) CpuUtilizationInput() *ComputeRegionAutoscalerAutoscalingPolicyCpuUtilization {
	var returns *ComputeRegionAutoscalerAutoscalingPolicyCpuUtilization
	_jsii_.Get(
		j,
		"cpuUtilizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) InternalValue() *ComputeRegionAutoscalerAutoscalingPolicy {
	var returns *ComputeRegionAutoscalerAutoscalingPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) LoadBalancingUtilization() ComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationOutputReference {
	var returns ComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilizationOutputReference
	_jsii_.Get(
		j,
		"loadBalancingUtilization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) LoadBalancingUtilizationInput() *ComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilization {
	var returns *ComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilization
	_jsii_.Get(
		j,
		"loadBalancingUtilizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) MaxReplicas() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxReplicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) MaxReplicasInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxReplicasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) Metric() ComputeRegionAutoscalerAutoscalingPolicyMetricList {
	var returns ComputeRegionAutoscalerAutoscalingPolicyMetricList
	_jsii_.Get(
		j,
		"metric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) MetricInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) MinReplicas() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minReplicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) MinReplicasInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minReplicasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) ScaleInControl() ComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference {
	var returns ComputeRegionAutoscalerAutoscalingPolicyScaleInControlOutputReference
	_jsii_.Get(
		j,
		"scaleInControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) ScaleInControlInput() *ComputeRegionAutoscalerAutoscalingPolicyScaleInControl {
	var returns *ComputeRegionAutoscalerAutoscalingPolicyScaleInControl
	_jsii_.Get(
		j,
		"scaleInControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) ScalingSchedules() ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesList {
	var returns ComputeRegionAutoscalerAutoscalingPolicyScalingSchedulesList
	_jsii_.Get(
		j,
		"scalingSchedules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) ScalingSchedulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scalingSchedulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeRegionAutoscalerAutoscalingPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeRegionAutoscalerAutoscalingPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewComputeRegionAutoscalerAutoscalingPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionAutoscaler.ComputeRegionAutoscalerAutoscalingPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeRegionAutoscalerAutoscalingPolicyOutputReference_Override(c ComputeRegionAutoscalerAutoscalingPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionAutoscaler.ComputeRegionAutoscalerAutoscalingPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference)SetCooldownPeriod(val *float64) {
	if err := j.validateSetCooldownPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cooldownPeriod",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference)SetInternalValue(val *ComputeRegionAutoscalerAutoscalingPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference)SetMaxReplicas(val *float64) {
	if err := j.validateSetMaxReplicasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxReplicas",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference)SetMinReplicas(val *float64) {
	if err := j.validateSetMinReplicasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minReplicas",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) PutCpuUtilization(value *ComputeRegionAutoscalerAutoscalingPolicyCpuUtilization) {
	if err := c.validatePutCpuUtilizationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCpuUtilization",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) PutLoadBalancingUtilization(value *ComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilization) {
	if err := c.validatePutLoadBalancingUtilizationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLoadBalancingUtilization",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) PutMetric(value interface{}) {
	if err := c.validatePutMetricParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMetric",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) PutScaleInControl(value *ComputeRegionAutoscalerAutoscalingPolicyScaleInControl) {
	if err := c.validatePutScaleInControlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putScaleInControl",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) PutScalingSchedules(value interface{}) {
	if err := c.validatePutScalingSchedulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putScalingSchedules",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) ResetCooldownPeriod() {
	_jsii_.InvokeVoid(
		c,
		"resetCooldownPeriod",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) ResetCpuUtilization() {
	_jsii_.InvokeVoid(
		c,
		"resetCpuUtilization",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) ResetLoadBalancingUtilization() {
	_jsii_.InvokeVoid(
		c,
		"resetLoadBalancingUtilization",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) ResetMetric() {
	_jsii_.InvokeVoid(
		c,
		"resetMetric",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) ResetMode() {
	_jsii_.InvokeVoid(
		c,
		"resetMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) ResetScaleInControl() {
	_jsii_.InvokeVoid(
		c,
		"resetScaleInControl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) ResetScalingSchedules() {
	_jsii_.InvokeVoid(
		c,
		"resetScalingSchedules",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeRegionAutoscalerAutoscalingPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

