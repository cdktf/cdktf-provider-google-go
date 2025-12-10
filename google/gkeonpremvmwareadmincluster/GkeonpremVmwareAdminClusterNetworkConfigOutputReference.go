// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonpremvmwareadmincluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/gkeonpremvmwareadmincluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GkeonpremVmwareAdminClusterNetworkConfigOutputReference interface {
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
	DhcpIpConfig() GkeonpremVmwareAdminClusterNetworkConfigDhcpIpConfigOutputReference
	DhcpIpConfigInput() *GkeonpremVmwareAdminClusterNetworkConfigDhcpIpConfig
	// Experimental.
	Fqn() *string
	HaControlPlaneConfig() GkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigOutputReference
	HaControlPlaneConfigInput() *GkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfig
	HostConfig() GkeonpremVmwareAdminClusterNetworkConfigHostConfigOutputReference
	HostConfigInput() *GkeonpremVmwareAdminClusterNetworkConfigHostConfig
	InternalValue() *GkeonpremVmwareAdminClusterNetworkConfig
	SetInternalValue(val *GkeonpremVmwareAdminClusterNetworkConfig)
	PodAddressCidrBlocks() *[]*string
	SetPodAddressCidrBlocks(val *[]*string)
	PodAddressCidrBlocksInput() *[]*string
	ServiceAddressCidrBlocks() *[]*string
	SetServiceAddressCidrBlocks(val *[]*string)
	ServiceAddressCidrBlocksInput() *[]*string
	StaticIpConfig() GkeonpremVmwareAdminClusterNetworkConfigStaticIpConfigOutputReference
	StaticIpConfigInput() *GkeonpremVmwareAdminClusterNetworkConfigStaticIpConfig
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VcenterNetwork() *string
	SetVcenterNetwork(val *string)
	VcenterNetworkInput() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutDhcpIpConfig(value *GkeonpremVmwareAdminClusterNetworkConfigDhcpIpConfig)
	PutHaControlPlaneConfig(value *GkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfig)
	PutHostConfig(value *GkeonpremVmwareAdminClusterNetworkConfigHostConfig)
	PutStaticIpConfig(value *GkeonpremVmwareAdminClusterNetworkConfigStaticIpConfig)
	ResetDhcpIpConfig()
	ResetHaControlPlaneConfig()
	ResetHostConfig()
	ResetStaticIpConfig()
	ResetVcenterNetwork()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GkeonpremVmwareAdminClusterNetworkConfigOutputReference
type jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference) DhcpIpConfig() GkeonpremVmwareAdminClusterNetworkConfigDhcpIpConfigOutputReference {
	var returns GkeonpremVmwareAdminClusterNetworkConfigDhcpIpConfigOutputReference
	_jsii_.Get(
		j,
		"dhcpIpConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference) DhcpIpConfigInput() *GkeonpremVmwareAdminClusterNetworkConfigDhcpIpConfig {
	var returns *GkeonpremVmwareAdminClusterNetworkConfigDhcpIpConfig
	_jsii_.Get(
		j,
		"dhcpIpConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference) HaControlPlaneConfig() GkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigOutputReference {
	var returns GkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfigOutputReference
	_jsii_.Get(
		j,
		"haControlPlaneConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference) HaControlPlaneConfigInput() *GkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfig {
	var returns *GkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfig
	_jsii_.Get(
		j,
		"haControlPlaneConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference) HostConfig() GkeonpremVmwareAdminClusterNetworkConfigHostConfigOutputReference {
	var returns GkeonpremVmwareAdminClusterNetworkConfigHostConfigOutputReference
	_jsii_.Get(
		j,
		"hostConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference) HostConfigInput() *GkeonpremVmwareAdminClusterNetworkConfigHostConfig {
	var returns *GkeonpremVmwareAdminClusterNetworkConfigHostConfig
	_jsii_.Get(
		j,
		"hostConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference) InternalValue() *GkeonpremVmwareAdminClusterNetworkConfig {
	var returns *GkeonpremVmwareAdminClusterNetworkConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference) PodAddressCidrBlocks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"podAddressCidrBlocks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference) PodAddressCidrBlocksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"podAddressCidrBlocksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference) ServiceAddressCidrBlocks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serviceAddressCidrBlocks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference) ServiceAddressCidrBlocksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serviceAddressCidrBlocksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference) StaticIpConfig() GkeonpremVmwareAdminClusterNetworkConfigStaticIpConfigOutputReference {
	var returns GkeonpremVmwareAdminClusterNetworkConfigStaticIpConfigOutputReference
	_jsii_.Get(
		j,
		"staticIpConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference) StaticIpConfigInput() *GkeonpremVmwareAdminClusterNetworkConfigStaticIpConfig {
	var returns *GkeonpremVmwareAdminClusterNetworkConfigStaticIpConfig
	_jsii_.Get(
		j,
		"staticIpConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference) VcenterNetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vcenterNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference) VcenterNetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vcenterNetworkInput",
		&returns,
	)
	return returns
}


func NewGkeonpremVmwareAdminClusterNetworkConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GkeonpremVmwareAdminClusterNetworkConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGkeonpremVmwareAdminClusterNetworkConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.gkeonpremVmwareAdminCluster.GkeonpremVmwareAdminClusterNetworkConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGkeonpremVmwareAdminClusterNetworkConfigOutputReference_Override(g GkeonpremVmwareAdminClusterNetworkConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.gkeonpremVmwareAdminCluster.GkeonpremVmwareAdminClusterNetworkConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference)SetInternalValue(val *GkeonpremVmwareAdminClusterNetworkConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference)SetPodAddressCidrBlocks(val *[]*string) {
	if err := j.validateSetPodAddressCidrBlocksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"podAddressCidrBlocks",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference)SetServiceAddressCidrBlocks(val *[]*string) {
	if err := j.validateSetServiceAddressCidrBlocksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAddressCidrBlocks",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference)SetVcenterNetwork(val *string) {
	if err := j.validateSetVcenterNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vcenterNetwork",
		val,
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference) PutDhcpIpConfig(value *GkeonpremVmwareAdminClusterNetworkConfigDhcpIpConfig) {
	if err := g.validatePutDhcpIpConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDhcpIpConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference) PutHaControlPlaneConfig(value *GkeonpremVmwareAdminClusterNetworkConfigHaControlPlaneConfig) {
	if err := g.validatePutHaControlPlaneConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHaControlPlaneConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference) PutHostConfig(value *GkeonpremVmwareAdminClusterNetworkConfigHostConfig) {
	if err := g.validatePutHostConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHostConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference) PutStaticIpConfig(value *GkeonpremVmwareAdminClusterNetworkConfigStaticIpConfig) {
	if err := g.validatePutStaticIpConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStaticIpConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference) ResetDhcpIpConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetDhcpIpConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference) ResetHaControlPlaneConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetHaControlPlaneConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference) ResetHostConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetHostConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference) ResetStaticIpConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetStaticIpConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference) ResetVcenterNetwork() {
	_jsii_.InvokeVoid(
		g,
		"resetVcenterNetwork",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareAdminClusterNetworkConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

