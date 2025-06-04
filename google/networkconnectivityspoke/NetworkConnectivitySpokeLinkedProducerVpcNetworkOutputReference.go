// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkconnectivityspoke

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/networkconnectivityspoke/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference interface {
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
	ExcludeExportRanges() *[]*string
	SetExcludeExportRanges(val *[]*string)
	ExcludeExportRangesInput() *[]*string
	// Experimental.
	Fqn() *string
	IncludeExportRanges() *[]*string
	SetIncludeExportRanges(val *[]*string)
	IncludeExportRangesInput() *[]*string
	InternalValue() *NetworkConnectivitySpokeLinkedProducerVpcNetwork
	SetInternalValue(val *NetworkConnectivitySpokeLinkedProducerVpcNetwork)
	Network() *string
	SetNetwork(val *string)
	NetworkInput() *string
	Peering() *string
	SetPeering(val *string)
	PeeringInput() *string
	ProducerNetwork() *string
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
	ResetExcludeExportRanges()
	ResetIncludeExportRanges()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference
type jsiiProxy_NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) ExcludeExportRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeExportRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) ExcludeExportRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeExportRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) IncludeExportRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeExportRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) IncludeExportRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeExportRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) InternalValue() *NetworkConnectivitySpokeLinkedProducerVpcNetwork {
	var returns *NetworkConnectivitySpokeLinkedProducerVpcNetwork
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) Peering() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peering",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) PeeringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peeringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) ProducerNetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"producerNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.networkConnectivitySpoke.NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference_Override(n NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.networkConnectivitySpoke.NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference)SetExcludeExportRanges(val *[]*string) {
	if err := j.validateSetExcludeExportRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeExportRanges",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference)SetIncludeExportRanges(val *[]*string) {
	if err := j.validateSetIncludeExportRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeExportRanges",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference)SetInternalValue(val *NetworkConnectivitySpokeLinkedProducerVpcNetwork) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference)SetPeering(val *string) {
	if err := j.validateSetPeeringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peering",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) ResetExcludeExportRanges() {
	_jsii_.InvokeVoid(
		n,
		"resetExcludeExportRanges",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) ResetIncludeExportRanges() {
	_jsii_.InvokeVoid(
		n,
		"resetIncludeExportRanges",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (n *jsiiProxy_NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

