// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocgdcsparkapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/dataprocgdcsparkapplication/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference interface {
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
	InternalValue() *DataprocGdcSparkApplicationSparkSqlApplicationConfig
	SetInternalValue(val *DataprocGdcSparkApplicationSparkSqlApplicationConfig)
	JarFileUris() *[]*string
	SetJarFileUris(val *[]*string)
	JarFileUrisInput() *[]*string
	QueryFileUri() *string
	SetQueryFileUri(val *string)
	QueryFileUriInput() *string
	QueryList() DataprocGdcSparkApplicationSparkSqlApplicationConfigQueryListStructOutputReference
	QueryListInput() *DataprocGdcSparkApplicationSparkSqlApplicationConfigQueryListStruct
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
	PutQueryList(value *DataprocGdcSparkApplicationSparkSqlApplicationConfigQueryListStruct)
	ResetJarFileUris()
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

// The jsii proxy struct for DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference
type jsiiProxy_DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference) InternalValue() *DataprocGdcSparkApplicationSparkSqlApplicationConfig {
	var returns *DataprocGdcSparkApplicationSparkSqlApplicationConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference) JarFileUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jarFileUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference) JarFileUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jarFileUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference) QueryFileUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryFileUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference) QueryFileUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryFileUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference) QueryList() DataprocGdcSparkApplicationSparkSqlApplicationConfigQueryListStructOutputReference {
	var returns DataprocGdcSparkApplicationSparkSqlApplicationConfigQueryListStructOutputReference
	_jsii_.Get(
		j,
		"queryList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference) QueryListInput() *DataprocGdcSparkApplicationSparkSqlApplicationConfigQueryListStruct {
	var returns *DataprocGdcSparkApplicationSparkSqlApplicationConfigQueryListStruct
	_jsii_.Get(
		j,
		"queryListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference) ScriptVariables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"scriptVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference) ScriptVariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"scriptVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataprocGdcSparkApplication.DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference_Override(d DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataprocGdcSparkApplication.DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference)SetInternalValue(val *DataprocGdcSparkApplicationSparkSqlApplicationConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference)SetJarFileUris(val *[]*string) {
	if err := j.validateSetJarFileUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jarFileUris",
		val,
	)
}

func (j *jsiiProxy_DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference)SetQueryFileUri(val *string) {
	if err := j.validateSetQueryFileUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryFileUri",
		val,
	)
}

func (j *jsiiProxy_DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference)SetScriptVariables(val *map[string]*string) {
	if err := j.validateSetScriptVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scriptVariables",
		val,
	)
}

func (j *jsiiProxy_DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference) PutQueryList(value *DataprocGdcSparkApplicationSparkSqlApplicationConfigQueryListStruct) {
	if err := d.validatePutQueryListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putQueryList",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference) ResetJarFileUris() {
	_jsii_.InvokeVoid(
		d,
		"resetJarFileUris",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference) ResetQueryFileUri() {
	_jsii_.InvokeVoid(
		d,
		"resetQueryFileUri",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference) ResetQueryList() {
	_jsii_.InvokeVoid(
		d,
		"resetQueryList",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference) ResetScriptVariables() {
	_jsii_.InvokeVoid(
		d,
		"resetScriptVariables",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataprocGdcSparkApplicationSparkSqlApplicationConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

