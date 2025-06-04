// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privatecacertificate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/privatecacertificate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference interface {
	cdktf.ComplexObject
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
	DnsNames() *[]*string
	SetDnsNames(val *[]*string)
	DnsNamesInput() *[]*string
	EmailAddresses() *[]*string
	SetEmailAddresses(val *[]*string)
	EmailAddressesInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *PrivatecaCertificateConfigSubjectConfigSubjectAltName
	SetInternalValue(val *PrivatecaCertificateConfigSubjectConfigSubjectAltName)
	IpAddresses() *[]*string
	SetIpAddresses(val *[]*string)
	IpAddressesInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Uris() *[]*string
	SetUris(val *[]*string)
	UrisInput() *[]*string
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
	ResetDnsNames()
	ResetEmailAddresses()
	ResetIpAddresses()
	ResetUris()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference
type jsiiProxy_PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference) DnsNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference) DnsNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference) EmailAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"emailAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference) EmailAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"emailAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference) InternalValue() *PrivatecaCertificateConfigSubjectConfigSubjectAltName {
	var returns *PrivatecaCertificateConfigSubjectConfigSubjectAltName
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference) IpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference) IpAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference) Uris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"uris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference) UrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"urisInput",
		&returns,
	)
	return returns
}


func NewPrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference {
	_init_.Initialize()

	if err := validateNewPrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.privatecaCertificate.PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference_Override(p PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.privatecaCertificate.PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference)SetDnsNames(val *[]*string) {
	if err := j.validateSetDnsNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsNames",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference)SetEmailAddresses(val *[]*string) {
	if err := j.validateSetEmailAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailAddresses",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference)SetInternalValue(val *PrivatecaCertificateConfigSubjectConfigSubjectAltName) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference)SetIpAddresses(val *[]*string) {
	if err := j.validateSetIpAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddresses",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference)SetUris(val *[]*string) {
	if err := j.validateSetUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uris",
		val,
	)
}

func (p *jsiiProxy_PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference) ResetDnsNames() {
	_jsii_.InvokeVoid(
		p,
		"resetDnsNames",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference) ResetEmailAddresses() {
	_jsii_.InvokeVoid(
		p,
		"resetEmailAddresses",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference) ResetIpAddresses() {
	_jsii_.InvokeVoid(
		p,
		"resetIpAddresses",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference) ResetUris() {
	_jsii_.InvokeVoid(
		p,
		"resetUris",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PrivatecaCertificateConfigSubjectConfigSubjectAltNameOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

