// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package blockchainnodeengineblockchainnodes

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/blockchainnodeengineblockchainnodes/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsOutputReference interface {
	cdktf.ComplexObject
	BeaconApiEndpoint() *string
	BeaconPrometheusMetricsApiEndpoint() *string
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
	ExecutionClientPrometheusMetricsApiEndpoint() *string
	// Experimental.
	Fqn() *string
	InternalValue() *BlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpoints
	SetInternalValue(val *BlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpoints)
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsOutputReference
type jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsOutputReference) BeaconApiEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"beaconApiEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsOutputReference) BeaconPrometheusMetricsApiEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"beaconPrometheusMetricsApiEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsOutputReference) ExecutionClientPrometheusMetricsApiEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionClientPrometheusMetricsApiEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsOutputReference) InternalValue() *BlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpoints {
	var returns *BlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpoints
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsOutputReference {
	_init_.Initialize()

	if err := validateNewBlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.blockchainNodeEngineBlockchainNodes.BlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsOutputReference_Override(b BlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.blockchainNodeEngineBlockchainNodes.BlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsOutputReference)SetInternalValue(val *BlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpoints) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (b *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

