// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonprembaremetalcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/gkeonprembaremetalcluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GkeonpremBareMetalClusterLoadBalancerOutputReference interface {
	cdktf.ComplexObject
	BgpLbConfig() GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference
	BgpLbConfigInput() *GkeonpremBareMetalClusterLoadBalancerBgpLbConfig
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
	InternalValue() *GkeonpremBareMetalClusterLoadBalancer
	SetInternalValue(val *GkeonpremBareMetalClusterLoadBalancer)
	ManualLbConfig() GkeonpremBareMetalClusterLoadBalancerManualLbConfigOutputReference
	ManualLbConfigInput() *GkeonpremBareMetalClusterLoadBalancerManualLbConfig
	MetalLbConfig() GkeonpremBareMetalClusterLoadBalancerMetalLbConfigOutputReference
	MetalLbConfigInput() *GkeonpremBareMetalClusterLoadBalancerMetalLbConfig
	PortConfig() GkeonpremBareMetalClusterLoadBalancerPortConfigOutputReference
	PortConfigInput() *GkeonpremBareMetalClusterLoadBalancerPortConfig
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VipConfig() GkeonpremBareMetalClusterLoadBalancerVipConfigOutputReference
	VipConfigInput() *GkeonpremBareMetalClusterLoadBalancerVipConfig
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
	PutBgpLbConfig(value *GkeonpremBareMetalClusterLoadBalancerBgpLbConfig)
	PutManualLbConfig(value *GkeonpremBareMetalClusterLoadBalancerManualLbConfig)
	PutMetalLbConfig(value *GkeonpremBareMetalClusterLoadBalancerMetalLbConfig)
	PutPortConfig(value *GkeonpremBareMetalClusterLoadBalancerPortConfig)
	PutVipConfig(value *GkeonpremBareMetalClusterLoadBalancerVipConfig)
	ResetBgpLbConfig()
	ResetManualLbConfig()
	ResetMetalLbConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GkeonpremBareMetalClusterLoadBalancerOutputReference
type jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference) BgpLbConfig() GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference {
	var returns GkeonpremBareMetalClusterLoadBalancerBgpLbConfigOutputReference
	_jsii_.Get(
		j,
		"bgpLbConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference) BgpLbConfigInput() *GkeonpremBareMetalClusterLoadBalancerBgpLbConfig {
	var returns *GkeonpremBareMetalClusterLoadBalancerBgpLbConfig
	_jsii_.Get(
		j,
		"bgpLbConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference) InternalValue() *GkeonpremBareMetalClusterLoadBalancer {
	var returns *GkeonpremBareMetalClusterLoadBalancer
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference) ManualLbConfig() GkeonpremBareMetalClusterLoadBalancerManualLbConfigOutputReference {
	var returns GkeonpremBareMetalClusterLoadBalancerManualLbConfigOutputReference
	_jsii_.Get(
		j,
		"manualLbConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference) ManualLbConfigInput() *GkeonpremBareMetalClusterLoadBalancerManualLbConfig {
	var returns *GkeonpremBareMetalClusterLoadBalancerManualLbConfig
	_jsii_.Get(
		j,
		"manualLbConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference) MetalLbConfig() GkeonpremBareMetalClusterLoadBalancerMetalLbConfigOutputReference {
	var returns GkeonpremBareMetalClusterLoadBalancerMetalLbConfigOutputReference
	_jsii_.Get(
		j,
		"metalLbConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference) MetalLbConfigInput() *GkeonpremBareMetalClusterLoadBalancerMetalLbConfig {
	var returns *GkeonpremBareMetalClusterLoadBalancerMetalLbConfig
	_jsii_.Get(
		j,
		"metalLbConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference) PortConfig() GkeonpremBareMetalClusterLoadBalancerPortConfigOutputReference {
	var returns GkeonpremBareMetalClusterLoadBalancerPortConfigOutputReference
	_jsii_.Get(
		j,
		"portConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference) PortConfigInput() *GkeonpremBareMetalClusterLoadBalancerPortConfig {
	var returns *GkeonpremBareMetalClusterLoadBalancerPortConfig
	_jsii_.Get(
		j,
		"portConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference) VipConfig() GkeonpremBareMetalClusterLoadBalancerVipConfigOutputReference {
	var returns GkeonpremBareMetalClusterLoadBalancerVipConfigOutputReference
	_jsii_.Get(
		j,
		"vipConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference) VipConfigInput() *GkeonpremBareMetalClusterLoadBalancerVipConfig {
	var returns *GkeonpremBareMetalClusterLoadBalancerVipConfig
	_jsii_.Get(
		j,
		"vipConfigInput",
		&returns,
	)
	return returns
}


func NewGkeonpremBareMetalClusterLoadBalancerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GkeonpremBareMetalClusterLoadBalancerOutputReference {
	_init_.Initialize()

	if err := validateNewGkeonpremBareMetalClusterLoadBalancerOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.gkeonpremBareMetalCluster.GkeonpremBareMetalClusterLoadBalancerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGkeonpremBareMetalClusterLoadBalancerOutputReference_Override(g GkeonpremBareMetalClusterLoadBalancerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.gkeonpremBareMetalCluster.GkeonpremBareMetalClusterLoadBalancerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference)SetInternalValue(val *GkeonpremBareMetalClusterLoadBalancer) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference) PutBgpLbConfig(value *GkeonpremBareMetalClusterLoadBalancerBgpLbConfig) {
	if err := g.validatePutBgpLbConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBgpLbConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference) PutManualLbConfig(value *GkeonpremBareMetalClusterLoadBalancerManualLbConfig) {
	if err := g.validatePutManualLbConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putManualLbConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference) PutMetalLbConfig(value *GkeonpremBareMetalClusterLoadBalancerMetalLbConfig) {
	if err := g.validatePutMetalLbConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMetalLbConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference) PutPortConfig(value *GkeonpremBareMetalClusterLoadBalancerPortConfig) {
	if err := g.validatePutPortConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPortConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference) PutVipConfig(value *GkeonpremBareMetalClusterLoadBalancerVipConfig) {
	if err := g.validatePutVipConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVipConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference) ResetBgpLbConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetBgpLbConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference) ResetManualLbConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetManualLbConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference) ResetMetalLbConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetMetalLbConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

