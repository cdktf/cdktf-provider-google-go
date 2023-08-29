// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privatecacertificateauthority

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v9/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v9/privatecacertificateauthority/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PrivatecaCertificateAuthorityConfigX509ConfigOutputReference interface {
	cdktf.ComplexObject
	AdditionalExtensions() PrivatecaCertificateAuthorityConfigX509ConfigAdditionalExtensionsList
	AdditionalExtensionsInput() interface{}
	AiaOcspServers() *[]*string
	SetAiaOcspServers(val *[]*string)
	AiaOcspServersInput() *[]*string
	CaOptions() PrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference
	CaOptionsInput() *PrivatecaCertificateAuthorityConfigX509ConfigCaOptions
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
	InternalValue() *PrivatecaCertificateAuthorityConfigX509Config
	SetInternalValue(val *PrivatecaCertificateAuthorityConfigX509Config)
	KeyUsage() PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference
	KeyUsageInput() *PrivatecaCertificateAuthorityConfigX509ConfigKeyUsage
	NameConstraints() PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference
	NameConstraintsInput() *PrivatecaCertificateAuthorityConfigX509ConfigNameConstraints
	PolicyIds() PrivatecaCertificateAuthorityConfigX509ConfigPolicyIdsList
	PolicyIdsInput() interface{}
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
	PutAdditionalExtensions(value interface{})
	PutCaOptions(value *PrivatecaCertificateAuthorityConfigX509ConfigCaOptions)
	PutKeyUsage(value *PrivatecaCertificateAuthorityConfigX509ConfigKeyUsage)
	PutNameConstraints(value *PrivatecaCertificateAuthorityConfigX509ConfigNameConstraints)
	PutPolicyIds(value interface{})
	ResetAdditionalExtensions()
	ResetAiaOcspServers()
	ResetNameConstraints()
	ResetPolicyIds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PrivatecaCertificateAuthorityConfigX509ConfigOutputReference
type jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference) AdditionalExtensions() PrivatecaCertificateAuthorityConfigX509ConfigAdditionalExtensionsList {
	var returns PrivatecaCertificateAuthorityConfigX509ConfigAdditionalExtensionsList
	_jsii_.Get(
		j,
		"additionalExtensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference) AdditionalExtensionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalExtensionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference) AiaOcspServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aiaOcspServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference) AiaOcspServersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aiaOcspServersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference) CaOptions() PrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference {
	var returns PrivatecaCertificateAuthorityConfigX509ConfigCaOptionsOutputReference
	_jsii_.Get(
		j,
		"caOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference) CaOptionsInput() *PrivatecaCertificateAuthorityConfigX509ConfigCaOptions {
	var returns *PrivatecaCertificateAuthorityConfigX509ConfigCaOptions
	_jsii_.Get(
		j,
		"caOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference) InternalValue() *PrivatecaCertificateAuthorityConfigX509Config {
	var returns *PrivatecaCertificateAuthorityConfigX509Config
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference) KeyUsage() PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference {
	var returns PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference
	_jsii_.Get(
		j,
		"keyUsage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference) KeyUsageInput() *PrivatecaCertificateAuthorityConfigX509ConfigKeyUsage {
	var returns *PrivatecaCertificateAuthorityConfigX509ConfigKeyUsage
	_jsii_.Get(
		j,
		"keyUsageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference) NameConstraints() PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference {
	var returns PrivatecaCertificateAuthorityConfigX509ConfigNameConstraintsOutputReference
	_jsii_.Get(
		j,
		"nameConstraints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference) NameConstraintsInput() *PrivatecaCertificateAuthorityConfigX509ConfigNameConstraints {
	var returns *PrivatecaCertificateAuthorityConfigX509ConfigNameConstraints
	_jsii_.Get(
		j,
		"nameConstraintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference) PolicyIds() PrivatecaCertificateAuthorityConfigX509ConfigPolicyIdsList {
	var returns PrivatecaCertificateAuthorityConfigX509ConfigPolicyIdsList
	_jsii_.Get(
		j,
		"policyIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference) PolicyIdsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policyIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPrivatecaCertificateAuthorityConfigX509ConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PrivatecaCertificateAuthorityConfigX509ConfigOutputReference {
	_init_.Initialize()

	if err := validateNewPrivatecaCertificateAuthorityConfigX509ConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.privatecaCertificateAuthority.PrivatecaCertificateAuthorityConfigX509ConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPrivatecaCertificateAuthorityConfigX509ConfigOutputReference_Override(p PrivatecaCertificateAuthorityConfigX509ConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.privatecaCertificateAuthority.PrivatecaCertificateAuthorityConfigX509ConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference)SetAiaOcspServers(val *[]*string) {
	if err := j.validateSetAiaOcspServersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aiaOcspServers",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference)SetInternalValue(val *PrivatecaCertificateAuthorityConfigX509Config) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference) PutAdditionalExtensions(value interface{}) {
	if err := p.validatePutAdditionalExtensionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAdditionalExtensions",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference) PutCaOptions(value *PrivatecaCertificateAuthorityConfigX509ConfigCaOptions) {
	if err := p.validatePutCaOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCaOptions",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference) PutKeyUsage(value *PrivatecaCertificateAuthorityConfigX509ConfigKeyUsage) {
	if err := p.validatePutKeyUsageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putKeyUsage",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference) PutNameConstraints(value *PrivatecaCertificateAuthorityConfigX509ConfigNameConstraints) {
	if err := p.validatePutNameConstraintsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putNameConstraints",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference) PutPolicyIds(value interface{}) {
	if err := p.validatePutPolicyIdsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putPolicyIds",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference) ResetAdditionalExtensions() {
	_jsii_.InvokeVoid(
		p,
		"resetAdditionalExtensions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference) ResetAiaOcspServers() {
	_jsii_.InvokeVoid(
		p,
		"resetAiaOcspServers",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference) ResetNameConstraints() {
	_jsii_.InvokeVoid(
		p,
		"resetNameConstraints",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference) ResetPolicyIds() {
	_jsii_.InvokeVoid(
		p,
		"resetPolicyIds",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

