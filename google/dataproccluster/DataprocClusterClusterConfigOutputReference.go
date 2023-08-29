// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataproccluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v9/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v9/dataproccluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataprocClusterClusterConfigOutputReference interface {
	cdktf.ComplexObject
	AutoscalingConfig() DataprocClusterClusterConfigAutoscalingConfigOutputReference
	AutoscalingConfigInput() *DataprocClusterClusterConfigAutoscalingConfig
	Bucket() *string
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
	DataprocMetricConfig() DataprocClusterClusterConfigDataprocMetricConfigOutputReference
	DataprocMetricConfigInput() *DataprocClusterClusterConfigDataprocMetricConfig
	EncryptionConfig() DataprocClusterClusterConfigEncryptionConfigOutputReference
	EncryptionConfigInput() *DataprocClusterClusterConfigEncryptionConfig
	EndpointConfig() DataprocClusterClusterConfigEndpointConfigOutputReference
	EndpointConfigInput() *DataprocClusterClusterConfigEndpointConfig
	// Experimental.
	Fqn() *string
	GceClusterConfig() DataprocClusterClusterConfigGceClusterConfigOutputReference
	GceClusterConfigInput() *DataprocClusterClusterConfigGceClusterConfig
	InitializationAction() DataprocClusterClusterConfigInitializationActionList
	InitializationActionInput() interface{}
	InternalValue() *DataprocClusterClusterConfig
	SetInternalValue(val *DataprocClusterClusterConfig)
	LifecycleConfig() DataprocClusterClusterConfigLifecycleConfigOutputReference
	LifecycleConfigInput() *DataprocClusterClusterConfigLifecycleConfig
	MasterConfig() DataprocClusterClusterConfigMasterConfigOutputReference
	MasterConfigInput() *DataprocClusterClusterConfigMasterConfig
	MetastoreConfig() DataprocClusterClusterConfigMetastoreConfigOutputReference
	MetastoreConfigInput() *DataprocClusterClusterConfigMetastoreConfig
	PreemptibleWorkerConfig() DataprocClusterClusterConfigPreemptibleWorkerConfigOutputReference
	PreemptibleWorkerConfigInput() *DataprocClusterClusterConfigPreemptibleWorkerConfig
	SecurityConfig() DataprocClusterClusterConfigSecurityConfigOutputReference
	SecurityConfigInput() *DataprocClusterClusterConfigSecurityConfig
	SoftwareConfig() DataprocClusterClusterConfigSoftwareConfigOutputReference
	SoftwareConfigInput() *DataprocClusterClusterConfigSoftwareConfig
	StagingBucket() *string
	SetStagingBucket(val *string)
	StagingBucketInput() *string
	TempBucket() *string
	SetTempBucket(val *string)
	TempBucketInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WorkerConfig() DataprocClusterClusterConfigWorkerConfigOutputReference
	WorkerConfigInput() *DataprocClusterClusterConfigWorkerConfig
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
	PutAutoscalingConfig(value *DataprocClusterClusterConfigAutoscalingConfig)
	PutDataprocMetricConfig(value *DataprocClusterClusterConfigDataprocMetricConfig)
	PutEncryptionConfig(value *DataprocClusterClusterConfigEncryptionConfig)
	PutEndpointConfig(value *DataprocClusterClusterConfigEndpointConfig)
	PutGceClusterConfig(value *DataprocClusterClusterConfigGceClusterConfig)
	PutInitializationAction(value interface{})
	PutLifecycleConfig(value *DataprocClusterClusterConfigLifecycleConfig)
	PutMasterConfig(value *DataprocClusterClusterConfigMasterConfig)
	PutMetastoreConfig(value *DataprocClusterClusterConfigMetastoreConfig)
	PutPreemptibleWorkerConfig(value *DataprocClusterClusterConfigPreemptibleWorkerConfig)
	PutSecurityConfig(value *DataprocClusterClusterConfigSecurityConfig)
	PutSoftwareConfig(value *DataprocClusterClusterConfigSoftwareConfig)
	PutWorkerConfig(value *DataprocClusterClusterConfigWorkerConfig)
	ResetAutoscalingConfig()
	ResetDataprocMetricConfig()
	ResetEncryptionConfig()
	ResetEndpointConfig()
	ResetGceClusterConfig()
	ResetInitializationAction()
	ResetLifecycleConfig()
	ResetMasterConfig()
	ResetMetastoreConfig()
	ResetPreemptibleWorkerConfig()
	ResetSecurityConfig()
	ResetSoftwareConfig()
	ResetStagingBucket()
	ResetTempBucket()
	ResetWorkerConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataprocClusterClusterConfigOutputReference
