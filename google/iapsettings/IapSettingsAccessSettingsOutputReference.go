// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iapsettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/iapsettings/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IapSettingsAccessSettingsOutputReference interface {
	cdktf.ComplexObject
	AllowedDomainsSettings() IapSettingsAccessSettingsAllowedDomainsSettingsOutputReference
	AllowedDomainsSettingsInput() *IapSettingsAccessSettingsAllowedDomainsSettings
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
	CorsSettings() IapSettingsAccessSettingsCorsSettingsOutputReference
	CorsSettingsInput() *IapSettingsAccessSettingsCorsSettings
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	GcipSettings() IapSettingsAccessSettingsGcipSettingsOutputReference
	GcipSettingsInput() *IapSettingsAccessSettingsGcipSettings
	IdentitySources() *[]*string
	SetIdentitySources(val *[]*string)
	IdentitySourcesInput() *[]*string
	InternalValue() *IapSettingsAccessSettings
	SetInternalValue(val *IapSettingsAccessSettings)
	OauthSettings() IapSettingsAccessSettingsOauthSettingsOutputReference
	OauthSettingsInput() *IapSettingsAccessSettingsOauthSettings
	ReauthSettings() IapSettingsAccessSettingsReauthSettingsOutputReference
	ReauthSettingsInput() *IapSettingsAccessSettingsReauthSettings
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WorkforceIdentitySettings() IapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference
	WorkforceIdentitySettingsInput() *IapSettingsAccessSettingsWorkforceIdentitySettings
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
	PutAllowedDomainsSettings(value *IapSettingsAccessSettingsAllowedDomainsSettings)
	PutCorsSettings(value *IapSettingsAccessSettingsCorsSettings)
	PutGcipSettings(value *IapSettingsAccessSettingsGcipSettings)
	PutOauthSettings(value *IapSettingsAccessSettingsOauthSettings)
	PutReauthSettings(value *IapSettingsAccessSettingsReauthSettings)
	PutWorkforceIdentitySettings(value *IapSettingsAccessSettingsWorkforceIdentitySettings)
	ResetAllowedDomainsSettings()
	ResetCorsSettings()
	ResetGcipSettings()
	ResetIdentitySources()
	ResetOauthSettings()
	ResetReauthSettings()
	ResetWorkforceIdentitySettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IapSettingsAccessSettingsOutputReference
