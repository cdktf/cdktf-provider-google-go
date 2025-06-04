// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package netappvolume

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/netappvolume/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetappVolumeHybridReplicationParametersOutputReference interface {
	cdktf.ComplexObject
	ClusterLocation() *string
	SetClusterLocation(val *string)
	ClusterLocationInput() *string
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *NetappVolumeHybridReplicationParameters
	SetInternalValue(val *NetappVolumeHybridReplicationParameters)
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	PeerClusterName() *string
	SetPeerClusterName(val *string)
	PeerClusterNameInput() *string
	PeerIpAddresses() *string
	SetPeerIpAddresses(val *string)
	PeerIpAddressesInput() *string
	PeerSvmName() *string
	SetPeerSvmName(val *string)
	PeerSvmNameInput() *string
	PeerVolumeName() *string
	SetPeerVolumeName(val *string)
	PeerVolumeNameInput() *string
	Replication() *string
	SetReplication(val *string)
	ReplicationInput() *string
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
	ResetClusterLocation()
	ResetDescription()
	ResetLabels()
	ResetPeerClusterName()
	ResetPeerIpAddresses()
	ResetPeerSvmName()
	ResetPeerVolumeName()
	ResetReplication()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetappVolumeHybridReplicationParametersOutputReference
type jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference) ClusterLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference) ClusterLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference) InternalValue() *NetappVolumeHybridReplicationParameters {
	var returns *NetappVolumeHybridReplicationParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference) PeerClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerClusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference) PeerClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerClusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference) PeerIpAddresses() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference) PeerIpAddressesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerIpAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference) PeerSvmName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerSvmName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference) PeerSvmNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerSvmNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference) PeerVolumeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerVolumeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference) PeerVolumeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerVolumeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference) Replication() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference) ReplicationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNetappVolumeHybridReplicationParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NetappVolumeHybridReplicationParametersOutputReference {
	_init_.Initialize()

	if err := validateNewNetappVolumeHybridReplicationParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.netappVolume.NetappVolumeHybridReplicationParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetappVolumeHybridReplicationParametersOutputReference_Override(n NetappVolumeHybridReplicationParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.netappVolume.NetappVolumeHybridReplicationParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference)SetClusterLocation(val *string) {
	if err := j.validateSetClusterLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterLocation",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference)SetInternalValue(val *NetappVolumeHybridReplicationParameters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference)SetPeerClusterName(val *string) {
	if err := j.validateSetPeerClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerClusterName",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference)SetPeerIpAddresses(val *string) {
	if err := j.validateSetPeerIpAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerIpAddresses",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference)SetPeerSvmName(val *string) {
	if err := j.validateSetPeerSvmNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerSvmName",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference)SetPeerVolumeName(val *string) {
	if err := j.validateSetPeerVolumeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerVolumeName",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference)SetReplication(val *string) {
	if err := j.validateSetReplicationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replication",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference) ResetClusterLocation() {
	_jsii_.InvokeVoid(
		n,
		"resetClusterLocation",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		n,
		"resetDescription",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		n,
		"resetLabels",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference) ResetPeerClusterName() {
	_jsii_.InvokeVoid(
		n,
		"resetPeerClusterName",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference) ResetPeerIpAddresses() {
	_jsii_.InvokeVoid(
		n,
		"resetPeerIpAddresses",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference) ResetPeerSvmName() {
	_jsii_.InvokeVoid(
		n,
		"resetPeerSvmName",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference) ResetPeerVolumeName() {
	_jsii_.InvokeVoid(
		n,
		"resetPeerVolumeName",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference) ResetReplication() {
	_jsii_.InvokeVoid(
		n,
		"resetReplication",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (n *jsiiProxy_NetappVolumeHybridReplicationParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

