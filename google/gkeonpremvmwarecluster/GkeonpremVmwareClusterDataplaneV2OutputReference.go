// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonpremvmwarecluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/gkeonpremvmwarecluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GkeonpremVmwareClusterDataplaneV2OutputReference interface {
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
	DataplaneV2Enabled() interface{}
	SetDataplaneV2Enabled(val interface{})
	DataplaneV2EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GkeonpremVmwareClusterDataplaneV2
	SetInternalValue(val *GkeonpremVmwareClusterDataplaneV2)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WindowsDataplaneV2Enabled() interface{}
	SetWindowsDataplaneV2Enabled(val interface{})
	WindowsDataplaneV2EnabledInput() interface{}
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
	ResetAdvancedNetworking()
	ResetDataplaneV2Enabled()
	ResetWindowsDataplaneV2Enabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GkeonpremVmwareClusterDataplaneV2OutputReference
type jsiiProxy_GkeonpremVmwareClusterDataplaneV2OutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GkeonpremVmwareClusterDataplaneV2OutputReference) AdvancedNetworking() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"advancedNetworking",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterDataplaneV2OutputReference) AdvancedNetworkingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"advancedNetworkingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterDataplaneV2OutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterDataplaneV2OutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterDataplaneV2OutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterDataplaneV2OutputReference) DataplaneV2Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataplaneV2Enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterDataplaneV2OutputReference) DataplaneV2EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataplaneV2EnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterDataplaneV2OutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterDataplaneV2OutputReference) InternalValue() *GkeonpremVmwareClusterDataplaneV2 {
	var returns *GkeonpremVmwareClusterDataplaneV2
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterDataplaneV2OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterDataplaneV2OutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterDataplaneV2OutputReference) WindowsDataplaneV2Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"windowsDataplaneV2Enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareClusterDataplaneV2OutputReference) WindowsDataplaneV2EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"windowsDataplaneV2EnabledInput",
		&returns,
	)
	return returns
}


func NewGkeonpremVmwareClusterDataplaneV2OutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GkeonpremVmwareClusterDataplaneV2OutputReference {
	_init_.Initialize()

	if err := validateNewGkeonpremVmwareClusterDataplaneV2OutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GkeonpremVmwareClusterDataplaneV2OutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.gkeonpremVmwareCluster.GkeonpremVmwareClusterDataplaneV2OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGkeonpremVmwareClusterDataplaneV2OutputReference_Override(g GkeonpremVmwareClusterDataplaneV2OutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.gkeonpremVmwareCluster.GkeonpremVmwareClusterDataplaneV2OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GkeonpremVmwareClusterDataplaneV2OutputReference)SetAdvancedNetworking(val interface{}) {
	if err := j.validateSetAdvancedNetworkingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"advancedNetworking",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareClusterDataplaneV2OutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareClusterDataplaneV2OutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareClusterDataplaneV2OutputReference)SetDataplaneV2Enabled(val interface{}) {
	if err := j.validateSetDataplaneV2EnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataplaneV2Enabled",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareClusterDataplaneV2OutputReference)SetInternalValue(val *GkeonpremVmwareClusterDataplaneV2) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareClusterDataplaneV2OutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareClusterDataplaneV2OutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareClusterDataplaneV2OutputReference)SetWindowsDataplaneV2Enabled(val interface{}) {
	if err := j.validateSetWindowsDataplaneV2EnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"windowsDataplaneV2Enabled",
		val,
	)
}

func (g *jsiiProxy_GkeonpremVmwareClusterDataplaneV2OutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareClusterDataplaneV2OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GkeonpremVmwareClusterDataplaneV2OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GkeonpremVmwareClusterDataplaneV2OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GkeonpremVmwareClusterDataplaneV2OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GkeonpremVmwareClusterDataplaneV2OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GkeonpremVmwareClusterDataplaneV2OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GkeonpremVmwareClusterDataplaneV2OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GkeonpremVmwareClusterDataplaneV2OutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GkeonpremVmwareClusterDataplaneV2OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GkeonpremVmwareClusterDataplaneV2OutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareClusterDataplaneV2OutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GkeonpremVmwareClusterDataplaneV2OutputReference) ResetAdvancedNetworking() {
	_jsii_.InvokeVoid(
		g,
		"resetAdvancedNetworking",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareClusterDataplaneV2OutputReference) ResetDataplaneV2Enabled() {
	_jsii_.InvokeVoid(
		g,
		"resetDataplaneV2Enabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareClusterDataplaneV2OutputReference) ResetWindowsDataplaneV2Enabled() {
	_jsii_.InvokeVoid(
		g,
		"resetWindowsDataplaneV2Enabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareClusterDataplaneV2OutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GkeonpremVmwareClusterDataplaneV2OutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

