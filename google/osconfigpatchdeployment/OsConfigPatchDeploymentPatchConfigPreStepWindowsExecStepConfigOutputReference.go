// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigpatchdeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/osconfigpatchdeployment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference interface {
	cdktf.ComplexObject
	AllowedSuccessCodes() *[]*float64
	SetAllowedSuccessCodes(val *[]*float64)
	AllowedSuccessCodesInput() *[]*float64
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
	GcsObject() OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObjectOutputReference
	GcsObjectInput() *OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject
	InternalValue() *OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfig
	SetInternalValue(val *OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfig)
	Interpreter() *string
	SetInterpreter(val *string)
	InterpreterInput() *string
	LocalPath() *string
	SetLocalPath(val *string)
	LocalPathInput() *string
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
	PutGcsObject(value *OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject)
	ResetAllowedSuccessCodes()
	ResetGcsObject()
	ResetInterpreter()
	ResetLocalPath()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference
type jsiiProxy_OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference) AllowedSuccessCodes() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"allowedSuccessCodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference) AllowedSuccessCodesInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"allowedSuccessCodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference) GcsObject() OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObjectOutputReference {
	var returns OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObjectOutputReference
	_jsii_.Get(
		j,
		"gcsObject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference) GcsObjectInput() *OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject {
	var returns *OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject
	_jsii_.Get(
		j,
		"gcsObjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference) InternalValue() *OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfig {
	var returns *OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference) Interpreter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interpreter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference) InterpreterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interpreterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference) LocalPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference) LocalPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference {
	_init_.Initialize()

	if err := validateNewOsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.osConfigPatchDeployment.OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference_Override(o OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.osConfigPatchDeployment.OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference)SetAllowedSuccessCodes(val *[]*float64) {
	if err := j.validateSetAllowedSuccessCodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedSuccessCodes",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference)SetInternalValue(val *OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference)SetInterpreter(val *string) {
	if err := j.validateSetInterpreterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interpreter",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference)SetLocalPath(val *string) {
	if err := j.validateSetLocalPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localPath",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference) PutGcsObject(value *OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject) {
	if err := o.validatePutGcsObjectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putGcsObject",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference) ResetAllowedSuccessCodes() {
	_jsii_.InvokeVoid(
		o,
		"resetAllowedSuccessCodes",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference) ResetGcsObject() {
	_jsii_.InvokeVoid(
		o,
		"resetGcsObject",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference) ResetInterpreter() {
	_jsii_.InvokeVoid(
		o,
		"resetInterpreter",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference) ResetLocalPath() {
	_jsii_.InvokeVoid(
		o,
		"resetLocalPath",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

