// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networksecuritymirroringendpointgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/networksecuritymirroringendpointgroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkSecurityMirroringEndpointGroupConnectedDeploymentGroupsLocationsList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) NetworkSecurityMirroringEndpointGroupConnectedDeploymentGroupsLocationsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkSecurityMirroringEndpointGroupConnectedDeploymentGroupsLocationsList
type jsiiProxy_NetworkSecurityMirroringEndpointGroupConnectedDeploymentGroupsLocationsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_NetworkSecurityMirroringEndpointGroupConnectedDeploymentGroupsLocationsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityMirroringEndpointGroupConnectedDeploymentGroupsLocationsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityMirroringEndpointGroupConnectedDeploymentGroupsLocationsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityMirroringEndpointGroupConnectedDeploymentGroupsLocationsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityMirroringEndpointGroupConnectedDeploymentGroupsLocationsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewNetworkSecurityMirroringEndpointGroupConnectedDeploymentGroupsLocationsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) NetworkSecurityMirroringEndpointGroupConnectedDeploymentGroupsLocationsList {
	_init_.Initialize()

	if err := validateNewNetworkSecurityMirroringEndpointGroupConnectedDeploymentGroupsLocationsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkSecurityMirroringEndpointGroupConnectedDeploymentGroupsLocationsList{}

	_jsii_.Create(
		"@cdktf/provider-google.networkSecurityMirroringEndpointGroup.NetworkSecurityMirroringEndpointGroupConnectedDeploymentGroupsLocationsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewNetworkSecurityMirroringEndpointGroupConnectedDeploymentGroupsLocationsList_Override(n NetworkSecurityMirroringEndpointGroupConnectedDeploymentGroupsLocationsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.networkSecurityMirroringEndpointGroup.NetworkSecurityMirroringEndpointGroupConnectedDeploymentGroupsLocationsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		n,
	)
}

func (j *jsiiProxy_NetworkSecurityMirroringEndpointGroupConnectedDeploymentGroupsLocationsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityMirroringEndpointGroupConnectedDeploymentGroupsLocationsList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityMirroringEndpointGroupConnectedDeploymentGroupsLocationsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (n *jsiiProxy_NetworkSecurityMirroringEndpointGroupConnectedDeploymentGroupsLocationsList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := n.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		n,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityMirroringEndpointGroupConnectedDeploymentGroupsLocationsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityMirroringEndpointGroupConnectedDeploymentGroupsLocationsList) Get(index *float64) NetworkSecurityMirroringEndpointGroupConnectedDeploymentGroupsLocationsOutputReference {
	if err := n.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns NetworkSecurityMirroringEndpointGroupConnectedDeploymentGroupsLocationsOutputReference

	_jsii_.Invoke(
		n,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityMirroringEndpointGroupConnectedDeploymentGroupsLocationsList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (n *jsiiProxy_NetworkSecurityMirroringEndpointGroupConnectedDeploymentGroupsLocationsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

