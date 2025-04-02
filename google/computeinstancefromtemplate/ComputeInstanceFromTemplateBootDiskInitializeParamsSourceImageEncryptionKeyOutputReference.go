// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeinstancefromtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/computeinstancefromtemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference interface {
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
	InternalValue() *ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKey
	SetInternalValue(val *ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKey)
	KmsKeySelfLink() *string
	SetKmsKeySelfLink(val *string)
	KmsKeySelfLinkInput() *string
	KmsKeyServiceAccount() *string
	SetKmsKeyServiceAccount(val *string)
	KmsKeyServiceAccountInput() *string
	RawKey() *string
	SetRawKey(val *string)
	RawKeyInput() *string
	RsaEncryptedKey() *string
	SetRsaEncryptedKey(val *string)
	RsaEncryptedKeyInput() *string
	Sha256() *string
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
	ResetKmsKeySelfLink()
	ResetKmsKeyServiceAccount()
	ResetRawKey()
	ResetRsaEncryptedKey()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference
type jsiiProxy_ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference) InternalValue() *ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKey {
	var returns *ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKey
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference) KmsKeySelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeySelfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference) KmsKeySelfLinkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeySelfLinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference) KmsKeyServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyServiceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference) KmsKeyServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyServiceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference) RawKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rawKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference) RawKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rawKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference) RsaEncryptedKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rsaEncryptedKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference) RsaEncryptedKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rsaEncryptedKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference) Sha256() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sha256",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference {
	_init_.Initialize()

	if err := validateNewComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeInstanceFromTemplate.ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference_Override(c ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeInstanceFromTemplate.ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference)SetInternalValue(val *ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKey) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference)SetKmsKeySelfLink(val *string) {
	if err := j.validateSetKmsKeySelfLinkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeySelfLink",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference)SetKmsKeyServiceAccount(val *string) {
	if err := j.validateSetKmsKeyServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyServiceAccount",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference)SetRawKey(val *string) {
	if err := j.validateSetRawKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rawKey",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference)SetRsaEncryptedKey(val *string) {
	if err := j.validateSetRsaEncryptedKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rsaEncryptedKey",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference) ResetKmsKeySelfLink() {
	_jsii_.InvokeVoid(
		c,
		"resetKmsKeySelfLink",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference) ResetKmsKeyServiceAccount() {
	_jsii_.InvokeVoid(
		c,
		"resetKmsKeyServiceAccount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference) ResetRawKey() {
	_jsii_.InvokeVoid(
		c,
		"resetRawKey",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference) ResetRsaEncryptedKey() {
	_jsii_.InvokeVoid(
		c,
		"resetRsaEncryptedKey",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

