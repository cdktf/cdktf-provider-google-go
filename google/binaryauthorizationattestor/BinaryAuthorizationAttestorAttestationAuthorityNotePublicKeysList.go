package binaryauthorizationattestor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v7/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v7/binaryauthorizationattestor/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) BinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysList
type jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) BinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysList {
	_init_.Initialize()

	if err := validateNewBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysList{}

	_jsii_.Create(
		"@cdktf/provider-google.binaryAuthorizationAttestor.BinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewBinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysList_Override(b BinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.binaryAuthorizationAttestor.BinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		b,
	)
}

func (j *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (b *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysList) Get(index *float64) BinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference {
	if err := b.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns BinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysOutputReference

	_jsii_.Invoke(
		b,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (b *jsiiProxy_BinaryAuthorizationAttestorAttestationAuthorityNotePublicKeysList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

