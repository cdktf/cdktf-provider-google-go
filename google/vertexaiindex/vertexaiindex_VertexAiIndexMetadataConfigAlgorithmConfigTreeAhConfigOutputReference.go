package vertexaiindex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v4/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v4/vertexaiindex/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference interface {
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
	InternalValue() *VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfig
	SetInternalValue(val *VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfig)
	LeafNodeEmbeddingCount() *float64
	SetLeafNodeEmbeddingCount(val *float64)
	LeafNodeEmbeddingCountInput() *float64
	LeafNodesToSearchPercent() *float64
	SetLeafNodesToSearchPercent(val *float64)
	LeafNodesToSearchPercentInput() *float64
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
	ResetLeafNodeEmbeddingCount()
	ResetLeafNodesToSearchPercent()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference
type jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) InternalValue() *VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfig {
	var returns *VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) LeafNodeEmbeddingCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"leafNodeEmbeddingCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) LeafNodeEmbeddingCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"leafNodeEmbeddingCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) LeafNodesToSearchPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"leafNodesToSearchPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) LeafNodesToSearchPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"leafNodesToSearchPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference {
	_init_.Initialize()

	if err := validateNewVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.vertexAiIndex.VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference_Override(v VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.vertexAiIndex.VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference)SetInternalValue(val *VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference)SetLeafNodeEmbeddingCount(val *float64) {
	if err := j.validateSetLeafNodeEmbeddingCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"leafNodeEmbeddingCount",
		val,
	)
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference)SetLeafNodesToSearchPercent(val *float64) {
	if err := j.validateSetLeafNodesToSearchPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"leafNodesToSearchPercent",
		val,
	)
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) ResetLeafNodeEmbeddingCount() {
	_jsii_.InvokeVoid(
		v,
		"resetLeafNodeEmbeddingCount",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) ResetLeafNodesToSearchPercent() {
	_jsii_.InvokeVoid(
		v,
		"resetLeafNodesToSearchPercent",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

