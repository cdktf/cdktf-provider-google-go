// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containernodepool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/containernodepool/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerNodePoolNetworkConfigOutputReference interface {
	cdktf.ComplexObject
	AdditionalNodeNetworkConfigs() ContainerNodePoolNetworkConfigAdditionalNodeNetworkConfigsList
	AdditionalNodeNetworkConfigsInput() interface{}
	AdditionalPodNetworkConfigs() ContainerNodePoolNetworkConfigAdditionalPodNetworkConfigsList
	AdditionalPodNetworkConfigsInput() interface{}
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
	CreatePodRange() interface{}
	SetCreatePodRange(val interface{})
	CreatePodRangeInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EnablePrivateNodes() interface{}
	SetEnablePrivateNodes(val interface{})
	EnablePrivateNodesInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *ContainerNodePoolNetworkConfig
	SetInternalValue(val *ContainerNodePoolNetworkConfig)
	NetworkPerformanceConfig() ContainerNodePoolNetworkConfigNetworkPerformanceConfigOutputReference
	NetworkPerformanceConfigInput() *ContainerNodePoolNetworkConfigNetworkPerformanceConfig
	PodCidrOverprovisionConfig() ContainerNodePoolNetworkConfigPodCidrOverprovisionConfigOutputReference
	PodCidrOverprovisionConfigInput() *ContainerNodePoolNetworkConfigPodCidrOverprovisionConfig
	PodIpv4CidrBlock() *string
	SetPodIpv4CidrBlock(val *string)
	PodIpv4CidrBlockInput() *string
	PodRange() *string
	SetPodRange(val *string)
	PodRangeInput() *string
	Subnetwork() *string
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
	PutAdditionalNodeNetworkConfigs(value interface{})
	PutAdditionalPodNetworkConfigs(value interface{})
	PutNetworkPerformanceConfig(value *ContainerNodePoolNetworkConfigNetworkPerformanceConfig)
	PutPodCidrOverprovisionConfig(value *ContainerNodePoolNetworkConfigPodCidrOverprovisionConfig)
	ResetAdditionalNodeNetworkConfigs()
	ResetAdditionalPodNetworkConfigs()
	ResetCreatePodRange()
	ResetEnablePrivateNodes()
	ResetNetworkPerformanceConfig()
	ResetPodCidrOverprovisionConfig()
	ResetPodIpv4CidrBlock()
	ResetPodRange()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerNodePoolNetworkConfigOutputReference
type jsiiProxy_ContainerNodePoolNetworkConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) AdditionalNodeNetworkConfigs() ContainerNodePoolNetworkConfigAdditionalNodeNetworkConfigsList {
	var returns ContainerNodePoolNetworkConfigAdditionalNodeNetworkConfigsList
	_jsii_.Get(
		j,
		"additionalNodeNetworkConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) AdditionalNodeNetworkConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalNodeNetworkConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) AdditionalPodNetworkConfigs() ContainerNodePoolNetworkConfigAdditionalPodNetworkConfigsList {
	var returns ContainerNodePoolNetworkConfigAdditionalPodNetworkConfigsList
	_jsii_.Get(
		j,
		"additionalPodNetworkConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) AdditionalPodNetworkConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalPodNetworkConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) CreatePodRange() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createPodRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) CreatePodRangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createPodRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) EnablePrivateNodes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePrivateNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) EnablePrivateNodesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePrivateNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) InternalValue() *ContainerNodePoolNetworkConfig {
	var returns *ContainerNodePoolNetworkConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) NetworkPerformanceConfig() ContainerNodePoolNetworkConfigNetworkPerformanceConfigOutputReference {
	var returns ContainerNodePoolNetworkConfigNetworkPerformanceConfigOutputReference
	_jsii_.Get(
		j,
		"networkPerformanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) NetworkPerformanceConfigInput() *ContainerNodePoolNetworkConfigNetworkPerformanceConfig {
	var returns *ContainerNodePoolNetworkConfigNetworkPerformanceConfig
	_jsii_.Get(
		j,
		"networkPerformanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) PodCidrOverprovisionConfig() ContainerNodePoolNetworkConfigPodCidrOverprovisionConfigOutputReference {
	var returns ContainerNodePoolNetworkConfigPodCidrOverprovisionConfigOutputReference
	_jsii_.Get(
		j,
		"podCidrOverprovisionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) PodCidrOverprovisionConfigInput() *ContainerNodePoolNetworkConfigPodCidrOverprovisionConfig {
	var returns *ContainerNodePoolNetworkConfigPodCidrOverprovisionConfig
	_jsii_.Get(
		j,
		"podCidrOverprovisionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) PodIpv4CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"podIpv4CidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) PodIpv4CidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"podIpv4CidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) PodRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"podRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) PodRangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"podRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) Subnetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewContainerNodePoolNetworkConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContainerNodePoolNetworkConfigOutputReference {
	_init_.Initialize()

	if err := validateNewContainerNodePoolNetworkConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerNodePoolNetworkConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.containerNodePool.ContainerNodePoolNetworkConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerNodePoolNetworkConfigOutputReference_Override(c ContainerNodePoolNetworkConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.containerNodePool.ContainerNodePoolNetworkConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference)SetCreatePodRange(val interface{}) {
	if err := j.validateSetCreatePodRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createPodRange",
		val,
	)
}

func (j *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference)SetEnablePrivateNodes(val interface{}) {
	if err := j.validateSetEnablePrivateNodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePrivateNodes",
		val,
	)
}

func (j *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference)SetInternalValue(val *ContainerNodePoolNetworkConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference)SetPodIpv4CidrBlock(val *string) {
	if err := j.validateSetPodIpv4CidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"podIpv4CidrBlock",
		val,
	)
}

func (j *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference)SetPodRange(val *string) {
	if err := j.validateSetPodRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"podRange",
		val,
	)
}

func (j *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) PutAdditionalNodeNetworkConfigs(value interface{}) {
	if err := c.validatePutAdditionalNodeNetworkConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAdditionalNodeNetworkConfigs",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) PutAdditionalPodNetworkConfigs(value interface{}) {
	if err := c.validatePutAdditionalPodNetworkConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAdditionalPodNetworkConfigs",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) PutNetworkPerformanceConfig(value *ContainerNodePoolNetworkConfigNetworkPerformanceConfig) {
	if err := c.validatePutNetworkPerformanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNetworkPerformanceConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) PutPodCidrOverprovisionConfig(value *ContainerNodePoolNetworkConfigPodCidrOverprovisionConfig) {
	if err := c.validatePutPodCidrOverprovisionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPodCidrOverprovisionConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) ResetAdditionalNodeNetworkConfigs() {
	_jsii_.InvokeVoid(
		c,
		"resetAdditionalNodeNetworkConfigs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) ResetAdditionalPodNetworkConfigs() {
	_jsii_.InvokeVoid(
		c,
		"resetAdditionalPodNetworkConfigs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) ResetCreatePodRange() {
	_jsii_.InvokeVoid(
		c,
		"resetCreatePodRange",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) ResetEnablePrivateNodes() {
	_jsii_.InvokeVoid(
		c,
		"resetEnablePrivateNodes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) ResetNetworkPerformanceConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetNetworkPerformanceConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) ResetPodCidrOverprovisionConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetPodCidrOverprovisionConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) ResetPodIpv4CidrBlock() {
	_jsii_.InvokeVoid(
		c,
		"resetPodIpv4CidrBlock",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) ResetPodRange() {
	_jsii_.InvokeVoid(
		c,
		"resetPodRange",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ContainerNodePoolNetworkConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

