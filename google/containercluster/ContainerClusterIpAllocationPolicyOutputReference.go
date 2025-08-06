// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/containercluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerClusterIpAllocationPolicyOutputReference interface {
	cdktf.ComplexObject
	AdditionalIpRangesConfig() ContainerClusterIpAllocationPolicyAdditionalIpRangesConfigList
	AdditionalIpRangesConfigInput() interface{}
	AdditionalPodRangesConfig() ContainerClusterIpAllocationPolicyAdditionalPodRangesConfigOutputReference
	AdditionalPodRangesConfigInput() *ContainerClusterIpAllocationPolicyAdditionalPodRangesConfig
	ClusterIpv4CidrBlock() *string
	SetClusterIpv4CidrBlock(val *string)
	ClusterIpv4CidrBlockInput() *string
	ClusterSecondaryRangeName() *string
	SetClusterSecondaryRangeName(val *string)
	ClusterSecondaryRangeNameInput() *string
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
	InternalValue() *ContainerClusterIpAllocationPolicy
	SetInternalValue(val *ContainerClusterIpAllocationPolicy)
	PodCidrOverprovisionConfig() ContainerClusterIpAllocationPolicyPodCidrOverprovisionConfigOutputReference
	PodCidrOverprovisionConfigInput() *ContainerClusterIpAllocationPolicyPodCidrOverprovisionConfig
	ServicesIpv4CidrBlock() *string
	SetServicesIpv4CidrBlock(val *string)
	ServicesIpv4CidrBlockInput() *string
	ServicesSecondaryRangeName() *string
	SetServicesSecondaryRangeName(val *string)
	ServicesSecondaryRangeNameInput() *string
	StackType() *string
	SetStackType(val *string)
	StackTypeInput() *string
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
	PutAdditionalIpRangesConfig(value interface{})
	PutAdditionalPodRangesConfig(value *ContainerClusterIpAllocationPolicyAdditionalPodRangesConfig)
	PutPodCidrOverprovisionConfig(value *ContainerClusterIpAllocationPolicyPodCidrOverprovisionConfig)
	ResetAdditionalIpRangesConfig()
	ResetAdditionalPodRangesConfig()
	ResetClusterIpv4CidrBlock()
	ResetClusterSecondaryRangeName()
	ResetPodCidrOverprovisionConfig()
	ResetServicesIpv4CidrBlock()
	ResetServicesSecondaryRangeName()
	ResetStackType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerClusterIpAllocationPolicyOutputReference
type jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) AdditionalIpRangesConfig() ContainerClusterIpAllocationPolicyAdditionalIpRangesConfigList {
	var returns ContainerClusterIpAllocationPolicyAdditionalIpRangesConfigList
	_jsii_.Get(
		j,
		"additionalIpRangesConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) AdditionalIpRangesConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalIpRangesConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) AdditionalPodRangesConfig() ContainerClusterIpAllocationPolicyAdditionalPodRangesConfigOutputReference {
	var returns ContainerClusterIpAllocationPolicyAdditionalPodRangesConfigOutputReference
	_jsii_.Get(
		j,
		"additionalPodRangesConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) AdditionalPodRangesConfigInput() *ContainerClusterIpAllocationPolicyAdditionalPodRangesConfig {
	var returns *ContainerClusterIpAllocationPolicyAdditionalPodRangesConfig
	_jsii_.Get(
		j,
		"additionalPodRangesConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) ClusterIpv4CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIpv4CidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) ClusterIpv4CidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIpv4CidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) ClusterSecondaryRangeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterSecondaryRangeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) ClusterSecondaryRangeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterSecondaryRangeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) InternalValue() *ContainerClusterIpAllocationPolicy {
	var returns *ContainerClusterIpAllocationPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) PodCidrOverprovisionConfig() ContainerClusterIpAllocationPolicyPodCidrOverprovisionConfigOutputReference {
	var returns ContainerClusterIpAllocationPolicyPodCidrOverprovisionConfigOutputReference
	_jsii_.Get(
		j,
		"podCidrOverprovisionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) PodCidrOverprovisionConfigInput() *ContainerClusterIpAllocationPolicyPodCidrOverprovisionConfig {
	var returns *ContainerClusterIpAllocationPolicyPodCidrOverprovisionConfig
	_jsii_.Get(
		j,
		"podCidrOverprovisionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) ServicesIpv4CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicesIpv4CidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) ServicesIpv4CidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicesIpv4CidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) ServicesSecondaryRangeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicesSecondaryRangeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) ServicesSecondaryRangeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicesSecondaryRangeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) StackType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) StackTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewContainerClusterIpAllocationPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContainerClusterIpAllocationPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewContainerClusterIpAllocationPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.containerCluster.ContainerClusterIpAllocationPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerClusterIpAllocationPolicyOutputReference_Override(c ContainerClusterIpAllocationPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.containerCluster.ContainerClusterIpAllocationPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference)SetClusterIpv4CidrBlock(val *string) {
	if err := j.validateSetClusterIpv4CidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterIpv4CidrBlock",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference)SetClusterSecondaryRangeName(val *string) {
	if err := j.validateSetClusterSecondaryRangeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterSecondaryRangeName",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference)SetInternalValue(val *ContainerClusterIpAllocationPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference)SetServicesIpv4CidrBlock(val *string) {
	if err := j.validateSetServicesIpv4CidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"servicesIpv4CidrBlock",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference)SetServicesSecondaryRangeName(val *string) {
	if err := j.validateSetServicesSecondaryRangeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"servicesSecondaryRangeName",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference)SetStackType(val *string) {
	if err := j.validateSetStackTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stackType",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) PutAdditionalIpRangesConfig(value interface{}) {
	if err := c.validatePutAdditionalIpRangesConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAdditionalIpRangesConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) PutAdditionalPodRangesConfig(value *ContainerClusterIpAllocationPolicyAdditionalPodRangesConfig) {
	if err := c.validatePutAdditionalPodRangesConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAdditionalPodRangesConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) PutPodCidrOverprovisionConfig(value *ContainerClusterIpAllocationPolicyPodCidrOverprovisionConfig) {
	if err := c.validatePutPodCidrOverprovisionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPodCidrOverprovisionConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) ResetAdditionalIpRangesConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetAdditionalIpRangesConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) ResetAdditionalPodRangesConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetAdditionalPodRangesConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) ResetClusterIpv4CidrBlock() {
	_jsii_.InvokeVoid(
		c,
		"resetClusterIpv4CidrBlock",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) ResetClusterSecondaryRangeName() {
	_jsii_.InvokeVoid(
		c,
		"resetClusterSecondaryRangeName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) ResetPodCidrOverprovisionConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetPodCidrOverprovisionConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) ResetServicesIpv4CidrBlock() {
	_jsii_.InvokeVoid(
		c,
		"resetServicesIpv4CidrBlock",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) ResetServicesSecondaryRangeName() {
	_jsii_.InvokeVoid(
		c,
		"resetServicesSecondaryRangeName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) ResetStackType() {
	_jsii_.InvokeVoid(
		c,
		"resetStackType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ContainerClusterIpAllocationPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

