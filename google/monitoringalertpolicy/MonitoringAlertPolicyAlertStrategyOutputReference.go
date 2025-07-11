// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoringalertpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/monitoringalertpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MonitoringAlertPolicyAlertStrategyOutputReference interface {
	cdktf.ComplexObject
	AutoClose() *string
	SetAutoClose(val *string)
	AutoCloseInput() *string
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
	InternalValue() *MonitoringAlertPolicyAlertStrategy
	SetInternalValue(val *MonitoringAlertPolicyAlertStrategy)
	NotificationChannelStrategy() MonitoringAlertPolicyAlertStrategyNotificationChannelStrategyList
	NotificationChannelStrategyInput() interface{}
	NotificationPrompts() *[]*string
	SetNotificationPrompts(val *[]*string)
	NotificationPromptsInput() *[]*string
	NotificationRateLimit() MonitoringAlertPolicyAlertStrategyNotificationRateLimitOutputReference
	NotificationRateLimitInput() *MonitoringAlertPolicyAlertStrategyNotificationRateLimit
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
	PutNotificationChannelStrategy(value interface{})
	PutNotificationRateLimit(value *MonitoringAlertPolicyAlertStrategyNotificationRateLimit)
	ResetAutoClose()
	ResetNotificationChannelStrategy()
	ResetNotificationPrompts()
	ResetNotificationRateLimit()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitoringAlertPolicyAlertStrategyOutputReference
type jsiiProxy_MonitoringAlertPolicyAlertStrategyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitoringAlertPolicyAlertStrategyOutputReference) AutoClose() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoClose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringAlertPolicyAlertStrategyOutputReference) AutoCloseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoCloseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringAlertPolicyAlertStrategyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringAlertPolicyAlertStrategyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringAlertPolicyAlertStrategyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringAlertPolicyAlertStrategyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringAlertPolicyAlertStrategyOutputReference) InternalValue() *MonitoringAlertPolicyAlertStrategy {
	var returns *MonitoringAlertPolicyAlertStrategy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringAlertPolicyAlertStrategyOutputReference) NotificationChannelStrategy() MonitoringAlertPolicyAlertStrategyNotificationChannelStrategyList {
	var returns MonitoringAlertPolicyAlertStrategyNotificationChannelStrategyList
	_jsii_.Get(
		j,
		"notificationChannelStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringAlertPolicyAlertStrategyOutputReference) NotificationChannelStrategyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notificationChannelStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringAlertPolicyAlertStrategyOutputReference) NotificationPrompts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notificationPrompts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringAlertPolicyAlertStrategyOutputReference) NotificationPromptsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notificationPromptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringAlertPolicyAlertStrategyOutputReference) NotificationRateLimit() MonitoringAlertPolicyAlertStrategyNotificationRateLimitOutputReference {
	var returns MonitoringAlertPolicyAlertStrategyNotificationRateLimitOutputReference
	_jsii_.Get(
		j,
		"notificationRateLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringAlertPolicyAlertStrategyOutputReference) NotificationRateLimitInput() *MonitoringAlertPolicyAlertStrategyNotificationRateLimit {
	var returns *MonitoringAlertPolicyAlertStrategyNotificationRateLimit
	_jsii_.Get(
		j,
		"notificationRateLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringAlertPolicyAlertStrategyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringAlertPolicyAlertStrategyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMonitoringAlertPolicyAlertStrategyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MonitoringAlertPolicyAlertStrategyOutputReference {
	_init_.Initialize()

	if err := validateNewMonitoringAlertPolicyAlertStrategyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitoringAlertPolicyAlertStrategyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.monitoringAlertPolicy.MonitoringAlertPolicyAlertStrategyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitoringAlertPolicyAlertStrategyOutputReference_Override(m MonitoringAlertPolicyAlertStrategyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.monitoringAlertPolicy.MonitoringAlertPolicyAlertStrategyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitoringAlertPolicyAlertStrategyOutputReference)SetAutoClose(val *string) {
	if err := j.validateSetAutoCloseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoClose",
		val,
	)
}

func (j *jsiiProxy_MonitoringAlertPolicyAlertStrategyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitoringAlertPolicyAlertStrategyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitoringAlertPolicyAlertStrategyOutputReference)SetInternalValue(val *MonitoringAlertPolicyAlertStrategy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitoringAlertPolicyAlertStrategyOutputReference)SetNotificationPrompts(val *[]*string) {
	if err := j.validateSetNotificationPromptsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationPrompts",
		val,
	)
}

func (j *jsiiProxy_MonitoringAlertPolicyAlertStrategyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitoringAlertPolicyAlertStrategyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MonitoringAlertPolicyAlertStrategyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringAlertPolicyAlertStrategyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringAlertPolicyAlertStrategyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringAlertPolicyAlertStrategyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringAlertPolicyAlertStrategyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringAlertPolicyAlertStrategyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringAlertPolicyAlertStrategyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringAlertPolicyAlertStrategyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringAlertPolicyAlertStrategyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringAlertPolicyAlertStrategyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringAlertPolicyAlertStrategyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringAlertPolicyAlertStrategyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringAlertPolicyAlertStrategyOutputReference) PutNotificationChannelStrategy(value interface{}) {
	if err := m.validatePutNotificationChannelStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putNotificationChannelStrategy",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitoringAlertPolicyAlertStrategyOutputReference) PutNotificationRateLimit(value *MonitoringAlertPolicyAlertStrategyNotificationRateLimit) {
	if err := m.validatePutNotificationRateLimitParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putNotificationRateLimit",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitoringAlertPolicyAlertStrategyOutputReference) ResetAutoClose() {
	_jsii_.InvokeVoid(
		m,
		"resetAutoClose",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitoringAlertPolicyAlertStrategyOutputReference) ResetNotificationChannelStrategy() {
	_jsii_.InvokeVoid(
		m,
		"resetNotificationChannelStrategy",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitoringAlertPolicyAlertStrategyOutputReference) ResetNotificationPrompts() {
	_jsii_.InvokeVoid(
		m,
		"resetNotificationPrompts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitoringAlertPolicyAlertStrategyOutputReference) ResetNotificationRateLimit() {
	_jsii_.InvokeVoid(
		m,
		"resetNotificationRateLimit",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitoringAlertPolicyAlertStrategyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringAlertPolicyAlertStrategyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

