// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appengineflexibleappversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v9/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v9/appengineflexibleappversion/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppEngineFlexibleAppVersionAutomaticScalingOutputReference interface {
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
	CoolDownPeriod() *string
	SetCoolDownPeriod(val *string)
	CoolDownPeriodInput() *string
	CpuUtilization() AppEngineFlexibleAppVersionAutomaticScalingCpuUtilizationOutputReference
	CpuUtilizationInput() *AppEngineFlexibleAppVersionAutomaticScalingCpuUtilization
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DiskUtilization() AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference
	DiskUtilizationInput() *AppEngineFlexibleAppVersionAutomaticScalingDiskUtilization
	// Experimental.
	Fqn() *string
	InternalValue() *AppEngineFlexibleAppVersionAutomaticScaling
	SetInternalValue(val *AppEngineFlexibleAppVersionAutomaticScaling)
	MaxConcurrentRequests() *float64
	SetMaxConcurrentRequests(val *float64)
	MaxConcurrentRequestsInput() *float64
	MaxIdleInstances() *float64
	SetMaxIdleInstances(val *float64)
	MaxIdleInstancesInput() *float64
	MaxPendingLatency() *string
	SetMaxPendingLatency(val *string)
	MaxPendingLatencyInput() *string
	MaxTotalInstances() *float64
	SetMaxTotalInstances(val *float64)
	MaxTotalInstancesInput() *float64
	MinIdleInstances() *float64
	SetMinIdleInstances(val *float64)
	MinIdleInstancesInput() *float64
	MinPendingLatency() *string
	SetMinPendingLatency(val *string)
	MinPendingLatencyInput() *string
	MinTotalInstances() *float64
	SetMinTotalInstances(val *float64)
	MinTotalInstancesInput() *float64
	NetworkUtilization() AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference
	NetworkUtilizationInput() *AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilization
	RequestUtilization() AppEngineFlexibleAppVersionAutomaticScalingRequestUtilizationOutputReference
	RequestUtilizationInput() *AppEngineFlexibleAppVersionAutomaticScalingRequestUtilization
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
	PutCpuUtilization(value *AppEngineFlexibleAppVersionAutomaticScalingCpuUtilization)
	PutDiskUtilization(value *AppEngineFlexibleAppVersionAutomaticScalingDiskUtilization)
	PutNetworkUtilization(value *AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilization)
	PutRequestUtilization(value *AppEngineFlexibleAppVersionAutomaticScalingRequestUtilization)
	ResetCoolDownPeriod()
	ResetDiskUtilization()
	ResetMaxConcurrentRequests()
	ResetMaxIdleInstances()
	ResetMaxPendingLatency()
	ResetMaxTotalInstances()
	ResetMinIdleInstances()
	ResetMinPendingLatency()
	ResetMinTotalInstances()
	ResetNetworkUtilization()
	ResetRequestUtilization()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppEngineFlexibleAppVersionAutomaticScalingOutputReference
type jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) CoolDownPeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coolDownPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) CoolDownPeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coolDownPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) CpuUtilization() AppEngineFlexibleAppVersionAutomaticScalingCpuUtilizationOutputReference {
	var returns AppEngineFlexibleAppVersionAutomaticScalingCpuUtilizationOutputReference
	_jsii_.Get(
		j,
		"cpuUtilization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) CpuUtilizationInput() *AppEngineFlexibleAppVersionAutomaticScalingCpuUtilization {
	var returns *AppEngineFlexibleAppVersionAutomaticScalingCpuUtilization
	_jsii_.Get(
		j,
		"cpuUtilizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) DiskUtilization() AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference {
	var returns AppEngineFlexibleAppVersionAutomaticScalingDiskUtilizationOutputReference
	_jsii_.Get(
		j,
		"diskUtilization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) DiskUtilizationInput() *AppEngineFlexibleAppVersionAutomaticScalingDiskUtilization {
	var returns *AppEngineFlexibleAppVersionAutomaticScalingDiskUtilization
	_jsii_.Get(
		j,
		"diskUtilizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) InternalValue() *AppEngineFlexibleAppVersionAutomaticScaling {
	var returns *AppEngineFlexibleAppVersionAutomaticScaling
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) MaxConcurrentRequests() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentRequests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) MaxConcurrentRequestsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentRequestsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) MaxIdleInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIdleInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) MaxIdleInstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIdleInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) MaxPendingLatency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxPendingLatency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) MaxPendingLatencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxPendingLatencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) MaxTotalInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTotalInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) MaxTotalInstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTotalInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) MinIdleInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minIdleInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) MinIdleInstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minIdleInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) MinPendingLatency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minPendingLatency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) MinPendingLatencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minPendingLatencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) MinTotalInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minTotalInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) MinTotalInstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minTotalInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) NetworkUtilization() AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference {
	var returns AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilizationOutputReference
	_jsii_.Get(
		j,
		"networkUtilization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) NetworkUtilizationInput() *AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilization {
	var returns *AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilization
	_jsii_.Get(
		j,
		"networkUtilizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) RequestUtilization() AppEngineFlexibleAppVersionAutomaticScalingRequestUtilizationOutputReference {
	var returns AppEngineFlexibleAppVersionAutomaticScalingRequestUtilizationOutputReference
	_jsii_.Get(
		j,
		"requestUtilization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) RequestUtilizationInput() *AppEngineFlexibleAppVersionAutomaticScalingRequestUtilization {
	var returns *AppEngineFlexibleAppVersionAutomaticScalingRequestUtilization
	_jsii_.Get(
		j,
		"requestUtilizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAppEngineFlexibleAppVersionAutomaticScalingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppEngineFlexibleAppVersionAutomaticScalingOutputReference {
	_init_.Initialize()

	if err := validateNewAppEngineFlexibleAppVersionAutomaticScalingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.appEngineFlexibleAppVersion.AppEngineFlexibleAppVersionAutomaticScalingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppEngineFlexibleAppVersionAutomaticScalingOutputReference_Override(a AppEngineFlexibleAppVersionAutomaticScalingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.appEngineFlexibleAppVersion.AppEngineFlexibleAppVersionAutomaticScalingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference)SetCoolDownPeriod(val *string) {
	if err := j.validateSetCoolDownPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"coolDownPeriod",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference)SetInternalValue(val *AppEngineFlexibleAppVersionAutomaticScaling) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference)SetMaxConcurrentRequests(val *float64) {
	if err := j.validateSetMaxConcurrentRequestsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrentRequests",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference)SetMaxIdleInstances(val *float64) {
	if err := j.validateSetMaxIdleInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxIdleInstances",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference)SetMaxPendingLatency(val *string) {
	if err := j.validateSetMaxPendingLatencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxPendingLatency",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference)SetMaxTotalInstances(val *float64) {
	if err := j.validateSetMaxTotalInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxTotalInstances",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference)SetMinIdleInstances(val *float64) {
	if err := j.validateSetMinIdleInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minIdleInstances",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference)SetMinPendingLatency(val *string) {
	if err := j.validateSetMinPendingLatencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minPendingLatency",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference)SetMinTotalInstances(val *float64) {
	if err := j.validateSetMinTotalInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minTotalInstances",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) PutCpuUtilization(value *AppEngineFlexibleAppVersionAutomaticScalingCpuUtilization) {
	if err := a.validatePutCpuUtilizationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCpuUtilization",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) PutDiskUtilization(value *AppEngineFlexibleAppVersionAutomaticScalingDiskUtilization) {
	if err := a.validatePutDiskUtilizationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDiskUtilization",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) PutNetworkUtilization(value *AppEngineFlexibleAppVersionAutomaticScalingNetworkUtilization) {
	if err := a.validatePutNetworkUtilizationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNetworkUtilization",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) PutRequestUtilization(value *AppEngineFlexibleAppVersionAutomaticScalingRequestUtilization) {
	if err := a.validatePutRequestUtilizationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRequestUtilization",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) ResetCoolDownPeriod() {
	_jsii_.InvokeVoid(
		a,
		"resetCoolDownPeriod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) ResetDiskUtilization() {
	_jsii_.InvokeVoid(
		a,
		"resetDiskUtilization",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) ResetMaxConcurrentRequests() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxConcurrentRequests",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) ResetMaxIdleInstances() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxIdleInstances",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) ResetMaxPendingLatency() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxPendingLatency",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) ResetMaxTotalInstances() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxTotalInstances",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) ResetMinIdleInstances() {
	_jsii_.InvokeVoid(
		a,
		"resetMinIdleInstances",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) ResetMinPendingLatency() {
	_jsii_.InvokeVoid(
		a,
		"resetMinPendingLatency",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) ResetMinTotalInstances() {
	_jsii_.InvokeVoid(
		a,
		"resetMinTotalInstances",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) ResetNetworkUtilization() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkUtilization",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) ResetRequestUtilization() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestUtilization",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionAutomaticScalingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

