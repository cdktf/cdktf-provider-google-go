// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pubsubtopic

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/pubsubtopic/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PubsubTopicIngestionDataSourceSettingsOutputReference interface {
	cdktf.ComplexObject
	AwsKinesis() PubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference
	AwsKinesisInput() *PubsubTopicIngestionDataSourceSettingsAwsKinesis
	AwsMsk() PubsubTopicIngestionDataSourceSettingsAwsMskOutputReference
	AwsMskInput() *PubsubTopicIngestionDataSourceSettingsAwsMsk
	AzureEventHubs() PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference
	AzureEventHubsInput() *PubsubTopicIngestionDataSourceSettingsAzureEventHubs
	CloudStorage() PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference
	CloudStorageInput() *PubsubTopicIngestionDataSourceSettingsCloudStorage
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
	ConfluentCloud() PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference
	ConfluentCloudInput() *PubsubTopicIngestionDataSourceSettingsConfluentCloud
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *PubsubTopicIngestionDataSourceSettings
	SetInternalValue(val *PubsubTopicIngestionDataSourceSettings)
	PlatformLogsSettings() PubsubTopicIngestionDataSourceSettingsPlatformLogsSettingsOutputReference
	PlatformLogsSettingsInput() *PubsubTopicIngestionDataSourceSettingsPlatformLogsSettings
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
	PutAwsKinesis(value *PubsubTopicIngestionDataSourceSettingsAwsKinesis)
	PutAwsMsk(value *PubsubTopicIngestionDataSourceSettingsAwsMsk)
	PutAzureEventHubs(value *PubsubTopicIngestionDataSourceSettingsAzureEventHubs)
	PutCloudStorage(value *PubsubTopicIngestionDataSourceSettingsCloudStorage)
	PutConfluentCloud(value *PubsubTopicIngestionDataSourceSettingsConfluentCloud)
	PutPlatformLogsSettings(value *PubsubTopicIngestionDataSourceSettingsPlatformLogsSettings)
	ResetAwsKinesis()
	ResetAwsMsk()
	ResetAzureEventHubs()
	ResetCloudStorage()
	ResetConfluentCloud()
	ResetPlatformLogsSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PubsubTopicIngestionDataSourceSettingsOutputReference
type jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference) AwsKinesis() PubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference {
	var returns PubsubTopicIngestionDataSourceSettingsAwsKinesisOutputReference
	_jsii_.Get(
		j,
		"awsKinesis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference) AwsKinesisInput() *PubsubTopicIngestionDataSourceSettingsAwsKinesis {
	var returns *PubsubTopicIngestionDataSourceSettingsAwsKinesis
	_jsii_.Get(
		j,
		"awsKinesisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference) AwsMsk() PubsubTopicIngestionDataSourceSettingsAwsMskOutputReference {
	var returns PubsubTopicIngestionDataSourceSettingsAwsMskOutputReference
	_jsii_.Get(
		j,
		"awsMsk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference) AwsMskInput() *PubsubTopicIngestionDataSourceSettingsAwsMsk {
	var returns *PubsubTopicIngestionDataSourceSettingsAwsMsk
	_jsii_.Get(
		j,
		"awsMskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference) AzureEventHubs() PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference {
	var returns PubsubTopicIngestionDataSourceSettingsAzureEventHubsOutputReference
	_jsii_.Get(
		j,
		"azureEventHubs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference) AzureEventHubsInput() *PubsubTopicIngestionDataSourceSettingsAzureEventHubs {
	var returns *PubsubTopicIngestionDataSourceSettingsAzureEventHubs
	_jsii_.Get(
		j,
		"azureEventHubsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference) CloudStorage() PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference {
	var returns PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference
	_jsii_.Get(
		j,
		"cloudStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference) CloudStorageInput() *PubsubTopicIngestionDataSourceSettingsCloudStorage {
	var returns *PubsubTopicIngestionDataSourceSettingsCloudStorage
	_jsii_.Get(
		j,
		"cloudStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference) ConfluentCloud() PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference {
	var returns PubsubTopicIngestionDataSourceSettingsConfluentCloudOutputReference
	_jsii_.Get(
		j,
		"confluentCloud",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference) ConfluentCloudInput() *PubsubTopicIngestionDataSourceSettingsConfluentCloud {
	var returns *PubsubTopicIngestionDataSourceSettingsConfluentCloud
	_jsii_.Get(
		j,
		"confluentCloudInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference) InternalValue() *PubsubTopicIngestionDataSourceSettings {
	var returns *PubsubTopicIngestionDataSourceSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference) PlatformLogsSettings() PubsubTopicIngestionDataSourceSettingsPlatformLogsSettingsOutputReference {
	var returns PubsubTopicIngestionDataSourceSettingsPlatformLogsSettingsOutputReference
	_jsii_.Get(
		j,
		"platformLogsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference) PlatformLogsSettingsInput() *PubsubTopicIngestionDataSourceSettingsPlatformLogsSettings {
	var returns *PubsubTopicIngestionDataSourceSettingsPlatformLogsSettings
	_jsii_.Get(
		j,
		"platformLogsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPubsubTopicIngestionDataSourceSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PubsubTopicIngestionDataSourceSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewPubsubTopicIngestionDataSourceSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.pubsubTopic.PubsubTopicIngestionDataSourceSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPubsubTopicIngestionDataSourceSettingsOutputReference_Override(p PubsubTopicIngestionDataSourceSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.pubsubTopic.PubsubTopicIngestionDataSourceSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference)SetInternalValue(val *PubsubTopicIngestionDataSourceSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference) PutAwsKinesis(value *PubsubTopicIngestionDataSourceSettingsAwsKinesis) {
	if err := p.validatePutAwsKinesisParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAwsKinesis",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference) PutAwsMsk(value *PubsubTopicIngestionDataSourceSettingsAwsMsk) {
	if err := p.validatePutAwsMskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAwsMsk",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference) PutAzureEventHubs(value *PubsubTopicIngestionDataSourceSettingsAzureEventHubs) {
	if err := p.validatePutAzureEventHubsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAzureEventHubs",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference) PutCloudStorage(value *PubsubTopicIngestionDataSourceSettingsCloudStorage) {
	if err := p.validatePutCloudStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCloudStorage",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference) PutConfluentCloud(value *PubsubTopicIngestionDataSourceSettingsConfluentCloud) {
	if err := p.validatePutConfluentCloudParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putConfluentCloud",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference) PutPlatformLogsSettings(value *PubsubTopicIngestionDataSourceSettingsPlatformLogsSettings) {
	if err := p.validatePutPlatformLogsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putPlatformLogsSettings",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference) ResetAwsKinesis() {
	_jsii_.InvokeVoid(
		p,
		"resetAwsKinesis",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference) ResetAwsMsk() {
	_jsii_.InvokeVoid(
		p,
		"resetAwsMsk",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference) ResetAzureEventHubs() {
	_jsii_.InvokeVoid(
		p,
		"resetAzureEventHubs",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference) ResetCloudStorage() {
	_jsii_.InvokeVoid(
		p,
		"resetCloudStorage",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference) ResetConfluentCloud() {
	_jsii_.InvokeVoid(
		p,
		"resetConfluentCloud",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference) ResetPlatformLogsSettings() {
	_jsii_.InvokeVoid(
		p,
		"resetPlatformLogsSettings",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

