// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocsessiontemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/dataprocsessiontemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference interface {
	cdktf.ComplexObject
	AuthenticationConfig() DataprocSessionTemplateEnvironmentConfigExecutionConfigAuthenticationConfigOutputReference
	AuthenticationConfigInput() *DataprocSessionTemplateEnvironmentConfigExecutionConfigAuthenticationConfig
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
	IdleTtl() *string
	SetIdleTtl(val *string)
	IdleTtlInput() *string
	InternalValue() *DataprocSessionTemplateEnvironmentConfigExecutionConfig
	SetInternalValue(val *DataprocSessionTemplateEnvironmentConfigExecutionConfig)
	KmsKey() *string
	SetKmsKey(val *string)
	KmsKeyInput() *string
	NetworkTags() *[]*string
	SetNetworkTags(val *[]*string)
	NetworkTagsInput() *[]*string
	ServiceAccount() *string
	SetServiceAccount(val *string)
	ServiceAccountInput() *string
	StagingBucket() *string
	SetStagingBucket(val *string)
	StagingBucketInput() *string
	SubnetworkUri() *string
	SetSubnetworkUri(val *string)
	SubnetworkUriInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Ttl() *string
	SetTtl(val *string)
	TtlInput() *string
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
	PutAuthenticationConfig(value *DataprocSessionTemplateEnvironmentConfigExecutionConfigAuthenticationConfig)
	ResetAuthenticationConfig()
	ResetIdleTtl()
	ResetKmsKey()
	ResetNetworkTags()
	ResetServiceAccount()
	ResetStagingBucket()
	ResetSubnetworkUri()
	ResetTtl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference
type jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) AuthenticationConfig() DataprocSessionTemplateEnvironmentConfigExecutionConfigAuthenticationConfigOutputReference {
	var returns DataprocSessionTemplateEnvironmentConfigExecutionConfigAuthenticationConfigOutputReference
	_jsii_.Get(
		j,
		"authenticationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) AuthenticationConfigInput() *DataprocSessionTemplateEnvironmentConfigExecutionConfigAuthenticationConfig {
	var returns *DataprocSessionTemplateEnvironmentConfigExecutionConfigAuthenticationConfig
	_jsii_.Get(
		j,
		"authenticationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) IdleTtl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idleTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) IdleTtlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idleTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) InternalValue() *DataprocSessionTemplateEnvironmentConfigExecutionConfig {
	var returns *DataprocSessionTemplateEnvironmentConfigExecutionConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) KmsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) KmsKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) NetworkTags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) NetworkTagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) ServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) StagingBucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stagingBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) StagingBucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stagingBucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) SubnetworkUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetworkUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) SubnetworkUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetworkUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) Ttl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) TtlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ttlInput",
		&returns,
	)
	return returns
}


func NewDataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataprocSessionTemplate.DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference_Override(d DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataprocSessionTemplate.DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference)SetIdleTtl(val *string) {
	if err := j.validateSetIdleTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idleTtl",
		val,
	)
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference)SetInternalValue(val *DataprocSessionTemplateEnvironmentConfigExecutionConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference)SetKmsKey(val *string) {
	if err := j.validateSetKmsKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKey",
		val,
	)
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference)SetNetworkTags(val *[]*string) {
	if err := j.validateSetNetworkTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkTags",
		val,
	)
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference)SetServiceAccount(val *string) {
	if err := j.validateSetServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccount",
		val,
	)
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference)SetStagingBucket(val *string) {
	if err := j.validateSetStagingBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stagingBucket",
		val,
	)
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference)SetSubnetworkUri(val *string) {
	if err := j.validateSetSubnetworkUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetworkUri",
		val,
	)
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference)SetTtl(val *string) {
	if err := j.validateSetTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ttl",
		val,
	)
}

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) PutAuthenticationConfig(value *DataprocSessionTemplateEnvironmentConfigExecutionConfigAuthenticationConfig) {
	if err := d.validatePutAuthenticationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAuthenticationConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) ResetAuthenticationConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetAuthenticationConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) ResetIdleTtl() {
	_jsii_.InvokeVoid(
		d,
		"resetIdleTtl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) ResetKmsKey() {
	_jsii_.InvokeVoid(
		d,
		"resetKmsKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) ResetNetworkTags() {
	_jsii_.InvokeVoid(
		d,
		"resetNetworkTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) ResetStagingBucket() {
	_jsii_.InvokeVoid(
		d,
		"resetStagingBucket",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) ResetSubnetworkUri() {
	_jsii_.InvokeVoid(
		d,
		"resetSubnetworkUri",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) ResetTtl() {
	_jsii_.InvokeVoid(
		d,
		"resetTtl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigExecutionConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

