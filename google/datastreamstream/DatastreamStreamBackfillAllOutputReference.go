// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datastreamstream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v10/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v10/datastreamstream/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatastreamStreamBackfillAllOutputReference interface {
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
	InternalValue() *DatastreamStreamBackfillAll
	SetInternalValue(val *DatastreamStreamBackfillAll)
	MysqlExcludedObjects() DatastreamStreamBackfillAllMysqlExcludedObjectsOutputReference
	MysqlExcludedObjectsInput() *DatastreamStreamBackfillAllMysqlExcludedObjects
	OracleExcludedObjects() DatastreamStreamBackfillAllOracleExcludedObjectsOutputReference
	OracleExcludedObjectsInput() *DatastreamStreamBackfillAllOracleExcludedObjects
	PostgresqlExcludedObjects() DatastreamStreamBackfillAllPostgresqlExcludedObjectsOutputReference
	PostgresqlExcludedObjectsInput() *DatastreamStreamBackfillAllPostgresqlExcludedObjects
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
	PutMysqlExcludedObjects(value *DatastreamStreamBackfillAllMysqlExcludedObjects)
	PutOracleExcludedObjects(value *DatastreamStreamBackfillAllOracleExcludedObjects)
	PutPostgresqlExcludedObjects(value *DatastreamStreamBackfillAllPostgresqlExcludedObjects)
	ResetMysqlExcludedObjects()
	ResetOracleExcludedObjects()
	ResetPostgresqlExcludedObjects()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatastreamStreamBackfillAllOutputReference
type jsiiProxy_DatastreamStreamBackfillAllOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatastreamStreamBackfillAllOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamBackfillAllOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamBackfillAllOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamBackfillAllOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamBackfillAllOutputReference) InternalValue() *DatastreamStreamBackfillAll {
	var returns *DatastreamStreamBackfillAll
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamBackfillAllOutputReference) MysqlExcludedObjects() DatastreamStreamBackfillAllMysqlExcludedObjectsOutputReference {
	var returns DatastreamStreamBackfillAllMysqlExcludedObjectsOutputReference
	_jsii_.Get(
		j,
		"mysqlExcludedObjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamBackfillAllOutputReference) MysqlExcludedObjectsInput() *DatastreamStreamBackfillAllMysqlExcludedObjects {
	var returns *DatastreamStreamBackfillAllMysqlExcludedObjects
	_jsii_.Get(
		j,
		"mysqlExcludedObjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamBackfillAllOutputReference) OracleExcludedObjects() DatastreamStreamBackfillAllOracleExcludedObjectsOutputReference {
	var returns DatastreamStreamBackfillAllOracleExcludedObjectsOutputReference
	_jsii_.Get(
		j,
		"oracleExcludedObjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamBackfillAllOutputReference) OracleExcludedObjectsInput() *DatastreamStreamBackfillAllOracleExcludedObjects {
	var returns *DatastreamStreamBackfillAllOracleExcludedObjects
	_jsii_.Get(
		j,
		"oracleExcludedObjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamBackfillAllOutputReference) PostgresqlExcludedObjects() DatastreamStreamBackfillAllPostgresqlExcludedObjectsOutputReference {
	var returns DatastreamStreamBackfillAllPostgresqlExcludedObjectsOutputReference
	_jsii_.Get(
		j,
		"postgresqlExcludedObjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamBackfillAllOutputReference) PostgresqlExcludedObjectsInput() *DatastreamStreamBackfillAllPostgresqlExcludedObjects {
	var returns *DatastreamStreamBackfillAllPostgresqlExcludedObjects
	_jsii_.Get(
		j,
		"postgresqlExcludedObjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamBackfillAllOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamBackfillAllOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDatastreamStreamBackfillAllOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatastreamStreamBackfillAllOutputReference {
	_init_.Initialize()

	if err := validateNewDatastreamStreamBackfillAllOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatastreamStreamBackfillAllOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.datastreamStream.DatastreamStreamBackfillAllOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatastreamStreamBackfillAllOutputReference_Override(d DatastreamStreamBackfillAllOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.datastreamStream.DatastreamStreamBackfillAllOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatastreamStreamBackfillAllOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamBackfillAllOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamBackfillAllOutputReference)SetInternalValue(val *DatastreamStreamBackfillAll) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamBackfillAllOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamBackfillAllOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatastreamStreamBackfillAllOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamStreamBackfillAllOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatastreamStreamBackfillAllOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatastreamStreamBackfillAllOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatastreamStreamBackfillAllOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatastreamStreamBackfillAllOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatastreamStreamBackfillAllOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatastreamStreamBackfillAllOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatastreamStreamBackfillAllOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatastreamStreamBackfillAllOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatastreamStreamBackfillAllOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamStreamBackfillAllOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatastreamStreamBackfillAllOutputReference) PutMysqlExcludedObjects(value *DatastreamStreamBackfillAllMysqlExcludedObjects) {
	if err := d.validatePutMysqlExcludedObjectsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMysqlExcludedObjects",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatastreamStreamBackfillAllOutputReference) PutOracleExcludedObjects(value *DatastreamStreamBackfillAllOracleExcludedObjects) {
	if err := d.validatePutOracleExcludedObjectsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOracleExcludedObjects",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatastreamStreamBackfillAllOutputReference) PutPostgresqlExcludedObjects(value *DatastreamStreamBackfillAllPostgresqlExcludedObjects) {
	if err := d.validatePutPostgresqlExcludedObjectsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPostgresqlExcludedObjects",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatastreamStreamBackfillAllOutputReference) ResetMysqlExcludedObjects() {
	_jsii_.InvokeVoid(
		d,
		"resetMysqlExcludedObjects",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamStreamBackfillAllOutputReference) ResetOracleExcludedObjects() {
	_jsii_.InvokeVoid(
		d,
		"resetOracleExcludedObjects",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamStreamBackfillAllOutputReference) ResetPostgresqlExcludedObjects() {
	_jsii_.InvokeVoid(
		d,
		"resetPostgresqlExcludedObjects",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamStreamBackfillAllOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DatastreamStreamBackfillAllOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

