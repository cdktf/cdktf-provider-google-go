// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagetransferjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/storagetransferjob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference interface {
	cdktf.ComplexObject
	Acl() *string
	SetAcl(val *string)
	AclInput() *string
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
	Gid() *string
	SetGid(val *string)
	GidInput() *string
	InternalValue() *StorageTransferJobTransferSpecTransferOptionsMetadataOptions
	SetInternalValue(val *StorageTransferJobTransferSpecTransferOptionsMetadataOptions)
	KmsKey() *string
	SetKmsKey(val *string)
	KmsKeyInput() *string
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	StorageClass() *string
	SetStorageClass(val *string)
	StorageClassInput() *string
	Symlink() *string
	SetSymlink(val *string)
	SymlinkInput() *string
	TemporaryHold() *string
	SetTemporaryHold(val *string)
	TemporaryHoldInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeCreated() *string
	SetTimeCreated(val *string)
	TimeCreatedInput() *string
	Uid() *string
	SetUid(val *string)
	UidInput() *string
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
	ResetAcl()
	ResetGid()
	ResetKmsKey()
	ResetMode()
	ResetStorageClass()
	ResetSymlink()
	ResetTemporaryHold()
	ResetTimeCreated()
	ResetUid()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference
type jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) Acl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) AclInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aclInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) Gid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) GidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) InternalValue() *StorageTransferJobTransferSpecTransferOptionsMetadataOptions {
	var returns *StorageTransferJobTransferSpecTransferOptionsMetadataOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) KmsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) KmsKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) StorageClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) StorageClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) Symlink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"symlink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) SymlinkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"symlinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) TemporaryHold() *string {
	var returns *string
	_jsii_.Get(
		j,
		"temporaryHold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) TemporaryHoldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"temporaryHoldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) TimeCreated() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeCreated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) TimeCreatedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeCreatedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) Uid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) UidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uidInput",
		&returns,
	)
	return returns
}


func NewStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.storageTransferJob.StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference_Override(s StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.storageTransferJob.StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference)SetAcl(val *string) {
	if err := j.validateSetAclParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acl",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference)SetGid(val *string) {
	if err := j.validateSetGidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gid",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference)SetInternalValue(val *StorageTransferJobTransferSpecTransferOptionsMetadataOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference)SetKmsKey(val *string) {
	if err := j.validateSetKmsKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKey",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference)SetStorageClass(val *string) {
	if err := j.validateSetStorageClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageClass",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference)SetSymlink(val *string) {
	if err := j.validateSetSymlinkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"symlink",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference)SetTemporaryHold(val *string) {
	if err := j.validateSetTemporaryHoldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"temporaryHold",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference)SetTimeCreated(val *string) {
	if err := j.validateSetTimeCreatedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeCreated",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference)SetUid(val *string) {
	if err := j.validateSetUidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uid",
		val,
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) ResetAcl() {
	_jsii_.InvokeVoid(
		s,
		"resetAcl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) ResetGid() {
	_jsii_.InvokeVoid(
		s,
		"resetGid",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) ResetKmsKey() {
	_jsii_.InvokeVoid(
		s,
		"resetKmsKey",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) ResetMode() {
	_jsii_.InvokeVoid(
		s,
		"resetMode",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) ResetStorageClass() {
	_jsii_.InvokeVoid(
		s,
		"resetStorageClass",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) ResetSymlink() {
	_jsii_.InvokeVoid(
		s,
		"resetSymlink",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) ResetTemporaryHold() {
	_jsii_.InvokeVoid(
		s,
		"resetTemporaryHold",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) ResetTimeCreated() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeCreated",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) ResetUid() {
	_jsii_.InvokeVoid(
		s,
		"resetUid",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_StorageTransferJobTransferSpecTransferOptionsMetadataOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

