// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package composerenvironment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v9/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v9/composerenvironment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComposerEnvironmentConfigAOutputReference interface {
	cdktf.ComplexObject
	AirflowUri() *string
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
	DagGcsPrefix() *string
	DatabaseConfig() ComposerEnvironmentConfigDatabaseConfigOutputReference
	DatabaseConfigInput() *ComposerEnvironmentConfigDatabaseConfig
	EncryptionConfig() ComposerEnvironmentConfigEncryptionConfigOutputReference
	EncryptionConfigInput() *ComposerEnvironmentConfigEncryptionConfig
	EnvironmentSize() *string
	SetEnvironmentSize(val *string)
	EnvironmentSizeInput() *string
	// Experimental.
	Fqn() *string
	GkeCluster() *string
	InternalValue() *ComposerEnvironmentConfigA
	SetInternalValue(val *ComposerEnvironmentConfigA)
	MaintenanceWindow() ComposerEnvironmentConfigMaintenanceWindowOutputReference
	MaintenanceWindowInput() *ComposerEnvironmentConfigMaintenanceWindow
	MasterAuthorizedNetworksConfig() ComposerEnvironmentConfigMasterAuthorizedNetworksConfigOutputReference
	MasterAuthorizedNetworksConfigInput() *ComposerEnvironmentConfigMasterAuthorizedNetworksConfig
	NodeConfig() ComposerEnvironmentConfigNodeConfigOutputReference
	NodeConfigInput() *ComposerEnvironmentConfigNodeConfig
	NodeCount() *float64
	SetNodeCount(val *float64)
	NodeCountInput() *float64
	PrivateEnvironmentConfig() ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference
	PrivateEnvironmentConfigInput() *ComposerEnvironmentConfigPrivateEnvironmentConfig
	RecoveryConfig() ComposerEnvironmentConfigRecoveryConfigOutputReference
	RecoveryConfigInput() *ComposerEnvironmentConfigRecoveryConfig
	ResilienceMode() *string
	SetResilienceMode(val *string)
	ResilienceModeInput() *string
	SoftwareConfig() ComposerEnvironmentConfigSoftwareConfigOutputReference
	SoftwareConfigInput() *ComposerEnvironmentConfigSoftwareConfig
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WebServerConfig() ComposerEnvironmentConfigWebServerConfigOutputReference
	WebServerConfigInput() *ComposerEnvironmentConfigWebServerConfig
	WebServerNetworkAccessControl() ComposerEnvironmentConfigWebServerNetworkAccessControlOutputReference
	WebServerNetworkAccessControlInput() *ComposerEnvironmentConfigWebServerNetworkAccessControl
	WorkloadsConfig() ComposerEnvironmentConfigWorkloadsConfigOutputReference
	WorkloadsConfigInput() *ComposerEnvironmentConfigWorkloadsConfig
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
	PutDatabaseConfig(value *ComposerEnvironmentConfigDatabaseConfig)
	PutEncryptionConfig(value *ComposerEnvironmentConfigEncryptionConfig)
	PutMaintenanceWindow(value *ComposerEnvironmentConfigMaintenanceWindow)
	PutMasterAuthorizedNetworksConfig(value *ComposerEnvironmentConfigMasterAuthorizedNetworksConfig)
	PutNodeConfig(value *ComposerEnvironmentConfigNodeConfig)
	PutPrivateEnvironmentConfig(value *ComposerEnvironmentConfigPrivateEnvironmentConfig)
	PutRecoveryConfig(value *ComposerEnvironmentConfigRecoveryConfig)
	PutSoftwareConfig(value *ComposerEnvironmentConfigSoftwareConfig)
	PutWebServerConfig(value *ComposerEnvironmentConfigWebServerConfig)
	PutWebServerNetworkAccessControl(value *ComposerEnvironmentConfigWebServerNetworkAccessControl)
	PutWorkloadsConfig(value *ComposerEnvironmentConfigWorkloadsConfig)
	ResetDatabaseConfig()
	ResetEncryptionConfig()
	ResetEnvironmentSize()
	ResetMaintenanceWindow()
	ResetMasterAuthorizedNetworksConfig()
	ResetNodeConfig()
	ResetNodeCount()
	ResetPrivateEnvironmentConfig()
	ResetRecoveryConfig()
	ResetResilienceMode()
	ResetSoftwareConfig()
	ResetWebServerConfig()
	ResetWebServerNetworkAccessControl()
	ResetWorkloadsConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComposerEnvironmentConfigAOutputReference
