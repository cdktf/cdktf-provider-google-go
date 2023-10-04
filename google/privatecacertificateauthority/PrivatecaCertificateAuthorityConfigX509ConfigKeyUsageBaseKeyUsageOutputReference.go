// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privatecacertificateauthority

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v10/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v10/privatecacertificateauthority/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference interface {
	cdktf.ComplexObject
	CertSign() interface{}
	SetCertSign(val interface{})
	CertSignInput() interface{}
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
	ContentCommitment() interface{}
	SetContentCommitment(val interface{})
	ContentCommitmentInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CrlSign() interface{}
	SetCrlSign(val interface{})
	CrlSignInput() interface{}
	DataEncipherment() interface{}
	SetDataEncipherment(val interface{})
	DataEnciphermentInput() interface{}
	DecipherOnly() interface{}
	SetDecipherOnly(val interface{})
	DecipherOnlyInput() interface{}
	DigitalSignature() interface{}
	SetDigitalSignature(val interface{})
	DigitalSignatureInput() interface{}
	EncipherOnly() interface{}
	SetEncipherOnly(val interface{})
	EncipherOnlyInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage
	SetInternalValue(val *PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage)
	KeyAgreement() interface{}
	SetKeyAgreement(val interface{})
	KeyAgreementInput() interface{}
	KeyEncipherment() interface{}
	SetKeyEncipherment(val interface{})
	KeyEnciphermentInput() interface{}
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
	ResetCertSign()
	ResetContentCommitment()
	ResetCrlSign()
	ResetDataEncipherment()
	ResetDecipherOnly()
	ResetDigitalSignature()
	ResetEncipherOnly()
	ResetKeyAgreement()
	ResetKeyEncipherment()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference
type jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) CertSign() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"certSign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) CertSignInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"certSignInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) ContentCommitment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentCommitment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) ContentCommitmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentCommitmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) CrlSign() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crlSign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) CrlSignInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crlSignInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) DataEncipherment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataEncipherment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) DataEnciphermentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataEnciphermentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) DecipherOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"decipherOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) DecipherOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"decipherOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) DigitalSignature() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"digitalSignature",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) DigitalSignatureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"digitalSignatureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) EncipherOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encipherOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) EncipherOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encipherOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) InternalValue() *PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage {
	var returns *PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) KeyAgreement() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keyAgreement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) KeyAgreementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keyAgreementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) KeyEncipherment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keyEncipherment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) KeyEnciphermentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keyEnciphermentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference {
	_init_.Initialize()

	if err := validateNewPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.privatecaCertificateAuthority.PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference_Override(p PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.privatecaCertificateAuthority.PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference)SetCertSign(val interface{}) {
	if err := j.validateSetCertSignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certSign",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference)SetContentCommitment(val interface{}) {
	if err := j.validateSetContentCommitmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentCommitment",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference)SetCrlSign(val interface{}) {
	if err := j.validateSetCrlSignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crlSign",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference)SetDataEncipherment(val interface{}) {
	if err := j.validateSetDataEnciphermentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataEncipherment",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference)SetDecipherOnly(val interface{}) {
	if err := j.validateSetDecipherOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"decipherOnly",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference)SetDigitalSignature(val interface{}) {
	if err := j.validateSetDigitalSignatureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"digitalSignature",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference)SetEncipherOnly(val interface{}) {
	if err := j.validateSetEncipherOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encipherOnly",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference)SetInternalValue(val *PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference)SetKeyAgreement(val interface{}) {
	if err := j.validateSetKeyAgreementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyAgreement",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference)SetKeyEncipherment(val interface{}) {
	if err := j.validateSetKeyEnciphermentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyEncipherment",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) ResetCertSign() {
	_jsii_.InvokeVoid(
		p,
		"resetCertSign",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) ResetContentCommitment() {
	_jsii_.InvokeVoid(
		p,
		"resetContentCommitment",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) ResetCrlSign() {
	_jsii_.InvokeVoid(
		p,
		"resetCrlSign",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) ResetDataEncipherment() {
	_jsii_.InvokeVoid(
		p,
		"resetDataEncipherment",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) ResetDecipherOnly() {
	_jsii_.InvokeVoid(
		p,
		"resetDecipherOnly",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) ResetDigitalSignature() {
	_jsii_.InvokeVoid(
		p,
		"resetDigitalSignature",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) ResetEncipherOnly() {
	_jsii_.InvokeVoid(
		p,
		"resetEncipherOnly",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) ResetKeyAgreement() {
	_jsii_.InvokeVoid(
		p,
		"resetKeyAgreement",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) ResetKeyEncipherment() {
	_jsii_.InvokeVoid(
		p,
		"resetKeyEncipherment",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

