// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pubsubtopic

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/pubsubtopic/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference interface {
	cdktf.ComplexObject
	AvroFormat() PubsubTopicIngestionDataSourceSettingsCloudStorageAvroFormatOutputReference
	AvroFormatInput() *PubsubTopicIngestionDataSourceSettingsCloudStorageAvroFormat
	Bucket() *string
	SetBucket(val *string)
	BucketInput() *string
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
	InternalValue() *PubsubTopicIngestionDataSourceSettingsCloudStorage
	SetInternalValue(val *PubsubTopicIngestionDataSourceSettingsCloudStorage)
	MatchGlob() *string
	SetMatchGlob(val *string)
	MatchGlobInput() *string
	MinimumObjectCreateTime() *string
	SetMinimumObjectCreateTime(val *string)
	MinimumObjectCreateTimeInput() *string
	PubsubAvroFormat() PubsubTopicIngestionDataSourceSettingsCloudStoragePubsubAvroFormatOutputReference
	PubsubAvroFormatInput() *PubsubTopicIngestionDataSourceSettingsCloudStoragePubsubAvroFormat
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TextFormat() PubsubTopicIngestionDataSourceSettingsCloudStorageTextFormatOutputReference
	TextFormatInput() *PubsubTopicIngestionDataSourceSettingsCloudStorageTextFormat
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
	PutAvroFormat(value *PubsubTopicIngestionDataSourceSettingsCloudStorageAvroFormat)
	PutPubsubAvroFormat(value *PubsubTopicIngestionDataSourceSettingsCloudStoragePubsubAvroFormat)
	PutTextFormat(value *PubsubTopicIngestionDataSourceSettingsCloudStorageTextFormat)
	ResetAvroFormat()
	ResetMatchGlob()
	ResetMinimumObjectCreateTime()
	ResetPubsubAvroFormat()
	ResetTextFormat()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference
type jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) AvroFormat() PubsubTopicIngestionDataSourceSettingsCloudStorageAvroFormatOutputReference {
	var returns PubsubTopicIngestionDataSourceSettingsCloudStorageAvroFormatOutputReference
	_jsii_.Get(
		j,
		"avroFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) AvroFormatInput() *PubsubTopicIngestionDataSourceSettingsCloudStorageAvroFormat {
	var returns *PubsubTopicIngestionDataSourceSettingsCloudStorageAvroFormat
	_jsii_.Get(
		j,
		"avroFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) InternalValue() *PubsubTopicIngestionDataSourceSettingsCloudStorage {
	var returns *PubsubTopicIngestionDataSourceSettingsCloudStorage
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) MatchGlob() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchGlob",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) MatchGlobInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchGlobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) MinimumObjectCreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumObjectCreateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) MinimumObjectCreateTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumObjectCreateTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) PubsubAvroFormat() PubsubTopicIngestionDataSourceSettingsCloudStoragePubsubAvroFormatOutputReference {
	var returns PubsubTopicIngestionDataSourceSettingsCloudStoragePubsubAvroFormatOutputReference
	_jsii_.Get(
		j,
		"pubsubAvroFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) PubsubAvroFormatInput() *PubsubTopicIngestionDataSourceSettingsCloudStoragePubsubAvroFormat {
	var returns *PubsubTopicIngestionDataSourceSettingsCloudStoragePubsubAvroFormat
	_jsii_.Get(
		j,
		"pubsubAvroFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) TextFormat() PubsubTopicIngestionDataSourceSettingsCloudStorageTextFormatOutputReference {
	var returns PubsubTopicIngestionDataSourceSettingsCloudStorageTextFormatOutputReference
	_jsii_.Get(
		j,
		"textFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) TextFormatInput() *PubsubTopicIngestionDataSourceSettingsCloudStorageTextFormat {
	var returns *PubsubTopicIngestionDataSourceSettingsCloudStorageTextFormat
	_jsii_.Get(
		j,
		"textFormatInput",
		&returns,
	)
	return returns
}


func NewPubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference {
	_init_.Initialize()

	if err := validateNewPubsubTopicIngestionDataSourceSettingsCloudStorageOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.pubsubTopic.PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference_Override(p PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.pubsubTopic.PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference)SetBucket(val *string) {
	if err := j.validateSetBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucket",
		val,
	)
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference)SetInternalValue(val *PubsubTopicIngestionDataSourceSettingsCloudStorage) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference)SetMatchGlob(val *string) {
	if err := j.validateSetMatchGlobParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchGlob",
		val,
	)
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference)SetMinimumObjectCreateTime(val *string) {
	if err := j.validateSetMinimumObjectCreateTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumObjectCreateTime",
		val,
	)
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) PutAvroFormat(value *PubsubTopicIngestionDataSourceSettingsCloudStorageAvroFormat) {
	if err := p.validatePutAvroFormatParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAvroFormat",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) PutPubsubAvroFormat(value *PubsubTopicIngestionDataSourceSettingsCloudStoragePubsubAvroFormat) {
	if err := p.validatePutPubsubAvroFormatParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putPubsubAvroFormat",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) PutTextFormat(value *PubsubTopicIngestionDataSourceSettingsCloudStorageTextFormat) {
	if err := p.validatePutTextFormatParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putTextFormat",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) ResetAvroFormat() {
	_jsii_.InvokeVoid(
		p,
		"resetAvroFormat",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) ResetMatchGlob() {
	_jsii_.InvokeVoid(
		p,
		"resetMatchGlob",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) ResetMinimumObjectCreateTime() {
	_jsii_.InvokeVoid(
		p,
		"resetMinimumObjectCreateTime",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) ResetPubsubAvroFormat() {
	_jsii_.InvokeVoid(
		p,
		"resetPubsubAvroFormat",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) ResetTextFormat() {
	_jsii_.InvokeVoid(
		p,
		"resetTextFormat",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PubsubTopicIngestionDataSourceSettingsCloudStorageOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

