// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privilegedaccessmanagerentitlement

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/privilegedaccessmanagerentitlement/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference interface {
	cdktf.ComplexObject
	ApprovalsNeeded() *float64
	SetApprovalsNeeded(val *float64)
	ApprovalsNeededInput() *float64
	ApproverEmailRecipients() *[]*string
	SetApproverEmailRecipients(val *[]*string)
	ApproverEmailRecipientsInput() *[]*string
	Approvers() PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsApproversOutputReference
	ApproversInput() *PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsApprovers
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
	PutApprovers(value *PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsApprovers)
	ResetApprovalsNeeded()
	ResetApproverEmailRecipients()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference
type jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) ApprovalsNeeded() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"approvalsNeeded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) ApprovalsNeededInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"approvalsNeededInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) ApproverEmailRecipients() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"approverEmailRecipients",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) ApproverEmailRecipientsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"approverEmailRecipientsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) Approvers() PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsApproversOutputReference {
	var returns PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsApproversOutputReference
	_jsii_.Get(
		j,
		"approvers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) ApproversInput() *PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsApprovers {
	var returns *PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsApprovers
	_jsii_.Get(
		j,
		"approversInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference {
	_init_.Initialize()

	if err := validateNewPrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.privilegedAccessManagerEntitlement.PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference_Override(p PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.privilegedAccessManagerEntitlement.PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference)SetApprovalsNeeded(val *float64) {
	if err := j.validateSetApprovalsNeededParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"approvalsNeeded",
		val,
	)
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference)SetApproverEmailRecipients(val *[]*string) {
	if err := j.validateSetApproverEmailRecipientsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"approverEmailRecipients",
		val,
	)
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) PutApprovers(value *PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsApprovers) {
	if err := p.validatePutApproversParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putApprovers",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) ResetApprovalsNeeded() {
	_jsii_.InvokeVoid(
		p,
		"resetApprovalsNeeded",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) ResetApproverEmailRecipients() {
	_jsii_.InvokeVoid(
		p,
		"resetApproverEmailRecipients",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementApprovalWorkflowManualApprovalsStepsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

