// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageinsightsreportconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/storageinsightsreportconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageInsightsReportConfigObjectMetadataReportOptionsOutputReference interface {
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
	InternalValue() *StorageInsightsReportConfigObjectMetadataReportOptions
	SetInternalValue(val *StorageInsightsReportConfigObjectMetadataReportOptions)
	MetadataFields() *[]*string
	SetMetadataFields(val *[]*string)
	MetadataFieldsInput() *[]*string
	StorageDestinationOptions() StorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference
	StorageDestinationOptionsInput() *StorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptions
	StorageFilters() StorageInsightsReportConfigObjectMetadataReportOptionsStorageFiltersOutputReference
	StorageFiltersInput() *StorageInsightsReportConfigObjectMetadataReportOptionsStorageFilters
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
	PutStorageDestinationOptions(value *StorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptions)
	PutStorageFilters(value *StorageInsightsReportConfigObjectMetadataReportOptionsStorageFilters)
	ResetStorageFilters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageInsightsReportConfigObjectMetadataReportOptionsOutputReference
type jsiiProxy_StorageInsightsReportConfigObjectMetadataReportOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) InternalValue() *StorageInsightsReportConfigObjectMetadataReportOptions {
	var returns *StorageInsightsReportConfigObjectMetadataReportOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) MetadataFields() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"metadataFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) MetadataFieldsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"metadataFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) StorageDestinationOptions() StorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference {
	var returns StorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference
	_jsii_.Get(
		j,
		"storageDestinationOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) StorageDestinationOptionsInput() *StorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptions {
	var returns *StorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptions
	_jsii_.Get(
		j,
		"storageDestinationOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) StorageFilters() StorageInsightsReportConfigObjectMetadataReportOptionsStorageFiltersOutputReference {
	var returns StorageInsightsReportConfigObjectMetadataReportOptionsStorageFiltersOutputReference
	_jsii_.Get(
		j,
		"storageFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) StorageFiltersInput() *StorageInsightsReportConfigObjectMetadataReportOptionsStorageFilters {
	var returns *StorageInsightsReportConfigObjectMetadataReportOptionsStorageFilters
	_jsii_.Get(
		j,
		"storageFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageInsightsReportConfigObjectMetadataReportOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewStorageInsightsReportConfigObjectMetadataReportOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StorageInsightsReportConfigObjectMetadataReportOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.storageInsightsReportConfig.StorageInsightsReportConfigObjectMetadataReportOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageInsightsReportConfigObjectMetadataReportOptionsOutputReference_Override(s StorageInsightsReportConfigObjectMetadataReportOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.storageInsightsReportConfig.StorageInsightsReportConfigObjectMetadataReportOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageInsightsReportConfigObjectMetadataReportOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageInsightsReportConfigObjectMetadataReportOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageInsightsReportConfigObjectMetadataReportOptionsOutputReference)SetInternalValue(val *StorageInsightsReportConfigObjectMetadataReportOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageInsightsReportConfigObjectMetadataReportOptionsOutputReference)SetMetadataFields(val *[]*string) {
	if err := j.validateSetMetadataFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadataFields",
		val,
	)
}

func (j *jsiiProxy_StorageInsightsReportConfigObjectMetadataReportOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageInsightsReportConfigObjectMetadataReportOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) PutStorageDestinationOptions(value *StorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptions) {
	if err := s.validatePutStorageDestinationOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putStorageDestinationOptions",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) PutStorageFilters(value *StorageInsightsReportConfigObjectMetadataReportOptionsStorageFilters) {
	if err := s.validatePutStorageFiltersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putStorageFilters",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) ResetStorageFilters() {
	_jsii_.InvokeVoid(
		s,
		"resetStorageFilters",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_StorageInsightsReportConfigObjectMetadataReportOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

