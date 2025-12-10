// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vertexaiindex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/vertexaiindex/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VertexAiIndexMetadataConfigAlgorithmConfigOutputReference interface {
	cdktf.ComplexObject
	BruteForceConfig() VertexAiIndexMetadataConfigAlgorithmConfigBruteForceConfigOutputReference
	BruteForceConfigInput() *VertexAiIndexMetadataConfigAlgorithmConfigBruteForceConfig
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
	InternalValue() *VertexAiIndexMetadataConfigAlgorithmConfig
	SetInternalValue(val *VertexAiIndexMetadataConfigAlgorithmConfig)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TreeAhConfig() VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference
	TreeAhConfigInput() *VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfig
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutBruteForceConfig(value *VertexAiIndexMetadataConfigAlgorithmConfigBruteForceConfig)
	PutTreeAhConfig(value *VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfig)
	ResetBruteForceConfig()
	ResetTreeAhConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VertexAiIndexMetadataConfigAlgorithmConfigOutputReference
type jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigOutputReference) BruteForceConfig() VertexAiIndexMetadataConfigAlgorithmConfigBruteForceConfigOutputReference {
	var returns VertexAiIndexMetadataConfigAlgorithmConfigBruteForceConfigOutputReference
	_jsii_.Get(
		j,
		"bruteForceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigOutputReference) BruteForceConfigInput() *VertexAiIndexMetadataConfigAlgorithmConfigBruteForceConfig {
	var returns *VertexAiIndexMetadataConfigAlgorithmConfigBruteForceConfig
	_jsii_.Get(
		j,
		"bruteForceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigOutputReference) InternalValue() *VertexAiIndexMetadataConfigAlgorithmConfig {
	var returns *VertexAiIndexMetadataConfigAlgorithmConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigOutputReference) TreeAhConfig() VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference {
	var returns VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfigOutputReference
	_jsii_.Get(
		j,
		"treeAhConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigOutputReference) TreeAhConfigInput() *VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfig {
	var returns *VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfig
	_jsii_.Get(
		j,
		"treeAhConfigInput",
		&returns,
	)
	return returns
}


func NewVertexAiIndexMetadataConfigAlgorithmConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VertexAiIndexMetadataConfigAlgorithmConfigOutputReference {
	_init_.Initialize()

	if err := validateNewVertexAiIndexMetadataConfigAlgorithmConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.vertexAiIndex.VertexAiIndexMetadataConfigAlgorithmConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVertexAiIndexMetadataConfigAlgorithmConfigOutputReference_Override(v VertexAiIndexMetadataConfigAlgorithmConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.vertexAiIndex.VertexAiIndexMetadataConfigAlgorithmConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigOutputReference)SetInternalValue(val *VertexAiIndexMetadataConfigAlgorithmConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigOutputReference) PutBruteForceConfig(value *VertexAiIndexMetadataConfigAlgorithmConfigBruteForceConfig) {
	if err := v.validatePutBruteForceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putBruteForceConfig",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigOutputReference) PutTreeAhConfig(value *VertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfig) {
	if err := v.validatePutTreeAhConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putTreeAhConfig",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigOutputReference) ResetBruteForceConfig() {
	_jsii_.InvokeVoid(
		v,
		"resetBruteForceConfig",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigOutputReference) ResetTreeAhConfig() {
	_jsii_.InvokeVoid(
		v,
		"resetTreeAhConfig",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := v.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiIndexMetadataConfigAlgorithmConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

