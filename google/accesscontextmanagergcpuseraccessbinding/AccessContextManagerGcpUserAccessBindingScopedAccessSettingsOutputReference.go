// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accesscontextmanagergcpuseraccessbinding

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/accesscontextmanagergcpuseraccessbinding/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference interface {
	cdktf.ComplexObject
	ActiveSettings() AccessContextManagerGcpUserAccessBindingScopedAccessSettingsActiveSettingsOutputReference
	ActiveSettingsInput() *AccessContextManagerGcpUserAccessBindingScopedAccessSettingsActiveSettings
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
	DryRunSettings() AccessContextManagerGcpUserAccessBindingScopedAccessSettingsDryRunSettingsOutputReference
	DryRunSettingsInput() *AccessContextManagerGcpUserAccessBindingScopedAccessSettingsDryRunSettings
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Scope() AccessContextManagerGcpUserAccessBindingScopedAccessSettingsScopeOutputReference
	ScopeInput() *AccessContextManagerGcpUserAccessBindingScopedAccessSettingsScope
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
	PutActiveSettings(value *AccessContextManagerGcpUserAccessBindingScopedAccessSettingsActiveSettings)
	PutDryRunSettings(value *AccessContextManagerGcpUserAccessBindingScopedAccessSettingsDryRunSettings)
	PutScope(value *AccessContextManagerGcpUserAccessBindingScopedAccessSettingsScope)
	ResetActiveSettings()
	ResetDryRunSettings()
	ResetScope()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference
type jsiiProxy_AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) ActiveSettings() AccessContextManagerGcpUserAccessBindingScopedAccessSettingsActiveSettingsOutputReference {
	var returns AccessContextManagerGcpUserAccessBindingScopedAccessSettingsActiveSettingsOutputReference
	_jsii_.Get(
		j,
		"activeSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) ActiveSettingsInput() *AccessContextManagerGcpUserAccessBindingScopedAccessSettingsActiveSettings {
	var returns *AccessContextManagerGcpUserAccessBindingScopedAccessSettingsActiveSettings
	_jsii_.Get(
		j,
		"activeSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) DryRunSettings() AccessContextManagerGcpUserAccessBindingScopedAccessSettingsDryRunSettingsOutputReference {
	var returns AccessContextManagerGcpUserAccessBindingScopedAccessSettingsDryRunSettingsOutputReference
	_jsii_.Get(
		j,
		"dryRunSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) DryRunSettingsInput() *AccessContextManagerGcpUserAccessBindingScopedAccessSettingsDryRunSettings {
	var returns *AccessContextManagerGcpUserAccessBindingScopedAccessSettingsDryRunSettings
	_jsii_.Get(
		j,
		"dryRunSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) Scope() AccessContextManagerGcpUserAccessBindingScopedAccessSettingsScopeOutputReference {
	var returns AccessContextManagerGcpUserAccessBindingScopedAccessSettingsScopeOutputReference
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) ScopeInput() *AccessContextManagerGcpUserAccessBindingScopedAccessSettingsScope {
	var returns *AccessContextManagerGcpUserAccessBindingScopedAccessSettingsScope
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.accessContextManagerGcpUserAccessBinding.AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference_Override(a AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.accessContextManagerGcpUserAccessBinding.AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) PutActiveSettings(value *AccessContextManagerGcpUserAccessBindingScopedAccessSettingsActiveSettings) {
	if err := a.validatePutActiveSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putActiveSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) PutDryRunSettings(value *AccessContextManagerGcpUserAccessBindingScopedAccessSettingsDryRunSettings) {
	if err := a.validatePutDryRunSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDryRunSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) PutScope(value *AccessContextManagerGcpUserAccessBindingScopedAccessSettingsScope) {
	if err := a.validatePutScopeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putScope",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) ResetActiveSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetActiveSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) ResetDryRunSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetDryRunSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) ResetScope() {
	_jsii_.InvokeVoid(
		a,
		"resetScope",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AccessContextManagerGcpUserAccessBindingScopedAccessSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

