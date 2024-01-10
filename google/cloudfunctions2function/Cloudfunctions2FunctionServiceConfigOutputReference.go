// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfunctions2function

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v13/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v13/cloudfunctions2function/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Cloudfunctions2FunctionServiceConfigOutputReference interface {
	cdktf.ComplexObject
	AllTrafficOnLatestRevision() interface{}
	SetAllTrafficOnLatestRevision(val interface{})
	AllTrafficOnLatestRevisionInput() interface{}
	AvailableCpu() *string
	SetAvailableCpu(val *string)
	AvailableCpuInput() *string
	AvailableMemory() *string
	SetAvailableMemory(val *string)
	AvailableMemoryInput() *string
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
	EnvironmentVariables() *map[string]*string
	SetEnvironmentVariables(val *map[string]*string)
	EnvironmentVariablesInput() *map[string]*string
	// Experimental.
	Fqn() *string
	GcfUri() *string
	IngressSettings() *string
	SetIngressSettings(val *string)
	IngressSettingsInput() *string
	InternalValue() *Cloudfunctions2FunctionServiceConfig
	SetInternalValue(val *Cloudfunctions2FunctionServiceConfig)
	MaxInstanceCount() *float64
	SetMaxInstanceCount(val *float64)
	MaxInstanceCountInput() *float64
	MaxInstanceRequestConcurrency() *float64
	SetMaxInstanceRequestConcurrency(val *float64)
	MaxInstanceRequestConcurrencyInput() *float64
	MinInstanceCount() *float64
	SetMinInstanceCount(val *float64)
	MinInstanceCountInput() *float64
	SecretEnvironmentVariables() Cloudfunctions2FunctionServiceConfigSecretEnvironmentVariablesList
	SecretEnvironmentVariablesInput() interface{}
	SecretVolumes() Cloudfunctions2FunctionServiceConfigSecretVolumesList
	SecretVolumesInput() interface{}
	Service() *string
	SetService(val *string)
	ServiceAccountEmail() *string
	SetServiceAccountEmail(val *string)
	ServiceAccountEmailInput() *string
	ServiceInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeoutSeconds() *float64
	SetTimeoutSeconds(val *float64)
	TimeoutSecondsInput() *float64
	Uri() *string
	VpcConnector() *string
	SetVpcConnector(val *string)
	VpcConnectorEgressSettings() *string
	SetVpcConnectorEgressSettings(val *string)
	VpcConnectorEgressSettingsInput() *string
	VpcConnectorInput() *string
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
	PutSecretEnvironmentVariables(value interface{})
	PutSecretVolumes(value interface{})
	ResetAllTrafficOnLatestRevision()
	ResetAvailableCpu()
	ResetAvailableMemory()
	ResetEnvironmentVariables()
	ResetIngressSettings()
	ResetMaxInstanceCount()
	ResetMaxInstanceRequestConcurrency()
	ResetMinInstanceCount()
	ResetSecretEnvironmentVariables()
	ResetSecretVolumes()
	ResetService()
	ResetServiceAccountEmail()
	ResetTimeoutSeconds()
	ResetVpcConnector()
	ResetVpcConnectorEgressSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Cloudfunctions2FunctionServiceConfigOutputReference
type jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) AllTrafficOnLatestRevision() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allTrafficOnLatestRevision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) AllTrafficOnLatestRevisionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allTrafficOnLatestRevisionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) AvailableCpu() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availableCpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) AvailableCpuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availableCpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) AvailableMemory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availableMemory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) AvailableMemoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availableMemoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) EnvironmentVariables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) EnvironmentVariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) GcfUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcfUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) IngressSettings() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ingressSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) IngressSettingsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ingressSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) InternalValue() *Cloudfunctions2FunctionServiceConfig {
	var returns *Cloudfunctions2FunctionServiceConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) MaxInstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInstanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) MaxInstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInstanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) MaxInstanceRequestConcurrency() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInstanceRequestConcurrency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) MaxInstanceRequestConcurrencyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxInstanceRequestConcurrencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) MinInstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minInstanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) MinInstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minInstanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) SecretEnvironmentVariables() Cloudfunctions2FunctionServiceConfigSecretEnvironmentVariablesList {
	var returns Cloudfunctions2FunctionServiceConfigSecretEnvironmentVariablesList
	_jsii_.Get(
		j,
		"secretEnvironmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) SecretEnvironmentVariablesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretEnvironmentVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) SecretVolumes() Cloudfunctions2FunctionServiceConfigSecretVolumesList {
	var returns Cloudfunctions2FunctionServiceConfigSecretVolumesList
	_jsii_.Get(
		j,
		"secretVolumes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) SecretVolumesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretVolumesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) Service() *string {
	var returns *string
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) ServiceAccountEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) ServiceAccountEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) ServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) TimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) TimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) VpcConnector() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcConnector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) VpcConnectorEgressSettings() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcConnectorEgressSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) VpcConnectorEgressSettingsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcConnectorEgressSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) VpcConnectorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcConnectorInput",
		&returns,
	)
	return returns
}


