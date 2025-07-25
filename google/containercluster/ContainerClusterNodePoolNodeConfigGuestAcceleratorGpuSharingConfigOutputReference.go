// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/containercluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReference interface {
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	GpuSharingStrategy() *string
	SetGpuSharingStrategy(val *string)
	GpuSharingStrategyInput() *string
	InternalValue() *ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfig
	SetInternalValue(val *ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfig)
	MaxSharedClientsPerGpu() *float64
	SetMaxSharedClientsPerGpu(val *float64)
	MaxSharedClientsPerGpuInput() *float64
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReference
type jsiiProxy_ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReference) GpuSharingStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gpuSharingStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReference) GpuSharingStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gpuSharingStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReference) InternalValue() *ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfig {
	var returns *ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReference) MaxSharedClientsPerGpu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSharedClientsPerGpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReference) MaxSharedClientsPerGpuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSharedClientsPerGpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReference {
	_init_.Initialize()

	if err := validateNewContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.containerCluster.ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReference_Override(c ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.containerCluster.ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReference)SetGpuSharingStrategy(val *string) {
	if err := j.validateSetGpuSharingStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gpuSharingStrategy",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReference)SetInternalValue(val *ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReference)SetMaxSharedClientsPerGpu(val *float64) {
	if err := j.validateSetMaxSharedClientsPerGpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSharedClientsPerGpu",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

