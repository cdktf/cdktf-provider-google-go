package datalosspreventionstoredinfotype

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v6/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v6/datalosspreventionstoredinfotype/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference interface {
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
	Field() DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldFieldOutputReference
	FieldInput() *DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldField
	// Experimental.
	Fqn() *string
	InternalValue() *DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryField
	SetInternalValue(val *DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryField)
	Table() DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldTableOutputReference
	TableInput() *DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable
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
	PutField(value *DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldField)
	PutTable(value *DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference
type jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference) Field() DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldFieldOutputReference {
	var returns DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldFieldOutputReference
	_jsii_.Get(
		j,
		"field",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference) FieldInput() *DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldField {
	var returns *DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldField
	_jsii_.Get(
		j,
		"fieldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference) InternalValue() *DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryField {
	var returns *DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryField
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference) Table() DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldTableOutputReference {
	var returns DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldTableOutputReference
	_jsii_.Get(
		j,
		"table",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference) TableInput() *DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable {
	var returns *DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable
	_jsii_.Get(
		j,
		"tableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference {
	_init_.Initialize()

	if err := validateNewDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataLossPreventionStoredInfoType.DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference_Override(d DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataLossPreventionStoredInfoType.DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference)SetInternalValue(val *DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryField) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference) PutField(value *DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldField) {
	if err := d.validatePutFieldParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putField",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference) PutTable(value *DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldTable) {
	if err := d.validatePutTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTable",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataLossPreventionStoredInfoTypeLargeCustomDictionaryBigQueryFieldOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

