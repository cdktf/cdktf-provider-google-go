// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigqueryjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v13/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v13/bigqueryjob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BigqueryJobExtractOutputReference interface {
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
	Compression() *string
	SetCompression(val *string)
	CompressionInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DestinationFormat() *string
	SetDestinationFormat(val *string)
	DestinationFormatInput() *string
	DestinationUris() *[]*string
	SetDestinationUris(val *[]*string)
	DestinationUrisInput() *[]*string
	FieldDelimiter() *string
	SetFieldDelimiter(val *string)
	FieldDelimiterInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *BigqueryJobExtract
	SetInternalValue(val *BigqueryJobExtract)
	PrintHeader() interface{}
	SetPrintHeader(val interface{})
	PrintHeaderInput() interface{}
	SourceModel() BigqueryJobExtractSourceModelOutputReference
	SourceModelInput() *BigqueryJobExtractSourceModel
	SourceTable() BigqueryJobExtractSourceTableOutputReference
	SourceTableInput() *BigqueryJobExtractSourceTable
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UseAvroLogicalTypes() interface{}
	SetUseAvroLogicalTypes(val interface{})
	UseAvroLogicalTypesInput() interface{}
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
	PutSourceModel(value *BigqueryJobExtractSourceModel)
	PutSourceTable(value *BigqueryJobExtractSourceTable)
	ResetCompression()
	ResetDestinationFormat()
	ResetFieldDelimiter()
	ResetPrintHeader()
	ResetSourceModel()
	ResetSourceTable()
	ResetUseAvroLogicalTypes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BigqueryJobExtractOutputReference
type jsiiProxy_BigqueryJobExtractOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BigqueryJobExtractOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobExtractOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobExtractOutputReference) Compression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobExtractOutputReference) CompressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobExtractOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobExtractOutputReference) DestinationFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobExtractOutputReference) DestinationFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobExtractOutputReference) DestinationUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobExtractOutputReference) DestinationUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobExtractOutputReference) FieldDelimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobExtractOutputReference) FieldDelimiterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldDelimiterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobExtractOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobExtractOutputReference) InternalValue() *BigqueryJobExtract {
	var returns *BigqueryJobExtract
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobExtractOutputReference) PrintHeader() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"printHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobExtractOutputReference) PrintHeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"printHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobExtractOutputReference) SourceModel() BigqueryJobExtractSourceModelOutputReference {
	var returns BigqueryJobExtractSourceModelOutputReference
	_jsii_.Get(
		j,
		"sourceModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobExtractOutputReference) SourceModelInput() *BigqueryJobExtractSourceModel {
	var returns *BigqueryJobExtractSourceModel
	_jsii_.Get(
		j,
		"sourceModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobExtractOutputReference) SourceTable() BigqueryJobExtractSourceTableOutputReference {
	var returns BigqueryJobExtractSourceTableOutputReference
	_jsii_.Get(
		j,
		"sourceTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobExtractOutputReference) SourceTableInput() *BigqueryJobExtractSourceTable {
	var returns *BigqueryJobExtractSourceTable
	_jsii_.Get(
		j,
		"sourceTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobExtractOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobExtractOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobExtractOutputReference) UseAvroLogicalTypes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAvroLogicalTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobExtractOutputReference) UseAvroLogicalTypesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAvroLogicalTypesInput",
		&returns,
	)
	return returns
}


func NewBigqueryJobExtractOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BigqueryJobExtractOutputReference {
	_init_.Initialize()

	if err := validateNewBigqueryJobExtractOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BigqueryJobExtractOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.bigqueryJob.BigqueryJobExtractOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBigqueryJobExtractOutputReference_Override(b BigqueryJobExtractOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.bigqueryJob.BigqueryJobExtractOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BigqueryJobExtractOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobExtractOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobExtractOutputReference)SetCompression(val *string) {
	if err := j.validateSetCompressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compression",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobExtractOutputReference)SetDestinationFormat(val *string) {
	if err := j.validateSetDestinationFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationFormat",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobExtractOutputReference)SetDestinationUris(val *[]*string) {
	if err := j.validateSetDestinationUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationUris",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobExtractOutputReference)SetFieldDelimiter(val *string) {
	if err := j.validateSetFieldDelimiterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldDelimiter",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobExtractOutputReference)SetInternalValue(val *BigqueryJobExtract) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobExtractOutputReference)SetPrintHeader(val interface{}) {
	if err := j.validateSetPrintHeaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"printHeader",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobExtractOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobExtractOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobExtractOutputReference)SetUseAvroLogicalTypes(val interface{}) {
	if err := j.validateSetUseAvroLogicalTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useAvroLogicalTypes",
		val,
	)
}

func (b *jsiiProxy_BigqueryJobExtractOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryJobExtractOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryJobExtractOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryJobExtractOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryJobExtractOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryJobExtractOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryJobExtractOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryJobExtractOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryJobExtractOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryJobExtractOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryJobExtractOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryJobExtractOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryJobExtractOutputReference) PutSourceModel(value *BigqueryJobExtractSourceModel) {
	if err := b.validatePutSourceModelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putSourceModel",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BigqueryJobExtractOutputReference) PutSourceTable(value *BigqueryJobExtractSourceTable) {
	if err := b.validatePutSourceTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putSourceTable",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BigqueryJobExtractOutputReference) ResetCompression() {
	_jsii_.InvokeVoid(
		b,
		"resetCompression",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryJobExtractOutputReference) ResetDestinationFormat() {
	_jsii_.InvokeVoid(
		b,
		"resetDestinationFormat",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryJobExtractOutputReference) ResetFieldDelimiter() {
	_jsii_.InvokeVoid(
		b,
		"resetFieldDelimiter",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryJobExtractOutputReference) ResetPrintHeader() {
	_jsii_.InvokeVoid(
		b,
		"resetPrintHeader",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryJobExtractOutputReference) ResetSourceModel() {
	_jsii_.InvokeVoid(
		b,
		"resetSourceModel",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryJobExtractOutputReference) ResetSourceTable() {
	_jsii_.InvokeVoid(
		b,
		"resetSourceTable",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryJobExtractOutputReference) ResetUseAvroLogicalTypes() {
	_jsii_.InvokeVoid(
		b,
		"resetUseAvroLogicalTypes",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryJobExtractOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := b.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryJobExtractOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

