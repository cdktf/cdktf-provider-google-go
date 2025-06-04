// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagooglebigquerytable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/datagooglebigquerytable/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleBigqueryTableExternalDataConfigurationOutputReference interface {
	cdktf.ComplexObject
	Autodetect() cdktf.IResolvable
	AvroOptions() DataGoogleBigqueryTableExternalDataConfigurationAvroOptionsList
	BigtableOptions() DataGoogleBigqueryTableExternalDataConfigurationBigtableOptionsList
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
	Compression() *string
	ConnectionId() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CsvOptions() DataGoogleBigqueryTableExternalDataConfigurationCsvOptionsList
	FileSetSpecType() *string
	// Experimental.
	Fqn() *string
	GoogleSheetsOptions() DataGoogleBigqueryTableExternalDataConfigurationGoogleSheetsOptionsList
	HivePartitioningOptions() DataGoogleBigqueryTableExternalDataConfigurationHivePartitioningOptionsList
	IgnoreUnknownValues() cdktf.IResolvable
	InternalValue() *DataGoogleBigqueryTableExternalDataConfiguration
	SetInternalValue(val *DataGoogleBigqueryTableExternalDataConfiguration)
	JsonExtension() *string
	JsonOptions() DataGoogleBigqueryTableExternalDataConfigurationJsonOptionsList
	MaxBadRecords() *float64
	MetadataCacheMode() *string
	ObjectMetadata() *string
	ParquetOptions() DataGoogleBigqueryTableExternalDataConfigurationParquetOptionsList
	ReferenceFileSchemaUri() *string
	Schema() *string
	SourceFormat() *string
	SourceUris() *[]*string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataGoogleBigqueryTableExternalDataConfigurationOutputReference
type jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference) Autodetect() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"autodetect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference) AvroOptions() DataGoogleBigqueryTableExternalDataConfigurationAvroOptionsList {
	var returns DataGoogleBigqueryTableExternalDataConfigurationAvroOptionsList
	_jsii_.Get(
		j,
		"avroOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference) BigtableOptions() DataGoogleBigqueryTableExternalDataConfigurationBigtableOptionsList {
	var returns DataGoogleBigqueryTableExternalDataConfigurationBigtableOptionsList
	_jsii_.Get(
		j,
		"bigtableOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference) Compression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference) ConnectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference) CsvOptions() DataGoogleBigqueryTableExternalDataConfigurationCsvOptionsList {
	var returns DataGoogleBigqueryTableExternalDataConfigurationCsvOptionsList
	_jsii_.Get(
		j,
		"csvOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference) FileSetSpecType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSetSpecType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference) GoogleSheetsOptions() DataGoogleBigqueryTableExternalDataConfigurationGoogleSheetsOptionsList {
	var returns DataGoogleBigqueryTableExternalDataConfigurationGoogleSheetsOptionsList
	_jsii_.Get(
		j,
		"googleSheetsOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference) HivePartitioningOptions() DataGoogleBigqueryTableExternalDataConfigurationHivePartitioningOptionsList {
	var returns DataGoogleBigqueryTableExternalDataConfigurationHivePartitioningOptionsList
	_jsii_.Get(
		j,
		"hivePartitioningOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference) IgnoreUnknownValues() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"ignoreUnknownValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference) InternalValue() *DataGoogleBigqueryTableExternalDataConfiguration {
	var returns *DataGoogleBigqueryTableExternalDataConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference) JsonExtension() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jsonExtension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference) JsonOptions() DataGoogleBigqueryTableExternalDataConfigurationJsonOptionsList {
	var returns DataGoogleBigqueryTableExternalDataConfigurationJsonOptionsList
	_jsii_.Get(
		j,
		"jsonOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference) MaxBadRecords() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBadRecords",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference) MetadataCacheMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataCacheMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference) ObjectMetadata() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference) ParquetOptions() DataGoogleBigqueryTableExternalDataConfigurationParquetOptionsList {
	var returns DataGoogleBigqueryTableExternalDataConfigurationParquetOptionsList
	_jsii_.Get(
		j,
		"parquetOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference) ReferenceFileSchemaUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"referenceFileSchemaUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference) Schema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference) SourceFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference) SourceUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sourceUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataGoogleBigqueryTableExternalDataConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataGoogleBigqueryTableExternalDataConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewDataGoogleBigqueryTableExternalDataConfigurationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleBigqueryTable.DataGoogleBigqueryTableExternalDataConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataGoogleBigqueryTableExternalDataConfigurationOutputReference_Override(d DataGoogleBigqueryTableExternalDataConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleBigqueryTable.DataGoogleBigqueryTableExternalDataConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference)SetInternalValue(val *DataGoogleBigqueryTableExternalDataConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataGoogleBigqueryTableExternalDataConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

