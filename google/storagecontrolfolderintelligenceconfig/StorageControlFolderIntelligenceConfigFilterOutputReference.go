// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagecontrolfolderintelligenceconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/storagecontrolfolderintelligenceconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageControlFolderIntelligenceConfigFilterOutputReference interface {
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
	ExcludedCloudStorageBuckets() StorageControlFolderIntelligenceConfigFilterExcludedCloudStorageBucketsOutputReference
	ExcludedCloudStorageBucketsInput() *StorageControlFolderIntelligenceConfigFilterExcludedCloudStorageBuckets
	ExcludedCloudStorageLocations() StorageControlFolderIntelligenceConfigFilterExcludedCloudStorageLocationsOutputReference
	ExcludedCloudStorageLocationsInput() *StorageControlFolderIntelligenceConfigFilterExcludedCloudStorageLocations
	// Experimental.
	Fqn() *string
	IncludedCloudStorageBuckets() StorageControlFolderIntelligenceConfigFilterIncludedCloudStorageBucketsOutputReference
	IncludedCloudStorageBucketsInput() *StorageControlFolderIntelligenceConfigFilterIncludedCloudStorageBuckets
	IncludedCloudStorageLocations() StorageControlFolderIntelligenceConfigFilterIncludedCloudStorageLocationsOutputReference
	IncludedCloudStorageLocationsInput() *StorageControlFolderIntelligenceConfigFilterIncludedCloudStorageLocations
	InternalValue() *StorageControlFolderIntelligenceConfigFilter
	SetInternalValue(val *StorageControlFolderIntelligenceConfigFilter)
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
	PutExcludedCloudStorageBuckets(value *StorageControlFolderIntelligenceConfigFilterExcludedCloudStorageBuckets)
	PutExcludedCloudStorageLocations(value *StorageControlFolderIntelligenceConfigFilterExcludedCloudStorageLocations)
	PutIncludedCloudStorageBuckets(value *StorageControlFolderIntelligenceConfigFilterIncludedCloudStorageBuckets)
	PutIncludedCloudStorageLocations(value *StorageControlFolderIntelligenceConfigFilterIncludedCloudStorageLocations)
	ResetExcludedCloudStorageBuckets()
	ResetExcludedCloudStorageLocations()
	ResetIncludedCloudStorageBuckets()
	ResetIncludedCloudStorageLocations()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageControlFolderIntelligenceConfigFilterOutputReference
