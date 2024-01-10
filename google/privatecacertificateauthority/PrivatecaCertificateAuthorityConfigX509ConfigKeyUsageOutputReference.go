// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privatecacertificateauthority

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v13/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v13/privatecacertificateauthority/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference interface {
	cdktf.ComplexObject
	BaseKeyUsage() PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference
	BaseKeyUsageInput() *PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage
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
	ExtendedKeyUsage() PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference
	ExtendedKeyUsageInput() *PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage
	// Experimental.
	Fqn() *string
	InternalValue() *PrivatecaCertificateAuthorityConfigX509ConfigKeyUsage
	SetInternalValue(val *PrivatecaCertificateAuthorityConfigX509ConfigKeyUsage)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UnknownExtendedKeyUsages() PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesList
	UnknownExtendedKeyUsagesInput() interface{}
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
	PutBaseKeyUsage(value *PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage)
	PutExtendedKeyUsage(value *PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage)
	PutUnknownExtendedKeyUsages(value interface{})
	ResetUnknownExtendedKeyUsages()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference
type jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) BaseKeyUsage() PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference {
	var returns PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsageOutputReference
	_jsii_.Get(
		j,
		"baseKeyUsage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) BaseKeyUsageInput() *PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage {
	var returns *PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage
	_jsii_.Get(
		j,
		"baseKeyUsageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) ExtendedKeyUsage() PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference {
	var returns PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsageOutputReference
	_jsii_.Get(
		j,
		"extendedKeyUsage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) ExtendedKeyUsageInput() *PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage {
	var returns *PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage
	_jsii_.Get(
		j,
		"extendedKeyUsageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) InternalValue() *PrivatecaCertificateAuthorityConfigX509ConfigKeyUsage {
	var returns *PrivatecaCertificateAuthorityConfigX509ConfigKeyUsage
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) UnknownExtendedKeyUsages() PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesList {
	var returns PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageUnknownExtendedKeyUsagesList
	_jsii_.Get(
		j,
		"unknownExtendedKeyUsages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) UnknownExtendedKeyUsagesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unknownExtendedKeyUsagesInput",
		&returns,
	)
	return returns
}


func NewPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference {
	_init_.Initialize()

	if err := validateNewPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.privatecaCertificateAuthority.PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference_Override(p PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.privatecaCertificateAuthority.PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference)SetInternalValue(val *PrivatecaCertificateAuthorityConfigX509ConfigKeyUsage) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) PutBaseKeyUsage(value *PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageBaseKeyUsage) {
	if err := p.validatePutBaseKeyUsageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putBaseKeyUsage",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) PutExtendedKeyUsage(value *PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageExtendedKeyUsage) {
	if err := p.validatePutExtendedKeyUsageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putExtendedKeyUsage",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) PutUnknownExtendedKeyUsages(value interface{}) {
	if err := p.validatePutUnknownExtendedKeyUsagesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putUnknownExtendedKeyUsages",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) ResetUnknownExtendedKeyUsages() {
	_jsii_.InvokeVoid(
		p,
		"resetUnknownExtendedKeyUsages",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PrivatecaCertificateAuthorityConfigX509ConfigKeyUsageOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

