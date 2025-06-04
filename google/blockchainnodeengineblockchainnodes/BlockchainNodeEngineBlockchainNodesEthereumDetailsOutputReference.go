// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package blockchainnodeengineblockchainnodes

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/blockchainnodeengineblockchainnodes/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference interface {
	cdktf.ComplexObject
	AdditionalEndpoints() BlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsList
	ApiEnableAdmin() interface{}
	SetApiEnableAdmin(val interface{})
	ApiEnableAdminInput() interface{}
	ApiEnableDebug() interface{}
	SetApiEnableDebug(val interface{})
	ApiEnableDebugInput() interface{}
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
	ConsensusClient() *string
	SetConsensusClient(val *string)
	ConsensusClientInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ExecutionClient() *string
	SetExecutionClient(val *string)
	ExecutionClientInput() *string
	FetchhDetails() BlockchainNodeEngineBlockchainNodesEthereumDetailsGethDetailsOutputReference
	FetchhDetailsInput() *BlockchainNodeEngineBlockchainNodesEthereumDetailsGethDetails
	// Experimental.
	Fqn() *string
	InternalValue() *BlockchainNodeEngineBlockchainNodesEthereumDetails
	SetInternalValue(val *BlockchainNodeEngineBlockchainNodesEthereumDetails)
	Network() *string
	SetNetwork(val *string)
	NetworkInput() *string
	NodeType() *string
	SetNodeType(val *string)
	NodeTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ValidatorConfig() BlockchainNodeEngineBlockchainNodesEthereumDetailsValidatorConfigOutputReference
	ValidatorConfigInput() *BlockchainNodeEngineBlockchainNodesEthereumDetailsValidatorConfig
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
	PutFetchhDetails(value *BlockchainNodeEngineBlockchainNodesEthereumDetailsGethDetails)
	PutValidatorConfig(value *BlockchainNodeEngineBlockchainNodesEthereumDetailsValidatorConfig)
	ResetApiEnableAdmin()
	ResetApiEnableDebug()
	ResetConsensusClient()
	ResetExecutionClient()
	ResetFetchhDetails()
	ResetNetwork()
	ResetNodeType()
	ResetValidatorConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference
type jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) AdditionalEndpoints() BlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsList {
	var returns BlockchainNodeEngineBlockchainNodesEthereumDetailsAdditionalEndpointsList
	_jsii_.Get(
		j,
		"additionalEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) ApiEnableAdmin() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"apiEnableAdmin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) ApiEnableAdminInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"apiEnableAdminInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) ApiEnableDebug() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"apiEnableDebug",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) ApiEnableDebugInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"apiEnableDebugInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) ConsensusClient() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consensusClient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) ConsensusClientInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consensusClientInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) ExecutionClient() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionClient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) ExecutionClientInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionClientInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) FetchhDetails() BlockchainNodeEngineBlockchainNodesEthereumDetailsGethDetailsOutputReference {
	var returns BlockchainNodeEngineBlockchainNodesEthereumDetailsGethDetailsOutputReference
	_jsii_.Get(
		j,
		"fetchhDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) FetchhDetailsInput() *BlockchainNodeEngineBlockchainNodesEthereumDetailsGethDetails {
	var returns *BlockchainNodeEngineBlockchainNodesEthereumDetailsGethDetails
	_jsii_.Get(
		j,
		"fetchhDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) InternalValue() *BlockchainNodeEngineBlockchainNodesEthereumDetails {
	var returns *BlockchainNodeEngineBlockchainNodesEthereumDetails
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) NodeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) NodeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) ValidatorConfig() BlockchainNodeEngineBlockchainNodesEthereumDetailsValidatorConfigOutputReference {
	var returns BlockchainNodeEngineBlockchainNodesEthereumDetailsValidatorConfigOutputReference
	_jsii_.Get(
		j,
		"validatorConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) ValidatorConfigInput() *BlockchainNodeEngineBlockchainNodesEthereumDetailsValidatorConfig {
	var returns *BlockchainNodeEngineBlockchainNodesEthereumDetailsValidatorConfig
	_jsii_.Get(
		j,
		"validatorConfigInput",
		&returns,
	)
	return returns
}


func NewBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference {
	_init_.Initialize()

	if err := validateNewBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.blockchainNodeEngineBlockchainNodes.BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference_Override(b BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.blockchainNodeEngineBlockchainNodes.BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference)SetApiEnableAdmin(val interface{}) {
	if err := j.validateSetApiEnableAdminParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiEnableAdmin",
		val,
	)
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference)SetApiEnableDebug(val interface{}) {
	if err := j.validateSetApiEnableDebugParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiEnableDebug",
		val,
	)
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference)SetConsensusClient(val *string) {
	if err := j.validateSetConsensusClientParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consensusClient",
		val,
	)
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference)SetExecutionClient(val *string) {
	if err := j.validateSetExecutionClientParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionClient",
		val,
	)
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference)SetInternalValue(val *BlockchainNodeEngineBlockchainNodesEthereumDetails) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference)SetNodeType(val *string) {
	if err := j.validateSetNodeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeType",
		val,
	)
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) PutFetchhDetails(value *BlockchainNodeEngineBlockchainNodesEthereumDetailsGethDetails) {
	if err := b.validatePutFetchhDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putFetchhDetails",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) PutValidatorConfig(value *BlockchainNodeEngineBlockchainNodesEthereumDetailsValidatorConfig) {
	if err := b.validatePutValidatorConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putValidatorConfig",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) ResetApiEnableAdmin() {
	_jsii_.InvokeVoid(
		b,
		"resetApiEnableAdmin",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) ResetApiEnableDebug() {
	_jsii_.InvokeVoid(
		b,
		"resetApiEnableDebug",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) ResetConsensusClient() {
	_jsii_.InvokeVoid(
		b,
		"resetConsensusClient",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) ResetExecutionClient() {
	_jsii_.InvokeVoid(
		b,
		"resetExecutionClient",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) ResetFetchhDetails() {
	_jsii_.InvokeVoid(
		b,
		"resetFetchhDetails",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) ResetNetwork() {
	_jsii_.InvokeVoid(
		b,
		"resetNetwork",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) ResetNodeType() {
	_jsii_.InvokeVoid(
		b,
		"resetNodeType",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) ResetValidatorConfig() {
	_jsii_.InvokeVoid(
		b,
		"resetValidatorConfig",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (b *jsiiProxy_BlockchainNodeEngineBlockchainNodesEthereumDetailsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

