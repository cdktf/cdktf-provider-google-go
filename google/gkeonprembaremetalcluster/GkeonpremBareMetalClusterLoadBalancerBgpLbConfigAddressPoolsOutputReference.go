// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonprembaremetalcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v12/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v12/gkeonprembaremetalcluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference interface {
	cdktf.ComplexObject
	Addresses() *[]*string
	SetAddresses(val *[]*string)
	AddressesInput() *[]*string
	AvoidBuggyIps() interface{}
	SetAvoidBuggyIps(val interface{})
	AvoidBuggyIpsInput() interface{}
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ManualAssign() *string
	SetManualAssign(val *string)
	ManualAssignInput() *string
	Pool() *string
	SetPool(val *string)
	PoolInput() *string
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
	ResetAvoidBuggyIps()
	ResetManualAssign()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference
type jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference) Addresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference) AddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"addressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference) AvoidBuggyIps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"avoidBuggyIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference) AvoidBuggyIpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"avoidBuggyIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference) ManualAssign() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manualAssign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference) ManualAssignInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manualAssignInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference) Pool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference) PoolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"poolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference {
	_init_.Initialize()

	if err := validateNewGkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.gkeonpremBareMetalCluster.GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference_Override(g GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.gkeonpremBareMetalCluster.GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference)SetAddresses(val *[]*string) {
	if err := j.validateSetAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addresses",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference)SetAvoidBuggyIps(val interface{}) {
	if err := j.validateSetAvoidBuggyIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"avoidBuggyIps",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference)SetManualAssign(val *string) {
	if err := j.validateSetManualAssignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manualAssign",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference)SetPool(val *string) {
	if err := j.validateSetPoolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pool",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference) ResetAvoidBuggyIps() {
	_jsii_.InvokeVoid(
		g,
		"resetAvoidBuggyIps",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference) ResetManualAssign() {
	_jsii_.InvokeVoid(
		g,
		"resetManualAssign",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GkeonpremBareMetalClusterLoadBalancerBgpLbConfigAddressPoolsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

