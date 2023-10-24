// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonpremvmwarecluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v12/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v12/gkeonpremvmwarecluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GkeonpremVmwareClusterNetworkConfigOutputReference interface {
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
	ControlPlaneV2Config() GkeonpremVmwareClusterNetworkConfigControlPlaneV2ConfigOutputReference
	ControlPlaneV2ConfigInput() *GkeonpremVmwareClusterNetworkConfigControlPlaneV2Config
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DhcpIpConfig() GkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReference
	DhcpIpConfigInput() *GkeonpremVmwareClusterNetworkConfigDhcpIpConfig
	// Experimental.
	Fqn() *string
	HostConfig() GkeonpremVmwareClusterNetworkConfigHostConfigOutputReference
	HostConfigInput() *GkeonpremVmwareClusterNetworkConfigHostConfig
	InternalValue() *GkeonpremVmwareClusterNetworkConfig
	SetInternalValue(val *GkeonpremVmwareClusterNetworkConfig)
	PodAddressCidrBlocks() *[]*string
	SetPodAddressCidrBlocks(val *[]*string)
	PodAddressCidrBlocksInput() *[]*string
	ServiceAddressCidrBlocks() *[]*string
	SetServiceAddressCidrBlocks(val *[]*string)
	ServiceAddressCidrBlocksInput() *[]*string
	StaticIpConfig() GkeonpremVmwareClusterNetworkConfigStaticIpConfigOutputReference
	StaticIpConfigInput() *GkeonpremVmwareClusterNetworkConfigStaticIpConfig
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VcenterNetwork() *string
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
	PutControlPlaneV2Config(value *GkeonpremVmwareClusterNetworkConfigControlPlaneV2Config)
	PutDhcpIpConfig(value *GkeonpremVmwareClusterNetworkConfigDhcpIpConfig)
	PutHostConfig(value *GkeonpremVmwareClusterNetworkConfigHostConfig)
	PutStaticIpConfig(value *GkeonpremVmwareClusterNetworkConfigStaticIpConfig)
	ResetControlPlaneV2Config()
	ResetDhcpIpConfig()
	ResetHostConfig()
	ResetStaticIpConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GkeonpremVmwareClusterNetworkConfigOutputReference
type jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference) ControlPlaneV2Config() GkeonpremVmwareClusterNetworkConfigControlPlaneV2ConfigOutputReference {
	var returns GkeonpremVmwareClusterNetworkConfigControlPlaneV2ConfigOutputReference
	_jsii_.Get(
		j,
		"controlPlaneV2Config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference) ControlPlaneV2ConfigInput() *GkeonpremVmwareClusterNetworkConfigControlPlaneV2Config {
	var returns *GkeonpremVmwareClusterNetworkConfigControlPlaneV2Config
	_jsii_.Get(
		j,
		"controlPlaneV2ConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference) DhcpIpConfig() GkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReference {
	var returns GkeonpremVmwareClusterNetworkConfigDhcpIpConfigOutputReference
	_jsii_.Get(
		j,
		"dhcpIpConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference) DhcpIpConfigInput() *GkeonpremVmwareClusterNetworkConfigDhcpIpConfig {
	var returns *GkeonpremVmwareClusterNetworkConfigDhcpIpConfig
	_jsii_.Get(
		j,
		"dhcpIpConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference) HostConfig() GkeonpremVmwareClusterNetworkConfigHostConfigOutputReference {
	var returns GkeonpremVmwareClusterNetworkConfigHostConfigOutputReference
	_jsii_.Get(
		j,
		"hostConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference) HostConfigInput() *GkeonpremVmwareClusterNetworkConfigHostConfig {
	var returns *GkeonpremVmwareClusterNetworkConfigHostConfig
	_jsii_.Get(
		j,
		"hostConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference) InternalValue() *GkeonpremVmwareClusterNetworkConfig {
	var returns *GkeonpremVmwareClusterNetworkConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference) PodAddressCidrBlocks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"podAddressCidrBlocks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference) PodAddressCidrBlocksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"podAddressCidrBlocksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference) ServiceAddressCidrBlocks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serviceAddressCidrBlocks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference) ServiceAddressCidrBlocksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serviceAddressCidrBlocksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference) StaticIpConfig() GkeonpremVmwareClusterNetworkConfigStaticIpConfigOutputReference {
	var returns GkeonpremVmwareClusterNetworkConfigStaticIpConfigOutputReference
	_jsii_.Get(
		j,
		"staticIpConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference) StaticIpConfigInput() *GkeonpremVmwareClusterNetworkConfigStaticIpConfig {
	var returns *GkeonpremVmwareClusterNetworkConfigStaticIpConfig
	_jsii_.Get(
		j,
		"staticIpConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference) VcenterNetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vcenterNetwork",
		&returns,
	)
	return returns
}


func NewGkeonpremVmwareClusterNetworkConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GkeonpremVmwareClusterNetworkConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGkeonpremVmwareClusterNetworkConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.gkeonpremVmwareCluster.GkeonpremVmwareClusterNetworkConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGkeonpremVmwareClusterNetworkConfigOutputReference_Override(g GkeonpremVmwareClusterNetworkConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.gkeonpremVmwareCluster.GkeonpremVmwareClusterNetworkConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference)SetInternalValue(val *GkeonpremVmwareClusterNetworkConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference)SetPodAddressCidrBlocks(val *[]*string) {
	if err := j.validateSetPodAddressCidrBlocksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"podAddressCidrBlocks",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference)SetServiceAddressCidrBlocks(val *[]*string) {
	if err := j.validateSetServiceAddressCidrBlocksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAddressCidrBlocks",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference) PutControlPlaneV2Config(value *GkeonpremVmwareClusterNetworkConfigControlPlaneV2Config) {
	if err := g.validatePutControlPlaneV2ConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putControlPlaneV2Config",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference) PutDhcpIpConfig(value *GkeonpremVmwareClusterNetworkConfigDhcpIpConfig) {
	if err := g.validatePutDhcpIpConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDhcpIpConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference) PutHostConfig(value *GkeonpremVmwareClusterNetworkConfigHostConfig) {
	if err := g.validatePutHostConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHostConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference) PutStaticIpConfig(value *GkeonpremVmwareClusterNetworkConfigStaticIpConfig) {
	if err := g.validatePutStaticIpConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStaticIpConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference) ResetControlPlaneV2Config() {
	_jsii_.InvokeVoid(
		g,
		"resetControlPlaneV2Config",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference) ResetDhcpIpConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetDhcpIpConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference) ResetHostConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetHostConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference) ResetStaticIpConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetStaticIpConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareClusterNetworkConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

