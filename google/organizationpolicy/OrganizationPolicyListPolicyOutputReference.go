// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package organizationpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v13/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v13/organizationpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrganizationPolicyListPolicyOutputReference interface {
	cdktf.ComplexObject
	Allow() OrganizationPolicyListPolicyAllowOutputReference
	AllowInput() *OrganizationPolicyListPolicyAllow
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
	Deny() OrganizationPolicyListPolicyDenyOutputReference
	DenyInput() *OrganizationPolicyListPolicyDeny
	// Experimental.
	Fqn() *string
	InheritFromParent() interface{}
	SetInheritFromParent(val interface{})
	InheritFromParentInput() interface{}
	InternalValue() *OrganizationPolicyListPolicy
	SetInternalValue(val *OrganizationPolicyListPolicy)
	SuggestedValue() *string
	SetSuggestedValue(val *string)
	SuggestedValueInput() *string
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
	PutAllow(value *OrganizationPolicyListPolicyAllow)
	PutDeny(value *OrganizationPolicyListPolicyDeny)
	ResetAllow()
	ResetDeny()
	ResetInheritFromParent()
	ResetSuggestedValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OrganizationPolicyListPolicyOutputReference
type jsiiProxy_OrganizationPolicyListPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OrganizationPolicyListPolicyOutputReference) Allow() OrganizationPolicyListPolicyAllowOutputReference {
	var returns OrganizationPolicyListPolicyAllowOutputReference
	_jsii_.Get(
		j,
		"allow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationPolicyListPolicyOutputReference) AllowInput() *OrganizationPolicyListPolicyAllow {
	var returns *OrganizationPolicyListPolicyAllow
	_jsii_.Get(
		j,
		"allowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationPolicyListPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationPolicyListPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationPolicyListPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationPolicyListPolicyOutputReference) Deny() OrganizationPolicyListPolicyDenyOutputReference {
	var returns OrganizationPolicyListPolicyDenyOutputReference
	_jsii_.Get(
		j,
		"deny",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationPolicyListPolicyOutputReference) DenyInput() *OrganizationPolicyListPolicyDeny {
	var returns *OrganizationPolicyListPolicyDeny
	_jsii_.Get(
		j,
		"denyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationPolicyListPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationPolicyListPolicyOutputReference) InheritFromParent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inheritFromParent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationPolicyListPolicyOutputReference) InheritFromParentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inheritFromParentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationPolicyListPolicyOutputReference) InternalValue() *OrganizationPolicyListPolicy {
	var returns *OrganizationPolicyListPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationPolicyListPolicyOutputReference) SuggestedValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suggestedValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationPolicyListPolicyOutputReference) SuggestedValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suggestedValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationPolicyListPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationPolicyListPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOrganizationPolicyListPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OrganizationPolicyListPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewOrganizationPolicyListPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OrganizationPolicyListPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.organizationPolicy.OrganizationPolicyListPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOrganizationPolicyListPolicyOutputReference_Override(o OrganizationPolicyListPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.organizationPolicy.OrganizationPolicyListPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OrganizationPolicyListPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OrganizationPolicyListPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OrganizationPolicyListPolicyOutputReference)SetInheritFromParent(val interface{}) {
	if err := j.validateSetInheritFromParentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inheritFromParent",
		val,
	)
}

func (j *jsiiProxy_OrganizationPolicyListPolicyOutputReference)SetInternalValue(val *OrganizationPolicyListPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OrganizationPolicyListPolicyOutputReference)SetSuggestedValue(val *string) {
	if err := j.validateSetSuggestedValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suggestedValue",
		val,
	)
}

func (j *jsiiProxy_OrganizationPolicyListPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OrganizationPolicyListPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OrganizationPolicyListPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationPolicyListPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationPolicyListPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationPolicyListPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationPolicyListPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationPolicyListPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationPolicyListPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationPolicyListPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationPolicyListPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationPolicyListPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationPolicyListPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationPolicyListPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationPolicyListPolicyOutputReference) PutAllow(value *OrganizationPolicyListPolicyAllow) {
	if err := o.validatePutAllowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAllow",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OrganizationPolicyListPolicyOutputReference) PutDeny(value *OrganizationPolicyListPolicyDeny) {
	if err := o.validatePutDenyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putDeny",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OrganizationPolicyListPolicyOutputReference) ResetAllow() {
	_jsii_.InvokeVoid(
		o,
		"resetAllow",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationPolicyListPolicyOutputReference) ResetDeny() {
	_jsii_.InvokeVoid(
		o,
		"resetDeny",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationPolicyListPolicyOutputReference) ResetInheritFromParent() {
	_jsii_.InvokeVoid(
		o,
		"resetInheritFromParent",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationPolicyListPolicyOutputReference) ResetSuggestedValue() {
	_jsii_.InvokeVoid(
		o,
		"resetSuggestedValue",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationPolicyListPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationPolicyListPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

