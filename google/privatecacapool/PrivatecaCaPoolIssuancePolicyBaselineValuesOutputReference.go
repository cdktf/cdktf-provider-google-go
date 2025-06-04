// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privatecacapool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/privatecacapool/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference interface {
	cdktf.ComplexObject
	AdditionalExtensions() PrivatecaCaPoolIssuancePolicyBaselineValuesAdditionalExtensionsList
	AdditionalExtensionsInput() interface{}
	AiaOcspServers() *[]*string
	SetAiaOcspServers(val *[]*string)
	AiaOcspServersInput() *[]*string
	CaOptions() PrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference
	CaOptionsInput() *PrivatecaCaPoolIssuancePolicyBaselineValuesCaOptions
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
	InternalValue() *PrivatecaCaPoolIssuancePolicyBaselineValues
	SetInternalValue(val *PrivatecaCaPoolIssuancePolicyBaselineValues)
	KeyUsage() PrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageOutputReference
	KeyUsageInput() *PrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsage
	NameConstraints() PrivatecaCaPoolIssuancePolicyBaselineValuesNameConstraintsOutputReference
	NameConstraintsInput() *PrivatecaCaPoolIssuancePolicyBaselineValuesNameConstraints
	PolicyIds() PrivatecaCaPoolIssuancePolicyBaselineValuesPolicyIdsList
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
	PutCaOptions(value *PrivatecaCaPoolIssuancePolicyBaselineValuesCaOptions)
	PutKeyUsage(value *PrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsage)
	PutNameConstraints(value *PrivatecaCaPoolIssuancePolicyBaselineValuesNameConstraints)
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

// The jsii proxy struct for PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference
type jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference) AdditionalExtensions() PrivatecaCaPoolIssuancePolicyBaselineValuesAdditionalExtensionsList {
	var returns PrivatecaCaPoolIssuancePolicyBaselineValuesAdditionalExtensionsList
	_jsii_.Get(
		j,
		"additionalExtensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference) AdditionalExtensionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalExtensionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference) AiaOcspServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aiaOcspServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference) AiaOcspServersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aiaOcspServersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference) CaOptions() PrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference {
	var returns PrivatecaCaPoolIssuancePolicyBaselineValuesCaOptionsOutputReference
	_jsii_.Get(
		j,
		"caOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference) CaOptionsInput() *PrivatecaCaPoolIssuancePolicyBaselineValuesCaOptions {
	var returns *PrivatecaCaPoolIssuancePolicyBaselineValuesCaOptions
	_jsii_.Get(
		j,
		"caOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference) InternalValue() *PrivatecaCaPoolIssuancePolicyBaselineValues {
	var returns *PrivatecaCaPoolIssuancePolicyBaselineValues
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference) KeyUsage() PrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageOutputReference {
	var returns PrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsageOutputReference
	_jsii_.Get(
		j,
		"keyUsage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference) KeyUsageInput() *PrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsage {
	var returns *PrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsage
	_jsii_.Get(
		j,
		"keyUsageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference) NameConstraints() PrivatecaCaPoolIssuancePolicyBaselineValuesNameConstraintsOutputReference {
	var returns PrivatecaCaPoolIssuancePolicyBaselineValuesNameConstraintsOutputReference
	_jsii_.Get(
		j,
		"nameConstraints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference) NameConstraintsInput() *PrivatecaCaPoolIssuancePolicyBaselineValuesNameConstraints {
	var returns *PrivatecaCaPoolIssuancePolicyBaselineValuesNameConstraints
	_jsii_.Get(
		j,
		"nameConstraintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference) PolicyIds() PrivatecaCaPoolIssuancePolicyBaselineValuesPolicyIdsList {
	var returns PrivatecaCaPoolIssuancePolicyBaselineValuesPolicyIdsList
	_jsii_.Get(
		j,
		"policyIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference) PolicyIdsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policyIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference {
	_init_.Initialize()

	if err := validateNewPrivatecaCaPoolIssuancePolicyBaselineValuesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.privatecaCaPool.PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference_Override(p PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.privatecaCaPool.PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference)SetAiaOcspServers(val *[]*string) {
	if err := j.validateSetAiaOcspServersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aiaOcspServers",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference)SetInternalValue(val *PrivatecaCaPoolIssuancePolicyBaselineValues) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference) PutAdditionalExtensions(value interface{}) {
	if err := p.validatePutAdditionalExtensionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAdditionalExtensions",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference) PutCaOptions(value *PrivatecaCaPoolIssuancePolicyBaselineValuesCaOptions) {
	if err := p.validatePutCaOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCaOptions",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference) PutKeyUsage(value *PrivatecaCaPoolIssuancePolicyBaselineValuesKeyUsage) {
	if err := p.validatePutKeyUsageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putKeyUsage",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference) PutNameConstraints(value *PrivatecaCaPoolIssuancePolicyBaselineValuesNameConstraints) {
	if err := p.validatePutNameConstraintsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putNameConstraints",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference) PutPolicyIds(value interface{}) {
	if err := p.validatePutPolicyIdsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putPolicyIds",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference) ResetAdditionalExtensions() {
	_jsii_.InvokeVoid(
		p,
		"resetAdditionalExtensions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference) ResetAiaOcspServers() {
	_jsii_.InvokeVoid(
		p,
		"resetAiaOcspServers",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference) ResetNameConstraints() {
	_jsii_.InvokeVoid(
		p,
		"resetNameConstraints",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference) ResetPolicyIds() {
	_jsii_.InvokeVoid(
		p,
		"resetPolicyIds",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyBaselineValuesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

