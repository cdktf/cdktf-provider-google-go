// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datastreamstream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/datastreamstream/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference interface {
	cdktf.ComplexObject
	Column() *string
	SetColumn(val *string)
	ColumnInput() *string
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
	DataType() *string
	SetDataType(val *string)
	DataTypeInput() *string
	Encoding() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Length() *float64
	Nullable() cdktf.IResolvable
	OrdinalPosition() *float64
	Precision() *float64
	PrimaryKey() cdktf.IResolvable
	Scale() *float64
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
	ResetColumn()
	ResetDataType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference
type jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference) Column() *string {
	var returns *string
	_jsii_.Get(
		j,
		"column",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference) ColumnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"columnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference) DataType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference) DataTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference) Encoding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encoding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference) Length() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"length",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference) Nullable() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"nullable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference) OrdinalPosition() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ordinalPosition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference) Precision() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"precision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference) PrimaryKey() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"primaryKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference) Scale() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference {
	_init_.Initialize()

	if err := validateNewDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.datastreamStream.DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference_Override(d DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.datastreamStream.DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference)SetColumn(val *string) {
	if err := j.validateSetColumnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"column",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference)SetDataType(val *string) {
	if err := j.validateSetDataTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataType",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference) ResetColumn() {
	_jsii_.InvokeVoid(
		d,
		"resetColumn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference) ResetDataType() {
	_jsii_.InvokeVoid(
		d,
		"resetDataType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

