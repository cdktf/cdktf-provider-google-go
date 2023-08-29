// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iamdenypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v9/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v9/iamdenypolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IamDenyPolicyRulesDenyRuleOutputReference interface {
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
	DenialCondition() IamDenyPolicyRulesDenyRuleDenialConditionOutputReference
	DenialConditionInput() *IamDenyPolicyRulesDenyRuleDenialCondition
	DeniedPermissions() *[]*string
	SetDeniedPermissions(val *[]*string)
	DeniedPermissionsInput() *[]*string
	DeniedPrincipals() *[]*string
	SetDeniedPrincipals(val *[]*string)
	DeniedPrincipalsInput() *[]*string
	ExceptionPermissions() *[]*string
	SetExceptionPermissions(val *[]*string)
	ExceptionPermissionsInput() *[]*string
	ExceptionPrincipals() *[]*string
	SetExceptionPrincipals(val *[]*string)
	ExceptionPrincipalsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *IamDenyPolicyRulesDenyRule
	SetInternalValue(val *IamDenyPolicyRulesDenyRule)
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
	PutDenialCondition(value *IamDenyPolicyRulesDenyRuleDenialCondition)
	ResetDenialCondition()
	ResetDeniedPermissions()
	ResetDeniedPrincipals()
	ResetExceptionPermissions()
	ResetExceptionPrincipals()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IamDenyPolicyRulesDenyRuleOutputReference
type jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference) DenialCondition() IamDenyPolicyRulesDenyRuleDenialConditionOutputReference {
	var returns IamDenyPolicyRulesDenyRuleDenialConditionOutputReference
	_jsii_.Get(
		j,
		"denialCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference) DenialConditionInput() *IamDenyPolicyRulesDenyRuleDenialCondition {
	var returns *IamDenyPolicyRulesDenyRuleDenialCondition
	_jsii_.Get(
		j,
		"denialConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference) DeniedPermissions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"deniedPermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference) DeniedPermissionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"deniedPermissionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference) DeniedPrincipals() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"deniedPrincipals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference) DeniedPrincipalsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"deniedPrincipalsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference) ExceptionPermissions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exceptionPermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference) ExceptionPermissionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exceptionPermissionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference) ExceptionPrincipals() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exceptionPrincipals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference) ExceptionPrincipalsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exceptionPrincipalsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference) InternalValue() *IamDenyPolicyRulesDenyRule {
	var returns *IamDenyPolicyRulesDenyRule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIamDenyPolicyRulesDenyRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IamDenyPolicyRulesDenyRuleOutputReference {
	_init_.Initialize()

	if err := validateNewIamDenyPolicyRulesDenyRuleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.iamDenyPolicy.IamDenyPolicyRulesDenyRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIamDenyPolicyRulesDenyRuleOutputReference_Override(i IamDenyPolicyRulesDenyRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.iamDenyPolicy.IamDenyPolicyRulesDenyRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference)SetDeniedPermissions(val *[]*string) {
	if err := j.validateSetDeniedPermissionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deniedPermissions",
		val,
	)
}

func (j *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference)SetDeniedPrincipals(val *[]*string) {
	if err := j.validateSetDeniedPrincipalsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deniedPrincipals",
		val,
	)
}

func (j *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference)SetExceptionPermissions(val *[]*string) {
	if err := j.validateSetExceptionPermissionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exceptionPermissions",
		val,
	)
}

func (j *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference)SetExceptionPrincipals(val *[]*string) {
	if err := j.validateSetExceptionPrincipalsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exceptionPrincipals",
		val,
	)
}

func (j *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference)SetInternalValue(val *IamDenyPolicyRulesDenyRule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference) PutDenialCondition(value *IamDenyPolicyRulesDenyRuleDenialCondition) {
	if err := i.validatePutDenialConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putDenialCondition",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference) ResetDenialCondition() {
	_jsii_.InvokeVoid(
		i,
		"resetDenialCondition",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference) ResetDeniedPermissions() {
	_jsii_.InvokeVoid(
		i,
		"resetDeniedPermissions",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference) ResetDeniedPrincipals() {
	_jsii_.InvokeVoid(
		i,
		"resetDeniedPrincipals",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference) ResetExceptionPermissions() {
	_jsii_.InvokeVoid(
		i,
		"resetExceptionPermissions",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference) ResetExceptionPrincipals() {
	_jsii_.InvokeVoid(
		i,
		"resetExceptionPrincipals",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (i *jsiiProxy_IamDenyPolicyRulesDenyRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