type jsiiProxy_ComposerEnvironmentConfigAOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference) AirflowUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"airflowUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference) DagGcsPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dagGcsPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference) DatabaseConfig() ComposerEnvironmentConfigDatabaseConfigOutputReference {
	var returns ComposerEnvironmentConfigDatabaseConfigOutputReference
	_jsii_.Get(
		j,
		"databaseConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference) DatabaseConfigInput() *ComposerEnvironmentConfigDatabaseConfig {
	var returns *ComposerEnvironmentConfigDatabaseConfig
	_jsii_.Get(
		j,
		"databaseConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference) EncryptionConfig() ComposerEnvironmentConfigEncryptionConfigOutputReference {
	var returns ComposerEnvironmentConfigEncryptionConfigOutputReference
	_jsii_.Get(
		j,
		"encryptionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference) EncryptionConfigInput() *ComposerEnvironmentConfigEncryptionConfig {
	var returns *ComposerEnvironmentConfigEncryptionConfig
	_jsii_.Get(
		j,
		"encryptionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference) EnvironmentSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference) EnvironmentSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference) GkeCluster() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gkeCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference) InternalValue() *ComposerEnvironmentConfigA {
	var returns *ComposerEnvironmentConfigA
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference) MaintenanceWindow() ComposerEnvironmentConfigMaintenanceWindowOutputReference {
	var returns ComposerEnvironmentConfigMaintenanceWindowOutputReference
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference) MaintenanceWindowInput() *ComposerEnvironmentConfigMaintenanceWindow {
	var returns *ComposerEnvironmentConfigMaintenanceWindow
	_jsii_.Get(
		j,
		"maintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference) MasterAuthorizedNetworksConfig() ComposerEnvironmentConfigMasterAuthorizedNetworksConfigOutputReference {
	var returns ComposerEnvironmentConfigMasterAuthorizedNetworksConfigOutputReference
	_jsii_.Get(
		j,
		"masterAuthorizedNetworksConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference) MasterAuthorizedNetworksConfigInput() *ComposerEnvironmentConfigMasterAuthorizedNetworksConfig {
	var returns *ComposerEnvironmentConfigMasterAuthorizedNetworksConfig
	_jsii_.Get(
		j,
		"masterAuthorizedNetworksConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference) NodeConfig() ComposerEnvironmentConfigNodeConfigOutputReference {
	var returns ComposerEnvironmentConfigNodeConfigOutputReference
	_jsii_.Get(
		j,
		"nodeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference) NodeConfigInput() *ComposerEnvironmentConfigNodeConfig {
	var returns *ComposerEnvironmentConfigNodeConfig
	_jsii_.Get(
		j,
		"nodeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference) NodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference) NodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference) PrivateEnvironmentConfig() ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference {
	var returns ComposerEnvironmentConfigPrivateEnvironmentConfigOutputReference
	_jsii_.Get(
		j,
		"privateEnvironmentConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference) PrivateEnvironmentConfigInput() *ComposerEnvironmentConfigPrivateEnvironmentConfig {
	var returns *ComposerEnvironmentConfigPrivateEnvironmentConfig
	_jsii_.Get(
		j,
		"privateEnvironmentConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference) RecoveryConfig() ComposerEnvironmentConfigRecoveryConfigOutputReference {
	var returns ComposerEnvironmentConfigRecoveryConfigOutputReference
	_jsii_.Get(
		j,
		"recoveryConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference) RecoveryConfigInput() *ComposerEnvironmentConfigRecoveryConfig {
	var returns *ComposerEnvironmentConfigRecoveryConfig
	_jsii_.Get(
		j,
		"recoveryConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference) ResilienceMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resilienceMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference) ResilienceModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resilienceModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference) SoftwareConfig() ComposerEnvironmentConfigSoftwareConfigOutputReference {
	var returns ComposerEnvironmentConfigSoftwareConfigOutputReference
	_jsii_.Get(
		j,
		"softwareConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference) SoftwareConfigInput() *ComposerEnvironmentConfigSoftwareConfig {
	var returns *ComposerEnvironmentConfigSoftwareConfig
	_jsii_.Get(
		j,
		"softwareConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference) WebServerConfig() ComposerEnvironmentConfigWebServerConfigOutputReference {
	var returns ComposerEnvironmentConfigWebServerConfigOutputReference
	_jsii_.Get(
		j,
		"webServerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference) WebServerConfigInput() *ComposerEnvironmentConfigWebServerConfig {
	var returns *ComposerEnvironmentConfigWebServerConfig
	_jsii_.Get(
		j,
		"webServerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference) WebServerNetworkAccessControl() ComposerEnvironmentConfigWebServerNetworkAccessControlOutputReference {
	var returns ComposerEnvironmentConfigWebServerNetworkAccessControlOutputReference
	_jsii_.Get(
		j,
		"webServerNetworkAccessControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference) WebServerNetworkAccessControlInput() *ComposerEnvironmentConfigWebServerNetworkAccessControl {
	var returns *ComposerEnvironmentConfigWebServerNetworkAccessControl
	_jsii_.Get(
		j,
		"webServerNetworkAccessControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference) WorkloadsConfig() ComposerEnvironmentConfigWorkloadsConfigOutputReference {
	var returns ComposerEnvironmentConfigWorkloadsConfigOutputReference
	_jsii_.Get(
		j,
		"workloadsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference) WorkloadsConfigInput() *ComposerEnvironmentConfigWorkloadsConfig {
	var returns *ComposerEnvironmentConfigWorkloadsConfig
	_jsii_.Get(
		j,
		"workloadsConfigInput",
		&returns,
	)
	return returns
}


func NewComposerEnvironmentConfigAOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComposerEnvironmentConfigAOutputReference {
	_init_.Initialize()

	if err := validateNewComposerEnvironmentConfigAOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComposerEnvironmentConfigAOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.composerEnvironment.ComposerEnvironmentConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComposerEnvironmentConfigAOutputReference_Override(c ComposerEnvironmentConfigAOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.composerEnvironment.ComposerEnvironmentConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference)SetEnvironmentSize(val *string) {
	if err := j.validateSetEnvironmentSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentSize",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference)SetInternalValue(val *ComposerEnvironmentConfigA) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference)SetNodeCount(val *float64) {
	if err := j.validateSetNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeCount",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference)SetResilienceMode(val *string) {
	if err := j.validateSetResilienceModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resilienceMode",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigAOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigAOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComposerEnvironmentConfigAOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComposerEnvironmentConfigAOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComposerEnvironmentConfigAOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComposerEnvironmentConfigAOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComposerEnvironmentConfigAOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComposerEnvironmentConfigAOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComposerEnvironmentConfigAOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComposerEnvironmentConfigAOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComposerEnvironmentConfigAOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComposerEnvironmentConfigAOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComposerEnvironmentConfigAOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComposerEnvironmentConfigAOutputReference) PutDatabaseConfig(value *ComposerEnvironmentConfigDatabaseConfig) {
	if err := c.validatePutDatabaseConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDatabaseConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigAOutputReference) PutEncryptionConfig(value *ComposerEnvironmentConfigEncryptionConfig) {
	if err := c.validatePutEncryptionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEncryptionConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigAOutputReference) PutMaintenanceWindow(value *ComposerEnvironmentConfigMaintenanceWindow) {
	if err := c.validatePutMaintenanceWindowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMaintenanceWindow",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigAOutputReference) PutMasterAuthorizedNetworksConfig(value *ComposerEnvironmentConfigMasterAuthorizedNetworksConfig) {
	if err := c.validatePutMasterAuthorizedNetworksConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMasterAuthorizedNetworksConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigAOutputReference) PutNodeConfig(value *ComposerEnvironmentConfigNodeConfig) {
	if err := c.validatePutNodeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNodeConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigAOutputReference) PutPrivateEnvironmentConfig(value *ComposerEnvironmentConfigPrivateEnvironmentConfig) {
	if err := c.validatePutPrivateEnvironmentConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPrivateEnvironmentConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigAOutputReference) PutRecoveryConfig(value *ComposerEnvironmentConfigRecoveryConfig) {
	if err := c.validatePutRecoveryConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRecoveryConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigAOutputReference) PutSoftwareConfig(value *ComposerEnvironmentConfigSoftwareConfig) {
	if err := c.validatePutSoftwareConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSoftwareConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigAOutputReference) PutWebServerConfig(value *ComposerEnvironmentConfigWebServerConfig) {
	if err := c.validatePutWebServerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putWebServerConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigAOutputReference) PutWebServerNetworkAccessControl(value *ComposerEnvironmentConfigWebServerNetworkAccessControl) {
	if err := c.validatePutWebServerNetworkAccessControlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putWebServerNetworkAccessControl",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigAOutputReference) PutWorkloadsConfig(value *ComposerEnvironmentConfigWorkloadsConfig) {
	if err := c.validatePutWorkloadsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putWorkloadsConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigAOutputReference) ResetDatabaseConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetDatabaseConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigAOutputReference) ResetEncryptionConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetEncryptionConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigAOutputReference) ResetEnvironmentSize() {
	_jsii_.InvokeVoid(
		c,
		"resetEnvironmentSize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigAOutputReference) ResetMaintenanceWindow() {
	_jsii_.InvokeVoid(
		c,
		"resetMaintenanceWindow",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigAOutputReference) ResetMasterAuthorizedNetworksConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetMasterAuthorizedNetworksConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigAOutputReference) ResetNodeConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetNodeConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigAOutputReference) ResetNodeCount() {
	_jsii_.InvokeVoid(
		c,
		"resetNodeCount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigAOutputReference) ResetPrivateEnvironmentConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetPrivateEnvironmentConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigAOutputReference) ResetRecoveryConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetRecoveryConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigAOutputReference) ResetResilienceMode() {
	_jsii_.InvokeVoid(
		c,
		"resetResilienceMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigAOutputReference) ResetSoftwareConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetSoftwareConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigAOutputReference) ResetWebServerConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetWebServerConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigAOutputReference) ResetWebServerNetworkAccessControl() {
	_jsii_.InvokeVoid(
		c,
		"resetWebServerNetworkAccessControl",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigAOutputReference) ResetWorkloadsConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetWorkloadsConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigAOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComposerEnvironmentConfigAOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

