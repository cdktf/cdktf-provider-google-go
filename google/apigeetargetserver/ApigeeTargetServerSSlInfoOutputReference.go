// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigeetargetserver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v12/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v12/apigeetargetserver/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigeeTargetServerSSlInfoOutputReference interface {
	cdktf.ComplexObject
	Ciphers() *[]*string
	SetCiphers(val *[]*string)
	CiphersInput() *[]*string
	ClientAuthEnabled() interface{}
	SetClientAuthEnabled(val interface{})
	ClientAuthEnabledInput() interface{}
	CommonName() ApigeeTargetServerSSlInfoCommonNameOutputReference
	CommonNameInput() *ApigeeTargetServerSSlInfoCommonName
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	IgnoreValidationErrors() interface{}
	SetIgnoreValidationErrors(val interface{})
	IgnoreValidationErrorsInput() interface{}
	InternalValue() *ApigeeTargetServerSSlInfo
	SetInternalValue(val *ApigeeTargetServerSSlInfo)
	KeyAlias() *string
	SetKeyAlias(val *string)
	KeyAliasInput() *string
	KeyStore() *string
	SetKeyStore(val *string)
	KeyStoreInput() *string
	Protocols() *[]*string
	SetProtocols(val *[]*string)
	ProtocolsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TrustStore() *string
	SetTrustStore(val *string)
	TrustStoreInput() *string
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
	PutCommonName(value *ApigeeTargetServerSSlInfoCommonName)
	ResetCiphers()
	ResetClientAuthEnabled()
	ResetCommonName()
	ResetIgnoreValidationErrors()
	ResetKeyAlias()
	ResetKeyStore()
	ResetProtocols()
	ResetTrustStore()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApigeeTargetServerSSlInfoOutputReference
type jsiiProxy_ApigeeTargetServerSSlInfoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) Ciphers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ciphers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) CiphersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ciphersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) ClientAuthEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientAuthEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) ClientAuthEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientAuthEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) CommonName() ApigeeTargetServerSSlInfoCommonNameOutputReference {
	var returns ApigeeTargetServerSSlInfoCommonNameOutputReference
	_jsii_.Get(
		j,
		"commonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) CommonNameInput() *ApigeeTargetServerSSlInfoCommonName {
	var returns *ApigeeTargetServerSSlInfoCommonName
	_jsii_.Get(
		j,
		"commonNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) IgnoreValidationErrors() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreValidationErrors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) IgnoreValidationErrorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreValidationErrorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) InternalValue() *ApigeeTargetServerSSlInfo {
	var returns *ApigeeTargetServerSSlInfo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) KeyAlias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) KeyAliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyAliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) KeyStore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) KeyStoreInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyStoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) Protocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) ProtocolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) TrustStore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) TrustStoreInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustStoreInput",
		&returns,
	)
	return returns
}


func NewApigeeTargetServerSSlInfoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApigeeTargetServerSSlInfoOutputReference {
	_init_.Initialize()

	if err := validateNewApigeeTargetServerSSlInfoOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApigeeTargetServerSSlInfoOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.apigeeTargetServer.ApigeeTargetServerSSlInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApigeeTargetServerSSlInfoOutputReference_Override(a ApigeeTargetServerSSlInfoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.apigeeTargetServer.ApigeeTargetServerSSlInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference)SetCiphers(val *[]*string) {
	if err := j.validateSetCiphersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ciphers",
		val,
	)
}

func (j *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference)SetClientAuthEnabled(val interface{}) {
	if err := j.validateSetClientAuthEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientAuthEnabled",
		val,
	)
}

func (j *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference)SetIgnoreValidationErrors(val interface{}) {
	if err := j.validateSetIgnoreValidationErrorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreValidationErrors",
		val,
	)
}

func (j *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference)SetInternalValue(val *ApigeeTargetServerSSlInfo) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference)SetKeyAlias(val *string) {
	if err := j.validateSetKeyAliasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyAlias",
		val,
	)
}

func (j *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference)SetKeyStore(val *string) {
	if err := j.validateSetKeyStoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyStore",
		val,
	)
}

func (j *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference)SetProtocols(val *[]*string) {
	if err := j.validateSetProtocolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocols",
		val,
	)
}

func (j *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference)SetTrustStore(val *string) {
	if err := j.validateSetTrustStoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustStore",
		val,
	)
}

func (a *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) PutCommonName(value *ApigeeTargetServerSSlInfoCommonName) {
	if err := a.validatePutCommonNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCommonName",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) ResetCiphers() {
	_jsii_.InvokeVoid(
		a,
		"resetCiphers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) ResetClientAuthEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetClientAuthEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) ResetCommonName() {
	_jsii_.InvokeVoid(
		a,
		"resetCommonName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) ResetIgnoreValidationErrors() {
	_jsii_.InvokeVoid(
		a,
		"resetIgnoreValidationErrors",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) ResetKeyAlias() {
	_jsii_.InvokeVoid(
		a,
		"resetKeyAlias",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) ResetKeyStore() {
	_jsii_.InvokeVoid(
		a,
		"resetKeyStore",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) ResetProtocols() {
	_jsii_.InvokeVoid(
		a,
		"resetProtocols",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) ResetTrustStore() {
	_jsii_.InvokeVoid(
		a,
		"resetTrustStore",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApigeeTargetServerSSlInfoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

