// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package folderorganizationpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/folderorganizationpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FolderOrganizationPolicyListPolicyOutputReference interface {
	cdktf.ComplexObject
	Allow() FolderOrganizationPolicyListPolicyAllowOutputReference
	AllowInput() *FolderOrganizationPolicyListPolicyAllow
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
	Deny() FolderOrganizationPolicyListPolicyDenyOutputReference
	DenyInput() *FolderOrganizationPolicyListPolicyDeny
	// Experimental.
	Fqn() *string
	InheritFromParent() interface{}
	SetInheritFromParent(val interface{})
	InheritFromParentInput() interface{}
	InternalValue() *FolderOrganizationPolicyListPolicy
	SetInternalValue(val *FolderOrganizationPolicyListPolicy)
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
	PutAllow(value *FolderOrganizationPolicyListPolicyAllow)
	PutDeny(value *FolderOrganizationPolicyListPolicyDeny)
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

// The jsii proxy struct for FolderOrganizationPolicyListPolicyOutputReference
type jsiiProxy_FolderOrganizationPolicyListPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FolderOrganizationPolicyListPolicyOutputReference) Allow() FolderOrganizationPolicyListPolicyAllowOutputReference {
	var returns FolderOrganizationPolicyListPolicyAllowOutputReference
	_jsii_.Get(
		j,
		"allow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FolderOrganizationPolicyListPolicyOutputReference) AllowInput() *FolderOrganizationPolicyListPolicyAllow {
	var returns *FolderOrganizationPolicyListPolicyAllow
	_jsii_.Get(
		j,
		"allowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FolderOrganizationPolicyListPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FolderOrganizationPolicyListPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FolderOrganizationPolicyListPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FolderOrganizationPolicyListPolicyOutputReference) Deny() FolderOrganizationPolicyListPolicyDenyOutputReference {
	var returns FolderOrganizationPolicyListPolicyDenyOutputReference
	_jsii_.Get(
		j,
		"deny",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FolderOrganizationPolicyListPolicyOutputReference) DenyInput() *FolderOrganizationPolicyListPolicyDeny {
	var returns *FolderOrganizationPolicyListPolicyDeny
	_jsii_.Get(
		j,
		"denyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FolderOrganizationPolicyListPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FolderOrganizationPolicyListPolicyOutputReference) InheritFromParent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inheritFromParent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FolderOrganizationPolicyListPolicyOutputReference) InheritFromParentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inheritFromParentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FolderOrganizationPolicyListPolicyOutputReference) InternalValue() *FolderOrganizationPolicyListPolicy {
	var returns *FolderOrganizationPolicyListPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FolderOrganizationPolicyListPolicyOutputReference) SuggestedValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suggestedValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FolderOrganizationPolicyListPolicyOutputReference) SuggestedValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suggestedValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FolderOrganizationPolicyListPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FolderOrganizationPolicyListPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewFolderOrganizationPolicyListPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) FolderOrganizationPolicyListPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewFolderOrganizationPolicyListPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_FolderOrganizationPolicyListPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.folderOrganizationPolicy.FolderOrganizationPolicyListPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFolderOrganizationPolicyListPolicyOutputReference_Override(f FolderOrganizationPolicyListPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.folderOrganizationPolicy.FolderOrganizationPolicyListPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FolderOrganizationPolicyListPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FolderOrganizationPolicyListPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FolderOrganizationPolicyListPolicyOutputReference)SetInheritFromParent(val interface{}) {
	if err := j.validateSetInheritFromParentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inheritFromParent",
		val,
	)
}

func (j *jsiiProxy_FolderOrganizationPolicyListPolicyOutputReference)SetInternalValue(val *FolderOrganizationPolicyListPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FolderOrganizationPolicyListPolicyOutputReference)SetSuggestedValue(val *string) {
	if err := j.validateSetSuggestedValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suggestedValue",
		val,
	)
}

func (j *jsiiProxy_FolderOrganizationPolicyListPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FolderOrganizationPolicyListPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FolderOrganizationPolicyListPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FolderOrganizationPolicyListPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FolderOrganizationPolicyListPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FolderOrganizationPolicyListPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FolderOrganizationPolicyListPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FolderOrganizationPolicyListPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FolderOrganizationPolicyListPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FolderOrganizationPolicyListPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FolderOrganizationPolicyListPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FolderOrganizationPolicyListPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FolderOrganizationPolicyListPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FolderOrganizationPolicyListPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FolderOrganizationPolicyListPolicyOutputReference) PutAllow(value *FolderOrganizationPolicyListPolicyAllow) {
	if err := f.validatePutAllowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putAllow",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FolderOrganizationPolicyListPolicyOutputReference) PutDeny(value *FolderOrganizationPolicyListPolicyDeny) {
	if err := f.validatePutDenyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putDeny",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FolderOrganizationPolicyListPolicyOutputReference) ResetAllow() {
	_jsii_.InvokeVoid(
		f,
		"resetAllow",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FolderOrganizationPolicyListPolicyOutputReference) ResetDeny() {
	_jsii_.InvokeVoid(
		f,
		"resetDeny",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FolderOrganizationPolicyListPolicyOutputReference) ResetInheritFromParent() {
	_jsii_.InvokeVoid(
		f,
		"resetInheritFromParent",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FolderOrganizationPolicyListPolicyOutputReference) ResetSuggestedValue() {
	_jsii_.InvokeVoid(
		f,
		"resetSuggestedValue",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FolderOrganizationPolicyListPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := f.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FolderOrganizationPolicyListPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

