package dataprocworkflowtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v3/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v3/dataprocworkflowtemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference interface {
	cdktf.ComplexObject
	AutoscalingConfig() DataprocWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfigOutputReference
	AutoscalingConfigInput() *DataprocWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig
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
	EncryptionConfig() DataprocWorkflowTemplatePlacementManagedClusterConfigEncryptionConfigOutputReference
	EncryptionConfigInput() *DataprocWorkflowTemplatePlacementManagedClusterConfigEncryptionConfig
	EndpointConfig() DataprocWorkflowTemplatePlacementManagedClusterConfigEndpointConfigOutputReference
	EndpointConfigInput() *DataprocWorkflowTemplatePlacementManagedClusterConfigEndpointConfig
	// Experimental.
	Fqn() *string
	GceClusterConfig() DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigOutputReference
	GceClusterConfigInput() *DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfig
	InitializationActions() DataprocWorkflowTemplatePlacementManagedClusterConfigInitializationActionsList
	InitializationActionsInput() interface{}
	InternalValue() *DataprocWorkflowTemplatePlacementManagedClusterConfig
	SetInternalValue(val *DataprocWorkflowTemplatePlacementManagedClusterConfig)
	LifecycleConfig() DataprocWorkflowTemplatePlacementManagedClusterConfigLifecycleConfigOutputReference
	LifecycleConfigInput() *DataprocWorkflowTemplatePlacementManagedClusterConfigLifecycleConfig
	MasterConfig() DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigOutputReference
	MasterConfigInput() *DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfig
	SecondaryWorkerConfig() DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference
	SecondaryWorkerConfigInput() *DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig
	SecurityConfig() DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigOutputReference
	SecurityConfigInput() *DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfig
	SoftwareConfig() DataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOutputReference
	SoftwareConfigInput() *DataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfig
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
	WorkerConfig() DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference
	WorkerConfigInput() *DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfig
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
	PutAutoscalingConfig(value *DataprocWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig)
	PutEncryptionConfig(value *DataprocWorkflowTemplatePlacementManagedClusterConfigEncryptionConfig)
	PutEndpointConfig(value *DataprocWorkflowTemplatePlacementManagedClusterConfigEndpointConfig)
	PutGceClusterConfig(value *DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfig)
	PutInitializationActions(value interface{})
	PutLifecycleConfig(value *DataprocWorkflowTemplatePlacementManagedClusterConfigLifecycleConfig)
	PutMasterConfig(value *DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfig)
	PutSecondaryWorkerConfig(value *DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig)
	PutSecurityConfig(value *DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfig)
	PutSoftwareConfig(value *DataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfig)
	PutWorkerConfig(value *DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfig)
	ResetAutoscalingConfig()
	ResetEncryptionConfig()
	ResetEndpointConfig()
	ResetGceClusterConfig()
	ResetInitializationActions()
	ResetLifecycleConfig()
	ResetMasterConfig()
	ResetSecondaryWorkerConfig()
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

// The jsii proxy struct for DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference
type jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) AutoscalingConfig() DataprocWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfigOutputReference {
	var returns DataprocWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfigOutputReference
	_jsii_.Get(
		j,
		"autoscalingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) AutoscalingConfigInput() *DataprocWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig {
	var returns *DataprocWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig
	_jsii_.Get(
		j,
		"autoscalingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) EncryptionConfig() DataprocWorkflowTemplatePlacementManagedClusterConfigEncryptionConfigOutputReference {
	var returns DataprocWorkflowTemplatePlacementManagedClusterConfigEncryptionConfigOutputReference
	_jsii_.Get(
		j,
		"encryptionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) EncryptionConfigInput() *DataprocWorkflowTemplatePlacementManagedClusterConfigEncryptionConfig {
	var returns *DataprocWorkflowTemplatePlacementManagedClusterConfigEncryptionConfig
	_jsii_.Get(
		j,
		"encryptionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) EndpointConfig() DataprocWorkflowTemplatePlacementManagedClusterConfigEndpointConfigOutputReference {
	var returns DataprocWorkflowTemplatePlacementManagedClusterConfigEndpointConfigOutputReference
	_jsii_.Get(
		j,
		"endpointConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) EndpointConfigInput() *DataprocWorkflowTemplatePlacementManagedClusterConfigEndpointConfig {
	var returns *DataprocWorkflowTemplatePlacementManagedClusterConfigEndpointConfig
	_jsii_.Get(
		j,
		"endpointConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) GceClusterConfig() DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigOutputReference {
	var returns DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigOutputReference
	_jsii_.Get(
		j,
		"gceClusterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) GceClusterConfigInput() *DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfig {
	var returns *DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfig
	_jsii_.Get(
		j,
		"gceClusterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) InitializationActions() DataprocWorkflowTemplatePlacementManagedClusterConfigInitializationActionsList {
	var returns DataprocWorkflowTemplatePlacementManagedClusterConfigInitializationActionsList
	_jsii_.Get(
		j,
		"initializationActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) InitializationActionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initializationActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) InternalValue() *DataprocWorkflowTemplatePlacementManagedClusterConfig {
	var returns *DataprocWorkflowTemplatePlacementManagedClusterConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) LifecycleConfig() DataprocWorkflowTemplatePlacementManagedClusterConfigLifecycleConfigOutputReference {
	var returns DataprocWorkflowTemplatePlacementManagedClusterConfigLifecycleConfigOutputReference
	_jsii_.Get(
		j,
		"lifecycleConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) LifecycleConfigInput() *DataprocWorkflowTemplatePlacementManagedClusterConfigLifecycleConfig {
	var returns *DataprocWorkflowTemplatePlacementManagedClusterConfigLifecycleConfig
	_jsii_.Get(
		j,
		"lifecycleConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) MasterConfig() DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigOutputReference {
	var returns DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigOutputReference
	_jsii_.Get(
		j,
		"masterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) MasterConfigInput() *DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfig {
	var returns *DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfig
	_jsii_.Get(
		j,
		"masterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) SecondaryWorkerConfig() DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference {
	var returns DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference
	_jsii_.Get(
		j,
		"secondaryWorkerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) SecondaryWorkerConfigInput() *DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig {
	var returns *DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig
	_jsii_.Get(
		j,
		"secondaryWorkerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) SecurityConfig() DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigOutputReference {
	var returns DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigOutputReference
	_jsii_.Get(
		j,
		"securityConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) SecurityConfigInput() *DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfig {
	var returns *DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfig
	_jsii_.Get(
		j,
		"securityConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) SoftwareConfig() DataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOutputReference {
	var returns DataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOutputReference
	_jsii_.Get(
		j,
		"softwareConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) SoftwareConfigInput() *DataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfig {
	var returns *DataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfig
	_jsii_.Get(
		j,
		"softwareConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) StagingBucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stagingBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) StagingBucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stagingBucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) TempBucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tempBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) TempBucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tempBucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) WorkerConfig() DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference {
	var returns DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference
	_jsii_.Get(
		j,
		"workerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) WorkerConfigInput() *DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfig {
	var returns *DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfig
	_jsii_.Get(
		j,
		"workerConfigInput",
		&returns,
	)
	return returns
}


func NewDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataprocWorkflowTemplate.DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference_Override(d DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataprocWorkflowTemplate.DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference)SetInternalValue(val *DataprocWorkflowTemplatePlacementManagedClusterConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference)SetStagingBucket(val *string) {
	if err := j.validateSetStagingBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stagingBucket",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference)SetTempBucket(val *string) {
	if err := j.validateSetTempBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tempBucket",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) PutAutoscalingConfig(value *DataprocWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig) {
	if err := d.validatePutAutoscalingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAutoscalingConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) PutEncryptionConfig(value *DataprocWorkflowTemplatePlacementManagedClusterConfigEncryptionConfig) {
	if err := d.validatePutEncryptionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEncryptionConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) PutEndpointConfig(value *DataprocWorkflowTemplatePlacementManagedClusterConfigEndpointConfig) {
	if err := d.validatePutEndpointConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEndpointConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) PutGceClusterConfig(value *DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfig) {
	if err := d.validatePutGceClusterConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGceClusterConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) PutInitializationActions(value interface{}) {
	if err := d.validatePutInitializationActionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putInitializationActions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) PutLifecycleConfig(value *DataprocWorkflowTemplatePlacementManagedClusterConfigLifecycleConfig) {
	if err := d.validatePutLifecycleConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLifecycleConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) PutMasterConfig(value *DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfig) {
	if err := d.validatePutMasterConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMasterConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) PutSecondaryWorkerConfig(value *DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig) {
	if err := d.validatePutSecondaryWorkerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSecondaryWorkerConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) PutSecurityConfig(value *DataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfig) {
	if err := d.validatePutSecurityConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSecurityConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) PutSoftwareConfig(value *DataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfig) {
	if err := d.validatePutSoftwareConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSoftwareConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) PutWorkerConfig(value *DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfig) {
	if err := d.validatePutWorkerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWorkerConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) ResetAutoscalingConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetAutoscalingConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) ResetEncryptionConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetEncryptionConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) ResetEndpointConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetEndpointConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) ResetGceClusterConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetGceClusterConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) ResetInitializationActions() {
	_jsii_.InvokeVoid(
		d,
		"resetInitializationActions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) ResetLifecycleConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetLifecycleConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) ResetMasterConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetMasterConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) ResetSecondaryWorkerConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetSecondaryWorkerConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) ResetSecurityConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetSecurityConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) ResetSoftwareConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetSoftwareConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) ResetStagingBucket() {
	_jsii_.InvokeVoid(
		d,
		"resetStagingBucket",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) ResetTempBucket() {
	_jsii_.InvokeVoid(
		d,
		"resetTempBucket",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) ResetWorkerConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetWorkerConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