func NewCloudfunctions2FunctionServiceConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Cloudfunctions2FunctionServiceConfigOutputReference {
	_init_.Initialize()

	if err := validateNewCloudfunctions2FunctionServiceConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cloudfunctions2Function.Cloudfunctions2FunctionServiceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudfunctions2FunctionServiceConfigOutputReference_Override(c Cloudfunctions2FunctionServiceConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cloudfunctions2Function.Cloudfunctions2FunctionServiceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference)SetAllTrafficOnLatestRevision(val interface{}) {
	if err := j.validateSetAllTrafficOnLatestRevisionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allTrafficOnLatestRevision",
		val,
	)
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference)SetAvailableCpu(val *string) {
	if err := j.validateSetAvailableCpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availableCpu",
		val,
	)
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference)SetAvailableMemory(val *string) {
	if err := j.validateSetAvailableMemoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availableMemory",
		val,
	)
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference)SetEnvironmentVariables(val *map[string]*string) {
	if err := j.validateSetEnvironmentVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentVariables",
		val,
	)
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference)SetIngressSettings(val *string) {
	if err := j.validateSetIngressSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ingressSettings",
		val,
	)
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference)SetInternalValue(val *Cloudfunctions2FunctionServiceConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference)SetMaxInstanceCount(val *float64) {
	if err := j.validateSetMaxInstanceCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxInstanceCount",
		val,
	)
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference)SetMaxInstanceRequestConcurrency(val *float64) {
	if err := j.validateSetMaxInstanceRequestConcurrencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxInstanceRequestConcurrency",
		val,
	)
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference)SetMinInstanceCount(val *float64) {
	if err := j.validateSetMinInstanceCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minInstanceCount",
		val,
	)
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference)SetService(val *string) {
	if err := j.validateSetServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"service",
		val,
	)
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference)SetServiceAccountEmail(val *string) {
	if err := j.validateSetServiceAccountEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccountEmail",
		val,
	)
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference)SetTimeoutSeconds(val *float64) {
	if err := j.validateSetTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutSeconds",
		val,
	)
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference)SetVpcConnector(val *string) {
	if err := j.validateSetVpcConnectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcConnector",
		val,
	)
}

func (j *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference)SetVpcConnectorEgressSettings(val *string) {
	if err := j.validateSetVpcConnectorEgressSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcConnectorEgressSettings",
		val,
	)
}

func (c *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) PutSecretEnvironmentVariables(value interface{}) {
	if err := c.validatePutSecretEnvironmentVariablesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSecretEnvironmentVariables",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) PutSecretVolumes(value interface{}) {
	if err := c.validatePutSecretVolumesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSecretVolumes",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) ResetAllTrafficOnLatestRevision() {
	_jsii_.InvokeVoid(
		c,
		"resetAllTrafficOnLatestRevision",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) ResetAvailableCpu() {
	_jsii_.InvokeVoid(
		c,
		"resetAvailableCpu",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) ResetAvailableMemory() {
	_jsii_.InvokeVoid(
		c,
		"resetAvailableMemory",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) ResetEnvironmentVariables() {
	_jsii_.InvokeVoid(
		c,
		"resetEnvironmentVariables",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) ResetIngressSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetIngressSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) ResetMaxInstanceCount() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxInstanceCount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) ResetMaxInstanceRequestConcurrency() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxInstanceRequestConcurrency",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) ResetMinInstanceCount() {
	_jsii_.InvokeVoid(
		c,
		"resetMinInstanceCount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) ResetSecretEnvironmentVariables() {
	_jsii_.InvokeVoid(
		c,
		"resetSecretEnvironmentVariables",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) ResetSecretVolumes() {
	_jsii_.InvokeVoid(
		c,
		"resetSecretVolumes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) ResetService() {
	_jsii_.InvokeVoid(
		c,
		"resetService",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) ResetServiceAccountEmail() {
	_jsii_.InvokeVoid(
		c,
		"resetServiceAccountEmail",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) ResetTimeoutSeconds() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeoutSeconds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) ResetVpcConnector() {
	_jsii_.InvokeVoid(
		c,
		"resetVpcConnector",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) ResetVpcConnectorEgressSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetVpcConnectorEgressSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_Cloudfunctions2FunctionServiceConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

