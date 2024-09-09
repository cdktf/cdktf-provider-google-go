// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vertexaifeatureonlinestore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/vertexaifeatureonlinestore/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference interface {
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
	CpuUtilizationTarget() *float64
	SetCpuUtilizationTarget(val *float64)
	CpuUtilizationTargetInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *VertexAiFeatureOnlineStoreBigtableAutoScaling
	SetInternalValue(val *VertexAiFeatureOnlineStoreBigtableAutoScaling)
	MaxNodeCount() *float64
	SetMaxNodeCount(val *float64)
	MaxNodeCountInput() *float64
	MinNodeCount() *float64
	SetMinNodeCount(val *float64)
	MinNodeCountInput() *float64
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
	ResetCpuUtilizationTarget()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference
type jsiiProxy_VertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference) CpuUtilizationTarget() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuUtilizationTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference) CpuUtilizationTargetInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuUtilizationTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference) InternalValue() *VertexAiFeatureOnlineStoreBigtableAutoScaling {
	var returns *VertexAiFeatureOnlineStoreBigtableAutoScaling
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference) MaxNodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxNodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference) MaxNodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxNodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference) MinNodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference) MinNodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference {
	_init_.Initialize()

	if err := validateNewVertexAiFeatureOnlineStoreBigtableAutoScalingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.vertexAiFeatureOnlineStore.VertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference_Override(v VertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.vertexAiFeatureOnlineStore.VertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference)SetCpuUtilizationTarget(val *float64) {
	if err := j.validateSetCpuUtilizationTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuUtilizationTarget",
		val,
	)
}

func (j *jsiiProxy_VertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference)SetInternalValue(val *VertexAiFeatureOnlineStoreBigtableAutoScaling) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference)SetMaxNodeCount(val *float64) {
	if err := j.validateSetMaxNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxNodeCount",
		val,
	)
}

func (j *jsiiProxy_VertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference)SetMinNodeCount(val *float64) {
	if err := j.validateSetMinNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minNodeCount",
		val,
	)
}

func (j *jsiiProxy_VertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference) ResetCpuUtilizationTarget() {
	_jsii_.InvokeVoid(
		v,
		"resetCpuUtilizationTarget",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VertexAiFeatureOnlineStoreBigtableAutoScalingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

