// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeinstancefromtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/computeinstancefromtemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeInstanceFromTemplateBootDiskOutputReference interface {
	cdktf.ComplexObject
	AutoDelete() interface{}
	SetAutoDelete(val interface{})
	AutoDeleteInput() interface{}
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
	DeviceName() *string
	SetDeviceName(val *string)
	DeviceNameInput() *string
	DiskEncryptionKeyRaw() *string
	SetDiskEncryptionKeyRaw(val *string)
	DiskEncryptionKeyRawInput() *string
	DiskEncryptionKeyRsa() *string
	SetDiskEncryptionKeyRsa(val *string)
	DiskEncryptionKeyRsaInput() *string
	DiskEncryptionKeySha256() *string
	DiskEncryptionServiceAccount() *string
	SetDiskEncryptionServiceAccount(val *string)
	DiskEncryptionServiceAccountInput() *string
	// Experimental.
	Fqn() *string
	GuestOsFeatures() *[]*string
	SetGuestOsFeatures(val *[]*string)
	GuestOsFeaturesInput() *[]*string
	InitializeParams() ComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference
	InitializeParamsInput() *ComputeInstanceFromTemplateBootDiskInitializeParams
	Interface() *string
	SetInterface(val *string)
	InterfaceInput() *string
	InternalValue() *ComputeInstanceFromTemplateBootDisk
	SetInternalValue(val *ComputeInstanceFromTemplateBootDisk)
	KmsKeySelfLink() *string
	SetKmsKeySelfLink(val *string)
	KmsKeySelfLinkInput() *string
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	Source() *string
	SetSource(val *string)
	SourceInput() *string
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
	PutInitializeParams(value *ComputeInstanceFromTemplateBootDiskInitializeParams)
	ResetAutoDelete()
	ResetDeviceName()
	ResetDiskEncryptionKeyRaw()
	ResetDiskEncryptionKeyRsa()
	ResetDiskEncryptionServiceAccount()
	ResetGuestOsFeatures()
	ResetInitializeParams()
	ResetInterface()
	ResetKmsKeySelfLink()
	ResetMode()
	ResetSource()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeInstanceFromTemplateBootDiskOutputReference
type jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) AutoDelete() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) AutoDeleteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) DeviceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) DeviceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) DiskEncryptionKeyRaw() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryptionKeyRaw",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) DiskEncryptionKeyRawInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryptionKeyRawInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) DiskEncryptionKeyRsa() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryptionKeyRsa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) DiskEncryptionKeyRsaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryptionKeyRsaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) DiskEncryptionKeySha256() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryptionKeySha256",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) DiskEncryptionServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryptionServiceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) DiskEncryptionServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryptionServiceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) GuestOsFeatures() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"guestOsFeatures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) GuestOsFeaturesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"guestOsFeaturesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) InitializeParams() ComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference {
	var returns ComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference
	_jsii_.Get(
		j,
		"initializeParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) InitializeParamsInput() *ComputeInstanceFromTemplateBootDiskInitializeParams {
	var returns *ComputeInstanceFromTemplateBootDiskInitializeParams
	_jsii_.Get(
		j,
		"initializeParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) Interface() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) InterfaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interfaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) InternalValue() *ComputeInstanceFromTemplateBootDisk {
	var returns *ComputeInstanceFromTemplateBootDisk
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) KmsKeySelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeySelfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) KmsKeySelfLinkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeySelfLinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) SourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeInstanceFromTemplateBootDiskOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeInstanceFromTemplateBootDiskOutputReference {
	_init_.Initialize()

	if err := validateNewComputeInstanceFromTemplateBootDiskOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeInstanceFromTemplate.ComputeInstanceFromTemplateBootDiskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeInstanceFromTemplateBootDiskOutputReference_Override(c ComputeInstanceFromTemplateBootDiskOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeInstanceFromTemplate.ComputeInstanceFromTemplateBootDiskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference)SetAutoDelete(val interface{}) {
	if err := j.validateSetAutoDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoDelete",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference)SetDeviceName(val *string) {
	if err := j.validateSetDeviceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceName",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference)SetDiskEncryptionKeyRaw(val *string) {
	if err := j.validateSetDiskEncryptionKeyRawParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskEncryptionKeyRaw",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference)SetDiskEncryptionKeyRsa(val *string) {
	if err := j.validateSetDiskEncryptionKeyRsaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskEncryptionKeyRsa",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference)SetDiskEncryptionServiceAccount(val *string) {
	if err := j.validateSetDiskEncryptionServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskEncryptionServiceAccount",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference)SetGuestOsFeatures(val *[]*string) {
	if err := j.validateSetGuestOsFeaturesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"guestOsFeatures",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference)SetInterface(val *string) {
	if err := j.validateSetInterfaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interface",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference)SetInternalValue(val *ComputeInstanceFromTemplateBootDisk) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference)SetKmsKeySelfLink(val *string) {
	if err := j.validateSetKmsKeySelfLinkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeySelfLink",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference)SetSource(val *string) {
	if err := j.validateSetSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) PutInitializeParams(value *ComputeInstanceFromTemplateBootDiskInitializeParams) {
	if err := c.validatePutInitializeParamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putInitializeParams",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) ResetAutoDelete() {
	_jsii_.InvokeVoid(
		c,
		"resetAutoDelete",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) ResetDeviceName() {
	_jsii_.InvokeVoid(
		c,
		"resetDeviceName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) ResetDiskEncryptionKeyRaw() {
	_jsii_.InvokeVoid(
		c,
		"resetDiskEncryptionKeyRaw",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) ResetDiskEncryptionKeyRsa() {
	_jsii_.InvokeVoid(
		c,
		"resetDiskEncryptionKeyRsa",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) ResetDiskEncryptionServiceAccount() {
	_jsii_.InvokeVoid(
		c,
		"resetDiskEncryptionServiceAccount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) ResetGuestOsFeatures() {
	_jsii_.InvokeVoid(
		c,
		"resetGuestOsFeatures",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) ResetInitializeParams() {
	_jsii_.InvokeVoid(
		c,
		"resetInitializeParams",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) ResetInterface() {
	_jsii_.InvokeVoid(
		c,
		"resetInterface",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) ResetKmsKeySelfLink() {
	_jsii_.InvokeVoid(
		c,
		"resetKmsKeySelfLink",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) ResetMode() {
	_jsii_.InvokeVoid(
		c,
		"resetMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) ResetSource() {
	_jsii_.InvokeVoid(
		c,
		"resetSource",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeInstanceFromTemplateBootDiskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

