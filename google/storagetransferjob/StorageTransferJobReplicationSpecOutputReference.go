// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagetransferjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/storagetransferjob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageTransferJobReplicationSpecOutputReference interface {
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
	GcsDataSink() StorageTransferJobReplicationSpecGcsDataSinkOutputReference
	GcsDataSinkInput() *StorageTransferJobReplicationSpecGcsDataSink
	GcsDataSource() StorageTransferJobReplicationSpecGcsDataSourceOutputReference
	GcsDataSourceInput() *StorageTransferJobReplicationSpecGcsDataSource
	InternalValue() *StorageTransferJobReplicationSpec
	SetInternalValue(val *StorageTransferJobReplicationSpec)
	ObjectConditions() StorageTransferJobReplicationSpecObjectConditionsOutputReference
	ObjectConditionsInput() *StorageTransferJobReplicationSpecObjectConditions
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TransferOptions() StorageTransferJobReplicationSpecTransferOptionsOutputReference
	TransferOptionsInput() *StorageTransferJobReplicationSpecTransferOptions
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
	PutGcsDataSink(value *StorageTransferJobReplicationSpecGcsDataSink)
	PutGcsDataSource(value *StorageTransferJobReplicationSpecGcsDataSource)
	PutObjectConditions(value *StorageTransferJobReplicationSpecObjectConditions)
	PutTransferOptions(value *StorageTransferJobReplicationSpecTransferOptions)
	ResetGcsDataSink()
	ResetGcsDataSource()
	ResetObjectConditions()
	ResetTransferOptions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageTransferJobReplicationSpecOutputReference
type jsiiProxy_StorageTransferJobReplicationSpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecOutputReference) GcsDataSink() StorageTransferJobReplicationSpecGcsDataSinkOutputReference {
	var returns StorageTransferJobReplicationSpecGcsDataSinkOutputReference
	_jsii_.Get(
		j,
		"gcsDataSink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecOutputReference) GcsDataSinkInput() *StorageTransferJobReplicationSpecGcsDataSink {
	var returns *StorageTransferJobReplicationSpecGcsDataSink
	_jsii_.Get(
		j,
		"gcsDataSinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecOutputReference) GcsDataSource() StorageTransferJobReplicationSpecGcsDataSourceOutputReference {
	var returns StorageTransferJobReplicationSpecGcsDataSourceOutputReference
	_jsii_.Get(
		j,
		"gcsDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecOutputReference) GcsDataSourceInput() *StorageTransferJobReplicationSpecGcsDataSource {
	var returns *StorageTransferJobReplicationSpecGcsDataSource
	_jsii_.Get(
		j,
		"gcsDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecOutputReference) InternalValue() *StorageTransferJobReplicationSpec {
	var returns *StorageTransferJobReplicationSpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecOutputReference) ObjectConditions() StorageTransferJobReplicationSpecObjectConditionsOutputReference {
	var returns StorageTransferJobReplicationSpecObjectConditionsOutputReference
	_jsii_.Get(
		j,
		"objectConditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecOutputReference) ObjectConditionsInput() *StorageTransferJobReplicationSpecObjectConditions {
	var returns *StorageTransferJobReplicationSpecObjectConditions
	_jsii_.Get(
		j,
		"objectConditionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecOutputReference) TransferOptions() StorageTransferJobReplicationSpecTransferOptionsOutputReference {
	var returns StorageTransferJobReplicationSpecTransferOptionsOutputReference
	_jsii_.Get(
		j,
		"transferOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecOutputReference) TransferOptionsInput() *StorageTransferJobReplicationSpecTransferOptions {
	var returns *StorageTransferJobReplicationSpecTransferOptions
	_jsii_.Get(
		j,
		"transferOptionsInput",
		&returns,
	)
	return returns
}


func NewStorageTransferJobReplicationSpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageTransferJobReplicationSpecOutputReference {
	_init_.Initialize()

	if err := validateNewStorageTransferJobReplicationSpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StorageTransferJobReplicationSpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.storageTransferJob.StorageTransferJobReplicationSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageTransferJobReplicationSpecOutputReference_Override(s StorageTransferJobReplicationSpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.storageTransferJob.StorageTransferJobReplicationSpecOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecOutputReference)SetInternalValue(val *StorageTransferJobReplicationSpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobReplicationSpecOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StorageTransferJobReplicationSpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageTransferJobReplicationSpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StorageTransferJobReplicationSpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StorageTransferJobReplicationSpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StorageTransferJobReplicationSpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StorageTransferJobReplicationSpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StorageTransferJobReplicationSpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StorageTransferJobReplicationSpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StorageTransferJobReplicationSpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StorageTransferJobReplicationSpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StorageTransferJobReplicationSpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageTransferJobReplicationSpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StorageTransferJobReplicationSpecOutputReference) PutGcsDataSink(value *StorageTransferJobReplicationSpecGcsDataSink) {
	if err := s.validatePutGcsDataSinkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putGcsDataSink",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageTransferJobReplicationSpecOutputReference) PutGcsDataSource(value *StorageTransferJobReplicationSpecGcsDataSource) {
	if err := s.validatePutGcsDataSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putGcsDataSource",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageTransferJobReplicationSpecOutputReference) PutObjectConditions(value *StorageTransferJobReplicationSpecObjectConditions) {
	if err := s.validatePutObjectConditionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putObjectConditions",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageTransferJobReplicationSpecOutputReference) PutTransferOptions(value *StorageTransferJobReplicationSpecTransferOptions) {
	if err := s.validatePutTransferOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTransferOptions",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageTransferJobReplicationSpecOutputReference) ResetGcsDataSink() {
	_jsii_.InvokeVoid(
		s,
		"resetGcsDataSink",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobReplicationSpecOutputReference) ResetGcsDataSource() {
	_jsii_.InvokeVoid(
		s,
		"resetGcsDataSource",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobReplicationSpecOutputReference) ResetObjectConditions() {
	_jsii_.InvokeVoid(
		s,
		"resetObjectConditions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobReplicationSpecOutputReference) ResetTransferOptions() {
	_jsii_.InvokeVoid(
		s,
		"resetTransferOptions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobReplicationSpecOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_StorageTransferJobReplicationSpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

