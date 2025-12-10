// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregionbackendservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/computeregionbackendservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeRegionBackendServiceOutlierDetectionOutputReference interface {
	cdktf.ComplexObject
	BaseEjectionTime() ComputeRegionBackendServiceOutlierDetectionBaseEjectionTimeOutputReference
	BaseEjectionTimeInput() *ComputeRegionBackendServiceOutlierDetectionBaseEjectionTime
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
	ConsecutiveErrors() *float64
	SetConsecutiveErrors(val *float64)
	ConsecutiveErrorsInput() *float64
	ConsecutiveGatewayFailure() *float64
	SetConsecutiveGatewayFailure(val *float64)
	ConsecutiveGatewayFailureInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EnforcingConsecutiveErrors() *float64
	SetEnforcingConsecutiveErrors(val *float64)
	EnforcingConsecutiveErrorsInput() *float64
	EnforcingConsecutiveGatewayFailure() *float64
	SetEnforcingConsecutiveGatewayFailure(val *float64)
	EnforcingConsecutiveGatewayFailureInput() *float64
	EnforcingSuccessRate() *float64
	SetEnforcingSuccessRate(val *float64)
	EnforcingSuccessRateInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *ComputeRegionBackendServiceOutlierDetection
	SetInternalValue(val *ComputeRegionBackendServiceOutlierDetection)
	Interval() ComputeRegionBackendServiceOutlierDetectionIntervalOutputReference
	IntervalInput() *ComputeRegionBackendServiceOutlierDetectionInterval
	MaxEjectionPercent() *float64
	SetMaxEjectionPercent(val *float64)
	MaxEjectionPercentInput() *float64
	SuccessRateMinimumHosts() *float64
	SetSuccessRateMinimumHosts(val *float64)
	SuccessRateMinimumHostsInput() *float64
	SuccessRateRequestVolume() *float64
	SetSuccessRateRequestVolume(val *float64)
	SuccessRateRequestVolumeInput() *float64
	SuccessRateStdevFactor() *float64
	SetSuccessRateStdevFactor(val *float64)
	SuccessRateStdevFactorInput() *float64
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutBaseEjectionTime(value *ComputeRegionBackendServiceOutlierDetectionBaseEjectionTime)
	PutInterval(value *ComputeRegionBackendServiceOutlierDetectionInterval)
	ResetBaseEjectionTime()
	ResetConsecutiveErrors()
	ResetConsecutiveGatewayFailure()
	ResetEnforcingConsecutiveErrors()
	ResetEnforcingConsecutiveGatewayFailure()
	ResetEnforcingSuccessRate()
	ResetInterval()
	ResetMaxEjectionPercent()
	ResetSuccessRateMinimumHosts()
	ResetSuccessRateRequestVolume()
	ResetSuccessRateStdevFactor()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeRegionBackendServiceOutlierDetectionOutputReference
type jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) BaseEjectionTime() ComputeRegionBackendServiceOutlierDetectionBaseEjectionTimeOutputReference {
	var returns ComputeRegionBackendServiceOutlierDetectionBaseEjectionTimeOutputReference
	_jsii_.Get(
		j,
		"baseEjectionTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) BaseEjectionTimeInput() *ComputeRegionBackendServiceOutlierDetectionBaseEjectionTime {
	var returns *ComputeRegionBackendServiceOutlierDetectionBaseEjectionTime
	_jsii_.Get(
		j,
		"baseEjectionTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) ConsecutiveErrors() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"consecutiveErrors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) ConsecutiveErrorsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"consecutiveErrorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) ConsecutiveGatewayFailure() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"consecutiveGatewayFailure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) ConsecutiveGatewayFailureInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"consecutiveGatewayFailureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) EnforcingConsecutiveErrors() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"enforcingConsecutiveErrors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) EnforcingConsecutiveErrorsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"enforcingConsecutiveErrorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) EnforcingConsecutiveGatewayFailure() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"enforcingConsecutiveGatewayFailure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) EnforcingConsecutiveGatewayFailureInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"enforcingConsecutiveGatewayFailureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) EnforcingSuccessRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"enforcingSuccessRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) EnforcingSuccessRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"enforcingSuccessRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) InternalValue() *ComputeRegionBackendServiceOutlierDetection {
	var returns *ComputeRegionBackendServiceOutlierDetection
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) Interval() ComputeRegionBackendServiceOutlierDetectionIntervalOutputReference {
	var returns ComputeRegionBackendServiceOutlierDetectionIntervalOutputReference
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) IntervalInput() *ComputeRegionBackendServiceOutlierDetectionInterval {
	var returns *ComputeRegionBackendServiceOutlierDetectionInterval
	_jsii_.Get(
		j,
		"intervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) MaxEjectionPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxEjectionPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) MaxEjectionPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxEjectionPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) SuccessRateMinimumHosts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successRateMinimumHosts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) SuccessRateMinimumHostsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successRateMinimumHostsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) SuccessRateRequestVolume() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successRateRequestVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) SuccessRateRequestVolumeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successRateRequestVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) SuccessRateStdevFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successRateStdevFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) SuccessRateStdevFactorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successRateStdevFactorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeRegionBackendServiceOutlierDetectionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeRegionBackendServiceOutlierDetectionOutputReference {
	_init_.Initialize()

	if err := validateNewComputeRegionBackendServiceOutlierDetectionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionBackendService.ComputeRegionBackendServiceOutlierDetectionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeRegionBackendServiceOutlierDetectionOutputReference_Override(c ComputeRegionBackendServiceOutlierDetectionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionBackendService.ComputeRegionBackendServiceOutlierDetectionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference)SetConsecutiveErrors(val *float64) {
	if err := j.validateSetConsecutiveErrorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consecutiveErrors",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference)SetConsecutiveGatewayFailure(val *float64) {
	if err := j.validateSetConsecutiveGatewayFailureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consecutiveGatewayFailure",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference)SetEnforcingConsecutiveErrors(val *float64) {
	if err := j.validateSetEnforcingConsecutiveErrorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforcingConsecutiveErrors",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference)SetEnforcingConsecutiveGatewayFailure(val *float64) {
	if err := j.validateSetEnforcingConsecutiveGatewayFailureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforcingConsecutiveGatewayFailure",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference)SetEnforcingSuccessRate(val *float64) {
	if err := j.validateSetEnforcingSuccessRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforcingSuccessRate",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference)SetInternalValue(val *ComputeRegionBackendServiceOutlierDetection) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference)SetMaxEjectionPercent(val *float64) {
	if err := j.validateSetMaxEjectionPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxEjectionPercent",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference)SetSuccessRateMinimumHosts(val *float64) {
	if err := j.validateSetSuccessRateMinimumHostsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"successRateMinimumHosts",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference)SetSuccessRateRequestVolume(val *float64) {
	if err := j.validateSetSuccessRateRequestVolumeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"successRateRequestVolume",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference)SetSuccessRateStdevFactor(val *float64) {
	if err := j.validateSetSuccessRateStdevFactorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"successRateStdevFactor",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) PutBaseEjectionTime(value *ComputeRegionBackendServiceOutlierDetectionBaseEjectionTime) {
	if err := c.validatePutBaseEjectionTimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBaseEjectionTime",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) PutInterval(value *ComputeRegionBackendServiceOutlierDetectionInterval) {
	if err := c.validatePutIntervalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putInterval",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) ResetBaseEjectionTime() {
	_jsii_.InvokeVoid(
		c,
		"resetBaseEjectionTime",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) ResetConsecutiveErrors() {
	_jsii_.InvokeVoid(
		c,
		"resetConsecutiveErrors",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) ResetConsecutiveGatewayFailure() {
	_jsii_.InvokeVoid(
		c,
		"resetConsecutiveGatewayFailure",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) ResetEnforcingConsecutiveErrors() {
	_jsii_.InvokeVoid(
		c,
		"resetEnforcingConsecutiveErrors",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) ResetEnforcingConsecutiveGatewayFailure() {
	_jsii_.InvokeVoid(
		c,
		"resetEnforcingConsecutiveGatewayFailure",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) ResetEnforcingSuccessRate() {
	_jsii_.InvokeVoid(
		c,
		"resetEnforcingSuccessRate",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) ResetInterval() {
	_jsii_.InvokeVoid(
		c,
		"resetInterval",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) ResetMaxEjectionPercent() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxEjectionPercent",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) ResetSuccessRateMinimumHosts() {
	_jsii_.InvokeVoid(
		c,
		"resetSuccessRateMinimumHosts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) ResetSuccessRateRequestVolume() {
	_jsii_.InvokeVoid(
		c,
		"resetSuccessRateRequestVolume",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) ResetSuccessRateStdevFactor() {
	_jsii_.InvokeVoid(
		c,
		"resetSuccessRateStdevFactor",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionBackendServiceOutlierDetectionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

