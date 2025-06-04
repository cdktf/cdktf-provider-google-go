// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transcoderjobtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/transcoderjobtemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference interface {
	cdktf.ComplexObject
	Clearkey() TranscoderJobTemplateConfigEncryptionsDrmSystemsClearkeyOutputReference
	ClearkeyInput() *TranscoderJobTemplateConfigEncryptionsDrmSystemsClearkey
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
	Fairplay() TranscoderJobTemplateConfigEncryptionsDrmSystemsFairplayOutputReference
	FairplayInput() *TranscoderJobTemplateConfigEncryptionsDrmSystemsFairplay
	// Experimental.
	Fqn() *string
	InternalValue() *TranscoderJobTemplateConfigEncryptionsDrmSystems
	SetInternalValue(val *TranscoderJobTemplateConfigEncryptionsDrmSystems)
	Playready() TranscoderJobTemplateConfigEncryptionsDrmSystemsPlayreadyOutputReference
	PlayreadyInput() *TranscoderJobTemplateConfigEncryptionsDrmSystemsPlayready
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Widevine() TranscoderJobTemplateConfigEncryptionsDrmSystemsWidevineOutputReference
	WidevineInput() *TranscoderJobTemplateConfigEncryptionsDrmSystemsWidevine
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
	PutClearkey(value *TranscoderJobTemplateConfigEncryptionsDrmSystemsClearkey)
	PutFairplay(value *TranscoderJobTemplateConfigEncryptionsDrmSystemsFairplay)
	PutPlayready(value *TranscoderJobTemplateConfigEncryptionsDrmSystemsPlayready)
	PutWidevine(value *TranscoderJobTemplateConfigEncryptionsDrmSystemsWidevine)
	ResetClearkey()
	ResetFairplay()
	ResetPlayready()
	ResetWidevine()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference
type jsiiProxy_TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) Clearkey() TranscoderJobTemplateConfigEncryptionsDrmSystemsClearkeyOutputReference {
	var returns TranscoderJobTemplateConfigEncryptionsDrmSystemsClearkeyOutputReference
	_jsii_.Get(
		j,
		"clearkey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) ClearkeyInput() *TranscoderJobTemplateConfigEncryptionsDrmSystemsClearkey {
	var returns *TranscoderJobTemplateConfigEncryptionsDrmSystemsClearkey
	_jsii_.Get(
		j,
		"clearkeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) Fairplay() TranscoderJobTemplateConfigEncryptionsDrmSystemsFairplayOutputReference {
	var returns TranscoderJobTemplateConfigEncryptionsDrmSystemsFairplayOutputReference
	_jsii_.Get(
		j,
		"fairplay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) FairplayInput() *TranscoderJobTemplateConfigEncryptionsDrmSystemsFairplay {
	var returns *TranscoderJobTemplateConfigEncryptionsDrmSystemsFairplay
	_jsii_.Get(
		j,
		"fairplayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) InternalValue() *TranscoderJobTemplateConfigEncryptionsDrmSystems {
	var returns *TranscoderJobTemplateConfigEncryptionsDrmSystems
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) Playready() TranscoderJobTemplateConfigEncryptionsDrmSystemsPlayreadyOutputReference {
	var returns TranscoderJobTemplateConfigEncryptionsDrmSystemsPlayreadyOutputReference
	_jsii_.Get(
		j,
		"playready",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) PlayreadyInput() *TranscoderJobTemplateConfigEncryptionsDrmSystemsPlayready {
	var returns *TranscoderJobTemplateConfigEncryptionsDrmSystemsPlayready
	_jsii_.Get(
		j,
		"playreadyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) Widevine() TranscoderJobTemplateConfigEncryptionsDrmSystemsWidevineOutputReference {
	var returns TranscoderJobTemplateConfigEncryptionsDrmSystemsWidevineOutputReference
	_jsii_.Get(
		j,
		"widevine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) WidevineInput() *TranscoderJobTemplateConfigEncryptionsDrmSystemsWidevine {
	var returns *TranscoderJobTemplateConfigEncryptionsDrmSystemsWidevine
	_jsii_.Get(
		j,
		"widevineInput",
		&returns,
	)
	return returns
}


func NewTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference {
	_init_.Initialize()

	if err := validateNewTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.transcoderJobTemplate.TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewTranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference_Override(t TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.transcoderJobTemplate.TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference)SetInternalValue(val *TranscoderJobTemplateConfigEncryptionsDrmSystems) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) PutClearkey(value *TranscoderJobTemplateConfigEncryptionsDrmSystemsClearkey) {
	if err := t.validatePutClearkeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putClearkey",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) PutFairplay(value *TranscoderJobTemplateConfigEncryptionsDrmSystemsFairplay) {
	if err := t.validatePutFairplayParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFairplay",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) PutPlayready(value *TranscoderJobTemplateConfigEncryptionsDrmSystemsPlayready) {
	if err := t.validatePutPlayreadyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPlayready",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) PutWidevine(value *TranscoderJobTemplateConfigEncryptionsDrmSystemsWidevine) {
	if err := t.validatePutWidevineParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putWidevine",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) ResetClearkey() {
	_jsii_.InvokeVoid(
		t,
		"resetClearkey",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) ResetFairplay() {
	_jsii_.InvokeVoid(
		t,
		"resetFairplay",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) ResetPlayready() {
	_jsii_.InvokeVoid(
		t,
		"resetPlayready",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) ResetWidevine() {
	_jsii_.InvokeVoid(
		t,
		"resetWidevine",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (t *jsiiProxy_TranscoderJobTemplateConfigEncryptionsDrmSystemsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

