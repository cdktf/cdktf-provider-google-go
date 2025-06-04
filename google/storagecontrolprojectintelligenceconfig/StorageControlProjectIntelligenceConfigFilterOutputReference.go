// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagecontrolprojectintelligenceconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/storagecontrolprojectintelligenceconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageControlProjectIntelligenceConfigFilterOutputReference interface {
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
	ExcludedCloudStorageBuckets() StorageControlProjectIntelligenceConfigFilterExcludedCloudStorageBucketsOutputReference
	ExcludedCloudStorageBucketsInput() *StorageControlProjectIntelligenceConfigFilterExcludedCloudStorageBuckets
	ExcludedCloudStorageLocations() StorageControlProjectIntelligenceConfigFilterExcludedCloudStorageLocationsOutputReference
	ExcludedCloudStorageLocationsInput() *StorageControlProjectIntelligenceConfigFilterExcludedCloudStorageLocations
	// Experimental.
	Fqn() *string
	IncludedCloudStorageBuckets() StorageControlProjectIntelligenceConfigFilterIncludedCloudStorageBucketsOutputReference
	IncludedCloudStorageBucketsInput() *StorageControlProjectIntelligenceConfigFilterIncludedCloudStorageBuckets
	IncludedCloudStorageLocations() StorageControlProjectIntelligenceConfigFilterIncludedCloudStorageLocationsOutputReference
	IncludedCloudStorageLocationsInput() *StorageControlProjectIntelligenceConfigFilterIncludedCloudStorageLocations
	InternalValue() *StorageControlProjectIntelligenceConfigFilter
	SetInternalValue(val *StorageControlProjectIntelligenceConfigFilter)
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
	PutExcludedCloudStorageBuckets(value *StorageControlProjectIntelligenceConfigFilterExcludedCloudStorageBuckets)
	PutExcludedCloudStorageLocations(value *StorageControlProjectIntelligenceConfigFilterExcludedCloudStorageLocations)
	PutIncludedCloudStorageBuckets(value *StorageControlProjectIntelligenceConfigFilterIncludedCloudStorageBuckets)
	PutIncludedCloudStorageLocations(value *StorageControlProjectIntelligenceConfigFilterIncludedCloudStorageLocations)
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

