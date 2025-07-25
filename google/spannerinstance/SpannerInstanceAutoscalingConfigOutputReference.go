// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spannerinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/spannerinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SpannerInstanceAutoscalingConfigOutputReference interface {
	cdktf.ComplexObject
	AsymmetricAutoscalingOptions() SpannerInstanceAutoscalingConfigAsymmetricAutoscalingOptionsList
	AsymmetricAutoscalingOptionsInput() interface{}
	AutoscalingLimits() SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference
	AutoscalingLimitsInput() *SpannerInstanceAutoscalingConfigAutoscalingLimits
	AutoscalingTargets() SpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference
	AutoscalingTargetsInput() *SpannerInstanceAutoscalingConfigAutoscalingTargets
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
	InternalValue() *SpannerInstanceAutoscalingConfig
	SetInternalValue(val *SpannerInstanceAutoscalingConfig)
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
	PutAsymmetricAutoscalingOptions(value interface{})
	PutAutoscalingLimits(value *SpannerInstanceAutoscalingConfigAutoscalingLimits)
	PutAutoscalingTargets(value *SpannerInstanceAutoscalingConfigAutoscalingTargets)
	ResetAsymmetricAutoscalingOptions()
	ResetAutoscalingLimits()
	ResetAutoscalingTargets()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SpannerInstanceAutoscalingConfigOutputReference
type jsiiProxy_SpannerInstanceAutoscalingConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SpannerInstanceAutoscalingConfigOutputReference) AsymmetricAutoscalingOptions() SpannerInstanceAutoscalingConfigAsymmetricAutoscalingOptionsList {
	var returns SpannerInstanceAutoscalingConfigAsymmetricAutoscalingOptionsList
	_jsii_.Get(
		j,
		"asymmetricAutoscalingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstanceAutoscalingConfigOutputReference) AsymmetricAutoscalingOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"asymmetricAutoscalingOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstanceAutoscalingConfigOutputReference) AutoscalingLimits() SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference {
	var returns SpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference
	_jsii_.Get(
		j,
		"autoscalingLimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstanceAutoscalingConfigOutputReference) AutoscalingLimitsInput() *SpannerInstanceAutoscalingConfigAutoscalingLimits {
	var returns *SpannerInstanceAutoscalingConfigAutoscalingLimits
	_jsii_.Get(
		j,
		"autoscalingLimitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstanceAutoscalingConfigOutputReference) AutoscalingTargets() SpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference {
	var returns SpannerInstanceAutoscalingConfigAutoscalingTargetsOutputReference
	_jsii_.Get(
		j,
		"autoscalingTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstanceAutoscalingConfigOutputReference) AutoscalingTargetsInput() *SpannerInstanceAutoscalingConfigAutoscalingTargets {
	var returns *SpannerInstanceAutoscalingConfigAutoscalingTargets
	_jsii_.Get(
		j,
		"autoscalingTargetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstanceAutoscalingConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstanceAutoscalingConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstanceAutoscalingConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstanceAutoscalingConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstanceAutoscalingConfigOutputReference) InternalValue() *SpannerInstanceAutoscalingConfig {
	var returns *SpannerInstanceAutoscalingConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstanceAutoscalingConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpannerInstanceAutoscalingConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSpannerInstanceAutoscalingConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SpannerInstanceAutoscalingConfigOutputReference {
	_init_.Initialize()

	if err := validateNewSpannerInstanceAutoscalingConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SpannerInstanceAutoscalingConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.spannerInstance.SpannerInstanceAutoscalingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSpannerInstanceAutoscalingConfigOutputReference_Override(s SpannerInstanceAutoscalingConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.spannerInstance.SpannerInstanceAutoscalingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SpannerInstanceAutoscalingConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SpannerInstanceAutoscalingConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SpannerInstanceAutoscalingConfigOutputReference)SetInternalValue(val *SpannerInstanceAutoscalingConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SpannerInstanceAutoscalingConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SpannerInstanceAutoscalingConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SpannerInstanceAutoscalingConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpannerInstanceAutoscalingConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SpannerInstanceAutoscalingConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SpannerInstanceAutoscalingConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SpannerInstanceAutoscalingConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SpannerInstanceAutoscalingConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SpannerInstanceAutoscalingConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SpannerInstanceAutoscalingConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SpannerInstanceAutoscalingConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SpannerInstanceAutoscalingConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SpannerInstanceAutoscalingConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpannerInstanceAutoscalingConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SpannerInstanceAutoscalingConfigOutputReference) PutAsymmetricAutoscalingOptions(value interface{}) {
	if err := s.validatePutAsymmetricAutoscalingOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAsymmetricAutoscalingOptions",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpannerInstanceAutoscalingConfigOutputReference) PutAutoscalingLimits(value *SpannerInstanceAutoscalingConfigAutoscalingLimits) {
	if err := s.validatePutAutoscalingLimitsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAutoscalingLimits",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpannerInstanceAutoscalingConfigOutputReference) PutAutoscalingTargets(value *SpannerInstanceAutoscalingConfigAutoscalingTargets) {
	if err := s.validatePutAutoscalingTargetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAutoscalingTargets",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SpannerInstanceAutoscalingConfigOutputReference) ResetAsymmetricAutoscalingOptions() {
	_jsii_.InvokeVoid(
		s,
		"resetAsymmetricAutoscalingOptions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpannerInstanceAutoscalingConfigOutputReference) ResetAutoscalingLimits() {
	_jsii_.InvokeVoid(
		s,
		"resetAutoscalingLimits",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpannerInstanceAutoscalingConfigOutputReference) ResetAutoscalingTargets() {
	_jsii_.InvokeVoid(
		s,
		"resetAutoscalingTargets",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpannerInstanceAutoscalingConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SpannerInstanceAutoscalingConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

