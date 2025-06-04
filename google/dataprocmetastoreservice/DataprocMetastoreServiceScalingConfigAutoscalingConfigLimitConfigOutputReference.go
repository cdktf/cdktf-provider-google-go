// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocmetastoreservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/dataprocmetastoreservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference interface {
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
	InternalValue() *DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfig
	SetInternalValue(val *DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfig)
	MaxScalingFactor() *float64
	SetMaxScalingFactor(val *float64)
	MaxScalingFactorInput() *float64
	MinScalingFactor() *float64
	SetMinScalingFactor(val *float64)
	MinScalingFactorInput() *float64
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
	ResetMaxScalingFactor()
	ResetMinScalingFactor()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference
type jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) InternalValue() *DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfig {
	var returns *DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) MaxScalingFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxScalingFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) MaxScalingFactorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxScalingFactorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) MinScalingFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minScalingFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) MinScalingFactorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minScalingFactorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataprocMetastoreService.DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference_Override(d DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataprocMetastoreService.DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference)SetInternalValue(val *DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference)SetMaxScalingFactor(val *float64) {
	if err := j.validateSetMaxScalingFactorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxScalingFactor",
		val,
	)
}

func (j *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference)SetMinScalingFactor(val *float64) {
	if err := j.validateSetMinScalingFactorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minScalingFactor",
		val,
	)
}

func (j *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) ResetMaxScalingFactor() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxScalingFactor",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) ResetMinScalingFactor() {
	_jsii_.InvokeVoid(
		d,
		"resetMinScalingFactor",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

