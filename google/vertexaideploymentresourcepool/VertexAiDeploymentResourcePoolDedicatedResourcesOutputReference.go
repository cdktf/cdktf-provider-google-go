// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vertexaideploymentresourcepool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/vertexaideploymentresourcepool/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference interface {
	cdktf.ComplexObject
	AutoscalingMetricSpecs() VertexAiDeploymentResourcePoolDedicatedResourcesAutoscalingMetricSpecsList
	AutoscalingMetricSpecsInput() interface{}
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
	InternalValue() *VertexAiDeploymentResourcePoolDedicatedResources
	SetInternalValue(val *VertexAiDeploymentResourcePoolDedicatedResources)
	MachineSpec() VertexAiDeploymentResourcePoolDedicatedResourcesMachineSpecOutputReference
	MachineSpecInput() *VertexAiDeploymentResourcePoolDedicatedResourcesMachineSpec
	MaxReplicaCount() *float64
	SetMaxReplicaCount(val *float64)
	MaxReplicaCountInput() *float64
	MinReplicaCount() *float64
	SetMinReplicaCount(val *float64)
	MinReplicaCountInput() *float64
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
	PutAutoscalingMetricSpecs(value interface{})
	PutMachineSpec(value *VertexAiDeploymentResourcePoolDedicatedResourcesMachineSpec)
	ResetAutoscalingMetricSpecs()
	ResetMaxReplicaCount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference
type jsiiProxy_VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference) AutoscalingMetricSpecs() VertexAiDeploymentResourcePoolDedicatedResourcesAutoscalingMetricSpecsList {
	var returns VertexAiDeploymentResourcePoolDedicatedResourcesAutoscalingMetricSpecsList
	_jsii_.Get(
		j,
		"autoscalingMetricSpecs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference) AutoscalingMetricSpecsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoscalingMetricSpecsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference) InternalValue() *VertexAiDeploymentResourcePoolDedicatedResources {
	var returns *VertexAiDeploymentResourcePoolDedicatedResources
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference) MachineSpec() VertexAiDeploymentResourcePoolDedicatedResourcesMachineSpecOutputReference {
	var returns VertexAiDeploymentResourcePoolDedicatedResourcesMachineSpecOutputReference
	_jsii_.Get(
		j,
		"machineSpec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference) MachineSpecInput() *VertexAiDeploymentResourcePoolDedicatedResourcesMachineSpec {
	var returns *VertexAiDeploymentResourcePoolDedicatedResourcesMachineSpec
	_jsii_.Get(
		j,
		"machineSpecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference) MaxReplicaCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxReplicaCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference) MaxReplicaCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxReplicaCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference) MinReplicaCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minReplicaCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference) MinReplicaCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minReplicaCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVertexAiDeploymentResourcePoolDedicatedResourcesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference {
	_init_.Initialize()

	if err := validateNewVertexAiDeploymentResourcePoolDedicatedResourcesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.vertexAiDeploymentResourcePool.VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVertexAiDeploymentResourcePoolDedicatedResourcesOutputReference_Override(v VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.vertexAiDeploymentResourcePool.VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference)SetInternalValue(val *VertexAiDeploymentResourcePoolDedicatedResources) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference)SetMaxReplicaCount(val *float64) {
	if err := j.validateSetMaxReplicaCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxReplicaCount",
		val,
	)
}

func (j *jsiiProxy_VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference)SetMinReplicaCount(val *float64) {
	if err := j.validateSetMinReplicaCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minReplicaCount",
		val,
	)
}

func (j *jsiiProxy_VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference) PutAutoscalingMetricSpecs(value interface{}) {
	if err := v.validatePutAutoscalingMetricSpecsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putAutoscalingMetricSpecs",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference) PutMachineSpec(value *VertexAiDeploymentResourcePoolDedicatedResourcesMachineSpec) {
	if err := v.validatePutMachineSpecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putMachineSpec",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference) ResetAutoscalingMetricSpecs() {
	_jsii_.InvokeVoid(
		v,
		"resetAutoscalingMetricSpecs",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference) ResetMaxReplicaCount() {
	_jsii_.InvokeVoid(
		v,
		"resetMaxReplicaCount",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VertexAiDeploymentResourcePoolDedicatedResourcesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

