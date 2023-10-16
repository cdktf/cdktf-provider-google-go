// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataplextask

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v11/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v11/dataplextask/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataplexTaskSparkOutputReference interface {
	cdktf.ComplexObject
	ArchiveUris() *[]*string
	SetArchiveUris(val *[]*string)
	ArchiveUrisInput() *[]*string
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
	FileUris() *[]*string
	SetFileUris(val *[]*string)
	FileUrisInput() *[]*string
	// Experimental.
	Fqn() *string
	InfrastructureSpec() DataplexTaskSparkInfrastructureSpecOutputReference
	InfrastructureSpecInput() *DataplexTaskSparkInfrastructureSpec
	InternalValue() *DataplexTaskSpark
	SetInternalValue(val *DataplexTaskSpark)
	MainClass() *string
	SetMainClass(val *string)
	MainClassInput() *string
	MainJarFileUri() *string
	SetMainJarFileUri(val *string)
	MainJarFileUriInput() *string
	PythonScriptFile() *string
	SetPythonScriptFile(val *string)
	PythonScriptFileInput() *string
	SqlScript() *string
	SetSqlScript(val *string)
	SqlScriptFile() *string
	SetSqlScriptFile(val *string)
	SqlScriptFileInput() *string
	SqlScriptInput() *string
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
	PutInfrastructureSpec(value *DataplexTaskSparkInfrastructureSpec)
	ResetArchiveUris()
	ResetFileUris()
	ResetInfrastructureSpec()
	ResetMainClass()
	ResetMainJarFileUri()
	ResetPythonScriptFile()
	ResetSqlScript()
	ResetSqlScriptFile()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataplexTaskSparkOutputReference
type jsiiProxy_DataplexTaskSparkOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataplexTaskSparkOutputReference) ArchiveUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"archiveUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkOutputReference) ArchiveUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"archiveUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkOutputReference) FileUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"fileUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkOutputReference) FileUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"fileUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkOutputReference) InfrastructureSpec() DataplexTaskSparkInfrastructureSpecOutputReference {
	var returns DataplexTaskSparkInfrastructureSpecOutputReference
	_jsii_.Get(
		j,
		"infrastructureSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkOutputReference) InfrastructureSpecInput() *DataplexTaskSparkInfrastructureSpec {
	var returns *DataplexTaskSparkInfrastructureSpec
	_jsii_.Get(
		j,
		"infrastructureSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkOutputReference) InternalValue() *DataplexTaskSpark {
	var returns *DataplexTaskSpark
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkOutputReference) MainClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mainClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkOutputReference) MainClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mainClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkOutputReference) MainJarFileUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mainJarFileUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkOutputReference) MainJarFileUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mainJarFileUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkOutputReference) PythonScriptFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pythonScriptFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkOutputReference) PythonScriptFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pythonScriptFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkOutputReference) SqlScript() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkOutputReference) SqlScriptFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlScriptFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkOutputReference) SqlScriptFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlScriptFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkOutputReference) SqlScriptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlScriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexTaskSparkOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataplexTaskSparkOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataplexTaskSparkOutputReference {
	_init_.Initialize()

	if err := validateNewDataplexTaskSparkOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataplexTaskSparkOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataplexTask.DataplexTaskSparkOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataplexTaskSparkOutputReference_Override(d DataplexTaskSparkOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataplexTask.DataplexTaskSparkOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataplexTaskSparkOutputReference)SetArchiveUris(val *[]*string) {
	if err := j.validateSetArchiveUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"archiveUris",
		val,
	)
}

func (j *jsiiProxy_DataplexTaskSparkOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataplexTaskSparkOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataplexTaskSparkOutputReference)SetFileUris(val *[]*string) {
	if err := j.validateSetFileUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileUris",
		val,
	)
}

func (j *jsiiProxy_DataplexTaskSparkOutputReference)SetInternalValue(val *DataplexTaskSpark) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataplexTaskSparkOutputReference)SetMainClass(val *string) {
	if err := j.validateSetMainClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mainClass",
		val,
	)
}

func (j *jsiiProxy_DataplexTaskSparkOutputReference)SetMainJarFileUri(val *string) {
	if err := j.validateSetMainJarFileUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mainJarFileUri",
		val,
	)
}

func (j *jsiiProxy_DataplexTaskSparkOutputReference)SetPythonScriptFile(val *string) {
	if err := j.validateSetPythonScriptFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pythonScriptFile",
		val,
	)
}

func (j *jsiiProxy_DataplexTaskSparkOutputReference)SetSqlScript(val *string) {
	if err := j.validateSetSqlScriptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqlScript",
		val,
	)
}

func (j *jsiiProxy_DataplexTaskSparkOutputReference)SetSqlScriptFile(val *string) {
	if err := j.validateSetSqlScriptFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqlScriptFile",
		val,
	)
}

func (j *jsiiProxy_DataplexTaskSparkOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataplexTaskSparkOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataplexTaskSparkOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataplexTaskSparkOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataplexTaskSparkOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataplexTaskSparkOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataplexTaskSparkOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataplexTaskSparkOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataplexTaskSparkOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataplexTaskSparkOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataplexTaskSparkOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataplexTaskSparkOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataplexTaskSparkOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataplexTaskSparkOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataplexTaskSparkOutputReference) PutInfrastructureSpec(value *DataplexTaskSparkInfrastructureSpec) {
	if err := d.validatePutInfrastructureSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putInfrastructureSpec",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataplexTaskSparkOutputReference) ResetArchiveUris() {
	_jsii_.InvokeVoid(
		d,
		"resetArchiveUris",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexTaskSparkOutputReference) ResetFileUris() {
	_jsii_.InvokeVoid(
		d,
		"resetFileUris",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexTaskSparkOutputReference) ResetInfrastructureSpec() {
	_jsii_.InvokeVoid(
		d,
		"resetInfrastructureSpec",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexTaskSparkOutputReference) ResetMainClass() {
	_jsii_.InvokeVoid(
		d,
		"resetMainClass",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexTaskSparkOutputReference) ResetMainJarFileUri() {
	_jsii_.InvokeVoid(
		d,
		"resetMainJarFileUri",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexTaskSparkOutputReference) ResetPythonScriptFile() {
	_jsii_.InvokeVoid(
		d,
		"resetPythonScriptFile",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexTaskSparkOutputReference) ResetSqlScript() {
	_jsii_.InvokeVoid(
		d,
		"resetSqlScript",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexTaskSparkOutputReference) ResetSqlScriptFile() {
	_jsii_.InvokeVoid(
		d,
		"resetSqlScriptFile",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexTaskSparkOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataplexTaskSparkOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

