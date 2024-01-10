// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonprembaremetalcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v13/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v13/gkeonprembaremetalcluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GkeonpremBareMetalClusterNetworkConfigOutputReference interface {
	cdktf.ComplexObject
	AdvancedNetworking() interface{}
	SetAdvancedNetworking(val interface{})
	AdvancedNetworkingInput() interface{}
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
	InternalValue() *GkeonpremBareMetalClusterNetworkConfig
	SetInternalValue(val *GkeonpremBareMetalClusterNetworkConfig)
	IslandModeCidr() GkeonpremBareMetalClusterNetworkConfigIslandModeCidrOutputReference
	IslandModeCidrInput() *GkeonpremBareMetalClusterNetworkConfigIslandModeCidr
	MultipleNetworkInterfacesConfig() GkeonpremBareMetalClusterNetworkConfigMultipleNetworkInterfacesConfigOutputReference
	MultipleNetworkInterfacesConfigInput() *GkeonpremBareMetalClusterNetworkConfigMultipleNetworkInterfacesConfig
	SrIovConfig() GkeonpremBareMetalClusterNetworkConfigSrIovConfigOutputReference
	SrIovConfigInput() *GkeonpremBareMetalClusterNetworkConfigSrIovConfig
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
	PutIslandModeCidr(value *GkeonpremBareMetalClusterNetworkConfigIslandModeCidr)
	PutMultipleNetworkInterfacesConfig(value *GkeonpremBareMetalClusterNetworkConfigMultipleNetworkInterfacesConfig)
	PutSrIovConfig(value *GkeonpremBareMetalClusterNetworkConfigSrIovConfig)
	ResetAdvancedNetworking()
	ResetIslandModeCidr()
	ResetMultipleNetworkInterfacesConfig()
	ResetSrIovConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GkeonpremBareMetalClusterNetworkConfigOutputReference
type jsiiProxy_GkeonpremBareMetalClusterNetworkConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GkeonpremBareMetalClusterNetworkConfigOutputReference) AdvancedNetworking() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"advancedNetworking",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterNetworkConfigOutputReference) AdvancedNetworkingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"advancedNetworkingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterNetworkConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterNetworkConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterNetworkConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterNetworkConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterNetworkConfigOutputReference) InternalValue() *GkeonpremBareMetalClusterNetworkConfig {
	var returns *GkeonpremBareMetalClusterNetworkConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterNetworkConfigOutputReference) IslandModeCidr() GkeonpremBareMetalClusterNetworkConfigIslandModeCidrOutputReference {
	var returns GkeonpremBareMetalClusterNetworkConfigIslandModeCidrOutputReference
	_jsii_.Get(
		j,
		"islandModeCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterNetworkConfigOutputReference) IslandModeCidrInput() *GkeonpremBareMetalClusterNetworkConfigIslandModeCidr {
	var returns *GkeonpremBareMetalClusterNetworkConfigIslandModeCidr
	_jsii_.Get(
		j,
		"islandModeCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterNetworkConfigOutputReference) MultipleNetworkInterfacesConfig() GkeonpremBareMetalClusterNetworkConfigMultipleNetworkInterfacesConfigOutputReference {
	var returns GkeonpremBareMetalClusterNetworkConfigMultipleNetworkInterfacesConfigOutputReference
	_jsii_.Get(
		j,
		"multipleNetworkInterfacesConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterNetworkConfigOutputReference) MultipleNetworkInterfacesConfigInput() *GkeonpremBareMetalClusterNetworkConfigMultipleNetworkInterfacesConfig {
	var returns *GkeonpremBareMetalClusterNetworkConfigMultipleNetworkInterfacesConfig
	_jsii_.Get(
		j,
		"multipleNetworkInterfacesConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterNetworkConfigOutputReference) SrIovConfig() GkeonpremBareMetalClusterNetworkConfigSrIovConfigOutputReference {
	var returns GkeonpremBareMetalClusterNetworkConfigSrIovConfigOutputReference
	_jsii_.Get(
		j,
		"srIovConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterNetworkConfigOutputReference) SrIovConfigInput() *GkeonpremBareMetalClusterNetworkConfigSrIovConfig {
	var returns *GkeonpremBareMetalClusterNetworkConfigSrIovConfig
	_jsii_.Get(
		j,
		"srIovConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterNetworkConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterNetworkConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGkeonpremBareMetalClusterNetworkConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GkeonpremBareMetalClusterNetworkConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGkeonpremBareMetalClusterNetworkConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GkeonpremBareMetalClusterNetworkConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.gkeonpremBareMetalCluster.GkeonpremBareMetalClusterNetworkConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGkeonpremBareMetalClusterNetworkConfigOutputReference_Override(g GkeonpremBareMetalClusterNetworkConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.gkeonpremBareMetalCluster.GkeonpremBareMetalClusterNetworkConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalClusterNetworkConfigOutputReference)SetAdvancedNetworking(val interface{}) {
	if err := j.validateSetAdvancedNetworkingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"advancedNetworking",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalClusterNetworkConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalClusterNetworkConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalClusterNetworkConfigOutputReference)SetInternalValue(val *GkeonpremBareMetalClusterNetworkConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalClusterNetworkConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalClusterNetworkConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GkeonpremBareMetalClusterNetworkConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalClusterNetworkConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GkeonpremBareMetalClusterNetworkConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GkeonpremBareMetalClusterNetworkConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GkeonpremBareMetalClusterNetworkConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GkeonpremBareMetalClusterNetworkConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GkeonpremBareMetalClusterNetworkConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GkeonpremBareMetalClusterNetworkConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GkeonpremBareMetalClusterNetworkConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GkeonpremBareMetalClusterNetworkConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GkeonpremBareMetalClusterNetworkConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalClusterNetworkConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GkeonpremBareMetalClusterNetworkConfigOutputReference) PutIslandModeCidr(value *GkeonpremBareMetalClusterNetworkConfigIslandModeCidr) {
	if err := g.validatePutIslandModeCidrParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putIslandModeCidr",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalClusterNetworkConfigOutputReference) PutMultipleNetworkInterfacesConfig(value *GkeonpremBareMetalClusterNetworkConfigMultipleNetworkInterfacesConfig) {
	if err := g.validatePutMultipleNetworkInterfacesConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMultipleNetworkInterfacesConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalClusterNetworkConfigOutputReference) PutSrIovConfig(value *GkeonpremBareMetalClusterNetworkConfigSrIovConfig) {
	if err := g.validatePutSrIovConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSrIovConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalClusterNetworkConfigOutputReference) ResetAdvancedNetworking() {
	_jsii_.InvokeVoid(
		g,
		"resetAdvancedNetworking",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalClusterNetworkConfigOutputReference) ResetIslandModeCidr() {
	_jsii_.InvokeVoid(
		g,
		"resetIslandModeCidr",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalClusterNetworkConfigOutputReference) ResetMultipleNetworkInterfacesConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetMultipleNetworkInterfacesConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalClusterNetworkConfigOutputReference) ResetSrIovConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetSrIovConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalClusterNetworkConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GkeonpremBareMetalClusterNetworkConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

