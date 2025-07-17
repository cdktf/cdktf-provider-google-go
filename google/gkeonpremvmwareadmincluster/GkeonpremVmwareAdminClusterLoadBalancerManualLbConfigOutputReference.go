// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonpremvmwareadmincluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/gkeonpremvmwareadmincluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference interface {
	cdktf.ComplexObject
	AddonsNodePort() *float64
	SetAddonsNodePort(val *float64)
	AddonsNodePortInput() *float64
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
	ControlPlaneNodePort() *float64
	SetControlPlaneNodePort(val *float64)
	ControlPlaneNodePortInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	IngressHttpNodePort() *float64
	SetIngressHttpNodePort(val *float64)
	IngressHttpNodePortInput() *float64
	IngressHttpsNodePort() *float64
	SetIngressHttpsNodePort(val *float64)
	IngressHttpsNodePortInput() *float64
	InternalValue() *GkeonpremVmwareAdminClusterLoadBalancerManualLbConfig
	SetInternalValue(val *GkeonpremVmwareAdminClusterLoadBalancerManualLbConfig)
	KonnectivityServerNodePort() *float64
	SetKonnectivityServerNodePort(val *float64)
	KonnectivityServerNodePortInput() *float64
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
	ResetAddonsNodePort()
	ResetControlPlaneNodePort()
	ResetIngressHttpNodePort()
	ResetIngressHttpsNodePort()
	ResetKonnectivityServerNodePort()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference
type jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) AddonsNodePort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"addonsNodePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) AddonsNodePortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"addonsNodePortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) ControlPlaneNodePort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"controlPlaneNodePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) ControlPlaneNodePortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"controlPlaneNodePortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) IngressHttpNodePort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ingressHttpNodePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) IngressHttpNodePortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ingressHttpNodePortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) IngressHttpsNodePort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ingressHttpsNodePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) IngressHttpsNodePortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ingressHttpsNodePortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) InternalValue() *GkeonpremVmwareAdminClusterLoadBalancerManualLbConfig {
	var returns *GkeonpremVmwareAdminClusterLoadBalancerManualLbConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) KonnectivityServerNodePort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"konnectivityServerNodePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) KonnectivityServerNodePortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"konnectivityServerNodePortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.gkeonpremVmwareAdminCluster.GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference_Override(g GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.gkeonpremVmwareAdminCluster.GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference)SetAddonsNodePort(val *float64) {
	if err := j.validateSetAddonsNodePortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addonsNodePort",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference)SetControlPlaneNodePort(val *float64) {
	if err := j.validateSetControlPlaneNodePortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"controlPlaneNodePort",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference)SetIngressHttpNodePort(val *float64) {
	if err := j.validateSetIngressHttpNodePortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ingressHttpNodePort",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference)SetIngressHttpsNodePort(val *float64) {
	if err := j.validateSetIngressHttpsNodePortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ingressHttpsNodePort",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference)SetInternalValue(val *GkeonpremVmwareAdminClusterLoadBalancerManualLbConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference)SetKonnectivityServerNodePort(val *float64) {
	if err := j.validateSetKonnectivityServerNodePortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"konnectivityServerNodePort",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) ResetAddonsNodePort() {
	_jsii_.InvokeVoid(
		g,
		"resetAddonsNodePort",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) ResetControlPlaneNodePort() {
	_jsii_.InvokeVoid(
		g,
		"resetControlPlaneNodePort",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) ResetIngressHttpNodePort() {
	_jsii_.InvokeVoid(
		g,
		"resetIngressHttpNodePort",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) ResetIngressHttpsNodePort() {
	_jsii_.InvokeVoid(
		g,
		"resetIngressHttpsNodePort",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) ResetKonnectivityServerNodePort() {
	_jsii_.InvokeVoid(
		g,
		"resetKonnectivityServerNodePort",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GkeonpremVmwareAdminClusterLoadBalancerManualLbConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

