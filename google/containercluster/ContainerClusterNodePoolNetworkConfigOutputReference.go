// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v12/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v12/containercluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerClusterNodePoolNetworkConfigOutputReference interface {
	cdktf.ComplexObject
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
	InternalValue() *ContainerClusterNodePoolNetworkConfig
	SetInternalValue(val *ContainerClusterNodePoolNetworkConfig)
	NetworkPerformanceConfig() ContainerClusterNodePoolNetworkConfigNetworkPerformanceConfigOutputReference
	NetworkPerformanceConfigInput() *ContainerClusterNodePoolNetworkConfigNetworkPerformanceConfig
	PodCidrOverprovisionConfig() ContainerClusterNodePoolNetworkConfigPodCidrOverprovisionConfigOutputReference
	PodCidrOverprovisionConfigInput() *ContainerClusterNodePoolNetworkConfigPodCidrOverprovisionConfig
	PodIpv4CidrBlock() *string
	SetPodIpv4CidrBlock(val *string)
	PodIpv4CidrBlockInput() *string
	PodRange() *string
	SetPodRange(val *string)
	PodRangeInput() *string
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
	PutNetworkPerformanceConfig(value *ContainerClusterNodePoolNetworkConfigNetworkPerformanceConfig)
	PutPodCidrOverprovisionConfig(value *ContainerClusterNodePoolNetworkConfigPodCidrOverprovisionConfig)
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

// The jsii proxy struct for ContainerClusterNodePoolNetworkConfigOutputReference
type jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference) CreatePodRange() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createPodRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference) CreatePodRangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createPodRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference) EnablePrivateNodes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePrivateNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference) EnablePrivateNodesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePrivateNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference) InternalValue() *ContainerClusterNodePoolNetworkConfig {
	var returns *ContainerClusterNodePoolNetworkConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference) NetworkPerformanceConfig() ContainerClusterNodePoolNetworkConfigNetworkPerformanceConfigOutputReference {
	var returns ContainerClusterNodePoolNetworkConfigNetworkPerformanceConfigOutputReference
	_jsii_.Get(
		j,
		"networkPerformanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference) NetworkPerformanceConfigInput() *ContainerClusterNodePoolNetworkConfigNetworkPerformanceConfig {
	var returns *ContainerClusterNodePoolNetworkConfigNetworkPerformanceConfig
	_jsii_.Get(
		j,
		"networkPerformanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference) PodCidrOverprovisionConfig() ContainerClusterNodePoolNetworkConfigPodCidrOverprovisionConfigOutputReference {
	var returns ContainerClusterNodePoolNetworkConfigPodCidrOverprovisionConfigOutputReference
	_jsii_.Get(
		j,
		"podCidrOverprovisionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference) PodCidrOverprovisionConfigInput() *ContainerClusterNodePoolNetworkConfigPodCidrOverprovisionConfig {
	var returns *ContainerClusterNodePoolNetworkConfigPodCidrOverprovisionConfig
	_jsii_.Get(
		j,
		"podCidrOverprovisionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference) PodIpv4CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"podIpv4CidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference) PodIpv4CidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"podIpv4CidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference) PodRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"podRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference) PodRangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"podRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewContainerClusterNodePoolNetworkConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContainerClusterNodePoolNetworkConfigOutputReference {
	_init_.Initialize()

	if err := validateNewContainerClusterNodePoolNetworkConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.containerCluster.ContainerClusterNodePoolNetworkConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerClusterNodePoolNetworkConfigOutputReference_Override(c ContainerClusterNodePoolNetworkConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.containerCluster.ContainerClusterNodePoolNetworkConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference)SetCreatePodRange(val interface{}) {
	if err := j.validateSetCreatePodRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createPodRange",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference)SetEnablePrivateNodes(val interface{}) {
	if err := j.validateSetEnablePrivateNodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePrivateNodes",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference)SetInternalValue(val *ContainerClusterNodePoolNetworkConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference)SetPodIpv4CidrBlock(val *string) {
	if err := j.validateSetPodIpv4CidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"podIpv4CidrBlock",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference)SetPodRange(val *string) {
	if err := j.validateSetPodRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"podRange",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference) PutNetworkPerformanceConfig(value *ContainerClusterNodePoolNetworkConfigNetworkPerformanceConfig) {
	if err := c.validatePutNetworkPerformanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNetworkPerformanceConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference) PutPodCidrOverprovisionConfig(value *ContainerClusterNodePoolNetworkConfigPodCidrOverprovisionConfig) {
	if err := c.validatePutPodCidrOverprovisionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPodCidrOverprovisionConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference) ResetCreatePodRange() {
	_jsii_.InvokeVoid(
		c,
		"resetCreatePodRange",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference) ResetEnablePrivateNodes() {
	_jsii_.InvokeVoid(
		c,
		"resetEnablePrivateNodes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference) ResetNetworkPerformanceConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetNetworkPerformanceConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference) ResetPodCidrOverprovisionConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetPodCidrOverprovisionConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference) ResetPodIpv4CidrBlock() {
	_jsii_.InvokeVoid(
		c,
		"resetPodIpv4CidrBlock",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference) ResetPodRange() {
	_jsii_.InvokeVoid(
		c,
		"resetPodRange",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ContainerClusterNodePoolNetworkConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

