// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigquerytable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v13/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v13/bigquerytable/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference interface {
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
	Encoding() *string
	SetEncoding(val *string)
	EncodingInput() *string
	FieldName() *string
	SetFieldName(val *string)
	FieldNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OnlyReadLatest() interface{}
	SetOnlyReadLatest(val interface{})
	OnlyReadLatestInput() interface{}
	QualifierEncoded() *string
	SetQualifierEncoded(val *string)
	QualifierEncodedInput() *string
	QualifierString() *string
	SetQualifierString(val *string)
	QualifierStringInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetEncoding()
	ResetFieldName()
	ResetOnlyReadLatest()
	ResetQualifierEncoded()
	ResetQualifierString()
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference
type jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) Encoding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encoding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) EncodingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encodingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) FieldName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) FieldNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fieldNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) OnlyReadLatest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onlyReadLatest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) OnlyReadLatestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onlyReadLatestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) QualifierEncoded() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualifierEncoded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) QualifierEncodedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualifierEncodedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) QualifierString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualifierString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) QualifierStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualifierStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference {
	_init_.Initialize()

	if err := validateNewBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.bigqueryTable.BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference_Override(b BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.bigqueryTable.BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference)SetEncoding(val *string) {
	if err := j.validateSetEncodingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encoding",
		val,
	)
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference)SetFieldName(val *string) {
	if err := j.validateSetFieldNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fieldName",
		val,
	)
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference)SetOnlyReadLatest(val interface{}) {
	if err := j.validateSetOnlyReadLatestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onlyReadLatest",
		val,
	)
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference)SetQualifierEncoded(val *string) {
	if err := j.validateSetQualifierEncodedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"qualifierEncoded",
		val,
	)
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference)SetQualifierString(val *string) {
	if err := j.validateSetQualifierStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"qualifierString",
		val,
	)
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) ResetEncoding() {
	_jsii_.InvokeVoid(
		b,
		"resetEncoding",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) ResetFieldName() {
	_jsii_.InvokeVoid(
		b,
		"resetFieldName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) ResetOnlyReadLatest() {
	_jsii_.InvokeVoid(
		b,
		"resetOnlyReadLatest",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) ResetQualifierEncoded() {
	_jsii_.InvokeVoid(
		b,
		"resetQualifierEncoded",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) ResetQualifierString() {
	_jsii_.InvokeVoid(
		b,
		"resetQualifierString",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		b,
		"resetType",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

