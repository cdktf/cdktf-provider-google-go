// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privilegedaccessmanagerentitlement

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/privilegedaccessmanagerentitlement/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference interface {
	cdktf.ComplexObject
	AdminEmailRecipients() *[]*string
	SetAdminEmailRecipients(val *[]*string)
	AdminEmailRecipientsInput() *[]*string
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
	InternalValue() *PrivilegedAccessManagerEntitlementAdditionalNotificationTargets
	SetInternalValue(val *PrivilegedAccessManagerEntitlementAdditionalNotificationTargets)
	RequesterEmailRecipients() *[]*string
	SetRequesterEmailRecipients(val *[]*string)
	RequesterEmailRecipientsInput() *[]*string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	ResetAdminEmailRecipients()
	ResetRequesterEmailRecipients()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference
type jsiiProxy_PrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference) AdminEmailRecipients() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adminEmailRecipients",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference) AdminEmailRecipientsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adminEmailRecipientsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference) InternalValue() *PrivilegedAccessManagerEntitlementAdditionalNotificationTargets {
	var returns *PrivilegedAccessManagerEntitlementAdditionalNotificationTargets
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference) RequesterEmailRecipients() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requesterEmailRecipients",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference) RequesterEmailRecipientsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requesterEmailRecipientsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference {
	_init_.Initialize()

	if err := validateNewPrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.privilegedAccessManagerEntitlement.PrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference_Override(p PrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.privilegedAccessManagerEntitlement.PrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference)SetAdminEmailRecipients(val *[]*string) {
	if err := j.validateSetAdminEmailRecipientsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminEmailRecipients",
		val,
	)
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference)SetInternalValue(val *PrivilegedAccessManagerEntitlementAdditionalNotificationTargets) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference)SetRequesterEmailRecipients(val *[]*string) {
	if err := j.validateSetRequesterEmailRecipientsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requesterEmailRecipients",
		val,
	)
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference) ResetAdminEmailRecipients() {
	_jsii_.InvokeVoid(
		p,
		"resetAdminEmailRecipients",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference) ResetRequesterEmailRecipients() {
	_jsii_.InvokeVoid(
		p,
		"resetRequesterEmailRecipients",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PrivilegedAccessManagerEntitlementAdditionalNotificationTargetsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

