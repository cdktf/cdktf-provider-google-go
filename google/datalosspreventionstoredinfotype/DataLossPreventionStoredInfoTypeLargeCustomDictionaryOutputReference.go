// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventionstoredinfotype

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/datalosspreventionstoredinfotype/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference interface {
	cdktf.ComplexObject
	BigQueryField() DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference
	BigQueryFieldInput() *DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryField
	CloudStorageFileSet() DataLossPreventionStoredInfoTypeLargeCustomDictionaryCloudStorageFileSetOutputReference
	CloudStorageFileSetInput() *DataLossPreventionStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet
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
	InternalValue() *DataLossPreventionStoredInfoTypeLargeCustomDictionary
	SetInternalValue(val *DataLossPreventionStoredInfoTypeLargeCustomDictionary)
	OutputPath() DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputPathOutputReference
	OutputPathInput() *DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputPath
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
	PutBigQueryField(value *DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryField)
	PutCloudStorageFileSet(value *DataLossPreventionStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet)
	PutOutputPath(value *DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputPath)
	ResetBigQueryField()
	ResetCloudStorageFileSet()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference
type jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) BigQueryField() DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference {
	var returns DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference
	_jsii_.Get(
		j,
		"bigQueryField",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) BigQueryFieldInput() *DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryField {
	var returns *DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryField
	_jsii_.Get(
		j,
		"bigQueryFieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) CloudStorageFileSet() DataLossPreventionStoredInfoTypeLargeCustomDictionaryCloudStorageFileSetOutputReference {
	var returns DataLossPreventionStoredInfoTypeLargeCustomDictionaryCloudStorageFileSetOutputReference
	_jsii_.Get(
		j,
		"cloudStorageFileSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) CloudStorageFileSetInput() *DataLossPreventionStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet {
	var returns *DataLossPreventionStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet
	_jsii_.Get(
		j,
		"cloudStorageFileSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) InternalValue() *DataLossPreventionStoredInfoTypeLargeCustomDictionary {
	var returns *DataLossPreventionStoredInfoTypeLargeCustomDictionary
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) OutputPath() DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputPathOutputReference {
	var returns DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputPathOutputReference
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) OutputPathInput() *DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputPath {
	var returns *DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputPath
	_jsii_.Get(
		j,
		"outputPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference {
	_init_.Initialize()

	if err := validateNewDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataLossPreventionStoredInfoType.DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference_Override(d DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataLossPreventionStoredInfoType.DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference)SetInternalValue(val *DataLossPreventionStoredInfoTypeLargeCustomDictionary) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) PutBigQueryField(value *DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryField) {
	if err := d.validatePutBigQueryFieldParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putBigQueryField",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) PutCloudStorageFileSet(value *DataLossPreventionStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet) {
	if err := d.validatePutCloudStorageFileSetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCloudStorageFileSet",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) PutOutputPath(value *DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputPath) {
	if err := d.validatePutOutputPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOutputPath",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) ResetBigQueryField() {
	_jsii_.InvokeVoid(
		d,
		"resetBigQueryField",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) ResetCloudStorageFileSet() {
	_jsii_.InvokeVoid(
		d,
		"resetCloudStorageFileSet",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

