package computeautoscaler

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v4/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v4/computeautoscaler/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeAutoscalerAutoscalingPolicyOutputReference interface {
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
	CpuUtilization() ComputeAutoscalerAutoscalingPolicyCpuUtilizationOutputReference
	CpuUtilizationInput() *ComputeAutoscalerAutoscalingPolicyCpuUtilization
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *ComputeAutoscalerAutoscalingPolicy
	SetInternalValue(val *ComputeAutoscalerAutoscalingPolicy)
	LoadBalancingUtilization() ComputeAutoscalerAutoscalingPolicyLoadBalancingUtilizationOutputReference
	LoadBalancingUtilizationInput() *ComputeAutoscalerAutoscalingPolicyLoadBalancingUtilization
	MaxReplicas() *float64
	SetMaxReplicas(val *float64)
	MaxReplicasInput() *float64
	Metric() ComputeAutoscalerAutoscalingPolicyMetricList
	MetricInput() interface{}
	MinReplicas() *float64
	SetMinReplicas(val *float64)
	MinReplicasInput() *float64
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	ScaleInControl() ComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference
	ScaleInControlInput() *ComputeAutoscalerAutoscalingPolicyScaleInControl
	ScalingSchedules() ComputeAutoscalerAutoscalingPolicyScalingSchedulesList
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
	PutCpuUtilization(value *ComputeAutoscalerAutoscalingPolicyCpuUtilization)
	PutLoadBalancingUtilization(value *ComputeAutoscalerAutoscalingPolicyLoadBalancingUtilization)
	PutMetric(value interface{})
	PutScaleInControl(value *ComputeAutoscalerAutoscalingPolicyScaleInControl)
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

// The jsii proxy struct for ComputeAutoscalerAutoscalingPolicyOutputReference
type jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) CooldownPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cooldownPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) CooldownPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cooldownPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) CpuUtilization() ComputeAutoscalerAutoscalingPolicyCpuUtilizationOutputReference {
	var returns ComputeAutoscalerAutoscalingPolicyCpuUtilizationOutputReference
	_jsii_.Get(
		j,
		"cpuUtilization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) CpuUtilizationInput() *ComputeAutoscalerAutoscalingPolicyCpuUtilization {
	var returns *ComputeAutoscalerAutoscalingPolicyCpuUtilization
	_jsii_.Get(
		j,
		"cpuUtilizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) InternalValue() *ComputeAutoscalerAutoscalingPolicy {
	var returns *ComputeAutoscalerAutoscalingPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) LoadBalancingUtilization() ComputeAutoscalerAutoscalingPolicyLoadBalancingUtilizationOutputReference {
	var returns ComputeAutoscalerAutoscalingPolicyLoadBalancingUtilizationOutputReference
	_jsii_.Get(
		j,
		"loadBalancingUtilization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) LoadBalancingUtilizationInput() *ComputeAutoscalerAutoscalingPolicyLoadBalancingUtilization {
	var returns *ComputeAutoscalerAutoscalingPolicyLoadBalancingUtilization
	_jsii_.Get(
		j,
		"loadBalancingUtilizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) MaxReplicas() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxReplicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) MaxReplicasInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxReplicasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) Metric() ComputeAutoscalerAutoscalingPolicyMetricList {
	var returns ComputeAutoscalerAutoscalingPolicyMetricList
	_jsii_.Get(
		j,
		"metric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) MetricInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) MinReplicas() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minReplicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) MinReplicasInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minReplicasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) ScaleInControl() ComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference {
	var returns ComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference
	_jsii_.Get(
		j,
		"scaleInControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) ScaleInControlInput() *ComputeAutoscalerAutoscalingPolicyScaleInControl {
	var returns *ComputeAutoscalerAutoscalingPolicyScaleInControl
	_jsii_.Get(
		j,
		"scaleInControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) ScalingSchedules() ComputeAutoscalerAutoscalingPolicyScalingSchedulesList {
	var returns ComputeAutoscalerAutoscalingPolicyScalingSchedulesList
	_jsii_.Get(
		j,
		"scalingSchedules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) ScalingSchedulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scalingSchedulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeAutoscalerAutoscalingPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeAutoscalerAutoscalingPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewComputeAutoscalerAutoscalingPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeAutoscaler.ComputeAutoscalerAutoscalingPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeAutoscalerAutoscalingPolicyOutputReference_Override(c ComputeAutoscalerAutoscalingPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeAutoscaler.ComputeAutoscalerAutoscalingPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference)SetCooldownPeriod(val *float64) {
	if err := j.validateSetCooldownPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cooldownPeriod",
		val,
	)
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference)SetInternalValue(val *ComputeAutoscalerAutoscalingPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference)SetMaxReplicas(val *float64) {
	if err := j.validateSetMaxReplicasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxReplicas",
		val,
	)
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference)SetMinReplicas(val *float64) {
	if err := j.validateSetMinReplicasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minReplicas",
		val,
	)
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) PutCpuUtilization(value *ComputeAutoscalerAutoscalingPolicyCpuUtilization) {
	if err := c.validatePutCpuUtilizationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCpuUtilization",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) PutLoadBalancingUtilization(value *ComputeAutoscalerAutoscalingPolicyLoadBalancingUtilization) {
	if err := c.validatePutLoadBalancingUtilizationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLoadBalancingUtilization",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) PutMetric(value interface{}) {
	if err := c.validatePutMetricParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMetric",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) PutScaleInControl(value *ComputeAutoscalerAutoscalingPolicyScaleInControl) {
	if err := c.validatePutScaleInControlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putScaleInControl",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) PutScalingSchedules(value interface{}) {
	if err := c.validatePutScalingSchedulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putScalingSchedules",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) ResetCooldownPeriod() {
	_jsii_.InvokeVoid(
		c,
		"resetCooldownPeriod",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) ResetCpuUtilization() {
	_jsii_.InvokeVoid(
		c,
		"resetCpuUtilization",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) ResetLoadBalancingUtilization() {
	_jsii_.InvokeVoid(
		c,
		"resetLoadBalancingUtilization",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) ResetMetric() {
	_jsii_.InvokeVoid(
		c,
		"resetMetric",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) ResetMode() {
	_jsii_.InvokeVoid(
		c,
		"resetMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) ResetScaleInControl() {
	_jsii_.InvokeVoid(
		c,
		"resetScaleInControl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) ResetScalingSchedules() {
	_jsii_.InvokeVoid(
		c,
		"resetScalingSchedules",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeAutoscalerAutoscalingPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

