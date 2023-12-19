// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spannerinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v12/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v12/spannerinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference interface {
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
	InternalValue() *SpannerInstanceAutoscalingConfigAutoscalingLimits
	SetInternalValue(val *SpannerInstanceAutoscalingConfigAutoscalingLimits)
	MaxNodes() *float64
	SetMaxNodes(val *float64)
	MaxNodesInput() *float64
	MaxProcessingUnits() *float64
	SetMaxProcessingUnits(val *float64)
	MaxProcessingUnitsInput() *float64
	MinNodes() *float64
	SetMinNodes(val *float64)
	MinNodesInput() *float64
	MinProcessingUnits() *float64
	SetMinProcessingUnits(val *float64)
	MinProcessingUnitsInput() *float64
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
	ResetMaxNodes()
	ResetMaxProcessingUnits()
	ResetMinNodes()
	ResetMinProcessingUnits()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference
type jsiiProxy_SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) InternalValue() *SpannerInstanceAutoscalingConfigAutoscalingLimits {
	var returns *SpannerInstanceAutoscalingConfigAutoscalingLimits
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) MaxNodes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) MaxNodesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) MaxProcessingUnits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxProcessingUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) MaxProcessingUnitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxProcessingUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) MinNodes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) MinNodesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) MinProcessingUnits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minProcessingUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) MinProcessingUnitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minProcessingUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference {
	_init_.Initialize()

	if err := validateNewSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.spannerInstance.SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference_Override(s SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.spannerInstance.SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference)SetInternalValue(val *SpannerInstanceAutoscalingConfigAutoscalingLimits) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference)SetMaxNodes(val *float64) {
	if err := j.validateSetMaxNodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxNodes",
		val,
	)
}

func (j *jsiiProxy_SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference)SetMaxProcessingUnits(val *float64) {
	if err := j.validateSetMaxProcessingUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxProcessingUnits",
		val,
	)
}

func (j *jsiiProxy_SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference)SetMinNodes(val *float64) {
	if err := j.validateSetMinNodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minNodes",
		val,
	)
}

func (j *jsiiProxy_SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference)SetMinProcessingUnits(val *float64) {
	if err := j.validateSetMinProcessingUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minProcessingUnits",
		val,
	)
}

func (j *jsiiProxy_SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) ResetMaxNodes() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxNodes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) ResetMaxProcessingUnits() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxProcessingUnits",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) ResetMinNodes() {
	_jsii_.InvokeVoid(
		s,
		"resetMinNodes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) ResetMinProcessingUnits() {
	_jsii_.InvokeVoid(
		s,
		"resetMinProcessingUnits",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

