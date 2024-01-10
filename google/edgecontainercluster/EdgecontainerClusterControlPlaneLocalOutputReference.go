// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package edgecontainercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v13/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v13/edgecontainercluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EdgecontainerClusterControlPlaneLocalOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() *EdgecontainerClusterControlPlaneLocal
	SetInternalValue(val *EdgecontainerClusterControlPlaneLocal)
	MachineFilter() *string
	SetMachineFilter(val *string)
	MachineFilterInput() *string
	NodeCount() *float64
	SetNodeCount(val *float64)
	NodeCountInput() *float64
	NodeLocation() *string
	SetNodeLocation(val *string)
	NodeLocationInput() *string
	SharedDeploymentPolicy() *string
	SetSharedDeploymentPolicy(val *string)
	SharedDeploymentPolicyInput() *string
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
	ResetMachineFilter()
	ResetNodeCount()
	ResetNodeLocation()
	ResetSharedDeploymentPolicy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EdgecontainerClusterControlPlaneLocalOutputReference
type jsiiProxy_EdgecontainerClusterControlPlaneLocalOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EdgecontainerClusterControlPlaneLocalOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerClusterControlPlaneLocalOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerClusterControlPlaneLocalOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerClusterControlPlaneLocalOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerClusterControlPlaneLocalOutputReference) InternalValue() *EdgecontainerClusterControlPlaneLocal {
	var returns *EdgecontainerClusterControlPlaneLocal
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerClusterControlPlaneLocalOutputReference) MachineFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerClusterControlPlaneLocalOutputReference) MachineFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerClusterControlPlaneLocalOutputReference) NodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerClusterControlPlaneLocalOutputReference) NodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerClusterControlPlaneLocalOutputReference) NodeLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerClusterControlPlaneLocalOutputReference) NodeLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerClusterControlPlaneLocalOutputReference) SharedDeploymentPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedDeploymentPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerClusterControlPlaneLocalOutputReference) SharedDeploymentPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharedDeploymentPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerClusterControlPlaneLocalOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerClusterControlPlaneLocalOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEdgecontainerClusterControlPlaneLocalOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EdgecontainerClusterControlPlaneLocalOutputReference {
	_init_.Initialize()

	if err := validateNewEdgecontainerClusterControlPlaneLocalOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EdgecontainerClusterControlPlaneLocalOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.edgecontainerCluster.EdgecontainerClusterControlPlaneLocalOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEdgecontainerClusterControlPlaneLocalOutputReference_Override(e EdgecontainerClusterControlPlaneLocalOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.edgecontainerCluster.EdgecontainerClusterControlPlaneLocalOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EdgecontainerClusterControlPlaneLocalOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EdgecontainerClusterControlPlaneLocalOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EdgecontainerClusterControlPlaneLocalOutputReference)SetInternalValue(val *EdgecontainerClusterControlPlaneLocal) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EdgecontainerClusterControlPlaneLocalOutputReference)SetMachineFilter(val *string) {
	if err := j.validateSetMachineFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"machineFilter",
		val,
	)
}

func (j *jsiiProxy_EdgecontainerClusterControlPlaneLocalOutputReference)SetNodeCount(val *float64) {
	if err := j.validateSetNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeCount",
		val,
	)
}

func (j *jsiiProxy_EdgecontainerClusterControlPlaneLocalOutputReference)SetNodeLocation(val *string) {
	if err := j.validateSetNodeLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeLocation",
		val,
	)
}

func (j *jsiiProxy_EdgecontainerClusterControlPlaneLocalOutputReference)SetSharedDeploymentPolicy(val *string) {
	if err := j.validateSetSharedDeploymentPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedDeploymentPolicy",
		val,
	)
}

func (j *jsiiProxy_EdgecontainerClusterControlPlaneLocalOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EdgecontainerClusterControlPlaneLocalOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EdgecontainerClusterControlPlaneLocalOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgecontainerClusterControlPlaneLocalOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgecontainerClusterControlPlaneLocalOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgecontainerClusterControlPlaneLocalOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgecontainerClusterControlPlaneLocalOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgecontainerClusterControlPlaneLocalOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgecontainerClusterControlPlaneLocalOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgecontainerClusterControlPlaneLocalOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgecontainerClusterControlPlaneLocalOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgecontainerClusterControlPlaneLocalOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgecontainerClusterControlPlaneLocalOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgecontainerClusterControlPlaneLocalOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgecontainerClusterControlPlaneLocalOutputReference) ResetMachineFilter() {
	_jsii_.InvokeVoid(
		e,
		"resetMachineFilter",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgecontainerClusterControlPlaneLocalOutputReference) ResetNodeCount() {
	_jsii_.InvokeVoid(
		e,
		"resetNodeCount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgecontainerClusterControlPlaneLocalOutputReference) ResetNodeLocation() {
	_jsii_.InvokeVoid(
		e,
		"resetNodeLocation",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgecontainerClusterControlPlaneLocalOutputReference) ResetSharedDeploymentPolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetSharedDeploymentPolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgecontainerClusterControlPlaneLocalOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgecontainerClusterControlPlaneLocalOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

