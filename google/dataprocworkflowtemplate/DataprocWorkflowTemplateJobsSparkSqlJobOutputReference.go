// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocworkflowtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v12/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v12/dataprocworkflowtemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataprocWorkflowTemplateJobsSparkSqlJobOutputReference interface {
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
	InternalValue() *DataprocWorkflowTemplateJobsSparkSqlJob
	SetInternalValue(val *DataprocWorkflowTemplateJobsSparkSqlJob)
	JarFileUris() *[]*string
	SetJarFileUris(val *[]*string)
	JarFileUrisInput() *[]*string
	LoggingConfig() DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfigOutputReference
	LoggingConfigInput() *DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfig
	Properties() *map[string]*string
	SetProperties(val *map[string]*string)
	PropertiesInput() *map[string]*string
	QueryFileUri() *string
	SetQueryFileUri(val *string)
	QueryFileUriInput() *string
	QueryList() DataprocWorkflowTemplateJobsSparkSqlJobQueryListStructOutputReference
	QueryListInput() *DataprocWorkflowTemplateJobsSparkSqlJobQueryListStruct
	ScriptVariables() *map[string]*string
	SetScriptVariables(val *map[string]*string)
	ScriptVariablesInput() *map[string]*string
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
	PutLoggingConfig(value *DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfig)
	PutQueryList(value *DataprocWorkflowTemplateJobsSparkSqlJobQueryListStruct)
	ResetJarFileUris()
	ResetLoggingConfig()
	ResetProperties()
	ResetQueryFileUri()
	ResetQueryList()
	ResetScriptVariables()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataprocWorkflowTemplateJobsSparkSqlJobOutputReference
type jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference) InternalValue() *DataprocWorkflowTemplateJobsSparkSqlJob {
	var returns *DataprocWorkflowTemplateJobsSparkSqlJob
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference) JarFileUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jarFileUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference) JarFileUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jarFileUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference) LoggingConfig() DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfigOutputReference {
	var returns DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfigOutputReference
	_jsii_.Get(
		j,
		"loggingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference) LoggingConfigInput() *DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfig {
	var returns *DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfig
	_jsii_.Get(
		j,
		"loggingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference) Properties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference) PropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"propertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference) QueryFileUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryFileUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference) QueryFileUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryFileUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference) QueryList() DataprocWorkflowTemplateJobsSparkSqlJobQueryListStructOutputReference {
	var returns DataprocWorkflowTemplateJobsSparkSqlJobQueryListStructOutputReference
	_jsii_.Get(
		j,
		"queryList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference) QueryListInput() *DataprocWorkflowTemplateJobsSparkSqlJobQueryListStruct {
	var returns *DataprocWorkflowTemplateJobsSparkSqlJobQueryListStruct
	_jsii_.Get(
		j,
		"queryListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference) ScriptVariables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"scriptVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference) ScriptVariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"scriptVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataprocWorkflowTemplateJobsSparkSqlJobOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataprocWorkflowTemplateJobsSparkSqlJobOutputReference {
	_init_.Initialize()

	if err := validateNewDataprocWorkflowTemplateJobsSparkSqlJobOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataprocWorkflowTemplate.DataprocWorkflowTemplateJobsSparkSqlJobOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataprocWorkflowTemplateJobsSparkSqlJobOutputReference_Override(d DataprocWorkflowTemplateJobsSparkSqlJobOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataprocWorkflowTemplate.DataprocWorkflowTemplateJobsSparkSqlJobOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference)SetInternalValue(val *DataprocWorkflowTemplateJobsSparkSqlJob) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference)SetJarFileUris(val *[]*string) {
	if err := j.validateSetJarFileUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jarFileUris",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference)SetProperties(val *map[string]*string) {
	if err := j.validateSetPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"properties",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference)SetQueryFileUri(val *string) {
	if err := j.validateSetQueryFileUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryFileUri",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference)SetScriptVariables(val *map[string]*string) {
	if err := j.validateSetScriptVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scriptVariables",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference) PutLoggingConfig(value *DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfig) {
	if err := d.validatePutLoggingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLoggingConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference) PutQueryList(value *DataprocWorkflowTemplateJobsSparkSqlJobQueryListStruct) {
	if err := d.validatePutQueryListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putQueryList",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference) ResetJarFileUris() {
	_jsii_.InvokeVoid(
		d,
		"resetJarFileUris",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference) ResetLoggingConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetLoggingConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference) ResetProperties() {
	_jsii_.InvokeVoid(
		d,
		"resetProperties",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference) ResetQueryFileUri() {
	_jsii_.InvokeVoid(
		d,
		"resetQueryFileUri",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference) ResetQueryList() {
	_jsii_.InvokeVoid(
		d,
		"resetQueryList",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference) ResetScriptVariables() {
	_jsii_.InvokeVoid(
		d,
		"resetScriptVariables",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataprocWorkflowTemplateJobsSparkSqlJobOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

