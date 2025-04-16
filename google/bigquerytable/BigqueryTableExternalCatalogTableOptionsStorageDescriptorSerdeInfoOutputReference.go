// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigquerytable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/bigquerytable/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference interface {
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
	InternalValue() *BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfo
	SetInternalValue(val *BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfo)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Parameters() *map[string]*string
	SetParameters(val *map[string]*string)
	ParametersInput() *map[string]*string
	SerializationLibrary() *string
	SetSerializationLibrary(val *string)
	SerializationLibraryInput() *string
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
	ResetName()
	ResetParameters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference
type jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference) InternalValue() *BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfo {
	var returns *BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference) SerializationLibrary() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serializationLibrary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference) SerializationLibraryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serializationLibraryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference {
	_init_.Initialize()

	if err := validateNewBigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.bigqueryTable.BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference_Override(b BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.bigqueryTable.BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference)SetInternalValue(val *BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfo) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference)SetParameters(val *map[string]*string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference)SetSerializationLibrary(val *string) {
	if err := j.validateSetSerializationLibraryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serializationLibrary",
		val,
	)
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		b,
		"resetName",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference) ResetParameters() {
	_jsii_.InvokeVoid(
		b,
		"resetParameters",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

