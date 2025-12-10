// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package documentaiwarehousedocumentschema

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/documentaiwarehousedocumentschema/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference interface {
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
	DateTimeTypeOptions() DocumentAiWarehouseDocumentSchemaPropertyDefinitionsDateTimeTypeOptionsOutputReference
	DateTimeTypeOptionsInput() *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsDateTimeTypeOptions
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EnumTypeOptions() DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference
	EnumTypeOptionsInput() *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptions
	FloatTypeOptions() DocumentAiWarehouseDocumentSchemaPropertyDefinitionsFloatTypeOptionsOutputReference
	FloatTypeOptionsInput() *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsFloatTypeOptions
	// Experimental.
	Fqn() *string
	IntegerTypeOptions() DocumentAiWarehouseDocumentSchemaPropertyDefinitionsIntegerTypeOptionsOutputReference
	IntegerTypeOptionsInput() *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsIntegerTypeOptions
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IsFilterable() interface{}
	SetIsFilterable(val interface{})
	IsFilterableInput() interface{}
	IsMetadata() interface{}
	SetIsMetadata(val interface{})
	IsMetadataInput() interface{}
	IsRepeatable() interface{}
	SetIsRepeatable(val interface{})
	IsRepeatableInput() interface{}
	IsRequired() interface{}
	SetIsRequired(val interface{})
	IsRequiredInput() interface{}
	IsSearchable() interface{}
	SetIsSearchable(val interface{})
	IsSearchableInput() interface{}
	MapTypeOptions() DocumentAiWarehouseDocumentSchemaPropertyDefinitionsMapTypeOptionsOutputReference
	MapTypeOptionsInput() *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsMapTypeOptions
	Name() *string
	SetName(val *string)
	NameInput() *string
	PropertyTypeOptions() DocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsOutputReference
	PropertyTypeOptionsInput() *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptions
	RetrievalImportance() *string
	SetRetrievalImportance(val *string)
	RetrievalImportanceInput() *string
	SchemaSources() DocumentAiWarehouseDocumentSchemaPropertyDefinitionsSchemaSourcesList
	SchemaSourcesInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TextTypeOptions() DocumentAiWarehouseDocumentSchemaPropertyDefinitionsTextTypeOptionsOutputReference
	TextTypeOptionsInput() *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsTextTypeOptions
	TimestampTypeOptions() DocumentAiWarehouseDocumentSchemaPropertyDefinitionsTimestampTypeOptionsOutputReference
	TimestampTypeOptionsInput() *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsTimestampTypeOptions
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
	PutDateTimeTypeOptions(value *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsDateTimeTypeOptions)
	PutEnumTypeOptions(value *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptions)
	PutFloatTypeOptions(value *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsFloatTypeOptions)
	PutIntegerTypeOptions(value *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsIntegerTypeOptions)
	PutMapTypeOptions(value *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsMapTypeOptions)
	PutPropertyTypeOptions(value *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptions)
	PutSchemaSources(value interface{})
	PutTextTypeOptions(value *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsTextTypeOptions)
	PutTimestampTypeOptions(value *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsTimestampTypeOptions)
	ResetDateTimeTypeOptions()
	ResetDisplayName()
	ResetEnumTypeOptions()
	ResetFloatTypeOptions()
	ResetIntegerTypeOptions()
	ResetIsFilterable()
	ResetIsMetadata()
	ResetIsRepeatable()
	ResetIsRequired()
	ResetIsSearchable()
	ResetMapTypeOptions()
	ResetPropertyTypeOptions()
	ResetRetrievalImportance()
	ResetSchemaSources()
	ResetTextTypeOptions()
	ResetTimestampTypeOptions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference
type jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) DateTimeTypeOptions() DocumentAiWarehouseDocumentSchemaPropertyDefinitionsDateTimeTypeOptionsOutputReference {
	var returns DocumentAiWarehouseDocumentSchemaPropertyDefinitionsDateTimeTypeOptionsOutputReference
	_jsii_.Get(
		j,
		"dateTimeTypeOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) DateTimeTypeOptionsInput() *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsDateTimeTypeOptions {
	var returns *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsDateTimeTypeOptions
	_jsii_.Get(
		j,
		"dateTimeTypeOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) EnumTypeOptions() DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference {
	var returns DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference
	_jsii_.Get(
		j,
		"enumTypeOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) EnumTypeOptionsInput() *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptions {
	var returns *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptions
	_jsii_.Get(
		j,
		"enumTypeOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) FloatTypeOptions() DocumentAiWarehouseDocumentSchemaPropertyDefinitionsFloatTypeOptionsOutputReference {
	var returns DocumentAiWarehouseDocumentSchemaPropertyDefinitionsFloatTypeOptionsOutputReference
	_jsii_.Get(
		j,
		"floatTypeOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) FloatTypeOptionsInput() *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsFloatTypeOptions {
	var returns *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsFloatTypeOptions
	_jsii_.Get(
		j,
		"floatTypeOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) IntegerTypeOptions() DocumentAiWarehouseDocumentSchemaPropertyDefinitionsIntegerTypeOptionsOutputReference {
	var returns DocumentAiWarehouseDocumentSchemaPropertyDefinitionsIntegerTypeOptionsOutputReference
	_jsii_.Get(
		j,
		"integerTypeOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) IntegerTypeOptionsInput() *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsIntegerTypeOptions {
	var returns *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsIntegerTypeOptions
	_jsii_.Get(
		j,
		"integerTypeOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) IsFilterable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isFilterable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) IsFilterableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isFilterableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) IsMetadata() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) IsMetadataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) IsRepeatable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isRepeatable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) IsRepeatableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isRepeatableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) IsRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) IsRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) IsSearchable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSearchable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) IsSearchableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isSearchableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) MapTypeOptions() DocumentAiWarehouseDocumentSchemaPropertyDefinitionsMapTypeOptionsOutputReference {
	var returns DocumentAiWarehouseDocumentSchemaPropertyDefinitionsMapTypeOptionsOutputReference
	_jsii_.Get(
		j,
		"mapTypeOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) MapTypeOptionsInput() *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsMapTypeOptions {
	var returns *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsMapTypeOptions
	_jsii_.Get(
		j,
		"mapTypeOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) PropertyTypeOptions() DocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsOutputReference {
	var returns DocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsOutputReference
	_jsii_.Get(
		j,
		"propertyTypeOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) PropertyTypeOptionsInput() *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptions {
	var returns *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptions
	_jsii_.Get(
		j,
		"propertyTypeOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) RetrievalImportance() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retrievalImportance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) RetrievalImportanceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retrievalImportanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) SchemaSources() DocumentAiWarehouseDocumentSchemaPropertyDefinitionsSchemaSourcesList {
	var returns DocumentAiWarehouseDocumentSchemaPropertyDefinitionsSchemaSourcesList
	_jsii_.Get(
		j,
		"schemaSources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) SchemaSourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schemaSourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) TextTypeOptions() DocumentAiWarehouseDocumentSchemaPropertyDefinitionsTextTypeOptionsOutputReference {
	var returns DocumentAiWarehouseDocumentSchemaPropertyDefinitionsTextTypeOptionsOutputReference
	_jsii_.Get(
		j,
		"textTypeOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) TextTypeOptionsInput() *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsTextTypeOptions {
	var returns *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsTextTypeOptions
	_jsii_.Get(
		j,
		"textTypeOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) TimestampTypeOptions() DocumentAiWarehouseDocumentSchemaPropertyDefinitionsTimestampTypeOptionsOutputReference {
	var returns DocumentAiWarehouseDocumentSchemaPropertyDefinitionsTimestampTypeOptionsOutputReference
	_jsii_.Get(
		j,
		"timestampTypeOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) TimestampTypeOptionsInput() *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsTimestampTypeOptions {
	var returns *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsTimestampTypeOptions
	_jsii_.Get(
		j,
		"timestampTypeOptionsInput",
		&returns,
	)
	return returns
}


func NewDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference {
	_init_.Initialize()

	if err := validateNewDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.documentAiWarehouseDocumentSchema.DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference_Override(d DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.documentAiWarehouseDocumentSchema.DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference)SetIsFilterable(val interface{}) {
	if err := j.validateSetIsFilterableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isFilterable",
		val,
	)
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference)SetIsMetadata(val interface{}) {
	if err := j.validateSetIsMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isMetadata",
		val,
	)
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference)SetIsRepeatable(val interface{}) {
	if err := j.validateSetIsRepeatableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isRepeatable",
		val,
	)
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference)SetIsRequired(val interface{}) {
	if err := j.validateSetIsRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isRequired",
		val,
	)
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference)SetIsSearchable(val interface{}) {
	if err := j.validateSetIsSearchableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isSearchable",
		val,
	)
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference)SetRetrievalImportance(val *string) {
	if err := j.validateSetRetrievalImportanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retrievalImportance",
		val,
	)
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) PutDateTimeTypeOptions(value *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsDateTimeTypeOptions) {
	if err := d.validatePutDateTimeTypeOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDateTimeTypeOptions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) PutEnumTypeOptions(value *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptions) {
	if err := d.validatePutEnumTypeOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEnumTypeOptions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) PutFloatTypeOptions(value *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsFloatTypeOptions) {
	if err := d.validatePutFloatTypeOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFloatTypeOptions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) PutIntegerTypeOptions(value *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsIntegerTypeOptions) {
	if err := d.validatePutIntegerTypeOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putIntegerTypeOptions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) PutMapTypeOptions(value *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsMapTypeOptions) {
	if err := d.validatePutMapTypeOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMapTypeOptions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) PutPropertyTypeOptions(value *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptions) {
	if err := d.validatePutPropertyTypeOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPropertyTypeOptions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) PutSchemaSources(value interface{}) {
	if err := d.validatePutSchemaSourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSchemaSources",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) PutTextTypeOptions(value *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsTextTypeOptions) {
	if err := d.validatePutTextTypeOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTextTypeOptions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) PutTimestampTypeOptions(value *DocumentAiWarehouseDocumentSchemaPropertyDefinitionsTimestampTypeOptions) {
	if err := d.validatePutTimestampTypeOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimestampTypeOptions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) ResetDateTimeTypeOptions() {
	_jsii_.InvokeVoid(
		d,
		"resetDateTimeTypeOptions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) ResetDisplayName() {
	_jsii_.InvokeVoid(
		d,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) ResetEnumTypeOptions() {
	_jsii_.InvokeVoid(
		d,
		"resetEnumTypeOptions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) ResetFloatTypeOptions() {
	_jsii_.InvokeVoid(
		d,
		"resetFloatTypeOptions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) ResetIntegerTypeOptions() {
	_jsii_.InvokeVoid(
		d,
		"resetIntegerTypeOptions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) ResetIsFilterable() {
	_jsii_.InvokeVoid(
		d,
		"resetIsFilterable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) ResetIsMetadata() {
	_jsii_.InvokeVoid(
		d,
		"resetIsMetadata",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) ResetIsRepeatable() {
	_jsii_.InvokeVoid(
		d,
		"resetIsRepeatable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) ResetIsRequired() {
	_jsii_.InvokeVoid(
		d,
		"resetIsRequired",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) ResetIsSearchable() {
	_jsii_.InvokeVoid(
		d,
		"resetIsSearchable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) ResetMapTypeOptions() {
	_jsii_.InvokeVoid(
		d,
		"resetMapTypeOptions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) ResetPropertyTypeOptions() {
	_jsii_.InvokeVoid(
		d,
		"resetPropertyTypeOptions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) ResetRetrievalImportance() {
	_jsii_.InvokeVoid(
		d,
		"resetRetrievalImportance",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) ResetSchemaSources() {
	_jsii_.InvokeVoid(
		d,
		"resetSchemaSources",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) ResetTextTypeOptions() {
	_jsii_.InvokeVoid(
		d,
		"resetTextTypeOptions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) ResetTimestampTypeOptions() {
	_jsii_.InvokeVoid(
		d,
		"resetTimestampTypeOptions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

