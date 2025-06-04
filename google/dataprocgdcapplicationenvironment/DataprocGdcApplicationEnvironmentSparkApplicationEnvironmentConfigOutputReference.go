// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocgdcapplicationenvironment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/dataprocgdcapplicationenvironment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigOutputReference interface {
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
	DefaultProperties() *map[string]*string
	SetDefaultProperties(val *map[string]*string)
	DefaultPropertiesInput() *map[string]*string
	DefaultVersion() *string
	SetDefaultVersion(val *string)
	DefaultVersionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfig
	SetInternalValue(val *DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfig)
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
	ResetDefaultProperties()
	ResetDefaultVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigOutputReference
type jsiiProxy_DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigOutputReference) DefaultProperties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"defaultProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigOutputReference) DefaultPropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"defaultPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigOutputReference) DefaultVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigOutputReference) DefaultVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigOutputReference) InternalValue() *DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfig {
	var returns *DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataprocGdcApplicationEnvironment.DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigOutputReference_Override(d DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataprocGdcApplicationEnvironment.DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigOutputReference)SetDefaultProperties(val *map[string]*string) {
	if err := j.validateSetDefaultPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultProperties",
		val,
	)
}

func (j *jsiiProxy_DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigOutputReference)SetDefaultVersion(val *string) {
	if err := j.validateSetDefaultVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultVersion",
		val,
	)
}

func (j *jsiiProxy_DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigOutputReference)SetInternalValue(val *DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigOutputReference) ResetDefaultProperties() {
	_jsii_.InvokeVoid(
		d,
		"resetDefaultProperties",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigOutputReference) ResetDefaultVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetDefaultVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataprocGdcApplicationEnvironmentSparkApplicationEnvironmentConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

