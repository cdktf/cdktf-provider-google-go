package datalosspreventiondeidentifytemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v7/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v7/datalosspreventiondeidentifytemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference interface {
	cdktf.ComplexObject
	AllInfoTypes() DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsAllInfoTypesOutputReference
	AllInfoTypesInput() *DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsAllInfoTypes
	AllText() DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsAllTextOutputReference
	AllTextInput() *DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsAllText
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RedactionColor() DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsRedactionColorOutputReference
	RedactionColorInput() *DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsRedactionColor
	SelectedInfoTypes() DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsSelectedInfoTypesOutputReference
	SelectedInfoTypesInput() *DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsSelectedInfoTypes
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
	PutAllInfoTypes(value *DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsAllInfoTypes)
	PutAllText(value *DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsAllText)
	PutRedactionColor(value *DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsRedactionColor)
	PutSelectedInfoTypes(value *DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsSelectedInfoTypes)
	ResetAllInfoTypes()
	ResetAllText()
	ResetRedactionColor()
	ResetSelectedInfoTypes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference
type jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) AllInfoTypes() DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsAllInfoTypesOutputReference {
	var returns DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsAllInfoTypesOutputReference
	_jsii_.Get(
		j,
		"allInfoTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) AllInfoTypesInput() *DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsAllInfoTypes {
	var returns *DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsAllInfoTypes
	_jsii_.Get(
		j,
		"allInfoTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) AllText() DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsAllTextOutputReference {
	var returns DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsAllTextOutputReference
	_jsii_.Get(
		j,
		"allText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) AllTextInput() *DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsAllText {
	var returns *DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsAllText
	_jsii_.Get(
		j,
		"allTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) RedactionColor() DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsRedactionColorOutputReference {
	var returns DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsRedactionColorOutputReference
	_jsii_.Get(
		j,
		"redactionColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) RedactionColorInput() *DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsRedactionColor {
	var returns *DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsRedactionColor
	_jsii_.Get(
		j,
		"redactionColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) SelectedInfoTypes() DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsSelectedInfoTypesOutputReference {
	var returns DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsSelectedInfoTypesOutputReference
	_jsii_.Get(
		j,
		"selectedInfoTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) SelectedInfoTypesInput() *DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsSelectedInfoTypes {
	var returns *DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsSelectedInfoTypes
	_jsii_.Get(
		j,
		"selectedInfoTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference {
	_init_.Initialize()

	if err := validateNewDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataLossPreventionDeidentifyTemplate.DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference_Override(d DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataLossPreventionDeidentifyTemplate.DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) PutAllInfoTypes(value *DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsAllInfoTypes) {
	if err := d.validatePutAllInfoTypesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAllInfoTypes",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) PutAllText(value *DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsAllText) {
	if err := d.validatePutAllTextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAllText",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) PutRedactionColor(value *DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsRedactionColor) {
	if err := d.validatePutRedactionColorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRedactionColor",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) PutSelectedInfoTypes(value *DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsSelectedInfoTypes) {
	if err := d.validatePutSelectedInfoTypesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSelectedInfoTypes",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) ResetAllInfoTypes() {
	_jsii_.InvokeVoid(
		d,
		"resetAllInfoTypes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) ResetAllText() {
	_jsii_.InvokeVoid(
		d,
		"resetAllText",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) ResetRedactionColor() {
	_jsii_.InvokeVoid(
		d,
		"resetRedactionColor",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) ResetSelectedInfoTypes() {
	_jsii_.InvokeVoid(
		d,
		"resetSelectedInfoTypes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataLossPreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

