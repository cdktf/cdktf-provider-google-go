// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigeeaddonsconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/apigeeaddonsconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigeeAddonsConfigAddonsConfigOutputReference interface {
	cdktf.ComplexObject
	AdvancedApiOpsConfig() ApigeeAddonsConfigAddonsConfigAdvancedApiOpsConfigOutputReference
	AdvancedApiOpsConfigInput() *ApigeeAddonsConfigAddonsConfigAdvancedApiOpsConfig
	ApiSecurityConfig() ApigeeAddonsConfigAddonsConfigApiSecurityConfigOutputReference
	ApiSecurityConfigInput() *ApigeeAddonsConfigAddonsConfigApiSecurityConfig
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
	ConnectorsPlatformConfig() ApigeeAddonsConfigAddonsConfigConnectorsPlatformConfigOutputReference
	ConnectorsPlatformConfigInput() *ApigeeAddonsConfigAddonsConfigConnectorsPlatformConfig
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	IntegrationConfig() ApigeeAddonsConfigAddonsConfigIntegrationConfigOutputReference
	IntegrationConfigInput() *ApigeeAddonsConfigAddonsConfigIntegrationConfig
	InternalValue() *ApigeeAddonsConfigAddonsConfig
	SetInternalValue(val *ApigeeAddonsConfigAddonsConfig)
	MonetizationConfig() ApigeeAddonsConfigAddonsConfigMonetizationConfigOutputReference
	MonetizationConfigInput() *ApigeeAddonsConfigAddonsConfigMonetizationConfig
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
	PutAdvancedApiOpsConfig(value *ApigeeAddonsConfigAddonsConfigAdvancedApiOpsConfig)
	PutApiSecurityConfig(value *ApigeeAddonsConfigAddonsConfigApiSecurityConfig)
	PutConnectorsPlatformConfig(value *ApigeeAddonsConfigAddonsConfigConnectorsPlatformConfig)
	PutIntegrationConfig(value *ApigeeAddonsConfigAddonsConfigIntegrationConfig)
	PutMonetizationConfig(value *ApigeeAddonsConfigAddonsConfigMonetizationConfig)
	ResetAdvancedApiOpsConfig()
	ResetApiSecurityConfig()
	ResetConnectorsPlatformConfig()
	ResetIntegrationConfig()
	ResetMonetizationConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApigeeAddonsConfigAddonsConfigOutputReference
type jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference) AdvancedApiOpsConfig() ApigeeAddonsConfigAddonsConfigAdvancedApiOpsConfigOutputReference {
	var returns ApigeeAddonsConfigAddonsConfigAdvancedApiOpsConfigOutputReference
	_jsii_.Get(
		j,
		"advancedApiOpsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference) AdvancedApiOpsConfigInput() *ApigeeAddonsConfigAddonsConfigAdvancedApiOpsConfig {
	var returns *ApigeeAddonsConfigAddonsConfigAdvancedApiOpsConfig
	_jsii_.Get(
		j,
		"advancedApiOpsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference) ApiSecurityConfig() ApigeeAddonsConfigAddonsConfigApiSecurityConfigOutputReference {
	var returns ApigeeAddonsConfigAddonsConfigApiSecurityConfigOutputReference
	_jsii_.Get(
		j,
		"apiSecurityConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference) ApiSecurityConfigInput() *ApigeeAddonsConfigAddonsConfigApiSecurityConfig {
	var returns *ApigeeAddonsConfigAddonsConfigApiSecurityConfig
	_jsii_.Get(
		j,
		"apiSecurityConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference) ConnectorsPlatformConfig() ApigeeAddonsConfigAddonsConfigConnectorsPlatformConfigOutputReference {
	var returns ApigeeAddonsConfigAddonsConfigConnectorsPlatformConfigOutputReference
	_jsii_.Get(
		j,
		"connectorsPlatformConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference) ConnectorsPlatformConfigInput() *ApigeeAddonsConfigAddonsConfigConnectorsPlatformConfig {
	var returns *ApigeeAddonsConfigAddonsConfigConnectorsPlatformConfig
	_jsii_.Get(
		j,
		"connectorsPlatformConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference) IntegrationConfig() ApigeeAddonsConfigAddonsConfigIntegrationConfigOutputReference {
	var returns ApigeeAddonsConfigAddonsConfigIntegrationConfigOutputReference
	_jsii_.Get(
		j,
		"integrationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference) IntegrationConfigInput() *ApigeeAddonsConfigAddonsConfigIntegrationConfig {
	var returns *ApigeeAddonsConfigAddonsConfigIntegrationConfig
	_jsii_.Get(
		j,
		"integrationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference) InternalValue() *ApigeeAddonsConfigAddonsConfig {
	var returns *ApigeeAddonsConfigAddonsConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference) MonetizationConfig() ApigeeAddonsConfigAddonsConfigMonetizationConfigOutputReference {
	var returns ApigeeAddonsConfigAddonsConfigMonetizationConfigOutputReference
	_jsii_.Get(
		j,
		"monetizationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference) MonetizationConfigInput() *ApigeeAddonsConfigAddonsConfigMonetizationConfig {
	var returns *ApigeeAddonsConfigAddonsConfigMonetizationConfig
	_jsii_.Get(
		j,
		"monetizationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApigeeAddonsConfigAddonsConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApigeeAddonsConfigAddonsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewApigeeAddonsConfigAddonsConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.apigeeAddonsConfig.ApigeeAddonsConfigAddonsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApigeeAddonsConfigAddonsConfigOutputReference_Override(a ApigeeAddonsConfigAddonsConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.apigeeAddonsConfig.ApigeeAddonsConfigAddonsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference)SetInternalValue(val *ApigeeAddonsConfigAddonsConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference) PutAdvancedApiOpsConfig(value *ApigeeAddonsConfigAddonsConfigAdvancedApiOpsConfig) {
	if err := a.validatePutAdvancedApiOpsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAdvancedApiOpsConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference) PutApiSecurityConfig(value *ApigeeAddonsConfigAddonsConfigApiSecurityConfig) {
	if err := a.validatePutApiSecurityConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putApiSecurityConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference) PutConnectorsPlatformConfig(value *ApigeeAddonsConfigAddonsConfigConnectorsPlatformConfig) {
	if err := a.validatePutConnectorsPlatformConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putConnectorsPlatformConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference) PutIntegrationConfig(value *ApigeeAddonsConfigAddonsConfigIntegrationConfig) {
	if err := a.validatePutIntegrationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIntegrationConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference) PutMonetizationConfig(value *ApigeeAddonsConfigAddonsConfigMonetizationConfig) {
	if err := a.validatePutMonetizationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMonetizationConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference) ResetAdvancedApiOpsConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetAdvancedApiOpsConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference) ResetApiSecurityConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetApiSecurityConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference) ResetConnectorsPlatformConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetConnectorsPlatformConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference) ResetIntegrationConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetIntegrationConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference) ResetMonetizationConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetMonetizationConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApigeeAddonsConfigAddonsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

