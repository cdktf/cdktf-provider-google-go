// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appengineflexibleappversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/appengineflexibleappversion/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppEngineFlexibleAppVersionDeploymentOutputReference interface {
	cdktf.ComplexObject
	CloudBuildOptions() AppEngineFlexibleAppVersionDeploymentCloudBuildOptionsOutputReference
	CloudBuildOptionsInput() *AppEngineFlexibleAppVersionDeploymentCloudBuildOptions
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
	Container() AppEngineFlexibleAppVersionDeploymentContainerOutputReference
	ContainerInput() *AppEngineFlexibleAppVersionDeploymentContainer
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Files() AppEngineFlexibleAppVersionDeploymentFilesList
	FilesInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *AppEngineFlexibleAppVersionDeployment
	SetInternalValue(val *AppEngineFlexibleAppVersionDeployment)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Zip() AppEngineFlexibleAppVersionDeploymentZipOutputReference
	ZipInput() *AppEngineFlexibleAppVersionDeploymentZip
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
	PutCloudBuildOptions(value *AppEngineFlexibleAppVersionDeploymentCloudBuildOptions)
	PutContainer(value *AppEngineFlexibleAppVersionDeploymentContainer)
	PutFiles(value interface{})
	PutZip(value *AppEngineFlexibleAppVersionDeploymentZip)
	ResetCloudBuildOptions()
	ResetContainer()
	ResetFiles()
	ResetZip()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppEngineFlexibleAppVersionDeploymentOutputReference
type jsiiProxy_AppEngineFlexibleAppVersionDeploymentOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionDeploymentOutputReference) CloudBuildOptions() AppEngineFlexibleAppVersionDeploymentCloudBuildOptionsOutputReference {
	var returns AppEngineFlexibleAppVersionDeploymentCloudBuildOptionsOutputReference
	_jsii_.Get(
		j,
		"cloudBuildOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionDeploymentOutputReference) CloudBuildOptionsInput() *AppEngineFlexibleAppVersionDeploymentCloudBuildOptions {
	var returns *AppEngineFlexibleAppVersionDeploymentCloudBuildOptions
	_jsii_.Get(
		j,
		"cloudBuildOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionDeploymentOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionDeploymentOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionDeploymentOutputReference) Container() AppEngineFlexibleAppVersionDeploymentContainerOutputReference {
	var returns AppEngineFlexibleAppVersionDeploymentContainerOutputReference
	_jsii_.Get(
		j,
		"container",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionDeploymentOutputReference) ContainerInput() *AppEngineFlexibleAppVersionDeploymentContainer {
	var returns *AppEngineFlexibleAppVersionDeploymentContainer
	_jsii_.Get(
		j,
		"containerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionDeploymentOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionDeploymentOutputReference) Files() AppEngineFlexibleAppVersionDeploymentFilesList {
	var returns AppEngineFlexibleAppVersionDeploymentFilesList
	_jsii_.Get(
		j,
		"files",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionDeploymentOutputReference) FilesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionDeploymentOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionDeploymentOutputReference) InternalValue() *AppEngineFlexibleAppVersionDeployment {
	var returns *AppEngineFlexibleAppVersionDeployment
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionDeploymentOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionDeploymentOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionDeploymentOutputReference) Zip() AppEngineFlexibleAppVersionDeploymentZipOutputReference {
	var returns AppEngineFlexibleAppVersionDeploymentZipOutputReference
	_jsii_.Get(
		j,
		"zip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionDeploymentOutputReference) ZipInput() *AppEngineFlexibleAppVersionDeploymentZip {
	var returns *AppEngineFlexibleAppVersionDeploymentZip
	_jsii_.Get(
		j,
		"zipInput",
		&returns,
	)
	return returns
}


func NewAppEngineFlexibleAppVersionDeploymentOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppEngineFlexibleAppVersionDeploymentOutputReference {
	_init_.Initialize()

	if err := validateNewAppEngineFlexibleAppVersionDeploymentOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppEngineFlexibleAppVersionDeploymentOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.appEngineFlexibleAppVersion.AppEngineFlexibleAppVersionDeploymentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppEngineFlexibleAppVersionDeploymentOutputReference_Override(a AppEngineFlexibleAppVersionDeploymentOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.appEngineFlexibleAppVersion.AppEngineFlexibleAppVersionDeploymentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionDeploymentOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionDeploymentOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionDeploymentOutputReference)SetInternalValue(val *AppEngineFlexibleAppVersionDeployment) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionDeploymentOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionDeploymentOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionDeploymentOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionDeploymentOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionDeploymentOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionDeploymentOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionDeploymentOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionDeploymentOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionDeploymentOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionDeploymentOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionDeploymentOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionDeploymentOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionDeploymentOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionDeploymentOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionDeploymentOutputReference) PutCloudBuildOptions(value *AppEngineFlexibleAppVersionDeploymentCloudBuildOptions) {
	if err := a.validatePutCloudBuildOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCloudBuildOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionDeploymentOutputReference) PutContainer(value *AppEngineFlexibleAppVersionDeploymentContainer) {
	if err := a.validatePutContainerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putContainer",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionDeploymentOutputReference) PutFiles(value interface{}) {
	if err := a.validatePutFilesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFiles",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionDeploymentOutputReference) PutZip(value *AppEngineFlexibleAppVersionDeploymentZip) {
	if err := a.validatePutZipParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putZip",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionDeploymentOutputReference) ResetCloudBuildOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetCloudBuildOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionDeploymentOutputReference) ResetContainer() {
	_jsii_.InvokeVoid(
		a,
		"resetContainer",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionDeploymentOutputReference) ResetFiles() {
	_jsii_.InvokeVoid(
		a,
		"resetFiles",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionDeploymentOutputReference) ResetZip() {
	_jsii_.InvokeVoid(
		a,
		"resetZip",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionDeploymentOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionDeploymentOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

