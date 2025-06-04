// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package documentaiwarehousedocumentschema

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/documentaiwarehousedocumentschema/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference interface {
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
	InternalValue() *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptions
	SetInternalValue(val *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptions)
	PossibleValues() *[]*string
	SetPossibleValues(val *[]*string)
	PossibleValuesInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ValidationCheckDisabled() interface{}
	SetValidationCheckDisabled(val interface{})
	ValidationCheckDisabledInput() interface{}
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
	ResetValidationCheckDisabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference
type jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference) InternalValue() *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptions {
	var returns *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference) PossibleValues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"possibleValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference) PossibleValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"possibleValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference) ValidationCheckDisabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validationCheckDisabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference) ValidationCheckDisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validationCheckDisabledInput",
		&returns,
	)
	return returns
}


func NewDocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewDocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.documentAiWarehouseDocumentSchema.DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference_Override(d DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.documentAiWarehouseDocumentSchema.DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference)SetInternalValue(val *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference)SetPossibleValues(val *[]*string) {
	if err := j.validateSetPossibleValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"possibleValues",
		val,
	)
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference)SetValidationCheckDisabled(val interface{}) {
	if err := j.validateSetValidationCheckDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validationCheckDisabled",
		val,
	)
}

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference) ResetValidationCheckDisabled() {
	_jsii_.InvokeVoid(
		d,
		"resetValidationCheckDisabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

