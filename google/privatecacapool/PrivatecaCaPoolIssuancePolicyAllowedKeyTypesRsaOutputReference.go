// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privatecacapool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v13/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v13/privatecacapool/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsa
	SetInternalValue(val *PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsa)
	MaxModulusSize() *string
	SetMaxModulusSize(val *string)
	MaxModulusSizeInput() *string
	MinModulusSize() *string
	SetMinModulusSize(val *string)
	MinModulusSizeInput() *string
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
	ResetMaxModulusSize()
	ResetMinModulusSize()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference
type jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) InternalValue() *PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsa {
	var returns *PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsa
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) MaxModulusSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxModulusSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) MaxModulusSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxModulusSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) MinModulusSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minModulusSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) MinModulusSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minModulusSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference {
	_init_.Initialize()

	if err := validateNewPrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.privatecaCaPool.PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference_Override(p PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.privatecaCaPool.PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference)SetInternalValue(val *PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsa) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference)SetMaxModulusSize(val *string) {
	if err := j.validateSetMaxModulusSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxModulusSize",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference)SetMinModulusSize(val *string) {
	if err := j.validateSetMinModulusSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minModulusSize",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) ResetMaxModulusSize() {
	_jsii_.InvokeVoid(
		p,
		"resetMaxModulusSize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) ResetMinModulusSize() {
	_jsii_.InvokeVoid(
		p,
		"resetMinModulusSize",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PrivatecaCaPoolIssuancePolicyAllowedKeyTypesRsaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

