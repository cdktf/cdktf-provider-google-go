// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationconnectorsconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/integrationconnectorsconnection/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IntegrationConnectorsConnectionEventingConfigOutputReference interface {
	cdktf.ComplexObject
	AdditionalVariable() IntegrationConnectorsConnectionEventingConfigAdditionalVariableList
	AdditionalVariableInput() interface{}
	AuthConfig() IntegrationConnectorsConnectionEventingConfigAuthConfigOutputReference
	AuthConfigInput() *IntegrationConnectorsConnectionEventingConfigAuthConfig
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
	EnrichmentEnabled() interface{}
	SetEnrichmentEnabled(val interface{})
	EnrichmentEnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *IntegrationConnectorsConnectionEventingConfig
	SetInternalValue(val *IntegrationConnectorsConnectionEventingConfig)
	RegistrationDestinationConfig() IntegrationConnectorsConnectionEventingConfigRegistrationDestinationConfigOutputReference
	RegistrationDestinationConfigInput() *IntegrationConnectorsConnectionEventingConfigRegistrationDestinationConfig
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
	PutAdditionalVariable(value interface{})
	PutAuthConfig(value *IntegrationConnectorsConnectionEventingConfigAuthConfig)
	PutRegistrationDestinationConfig(value *IntegrationConnectorsConnectionEventingConfigRegistrationDestinationConfig)
	ResetAdditionalVariable()
	ResetAuthConfig()
	ResetEnrichmentEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IntegrationConnectorsConnectionEventingConfigOutputReference
type jsiiProxy_IntegrationConnectorsConnectionEventingConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IntegrationConnectorsConnectionEventingConfigOutputReference) AdditionalVariable() IntegrationConnectorsConnectionEventingConfigAdditionalVariableList {
	var returns IntegrationConnectorsConnectionEventingConfigAdditionalVariableList
	_jsii_.Get(
		j,
		"additionalVariable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionEventingConfigOutputReference) AdditionalVariableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalVariableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionEventingConfigOutputReference) AuthConfig() IntegrationConnectorsConnectionEventingConfigAuthConfigOutputReference {
	var returns IntegrationConnectorsConnectionEventingConfigAuthConfigOutputReference
	_jsii_.Get(
		j,
		"authConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionEventingConfigOutputReference) AuthConfigInput() *IntegrationConnectorsConnectionEventingConfigAuthConfig {
	var returns *IntegrationConnectorsConnectionEventingConfigAuthConfig
	_jsii_.Get(
		j,
		"authConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionEventingConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionEventingConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionEventingConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionEventingConfigOutputReference) EnrichmentEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enrichmentEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionEventingConfigOutputReference) EnrichmentEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enrichmentEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionEventingConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionEventingConfigOutputReference) InternalValue() *IntegrationConnectorsConnectionEventingConfig {
	var returns *IntegrationConnectorsConnectionEventingConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionEventingConfigOutputReference) RegistrationDestinationConfig() IntegrationConnectorsConnectionEventingConfigRegistrationDestinationConfigOutputReference {
	var returns IntegrationConnectorsConnectionEventingConfigRegistrationDestinationConfigOutputReference
	_jsii_.Get(
		j,
		"registrationDestinationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionEventingConfigOutputReference) RegistrationDestinationConfigInput() *IntegrationConnectorsConnectionEventingConfigRegistrationDestinationConfig {
	var returns *IntegrationConnectorsConnectionEventingConfigRegistrationDestinationConfig
	_jsii_.Get(
		j,
		"registrationDestinationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionEventingConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionEventingConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIntegrationConnectorsConnectionEventingConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IntegrationConnectorsConnectionEventingConfigOutputReference {
	_init_.Initialize()

	if err := validateNewIntegrationConnectorsConnectionEventingConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IntegrationConnectorsConnectionEventingConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.integrationConnectorsConnection.IntegrationConnectorsConnectionEventingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIntegrationConnectorsConnectionEventingConfigOutputReference_Override(i IntegrationConnectorsConnectionEventingConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.integrationConnectorsConnection.IntegrationConnectorsConnectionEventingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionEventingConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionEventingConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionEventingConfigOutputReference)SetEnrichmentEnabled(val interface{}) {
	if err := j.validateSetEnrichmentEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enrichmentEnabled",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionEventingConfigOutputReference)SetInternalValue(val *IntegrationConnectorsConnectionEventingConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionEventingConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionEventingConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionEventingConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationConnectorsConnectionEventingConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionEventingConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionEventingConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionEventingConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionEventingConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionEventingConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionEventingConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionEventingConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionEventingConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionEventingConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationConnectorsConnectionEventingConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionEventingConfigOutputReference) PutAdditionalVariable(value interface{}) {
	if err := i.validatePutAdditionalVariableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAdditionalVariable",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionEventingConfigOutputReference) PutAuthConfig(value *IntegrationConnectorsConnectionEventingConfigAuthConfig) {
	if err := i.validatePutAuthConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAuthConfig",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionEventingConfigOutputReference) PutRegistrationDestinationConfig(value *IntegrationConnectorsConnectionEventingConfigRegistrationDestinationConfig) {
	if err := i.validatePutRegistrationDestinationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putRegistrationDestinationConfig",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionEventingConfigOutputReference) ResetAdditionalVariable() {
	_jsii_.InvokeVoid(
		i,
		"resetAdditionalVariable",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionEventingConfigOutputReference) ResetAuthConfig() {
	_jsii_.InvokeVoid(
		i,
		"resetAuthConfig",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionEventingConfigOutputReference) ResetEnrichmentEnabled() {
	_jsii_.InvokeVoid(
		i,
		"resetEnrichmentEnabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionEventingConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionEventingConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

