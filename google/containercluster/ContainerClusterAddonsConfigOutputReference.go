// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v13/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v13/containercluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerClusterAddonsConfigOutputReference interface {
	cdktf.ComplexObject
	CloudrunConfig() ContainerClusterAddonsConfigCloudrunConfigOutputReference
	CloudrunConfigInput() *ContainerClusterAddonsConfigCloudrunConfig
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
	ConfigConnectorConfig() ContainerClusterAddonsConfigConfigConnectorConfigOutputReference
	ConfigConnectorConfigInput() *ContainerClusterAddonsConfigConfigConnectorConfig
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DnsCacheConfig() ContainerClusterAddonsConfigDnsCacheConfigOutputReference
	DnsCacheConfigInput() *ContainerClusterAddonsConfigDnsCacheConfig
	// Experimental.
	Fqn() *string
	GcePersistentDiskCsiDriverConfig() ContainerClusterAddonsConfigGcePersistentDiskCsiDriverConfigOutputReference
	GcePersistentDiskCsiDriverConfigInput() *ContainerClusterAddonsConfigGcePersistentDiskCsiDriverConfig
	GcpFilestoreCsiDriverConfig() ContainerClusterAddonsConfigGcpFilestoreCsiDriverConfigOutputReference
	GcpFilestoreCsiDriverConfigInput() *ContainerClusterAddonsConfigGcpFilestoreCsiDriverConfig
	GcsFuseCsiDriverConfig() ContainerClusterAddonsConfigGcsFuseCsiDriverConfigOutputReference
	GcsFuseCsiDriverConfigInput() *ContainerClusterAddonsConfigGcsFuseCsiDriverConfig
	GkeBackupAgentConfig() ContainerClusterAddonsConfigGkeBackupAgentConfigOutputReference
	GkeBackupAgentConfigInput() *ContainerClusterAddonsConfigGkeBackupAgentConfig
	HorizontalPodAutoscaling() ContainerClusterAddonsConfigHorizontalPodAutoscalingOutputReference
	HorizontalPodAutoscalingInput() *ContainerClusterAddonsConfigHorizontalPodAutoscaling
	HttpLoadBalancing() ContainerClusterAddonsConfigHttpLoadBalancingOutputReference
	HttpLoadBalancingInput() *ContainerClusterAddonsConfigHttpLoadBalancing
	InternalValue() *ContainerClusterAddonsConfig
	SetInternalValue(val *ContainerClusterAddonsConfig)
	NetworkPolicyConfig() ContainerClusterAddonsConfigNetworkPolicyConfigOutputReference
	NetworkPolicyConfigInput() *ContainerClusterAddonsConfigNetworkPolicyConfig
	RayOperatorConfig() ContainerClusterAddonsConfigRayOperatorConfigList
	RayOperatorConfigInput() interface{}
	StatefulHaConfig() ContainerClusterAddonsConfigStatefulHaConfigOutputReference
	StatefulHaConfigInput() *ContainerClusterAddonsConfigStatefulHaConfig
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
	PutCloudrunConfig(value *ContainerClusterAddonsConfigCloudrunConfig)
	PutConfigConnectorConfig(value *ContainerClusterAddonsConfigConfigConnectorConfig)
	PutDnsCacheConfig(value *ContainerClusterAddonsConfigDnsCacheConfig)
	PutGcePersistentDiskCsiDriverConfig(value *ContainerClusterAddonsConfigGcePersistentDiskCsiDriverConfig)
	PutGcpFilestoreCsiDriverConfig(value *ContainerClusterAddonsConfigGcpFilestoreCsiDriverConfig)
	PutGcsFuseCsiDriverConfig(value *ContainerClusterAddonsConfigGcsFuseCsiDriverConfig)
	PutGkeBackupAgentConfig(value *ContainerClusterAddonsConfigGkeBackupAgentConfig)
	PutHorizontalPodAutoscaling(value *ContainerClusterAddonsConfigHorizontalPodAutoscaling)
	PutHttpLoadBalancing(value *ContainerClusterAddonsConfigHttpLoadBalancing)
	PutNetworkPolicyConfig(value *ContainerClusterAddonsConfigNetworkPolicyConfig)
	PutRayOperatorConfig(value interface{})
	PutStatefulHaConfig(value *ContainerClusterAddonsConfigStatefulHaConfig)
	ResetCloudrunConfig()
	ResetConfigConnectorConfig()
	ResetDnsCacheConfig()
	ResetGcePersistentDiskCsiDriverConfig()
	ResetGcpFilestoreCsiDriverConfig()
	ResetGcsFuseCsiDriverConfig()
	ResetGkeBackupAgentConfig()
	ResetHorizontalPodAutoscaling()
	ResetHttpLoadBalancing()
	ResetNetworkPolicyConfig()
	ResetRayOperatorConfig()
	ResetStatefulHaConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerClusterAddonsConfigOutputReference
