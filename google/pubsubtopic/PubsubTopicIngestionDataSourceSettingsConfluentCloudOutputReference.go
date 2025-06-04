// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pubsubtopic

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/pubsubtopic/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference interface {
	cdktf.ComplexObject
	BootstrapServer() *string
	SetBootstrapServer(val *string)
	BootstrapServerInput() *string
	ClusterId() *string
	SetClusterId(val *string)
	ClusterIdInput() *string
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
	GcpServiceAccount() *string
	SetGcpServiceAccount(val *string)
	GcpServiceAccountInput() *string
	IdentityPoolId() *string
	SetIdentityPoolId(val *string)
	IdentityPoolIdInput() *string
	InternalValue() *PubsubTopicIngestionDataSourceSettingsConfluentCloud
	SetInternalValue(val *PubsubTopicIngestionDataSourceSettingsConfluentCloud)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Topic() *string
	SetTopic(val *string)
	TopicInput() *string
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
	ResetClusterId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference
type jsiiProxy_PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference) BootstrapServer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference) BootstrapServerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference) GcpServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcpServiceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference) GcpServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcpServiceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference) IdentityPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference) IdentityPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityPoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference) InternalValue() *PubsubTopicIngestionDataSourceSettingsConfluentCloud {
	var returns *PubsubTopicIngestionDataSourceSettingsConfluentCloud
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference) Topic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference) TopicInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicInput",
		&returns,
	)
	return returns
}


func NewPubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference {
	_init_.Initialize()

	if err := validateNewPubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.pubsubTopic.PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference_Override(p PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.pubsubTopic.PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference)SetBootstrapServer(val *string) {
	if err := j.validateSetBootstrapServerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootstrapServer",
		val,
	)
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference)SetClusterId(val *string) {
	if err := j.validateSetClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference)SetGcpServiceAccount(val *string) {
	if err := j.validateSetGcpServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gcpServiceAccount",
		val,
	)
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference)SetIdentityPoolId(val *string) {
	if err := j.validateSetIdentityPoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityPoolId",
		val,
	)
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference)SetInternalValue(val *PubsubTopicIngestionDataSourceSettingsConfluentCloud) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference)SetTopic(val *string) {
	if err := j.validateSetTopicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topic",
		val,
	)
}

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference) ResetClusterId() {
	_jsii_.InvokeVoid(
		p,
		"resetClusterId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

