// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package edgecontainercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v12/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v12/edgecontainercluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EdgecontainerClusterNetworkingOutputReference interface {
	cdktf.ComplexObject
	ClusterIpv4CidrBlocks() *[]*string
	SetClusterIpv4CidrBlocks(val *[]*string)
	ClusterIpv4CidrBlocksInput() *[]*string
	ClusterIpv6CidrBlocks() *[]*string
	SetClusterIpv6CidrBlocks(val *[]*string)
	ClusterIpv6CidrBlocksInput() *[]*string
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
	InternalValue() *EdgecontainerClusterNetworking
	SetInternalValue(val *EdgecontainerClusterNetworking)
	NetworkType() *string
	ServicesIpv4CidrBlocks() *[]*string
	SetServicesIpv4CidrBlocks(val *[]*string)
	ServicesIpv4CidrBlocksInput() *[]*string
	ServicesIpv6CidrBlocks() *[]*string
	SetServicesIpv6CidrBlocks(val *[]*string)
	ServicesIpv6CidrBlocksInput() *[]*string
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
	ResetClusterIpv6CidrBlocks()
	ResetServicesIpv6CidrBlocks()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EdgecontainerClusterNetworkingOutputReference
type jsiiProxy_EdgecontainerClusterNetworkingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EdgecontainerClusterNetworkingOutputReference) ClusterIpv4CidrBlocks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clusterIpv4CidrBlocks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerClusterNetworkingOutputReference) ClusterIpv4CidrBlocksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clusterIpv4CidrBlocksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerClusterNetworkingOutputReference) ClusterIpv6CidrBlocks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clusterIpv6CidrBlocks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerClusterNetworkingOutputReference) ClusterIpv6CidrBlocksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clusterIpv6CidrBlocksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerClusterNetworkingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerClusterNetworkingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerClusterNetworkingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerClusterNetworkingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerClusterNetworkingOutputReference) InternalValue() *EdgecontainerClusterNetworking {
	var returns *EdgecontainerClusterNetworking
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerClusterNetworkingOutputReference) NetworkType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerClusterNetworkingOutputReference) ServicesIpv4CidrBlocks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"servicesIpv4CidrBlocks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerClusterNetworkingOutputReference) ServicesIpv4CidrBlocksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"servicesIpv4CidrBlocksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerClusterNetworkingOutputReference) ServicesIpv6CidrBlocks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"servicesIpv6CidrBlocks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerClusterNetworkingOutputReference) ServicesIpv6CidrBlocksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"servicesIpv6CidrBlocksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerClusterNetworkingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerClusterNetworkingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEdgecontainerClusterNetworkingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EdgecontainerClusterNetworkingOutputReference {
	_init_.Initialize()

	if err := validateNewEdgecontainerClusterNetworkingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EdgecontainerClusterNetworkingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.edgecontainerCluster.EdgecontainerClusterNetworkingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEdgecontainerClusterNetworkingOutputReference_Override(e EdgecontainerClusterNetworkingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.edgecontainerCluster.EdgecontainerClusterNetworkingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EdgecontainerClusterNetworkingOutputReference)SetClusterIpv4CidrBlocks(val *[]*string) {
	if err := j.validateSetClusterIpv4CidrBlocksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterIpv4CidrBlocks",
		val,
	)
}

func (j *jsiiProxy_EdgecontainerClusterNetworkingOutputReference)SetClusterIpv6CidrBlocks(val *[]*string) {
	if err := j.validateSetClusterIpv6CidrBlocksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterIpv6CidrBlocks",
		val,
	)
}

func (j *jsiiProxy_EdgecontainerClusterNetworkingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EdgecontainerClusterNetworkingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EdgecontainerClusterNetworkingOutputReference)SetInternalValue(val *EdgecontainerClusterNetworking) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EdgecontainerClusterNetworkingOutputReference)SetServicesIpv4CidrBlocks(val *[]*string) {
	if err := j.validateSetServicesIpv4CidrBlocksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"servicesIpv4CidrBlocks",
		val,
	)
}

func (j *jsiiProxy_EdgecontainerClusterNetworkingOutputReference)SetServicesIpv6CidrBlocks(val *[]*string) {
	if err := j.validateSetServicesIpv6CidrBlocksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"servicesIpv6CidrBlocks",
		val,
	)
}

func (j *jsiiProxy_EdgecontainerClusterNetworkingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EdgecontainerClusterNetworkingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EdgecontainerClusterNetworkingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgecontainerClusterNetworkingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EdgecontainerClusterNetworkingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EdgecontainerClusterNetworkingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EdgecontainerClusterNetworkingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EdgecontainerClusterNetworkingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EdgecontainerClusterNetworkingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EdgecontainerClusterNetworkingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EdgecontainerClusterNetworkingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EdgecontainerClusterNetworkingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EdgecontainerClusterNetworkingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgecontainerClusterNetworkingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EdgecontainerClusterNetworkingOutputReference) ResetClusterIpv6CidrBlocks() {
	_jsii_.InvokeVoid(
		e,
		"resetClusterIpv6CidrBlocks",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgecontainerClusterNetworkingOutputReference) ResetServicesIpv6CidrBlocks() {
	_jsii_.InvokeVoid(
		e,
		"resetServicesIpv6CidrBlocks",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgecontainerClusterNetworkingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_EdgecontainerClusterNetworkingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

