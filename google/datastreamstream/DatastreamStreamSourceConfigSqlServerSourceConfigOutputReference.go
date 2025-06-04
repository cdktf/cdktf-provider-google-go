// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datastreamstream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/datastreamstream/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference interface {
	cdktf.ComplexObject
	ChangeTables() DatastreamStreamSourceConfigSqlServerSourceConfigChangeTablesOutputReference
	ChangeTablesInput() *DatastreamStreamSourceConfigSqlServerSourceConfigChangeTables
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
	ExcludeObjects() DatastreamStreamSourceConfigSqlServerSourceConfigExcludeObjectsOutputReference
	ExcludeObjectsInput() *DatastreamStreamSourceConfigSqlServerSourceConfigExcludeObjects
	// Experimental.
	Fqn() *string
	IncludeObjects() DatastreamStreamSourceConfigSqlServerSourceConfigIncludeObjectsOutputReference
	IncludeObjectsInput() *DatastreamStreamSourceConfigSqlServerSourceConfigIncludeObjects
	InternalValue() *DatastreamStreamSourceConfigSqlServerSourceConfig
	SetInternalValue(val *DatastreamStreamSourceConfigSqlServerSourceConfig)
	MaxConcurrentBackfillTasks() *float64
	SetMaxConcurrentBackfillTasks(val *float64)
	MaxConcurrentBackfillTasksInput() *float64
	MaxConcurrentCdcTasks() *float64
	SetMaxConcurrentCdcTasks(val *float64)
	MaxConcurrentCdcTasksInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TransactionLogs() DatastreamStreamSourceConfigSqlServerSourceConfigTransactionLogsOutputReference
	TransactionLogsInput() *DatastreamStreamSourceConfigSqlServerSourceConfigTransactionLogs
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
	PutChangeTables(value *DatastreamStreamSourceConfigSqlServerSourceConfigChangeTables)
	PutExcludeObjects(value *DatastreamStreamSourceConfigSqlServerSourceConfigExcludeObjects)
	PutIncludeObjects(value *DatastreamStreamSourceConfigSqlServerSourceConfigIncludeObjects)
	PutTransactionLogs(value *DatastreamStreamSourceConfigSqlServerSourceConfigTransactionLogs)
	ResetChangeTables()
	ResetExcludeObjects()
	ResetIncludeObjects()
	ResetMaxConcurrentBackfillTasks()
	ResetMaxConcurrentCdcTasks()
	ResetTransactionLogs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference
type jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference) ChangeTables() DatastreamStreamSourceConfigSqlServerSourceConfigChangeTablesOutputReference {
	var returns DatastreamStreamSourceConfigSqlServerSourceConfigChangeTablesOutputReference
	_jsii_.Get(
		j,
		"changeTables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference) ChangeTablesInput() *DatastreamStreamSourceConfigSqlServerSourceConfigChangeTables {
	var returns *DatastreamStreamSourceConfigSqlServerSourceConfigChangeTables
	_jsii_.Get(
		j,
		"changeTablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference) ExcludeObjects() DatastreamStreamSourceConfigSqlServerSourceConfigExcludeObjectsOutputReference {
	var returns DatastreamStreamSourceConfigSqlServerSourceConfigExcludeObjectsOutputReference
	_jsii_.Get(
		j,
		"excludeObjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference) ExcludeObjectsInput() *DatastreamStreamSourceConfigSqlServerSourceConfigExcludeObjects {
	var returns *DatastreamStreamSourceConfigSqlServerSourceConfigExcludeObjects
	_jsii_.Get(
		j,
		"excludeObjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference) IncludeObjects() DatastreamStreamSourceConfigSqlServerSourceConfigIncludeObjectsOutputReference {
	var returns DatastreamStreamSourceConfigSqlServerSourceConfigIncludeObjectsOutputReference
	_jsii_.Get(
		j,
		"includeObjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference) IncludeObjectsInput() *DatastreamStreamSourceConfigSqlServerSourceConfigIncludeObjects {
	var returns *DatastreamStreamSourceConfigSqlServerSourceConfigIncludeObjects
	_jsii_.Get(
		j,
		"includeObjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference) InternalValue() *DatastreamStreamSourceConfigSqlServerSourceConfig {
	var returns *DatastreamStreamSourceConfigSqlServerSourceConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference) MaxConcurrentBackfillTasks() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentBackfillTasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference) MaxConcurrentBackfillTasksInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentBackfillTasksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference) MaxConcurrentCdcTasks() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentCdcTasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference) MaxConcurrentCdcTasksInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentCdcTasksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference) TransactionLogs() DatastreamStreamSourceConfigSqlServerSourceConfigTransactionLogsOutputReference {
	var returns DatastreamStreamSourceConfigSqlServerSourceConfigTransactionLogsOutputReference
	_jsii_.Get(
		j,
		"transactionLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference) TransactionLogsInput() *DatastreamStreamSourceConfigSqlServerSourceConfigTransactionLogs {
	var returns *DatastreamStreamSourceConfigSqlServerSourceConfigTransactionLogs
	_jsii_.Get(
		j,
		"transactionLogsInput",
		&returns,
	)
	return returns
}


func NewDatastreamStreamSourceConfigSqlServerSourceConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDatastreamStreamSourceConfigSqlServerSourceConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.datastreamStream.DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatastreamStreamSourceConfigSqlServerSourceConfigOutputReference_Override(d DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.datastreamStream.DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference)SetInternalValue(val *DatastreamStreamSourceConfigSqlServerSourceConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference)SetMaxConcurrentBackfillTasks(val *float64) {
	if err := j.validateSetMaxConcurrentBackfillTasksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrentBackfillTasks",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference)SetMaxConcurrentCdcTasks(val *float64) {
	if err := j.validateSetMaxConcurrentCdcTasksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrentCdcTasks",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference) PutChangeTables(value *DatastreamStreamSourceConfigSqlServerSourceConfigChangeTables) {
	if err := d.validatePutChangeTablesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putChangeTables",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference) PutExcludeObjects(value *DatastreamStreamSourceConfigSqlServerSourceConfigExcludeObjects) {
	if err := d.validatePutExcludeObjectsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putExcludeObjects",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference) PutIncludeObjects(value *DatastreamStreamSourceConfigSqlServerSourceConfigIncludeObjects) {
	if err := d.validatePutIncludeObjectsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putIncludeObjects",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference) PutTransactionLogs(value *DatastreamStreamSourceConfigSqlServerSourceConfigTransactionLogs) {
	if err := d.validatePutTransactionLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTransactionLogs",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference) ResetChangeTables() {
	_jsii_.InvokeVoid(
		d,
		"resetChangeTables",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference) ResetExcludeObjects() {
	_jsii_.InvokeVoid(
		d,
		"resetExcludeObjects",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference) ResetIncludeObjects() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludeObjects",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference) ResetMaxConcurrentBackfillTasks() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxConcurrentBackfillTasks",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference) ResetMaxConcurrentCdcTasks() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxConcurrentCdcTasks",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference) ResetTransactionLogs() {
	_jsii_.InvokeVoid(
		d,
		"resetTransactionLogs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigSqlServerSourceConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

