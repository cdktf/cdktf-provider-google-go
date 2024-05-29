// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigquerytable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v13/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v13/bigquerytable/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference interface {
	cdktf.ComplexObject
	Column() BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnList
	ColumnInput() interface{}
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
	FamilyId() *string
	SetFamilyId(val *string)
	FamilyIdInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OnlyReadLatest() interface{}
	SetOnlyReadLatest(val interface{})
	OnlyReadLatestInput() interface{}
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
	PutColumn(value interface{})
	ResetColumn()
	ResetEncoding()
	ResetFamilyId()
	ResetOnlyReadLatest()
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

// The jsii proxy struct for BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference
type jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference) Column() BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnList {
	var returns BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyColumnList
	_jsii_.Get(
		j,
		"column",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference) ColumnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"columnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference) Encoding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encoding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference) EncodingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encodingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference) FamilyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"familyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference) FamilyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"familyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference) OnlyReadLatest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onlyReadLatest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference) OnlyReadLatestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onlyReadLatestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference {
	_init_.Initialize()

	if err := validateNewBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.bigqueryTable.BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference_Override(b BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.bigqueryTable.BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference)SetEncoding(val *string) {
	if err := j.validateSetEncodingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encoding",
		val,
	)
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference)SetFamilyId(val *string) {
	if err := j.validateSetFamilyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"familyId",
		val,
	)
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference)SetOnlyReadLatest(val interface{}) {
	if err := j.validateSetOnlyReadLatestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onlyReadLatest",
		val,
	)
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference) PutColumn(value interface{}) {
	if err := b.validatePutColumnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putColumn",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference) ResetColumn() {
	_jsii_.InvokeVoid(
		b,
		"resetColumn",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference) ResetEncoding() {
	_jsii_.InvokeVoid(
		b,
		"resetEncoding",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference) ResetFamilyId() {
	_jsii_.InvokeVoid(
		b,
		"resetFamilyId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference) ResetOnlyReadLatest() {
	_jsii_.InvokeVoid(
		b,
		"resetOnlyReadLatest",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		b,
		"resetType",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (b *jsiiProxy_BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamilyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

