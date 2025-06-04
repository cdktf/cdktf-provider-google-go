// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonprembaremetalcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/gkeonprembaremetalcluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference interface {
	cdktf.ComplexObject
	AddressPools() GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsList
	AddressPoolsInput() interface{}
	Asn() *float64
	SetAsn(val *float64)
	AsnInput() *float64
	BgpPeerConfigs() GkeonpremBareMetalClusterLoadBalancerBgpLbConfigBgpPeerConfigsList
	BgpPeerConfigsInput() interface{}
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
	InternalValue() *GkeonpremBareMetalClusterLoadBalancerBgpLbConfig
	SetInternalValue(val *GkeonpremBareMetalClusterLoadBalancerBgpLbConfig)
	LoadBalancerNodePoolConfig() GkeonpremBareMetalClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigOutputReference
	LoadBalancerNodePoolConfigInput() *GkeonpremBareMetalClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfig
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
	PutAddressPools(value interface{})
	PutBgpPeerConfigs(value interface{})
	PutLoadBalancerNodePoolConfig(value *GkeonpremBareMetalClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfig)
	ResetLoadBalancerNodePoolConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference
type jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) AddressPools() GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsList {
	var returns GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsList
	_jsii_.Get(
		j,
		"addressPools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) AddressPoolsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"addressPoolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) Asn() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"asn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) AsnInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"asnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) BgpPeerConfigs() GkeonpremBareMetalClusterLoadBalancerBgpLbConfigBgpPeerConfigsList {
	var returns GkeonpremBareMetalClusterLoadBalancerBgpLbConfigBgpPeerConfigsList
	_jsii_.Get(
		j,
		"bgpPeerConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) BgpPeerConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bgpPeerConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) InternalValue() *GkeonpremBareMetalClusterLoadBalancerBgpLbConfig {
	var returns *GkeonpremBareMetalClusterLoadBalancerBgpLbConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) LoadBalancerNodePoolConfig() GkeonpremBareMetalClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigOutputReference {
	var returns GkeonpremBareMetalClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigOutputReference
	_jsii_.Get(
		j,
		"loadBalancerNodePoolConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) LoadBalancerNodePoolConfigInput() *GkeonpremBareMetalClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfig {
	var returns *GkeonpremBareMetalClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfig
	_jsii_.Get(
		j,
		"loadBalancerNodePoolConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.gkeonpremBareMetalCluster.GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference_Override(g GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.gkeonpremBareMetalCluster.GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference)SetAsn(val *float64) {
	if err := j.validateSetAsnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"asn",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference)SetInternalValue(val *GkeonpremBareMetalClusterLoadBalancerBgpLbConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) PutAddressPools(value interface{}) {
	if err := g.validatePutAddressPoolsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAddressPools",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) PutBgpPeerConfigs(value interface{}) {
	if err := g.validatePutBgpPeerConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBgpPeerConfigs",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) PutLoadBalancerNodePoolConfig(value *GkeonpremBareMetalClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfig) {
	if err := g.validatePutLoadBalancerNodePoolConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLoadBalancerNodePoolConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) ResetLoadBalancerNodePoolConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetLoadBalancerNodePoolConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

