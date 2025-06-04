// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transcoderjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/transcoderjob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference interface {
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
	EndTimeOffset() *string
	SetEndTimeOffset(val *string)
	EndTimeOffsetInput() *string
	FadeType() *string
	SetFadeType(val *string)
	FadeTypeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *TranscoderJobConfigOverlaysAnimationsAnimationFade
	SetInternalValue(val *TranscoderJobConfigOverlaysAnimationsAnimationFade)
	StartTimeOffset() *string
	SetStartTimeOffset(val *string)
	StartTimeOffsetInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Xy() TranscoderJobConfigOverlaysAnimationsAnimationFadeXyOutputReference
	XyInput() *TranscoderJobConfigOverlaysAnimationsAnimationFadeXy
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
	PutXy(value *TranscoderJobConfigOverlaysAnimationsAnimationFadeXy)
	ResetEndTimeOffset()
	ResetStartTimeOffset()
	ResetXy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference
type jsiiProxy_TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) EndTimeOffset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTimeOffset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) EndTimeOffsetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTimeOffsetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) FadeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fadeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) FadeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fadeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) InternalValue() *TranscoderJobConfigOverlaysAnimationsAnimationFade {
	var returns *TranscoderJobConfigOverlaysAnimationsAnimationFade
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) StartTimeOffset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTimeOffset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) StartTimeOffsetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTimeOffsetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) Xy() TranscoderJobConfigOverlaysAnimationsAnimationFadeXyOutputReference {
	var returns TranscoderJobConfigOverlaysAnimationsAnimationFadeXyOutputReference
	_jsii_.Get(
		j,
		"xy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) XyInput() *TranscoderJobConfigOverlaysAnimationsAnimationFadeXy {
	var returns *TranscoderJobConfigOverlaysAnimationsAnimationFadeXy
	_jsii_.Get(
		j,
		"xyInput",
		&returns,
	)
	return returns
}


func NewTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference {
	_init_.Initialize()

	if err := validateNewTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.transcoderJob.TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewTranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference_Override(t TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.transcoderJob.TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference)SetEndTimeOffset(val *string) {
	if err := j.validateSetEndTimeOffsetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endTimeOffset",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference)SetFadeType(val *string) {
	if err := j.validateSetFadeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fadeType",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference)SetInternalValue(val *TranscoderJobConfigOverlaysAnimationsAnimationFade) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference)SetStartTimeOffset(val *string) {
	if err := j.validateSetStartTimeOffsetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startTimeOffset",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) PutXy(value *TranscoderJobConfigOverlaysAnimationsAnimationFadeXy) {
	if err := t.validatePutXyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putXy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) ResetEndTimeOffset() {
	_jsii_.InvokeVoid(
		t,
		"resetEndTimeOffset",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) ResetStartTimeOffset() {
	_jsii_.InvokeVoid(
		t,
		"resetStartTimeOffset",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) ResetXy() {
	_jsii_.InvokeVoid(
		t,
		"resetXy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := t.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobConfigOverlaysAnimationsAnimationFadeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

