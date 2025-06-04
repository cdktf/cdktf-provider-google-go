// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventioninspecttemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/datalosspreventioninspecttemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference interface {
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
	Dictionary() DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesDictionaryOutputReference
	DictionaryInput() *DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesDictionary
	ExclusionType() *string
	SetExclusionType(val *string)
	ExclusionTypeInput() *string
	// Experimental.
	Fqn() *string
	InfoType() DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesInfoTypeOutputReference
	InfoTypeInput() *DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesInfoType
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Likelihood() *string
	SetLikelihood(val *string)
	LikelihoodInput() *string
	Regex() DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesRegexOutputReference
	RegexInput() *DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesRegex
	SensitivityScore() DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesSensitivityScoreOutputReference
	SensitivityScoreInput() *DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesSensitivityScore
	StoredType() DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesStoredTypeOutputReference
	StoredTypeInput() *DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesStoredType
	SurrogateType() DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesSurrogateTypeOutputReference
	SurrogateTypeInput() *DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesSurrogateType
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
	PutDictionary(value *DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesDictionary)
	PutInfoType(value *DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesInfoType)
	PutRegex(value *DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesRegex)
	PutSensitivityScore(value *DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesSensitivityScore)
	PutStoredType(value *DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesStoredType)
	PutSurrogateType(value *DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesSurrogateType)
	ResetDictionary()
	ResetExclusionType()
	ResetLikelihood()
	ResetRegex()
	ResetSensitivityScore()
	ResetStoredType()
	ResetSurrogateType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference
type jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) Dictionary() DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesDictionaryOutputReference {
	var returns DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesDictionaryOutputReference
	_jsii_.Get(
		j,
		"dictionary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) DictionaryInput() *DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesDictionary {
	var returns *DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesDictionary
	_jsii_.Get(
		j,
		"dictionaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) ExclusionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exclusionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) ExclusionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exclusionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) InfoType() DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesInfoTypeOutputReference {
	var returns DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesInfoTypeOutputReference
	_jsii_.Get(
		j,
		"infoType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) InfoTypeInput() *DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesInfoType {
	var returns *DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesInfoType
	_jsii_.Get(
		j,
		"infoTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) Likelihood() *string {
	var returns *string
	_jsii_.Get(
		j,
		"likelihood",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) LikelihoodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"likelihoodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) Regex() DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesRegexOutputReference {
	var returns DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesRegexOutputReference
	_jsii_.Get(
		j,
		"regex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) RegexInput() *DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesRegex {
	var returns *DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesRegex
	_jsii_.Get(
		j,
		"regexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) SensitivityScore() DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesSensitivityScoreOutputReference {
	var returns DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesSensitivityScoreOutputReference
	_jsii_.Get(
		j,
		"sensitivityScore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) SensitivityScoreInput() *DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesSensitivityScore {
	var returns *DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesSensitivityScore
	_jsii_.Get(
		j,
		"sensitivityScoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) StoredType() DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesStoredTypeOutputReference {
	var returns DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesStoredTypeOutputReference
	_jsii_.Get(
		j,
		"storedType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) StoredTypeInput() *DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesStoredType {
	var returns *DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesStoredType
	_jsii_.Get(
		j,
		"storedTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) SurrogateType() DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesSurrogateTypeOutputReference {
	var returns DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesSurrogateTypeOutputReference
	_jsii_.Get(
		j,
		"surrogateType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) SurrogateTypeInput() *DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesSurrogateType {
	var returns *DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesSurrogateType
	_jsii_.Get(
		j,
		"surrogateTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference {
	_init_.Initialize()

	if err := validateNewDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataLossPreventionInspectTemplate.DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference_Override(d DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataLossPreventionInspectTemplate.DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference)SetExclusionType(val *string) {
	if err := j.validateSetExclusionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exclusionType",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference)SetLikelihood(val *string) {
	if err := j.validateSetLikelihoodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"likelihood",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) PutDictionary(value *DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesDictionary) {
	if err := d.validatePutDictionaryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDictionary",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) PutInfoType(value *DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesInfoType) {
	if err := d.validatePutInfoTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putInfoType",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) PutRegex(value *DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesRegex) {
	if err := d.validatePutRegexParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRegex",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) PutSensitivityScore(value *DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesSensitivityScore) {
	if err := d.validatePutSensitivityScoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSensitivityScore",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) PutStoredType(value *DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesStoredType) {
	if err := d.validatePutStoredTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putStoredType",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) PutSurrogateType(value *DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesSurrogateType) {
	if err := d.validatePutSurrogateTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSurrogateType",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) ResetDictionary() {
	_jsii_.InvokeVoid(
		d,
		"resetDictionary",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) ResetExclusionType() {
	_jsii_.InvokeVoid(
		d,
		"resetExclusionType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) ResetLikelihood() {
	_jsii_.InvokeVoid(
		d,
		"resetLikelihood",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) ResetRegex() {
	_jsii_.InvokeVoid(
		d,
		"resetRegex",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) ResetSensitivityScore() {
	_jsii_.InvokeVoid(
		d,
		"resetSensitivityScore",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) ResetStoredType() {
	_jsii_.InvokeVoid(
		d,
		"resetStoredType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) ResetSurrogateType() {
	_jsii_.InvokeVoid(
		d,
		"resetSurrogateType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataLossPreventionInspectTemplateInspectConfigCustomInfoTypesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

