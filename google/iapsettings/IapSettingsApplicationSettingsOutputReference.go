// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iapsettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/iapsettings/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IapSettingsApplicationSettingsOutputReference interface {
	cdktf.ComplexObject
	AccessDeniedPageSettings() IapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference
	AccessDeniedPageSettingsInput() *IapSettingsApplicationSettingsAccessDeniedPageSettings
	AttributePropagationSettings() IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference
	AttributePropagationSettingsInput() *IapSettingsApplicationSettingsAttributePropagationSettings
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
	CookieDomain() *string
	SetCookieDomain(val *string)
	CookieDomainInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CsmSettings() IapSettingsApplicationSettingsCsmSettingsOutputReference
	CsmSettingsInput() *IapSettingsApplicationSettingsCsmSettings
	// Experimental.
	Fqn() *string
	InternalValue() *IapSettingsApplicationSettings
	SetInternalValue(val *IapSettingsApplicationSettings)
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
	PutAccessDeniedPageSettings(value *IapSettingsApplicationSettingsAccessDeniedPageSettings)
	PutAttributePropagationSettings(value *IapSettingsApplicationSettingsAttributePropagationSettings)
	PutCsmSettings(value *IapSettingsApplicationSettingsCsmSettings)
	ResetAccessDeniedPageSettings()
	ResetAttributePropagationSettings()
	ResetCookieDomain()
	ResetCsmSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IapSettingsApplicationSettingsOutputReference
type jsiiProxy_IapSettingsApplicationSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IapSettingsApplicationSettingsOutputReference) AccessDeniedPageSettings() IapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference {
	var returns IapSettingsApplicationSettingsAccessDeniedPageSettingsOutputReference
	_jsii_.Get(
		j,
		"accessDeniedPageSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsApplicationSettingsOutputReference) AccessDeniedPageSettingsInput() *IapSettingsApplicationSettingsAccessDeniedPageSettings {
	var returns *IapSettingsApplicationSettingsAccessDeniedPageSettings
	_jsii_.Get(
		j,
		"accessDeniedPageSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsApplicationSettingsOutputReference) AttributePropagationSettings() IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference {
	var returns IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference
	_jsii_.Get(
		j,
		"attributePropagationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsApplicationSettingsOutputReference) AttributePropagationSettingsInput() *IapSettingsApplicationSettingsAttributePropagationSettings {
	var returns *IapSettingsApplicationSettingsAttributePropagationSettings
	_jsii_.Get(
		j,
		"attributePropagationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsApplicationSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsApplicationSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsApplicationSettingsOutputReference) CookieDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cookieDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsApplicationSettingsOutputReference) CookieDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cookieDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsApplicationSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsApplicationSettingsOutputReference) CsmSettings() IapSettingsApplicationSettingsCsmSettingsOutputReference {
	var returns IapSettingsApplicationSettingsCsmSettingsOutputReference
	_jsii_.Get(
		j,
		"csmSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsApplicationSettingsOutputReference) CsmSettingsInput() *IapSettingsApplicationSettingsCsmSettings {
	var returns *IapSettingsApplicationSettingsCsmSettings
	_jsii_.Get(
		j,
		"csmSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsApplicationSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsApplicationSettingsOutputReference) InternalValue() *IapSettingsApplicationSettings {
	var returns *IapSettingsApplicationSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsApplicationSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsApplicationSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIapSettingsApplicationSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IapSettingsApplicationSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewIapSettingsApplicationSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IapSettingsApplicationSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.iapSettings.IapSettingsApplicationSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIapSettingsApplicationSettingsOutputReference_Override(i IapSettingsApplicationSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.iapSettings.IapSettingsApplicationSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IapSettingsApplicationSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IapSettingsApplicationSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IapSettingsApplicationSettingsOutputReference)SetCookieDomain(val *string) {
	if err := j.validateSetCookieDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cookieDomain",
		val,
	)
}

func (j *jsiiProxy_IapSettingsApplicationSettingsOutputReference)SetInternalValue(val *IapSettingsApplicationSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IapSettingsApplicationSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IapSettingsApplicationSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IapSettingsApplicationSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IapSettingsApplicationSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IapSettingsApplicationSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IapSettingsApplicationSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IapSettingsApplicationSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IapSettingsApplicationSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IapSettingsApplicationSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IapSettingsApplicationSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IapSettingsApplicationSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IapSettingsApplicationSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IapSettingsApplicationSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IapSettingsApplicationSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IapSettingsApplicationSettingsOutputReference) PutAccessDeniedPageSettings(value *IapSettingsApplicationSettingsAccessDeniedPageSettings) {
	if err := i.validatePutAccessDeniedPageSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAccessDeniedPageSettings",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IapSettingsApplicationSettingsOutputReference) PutAttributePropagationSettings(value *IapSettingsApplicationSettingsAttributePropagationSettings) {
	if err := i.validatePutAttributePropagationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAttributePropagationSettings",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IapSettingsApplicationSettingsOutputReference) PutCsmSettings(value *IapSettingsApplicationSettingsCsmSettings) {
	if err := i.validatePutCsmSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putCsmSettings",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IapSettingsApplicationSettingsOutputReference) ResetAccessDeniedPageSettings() {
	_jsii_.InvokeVoid(
		i,
		"resetAccessDeniedPageSettings",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IapSettingsApplicationSettingsOutputReference) ResetAttributePropagationSettings() {
	_jsii_.InvokeVoid(
		i,
		"resetAttributePropagationSettings",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IapSettingsApplicationSettingsOutputReference) ResetCookieDomain() {
	_jsii_.InvokeVoid(
		i,
		"resetCookieDomain",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IapSettingsApplicationSettingsOutputReference) ResetCsmSettings() {
	_jsii_.InvokeVoid(
		i,
		"resetCsmSettings",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IapSettingsApplicationSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := i.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IapSettingsApplicationSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

