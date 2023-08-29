// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package binaryauthorizationattestor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v9/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v9/binaryauthorizationattestor/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference interface {
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
	DelegationServiceAccountEmail() *string
	// Experimental.
	Fqn() *string
	InternalValue() *BinaryAuthorizationAttestorAttestationAuthorityNote
	SetInternalValue(val *BinaryAuthorizationAttestorAttestationAuthorityNote)
	NoteReference() *string
	SetNoteReference(val *string)
	NoteReferenceInput() *string
	PublicKeys() BinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysList
	PublicKeysInput() interface{}
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
	PutPublicKeys(value interface{})
	ResetPublicKeys()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference
type jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) DelegationServiceAccountEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delegationServiceAccountEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) InternalValue() *BinaryAuthorizationAttestorAttestationAuthorityNote {
	var returns *BinaryAuthorizationAttestorAttestationAuthorityNote
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) NoteReference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noteReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) NoteReferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noteReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) PublicKeys() BinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysList {
	var returns BinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysList
	_jsii_.Get(
		j,
		"publicKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) PublicKeysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publicKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference {
	_init_.Initialize()

	if err := validateNewBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.binaryAuthorizationAttestor.BinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference_Override(b BinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.binaryAuthorizationAttestor.BinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference)SetInternalValue(val *BinaryAuthorizationAttestorAttestationAuthorityNote) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference)SetNoteReference(val *string) {
	if err := j.validateSetNoteReferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noteReference",
		val,
	)
}

func (j *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) PutPublicKeys(value interface{}) {
	if err := b.validatePutPublicKeysParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putPublicKeys",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) ResetPublicKeys() {
	_jsii_.InvokeVoid(
		b,
		"resetPublicKeys",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := b.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNoteOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

