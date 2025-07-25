// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/containercluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference interface {
	cdktf.ComplexObject
	BootDiskKmsKey() *string
	SetBootDiskKmsKey(val *string)
	BootDiskKmsKeyInput() *string
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
	DiskSize() *float64
	SetDiskSize(val *float64)
	DiskSizeInput() *float64
	DiskType() *string
	SetDiskType(val *string)
	DiskTypeInput() *string
	// Experimental.
	Fqn() *string
	ImageType() *string
	SetImageType(val *string)
	ImageTypeInput() *string
	InternalValue() *ContainerClusterClusterAutoscalingAutoProvisioningDefaults
	SetInternalValue(val *ContainerClusterClusterAutoscalingAutoProvisioningDefaults)
	Management() ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference
	ManagementInput() *ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagement
	MinCpuPlatform() *string
	SetMinCpuPlatform(val *string)
	MinCpuPlatformInput() *string
	OauthScopes() *[]*string
	SetOauthScopes(val *[]*string)
	OauthScopesInput() *[]*string
	ServiceAccount() *string
	SetServiceAccount(val *string)
	ServiceAccountInput() *string
	ShieldedInstanceConfig() ContainerClusterClusterAutoscalingAutoProvisioningDefaultsShieldedInstanceConfigOutputReference
	ShieldedInstanceConfigInput() *ContainerClusterClusterAutoscalingAutoProvisioningDefaultsShieldedInstanceConfig
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UpgradeSettings() ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference
	UpgradeSettingsInput() *ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettings
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
	PutManagement(value *ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagement)
	PutShieldedInstanceConfig(value *ContainerClusterClusterAutoscalingAutoProvisioningDefaultsShieldedInstanceConfig)
	PutUpgradeSettings(value *ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettings)
	ResetBootDiskKmsKey()
	ResetDiskSize()
	ResetDiskType()
	ResetImageType()
	ResetManagement()
	ResetMinCpuPlatform()
	ResetOauthScopes()
	ResetServiceAccount()
	ResetShieldedInstanceConfig()
	ResetUpgradeSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference
type jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) BootDiskKmsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootDiskKmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) BootDiskKmsKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootDiskKmsKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) DiskSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) DiskSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) DiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) DiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) ImageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) ImageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) InternalValue() *ContainerClusterClusterAutoscalingAutoProvisioningDefaults {
	var returns *ContainerClusterClusterAutoscalingAutoProvisioningDefaults
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) Management() ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference {
	var returns ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagementOutputReference
	_jsii_.Get(
		j,
		"management",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) ManagementInput() *ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagement {
	var returns *ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagement
	_jsii_.Get(
		j,
		"managementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) MinCpuPlatform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCpuPlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) MinCpuPlatformInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCpuPlatformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) OauthScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"oauthScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) OauthScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"oauthScopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) ServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) ShieldedInstanceConfig() ContainerClusterClusterAutoscalingAutoProvisioningDefaultsShieldedInstanceConfigOutputReference {
	var returns ContainerClusterClusterAutoscalingAutoProvisioningDefaultsShieldedInstanceConfigOutputReference
	_jsii_.Get(
		j,
		"shieldedInstanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) ShieldedInstanceConfigInput() *ContainerClusterClusterAutoscalingAutoProvisioningDefaultsShieldedInstanceConfig {
	var returns *ContainerClusterClusterAutoscalingAutoProvisioningDefaultsShieldedInstanceConfig
	_jsii_.Get(
		j,
		"shieldedInstanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) UpgradeSettings() ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference {
	var returns ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference
	_jsii_.Get(
		j,
		"upgradeSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) UpgradeSettingsInput() *ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettings {
	var returns *ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettings
	_jsii_.Get(
		j,
		"upgradeSettingsInput",
		&returns,
	)
	return returns
}


func NewContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference {
	_init_.Initialize()

	if err := validateNewContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.containerCluster.ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference_Override(c ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.containerCluster.ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference)SetBootDiskKmsKey(val *string) {
	if err := j.validateSetBootDiskKmsKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootDiskKmsKey",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference)SetDiskSize(val *float64) {
	if err := j.validateSetDiskSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskSize",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference)SetDiskType(val *string) {
	if err := j.validateSetDiskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskType",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference)SetImageType(val *string) {
	if err := j.validateSetImageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageType",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference)SetInternalValue(val *ContainerClusterClusterAutoscalingAutoProvisioningDefaults) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference)SetMinCpuPlatform(val *string) {
	if err := j.validateSetMinCpuPlatformParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minCpuPlatform",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference)SetOauthScopes(val *[]*string) {
	if err := j.validateSetOauthScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthScopes",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference)SetServiceAccount(val *string) {
	if err := j.validateSetServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccount",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) PutManagement(value *ContainerClusterClusterAutoscalingAutoProvisioningDefaultsManagement) {
	if err := c.validatePutManagementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putManagement",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) PutShieldedInstanceConfig(value *ContainerClusterClusterAutoscalingAutoProvisioningDefaultsShieldedInstanceConfig) {
	if err := c.validatePutShieldedInstanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putShieldedInstanceConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) PutUpgradeSettings(value *ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettings) {
	if err := c.validatePutUpgradeSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putUpgradeSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) ResetBootDiskKmsKey() {
	_jsii_.InvokeVoid(
		c,
		"resetBootDiskKmsKey",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) ResetDiskSize() {
	_jsii_.InvokeVoid(
		c,
		"resetDiskSize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) ResetDiskType() {
	_jsii_.InvokeVoid(
		c,
		"resetDiskType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) ResetImageType() {
	_jsii_.InvokeVoid(
		c,
		"resetImageType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) ResetManagement() {
	_jsii_.InvokeVoid(
		c,
		"resetManagement",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) ResetMinCpuPlatform() {
	_jsii_.InvokeVoid(
		c,
		"resetMinCpuPlatform",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) ResetOauthScopes() {
	_jsii_.InvokeVoid(
		c,
		"resetOauthScopes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		c,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) ResetShieldedInstanceConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetShieldedInstanceConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) ResetUpgradeSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetUpgradeSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

