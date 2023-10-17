// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package billingbudget

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v12/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v12/billingbudget/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BillingBudgetAllUpdatesRuleOutputReference interface {
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
	DisableDefaultIamRecipients() interface{}
	SetDisableDefaultIamRecipients(val interface{})
	DisableDefaultIamRecipientsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *BillingBudgetAllUpdatesRule
	SetInternalValue(val *BillingBudgetAllUpdatesRule)
	MonitoringNotificationChannels() *[]*string
	SetMonitoringNotificationChannels(val *[]*string)
	MonitoringNotificationChannelsInput() *[]*string
	PubsubTopic() *string
	SetPubsubTopic(val *string)
	PubsubTopicInput() *string
	SchemaVersion() *string
	SetSchemaVersion(val *string)
	SchemaVersionInput() *string
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
	ResetDisableDefaultIamRecipients()
	ResetMonitoringNotificationChannels()
	ResetPubsubTopic()
	ResetSchemaVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BillingBudgetAllUpdatesRuleOutputReference
type jsiiProxy_BillingBudgetAllUpdatesRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BillingBudgetAllUpdatesRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetAllUpdatesRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetAllUpdatesRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetAllUpdatesRuleOutputReference) DisableDefaultIamRecipients() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableDefaultIamRecipients",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetAllUpdatesRuleOutputReference) DisableDefaultIamRecipientsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableDefaultIamRecipientsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetAllUpdatesRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetAllUpdatesRuleOutputReference) InternalValue() *BillingBudgetAllUpdatesRule {
	var returns *BillingBudgetAllUpdatesRule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetAllUpdatesRuleOutputReference) MonitoringNotificationChannels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"monitoringNotificationChannels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetAllUpdatesRuleOutputReference) MonitoringNotificationChannelsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"monitoringNotificationChannelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetAllUpdatesRuleOutputReference) PubsubTopic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pubsubTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetAllUpdatesRuleOutputReference) PubsubTopicInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pubsubTopicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetAllUpdatesRuleOutputReference) SchemaVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetAllUpdatesRuleOutputReference) SchemaVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetAllUpdatesRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingBudgetAllUpdatesRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBillingBudgetAllUpdatesRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BillingBudgetAllUpdatesRuleOutputReference {
	_init_.Initialize()

	if err := validateNewBillingBudgetAllUpdatesRuleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BillingBudgetAllUpdatesRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.billingBudget.BillingBudgetAllUpdatesRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBillingBudgetAllUpdatesRuleOutputReference_Override(b BillingBudgetAllUpdatesRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.billingBudget.BillingBudgetAllUpdatesRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BillingBudgetAllUpdatesRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BillingBudgetAllUpdatesRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BillingBudgetAllUpdatesRuleOutputReference)SetDisableDefaultIamRecipients(val interface{}) {
	if err := j.validateSetDisableDefaultIamRecipientsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableDefaultIamRecipients",
		val,
	)
}

func (j *jsiiProxy_BillingBudgetAllUpdatesRuleOutputReference)SetInternalValue(val *BillingBudgetAllUpdatesRule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BillingBudgetAllUpdatesRuleOutputReference)SetMonitoringNotificationChannels(val *[]*string) {
	if err := j.validateSetMonitoringNotificationChannelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoringNotificationChannels",
		val,
	)
}

func (j *jsiiProxy_BillingBudgetAllUpdatesRuleOutputReference)SetPubsubTopic(val *string) {
	if err := j.validateSetPubsubTopicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pubsubTopic",
		val,
	)
}

func (j *jsiiProxy_BillingBudgetAllUpdatesRuleOutputReference)SetSchemaVersion(val *string) {
	if err := j.validateSetSchemaVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schemaVersion",
		val,
	)
}

func (j *jsiiProxy_BillingBudgetAllUpdatesRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BillingBudgetAllUpdatesRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BillingBudgetAllUpdatesRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingBudgetAllUpdatesRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BillingBudgetAllUpdatesRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BillingBudgetAllUpdatesRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BillingBudgetAllUpdatesRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BillingBudgetAllUpdatesRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BillingBudgetAllUpdatesRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BillingBudgetAllUpdatesRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BillingBudgetAllUpdatesRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BillingBudgetAllUpdatesRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BillingBudgetAllUpdatesRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingBudgetAllUpdatesRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BillingBudgetAllUpdatesRuleOutputReference) ResetDisableDefaultIamRecipients() {
	_jsii_.InvokeVoid(
		b,
		"resetDisableDefaultIamRecipients",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BillingBudgetAllUpdatesRuleOutputReference) ResetMonitoringNotificationChannels() {
	_jsii_.InvokeVoid(
		b,
		"resetMonitoringNotificationChannels",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BillingBudgetAllUpdatesRuleOutputReference) ResetPubsubTopic() {
	_jsii_.InvokeVoid(
		b,
		"resetPubsubTopic",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BillingBudgetAllUpdatesRuleOutputReference) ResetSchemaVersion() {
	_jsii_.InvokeVoid(
		b,
		"resetSchemaVersion",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BillingBudgetAllUpdatesRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (b *jsiiProxy_BillingBudgetAllUpdatesRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

