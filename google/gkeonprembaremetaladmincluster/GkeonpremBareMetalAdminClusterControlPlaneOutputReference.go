// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonprembaremetaladmincluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/gkeonprembaremetaladmincluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GkeonpremBareMetalAdminClusterControlPlaneOutputReference interface {
	cdktf.ComplexObject
	ApiServerArgs() GkeonpremBareMetalAdminClusterControlPlaneApiServerArgsList
	ApiServerArgsInput() interface{}
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
	ControlPlaneNodePoolConfig() GkeonpremBareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigOutputReference
	ControlPlaneNodePoolConfigInput() *GkeonpremBareMetalAdminClusterControlPlaneControlPlaneNodePoolConfig
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GkeonpremBareMetalAdminClusterControlPlane
	SetInternalValue(val *GkeonpremBareMetalAdminClusterControlPlane)
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
	PutApiServerArgs(value interface{})
	PutControlPlaneNodePoolConfig(value *GkeonpremBareMetalAdminClusterControlPlaneControlPlaneNodePoolConfig)
	ResetApiServerArgs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GkeonpremBareMetalAdminClusterControlPlaneOutputReference
type jsiiProxy_GkeonpremBareMetalAdminClusterControlPlaneOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterControlPlaneOutputReference) ApiServerArgs() GkeonpremBareMetalAdminClusterControlPlaneApiServerArgsList {
	var returns GkeonpremBareMetalAdminClusterControlPlaneApiServerArgsList
	_jsii_.Get(
		j,
		"apiServerArgs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterControlPlaneOutputReference) ApiServerArgsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"apiServerArgsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterControlPlaneOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterControlPlaneOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterControlPlaneOutputReference) ControlPlaneNodePoolConfig() GkeonpremBareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigOutputReference {
	var returns GkeonpremBareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigOutputReference
	_jsii_.Get(
		j,
		"controlPlaneNodePoolConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterControlPlaneOutputReference) ControlPlaneNodePoolConfigInput() *GkeonpremBareMetalAdminClusterControlPlaneControlPlaneNodePoolConfig {
	var returns *GkeonpremBareMetalAdminClusterControlPlaneControlPlaneNodePoolConfig
	_jsii_.Get(
		j,
		"controlPlaneNodePoolConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterControlPlaneOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterControlPlaneOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterControlPlaneOutputReference) InternalValue() *GkeonpremBareMetalAdminClusterControlPlane {
	var returns *GkeonpremBareMetalAdminClusterControlPlane
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterControlPlaneOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterControlPlaneOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGkeonpremBareMetalAdminClusterControlPlaneOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GkeonpremBareMetalAdminClusterControlPlaneOutputReference {
	_init_.Initialize()

	if err := validateNewGkeonpremBareMetalAdminClusterControlPlaneOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GkeonpremBareMetalAdminClusterControlPlaneOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.gkeonpremBareMetalAdminCluster.GkeonpremBareMetalAdminClusterControlPlaneOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGkeonpremBareMetalAdminClusterControlPlaneOutputReference_Override(g GkeonpremBareMetalAdminClusterControlPlaneOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.gkeonpremBareMetalAdminCluster.GkeonpremBareMetalAdminClusterControlPlaneOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterControlPlaneOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterControlPlaneOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterControlPlaneOutputReference)SetInternalValue(val *GkeonpremBareMetalAdminClusterControlPlane) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterControlPlaneOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalAdminClusterControlPlaneOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterControlPlaneOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterControlPlaneOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterControlPlaneOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterControlPlaneOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterControlPlaneOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterControlPlaneOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterControlPlaneOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterControlPlaneOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterControlPlaneOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterControlPlaneOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterControlPlaneOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterControlPlaneOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterControlPlaneOutputReference) PutApiServerArgs(value interface{}) {
	if err := g.validatePutApiServerArgsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putApiServerArgs",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterControlPlaneOutputReference) PutControlPlaneNodePoolConfig(value *GkeonpremBareMetalAdminClusterControlPlaneControlPlaneNodePoolConfig) {
	if err := g.validatePutControlPlaneNodePoolConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putControlPlaneNodePoolConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterControlPlaneOutputReference) ResetApiServerArgs() {
	_jsii_.InvokeVoid(
		g,
		"resetApiServerArgs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterControlPlaneOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GkeonpremBareMetalAdminClusterControlPlaneOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

