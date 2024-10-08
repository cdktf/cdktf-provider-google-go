// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocbatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/dataprocbatch/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataprocBatchPysparkBatchOutputReference interface {
	cdktf.ComplexObject
	ArchiveUris() *[]*string
	SetArchiveUris(val *[]*string)
	ArchiveUrisInput() *[]*string
	Args() *[]*string
	SetArgs(val *[]*string)
	ArgsInput() *[]*string
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
	InternalValue() *DataprocBatchPysparkBatch
	SetInternalValue(val *DataprocBatchPysparkBatch)
	JarFileUris() *[]*string
	SetJarFileUris(val *[]*string)
	JarFileUrisInput() *[]*string
	MainPythonFileUri() *string
	SetMainPythonFileUri(val *string)
	MainPythonFileUriInput() *string
	PythonFileUris() *[]*string
	SetPythonFileUris(val *[]*string)
	PythonFileUrisInput() *[]*string
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
	ResetArchiveUris()
	ResetArgs()
	ResetFileUris()
	ResetJarFileUris()
	ResetMainPythonFileUri()
	ResetPythonFileUris()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataprocBatchPysparkBatchOutputReference
type jsiiProxy_DataprocBatchPysparkBatchOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataprocBatchPysparkBatchOutputReference) ArchiveUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"archiveUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocBatchPysparkBatchOutputReference) ArchiveUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"archiveUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocBatchPysparkBatchOutputReference) Args() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"args",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocBatchPysparkBatchOutputReference) ArgsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"argsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocBatchPysparkBatchOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocBatchPysparkBatchOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocBatchPysparkBatchOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocBatchPysparkBatchOutputReference) FileUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"fileUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocBatchPysparkBatchOutputReference) FileUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"fileUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocBatchPysparkBatchOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocBatchPysparkBatchOutputReference) InternalValue() *DataprocBatchPysparkBatch {
	var returns *DataprocBatchPysparkBatch
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocBatchPysparkBatchOutputReference) JarFileUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jarFileUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocBatchPysparkBatchOutputReference) JarFileUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jarFileUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocBatchPysparkBatchOutputReference) MainPythonFileUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mainPythonFileUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocBatchPysparkBatchOutputReference) MainPythonFileUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mainPythonFileUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocBatchPysparkBatchOutputReference) PythonFileUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pythonFileUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocBatchPysparkBatchOutputReference) PythonFileUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pythonFileUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocBatchPysparkBatchOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocBatchPysparkBatchOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataprocBatchPysparkBatchOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataprocBatchPysparkBatchOutputReference {
	_init_.Initialize()

	if err := validateNewDataprocBatchPysparkBatchOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataprocBatchPysparkBatchOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataprocBatch.DataprocBatchPysparkBatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataprocBatchPysparkBatchOutputReference_Override(d DataprocBatchPysparkBatchOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataprocBatch.DataprocBatchPysparkBatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataprocBatchPysparkBatchOutputReference)SetArchiveUris(val *[]*string) {
	if err := j.validateSetArchiveUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"archiveUris",
		val,
	)
}

func (j *jsiiProxy_DataprocBatchPysparkBatchOutputReference)SetArgs(val *[]*string) {
	if err := j.validateSetArgsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"args",
		val,
	)
}

func (j *jsiiProxy_DataprocBatchPysparkBatchOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataprocBatchPysparkBatchOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataprocBatchPysparkBatchOutputReference)SetFileUris(val *[]*string) {
	if err := j.validateSetFileUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileUris",
		val,
	)
}

func (j *jsiiProxy_DataprocBatchPysparkBatchOutputReference)SetInternalValue(val *DataprocBatchPysparkBatch) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataprocBatchPysparkBatchOutputReference)SetJarFileUris(val *[]*string) {
	if err := j.validateSetJarFileUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jarFileUris",
		val,
	)
}

func (j *jsiiProxy_DataprocBatchPysparkBatchOutputReference)SetMainPythonFileUri(val *string) {
	if err := j.validateSetMainPythonFileUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mainPythonFileUri",
		val,
	)
}

func (j *jsiiProxy_DataprocBatchPysparkBatchOutputReference)SetPythonFileUris(val *[]*string) {
	if err := j.validateSetPythonFileUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pythonFileUris",
		val,
	)
}

func (j *jsiiProxy_DataprocBatchPysparkBatchOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataprocBatchPysparkBatchOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataprocBatchPysparkBatchOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocBatchPysparkBatchOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataprocBatchPysparkBatchOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataprocBatchPysparkBatchOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataprocBatchPysparkBatchOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataprocBatchPysparkBatchOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataprocBatchPysparkBatchOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataprocBatchPysparkBatchOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataprocBatchPysparkBatchOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataprocBatchPysparkBatchOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataprocBatchPysparkBatchOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocBatchPysparkBatchOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataprocBatchPysparkBatchOutputReference) ResetArchiveUris() {
	_jsii_.InvokeVoid(
		d,
		"resetArchiveUris",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocBatchPysparkBatchOutputReference) ResetArgs() {
	_jsii_.InvokeVoid(
		d,
		"resetArgs",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocBatchPysparkBatchOutputReference) ResetFileUris() {
	_jsii_.InvokeVoid(
		d,
		"resetFileUris",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocBatchPysparkBatchOutputReference) ResetJarFileUris() {
	_jsii_.InvokeVoid(
		d,
		"resetJarFileUris",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocBatchPysparkBatchOutputReference) ResetMainPythonFileUri() {
	_jsii_.InvokeVoid(
		d,
		"resetMainPythonFileUri",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocBatchPysparkBatchOutputReference) ResetPythonFileUris() {
	_jsii_.InvokeVoid(
		d,
		"resetPythonFileUris",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocBatchPysparkBatchOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataprocBatchPysparkBatchOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

