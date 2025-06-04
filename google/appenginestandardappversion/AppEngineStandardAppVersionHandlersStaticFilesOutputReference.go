// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appenginestandardappversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/appenginestandardappversion/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppEngineStandardAppVersionHandlersStaticFilesOutputReference interface {
	cdktf.ComplexObject
	ApplicationReadable() interface{}
	SetApplicationReadable(val interface{})
	ApplicationReadableInput() interface{}
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
	Expiration() *string
	SetExpiration(val *string)
	ExpirationInput() *string
	// Experimental.
	Fqn() *string
	HttpHeaders() *map[string]*string
	SetHttpHeaders(val *map[string]*string)
	HttpHeadersInput() *map[string]*string
	InternalValue() *AppEngineStandardAppVersionHandlersStaticFiles
	SetInternalValue(val *AppEngineStandardAppVersionHandlersStaticFiles)
	MimeType() *string
	SetMimeType(val *string)
	MimeTypeInput() *string
	Path() *string
	SetPath(val *string)
	PathInput() *string
	RequireMatchingFile() interface{}
	SetRequireMatchingFile(val interface{})
	RequireMatchingFileInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UploadPathRegex() *string
	SetUploadPathRegex(val *string)
	UploadPathRegexInput() *string
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
	ResetApplicationReadable()
	ResetExpiration()
	ResetHttpHeaders()
	ResetMimeType()
	ResetPath()
	ResetRequireMatchingFile()
	ResetUploadPathRegex()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppEngineStandardAppVersionHandlersStaticFilesOutputReference
type jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference) ApplicationReadable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applicationReadable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference) ApplicationReadableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applicationReadableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference) Expiration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference) ExpirationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expirationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference) HttpHeaders() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"httpHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference) HttpHeadersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"httpHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference) InternalValue() *AppEngineStandardAppVersionHandlersStaticFiles {
	var returns *AppEngineStandardAppVersionHandlersStaticFiles
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference) MimeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mimeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference) MimeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mimeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference) RequireMatchingFile() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireMatchingFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference) RequireMatchingFileInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireMatchingFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference) UploadPathRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uploadPathRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference) UploadPathRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uploadPathRegexInput",
		&returns,
	)
	return returns
}


func NewAppEngineStandardAppVersionHandlersStaticFilesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppEngineStandardAppVersionHandlersStaticFilesOutputReference {
	_init_.Initialize()

	if err := validateNewAppEngineStandardAppVersionHandlersStaticFilesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.appEngineStandardAppVersion.AppEngineStandardAppVersionHandlersStaticFilesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppEngineStandardAppVersionHandlersStaticFilesOutputReference_Override(a AppEngineStandardAppVersionHandlersStaticFilesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.appEngineStandardAppVersion.AppEngineStandardAppVersionHandlersStaticFilesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference)SetApplicationReadable(val interface{}) {
	if err := j.validateSetApplicationReadableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationReadable",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference)SetExpiration(val *string) {
	if err := j.validateSetExpirationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expiration",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference)SetHttpHeaders(val *map[string]*string) {
	if err := j.validateSetHttpHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpHeaders",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference)SetInternalValue(val *AppEngineStandardAppVersionHandlersStaticFiles) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference)SetMimeType(val *string) {
	if err := j.validateSetMimeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mimeType",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference)SetRequireMatchingFile(val interface{}) {
	if err := j.validateSetRequireMatchingFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireMatchingFile",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference)SetUploadPathRegex(val *string) {
	if err := j.validateSetUploadPathRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uploadPathRegex",
		val,
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference) ResetApplicationReadable() {
	_jsii_.InvokeVoid(
		a,
		"resetApplicationReadable",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference) ResetExpiration() {
	_jsii_.InvokeVoid(
		a,
		"resetExpiration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference) ResetHttpHeaders() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpHeaders",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference) ResetMimeType() {
	_jsii_.InvokeVoid(
		a,
		"resetMimeType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		a,
		"resetPath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference) ResetRequireMatchingFile() {
	_jsii_.InvokeVoid(
		a,
		"resetRequireMatchingFile",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference) ResetUploadPathRegex() {
	_jsii_.InvokeVoid(
		a,
		"resetUploadPathRegex",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AppEngineStandardAppVersionHandlersStaticFilesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

