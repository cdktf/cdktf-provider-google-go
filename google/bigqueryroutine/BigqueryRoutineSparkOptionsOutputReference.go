// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigqueryroutine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/bigqueryroutine/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BigqueryRoutineSparkOptionsOutputReference interface {
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
	Connection() *string
	SetConnection(val *string)
	ConnectionInput() *string
	ContainerImage() *string
	SetContainerImage(val *string)
	ContainerImageInput() *string
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
	InternalValue() *BigqueryRoutineSparkOptions
	SetInternalValue(val *BigqueryRoutineSparkOptions)
	JarUris() *[]*string
	SetJarUris(val *[]*string)
	JarUrisInput() *[]*string
	MainClass() *string
	SetMainClass(val *string)
	MainClassInput() *string
	MainFileUri() *string
	SetMainFileUri(val *string)
	MainFileUriInput() *string
	Properties() *map[string]*string
	SetProperties(val *map[string]*string)
	PropertiesInput() *map[string]*string
	PyFileUris() *[]*string
	SetPyFileUris(val *[]*string)
	PyFileUrisInput() *[]*string
	RuntimeVersion() *string
	SetRuntimeVersion(val *string)
	RuntimeVersionInput() *string
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
	ResetArchiveUris()
	ResetConnection()
	ResetContainerImage()
	ResetFileUris()
	ResetJarUris()
	ResetMainClass()
	ResetMainFileUri()
	ResetProperties()
	ResetPyFileUris()
	ResetRuntimeVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BigqueryRoutineSparkOptionsOutputReference
type jsiiProxy_BigqueryRoutineSparkOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) ArchiveUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"archiveUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) ArchiveUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"archiveUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) Connection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) ConnectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) ContainerImage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) ContainerImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) FileUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"fileUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) FileUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"fileUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) InternalValue() *BigqueryRoutineSparkOptions {
	var returns *BigqueryRoutineSparkOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) JarUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jarUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) JarUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jarUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) MainClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mainClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) MainClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mainClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) MainFileUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mainFileUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) MainFileUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mainFileUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) Properties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) PropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"propertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) PyFileUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pyFileUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) PyFileUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pyFileUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) RuntimeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) RuntimeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBigqueryRoutineSparkOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BigqueryRoutineSparkOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewBigqueryRoutineSparkOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BigqueryRoutineSparkOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.bigqueryRoutine.BigqueryRoutineSparkOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBigqueryRoutineSparkOptionsOutputReference_Override(b BigqueryRoutineSparkOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.bigqueryRoutine.BigqueryRoutineSparkOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference)SetArchiveUris(val *[]*string) {
	if err := j.validateSetArchiveUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"archiveUris",
		val,
	)
}

func (j *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference)SetConnection(val *string) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference)SetContainerImage(val *string) {
	if err := j.validateSetContainerImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerImage",
		val,
	)
}

func (j *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference)SetFileUris(val *[]*string) {
	if err := j.validateSetFileUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileUris",
		val,
	)
}

func (j *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference)SetInternalValue(val *BigqueryRoutineSparkOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference)SetJarUris(val *[]*string) {
	if err := j.validateSetJarUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jarUris",
		val,
	)
}

func (j *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference)SetMainClass(val *string) {
	if err := j.validateSetMainClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mainClass",
		val,
	)
}

func (j *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference)SetMainFileUri(val *string) {
	if err := j.validateSetMainFileUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mainFileUri",
		val,
	)
}

func (j *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference)SetProperties(val *map[string]*string) {
	if err := j.validateSetPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"properties",
		val,
	)
}

func (j *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference)SetPyFileUris(val *[]*string) {
	if err := j.validateSetPyFileUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pyFileUris",
		val,
	)
}

func (j *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference)SetRuntimeVersion(val *string) {
	if err := j.validateSetRuntimeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeVersion",
		val,
	)
}

func (j *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) ResetArchiveUris() {
	_jsii_.InvokeVoid(
		b,
		"resetArchiveUris",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) ResetConnection() {
	_jsii_.InvokeVoid(
		b,
		"resetConnection",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) ResetContainerImage() {
	_jsii_.InvokeVoid(
		b,
		"resetContainerImage",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) ResetFileUris() {
	_jsii_.InvokeVoid(
		b,
		"resetFileUris",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) ResetJarUris() {
	_jsii_.InvokeVoid(
		b,
		"resetJarUris",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) ResetMainClass() {
	_jsii_.InvokeVoid(
		b,
		"resetMainClass",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) ResetMainFileUri() {
	_jsii_.InvokeVoid(
		b,
		"resetMainFileUri",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) ResetProperties() {
	_jsii_.InvokeVoid(
		b,
		"resetProperties",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) ResetPyFileUris() {
	_jsii_.InvokeVoid(
		b,
		"resetPyFileUris",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) ResetRuntimeVersion() {
	_jsii_.InvokeVoid(
		b,
		"resetRuntimeVersion",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := b.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryRoutineSparkOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

