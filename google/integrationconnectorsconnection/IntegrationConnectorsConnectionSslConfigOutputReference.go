// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationconnectorsconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v12/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v12/integrationconnectorsconnection/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IntegrationConnectorsConnectionSslConfigOutputReference interface {
	cdktf.ComplexObject
	AdditionalVariable() IntegrationConnectorsConnectionSslConfigAdditionalVariableList
	AdditionalVariableInput() interface{}
	ClientCertificate() IntegrationConnectorsConnectionSslConfigClientCertificateOutputReference
	ClientCertificateInput() *IntegrationConnectorsConnectionSslConfigClientCertificate
	ClientCertType() *string
	SetClientCertType(val *string)
	ClientCertTypeInput() *string
	ClientPrivateKey() IntegrationConnectorsConnectionSslConfigClientPrivateKeyOutputReference
	ClientPrivateKeyInput() *IntegrationConnectorsConnectionSslConfigClientPrivateKey
	ClientPrivateKeyPass() IntegrationConnectorsConnectionSslConfigClientPrivateKeyPassOutputReference
	ClientPrivateKeyPassInput() *IntegrationConnectorsConnectionSslConfigClientPrivateKeyPass
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
	InternalValue() *IntegrationConnectorsConnectionSslConfig
	SetInternalValue(val *IntegrationConnectorsConnectionSslConfig)
	PrivateServerCertificate() IntegrationConnectorsConnectionSslConfigPrivateServerCertificateOutputReference
	PrivateServerCertificateInput() *IntegrationConnectorsConnectionSslConfigPrivateServerCertificate
	ServerCertType() *string
	SetServerCertType(val *string)
	ServerCertTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TrustModel() *string
	SetTrustModel(val *string)
	TrustModelInput() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	UseSsl() interface{}
	SetUseSsl(val interface{})
	UseSslInput() interface{}
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
	PutClientCertificate(value *IntegrationConnectorsConnectionSslConfigClientCertificate)
	PutClientPrivateKey(value *IntegrationConnectorsConnectionSslConfigClientPrivateKey)
	PutClientPrivateKeyPass(value *IntegrationConnectorsConnectionSslConfigClientPrivateKeyPass)
	PutPrivateServerCertificate(value *IntegrationConnectorsConnectionSslConfigPrivateServerCertificate)
	ResetAdditionalVariable()
	ResetClientCertificate()
	ResetClientCertType()
	ResetClientPrivateKey()
	ResetClientPrivateKeyPass()
	ResetPrivateServerCertificate()
	ResetServerCertType()
	ResetTrustModel()
	ResetUseSsl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IntegrationConnectorsConnectionSslConfigOutputReference
type jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) AdditionalVariable() IntegrationConnectorsConnectionSslConfigAdditionalVariableList {
	var returns IntegrationConnectorsConnectionSslConfigAdditionalVariableList
	_jsii_.Get(
		j,
		"additionalVariable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) AdditionalVariableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalVariableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) ClientCertificate() IntegrationConnectorsConnectionSslConfigClientCertificateOutputReference {
	var returns IntegrationConnectorsConnectionSslConfigClientCertificateOutputReference
	_jsii_.Get(
		j,
		"clientCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) ClientCertificateInput() *IntegrationConnectorsConnectionSslConfigClientCertificate {
	var returns *IntegrationConnectorsConnectionSslConfigClientCertificate
	_jsii_.Get(
		j,
		"clientCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) ClientCertType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) ClientCertTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) ClientPrivateKey() IntegrationConnectorsConnectionSslConfigClientPrivateKeyOutputReference {
	var returns IntegrationConnectorsConnectionSslConfigClientPrivateKeyOutputReference
	_jsii_.Get(
		j,
		"clientPrivateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) ClientPrivateKeyInput() *IntegrationConnectorsConnectionSslConfigClientPrivateKey {
	var returns *IntegrationConnectorsConnectionSslConfigClientPrivateKey
	_jsii_.Get(
		j,
		"clientPrivateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) ClientPrivateKeyPass() IntegrationConnectorsConnectionSslConfigClientPrivateKeyPassOutputReference {
	var returns IntegrationConnectorsConnectionSslConfigClientPrivateKeyPassOutputReference
	_jsii_.Get(
		j,
		"clientPrivateKeyPass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) ClientPrivateKeyPassInput() *IntegrationConnectorsConnectionSslConfigClientPrivateKeyPass {
	var returns *IntegrationConnectorsConnectionSslConfigClientPrivateKeyPass
	_jsii_.Get(
		j,
		"clientPrivateKeyPassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) InternalValue() *IntegrationConnectorsConnectionSslConfig {
	var returns *IntegrationConnectorsConnectionSslConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) PrivateServerCertificate() IntegrationConnectorsConnectionSslConfigPrivateServerCertificateOutputReference {
	var returns IntegrationConnectorsConnectionSslConfigPrivateServerCertificateOutputReference
	_jsii_.Get(
		j,
		"privateServerCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) PrivateServerCertificateInput() *IntegrationConnectorsConnectionSslConfigPrivateServerCertificate {
	var returns *IntegrationConnectorsConnectionSslConfigPrivateServerCertificate
	_jsii_.Get(
		j,
		"privateServerCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) ServerCertType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverCertType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) ServerCertTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverCertTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) TrustModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) TrustModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) UseSsl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useSsl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) UseSslInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useSslInput",
		&returns,
	)
	return returns
}


func NewIntegrationConnectorsConnectionSslConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IntegrationConnectorsConnectionSslConfigOutputReference {
	_init_.Initialize()

	if err := validateNewIntegrationConnectorsConnectionSslConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.integrationConnectorsConnection.IntegrationConnectorsConnectionSslConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIntegrationConnectorsConnectionSslConfigOutputReference_Override(i IntegrationConnectorsConnectionSslConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.integrationConnectorsConnection.IntegrationConnectorsConnectionSslConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference)SetClientCertType(val *string) {
	if err := j.validateSetClientCertTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCertType",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference)SetInternalValue(val *IntegrationConnectorsConnectionSslConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference)SetServerCertType(val *string) {
	if err := j.validateSetServerCertTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverCertType",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference)SetTrustModel(val *string) {
	if err := j.validateSetTrustModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustModel",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference)SetUseSsl(val interface{}) {
	if err := j.validateSetUseSslParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useSsl",
		val,
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) PutAdditionalVariable(value interface{}) {
	if err := i.validatePutAdditionalVariableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAdditionalVariable",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) PutClientCertificate(value *IntegrationConnectorsConnectionSslConfigClientCertificate) {
	if err := i.validatePutClientCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putClientCertificate",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) PutClientPrivateKey(value *IntegrationConnectorsConnectionSslConfigClientPrivateKey) {
	if err := i.validatePutClientPrivateKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putClientPrivateKey",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) PutClientPrivateKeyPass(value *IntegrationConnectorsConnectionSslConfigClientPrivateKeyPass) {
	if err := i.validatePutClientPrivateKeyPassParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putClientPrivateKeyPass",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) PutPrivateServerCertificate(value *IntegrationConnectorsConnectionSslConfigPrivateServerCertificate) {
	if err := i.validatePutPrivateServerCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putPrivateServerCertificate",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) ResetAdditionalVariable() {
	_jsii_.InvokeVoid(
		i,
		"resetAdditionalVariable",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) ResetClientCertificate() {
	_jsii_.InvokeVoid(
		i,
		"resetClientCertificate",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) ResetClientCertType() {
	_jsii_.InvokeVoid(
		i,
		"resetClientCertType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) ResetClientPrivateKey() {
	_jsii_.InvokeVoid(
		i,
		"resetClientPrivateKey",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) ResetClientPrivateKeyPass() {
	_jsii_.InvokeVoid(
		i,
		"resetClientPrivateKeyPass",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) ResetPrivateServerCertificate() {
	_jsii_.InvokeVoid(
		i,
		"resetPrivateServerCertificate",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) ResetServerCertType() {
	_jsii_.InvokeVoid(
		i,
		"resetServerCertType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) ResetTrustModel() {
	_jsii_.InvokeVoid(
		i,
		"resetTrustModel",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) ResetUseSsl() {
	_jsii_.InvokeVoid(
		i,
		"resetUseSsl",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionSslConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

