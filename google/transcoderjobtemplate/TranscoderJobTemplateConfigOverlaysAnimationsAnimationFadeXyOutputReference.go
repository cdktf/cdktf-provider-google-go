// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transcoderjobtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/transcoderjobtemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXyOutputReference interface {
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
	InternalValue() *TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXy
	SetInternalValue(val *TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXy)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	X() *float64
	SetX(val *float64)
	XInput() *float64
	Y() *float64
	SetY(val *float64)
	YInput() *float64
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
	ResetX()
	ResetY()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXyOutputReference
type jsiiProxy_TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXyOutputReference) InternalValue() *TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXy {
	var returns *TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXyOutputReference) X() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"x",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXyOutputReference) XInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"xInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXyOutputReference) Y() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"y",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXyOutputReference) YInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"yInput",
		&returns,
	)
	return returns
}


func NewTranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXyOutputReference {
	_init_.Initialize()

	if err := validateNewTranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.transcoderJobTemplate.TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewTranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXyOutputReference_Override(t TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.transcoderJobTemplate.TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXyOutputReference)SetInternalValue(val *TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXyOutputReference)SetX(val *float64) {
	if err := j.validateSetXParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"x",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXyOutputReference)SetY(val *float64) {
	if err := j.validateSetYParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"y",
		val,
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (t *jsiiProxy_TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (t *jsiiProxy_TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXyOutputReference) ResetX() {
	_jsii_.InvokeVoid(
		t,
		"resetX",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXyOutputReference) ResetY() {
	_jsii_.InvokeVoid(
		t,
		"resetY",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (t *jsiiProxy_TranscoderJobTemplateConfigOverlaysAnimationsAnimationFadeXyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

