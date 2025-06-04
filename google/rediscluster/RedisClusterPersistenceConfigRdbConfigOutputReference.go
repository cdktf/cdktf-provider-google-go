// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rediscluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/rediscluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RedisClusterPersistenceConfigRdbConfigOutputReference interface {
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
	InternalValue() *RedisClusterPersistenceConfigRdbConfig
	SetInternalValue(val *RedisClusterPersistenceConfigRdbConfig)
	RdbSnapshotPeriod() *string
	SetRdbSnapshotPeriod(val *string)
	RdbSnapshotPeriodInput() *string
	RdbSnapshotStartTime() *string
	SetRdbSnapshotStartTime(val *string)
	RdbSnapshotStartTimeInput() *string
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
	ResetRdbSnapshotPeriod()
	ResetRdbSnapshotStartTime()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RedisClusterPersistenceConfigRdbConfigOutputReference
type jsiiProxy_RedisClusterPersistenceConfigRdbConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_RedisClusterPersistenceConfigRdbConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisClusterPersistenceConfigRdbConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisClusterPersistenceConfigRdbConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisClusterPersistenceConfigRdbConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisClusterPersistenceConfigRdbConfigOutputReference) InternalValue() *RedisClusterPersistenceConfigRdbConfig {
	var returns *RedisClusterPersistenceConfigRdbConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisClusterPersistenceConfigRdbConfigOutputReference) RdbSnapshotPeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdbSnapshotPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisClusterPersistenceConfigRdbConfigOutputReference) RdbSnapshotPeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdbSnapshotPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisClusterPersistenceConfigRdbConfigOutputReference) RdbSnapshotStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdbSnapshotStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisClusterPersistenceConfigRdbConfigOutputReference) RdbSnapshotStartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdbSnapshotStartTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisClusterPersistenceConfigRdbConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisClusterPersistenceConfigRdbConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewRedisClusterPersistenceConfigRdbConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) RedisClusterPersistenceConfigRdbConfigOutputReference {
	_init_.Initialize()

	if err := validateNewRedisClusterPersistenceConfigRdbConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_RedisClusterPersistenceConfigRdbConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.redisCluster.RedisClusterPersistenceConfigRdbConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRedisClusterPersistenceConfigRdbConfigOutputReference_Override(r RedisClusterPersistenceConfigRdbConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.redisCluster.RedisClusterPersistenceConfigRdbConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_RedisClusterPersistenceConfigRdbConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RedisClusterPersistenceConfigRdbConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RedisClusterPersistenceConfigRdbConfigOutputReference)SetInternalValue(val *RedisClusterPersistenceConfigRdbConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RedisClusterPersistenceConfigRdbConfigOutputReference)SetRdbSnapshotPeriod(val *string) {
	if err := j.validateSetRdbSnapshotPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rdbSnapshotPeriod",
		val,
	)
}

func (j *jsiiProxy_RedisClusterPersistenceConfigRdbConfigOutputReference)SetRdbSnapshotStartTime(val *string) {
	if err := j.validateSetRdbSnapshotStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rdbSnapshotStartTime",
		val,
	)
}

func (j *jsiiProxy_RedisClusterPersistenceConfigRdbConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RedisClusterPersistenceConfigRdbConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_RedisClusterPersistenceConfigRdbConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisClusterPersistenceConfigRdbConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisClusterPersistenceConfigRdbConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisClusterPersistenceConfigRdbConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisClusterPersistenceConfigRdbConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisClusterPersistenceConfigRdbConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisClusterPersistenceConfigRdbConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisClusterPersistenceConfigRdbConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisClusterPersistenceConfigRdbConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisClusterPersistenceConfigRdbConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisClusterPersistenceConfigRdbConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisClusterPersistenceConfigRdbConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisClusterPersistenceConfigRdbConfigOutputReference) ResetRdbSnapshotPeriod() {
	_jsii_.InvokeVoid(
		r,
		"resetRdbSnapshotPeriod",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisClusterPersistenceConfigRdbConfigOutputReference) ResetRdbSnapshotStartTime() {
	_jsii_.InvokeVoid(
		r,
		"resetRdbSnapshotStartTime",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisClusterPersistenceConfigRdbConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (r *jsiiProxy_RedisClusterPersistenceConfigRdbConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

