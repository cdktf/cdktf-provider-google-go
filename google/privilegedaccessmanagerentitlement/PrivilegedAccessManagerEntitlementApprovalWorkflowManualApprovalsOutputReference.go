// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privilegedaccessmanagerentitlement

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/privilegedaccessmanagerentitlement/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsOutputReference interface {
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
	InternalValue() *PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovals
	SetInternalValue(val *PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovals)
	RequireApproverJustification() interface{}
	SetRequireApproverJustification(val interface{})
	RequireApproverJustificationInput() interface{}
	Steps() PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsList
	StepsInput() interface{}
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
	PutSteps(value interface{})
	ResetRequireApproverJustification()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsOutputReference
type jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsOutputReference) InternalValue() *PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovals {
	var returns *PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovals
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsOutputReference) RequireApproverJustification() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireApproverJustification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsOutputReference) RequireApproverJustificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireApproverJustificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsOutputReference) Steps() PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsList {
	var returns PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsList
	_jsii_.Get(
		j,
		"steps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsOutputReference) StepsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stepsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsOutputReference {
	_init_.Initialize()

	if err := validateNewPrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.privilegedAccessManagerEntitlement.PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsOutputReference_Override(p PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.privilegedAccessManagerEntitlement.PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsOutputReference)SetInternalValue(val *PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovals) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsOutputReference)SetRequireApproverJustification(val interface{}) {
	if err := j.validateSetRequireApproverJustificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireApproverJustification",
		val,
	)
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsOutputReference) PutSteps(value interface{}) {
	if err := p.validatePutStepsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSteps",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsOutputReference) ResetRequireApproverJustification() {
	_jsii_.InvokeVoid(
		p,
		"resetRequireApproverJustification",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