type jsiiProxy_ContainerClusterAddonsConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerClusterAddonsConfigOutputReference) CloudrunConfig() ContainerClusterAddonsConfigCloudrunConfigOutputReference {
	var returns ContainerClusterAddonsConfigCloudrunConfigOutputReference
	_jsii_.Get(
		j,
		"cloudrunConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterAddonsConfigOutputReference) CloudrunConfigInput() *ContainerClusterAddonsConfigCloudrunConfig {
	var returns *ContainerClusterAddonsConfigCloudrunConfig
	_jsii_.Get(
		j,
		"cloudrunConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterAddonsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterAddonsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterAddonsConfigOutputReference) ConfigConnectorConfig() ContainerClusterAddonsConfigConfigConnectorConfigOutputReference {
	var returns ContainerClusterAddonsConfigConfigConnectorConfigOutputReference
	_jsii_.Get(
		j,
		"configConnectorConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterAddonsConfigOutputReference) ConfigConnectorConfigInput() *ContainerClusterAddonsConfigConfigConnectorConfig {
	var returns *ContainerClusterAddonsConfigConfigConnectorConfig
	_jsii_.Get(
		j,
		"configConnectorConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterAddonsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterAddonsConfigOutputReference) DnsCacheConfig() ContainerClusterAddonsConfigDnsCacheConfigOutputReference {
	var returns ContainerClusterAddonsConfigDnsCacheConfigOutputReference
	_jsii_.Get(
		j,
		"dnsCacheConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterAddonsConfigOutputReference) DnsCacheConfigInput() *ContainerClusterAddonsConfigDnsCacheConfig {
	var returns *ContainerClusterAddonsConfigDnsCacheConfig
	_jsii_.Get(
		j,
		"dnsCacheConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterAddonsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterAddonsConfigOutputReference) GcePersistentDiskCsiDriverConfig() ContainerClusterAddonsConfigGcePersistentDiskCsiDriverConfigOutputReference {
	var returns ContainerClusterAddonsConfigGcePersistentDiskCsiDriverConfigOutputReference
	_jsii_.Get(
		j,
		"gcePersistentDiskCsiDriverConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterAddonsConfigOutputReference) GcePersistentDiskCsiDriverConfigInput() *ContainerClusterAddonsConfigGcePersistentDiskCsiDriverConfig {
	var returns *ContainerClusterAddonsConfigGcePersistentDiskCsiDriverConfig
	_jsii_.Get(
		j,
		"gcePersistentDiskCsiDriverConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterAddonsConfigOutputReference) GcpFilestoreCsiDriverConfig() ContainerClusterAddonsConfigGcpFilestoreCsiDriverConfigOutputReference {
	var returns ContainerClusterAddonsConfigGcpFilestoreCsiDriverConfigOutputReference
	_jsii_.Get(
		j,
		"gcpFilestoreCsiDriverConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterAddonsConfigOutputReference) GcpFilestoreCsiDriverConfigInput() *ContainerClusterAddonsConfigGcpFilestoreCsiDriverConfig {
	var returns *ContainerClusterAddonsConfigGcpFilestoreCsiDriverConfig
	_jsii_.Get(
		j,
		"gcpFilestoreCsiDriverConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterAddonsConfigOutputReference) GcsFuseCsiDriverConfig() ContainerClusterAddonsConfigGcsFuseCsiDriverConfigOutputReference {
	var returns ContainerClusterAddonsConfigGcsFuseCsiDriverConfigOutputReference
	_jsii_.Get(
		j,
		"gcsFuseCsiDriverConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterAddonsConfigOutputReference) GcsFuseCsiDriverConfigInput() *ContainerClusterAddonsConfigGcsFuseCsiDriverConfig {
	var returns *ContainerClusterAddonsConfigGcsFuseCsiDriverConfig
	_jsii_.Get(
		j,
		"gcsFuseCsiDriverConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterAddonsConfigOutputReference) GkeBackupAgentConfig() ContainerClusterAddonsConfigGkeBackupAgentConfigOutputReference {
	var returns ContainerClusterAddonsConfigGkeBackupAgentConfigOutputReference
	_jsii_.Get(
		j,
		"gkeBackupAgentConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterAddonsConfigOutputReference) GkeBackupAgentConfigInput() *ContainerClusterAddonsConfigGkeBackupAgentConfig {
	var returns *ContainerClusterAddonsConfigGkeBackupAgentConfig
	_jsii_.Get(
		j,
		"gkeBackupAgentConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterAddonsConfigOutputReference) HorizontalPodAutoscaling() ContainerClusterAddonsConfigHorizontalPodAutoscalingOutputReference {
	var returns ContainerClusterAddonsConfigHorizontalPodAutoscalingOutputReference
	_jsii_.Get(
		j,
		"horizontalPodAutoscaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterAddonsConfigOutputReference) HorizontalPodAutoscalingInput() *ContainerClusterAddonsConfigHorizontalPodAutoscaling {
	var returns *ContainerClusterAddonsConfigHorizontalPodAutoscaling
	_jsii_.Get(
		j,
		"horizontalPodAutoscalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterAddonsConfigOutputReference) HttpLoadBalancing() ContainerClusterAddonsConfigHttpLoadBalancingOutputReference {
	var returns ContainerClusterAddonsConfigHttpLoadBalancingOutputReference
	_jsii_.Get(
		j,
		"httpLoadBalancing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterAddonsConfigOutputReference) HttpLoadBalancingInput() *ContainerClusterAddonsConfigHttpLoadBalancing {
	var returns *ContainerClusterAddonsConfigHttpLoadBalancing
	_jsii_.Get(
		j,
		"httpLoadBalancingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterAddonsConfigOutputReference) InternalValue() *ContainerClusterAddonsConfig {
	var returns *ContainerClusterAddonsConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterAddonsConfigOutputReference) NetworkPolicyConfig() ContainerClusterAddonsConfigNetworkPolicyConfigOutputReference {
	var returns ContainerClusterAddonsConfigNetworkPolicyConfigOutputReference
	_jsii_.Get(
		j,
		"networkPolicyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterAddonsConfigOutputReference) NetworkPolicyConfigInput() *ContainerClusterAddonsConfigNetworkPolicyConfig {
	var returns *ContainerClusterAddonsConfigNetworkPolicyConfig
	_jsii_.Get(
		j,
		"networkPolicyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterAddonsConfigOutputReference) RayOperatorConfig() ContainerClusterAddonsConfigRayOperatorConfigList {
	var returns ContainerClusterAddonsConfigRayOperatorConfigList
	_jsii_.Get(
		j,
		"rayOperatorConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterAddonsConfigOutputReference) RayOperatorConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rayOperatorConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterAddonsConfigOutputReference) StatefulHaConfig() ContainerClusterAddonsConfigStatefulHaConfigOutputReference {
	var returns ContainerClusterAddonsConfigStatefulHaConfigOutputReference
	_jsii_.Get(
		j,
		"statefulHaConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterAddonsConfigOutputReference) StatefulHaConfigInput() *ContainerClusterAddonsConfigStatefulHaConfig {
	var returns *ContainerClusterAddonsConfigStatefulHaConfig
	_jsii_.Get(
		j,
		"statefulHaConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterAddonsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterAddonsConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewContainerClusterAddonsConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContainerClusterAddonsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewContainerClusterAddonsConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerClusterAddonsConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.containerCluster.ContainerClusterAddonsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerClusterAddonsConfigOutputReference_Override(c ContainerClusterAddonsConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.containerCluster.ContainerClusterAddonsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerClusterAddonsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterAddonsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterAddonsConfigOutputReference)SetInternalValue(val *ContainerClusterAddonsConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterAddonsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterAddonsConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ContainerClusterAddonsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterAddonsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterAddonsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterAddonsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterAddonsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterAddonsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterAddonsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterAddonsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterAddonsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterAddonsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterAddonsConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterAddonsConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterAddonsConfigOutputReference) PutCloudrunConfig(value *ContainerClusterAddonsConfigCloudrunConfig) {
	if err := c.validatePutCloudrunConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCloudrunConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterAddonsConfigOutputReference) PutConfigConnectorConfig(value *ContainerClusterAddonsConfigConfigConnectorConfig) {
	if err := c.validatePutConfigConnectorConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putConfigConnectorConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterAddonsConfigOutputReference) PutDnsCacheConfig(value *ContainerClusterAddonsConfigDnsCacheConfig) {
	if err := c.validatePutDnsCacheConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDnsCacheConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterAddonsConfigOutputReference) PutGcePersistentDiskCsiDriverConfig(value *ContainerClusterAddonsConfigGcePersistentDiskCsiDriverConfig) {
	if err := c.validatePutGcePersistentDiskCsiDriverConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGcePersistentDiskCsiDriverConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterAddonsConfigOutputReference) PutGcpFilestoreCsiDriverConfig(value *ContainerClusterAddonsConfigGcpFilestoreCsiDriverConfig) {
	if err := c.validatePutGcpFilestoreCsiDriverConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGcpFilestoreCsiDriverConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterAddonsConfigOutputReference) PutGcsFuseCsiDriverConfig(value *ContainerClusterAddonsConfigGcsFuseCsiDriverConfig) {
	if err := c.validatePutGcsFuseCsiDriverConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGcsFuseCsiDriverConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterAddonsConfigOutputReference) PutGkeBackupAgentConfig(value *ContainerClusterAddonsConfigGkeBackupAgentConfig) {
	if err := c.validatePutGkeBackupAgentConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGkeBackupAgentConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterAddonsConfigOutputReference) PutHorizontalPodAutoscaling(value *ContainerClusterAddonsConfigHorizontalPodAutoscaling) {
	if err := c.validatePutHorizontalPodAutoscalingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHorizontalPodAutoscaling",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterAddonsConfigOutputReference) PutHttpLoadBalancing(value *ContainerClusterAddonsConfigHttpLoadBalancing) {
	if err := c.validatePutHttpLoadBalancingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHttpLoadBalancing",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterAddonsConfigOutputReference) PutNetworkPolicyConfig(value *ContainerClusterAddonsConfigNetworkPolicyConfig) {
	if err := c.validatePutNetworkPolicyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNetworkPolicyConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterAddonsConfigOutputReference) PutRayOperatorConfig(value interface{}) {
	if err := c.validatePutRayOperatorConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRayOperatorConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterAddonsConfigOutputReference) PutStatefulHaConfig(value *ContainerClusterAddonsConfigStatefulHaConfig) {
	if err := c.validatePutStatefulHaConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putStatefulHaConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterAddonsConfigOutputReference) ResetCloudrunConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetCloudrunConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterAddonsConfigOutputReference) ResetConfigConnectorConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetConfigConnectorConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterAddonsConfigOutputReference) ResetDnsCacheConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetDnsCacheConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterAddonsConfigOutputReference) ResetGcePersistentDiskCsiDriverConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetGcePersistentDiskCsiDriverConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterAddonsConfigOutputReference) ResetGcpFilestoreCsiDriverConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetGcpFilestoreCsiDriverConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterAddonsConfigOutputReference) ResetGcsFuseCsiDriverConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetGcsFuseCsiDriverConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterAddonsConfigOutputReference) ResetGkeBackupAgentConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetGkeBackupAgentConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterAddonsConfigOutputReference) ResetHorizontalPodAutoscaling() {
	_jsii_.InvokeVoid(
		c,
		"resetHorizontalPodAutoscaling",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterAddonsConfigOutputReference) ResetHttpLoadBalancing() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpLoadBalancing",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterAddonsConfigOutputReference) ResetNetworkPolicyConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetNetworkPolicyConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterAddonsConfigOutputReference) ResetRayOperatorConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetRayOperatorConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterAddonsConfigOutputReference) ResetStatefulHaConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetStatefulHaConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterAddonsConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterAddonsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