type jsiiProxy_IapSettingsAccessSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IapSettingsAccessSettingsOutputReference) AllowedDomainsSettings() IapSettingsAccessSettingsAllowedDomainsSettingsOutputReference {
	var returns IapSettingsAccessSettingsAllowedDomainsSettingsOutputReference
	_jsii_.Get(
		j,
		"allowedDomainsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsAccessSettingsOutputReference) AllowedDomainsSettingsInput() *IapSettingsAccessSettingsAllowedDomainsSettings {
	var returns *IapSettingsAccessSettingsAllowedDomainsSettings
	_jsii_.Get(
		j,
		"allowedDomainsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsAccessSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsAccessSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsAccessSettingsOutputReference) CorsSettings() IapSettingsAccessSettingsCorsSettingsOutputReference {
	var returns IapSettingsAccessSettingsCorsSettingsOutputReference
	_jsii_.Get(
		j,
		"corsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsAccessSettingsOutputReference) CorsSettingsInput() *IapSettingsAccessSettingsCorsSettings {
	var returns *IapSettingsAccessSettingsCorsSettings
	_jsii_.Get(
		j,
		"corsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsAccessSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsAccessSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsAccessSettingsOutputReference) GcipSettings() IapSettingsAccessSettingsGcipSettingsOutputReference {
	var returns IapSettingsAccessSettingsGcipSettingsOutputReference
	_jsii_.Get(
		j,
		"gcipSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsAccessSettingsOutputReference) GcipSettingsInput() *IapSettingsAccessSettingsGcipSettings {
	var returns *IapSettingsAccessSettingsGcipSettings
	_jsii_.Get(
		j,
		"gcipSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsAccessSettingsOutputReference) IdentitySources() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identitySources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsAccessSettingsOutputReference) IdentitySourcesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identitySourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsAccessSettingsOutputReference) InternalValue() *IapSettingsAccessSettings {
	var returns *IapSettingsAccessSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsAccessSettingsOutputReference) OauthSettings() IapSettingsAccessSettingsOauthSettingsOutputReference {
	var returns IapSettingsAccessSettingsOauthSettingsOutputReference
	_jsii_.Get(
		j,
		"oauthSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsAccessSettingsOutputReference) OauthSettingsInput() *IapSettingsAccessSettingsOauthSettings {
	var returns *IapSettingsAccessSettingsOauthSettings
	_jsii_.Get(
		j,
		"oauthSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsAccessSettingsOutputReference) ReauthSettings() IapSettingsAccessSettingsReauthSettingsOutputReference {
	var returns IapSettingsAccessSettingsReauthSettingsOutputReference
	_jsii_.Get(
		j,
		"reauthSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsAccessSettingsOutputReference) ReauthSettingsInput() *IapSettingsAccessSettingsReauthSettings {
	var returns *IapSettingsAccessSettingsReauthSettings
	_jsii_.Get(
		j,
		"reauthSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsAccessSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsAccessSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsAccessSettingsOutputReference) WorkforceIdentitySettings() IapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference {
	var returns IapSettingsAccessSettingsWorkforceIdentitySettingsOutputReference
	_jsii_.Get(
		j,
		"workforceIdentitySettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsAccessSettingsOutputReference) WorkforceIdentitySettingsInput() *IapSettingsAccessSettingsWorkforceIdentitySettings {
	var returns *IapSettingsAccessSettingsWorkforceIdentitySettings
	_jsii_.Get(
		j,
		"workforceIdentitySettingsInput",
		&returns,
	)
	return returns
}


func NewIapSettingsAccessSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IapSettingsAccessSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewIapSettingsAccessSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IapSettingsAccessSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.iapSettings.IapSettingsAccessSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIapSettingsAccessSettingsOutputReference_Override(i IapSettingsAccessSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.iapSettings.IapSettingsAccessSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IapSettingsAccessSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IapSettingsAccessSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IapSettingsAccessSettingsOutputReference)SetIdentitySources(val *[]*string) {
	if err := j.validateSetIdentitySourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identitySources",
		val,
	)
}

func (j *jsiiProxy_IapSettingsAccessSettingsOutputReference)SetInternalValue(val *IapSettingsAccessSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IapSettingsAccessSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IapSettingsAccessSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IapSettingsAccessSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IapSettingsAccessSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IapSettingsAccessSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IapSettingsAccessSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IapSettingsAccessSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IapSettingsAccessSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IapSettingsAccessSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IapSettingsAccessSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IapSettingsAccessSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IapSettingsAccessSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IapSettingsAccessSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IapSettingsAccessSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IapSettingsAccessSettingsOutputReference) PutAllowedDomainsSettings(value *IapSettingsAccessSettingsAllowedDomainsSettings) {
	if err := i.validatePutAllowedDomainsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAllowedDomainsSettings",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IapSettingsAccessSettingsOutputReference) PutCorsSettings(value *IapSettingsAccessSettingsCorsSettings) {
	if err := i.validatePutCorsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putCorsSettings",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IapSettingsAccessSettingsOutputReference) PutGcipSettings(value *IapSettingsAccessSettingsGcipSettings) {
	if err := i.validatePutGcipSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putGcipSettings",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IapSettingsAccessSettingsOutputReference) PutOauthSettings(value *IapSettingsAccessSettingsOauthSettings) {
	if err := i.validatePutOauthSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putOauthSettings",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IapSettingsAccessSettingsOutputReference) PutReauthSettings(value *IapSettingsAccessSettingsReauthSettings) {
	if err := i.validatePutReauthSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putReauthSettings",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IapSettingsAccessSettingsOutputReference) PutWorkforceIdentitySettings(value *IapSettingsAccessSettingsWorkforceIdentitySettings) {
	if err := i.validatePutWorkforceIdentitySettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putWorkforceIdentitySettings",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IapSettingsAccessSettingsOutputReference) ResetAllowedDomainsSettings() {
	_jsii_.InvokeVoid(
		i,
		"resetAllowedDomainsSettings",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IapSettingsAccessSettingsOutputReference) ResetCorsSettings() {
	_jsii_.InvokeVoid(
		i,
		"resetCorsSettings",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IapSettingsAccessSettingsOutputReference) ResetGcipSettings() {
	_jsii_.InvokeVoid(
		i,
		"resetGcipSettings",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IapSettingsAccessSettingsOutputReference) ResetIdentitySources() {
	_jsii_.InvokeVoid(
		i,
		"resetIdentitySources",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IapSettingsAccessSettingsOutputReference) ResetOauthSettings() {
	_jsii_.InvokeVoid(
		i,
		"resetOauthSettings",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IapSettingsAccessSettingsOutputReference) ResetReauthSettings() {
	_jsii_.InvokeVoid(
		i,
		"resetReauthSettings",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IapSettingsAccessSettingsOutputReference) ResetWorkforceIdentitySettings() {
	_jsii_.InvokeVoid(
		i,
		"resetWorkforceIdentitySettings",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IapSettingsAccessSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (i *jsiiProxy_IapSettingsAccessSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

