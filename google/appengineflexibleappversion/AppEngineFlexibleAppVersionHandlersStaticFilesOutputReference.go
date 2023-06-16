package appengineflexibleappversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v8/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v8/appengineflexibleappversion/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference interface {
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
	InternalValue() *AppEngineFlexibleAppVersionHandlersStaticFiles
	SetInternalValue(val *AppEngineFlexibleAppVersionHandlersStaticFiles)
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

// The jsii proxy struct for AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference
type jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) ApplicationReadable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applicationReadable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) ApplicationReadableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applicationReadableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) Expiration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) ExpirationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expirationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) HttpHeaders() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"httpHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) HttpHeadersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"httpHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) InternalValue() *AppEngineFlexibleAppVersionHandlersStaticFiles {
	var returns *AppEngineFlexibleAppVersionHandlersStaticFiles
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) MimeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mimeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) MimeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mimeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) RequireMatchingFile() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireMatchingFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) RequireMatchingFileInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireMatchingFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) UploadPathRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uploadPathRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) UploadPathRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uploadPathRegexInput",
		&returns,
	)
	return returns
}


func NewAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference {
	_init_.Initialize()

	if err := validateNewAppEngineFlexibleAppVersionHandlersStaticFilesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.appEngineFlexibleAppVersion.AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppEngineFlexibleAppVersionHandlersStaticFilesOutputReference_Override(a AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.appEngineFlexibleAppVersion.AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference)SetApplicationReadable(val interface{}) {
	if err := j.validateSetApplicationReadableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationReadable",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference)SetExpiration(val *string) {
	if err := j.validateSetExpirationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expiration",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference)SetHttpHeaders(val *map[string]*string) {
	if err := j.validateSetHttpHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpHeaders",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference)SetInternalValue(val *AppEngineFlexibleAppVersionHandlersStaticFiles) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference)SetMimeType(val *string) {
	if err := j.validateSetMimeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mimeType",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference)SetRequireMatchingFile(val interface{}) {
	if err := j.validateSetRequireMatchingFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireMatchingFile",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference)SetUploadPathRegex(val *string) {
	if err := j.validateSetUploadPathRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uploadPathRegex",
		val,
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) ResetApplicationReadable() {
	_jsii_.InvokeVoid(
		a,
		"resetApplicationReadable",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) ResetExpiration() {
	_jsii_.InvokeVoid(
		a,
		"resetExpiration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) ResetHttpHeaders() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpHeaders",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) ResetMimeType() {
	_jsii_.InvokeVoid(
		a,
		"resetMimeType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		a,
		"resetPath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) ResetRequireMatchingFile() {
	_jsii_.InvokeVoid(
		a,
		"resetRequireMatchingFile",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) ResetUploadPathRegex() {
	_jsii_.InvokeVoid(
		a,
		"resetUploadPathRegex",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

