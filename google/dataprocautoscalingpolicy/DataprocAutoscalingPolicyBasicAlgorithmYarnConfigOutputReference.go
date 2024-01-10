// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocautoscalingpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v13/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v13/dataprocautoscalingpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference interface {
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
	GracefulDecommissionTimeout() *string
	SetGracefulDecommissionTimeout(val *string)
	GracefulDecommissionTimeoutInput() *string
	InternalValue() *DataprocAutoscalingPolicyBasicAlgorithmYarnConfig
	SetInternalValue(val *DataprocAutoscalingPolicyBasicAlgorithmYarnConfig)
	ScaleDownFactor() *float64
	SetScaleDownFactor(val *float64)
	ScaleDownFactorInput() *float64
	ScaleDownMinWorkerFraction() *float64
	SetScaleDownMinWorkerFraction(val *float64)
	ScaleDownMinWorkerFractionInput() *float64
	ScaleUpFactor() *float64
	SetScaleUpFactor(val *float64)
	ScaleUpFactorInput() *float64
	ScaleUpMinWorkerFraction() *float64
	SetScaleUpMinWorkerFraction(val *float64)
	ScaleUpMinWorkerFractionInput() *float64
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
	ResetScaleDownMinWorkerFraction()
	ResetScaleUpMinWorkerFraction()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference
type jsiiProxy_DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) GracefulDecommissionTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gracefulDecommissionTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) GracefulDecommissionTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gracefulDecommissionTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) InternalValue() *DataprocAutoscalingPolicyBasicAlgorithmYarnConfig {
	var returns *DataprocAutoscalingPolicyBasicAlgorithmYarnConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) ScaleDownFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleDownFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) ScaleDownFactorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleDownFactorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) ScaleDownMinWorkerFraction() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleDownMinWorkerFraction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) ScaleDownMinWorkerFractionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleDownMinWorkerFractionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) ScaleUpFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleUpFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) ScaleUpFactorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleUpFactorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) ScaleUpMinWorkerFraction() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleUpMinWorkerFraction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) ScaleUpMinWorkerFractionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scaleUpMinWorkerFractionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataprocAutoscalingPolicy.DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference_Override(d DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataprocAutoscalingPolicy.DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference)SetGracefulDecommissionTimeout(val *string) {
	if err := j.validateSetGracefulDecommissionTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gracefulDecommissionTimeout",
		val,
	)
}

func (j *jsiiProxy_DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference)SetInternalValue(val *DataprocAutoscalingPolicyBasicAlgorithmYarnConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference)SetScaleDownFactor(val *float64) {
	if err := j.validateSetScaleDownFactorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleDownFactor",
		val,
	)
}

func (j *jsiiProxy_DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference)SetScaleDownMinWorkerFraction(val *float64) {
	if err := j.validateSetScaleDownMinWorkerFractionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleDownMinWorkerFraction",
		val,
	)
}

func (j *jsiiProxy_DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference)SetScaleUpFactor(val *float64) {
	if err := j.validateSetScaleUpFactorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleUpFactor",
		val,
	)
}

func (j *jsiiProxy_DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference)SetScaleUpMinWorkerFraction(val *float64) {
	if err := j.validateSetScaleUpMinWorkerFractionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleUpMinWorkerFraction",
		val,
	)
}

func (j *jsiiProxy_DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) ResetScaleDownMinWorkerFraction() {
	_jsii_.InvokeVoid(
		d,
		"resetScaleDownMinWorkerFraction",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) ResetScaleUpMinWorkerFraction() {
	_jsii_.InvokeVoid(
		d,
		"resetScaleUpMinWorkerFraction",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