// The jsii proxy struct for StorageControlProjectIntelligenceConfigFilterOutputReference
type jsiiProxy_StorageControlProjectIntelligenceConfigFilterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageControlProjectIntelligenceConfigFilterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageControlProjectIntelligenceConfigFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageControlProjectIntelligenceConfigFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageControlProjectIntelligenceConfigFilterOutputReference) ExcludedCloudStorageBuckets() StorageControlProjectIntelligenceConfigFilterExcludedCloudStorageBucketsOutputReference {
	var returns StorageControlProjectIntelligenceConfigFilterExcludedCloudStorageBucketsOutputReference
	_jsii_.Get(
		j,
		"excludedCloudStorageBuckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageControlProjectIntelligenceConfigFilterOutputReference) ExcludedCloudStorageBucketsInput() *StorageControlProjectIntelligenceConfigFilterExcludedCloudStorageBuckets {
	var returns *StorageControlProjectIntelligenceConfigFilterExcludedCloudStorageBuckets
	_jsii_.Get(
		j,
		"excludedCloudStorageBucketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageControlProjectIntelligenceConfigFilterOutputReference) ExcludedCloudStorageLocations() StorageControlProjectIntelligenceConfigFilterExcludedCloudStorageLocationsOutputReference {
	var returns StorageControlProjectIntelligenceConfigFilterExcludedCloudStorageLocationsOutputReference
	_jsii_.Get(
		j,
		"excludedCloudStorageLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageControlProjectIntelligenceConfigFilterOutputReference) ExcludedCloudStorageLocationsInput() *StorageControlProjectIntelligenceConfigFilterExcludedCloudStorageLocations {
	var returns *StorageControlProjectIntelligenceConfigFilterExcludedCloudStorageLocations
	_jsii_.Get(
		j,
		"excludedCloudStorageLocationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageControlProjectIntelligenceConfigFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageControlProjectIntelligenceConfigFilterOutputReference) IncludedCloudStorageBuckets() StorageControlProjectIntelligenceConfigFilterIncludedCloudStorageBucketsOutputReference {
	var returns StorageControlProjectIntelligenceConfigFilterIncludedCloudStorageBucketsOutputReference
	_jsii_.Get(
		j,
		"includedCloudStorageBuckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageControlProjectIntelligenceConfigFilterOutputReference) IncludedCloudStorageBucketsInput() *StorageControlProjectIntelligenceConfigFilterIncludedCloudStorageBuckets {
	var returns *StorageControlProjectIntelligenceConfigFilterIncludedCloudStorageBuckets
	_jsii_.Get(
		j,
		"includedCloudStorageBucketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageControlProjectIntelligenceConfigFilterOutputReference) IncludedCloudStorageLocations() StorageControlProjectIntelligenceConfigFilterIncludedCloudStorageLocationsOutputReference {
	var returns StorageControlProjectIntelligenceConfigFilterIncludedCloudStorageLocationsOutputReference
	_jsii_.Get(
		j,
		"includedCloudStorageLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageControlProjectIntelligenceConfigFilterOutputReference) IncludedCloudStorageLocationsInput() *StorageControlProjectIntelligenceConfigFilterIncludedCloudStorageLocations {
	var returns *StorageControlProjectIntelligenceConfigFilterIncludedCloudStorageLocations
	_jsii_.Get(
		j,
		"includedCloudStorageLocationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageControlProjectIntelligenceConfigFilterOutputReference) InternalValue() *StorageControlProjectIntelligenceConfigFilter {
	var returns *StorageControlProjectIntelligenceConfigFilter
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageControlProjectIntelligenceConfigFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageControlProjectIntelligenceConfigFilterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewStorageControlProjectIntelligenceConfigFilterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageControlProjectIntelligenceConfigFilterOutputReference {
	_init_.Initialize()

	if err := validateNewStorageControlProjectIntelligenceConfigFilterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StorageControlProjectIntelligenceConfigFilterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.storageControlProjectIntelligenceConfig.StorageControlProjectIntelligenceConfigFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageControlProjectIntelligenceConfigFilterOutputReference_Override(s StorageControlProjectIntelligenceConfigFilterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.storageControlProjectIntelligenceConfig.StorageControlProjectIntelligenceConfigFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageControlProjectIntelligenceConfigFilterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageControlProjectIntelligenceConfigFilterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageControlProjectIntelligenceConfigFilterOutputReference)SetInternalValue(val *StorageControlProjectIntelligenceConfigFilter) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageControlProjectIntelligenceConfigFilterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageControlProjectIntelligenceConfigFilterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StorageControlProjectIntelligenceConfigFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageControlProjectIntelligenceConfigFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StorageControlProjectIntelligenceConfigFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StorageControlProjectIntelligenceConfigFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StorageControlProjectIntelligenceConfigFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StorageControlProjectIntelligenceConfigFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StorageControlProjectIntelligenceConfigFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StorageControlProjectIntelligenceConfigFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StorageControlProjectIntelligenceConfigFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StorageControlProjectIntelligenceConfigFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StorageControlProjectIntelligenceConfigFilterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageControlProjectIntelligenceConfigFilterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StorageControlProjectIntelligenceConfigFilterOutputReference) PutExcludedCloudStorageBuckets(value *StorageControlProjectIntelligenceConfigFilterExcludedCloudStorageBuckets) {
	if err := s.validatePutExcludedCloudStorageBucketsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putExcludedCloudStorageBuckets",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageControlProjectIntelligenceConfigFilterOutputReference) PutExcludedCloudStorageLocations(value *StorageControlProjectIntelligenceConfigFilterExcludedCloudStorageLocations) {
	if err := s.validatePutExcludedCloudStorageLocationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putExcludedCloudStorageLocations",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageControlProjectIntelligenceConfigFilterOutputReference) PutIncludedCloudStorageBuckets(value *StorageControlProjectIntelligenceConfigFilterIncludedCloudStorageBuckets) {
	if err := s.validatePutIncludedCloudStorageBucketsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putIncludedCloudStorageBuckets",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageControlProjectIntelligenceConfigFilterOutputReference) PutIncludedCloudStorageLocations(value *StorageControlProjectIntelligenceConfigFilterIncludedCloudStorageLocations) {
	if err := s.validatePutIncludedCloudStorageLocationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putIncludedCloudStorageLocations",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageControlProjectIntelligenceConfigFilterOutputReference) ResetExcludedCloudStorageBuckets() {
	_jsii_.InvokeVoid(
		s,
		"resetExcludedCloudStorageBuckets",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageControlProjectIntelligenceConfigFilterOutputReference) ResetExcludedCloudStorageLocations() {
	_jsii_.InvokeVoid(
		s,
		"resetExcludedCloudStorageLocations",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageControlProjectIntelligenceConfigFilterOutputReference) ResetIncludedCloudStorageBuckets() {
	_jsii_.InvokeVoid(
		s,
		"resetIncludedCloudStorageBuckets",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageControlProjectIntelligenceConfigFilterOutputReference) ResetIncludedCloudStorageLocations() {
	_jsii_.InvokeVoid(
		s,
		"resetIncludedCloudStorageLocations",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageControlProjectIntelligenceConfigFilterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_StorageControlProjectIntelligenceConfigFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

