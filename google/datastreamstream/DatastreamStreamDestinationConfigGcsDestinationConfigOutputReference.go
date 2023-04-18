package datastreamstream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v6/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v6/datastreamstream/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference interface {
	cdktf.ComplexObject
	AvroFileFormat() DatastreamStreamDestinationConfigGcsDestinationConfigAvroFileFormatOutputReference
	AvroFileFormatInput() *DatastreamStreamDestinationConfigGcsDestinationConfigAvroFileFormat
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
	FileRotationInterval() *string
	SetFileRotationInterval(val *string)
	FileRotationIntervalInput() *string
	FileRotationMb() *float64
	SetFileRotationMb(val *float64)
	FileRotationMbInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *DatastreamStreamDestinationConfigGcsDestinationConfig
	SetInternalValue(val *DatastreamStreamDestinationConfigGcsDestinationConfig)
	JsonFileFormat() DatastreamStreamDestinationConfigGcsDestinationConfigJsonFileFormatOutputReference
	JsonFileFormatInput() *DatastreamStreamDestinationConfigGcsDestinationConfigJsonFileFormat
	Path() *string
	SetPath(val *string)
	PathInput() *string
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
	PutAvroFileFormat(value *DatastreamStreamDestinationConfigGcsDestinationConfigAvroFileFormat)
	PutJsonFileFormat(value *DatastreamStreamDestinationConfigGcsDestinationConfigJsonFileFormat)
	ResetAvroFileFormat()
	ResetFileRotationInterval()
	ResetFileRotationMb()
	ResetJsonFileFormat()
	ResetPath()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference
type jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference) AvroFileFormat() DatastreamStreamDestinationConfigGcsDestinationConfigAvroFileFormatOutputReference {
	var returns DatastreamStreamDestinationConfigGcsDestinationConfigAvroFileFormatOutputReference
	_jsii_.Get(
		j,
		"avroFileFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference) AvroFileFormatInput() *DatastreamStreamDestinationConfigGcsDestinationConfigAvroFileFormat {
	var returns *DatastreamStreamDestinationConfigGcsDestinationConfigAvroFileFormat
	_jsii_.Get(
		j,
		"avroFileFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference) FileRotationInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileRotationInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference) FileRotationIntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileRotationIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference) FileRotationMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fileRotationMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference) FileRotationMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fileRotationMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference) InternalValue() *DatastreamStreamDestinationConfigGcsDestinationConfig {
	var returns *DatastreamStreamDestinationConfigGcsDestinationConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference) JsonFileFormat() DatastreamStreamDestinationConfigGcsDestinationConfigJsonFileFormatOutputReference {
	var returns DatastreamStreamDestinationConfigGcsDestinationConfigJsonFileFormatOutputReference
	_jsii_.Get(
		j,
		"jsonFileFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference) JsonFileFormatInput() *DatastreamStreamDestinationConfigGcsDestinationConfigJsonFileFormat {
	var returns *DatastreamStreamDestinationConfigGcsDestinationConfigJsonFileFormat
	_jsii_.Get(
		j,
		"jsonFileFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDatastreamStreamDestinationConfigGcsDestinationConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDatastreamStreamDestinationConfigGcsDestinationConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.datastreamStream.DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatastreamStreamDestinationConfigGcsDestinationConfigOutputReference_Override(d DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.datastreamStream.DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference)SetFileRotationInterval(val *string) {
	if err := j.validateSetFileRotationIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileRotationInterval",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference)SetFileRotationMb(val *float64) {
	if err := j.validateSetFileRotationMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileRotationMb",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference)SetInternalValue(val *DatastreamStreamDestinationConfigGcsDestinationConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference) PutAvroFileFormat(value *DatastreamStreamDestinationConfigGcsDestinationConfigAvroFileFormat) {
	if err := d.validatePutAvroFileFormatParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAvroFileFormat",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference) PutJsonFileFormat(value *DatastreamStreamDestinationConfigGcsDestinationConfigJsonFileFormat) {
	if err := d.validatePutJsonFileFormatParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putJsonFileFormat",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference) ResetAvroFileFormat() {
	_jsii_.InvokeVoid(
		d,
		"resetAvroFileFormat",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference) ResetFileRotationInterval() {
	_jsii_.InvokeVoid(
		d,
		"resetFileRotationInterval",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference) ResetFileRotationMb() {
	_jsii_.InvokeVoid(
		d,
		"resetFileRotationMb",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference) ResetJsonFileFormat() {
	_jsii_.InvokeVoid(
		d,
		"resetJsonFileFormat",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		d,
		"resetPath",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DatastreamStreamDestinationConfigGcsDestinationConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

