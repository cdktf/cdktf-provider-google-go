package dataprocjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v5/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v5/dataprocjob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataprocJobPrestoConfigOutputReference interface {
	cdktf.ComplexObject
	ClientTags() *[]*string
	SetClientTags(val *[]*string)
	ClientTagsInput() *[]*string
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
	ContinueOnFailure() interface{}
	SetContinueOnFailure(val interface{})
	ContinueOnFailureInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *DataprocJobPrestoConfig
	SetInternalValue(val *DataprocJobPrestoConfig)
	LoggingConfig() DataprocJobPrestoConfigLoggingConfigOutputReference
	LoggingConfigInput() *DataprocJobPrestoConfigLoggingConfig
	OutputFormat() *string
	SetOutputFormat(val *string)
	OutputFormatInput() *string
	Properties() *map[string]*string
	SetProperties(val *map[string]*string)
	PropertiesInput() *map[string]*string
	QueryFileUri() *string
	SetQueryFileUri(val *string)
	QueryFileUriInput() *string
	QueryList() *[]*string
	SetQueryList(val *[]*string)
	QueryListInput() *[]*string
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
	PutLoggingConfig(value *DataprocJobPrestoConfigLoggingConfig)
	ResetClientTags()
	ResetContinueOnFailure()
	ResetLoggingConfig()
	ResetOutputFormat()
	ResetProperties()
	ResetQueryFileUri()
	ResetQueryList()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataprocJobPrestoConfigOutputReference
type jsiiProxy_DataprocJobPrestoConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataprocJobPrestoConfigOutputReference) ClientTags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clientTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJobPrestoConfigOutputReference) ClientTagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clientTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJobPrestoConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJobPrestoConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJobPrestoConfigOutputReference) ContinueOnFailure() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"continueOnFailure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJobPrestoConfigOutputReference) ContinueOnFailureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"continueOnFailureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJobPrestoConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJobPrestoConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJobPrestoConfigOutputReference) InternalValue() *DataprocJobPrestoConfig {
	var returns *DataprocJobPrestoConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJobPrestoConfigOutputReference) LoggingConfig() DataprocJobPrestoConfigLoggingConfigOutputReference {
	var returns DataprocJobPrestoConfigLoggingConfigOutputReference
	_jsii_.Get(
		j,
		"loggingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJobPrestoConfigOutputReference) LoggingConfigInput() *DataprocJobPrestoConfigLoggingConfig {
	var returns *DataprocJobPrestoConfigLoggingConfig
	_jsii_.Get(
		j,
		"loggingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJobPrestoConfigOutputReference) OutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJobPrestoConfigOutputReference) OutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJobPrestoConfigOutputReference) Properties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJobPrestoConfigOutputReference) PropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"propertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJobPrestoConfigOutputReference) QueryFileUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryFileUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJobPrestoConfigOutputReference) QueryFileUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryFileUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJobPrestoConfigOutputReference) QueryList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJobPrestoConfigOutputReference) QueryListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJobPrestoConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocJobPrestoConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataprocJobPrestoConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataprocJobPrestoConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataprocJobPrestoConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataprocJobPrestoConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataprocJob.DataprocJobPrestoConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataprocJobPrestoConfigOutputReference_Override(d DataprocJobPrestoConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataprocJob.DataprocJobPrestoConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataprocJobPrestoConfigOutputReference)SetClientTags(val *[]*string) {
	if err := j.validateSetClientTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientTags",
		val,
	)
}

func (j *jsiiProxy_DataprocJobPrestoConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataprocJobPrestoConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataprocJobPrestoConfigOutputReference)SetContinueOnFailure(val interface{}) {
	if err := j.validateSetContinueOnFailureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"continueOnFailure",
		val,
	)
}

func (j *jsiiProxy_DataprocJobPrestoConfigOutputReference)SetInternalValue(val *DataprocJobPrestoConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataprocJobPrestoConfigOutputReference)SetOutputFormat(val *string) {
	if err := j.validateSetOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputFormat",
		val,
	)
}

func (j *jsiiProxy_DataprocJobPrestoConfigOutputReference)SetProperties(val *map[string]*string) {
	if err := j.validateSetPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"properties",
		val,
	)
}

func (j *jsiiProxy_DataprocJobPrestoConfigOutputReference)SetQueryFileUri(val *string) {
	if err := j.validateSetQueryFileUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryFileUri",
		val,
	)
}

func (j *jsiiProxy_DataprocJobPrestoConfigOutputReference)SetQueryList(val *[]*string) {
	if err := j.validateSetQueryListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryList",
		val,
	)
}

func (j *jsiiProxy_DataprocJobPrestoConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataprocJobPrestoConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataprocJobPrestoConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocJobPrestoConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataprocJobPrestoConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataprocJobPrestoConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataprocJobPrestoConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataprocJobPrestoConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataprocJobPrestoConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataprocJobPrestoConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataprocJobPrestoConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataprocJobPrestoConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataprocJobPrestoConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocJobPrestoConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataprocJobPrestoConfigOutputReference) PutLoggingConfig(value *DataprocJobPrestoConfigLoggingConfig) {
	if err := d.validatePutLoggingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLoggingConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocJobPrestoConfigOutputReference) ResetClientTags() {
	_jsii_.InvokeVoid(
		d,
		"resetClientTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocJobPrestoConfigOutputReference) ResetContinueOnFailure() {
	_jsii_.InvokeVoid(
		d,
		"resetContinueOnFailure",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocJobPrestoConfigOutputReference) ResetLoggingConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetLoggingConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocJobPrestoConfigOutputReference) ResetOutputFormat() {
	_jsii_.InvokeVoid(
		d,
		"resetOutputFormat",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocJobPrestoConfigOutputReference) ResetProperties() {
	_jsii_.InvokeVoid(
		d,
		"resetProperties",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocJobPrestoConfigOutputReference) ResetQueryFileUri() {
	_jsii_.InvokeVoid(
		d,
		"resetQueryFileUri",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocJobPrestoConfigOutputReference) ResetQueryList() {
	_jsii_.InvokeVoid(
		d,
		"resetQueryList",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocJobPrestoConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataprocJobPrestoConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

