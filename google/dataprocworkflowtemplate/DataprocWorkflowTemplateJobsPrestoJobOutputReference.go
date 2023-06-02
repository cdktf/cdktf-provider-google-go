package dataprocworkflowtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v7/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v7/dataprocworkflowtemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataprocWorkflowTemplateJobsPrestoJobOutputReference interface {
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
	InternalValue() *DataprocWorkflowTemplateJobsPrestoJob
	SetInternalValue(val *DataprocWorkflowTemplateJobsPrestoJob)
	LoggingConfig() DataprocWorkflowTemplateJobsPrestoJobLoggingConfigOutputReference
	LoggingConfigInput() *DataprocWorkflowTemplateJobsPrestoJobLoggingConfig
	OutputFormat() *string
	SetOutputFormat(val *string)
	OutputFormatInput() *string
	Properties() *map[string]*string
	SetProperties(val *map[string]*string)
	PropertiesInput() *map[string]*string
	QueryFileUri() *string
	SetQueryFileUri(val *string)
	QueryFileUriInput() *string
	QueryList() DataprocWorkflowTemplateJobsPrestoJobQueryListStructOutputReference
	QueryListInput() *DataprocWorkflowTemplateJobsPrestoJobQueryListStruct
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
	PutLoggingConfig(value *DataprocWorkflowTemplateJobsPrestoJobLoggingConfig)
	PutQueryList(value *DataprocWorkflowTemplateJobsPrestoJobQueryListStruct)
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

// The jsii proxy struct for DataprocWorkflowTemplateJobsPrestoJobOutputReference
type jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference) ClientTags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clientTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference) ClientTagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clientTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference) ContinueOnFailure() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"continueOnFailure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference) ContinueOnFailureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"continueOnFailureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference) InternalValue() *DataprocWorkflowTemplateJobsPrestoJob {
	var returns *DataprocWorkflowTemplateJobsPrestoJob
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference) LoggingConfig() DataprocWorkflowTemplateJobsPrestoJobLoggingConfigOutputReference {
	var returns DataprocWorkflowTemplateJobsPrestoJobLoggingConfigOutputReference
	_jsii_.Get(
		j,
		"loggingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference) LoggingConfigInput() *DataprocWorkflowTemplateJobsPrestoJobLoggingConfig {
	var returns *DataprocWorkflowTemplateJobsPrestoJobLoggingConfig
	_jsii_.Get(
		j,
		"loggingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference) OutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference) OutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference) Properties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference) PropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"propertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference) QueryFileUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryFileUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference) QueryFileUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryFileUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference) QueryList() DataprocWorkflowTemplateJobsPrestoJobQueryListStructOutputReference {
	var returns DataprocWorkflowTemplateJobsPrestoJobQueryListStructOutputReference
	_jsii_.Get(
		j,
		"queryList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference) QueryListInput() *DataprocWorkflowTemplateJobsPrestoJobQueryListStruct {
	var returns *DataprocWorkflowTemplateJobsPrestoJobQueryListStruct
	_jsii_.Get(
		j,
		"queryListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataprocWorkflowTemplateJobsPrestoJobOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataprocWorkflowTemplateJobsPrestoJobOutputReference {
	_init_.Initialize()

	if err := validateNewDataprocWorkflowTemplateJobsPrestoJobOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataprocWorkflowTemplate.DataprocWorkflowTemplateJobsPrestoJobOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataprocWorkflowTemplateJobsPrestoJobOutputReference_Override(d DataprocWorkflowTemplateJobsPrestoJobOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataprocWorkflowTemplate.DataprocWorkflowTemplateJobsPrestoJobOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference)SetClientTags(val *[]*string) {
	if err := j.validateSetClientTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientTags",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference)SetContinueOnFailure(val interface{}) {
	if err := j.validateSetContinueOnFailureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"continueOnFailure",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference)SetInternalValue(val *DataprocWorkflowTemplateJobsPrestoJob) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference)SetOutputFormat(val *string) {
	if err := j.validateSetOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputFormat",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference)SetProperties(val *map[string]*string) {
	if err := j.validateSetPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"properties",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference)SetQueryFileUri(val *string) {
	if err := j.validateSetQueryFileUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryFileUri",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference) PutLoggingConfig(value *DataprocWorkflowTemplateJobsPrestoJobLoggingConfig) {
	if err := d.validatePutLoggingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLoggingConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference) PutQueryList(value *DataprocWorkflowTemplateJobsPrestoJobQueryListStruct) {
	if err := d.validatePutQueryListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putQueryList",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference) ResetClientTags() {
	_jsii_.InvokeVoid(
		d,
		"resetClientTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference) ResetContinueOnFailure() {
	_jsii_.InvokeVoid(
		d,
		"resetContinueOnFailure",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference) ResetLoggingConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetLoggingConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference) ResetOutputFormat() {
	_jsii_.InvokeVoid(
		d,
		"resetOutputFormat",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference) ResetProperties() {
	_jsii_.InvokeVoid(
		d,
		"resetProperties",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference) ResetQueryFileUri() {
	_jsii_.InvokeVoid(
		d,
		"resetQueryFileUri",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference) ResetQueryList() {
	_jsii_.InvokeVoid(
		d,
		"resetQueryList",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataprocWorkflowTemplateJobsPrestoJobOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

