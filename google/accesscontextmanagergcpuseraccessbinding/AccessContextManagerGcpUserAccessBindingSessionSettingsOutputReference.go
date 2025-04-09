// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesscontextmanagergcpuseraccessbinding

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/accesscontextmanagergcpuseraccessbinding/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference interface {
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
	InternalValue() *AccessContextManagerGcpUserAccessBindingSessionSettings
	SetInternalValue(val *AccessContextManagerGcpUserAccessBindingSessionSettings)
	MaxInactivity() *string
	SetMaxInactivity(val *string)
	MaxInactivityInput() *string
	SessionLength() *string
	SetSessionLength(val *string)
	SessionLengthEnabled() interface{}
	SetSessionLengthEnabled(val interface{})
	SessionLengthEnabledInput() interface{}
	SessionLengthInput() *string
	SessionReauthMethod() *string
	SetSessionReauthMethod(val *string)
	SessionReauthMethodInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UseOidcMaxAge() interface{}
	SetUseOidcMaxAge(val interface{})
	UseOidcMaxAgeInput() interface{}
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
	ResetMaxInactivity()
	ResetSessionLength()
	ResetSessionLengthEnabled()
	ResetSessionReauthMethod()
	ResetUseOidcMaxAge()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference
type jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference) InternalValue() *AccessContextManagerGcpUserAccessBindingSessionSettings {
	var returns *AccessContextManagerGcpUserAccessBindingSessionSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference) MaxInactivity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxInactivity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference) MaxInactivityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxInactivityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference) SessionLength() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference) SessionLengthEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sessionLengthEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference) SessionLengthEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sessionLengthEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference) SessionLengthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference) SessionReauthMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionReauthMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference) SessionReauthMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionReauthMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference) UseOidcMaxAge() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useOidcMaxAge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference) UseOidcMaxAgeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useOidcMaxAgeInput",
		&returns,
	)
	return returns
}


func NewAccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewAccessContextManagerGcpUserAccessBindingSessionSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.accessContextManagerGcpUserAccessBinding.AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference_Override(a AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.accessContextManagerGcpUserAccessBinding.AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference)SetInternalValue(val *AccessContextManagerGcpUserAccessBindingSessionSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference)SetMaxInactivity(val *string) {
	if err := j.validateSetMaxInactivityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxInactivity",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference)SetSessionLength(val *string) {
	if err := j.validateSetSessionLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionLength",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference)SetSessionLengthEnabled(val interface{}) {
	if err := j.validateSetSessionLengthEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionLengthEnabled",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference)SetSessionReauthMethod(val *string) {
	if err := j.validateSetSessionReauthMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionReauthMethod",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference)SetUseOidcMaxAge(val interface{}) {
	if err := j.validateSetUseOidcMaxAgeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useOidcMaxAge",
		val,
	)
}

func (a *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference) ResetMaxInactivity() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxInactivity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference) ResetSessionLength() {
	_jsii_.InvokeVoid(
		a,
		"resetSessionLength",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference) ResetSessionLengthEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetSessionLengthEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference) ResetSessionReauthMethod() {
	_jsii_.InvokeVoid(
		a,
		"resetSessionReauthMethod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference) ResetUseOidcMaxAge() {
	_jsii_.InvokeVoid(
		a,
		"resetUseOidcMaxAge",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AccessContextManagerGcpUserAccessBindingSessionSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

