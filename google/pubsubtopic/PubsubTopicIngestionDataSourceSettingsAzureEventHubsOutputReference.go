// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pubsubtopic

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/pubsubtopic/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference interface {
	cdktf.ComplexObject
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
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
	EventHub() *string
	SetEventHub(val *string)
	EventHubInput() *string
	// Experimental.
	Fqn() *string
	GcpServiceAccount() *string
	SetGcpServiceAccount(val *string)
	GcpServiceAccountInput() *string
	InternalValue() *PubsubTopicIngestionDataSourceSettingsAzureEventHubs
	SetInternalValue(val *PubsubTopicIngestionDataSourceSettingsAzureEventHubs)
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	ResourceGroup() *string
	SetResourceGroup(val *string)
	ResourceGroupInput() *string
	SubscriptionId() *string
	SetSubscriptionId(val *string)
	SubscriptionIdInput() *string
	TenantId() *string
	SetTenantId(val *string)
	TenantIdInput() *string
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
	ResetClientId()
	ResetEventHub()
	ResetGcpServiceAccount()
	ResetNamespace()
	ResetResourceGroup()
	ResetSubscriptionId()
	ResetTenantId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference
type jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) EventHub() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventHub",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) EventHubInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventHubInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) GcpServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcpServiceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) GcpServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcpServiceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) InternalValue() *PubsubTopicIngestionDataSourceSettingsAzureEventHubs {
	var returns *PubsubTopicIngestionDataSourceSettingsAzureEventHubs
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) ResourceGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) ResourceGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) SubscriptionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) SubscriptionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) TenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference {
	_init_.Initialize()

	if err := validateNewPubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.pubsubTopic.PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference_Override(p PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.pubsubTopic.PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference)SetEventHub(val *string) {
	if err := j.validateSetEventHubParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventHub",
		val,
	)
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference)SetGcpServiceAccount(val *string) {
	if err := j.validateSetGcpServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gcpServiceAccount",
		val,
	)
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference)SetInternalValue(val *PubsubTopicIngestionDataSourceSettingsAzureEventHubs) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference)SetResourceGroup(val *string) {
	if err := j.validateSetResourceGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroup",
		val,
	)
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference)SetSubscriptionId(val *string) {
	if err := j.validateSetSubscriptionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subscriptionId",
		val,
	)
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference)SetTenantId(val *string) {
	if err := j.validateSetTenantIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenantId",
		val,
	)
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) ResetClientId() {
	_jsii_.InvokeVoid(
		p,
		"resetClientId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) ResetEventHub() {
	_jsii_.InvokeVoid(
		p,
		"resetEventHub",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) ResetGcpServiceAccount() {
	_jsii_.InvokeVoid(
		p,
		"resetGcpServiceAccount",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) ResetNamespace() {
	_jsii_.InvokeVoid(
		p,
		"resetNamespace",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) ResetResourceGroup() {
	_jsii_.InvokeVoid(
		p,
		"resetResourceGroup",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) ResetSubscriptionId() {
	_jsii_.InvokeVoid(
		p,
		"resetSubscriptionId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) ResetTenantId() {
	_jsii_.InvokeVoid(
		p,
		"resetTenantId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