type jsiiProxy_DataprocClusterClusterConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataprocClusterClusterConfigOutputReference) AutoscalingConfig() DataprocClusterClusterConfigAutoscalingConfigOutputReference {
	var returns DataprocClusterClusterConfigAutoscalingConfigOutputReference
	_jsii_.Get(
		j,
		"autoscalingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigOutputReference) AutoscalingConfigInput() *DataprocClusterClusterConfigAutoscalingConfig {
	var returns *DataprocClusterClusterConfigAutoscalingConfig
	_jsii_.Get(
		j,
		"autoscalingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigOutputReference) Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigOutputReference) DataprocMetricConfig() DataprocClusterClusterConfigDataprocMetricConfigOutputReference {
	var returns DataprocClusterClusterConfigDataprocMetricConfigOutputReference
	_jsii_.Get(
		j,
		"dataprocMetricConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigOutputReference) DataprocMetricConfigInput() *DataprocClusterClusterConfigDataprocMetricConfig {
	var returns *DataprocClusterClusterConfigDataprocMetricConfig
	_jsii_.Get(
		j,
		"dataprocMetricConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigOutputReference) EncryptionConfig() DataprocClusterClusterConfigEncryptionConfigOutputReference {
	var returns DataprocClusterClusterConfigEncryptionConfigOutputReference
	_jsii_.Get(
		j,
		"encryptionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigOutputReference) EncryptionConfigInput() *DataprocClusterClusterConfigEncryptionConfig {
	var returns *DataprocClusterClusterConfigEncryptionConfig
	_jsii_.Get(
		j,
		"encryptionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigOutputReference) EndpointConfig() DataprocClusterClusterConfigEndpointConfigOutputReference {
	var returns DataprocClusterClusterConfigEndpointConfigOutputReference
	_jsii_.Get(
		j,
		"endpointConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigOutputReference) EndpointConfigInput() *DataprocClusterClusterConfigEndpointConfig {
	var returns *DataprocClusterClusterConfigEndpointConfig
	_jsii_.Get(
		j,
		"endpointConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigOutputReference) GceClusterConfig() DataprocClusterClusterConfigGceClusterConfigOutputReference {
	var returns DataprocClusterClusterConfigGceClusterConfigOutputReference
	_jsii_.Get(
		j,
		"gceClusterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigOutputReference) GceClusterConfigInput() *DataprocClusterClusterConfigGceClusterConfig {
	var returns *DataprocClusterClusterConfigGceClusterConfig
	_jsii_.Get(
		j,
		"gceClusterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigOutputReference) InitializationAction() DataprocClusterClusterConfigInitializationActionList {
	var returns DataprocClusterClusterConfigInitializationActionList
	_jsii_.Get(
		j,
		"initializationAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigOutputReference) InitializationActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initializationActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigOutputReference) InternalValue() *DataprocClusterClusterConfig {
	var returns *DataprocClusterClusterConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigOutputReference) LifecycleConfig() DataprocClusterClusterConfigLifecycleConfigOutputReference {
	var returns DataprocClusterClusterConfigLifecycleConfigOutputReference
	_jsii_.Get(
		j,
		"lifecycleConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigOutputReference) LifecycleConfigInput() *DataprocClusterClusterConfigLifecycleConfig {
	var returns *DataprocClusterClusterConfigLifecycleConfig
	_jsii_.Get(
		j,
		"lifecycleConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigOutputReference) MasterConfig() DataprocClusterClusterConfigMasterConfigOutputReference {
	var returns DataprocClusterClusterConfigMasterConfigOutputReference
	_jsii_.Get(
		j,
		"masterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigOutputReference) MasterConfigInput() *DataprocClusterClusterConfigMasterConfig {
	var returns *DataprocClusterClusterConfigMasterConfig
	_jsii_.Get(
		j,
		"masterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigOutputReference) MetastoreConfig() DataprocClusterClusterConfigMetastoreConfigOutputReference {
	var returns DataprocClusterClusterConfigMetastoreConfigOutputReference
	_jsii_.Get(
		j,
		"metastoreConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigOutputReference) MetastoreConfigInput() *DataprocClusterClusterConfigMetastoreConfig {
	var returns *DataprocClusterClusterConfigMetastoreConfig
	_jsii_.Get(
		j,
		"metastoreConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigOutputReference) PreemptibleWorkerConfig() DataprocClusterClusterConfigPreemptibleWorkerConfigOutputReference {
	var returns DataprocClusterClusterConfigPreemptibleWorkerConfigOutputReference
	_jsii_.Get(
		j,
		"preemptibleWorkerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigOutputReference) PreemptibleWorkerConfigInput() *DataprocClusterClusterConfigPreemptibleWorkerConfig {
	var returns *DataprocClusterClusterConfigPreemptibleWorkerConfig
	_jsii_.Get(
		j,
		"preemptibleWorkerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigOutputReference) SecurityConfig() DataprocClusterClusterConfigSecurityConfigOutputReference {
	var returns DataprocClusterClusterConfigSecurityConfigOutputReference
	_jsii_.Get(
		j,
		"securityConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigOutputReference) SecurityConfigInput() *DataprocClusterClusterConfigSecurityConfig {
	var returns *DataprocClusterClusterConfigSecurityConfig
	_jsii_.Get(
		j,
		"securityConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigOutputReference) SoftwareConfig() DataprocClusterClusterConfigSoftwareConfigOutputReference {
	var returns DataprocClusterClusterConfigSoftwareConfigOutputReference
	_jsii_.Get(
		j,
		"softwareConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigOutputReference) SoftwareConfigInput() *DataprocClusterClusterConfigSoftwareConfig {
	var returns *DataprocClusterClusterConfigSoftwareConfig
	_jsii_.Get(
		j,
		"softwareConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigOutputReference) StagingBucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stagingBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigOutputReference) StagingBucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stagingBucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigOutputReference) TempBucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tempBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigOutputReference) TempBucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tempBucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigOutputReference) WorkerConfig() DataprocClusterClusterConfigWorkerConfigOutputReference {
	var returns DataprocClusterClusterConfigWorkerConfigOutputReference
	_jsii_.Get(
		j,
		"workerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocClusterClusterConfigOutputReference) WorkerConfigInput() *DataprocClusterClusterConfigWorkerConfig {
	var returns *DataprocClusterClusterConfigWorkerConfig
	_jsii_.Get(
		j,
		"workerConfigInput",
		&returns,
	)
	return returns
}


func NewDataprocClusterClusterConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataprocClusterClusterConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataprocClusterClusterConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataprocClusterClusterConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataprocCluster.DataprocClusterClusterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataprocClusterClusterConfigOutputReference_Override(d DataprocClusterClusterConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataprocCluster.DataprocClusterClusterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigOutputReference)SetInternalValue(val *DataprocClusterClusterConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigOutputReference)SetStagingBucket(val *string) {
	if err := j.validateSetStagingBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stagingBucket",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigOutputReference)SetTempBucket(val *string) {
	if err := j.validateSetTempBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tempBucket",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataprocClusterClusterConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataprocClusterClusterConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataprocClusterClusterConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataprocClusterClusterConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataprocClusterClusterConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataprocClusterClusterConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataprocClusterClusterConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataprocClusterClusterConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataprocClusterClusterConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataprocClusterClusterConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocClusterClusterConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataprocClusterClusterConfigOutputReference) PutAutoscalingConfig(value *DataprocClusterClusterConfigAutoscalingConfig) {
	if err := d.validatePutAutoscalingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAutoscalingConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigOutputReference) PutDataprocMetricConfig(value *DataprocClusterClusterConfigDataprocMetricConfig) {
	if err := d.validatePutDataprocMetricConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDataprocMetricConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigOutputReference) PutEncryptionConfig(value *DataprocClusterClusterConfigEncryptionConfig) {
	if err := d.validatePutEncryptionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEncryptionConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigOutputReference) PutEndpointConfig(value *DataprocClusterClusterConfigEndpointConfig) {
	if err := d.validatePutEndpointConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEndpointConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigOutputReference) PutGceClusterConfig(value *DataprocClusterClusterConfigGceClusterConfig) {
	if err := d.validatePutGceClusterConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGceClusterConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigOutputReference) PutInitializationAction(value interface{}) {
	if err := d.validatePutInitializationActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putInitializationAction",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigOutputReference) PutLifecycleConfig(value *DataprocClusterClusterConfigLifecycleConfig) {
	if err := d.validatePutLifecycleConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLifecycleConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigOutputReference) PutMasterConfig(value *DataprocClusterClusterConfigMasterConfig) {
	if err := d.validatePutMasterConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMasterConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigOutputReference) PutMetastoreConfig(value *DataprocClusterClusterConfigMetastoreConfig) {
	if err := d.validatePutMetastoreConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMetastoreConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigOutputReference) PutPreemptibleWorkerConfig(value *DataprocClusterClusterConfigPreemptibleWorkerConfig) {
	if err := d.validatePutPreemptibleWorkerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPreemptibleWorkerConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigOutputReference) PutSecurityConfig(value *DataprocClusterClusterConfigSecurityConfig) {
	if err := d.validatePutSecurityConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSecurityConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigOutputReference) PutSoftwareConfig(value *DataprocClusterClusterConfigSoftwareConfig) {
	if err := d.validatePutSoftwareConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSoftwareConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigOutputReference) PutWorkerConfig(value *DataprocClusterClusterConfigWorkerConfig) {
	if err := d.validatePutWorkerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWorkerConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigOutputReference) ResetAutoscalingConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetAutoscalingConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigOutputReference) ResetDataprocMetricConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetDataprocMetricConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigOutputReference) ResetEncryptionConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetEncryptionConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigOutputReference) ResetEndpointConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetEndpointConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigOutputReference) ResetGceClusterConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetGceClusterConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigOutputReference) ResetInitializationAction() {
	_jsii_.InvokeVoid(
		d,
		"resetInitializationAction",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigOutputReference) ResetLifecycleConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetLifecycleConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigOutputReference) ResetMasterConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetMasterConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigOutputReference) ResetMetastoreConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetMetastoreConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigOutputReference) ResetPreemptibleWorkerConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetPreemptibleWorkerConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigOutputReference) ResetSecurityConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetSecurityConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigOutputReference) ResetSoftwareConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetSoftwareConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigOutputReference) ResetStagingBucket() {
	_jsii_.InvokeVoid(
		d,
		"resetStagingBucket",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigOutputReference) ResetTempBucket() {
	_jsii_.InvokeVoid(
		d,
		"resetTempBucket",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigOutputReference) ResetWorkerConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetWorkerConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocClusterClusterConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataprocClusterClusterConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

