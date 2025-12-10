// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocmetastoreservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/dataprocmetastoreservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference interface {
	cdktf.ComplexObject
	AutoscalingEnabled() interface{}
	SetAutoscalingEnabled(val interface{})
	AutoscalingEnabledInput() interface{}
	AutoscalingFactor() *float64
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
	InternalValue() *DataprocMetastoreServiceScalingConfigAutoscalingConfig
	SetInternalValue(val *DataprocMetastoreServiceScalingConfigAutoscalingConfig)
	LimitConfig() DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference
	LimitConfigInput() *DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfig
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
	PutLimitConfig(value *DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfig)
	ResetAutoscalingEnabled()
	ResetLimitConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference
type jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) AutoscalingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoscalingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) AutoscalingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoscalingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) AutoscalingFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoscalingFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) InternalValue() *DataprocMetastoreServiceScalingConfigAutoscalingConfig {
	var returns *DataprocMetastoreServiceScalingConfigAutoscalingConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) LimitConfig() DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference {
	var returns DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfigOutputReference
	_jsii_.Get(
		j,
		"limitConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) LimitConfigInput() *DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfig {
	var returns *DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfig
	_jsii_.Get(
		j,
		"limitConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataprocMetastoreService.DataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference_Override(d DataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataprocMetastoreService.DataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference)SetAutoscalingEnabled(val interface{}) {
	if err := j.validateSetAutoscalingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoscalingEnabled",
		val,
	)
}

func (j *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference)SetInternalValue(val *DataprocMetastoreServiceScalingConfigAutoscalingConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) PutLimitConfig(value *DataprocMetastoreServiceScalingConfigAutoscalingConfigLimitConfig) {
	if err := d.validatePutLimitConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLimitConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) ResetAutoscalingEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetAutoscalingEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) ResetLimitConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetLimitConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocMetastoreServiceScalingConfigAutoscalingConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

