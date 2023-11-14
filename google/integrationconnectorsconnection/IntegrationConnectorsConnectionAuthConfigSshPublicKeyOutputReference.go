// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationconnectorsconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v12/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v12/integrationconnectorsconnection/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference interface {
	cdktf.ComplexObject
	CertType() *string
	SetCertType(val *string)
	CertTypeInput() *string
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
	InternalValue() *IntegrationConnectorsConnectionAuthConfigSshPublicKey
	SetInternalValue(val *IntegrationConnectorsConnectionAuthConfigSshPublicKey)
	SshClientCert() IntegrationConnectorsConnectionAuthConfigSshPublicKeySshClientCertOutputReference
	SshClientCertInput() *IntegrationConnectorsConnectionAuthConfigSshPublicKeySshClientCert
	SshClientCertPass() IntegrationConnectorsConnectionAuthConfigSshPublicKeySshClientCertPassOutputReference
	SshClientCertPassInput() *IntegrationConnectorsConnectionAuthConfigSshPublicKeySshClientCertPass
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
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
	PutSshClientCert(value *IntegrationConnectorsConnectionAuthConfigSshPublicKeySshClientCert)
	PutSshClientCertPass(value *IntegrationConnectorsConnectionAuthConfigSshPublicKeySshClientCertPass)
	ResetCertType()
	ResetSshClientCert()
	ResetSshClientCertPass()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference
type jsiiProxy_IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference) CertType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference) CertTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference) InternalValue() *IntegrationConnectorsConnectionAuthConfigSshPublicKey {
	var returns *IntegrationConnectorsConnectionAuthConfigSshPublicKey
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference) SshClientCert() IntegrationConnectorsConnectionAuthConfigSshPublicKeySshClientCertOutputReference {
	var returns IntegrationConnectorsConnectionAuthConfigSshPublicKeySshClientCertOutputReference
	_jsii_.Get(
		j,
		"sshClientCert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference) SshClientCertInput() *IntegrationConnectorsConnectionAuthConfigSshPublicKeySshClientCert {
	var returns *IntegrationConnectorsConnectionAuthConfigSshPublicKeySshClientCert
	_jsii_.Get(
		j,
		"sshClientCertInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference) SshClientCertPass() IntegrationConnectorsConnectionAuthConfigSshPublicKeySshClientCertPassOutputReference {
	var returns IntegrationConnectorsConnectionAuthConfigSshPublicKeySshClientCertPassOutputReference
	_jsii_.Get(
		j,
		"sshClientCertPass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference) SshClientCertPassInput() *IntegrationConnectorsConnectionAuthConfigSshPublicKeySshClientCertPass {
	var returns *IntegrationConnectorsConnectionAuthConfigSshPublicKeySshClientCertPass
	_jsii_.Get(
		j,
		"sshClientCertPassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


func NewIntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference {
	_init_.Initialize()

	if err := validateNewIntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.integrationConnectorsConnection.IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference_Override(i IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.integrationConnectorsConnection.IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference)SetCertType(val *string) {
	if err := j.validateSetCertTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certType",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference)SetInternalValue(val *IntegrationConnectorsConnectionAuthConfigSshPublicKey) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference) PutSshClientCert(value *IntegrationConnectorsConnectionAuthConfigSshPublicKeySshClientCert) {
	if err := i.validatePutSshClientCertParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSshClientCert",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference) PutSshClientCertPass(value *IntegrationConnectorsConnectionAuthConfigSshPublicKeySshClientCertPass) {
	if err := i.validatePutSshClientCertPassParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSshClientCertPass",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference) ResetCertType() {
	_jsii_.InvokeVoid(
		i,
		"resetCertType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference) ResetSshClientCert() {
	_jsii_.InvokeVoid(
		i,
		"resetSshClientCert",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference) ResetSshClientCertPass() {
	_jsii_.InvokeVoid(
		i,
		"resetSshClientCertPass",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (i *jsiiProxy_IntegrationConnectorsConnectionAuthConfigSshPublicKeyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

