// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vertexaiindex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/vertexaiindex/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VertexAiIndexMetadataConfigOutputReference interface {
	cdktf.ComplexObject
	AlgorithmConfig() VertexAiIndexMetadataConfigAlgorithmConfigOutputReference
	AlgorithmConfigInput() *VertexAiIndexMetadataConfigAlgorithmConfig
	ApproximateNeighborsCount() *float64
	SetApproximateNeighborsCount(val *float64)
	ApproximateNeighborsCountInput() *float64
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
	Dimensions() *float64
	SetDimensions(val *float64)
	DimensionsInput() *float64
	DistanceMeasureType() *string
	SetDistanceMeasureType(val *string)
	DistanceMeasureTypeInput() *string
	FeatureNormType() *string
	SetFeatureNormType(val *string)
	FeatureNormTypeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *VertexAiIndexMetadataConfig
	SetInternalValue(val *VertexAiIndexMetadataConfig)
	ShardSize() *string
	SetShardSize(val *string)
	ShardSizeInput() *string
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
	PutAlgorithmConfig(value *VertexAiIndexMetadataConfigAlgorithmConfig)
	ResetAlgorithmConfig()
	ResetApproximateNeighborsCount()
	ResetDistanceMeasureType()
	ResetFeatureNormType()
	ResetShardSize()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VertexAiIndexMetadataConfigOutputReference
type jsiiProxy_VertexAiIndexMetadataConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigOutputReference) AlgorithmConfig() VertexAiIndexMetadataConfigAlgorithmConfigOutputReference {
	var returns VertexAiIndexMetadataConfigAlgorithmConfigOutputReference
	_jsii_.Get(
		j,
		"algorithmConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigOutputReference) AlgorithmConfigInput() *VertexAiIndexMetadataConfigAlgorithmConfig {
	var returns *VertexAiIndexMetadataConfigAlgorithmConfig
	_jsii_.Get(
		j,
		"algorithmConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigOutputReference) ApproximateNeighborsCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"approximateNeighborsCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigOutputReference) ApproximateNeighborsCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"approximateNeighborsCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigOutputReference) Dimensions() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dimensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigOutputReference) DimensionsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dimensionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigOutputReference) DistanceMeasureType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distanceMeasureType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigOutputReference) DistanceMeasureTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distanceMeasureTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigOutputReference) FeatureNormType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"featureNormType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigOutputReference) FeatureNormTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"featureNormTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigOutputReference) InternalValue() *VertexAiIndexMetadataConfig {
	var returns *VertexAiIndexMetadataConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigOutputReference) ShardSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shardSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigOutputReference) ShardSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shardSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVertexAiIndexMetadataConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VertexAiIndexMetadataConfigOutputReference {
	_init_.Initialize()

	if err := validateNewVertexAiIndexMetadataConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiIndexMetadataConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.vertexAiIndex.VertexAiIndexMetadataConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVertexAiIndexMetadataConfigOutputReference_Override(v VertexAiIndexMetadataConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.vertexAiIndex.VertexAiIndexMetadataConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigOutputReference)SetApproximateNeighborsCount(val *float64) {
	if err := j.validateSetApproximateNeighborsCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"approximateNeighborsCount",
		val,
	)
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigOutputReference)SetDimensions(val *float64) {
	if err := j.validateSetDimensionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dimensions",
		val,
	)
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigOutputReference)SetDistanceMeasureType(val *string) {
	if err := j.validateSetDistanceMeasureTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"distanceMeasureType",
		val,
	)
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigOutputReference)SetFeatureNormType(val *string) {
	if err := j.validateSetFeatureNormTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"featureNormType",
		val,
	)
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigOutputReference)SetInternalValue(val *VertexAiIndexMetadataConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigOutputReference)SetShardSize(val *string) {
	if err := j.validateSetShardSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shardSize",
		val,
	)
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VertexAiIndexMetadataConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexMetadataConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexMetadataConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexMetadataConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexMetadataConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexMetadataConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexMetadataConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexMetadataConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexMetadataConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexMetadataConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexMetadataConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexMetadataConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexMetadataConfigOutputReference) PutAlgorithmConfig(value *VertexAiIndexMetadataConfigAlgorithmConfig) {
	if err := v.validatePutAlgorithmConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putAlgorithmConfig",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiIndexMetadataConfigOutputReference) ResetAlgorithmConfig() {
	_jsii_.InvokeVoid(
		v,
		"resetAlgorithmConfig",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiIndexMetadataConfigOutputReference) ResetApproximateNeighborsCount() {
	_jsii_.InvokeVoid(
		v,
		"resetApproximateNeighborsCount",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiIndexMetadataConfigOutputReference) ResetDistanceMeasureType() {
	_jsii_.InvokeVoid(
		v,
		"resetDistanceMeasureType",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiIndexMetadataConfigOutputReference) ResetFeatureNormType() {
	_jsii_.InvokeVoid(
		v,
		"resetFeatureNormType",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiIndexMetadataConfigOutputReference) ResetShardSize() {
	_jsii_.InvokeVoid(
		v,
		"resetShardSize",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiIndexMetadataConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := v.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexMetadataConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

