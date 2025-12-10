// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagetransferjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/storagetransferjob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageTransferJobTransferSpecObjectConditionsOutputReference interface {
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
	InternalValue() *StorageTransferJobTransferSpecObjectConditions
	SetInternalValue(val *StorageTransferJobTransferSpecObjectConditions)
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	ResetExcludePrefixes()
	ResetIncludePrefixes()
	ResetLastModifiedBefore()
	ResetLastModifiedSince()
	ResetMaxTimeElapsedSinceLastModification()
	ResetMinTimeElapsedSinceLastModification()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageTransferJobTransferSpecObjectConditionsOutputReference
type jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference) ExcludePrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludePrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference) ExcludePrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludePrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference) IncludePrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includePrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference) IncludePrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includePrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference) InternalValue() *StorageTransferJobTransferSpecObjectConditions {
	var returns *StorageTransferJobTransferSpecObjectConditions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference) LastModifiedBefore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedBefore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference) LastModifiedBeforeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedBeforeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference) LastModifiedSince() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedSince",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference) LastModifiedSinceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedSinceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference) MaxTimeElapsedSinceLastModification() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxTimeElapsedSinceLastModification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference) MaxTimeElapsedSinceLastModificationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxTimeElapsedSinceLastModificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference) MinTimeElapsedSinceLastModification() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minTimeElapsedSinceLastModification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference) MinTimeElapsedSinceLastModificationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minTimeElapsedSinceLastModificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewStorageTransferJobTransferSpecObjectConditionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageTransferJobTransferSpecObjectConditionsOutputReference {
	_init_.Initialize()

	if err := validateNewStorageTransferJobTransferSpecObjectConditionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.storageTransferJob.StorageTransferJobTransferSpecObjectConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageTransferJobTransferSpecObjectConditionsOutputReference_Override(s StorageTransferJobTransferSpecObjectConditionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.storageTransferJob.StorageTransferJobTransferSpecObjectConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference)SetExcludePrefixes(val *[]*string) {
	if err := j.validateSetExcludePrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludePrefixes",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference)SetIncludePrefixes(val *[]*string) {
	if err := j.validateSetIncludePrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includePrefixes",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference)SetInternalValue(val *StorageTransferJobTransferSpecObjectConditions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference)SetLastModifiedBefore(val *string) {
	if err := j.validateSetLastModifiedBeforeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lastModifiedBefore",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference)SetLastModifiedSince(val *string) {
	if err := j.validateSetLastModifiedSinceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lastModifiedSince",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference)SetMaxTimeElapsedSinceLastModification(val *string) {
	if err := j.validateSetMaxTimeElapsedSinceLastModificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxTimeElapsedSinceLastModification",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference)SetMinTimeElapsedSinceLastModification(val *string) {
	if err := j.validateSetMinTimeElapsedSinceLastModificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minTimeElapsedSinceLastModification",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference) ResetExcludePrefixes() {
	_jsii_.InvokeVoid(
		s,
		"resetExcludePrefixes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference) ResetIncludePrefixes() {
	_jsii_.InvokeVoid(
		s,
		"resetIncludePrefixes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference) ResetLastModifiedBefore() {
	_jsii_.InvokeVoid(
		s,
		"resetLastModifiedBefore",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference) ResetLastModifiedSince() {
	_jsii_.InvokeVoid(
		s,
		"resetLastModifiedSince",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference) ResetMaxTimeElapsedSinceLastModification() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxTimeElapsedSinceLastModification",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference) ResetMinTimeElapsedSinceLastModification() {
	_jsii_.InvokeVoid(
		s,
		"resetMinTimeElapsedSinceLastModification",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageTransferJobTransferSpecObjectConditionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

