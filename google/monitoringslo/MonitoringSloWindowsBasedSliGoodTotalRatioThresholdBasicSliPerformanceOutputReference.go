// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoringslo

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/monitoringslo/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference interface {
	cdktf.ComplexObject
	Availability() MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceAvailabilityOutputReference
	AvailabilityInput() *MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceAvailability
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
	InternalValue() *MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformance
	SetInternalValue(val *MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformance)
	Latency() MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceLatencyOutputReference
	LatencyInput() *MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceLatency
	Location() *[]*string
	SetLocation(val *[]*string)
	LocationInput() *[]*string
	Method() *[]*string
	SetMethod(val *[]*string)
	MethodInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Version() *[]*string
	SetVersion(val *[]*string)
	VersionInput() *[]*string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutAvailability(value *MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceAvailability)
	PutLatency(value *MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceLatency)
	ResetAvailability()
	ResetLatency()
	ResetLocation()
	ResetMethod()
	ResetVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference
type jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) Availability() MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceAvailabilityOutputReference {
	var returns MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceAvailabilityOutputReference
	_jsii_.Get(
		j,
		"availability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) AvailabilityInput() *MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceAvailability {
	var returns *MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceAvailability
	_jsii_.Get(
		j,
		"availabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) InternalValue() *MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformance {
	var returns *MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformance
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) Latency() MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceLatencyOutputReference {
	var returns MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceLatencyOutputReference
	_jsii_.Get(
		j,
		"latency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) LatencyInput() *MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceLatency {
	var returns *MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceLatency
	_jsii_.Get(
		j,
		"latencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) Location() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) LocationInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) Method() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) MethodInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"methodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) Version() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) VersionInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


func NewMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference {
	_init_.Initialize()

	if err := validateNewMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.monitoringSlo.MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference_Override(m MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.monitoringSlo.MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference)SetInternalValue(val *MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformance) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference)SetLocation(val *[]*string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference)SetMethod(val *[]*string) {
	if err := j.validateSetMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"method",
		val,
	)
}

func (j *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference)SetVersion(val *[]*string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (m *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) PutAvailability(value *MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceAvailability) {
	if err := m.validatePutAvailabilityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAvailability",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) PutLatency(value *MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceLatency) {
	if err := m.validatePutLatencyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putLatency",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) ResetAvailability() {
	_jsii_.InvokeVoid(
		m,
		"resetAvailability",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) ResetLatency() {
	_jsii_.InvokeVoid(
		m,
		"resetLatency",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) ResetLocation() {
	_jsii_.InvokeVoid(
		m,
		"resetLocation",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) ResetMethod() {
	_jsii_.InvokeVoid(
		m,
		"resetMethod",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) ResetVersion() {
	_jsii_.InvokeVoid(
		m,
		"resetVersion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringSloWindowsBasedSliGoodTotalRatioThresholdBasicSliPerformanceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

