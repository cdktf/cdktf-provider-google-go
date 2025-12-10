// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudbuildtrigger

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/cloudbuildtrigger/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudbuildTriggerBuildArtifactsOutputReference interface {
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
	Images() *[]*string
	SetImages(val *[]*string)
	ImagesInput() *[]*string
	InternalValue() *CloudbuildTriggerBuildArtifacts
	SetInternalValue(val *CloudbuildTriggerBuildArtifacts)
	MavenArtifacts() CloudbuildTriggerBuildArtifactsMavenArtifactsList
	MavenArtifactsInput() interface{}
	NpmPackages() CloudbuildTriggerBuildArtifactsNpmPackagesList
	NpmPackagesInput() interface{}
	Objects() CloudbuildTriggerBuildArtifactsObjectsOutputReference
	ObjectsInput() *CloudbuildTriggerBuildArtifactsObjects
	PythonPackages() CloudbuildTriggerBuildArtifactsPythonPackagesList
	PythonPackagesInput() interface{}
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutMavenArtifacts(value interface{})
	PutNpmPackages(value interface{})
	PutObjects(value *CloudbuildTriggerBuildArtifactsObjects)
	PutPythonPackages(value interface{})
	ResetImages()
	ResetMavenArtifacts()
	ResetNpmPackages()
	ResetObjects()
	ResetPythonPackages()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudbuildTriggerBuildArtifactsOutputReference
type jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference) Images() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"images",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference) ImagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"imagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference) InternalValue() *CloudbuildTriggerBuildArtifacts {
	var returns *CloudbuildTriggerBuildArtifacts
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference) MavenArtifacts() CloudbuildTriggerBuildArtifactsMavenArtifactsList {
	var returns CloudbuildTriggerBuildArtifactsMavenArtifactsList
	_jsii_.Get(
		j,
		"mavenArtifacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference) MavenArtifactsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mavenArtifactsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference) NpmPackages() CloudbuildTriggerBuildArtifactsNpmPackagesList {
	var returns CloudbuildTriggerBuildArtifactsNpmPackagesList
	_jsii_.Get(
		j,
		"npmPackages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference) NpmPackagesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"npmPackagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference) Objects() CloudbuildTriggerBuildArtifactsObjectsOutputReference {
	var returns CloudbuildTriggerBuildArtifactsObjectsOutputReference
	_jsii_.Get(
		j,
		"objects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference) ObjectsInput() *CloudbuildTriggerBuildArtifactsObjects {
	var returns *CloudbuildTriggerBuildArtifactsObjects
	_jsii_.Get(
		j,
		"objectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference) PythonPackages() CloudbuildTriggerBuildArtifactsPythonPackagesList {
	var returns CloudbuildTriggerBuildArtifactsPythonPackagesList
	_jsii_.Get(
		j,
		"pythonPackages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference) PythonPackagesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pythonPackagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudbuildTriggerBuildArtifactsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CloudbuildTriggerBuildArtifactsOutputReference {
	_init_.Initialize()

	if err := validateNewCloudbuildTriggerBuildArtifactsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cloudbuildTrigger.CloudbuildTriggerBuildArtifactsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudbuildTriggerBuildArtifactsOutputReference_Override(c CloudbuildTriggerBuildArtifactsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cloudbuildTrigger.CloudbuildTriggerBuildArtifactsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference)SetImages(val *[]*string) {
	if err := j.validateSetImagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"images",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference)SetInternalValue(val *CloudbuildTriggerBuildArtifacts) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference) PutMavenArtifacts(value interface{}) {
	if err := c.validatePutMavenArtifactsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMavenArtifacts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference) PutNpmPackages(value interface{}) {
	if err := c.validatePutNpmPackagesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNpmPackages",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference) PutObjects(value *CloudbuildTriggerBuildArtifactsObjects) {
	if err := c.validatePutObjectsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putObjects",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference) PutPythonPackages(value interface{}) {
	if err := c.validatePutPythonPackagesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPythonPackages",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference) ResetImages() {
	_jsii_.InvokeVoid(
		c,
		"resetImages",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference) ResetMavenArtifacts() {
	_jsii_.InvokeVoid(
		c,
		"resetMavenArtifacts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference) ResetNpmPackages() {
	_jsii_.InvokeVoid(
		c,
		"resetNpmPackages",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference) ResetObjects() {
	_jsii_.InvokeVoid(
		c,
		"resetObjects",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference) ResetPythonPackages() {
	_jsii_.InvokeVoid(
		c,
		"resetPythonPackages",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTriggerBuildArtifactsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

