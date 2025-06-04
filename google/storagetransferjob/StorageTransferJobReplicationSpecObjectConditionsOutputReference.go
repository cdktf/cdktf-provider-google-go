// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagetransferjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/storagetransferjob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageTransferJobReplicationSpecObjectConditionsOutputReference interface {
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
	ExcludePrefixes() *[]*string
	SetExcludePrefixes(val *[]*string)
	ExcludePrefixesInput() *[]*string
	// Experimental.
	Fqn() *string
	IncludePrefixes() *[]*string
	SetIncludePrefixes(val *[]*string)
	IncludePrefixesInput() *[]*string
	InternalValue() *StorageTransferJobReplicationSpecObjectConditions
	SetInternalValue(val *StorageTransferJobReplicationSpecObjectConditions)
	LastModifiedBefore() *string
	SetLastModifiedBefore(val *string)
	LastModifiedBeforeInput() *string
	LastModifiedSince() *string
	SetLastModifiedSince(val *string)
	LastModifiedSinceInput() *string
	MaxTimeElapsedSinceLastModification() *string
	SetMaxTimeElapsedSinceLastModification(val *string)
	MaxTimeElapsedSinceLastModificationInput() *string
	MinTimeElapsedSinceLastModification() *string
	SetMinTimeElapsedSinceLastModification(val *string)
	MinTimeElapsedSinceLastModificationInput() *string
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
	ResetExcludePrefixes()
	ResetIncludePrefixes()
	ResetLastModifiedBefore()
	ResetLastModifiedSince()
	ResetMaxTimeElapsedSinceLastModification()
	ResetMinTimeElapsedSinceLastModification()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageTransferJobReplicationSpecObjectConditionsOutputReference
type jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference) ExcludePrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludePrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference) ExcludePrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludePrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference) IncludePrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includePrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference) IncludePrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includePrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference) InternalValue() *StorageTransferJobReplicationSpecObjectConditions {
	var returns *StorageTransferJobReplicationSpecObjectConditions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference) LastModifiedBefore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedBefore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference) LastModifiedBeforeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedBeforeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference) LastModifiedSince() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedSince",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference) LastModifiedSinceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedSinceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference) MaxTimeElapsedSinceLastModification() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxTimeElapsedSinceLastModification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference) MaxTimeElapsedSinceLastModificationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxTimeElapsedSinceLastModificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference) MinTimeElapsedSinceLastModification() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minTimeElapsedSinceLastModification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference) MinTimeElapsedSinceLastModificationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minTimeElapsedSinceLastModificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewStorageTransferJobReplicationSpecObjectConditionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageTransferJobReplicationSpecObjectConditionsOutputReference {
	_init_.Initialize()

	if err := validateNewStorageTransferJobReplicationSpecObjectConditionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.storageTransferJob.StorageTransferJobReplicationSpecObjectConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageTransferJobReplicationSpecObjectConditionsOutputReference_Override(s StorageTransferJobReplicationSpecObjectConditionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.storageTransferJob.StorageTransferJobReplicationSpecObjectConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference)SetExcludePrefixes(val *[]*string) {
	if err := j.validateSetExcludePrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludePrefixes",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference)SetIncludePrefixes(val *[]*string) {
	if err := j.validateSetIncludePrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includePrefixes",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference)SetInternalValue(val *StorageTransferJobReplicationSpecObjectConditions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference)SetLastModifiedBefore(val *string) {
	if err := j.validateSetLastModifiedBeforeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lastModifiedBefore",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference)SetLastModifiedSince(val *string) {
	if err := j.validateSetLastModifiedSinceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lastModifiedSince",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference)SetMaxTimeElapsedSinceLastModification(val *string) {
	if err := j.validateSetMaxTimeElapsedSinceLastModificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxTimeElapsedSinceLastModification",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference)SetMinTimeElapsedSinceLastModification(val *string) {
	if err := j.validateSetMinTimeElapsedSinceLastModificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minTimeElapsedSinceLastModification",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference) ResetExcludePrefixes() {
	_jsii_.InvokeVoid(
		s,
		"resetExcludePrefixes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference) ResetIncludePrefixes() {
	_jsii_.InvokeVoid(
		s,
		"resetIncludePrefixes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference) ResetLastModifiedBefore() {
	_jsii_.InvokeVoid(
		s,
		"resetLastModifiedBefore",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference) ResetLastModifiedSince() {
	_jsii_.InvokeVoid(
		s,
		"resetLastModifiedSince",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference) ResetMaxTimeElapsedSinceLastModification() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxTimeElapsedSinceLastModification",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference) ResetMinTimeElapsedSinceLastModification() {
	_jsii_.InvokeVoid(
		s,
		"resetMinTimeElapsedSinceLastModification",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_StorageTransferJobReplicationSpecObjectConditionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

