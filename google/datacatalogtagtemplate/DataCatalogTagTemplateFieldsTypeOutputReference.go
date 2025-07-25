// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacatalogtagtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/datacatalogtagtemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCatalogTagTemplateFieldsTypeOutputReference interface {
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
	EnumType() DataCatalogTagTemplateFieldsTypeEnumTypeOutputReference
	EnumTypeInput() *DataCatalogTagTemplateFieldsTypeEnumType
	// Experimental.
	Fqn() *string
	InternalValue() *DataCatalogTagTemplateFieldsType
	SetInternalValue(val *DataCatalogTagTemplateFieldsType)
	PrimitiveType() *string
	SetPrimitiveType(val *string)
	PrimitiveTypeInput() *string
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
	PutEnumType(value *DataCatalogTagTemplateFieldsTypeEnumType)
	ResetEnumType()
	ResetPrimitiveType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataCatalogTagTemplateFieldsTypeOutputReference
type jsiiProxy_DataCatalogTagTemplateFieldsTypeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCatalogTagTemplateFieldsTypeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCatalogTagTemplateFieldsTypeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCatalogTagTemplateFieldsTypeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCatalogTagTemplateFieldsTypeOutputReference) EnumType() DataCatalogTagTemplateFieldsTypeEnumTypeOutputReference {
	var returns DataCatalogTagTemplateFieldsTypeEnumTypeOutputReference
	_jsii_.Get(
		j,
		"enumType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCatalogTagTemplateFieldsTypeOutputReference) EnumTypeInput() *DataCatalogTagTemplateFieldsTypeEnumType {
	var returns *DataCatalogTagTemplateFieldsTypeEnumType
	_jsii_.Get(
		j,
		"enumTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCatalogTagTemplateFieldsTypeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCatalogTagTemplateFieldsTypeOutputReference) InternalValue() *DataCatalogTagTemplateFieldsType {
	var returns *DataCatalogTagTemplateFieldsType
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCatalogTagTemplateFieldsTypeOutputReference) PrimitiveType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primitiveType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCatalogTagTemplateFieldsTypeOutputReference) PrimitiveTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primitiveTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCatalogTagTemplateFieldsTypeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCatalogTagTemplateFieldsTypeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataCatalogTagTemplateFieldsTypeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataCatalogTagTemplateFieldsTypeOutputReference {
	_init_.Initialize()

	if err := validateNewDataCatalogTagTemplateFieldsTypeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCatalogTagTemplateFieldsTypeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataCatalogTagTemplate.DataCatalogTagTemplateFieldsTypeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataCatalogTagTemplateFieldsTypeOutputReference_Override(d DataCatalogTagTemplateFieldsTypeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataCatalogTagTemplate.DataCatalogTagTemplateFieldsTypeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataCatalogTagTemplateFieldsTypeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCatalogTagTemplateFieldsTypeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCatalogTagTemplateFieldsTypeOutputReference)SetInternalValue(val *DataCatalogTagTemplateFieldsType) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCatalogTagTemplateFieldsTypeOutputReference)SetPrimitiveType(val *string) {
	if err := j.validateSetPrimitiveTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primitiveType",
		val,
	)
}

func (j *jsiiProxy_DataCatalogTagTemplateFieldsTypeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCatalogTagTemplateFieldsTypeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCatalogTagTemplateFieldsTypeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCatalogTagTemplateFieldsTypeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCatalogTagTemplateFieldsTypeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCatalogTagTemplateFieldsTypeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCatalogTagTemplateFieldsTypeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCatalogTagTemplateFieldsTypeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCatalogTagTemplateFieldsTypeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCatalogTagTemplateFieldsTypeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCatalogTagTemplateFieldsTypeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCatalogTagTemplateFieldsTypeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCatalogTagTemplateFieldsTypeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCatalogTagTemplateFieldsTypeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCatalogTagTemplateFieldsTypeOutputReference) PutEnumType(value *DataCatalogTagTemplateFieldsTypeEnumType) {
	if err := d.validatePutEnumTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEnumType",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataCatalogTagTemplateFieldsTypeOutputReference) ResetEnumType() {
	_jsii_.InvokeVoid(
		d,
		"resetEnumType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCatalogTagTemplateFieldsTypeOutputReference) ResetPrimitiveType() {
	_jsii_.InvokeVoid(
		d,
		"resetPrimitiveType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCatalogTagTemplateFieldsTypeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataCatalogTagTemplateFieldsTypeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

