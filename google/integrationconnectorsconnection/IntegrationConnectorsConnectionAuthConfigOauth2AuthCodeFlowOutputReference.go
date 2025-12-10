// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationconnectorsconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/integrationconnectorsconnection/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference interface {
	cdktf.ComplexObject
	AuthUri() *string
	SetAuthUri(val *string)
	AuthUriInput() *string
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecret() IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowClientSecretOutputReference
	ClientSecretInput() *IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowClientSecret
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
	EnablePkce() interface{}
	SetEnablePkce(val interface{})
	EnablePkceInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlow
	SetInternalValue(val *IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlow)
	Scopes() *[]*string
	SetScopes(val *[]*string)
	ScopesInput() *[]*string
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
	PutClientSecret(value *IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowClientSecret)
	ResetAuthUri()
	ResetClientId()
	ResetClientSecret()
	ResetEnablePkce()
	ResetScopes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference
type jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) AuthUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) AuthUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) ClientSecret() IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowClientSecretOutputReference {
	var returns IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowClientSecretOutputReference
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) ClientSecretInput() *IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowClientSecret {
	var returns *IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowClientSecret
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) EnablePkce() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePkce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) EnablePkceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePkceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) InternalValue() *IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlow {
	var returns *IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlow
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) Scopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) ScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference {
	_init_.Initialize()

	if err := validateNewIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.integrationConnectorsConnection.IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference_Override(i IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.integrationConnectorsConnection.IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference)SetAuthUri(val *string) {
	if err := j.validateSetAuthUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authUri",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference)SetEnablePkce(val interface{}) {
	if err := j.validateSetEnablePkceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePkce",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference)SetInternalValue(val *IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlow) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference)SetScopes(val *[]*string) {
	if err := j.validateSetScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scopes",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) PutClientSecret(value *IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowClientSecret) {
	if err := i.validatePutClientSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putClientSecret",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) ResetAuthUri() {
	_jsii_.InvokeVoid(
		i,
		"resetAuthUri",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) ResetClientId() {
	_jsii_.InvokeVoid(
		i,
		"resetClientId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) ResetClientSecret() {
	_jsii_.InvokeVoid(
		i,
		"resetClientSecret",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) ResetEnablePkce() {
	_jsii_.InvokeVoid(
		i,
		"resetEnablePkce",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) ResetScopes() {
	_jsii_.InvokeVoid(
		i,
		"resetScopes",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := i.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigOauth2AuthCodeFlowOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