type jsiiProxy_StorageControlFolderIntelligenceConfigFilterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageControlFolderIntelligenceConfigFilterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageControlFolderIntelligenceConfigFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageControlFolderIntelligenceConfigFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageControlFolderIntelligenceConfigFilterOutputReference) ExcludedCloudStorageBuckets() StorageControlFolderIntelligenceConfigFilterExcludedCloudStorageBucketsOutputReference {
	var returns StorageControlFolderIntelligenceConfigFilterExcludedCloudStorageBucketsOutputReference
	_jsii_.Get(
		j,
		"excludedCloudStorageBuckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageControlFolderIntelligenceConfigFilterOutputReference) ExcludedCloudStorageBucketsInput() *StorageControlFolderIntelligenceConfigFilterExcludedCloudStorageBuckets {
	var returns *StorageControlFolderIntelligenceConfigFilterExcludedCloudStorageBuckets
	_jsii_.Get(
		j,
		"excludedCloudStorageBucketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageControlFolderIntelligenceConfigFilterOutputReference) ExcludedCloudStorageLocations() StorageControlFolderIntelligenceConfigFilterExcludedCloudStorageLocationsOutputReference {
	var returns StorageControlFolderIntelligenceConfigFilterExcludedCloudStorageLocationsOutputReference
	_jsii_.Get(
		j,
		"excludedCloudStorageLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageControlFolderIntelligenceConfigFilterOutputReference) ExcludedCloudStorageLocationsInput() *StorageControlFolderIntelligenceConfigFilterExcludedCloudStorageLocations {
	var returns *StorageControlFolderIntelligenceConfigFilterExcludedCloudStorageLocations
	_jsii_.Get(
		j,
		"excludedCloudStorageLocationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageControlFolderIntelligenceConfigFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageControlFolderIntelligenceConfigFilterOutputReference) IncludedCloudStorageBuckets() StorageControlFolderIntelligenceConfigFilterIncludedCloudStorageBucketsOutputReference {
	var returns StorageControlFolderIntelligenceConfigFilterIncludedCloudStorageBucketsOutputReference
	_jsii_.Get(
		j,
		"includedCloudStorageBuckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageControlFolderIntelligenceConfigFilterOutputReference) IncludedCloudStorageBucketsInput() *StorageControlFolderIntelligenceConfigFilterIncludedCloudStorageBuckets {
	var returns *StorageControlFolderIntelligenceConfigFilterIncludedCloudStorageBuckets
	_jsii_.Get(
		j,
		"includedCloudStorageBucketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageControlFolderIntelligenceConfigFilterOutputReference) IncludedCloudStorageLocations() StorageControlFolderIntelligenceConfigFilterIncludedCloudStorageLocationsOutputReference {
	var returns StorageControlFolderIntelligenceConfigFilterIncludedCloudStorageLocationsOutputReference
	_jsii_.Get(
		j,
		"includedCloudStorageLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageControlFolderIntelligenceConfigFilterOutputReference) IncludedCloudStorageLocationsInput() *StorageControlFolderIntelligenceConfigFilterIncludedCloudStorageLocations {
	var returns *StorageControlFolderIntelligenceConfigFilterIncludedCloudStorageLocations
	_jsii_.Get(
		j,
		"includedCloudStorageLocationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageControlFolderIntelligenceConfigFilterOutputReference) InternalValue() *StorageControlFolderIntelligenceConfigFilter {
	var returns *StorageControlFolderIntelligenceConfigFilter
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageControlFolderIntelligenceConfigFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageControlFolderIntelligenceConfigFilterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewStorageControlFolderIntelligenceConfigFilterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageControlFolderIntelligenceConfigFilterOutputReference {
	_init_.Initialize()

	if err := validateNewStorageControlFolderIntelligenceConfigFilterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StorageControlFolderIntelligenceConfigFilterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.storageControlFolderIntelligenceConfig.StorageControlFolderIntelligenceConfigFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageControlFolderIntelligenceConfigFilterOutputReference_Override(s StorageControlFolderIntelligenceConfigFilterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.storageControlFolderIntelligenceConfig.StorageControlFolderIntelligenceConfigFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageControlFolderIntelligenceConfigFilterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageControlFolderIntelligenceConfigFilterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageControlFolderIntelligenceConfigFilterOutputReference)SetInternalValue(val *StorageControlFolderIntelligenceConfigFilter) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageControlFolderIntelligenceConfigFilterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageControlFolderIntelligenceConfigFilterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StorageControlFolderIntelligenceConfigFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageControlFolderIntelligenceConfigFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StorageControlFolderIntelligenceConfigFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StorageControlFolderIntelligenceConfigFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StorageControlFolderIntelligenceConfigFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StorageControlFolderIntelligenceConfigFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StorageControlFolderIntelligenceConfigFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StorageControlFolderIntelligenceConfigFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StorageControlFolderIntelligenceConfigFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StorageControlFolderIntelligenceConfigFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StorageControlFolderIntelligenceConfigFilterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageControlFolderIntelligenceConfigFilterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StorageControlFolderIntelligenceConfigFilterOutputReference) PutExcludedCloudStorageBuckets(value *StorageControlFolderIntelligenceConfigFilterExcludedCloudStorageBuckets) {
	if err := s.validatePutExcludedCloudStorageBucketsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putExcludedCloudStorageBuckets",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageControlFolderIntelligenceConfigFilterOutputReference) PutExcludedCloudStorageLocations(value *StorageControlFolderIntelligenceConfigFilterExcludedCloudStorageLocations) {
	if err := s.validatePutExcludedCloudStorageLocationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putExcludedCloudStorageLocations",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageControlFolderIntelligenceConfigFilterOutputReference) PutIncludedCloudStorageBuckets(value *StorageControlFolderIntelligenceConfigFilterIncludedCloudStorageBuckets) {
	if err := s.validatePutIncludedCloudStorageBucketsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putIncludedCloudStorageBuckets",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageControlFolderIntelligenceConfigFilterOutputReference) PutIncludedCloudStorageLocations(value *StorageControlFolderIntelligenceConfigFilterIncludedCloudStorageLocations) {
	if err := s.validatePutIncludedCloudStorageLocationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putIncludedCloudStorageLocations",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageControlFolderIntelligenceConfigFilterOutputReference) ResetExcludedCloudStorageBuckets() {
	_jsii_.InvokeVoid(
		s,
		"resetExcludedCloudStorageBuckets",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageControlFolderIntelligenceConfigFilterOutputReference) ResetExcludedCloudStorageLocations() {
	_jsii_.InvokeVoid(
		s,
		"resetExcludedCloudStorageLocations",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageControlFolderIntelligenceConfigFilterOutputReference) ResetIncludedCloudStorageBuckets() {
	_jsii_.InvokeVoid(
		s,
		"resetIncludedCloudStorageBuckets",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageControlFolderIntelligenceConfigFilterOutputReference) ResetIncludedCloudStorageLocations() {
	_jsii_.InvokeVoid(
		s,
		"resetIncludedCloudStorageLocations",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageControlFolderIntelligenceConfigFilterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_StorageControlFolderIntelligenceConfigFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

