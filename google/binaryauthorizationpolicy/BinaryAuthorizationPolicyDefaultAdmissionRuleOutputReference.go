// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package binaryauthorizationpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v10/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v10/binaryauthorizationpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference interface {
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
	EnforcementMode() *string
	SetEnforcementMode(val *string)
	EnforcementModeInput() *string
	EvaluationMode() *string
	SetEvaluationMode(val *string)
	EvaluationModeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *BinaryAuthorizationPolicyDefaultAdmissionRule
	SetInternalValue(val *BinaryAuthorizationPolicyDefaultAdmissionRule)
	RequireAttestationsBy() *[]*string
	SetRequireAttestationsBy(val *[]*string)
	RequireAttestationsByInput() *[]*string
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
	ResetRequireAttestationsBy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference
type jsiiProxy_BinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) EnforcementMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enforcementMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) EnforcementModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enforcementModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) EvaluationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) EvaluationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) InternalValue() *BinaryAuthorizationPolicyDefaultAdmissionRule {
	var returns *BinaryAuthorizationPolicyDefaultAdmissionRule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) RequireAttestationsBy() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requireAttestationsBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) RequireAttestationsByInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requireAttestationsByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference {
	_init_.Initialize()

	if err := validateNewBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.binaryAuthorizationPolicy.BinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference_Override(b BinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.binaryAuthorizationPolicy.BinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference)SetEnforcementMode(val *string) {
	if err := j.validateSetEnforcementModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforcementMode",
		val,
	)
}

func (j *jsiiProxy_BinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference)SetEvaluationMode(val *string) {
	if err := j.validateSetEvaluationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evaluationMode",
		val,
	)
}

func (j *jsiiProxy_BinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference)SetInternalValue(val *BinaryAuthorizationPolicyDefaultAdmissionRule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference)SetRequireAttestationsBy(val *[]*string) {
	if err := j.validateSetRequireAttestationsByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireAttestationsBy",
		val,
	)
}

func (j *jsiiProxy_BinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) ResetRequireAttestationsBy() {
	_jsii_.InvokeVoid(
		b,
		"resetRequireAttestationsBy",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := b.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BinaryAuthorizationPolicyDefaultAdmissionRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

