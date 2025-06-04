// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigtableinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/bigtableinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BigtableInstanceClusterAutoscalingConfigOutputReference interface {
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
	CpuTarget() *float64
	SetCpuTarget(val *float64)
	CpuTargetInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *BigtableInstanceClusterAutoscalingConfig
	SetInternalValue(val *BigtableInstanceClusterAutoscalingConfig)
	MaxNodes() *float64
	SetMaxNodes(val *float64)
	MaxNodesInput() *float64
	MinNodes() *float64
	SetMinNodes(val *float64)
	MinNodesInput() *float64
	StorageTarget() *float64
	SetStorageTarget(val *float64)
	StorageTargetInput() *float64
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
	ResetStorageTarget()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BigtableInstanceClusterAutoscalingConfigOutputReference
type jsiiProxy_BigtableInstanceClusterAutoscalingConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BigtableInstanceClusterAutoscalingConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableInstanceClusterAutoscalingConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableInstanceClusterAutoscalingConfigOutputReference) CpuTarget() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableInstanceClusterAutoscalingConfigOutputReference) CpuTargetInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableInstanceClusterAutoscalingConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableInstanceClusterAutoscalingConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableInstanceClusterAutoscalingConfigOutputReference) InternalValue() *BigtableInstanceClusterAutoscalingConfig {
	var returns *BigtableInstanceClusterAutoscalingConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableInstanceClusterAutoscalingConfigOutputReference) MaxNodes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableInstanceClusterAutoscalingConfigOutputReference) MaxNodesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableInstanceClusterAutoscalingConfigOutputReference) MinNodes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableInstanceClusterAutoscalingConfigOutputReference) MinNodesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableInstanceClusterAutoscalingConfigOutputReference) StorageTarget() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableInstanceClusterAutoscalingConfigOutputReference) StorageTargetInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableInstanceClusterAutoscalingConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigtableInstanceClusterAutoscalingConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBigtableInstanceClusterAutoscalingConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BigtableInstanceClusterAutoscalingConfigOutputReference {
	_init_.Initialize()

	if err := validateNewBigtableInstanceClusterAutoscalingConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BigtableInstanceClusterAutoscalingConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.bigtableInstance.BigtableInstanceClusterAutoscalingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBigtableInstanceClusterAutoscalingConfigOutputReference_Override(b BigtableInstanceClusterAutoscalingConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.bigtableInstance.BigtableInstanceClusterAutoscalingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BigtableInstanceClusterAutoscalingConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BigtableInstanceClusterAutoscalingConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BigtableInstanceClusterAutoscalingConfigOutputReference)SetCpuTarget(val *float64) {
	if err := j.validateSetCpuTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuTarget",
		val,
	)
}

func (j *jsiiProxy_BigtableInstanceClusterAutoscalingConfigOutputReference)SetInternalValue(val *BigtableInstanceClusterAutoscalingConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BigtableInstanceClusterAutoscalingConfigOutputReference)SetMaxNodes(val *float64) {
	if err := j.validateSetMaxNodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxNodes",
		val,
	)
}

func (j *jsiiProxy_BigtableInstanceClusterAutoscalingConfigOutputReference)SetMinNodes(val *float64) {
	if err := j.validateSetMinNodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minNodes",
		val,
	)
}

func (j *jsiiProxy_BigtableInstanceClusterAutoscalingConfigOutputReference)SetStorageTarget(val *float64) {
	if err := j.validateSetStorageTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageTarget",
		val,
	)
}

func (j *jsiiProxy_BigtableInstanceClusterAutoscalingConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BigtableInstanceClusterAutoscalingConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BigtableInstanceClusterAutoscalingConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigtableInstanceClusterAutoscalingConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigtableInstanceClusterAutoscalingConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigtableInstanceClusterAutoscalingConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigtableInstanceClusterAutoscalingConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigtableInstanceClusterAutoscalingConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigtableInstanceClusterAutoscalingConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigtableInstanceClusterAutoscalingConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigtableInstanceClusterAutoscalingConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigtableInstanceClusterAutoscalingConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigtableInstanceClusterAutoscalingConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigtableInstanceClusterAutoscalingConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigtableInstanceClusterAutoscalingConfigOutputReference) ResetStorageTarget() {
	_jsii_.InvokeVoid(
		b,
		"resetStorageTarget",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigtableInstanceClusterAutoscalingConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := b.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigtableInstanceClusterAutoscalingConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

