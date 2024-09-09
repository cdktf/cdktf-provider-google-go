// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privilegedaccessmanagerentitlement

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/privilegedaccessmanagerentitlement/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference interface {
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
	InternalValue() *PrivilegedAccessManagerEntitlementRequesterJustificationConfig
	SetInternalValue(val *PrivilegedAccessManagerEntitlementRequesterJustificationConfig)
	NotMandatory() PrivilegedAccessManagerEntitlementRequesterJustificationConfigNotMandatoryOutputReference
	NotMandatoryInput() *PrivilegedAccessManagerEntitlementRequesterJustificationConfigNotMandatory
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Unstructured() PrivilegedAccessManagerEntitlementRequesterJustificationConfigUnstructuredOutputReference
	UnstructuredInput() *PrivilegedAccessManagerEntitlementRequesterJustificationConfigUnstructured
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
	PutNotMandatory(value *PrivilegedAccessManagerEntitlementRequesterJustificationConfigNotMandatory)
	PutUnstructured(value *PrivilegedAccessManagerEntitlementRequesterJustificationConfigUnstructured)
	ResetNotMandatory()
	ResetUnstructured()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference
type jsiiProxy_PrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference) InternalValue() *PrivilegedAccessManagerEntitlementRequesterJustificationConfig {
	var returns *PrivilegedAccessManagerEntitlementRequesterJustificationConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference) NotMandatory() PrivilegedAccessManagerEntitlementRequesterJustificationConfigNotMandatoryOutputReference {
	var returns PrivilegedAccessManagerEntitlementRequesterJustificationConfigNotMandatoryOutputReference
	_jsii_.Get(
		j,
		"notMandatory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference) NotMandatoryInput() *PrivilegedAccessManagerEntitlementRequesterJustificationConfigNotMandatory {
	var returns *PrivilegedAccessManagerEntitlementRequesterJustificationConfigNotMandatory
	_jsii_.Get(
		j,
		"notMandatoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference) Unstructured() PrivilegedAccessManagerEntitlementRequesterJustificationConfigUnstructuredOutputReference {
	var returns PrivilegedAccessManagerEntitlementRequesterJustificationConfigUnstructuredOutputReference
	_jsii_.Get(
		j,
		"unstructured",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference) UnstructuredInput() *PrivilegedAccessManagerEntitlementRequesterJustificationConfigUnstructured {
	var returns *PrivilegedAccessManagerEntitlementRequesterJustificationConfigUnstructured
	_jsii_.Get(
		j,
		"unstructuredInput",
		&returns,
	)
	return returns
}


func NewPrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference {
	_init_.Initialize()

	if err := validateNewPrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.privilegedAccessManagerEntitlement.PrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference_Override(p PrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.privilegedAccessManagerEntitlement.PrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference)SetInternalValue(val *PrivilegedAccessManagerEntitlementRequesterJustificationConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference) PutNotMandatory(value *PrivilegedAccessManagerEntitlementRequesterJustificationConfigNotMandatory) {
	if err := p.validatePutNotMandatoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putNotMandatory",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference) PutUnstructured(value *PrivilegedAccessManagerEntitlementRequesterJustificationConfigUnstructured) {
	if err := p.validatePutUnstructuredParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putUnstructured",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference) ResetNotMandatory() {
	_jsii_.InvokeVoid(
		p,
		"resetNotMandatory",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference) ResetUnstructured() {
	_jsii_.InvokeVoid(
		p,
		"resetUnstructured",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementRequesterJustificationConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

