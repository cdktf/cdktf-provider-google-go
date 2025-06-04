// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rediscluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/rediscluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RedisClusterCrossClusterReplicationConfigMembershipSecondaryClustersList interface {
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
	Get(index *float64) RedisClusterCrossClusterReplicationConfigMembershipSecondaryClustersOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RedisClusterCrossClusterReplicationConfigMembershipSecondaryClustersList
type jsiiProxy_RedisClusterCrossClusterReplicationConfigMembershipSecondaryClustersList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_RedisClusterCrossClusterReplicationConfigMembershipSecondaryClustersList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisClusterCrossClusterReplicationConfigMembershipSecondaryClustersList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisClusterCrossClusterReplicationConfigMembershipSecondaryClustersList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisClusterCrossClusterReplicationConfigMembershipSecondaryClustersList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisClusterCrossClusterReplicationConfigMembershipSecondaryClustersList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewRedisClusterCrossClusterReplicationConfigMembershipSecondaryClustersList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) RedisClusterCrossClusterReplicationConfigMembershipSecondaryClustersList {
	_init_.Initialize()

	if err := validateNewRedisClusterCrossClusterReplicationConfigMembershipSecondaryClustersListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_RedisClusterCrossClusterReplicationConfigMembershipSecondaryClustersList{}

	_jsii_.Create(
		"@cdktf/provider-google.redisCluster.RedisClusterCrossClusterReplicationConfigMembershipSecondaryClustersList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewRedisClusterCrossClusterReplicationConfigMembershipSecondaryClustersList_Override(r RedisClusterCrossClusterReplicationConfigMembershipSecondaryClustersList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.redisCluster.RedisClusterCrossClusterReplicationConfigMembershipSecondaryClustersList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		r,
	)
}

func (j *jsiiProxy_RedisClusterCrossClusterReplicationConfigMembershipSecondaryClustersList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RedisClusterCrossClusterReplicationConfigMembershipSecondaryClustersList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_RedisClusterCrossClusterReplicationConfigMembershipSecondaryClustersList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (r *jsiiProxy_RedisClusterCrossClusterReplicationConfigMembershipSecondaryClustersList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := r.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		r,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisClusterCrossClusterReplicationConfigMembershipSecondaryClustersList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisClusterCrossClusterReplicationConfigMembershipSecondaryClustersList) Get(index *float64) RedisClusterCrossClusterReplicationConfigMembershipSecondaryClustersOutputReference {
	if err := r.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns RedisClusterCrossClusterReplicationConfigMembershipSecondaryClustersOutputReference

	_jsii_.Invoke(
		r,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisClusterCrossClusterReplicationConfigMembershipSecondaryClustersList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := r.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisClusterCrossClusterReplicationConfigMembershipSecondaryClustersList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

