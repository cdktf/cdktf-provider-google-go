// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitypostureposture

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/securitypostureposture/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintOutputReference interface {
	cdktf.ComplexObject
	CannedConstraintId() *string
	SetCannedConstraintId(val *string)
	CannedConstraintIdInput() *string
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
	InternalValue() *SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraint
	SetInternalValue(val *SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraint)
	PolicyRules() SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintPolicyRulesList
	PolicyRulesInput() interface{}
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
	PutPolicyRules(value interface{})
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintOutputReference
type jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintOutputReference) CannedConstraintId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cannedConstraintId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintOutputReference) CannedConstraintIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cannedConstraintIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintOutputReference) InternalValue() *SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraint {
	var returns *SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraint
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintOutputReference) PolicyRules() SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintPolicyRulesList {
	var returns SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintPolicyRulesList
	_jsii_.Get(
		j,
		"policyRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintOutputReference) PolicyRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"policyRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintOutputReference {
	_init_.Initialize()

	if err := validateNewSecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.securityposturePosture.SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintOutputReference_Override(s SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.securityposturePosture.SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintOutputReference)SetCannedConstraintId(val *string) {
	if err := j.validateSetCannedConstraintIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cannedConstraintId",
		val,
	)
}

func (j *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintOutputReference)SetInternalValue(val *SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraint) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintOutputReference) PutPolicyRules(value interface{}) {
	if err := s.validatePutPolicyRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPolicyRules",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

