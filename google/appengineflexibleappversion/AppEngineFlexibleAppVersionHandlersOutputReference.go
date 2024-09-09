// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appengineflexibleappversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/appengineflexibleappversion/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppEngineFlexibleAppVersionHandlersOutputReference interface {
	cdktf.ComplexObject
	AuthFailAction() *string
	SetAuthFailAction(val *string)
	AuthFailActionInput() *string
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Login() *string
	SetLogin(val *string)
	LoginInput() *string
	RedirectHttpResponseCode() *string
	SetRedirectHttpResponseCode(val *string)
	RedirectHttpResponseCodeInput() *string
	Script() AppEngineFlexibleAppVersionHandlersScriptOutputReference
	ScriptInput() *AppEngineFlexibleAppVersionHandlersScript
	SecurityLevel() *string
	SetSecurityLevel(val *string)
	SecurityLevelInput() *string
	StaticFiles() AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference
	StaticFilesInput() *AppEngineFlexibleAppVersionHandlersStaticFiles
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UrlRegex() *string
	SetUrlRegex(val *string)
	UrlRegexInput() *string
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
	PutScript(value *AppEngineFlexibleAppVersionHandlersScript)
	PutStaticFiles(value *AppEngineFlexibleAppVersionHandlersStaticFiles)
	ResetAuthFailAction()
	ResetLogin()
	ResetRedirectHttpResponseCode()
	ResetScript()
	ResetSecurityLevel()
	ResetStaticFiles()
	ResetUrlRegex()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppEngineFlexibleAppVersionHandlersOutputReference
type jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference) AuthFailAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authFailAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference) AuthFailActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authFailActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference) Login() *string {
	var returns *string
	_jsii_.Get(
		j,
		"login",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference) LoginInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference) RedirectHttpResponseCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectHttpResponseCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference) RedirectHttpResponseCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redirectHttpResponseCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference) Script() AppEngineFlexibleAppVersionHandlersScriptOutputReference {
	var returns AppEngineFlexibleAppVersionHandlersScriptOutputReference
	_jsii_.Get(
		j,
		"script",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference) ScriptInput() *AppEngineFlexibleAppVersionHandlersScript {
	var returns *AppEngineFlexibleAppVersionHandlersScript
	_jsii_.Get(
		j,
		"scriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference) SecurityLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference) SecurityLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference) StaticFiles() AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference {
	var returns AppEngineFlexibleAppVersionHandlersStaticFilesOutputReference
	_jsii_.Get(
		j,
		"staticFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference) StaticFilesInput() *AppEngineFlexibleAppVersionHandlersStaticFiles {
	var returns *AppEngineFlexibleAppVersionHandlersStaticFiles
	_jsii_.Get(
		j,
		"staticFilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference) UrlRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference) UrlRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlRegexInput",
		&returns,
	)
	return returns
}


func NewAppEngineFlexibleAppVersionHandlersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AppEngineFlexibleAppVersionHandlersOutputReference {
	_init_.Initialize()

	if err := validateNewAppEngineFlexibleAppVersionHandlersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.appEngineFlexibleAppVersion.AppEngineFlexibleAppVersionHandlersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAppEngineFlexibleAppVersionHandlersOutputReference_Override(a AppEngineFlexibleAppVersionHandlersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.appEngineFlexibleAppVersion.AppEngineFlexibleAppVersionHandlersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference)SetAuthFailAction(val *string) {
	if err := j.validateSetAuthFailActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authFailAction",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference)SetLogin(val *string) {
	if err := j.validateSetLoginParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"login",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference)SetRedirectHttpResponseCode(val *string) {
	if err := j.validateSetRedirectHttpResponseCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redirectHttpResponseCode",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference)SetSecurityLevel(val *string) {
	if err := j.validateSetSecurityLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityLevel",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference)SetUrlRegex(val *string) {
	if err := j.validateSetUrlRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"urlRegex",
		val,
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference) PutScript(value *AppEngineFlexibleAppVersionHandlersScript) {
	if err := a.validatePutScriptParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putScript",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference) PutStaticFiles(value *AppEngineFlexibleAppVersionHandlersStaticFiles) {
	if err := a.validatePutStaticFilesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStaticFiles",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference) ResetAuthFailAction() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthFailAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference) ResetLogin() {
	_jsii_.InvokeVoid(
		a,
		"resetLogin",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference) ResetRedirectHttpResponseCode() {
	_jsii_.InvokeVoid(
		a,
		"resetRedirectHttpResponseCode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference) ResetScript() {
	_jsii_.InvokeVoid(
		a,
		"resetScript",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference) ResetSecurityLevel() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityLevel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference) ResetStaticFiles() {
	_jsii_.InvokeVoid(
		a,
		"resetStaticFiles",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference) ResetUrlRegex() {
	_jsii_.InvokeVoid(
		a,
		"resetUrlRegex",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersionHandlersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

