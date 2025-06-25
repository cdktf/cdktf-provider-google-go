// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocsessiontemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/dataprocsessiontemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference interface {
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
	InternalValue() *DataprocSessionTemplateEnvironmentConfigPeripheralsConfig
	SetInternalValue(val *DataprocSessionTemplateEnvironmentConfigPeripheralsConfig)
	MetastoreService() *string
	SetMetastoreService(val *string)
	MetastoreServiceInput() *string
	SparkHistoryServerConfig() DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference
	SparkHistoryServerConfigInput() *DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfig
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
	PutSparkHistoryServerConfig(value *DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfig)
	ResetMetastoreService()
	ResetSparkHistoryServerConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference
type jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) InternalValue() *DataprocSessionTemplateEnvironmentConfigPeripheralsConfig {
	var returns *DataprocSessionTemplateEnvironmentConfigPeripheralsConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) MetastoreService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metastoreService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) MetastoreServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metastoreServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) SparkHistoryServerConfig() DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference {
	var returns DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference
	_jsii_.Get(
		j,
		"sparkHistoryServerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) SparkHistoryServerConfigInput() *DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfig {
	var returns *DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfig
	_jsii_.Get(
		j,
		"sparkHistoryServerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataprocSessionTemplate.DataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference_Override(d DataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataprocSessionTemplate.DataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference)SetInternalValue(val *DataprocSessionTemplateEnvironmentConfigPeripheralsConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference)SetMetastoreService(val *string) {
	if err := j.validateSetMetastoreServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metastoreService",
		val,
	)
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) PutSparkHistoryServerConfig(value *DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfig) {
	if err := d.validatePutSparkHistoryServerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSparkHistoryServerConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) ResetMetastoreService() {
	_jsii_.InvokeVoid(
		d,
		"resetMetastoreService",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) ResetSparkHistoryServerConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetSparkHistoryServerConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

