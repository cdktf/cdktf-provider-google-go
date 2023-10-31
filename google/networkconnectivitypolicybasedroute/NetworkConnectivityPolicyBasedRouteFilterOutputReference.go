// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkconnectivitypolicybasedroute

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v12/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v12/networkconnectivitypolicybasedroute/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkConnectivityPolicyBasedRouteFilterOutputReference interface {
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
	DestRange() *string
	SetDestRange(val *string)
	DestRangeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *NetworkConnectivityPolicyBasedRouteFilter
	SetInternalValue(val *NetworkConnectivityPolicyBasedRouteFilter)
	IpProtocol() *string
	SetIpProtocol(val *string)
	IpProtocolInput() *string
	ProtocolVersion() *string
	SetProtocolVersion(val *string)
	ProtocolVersionInput() *string
	SrcRange() *string
	SetSrcRange(val *string)
	SrcRangeInput() *string
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
	ResetDestRange()
	ResetIpProtocol()
	ResetSrcRange()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkConnectivityPolicyBasedRouteFilterOutputReference
type jsiiProxy_NetworkConnectivityPolicyBasedRouteFilterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRouteFilterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRouteFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRouteFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRouteFilterOutputReference) DestRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRouteFilterOutputReference) DestRangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRouteFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRouteFilterOutputReference) InternalValue() *NetworkConnectivityPolicyBasedRouteFilter {
	var returns *NetworkConnectivityPolicyBasedRouteFilter
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRouteFilterOutputReference) IpProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRouteFilterOutputReference) IpProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRouteFilterOutputReference) ProtocolVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRouteFilterOutputReference) ProtocolVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRouteFilterOutputReference) SrcRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"srcRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRouteFilterOutputReference) SrcRangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"srcRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRouteFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRouteFilterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNetworkConnectivityPolicyBasedRouteFilterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NetworkConnectivityPolicyBasedRouteFilterOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkConnectivityPolicyBasedRouteFilterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkConnectivityPolicyBasedRouteFilterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.networkConnectivityPolicyBasedRoute.NetworkConnectivityPolicyBasedRouteFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetworkConnectivityPolicyBasedRouteFilterOutputReference_Override(n NetworkConnectivityPolicyBasedRouteFilterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.networkConnectivityPolicyBasedRoute.NetworkConnectivityPolicyBasedRouteFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRouteFilterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRouteFilterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRouteFilterOutputReference)SetDestRange(val *string) {
	if err := j.validateSetDestRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destRange",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRouteFilterOutputReference)SetInternalValue(val *NetworkConnectivityPolicyBasedRouteFilter) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRouteFilterOutputReference)SetIpProtocol(val *string) {
	if err := j.validateSetIpProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipProtocol",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRouteFilterOutputReference)SetProtocolVersion(val *string) {
	if err := j.validateSetProtocolVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocolVersion",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRouteFilterOutputReference)SetSrcRange(val *string) {
	if err := j.validateSetSrcRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"srcRange",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRouteFilterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRouteFilterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRouteFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRouteFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRouteFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRouteFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRouteFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRouteFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRouteFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRouteFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRouteFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRouteFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRouteFilterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRouteFilterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRouteFilterOutputReference) ResetDestRange() {
	_jsii_.InvokeVoid(
		n,
		"resetDestRange",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRouteFilterOutputReference) ResetIpProtocol() {
	_jsii_.InvokeVoid(
		n,
		"resetIpProtocol",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRouteFilterOutputReference) ResetSrcRange() {
	_jsii_.InvokeVoid(
		n,
		"resetSrcRange",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRouteFilterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := n.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRouteFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

