// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventionjobtrigger

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/datalosspreventionjobtrigger/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference interface {
	cdktf.ComplexObject
	BigQueryOptions() DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference
	BigQueryOptionsInput() *DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptions
	CloudStorageOptions() DataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference
	CloudStorageOptionsInput() *DataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptions
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
	DatastoreOptions() DataLossPreventionJobTriggerInspectJobStorageConfigDatastoreOptionsOutputReference
	DatastoreOptionsInput() *DataLossPreventionJobTriggerInspectJobStorageConfigDatastoreOptions
	// Experimental.
	Fqn() *string
	HybridOptions() DataLossPreventionJobTriggerInspectJobStorageConfigHybridOptionsOutputReference
	HybridOptionsInput() *DataLossPreventionJobTriggerInspectJobStorageConfigHybridOptions
	InternalValue() *DataLossPreventionJobTriggerInspectJobStorageConfig
	SetInternalValue(val *DataLossPreventionJobTriggerInspectJobStorageConfig)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimespanConfig() DataLossPreventionJobTriggerInspectJobStorageConfigTimespanConfigOutputReference
	TimespanConfigInput() *DataLossPreventionJobTriggerInspectJobStorageConfigTimespanConfig
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
	PutBigQueryOptions(value *DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptions)
	PutCloudStorageOptions(value *DataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptions)
	PutDatastoreOptions(value *DataLossPreventionJobTriggerInspectJobStorageConfigDatastoreOptions)
	PutHybridOptions(value *DataLossPreventionJobTriggerInspectJobStorageConfigHybridOptions)
	PutTimespanConfig(value *DataLossPreventionJobTriggerInspectJobStorageConfigTimespanConfig)
	ResetBigQueryOptions()
	ResetCloudStorageOptions()
	ResetDatastoreOptions()
	ResetHybridOptions()
	ResetTimespanConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference
type jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) BigQueryOptions() DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference {
	var returns DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptionsOutputReference
	_jsii_.Get(
		j,
		"bigQueryOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) BigQueryOptionsInput() *DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptions {
	var returns *DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptions
	_jsii_.Get(
		j,
		"bigQueryOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) CloudStorageOptions() DataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference {
	var returns DataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsOutputReference
	_jsii_.Get(
		j,
		"cloudStorageOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) CloudStorageOptionsInput() *DataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptions {
	var returns *DataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptions
	_jsii_.Get(
		j,
		"cloudStorageOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) DatastoreOptions() DataLossPreventionJobTriggerInspectJobStorageConfigDatastoreOptionsOutputReference {
	var returns DataLossPreventionJobTriggerInspectJobStorageConfigDatastoreOptionsOutputReference
	_jsii_.Get(
		j,
		"datastoreOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) DatastoreOptionsInput() *DataLossPreventionJobTriggerInspectJobStorageConfigDatastoreOptions {
	var returns *DataLossPreventionJobTriggerInspectJobStorageConfigDatastoreOptions
	_jsii_.Get(
		j,
		"datastoreOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) HybridOptions() DataLossPreventionJobTriggerInspectJobStorageConfigHybridOptionsOutputReference {
	var returns DataLossPreventionJobTriggerInspectJobStorageConfigHybridOptionsOutputReference
	_jsii_.Get(
		j,
		"hybridOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) HybridOptionsInput() *DataLossPreventionJobTriggerInspectJobStorageConfigHybridOptions {
	var returns *DataLossPreventionJobTriggerInspectJobStorageConfigHybridOptions
	_jsii_.Get(
		j,
		"hybridOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) InternalValue() *DataLossPreventionJobTriggerInspectJobStorageConfig {
	var returns *DataLossPreventionJobTriggerInspectJobStorageConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) TimespanConfig() DataLossPreventionJobTriggerInspectJobStorageConfigTimespanConfigOutputReference {
	var returns DataLossPreventionJobTriggerInspectJobStorageConfigTimespanConfigOutputReference
	_jsii_.Get(
		j,
		"timespanConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) TimespanConfigInput() *DataLossPreventionJobTriggerInspectJobStorageConfigTimespanConfig {
	var returns *DataLossPreventionJobTriggerInspectJobStorageConfigTimespanConfig
	_jsii_.Get(
		j,
		"timespanConfigInput",
		&returns,
	)
	return returns
}


func NewDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataLossPreventionJobTriggerInspectJobStorageConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataLossPreventionJobTrigger.DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataLossPreventionJobTriggerInspectJobStorageConfigOutputReference_Override(d DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataLossPreventionJobTrigger.DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference)SetInternalValue(val *DataLossPreventionJobTriggerInspectJobStorageConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) PutBigQueryOptions(value *DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptions) {
	if err := d.validatePutBigQueryOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putBigQueryOptions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) PutCloudStorageOptions(value *DataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptions) {
	if err := d.validatePutCloudStorageOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCloudStorageOptions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) PutDatastoreOptions(value *DataLossPreventionJobTriggerInspectJobStorageConfigDatastoreOptions) {
	if err := d.validatePutDatastoreOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDatastoreOptions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) PutHybridOptions(value *DataLossPreventionJobTriggerInspectJobStorageConfigHybridOptions) {
	if err := d.validatePutHybridOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putHybridOptions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) PutTimespanConfig(value *DataLossPreventionJobTriggerInspectJobStorageConfigTimespanConfig) {
	if err := d.validatePutTimespanConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimespanConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) ResetBigQueryOptions() {
	_jsii_.InvokeVoid(
		d,
		"resetBigQueryOptions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) ResetCloudStorageOptions() {
	_jsii_.InvokeVoid(
		d,
		"resetCloudStorageOptions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) ResetDatastoreOptions() {
	_jsii_.InvokeVoid(
		d,
		"resetDatastoreOptions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) ResetHybridOptions() {
	_jsii_.InvokeVoid(
		d,
		"resetHybridOptions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) ResetTimespanConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetTimespanConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataLossPreventionJobTriggerInspectJobStorageConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

