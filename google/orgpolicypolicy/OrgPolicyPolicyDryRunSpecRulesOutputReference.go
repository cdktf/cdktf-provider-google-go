// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package orgpolicypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v13/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v13/orgpolicypolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrgPolicyPolicyDryRunSpecRulesOutputReference interface {
	cdktf.ComplexObject
	AllowAll() *string
	SetAllowAll(val *string)
	AllowAllInput() *string
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
	Condition() OrgPolicyPolicyDryRunSpecRulesConditionOutputReference
	ConditionInput() *OrgPolicyPolicyDryRunSpecRulesCondition
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DenyAll() *string
	SetDenyAll(val *string)
	DenyAllInput() *string
	Enforce() *string
	SetEnforce(val *string)
	EnforceInput() *string
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
	Values() OrgPolicyPolicyDryRunSpecRulesValuesOutputReference
	ValuesInput() *OrgPolicyPolicyDryRunSpecRulesValues
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
	PutCondition(value *OrgPolicyPolicyDryRunSpecRulesCondition)
	PutValues(value *OrgPolicyPolicyDryRunSpecRulesValues)
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

// The jsii proxy struct for OrgPolicyPolicyDryRunSpecRulesOutputReference
type jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference) AllowAll() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference) AllowAllInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference) Condition() OrgPolicyPolicyDryRunSpecRulesConditionOutputReference {
	var returns OrgPolicyPolicyDryRunSpecRulesConditionOutputReference
	_jsii_.Get(
		j,
		"condition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference) ConditionInput() *OrgPolicyPolicyDryRunSpecRulesCondition {
	var returns *OrgPolicyPolicyDryRunSpecRulesCondition
	_jsii_.Get(
		j,
		"conditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference) DenyAll() *string {
	var returns *string
	_jsii_.Get(
		j,
		"denyAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference) DenyAllInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"denyAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference) Enforce() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enforce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference) EnforceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enforceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference) Values() OrgPolicyPolicyDryRunSpecRulesValuesOutputReference {
	var returns OrgPolicyPolicyDryRunSpecRulesValuesOutputReference
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference) ValuesInput() *OrgPolicyPolicyDryRunSpecRulesValues {
	var returns *OrgPolicyPolicyDryRunSpecRulesValues
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewOrgPolicyPolicyDryRunSpecRulesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) OrgPolicyPolicyDryRunSpecRulesOutputReference {
	_init_.Initialize()

	if err := validateNewOrgPolicyPolicyDryRunSpecRulesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.orgPolicyPolicy.OrgPolicyPolicyDryRunSpecRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewOrgPolicyPolicyDryRunSpecRulesOutputReference_Override(o OrgPolicyPolicyDryRunSpecRulesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.orgPolicyPolicy.OrgPolicyPolicyDryRunSpecRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference)SetAllowAll(val *string) {
	if err := j.validateSetAllowAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowAll",
		val,
	)
}

func (j *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference)SetDenyAll(val *string) {
	if err := j.validateSetDenyAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"denyAll",
		val,
	)
}

func (j *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference)SetEnforce(val *string) {
	if err := j.validateSetEnforceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforce",
		val,
	)
}

func (j *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference) PutCondition(value *OrgPolicyPolicyDryRunSpecRulesCondition) {
	if err := o.validatePutConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putCondition",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference) PutValues(value *OrgPolicyPolicyDryRunSpecRulesValues) {
	if err := o.validatePutValuesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putValues",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference) ResetAllowAll() {
	_jsii_.InvokeVoid(
		o,
		"resetAllowAll",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference) ResetCondition() {
	_jsii_.InvokeVoid(
		o,
		"resetCondition",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference) ResetDenyAll() {
	_jsii_.InvokeVoid(
		o,
		"resetDenyAll",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference) ResetEnforce() {
	_jsii_.InvokeVoid(
		o,
		"resetEnforce",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference) ResetValues() {
	_jsii_.InvokeVoid(
		o,
		"resetValues",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OrgPolicyPolicyDryRunSpecRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

