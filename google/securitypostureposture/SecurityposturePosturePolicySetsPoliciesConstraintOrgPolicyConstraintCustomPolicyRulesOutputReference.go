// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitypostureposture

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v13/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v13/securitypostureposture/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference interface {
	cdktf.ComplexObject
	AllowAll() interface{}
	SetAllowAll(val interface{})
	AllowAllInput() interface{}
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
	Condition() SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesConditionOutputReference
	ConditionInput() *SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesCondition
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DenyAll() interface{}
	SetDenyAll(val interface{})
	DenyAllInput() interface{}
	Enforce() interface{}
	SetEnforce(val interface{})
	EnforceInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Values() SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesValuesOutputReference
	ValuesInput() *SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesValues
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
	PutCondition(value *SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesCondition)
	PutValues(value *SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesValues)
	ResetAllowAll()
	ResetCondition()
	ResetDenyAll()
	ResetEnforce()
	ResetValues()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference
type jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference) AllowAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference) AllowAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference) Condition() SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesConditionOutputReference {
	var returns SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesConditionOutputReference
	_jsii_.Get(
		j,
		"condition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference) ConditionInput() *SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesCondition {
	var returns *SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesCondition
	_jsii_.Get(
		j,
		"conditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference) DenyAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"denyAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference) DenyAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"denyAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference) Enforce() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference) EnforceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference) Values() SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesValuesOutputReference {
	var returns SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesValuesOutputReference
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference) ValuesInput() *SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesValues {
	var returns *SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesValues
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewSecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference {
	_init_.Initialize()

	if err := validateNewSecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.securityposturePosture.SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference_Override(s SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.securityposturePosture.SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference)SetAllowAll(val interface{}) {
	if err := j.validateSetAllowAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowAll",
		val,
	)
}

func (j *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference)SetDenyAll(val interface{}) {
	if err := j.validateSetDenyAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"denyAll",
		val,
	)
}

func (j *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference)SetEnforce(val interface{}) {
	if err := j.validateSetEnforceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforce",
		val,
	)
}

func (j *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference) PutCondition(value *SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesCondition) {
	if err := s.validatePutConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCondition",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference) PutValues(value *SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesValues) {
	if err := s.validatePutValuesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putValues",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference) ResetAllowAll() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowAll",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference) ResetCondition() {
	_jsii_.InvokeVoid(
		s,
		"resetCondition",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference) ResetDenyAll() {
	_jsii_.InvokeVoid(
		s,
		"resetDenyAll",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference) ResetEnforce() {
	_jsii_.InvokeVoid(
		s,
		"resetEnforce",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference) ResetValues() {
	_jsii_.InvokeVoid(
		s,
		"resetValues",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

