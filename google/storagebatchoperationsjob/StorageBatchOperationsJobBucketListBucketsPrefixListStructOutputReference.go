// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagebatchoperationsjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/storagebatchoperationsjob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageBatchOperationsJobBucketListBucketsPrefixListStructOutputReference interface {
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
	IncludedObjectPrefixes() *[]*string
	SetIncludedObjectPrefixes(val *[]*string)
	IncludedObjectPrefixesInput() *[]*string
	InternalValue() *StorageBatchOperationsJobBucketListBucketsPrefixListStruct
	SetInternalValue(val *StorageBatchOperationsJobBucketListBucketsPrefixListStruct)
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
	ResetIncludedObjectPrefixes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageBatchOperationsJobBucketListBucketsPrefixListStructOutputReference
type jsiiProxy_StorageBatchOperationsJobBucketListBucketsPrefixListStructOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageBatchOperationsJobBucketListBucketsPrefixListStructOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJobBucketListBucketsPrefixListStructOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJobBucketListBucketsPrefixListStructOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJobBucketListBucketsPrefixListStructOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJobBucketListBucketsPrefixListStructOutputReference) IncludedObjectPrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedObjectPrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJobBucketListBucketsPrefixListStructOutputReference) IncludedObjectPrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedObjectPrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJobBucketListBucketsPrefixListStructOutputReference) InternalValue() *StorageBatchOperationsJobBucketListBucketsPrefixListStruct {
	var returns *StorageBatchOperationsJobBucketListBucketsPrefixListStruct
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJobBucketListBucketsPrefixListStructOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJobBucketListBucketsPrefixListStructOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewStorageBatchOperationsJobBucketListBucketsPrefixListStructOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageBatchOperationsJobBucketListBucketsPrefixListStructOutputReference {
	_init_.Initialize()

	if err := validateNewStorageBatchOperationsJobBucketListBucketsPrefixListStructOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StorageBatchOperationsJobBucketListBucketsPrefixListStructOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.storageBatchOperationsJob.StorageBatchOperationsJobBucketListBucketsPrefixListStructOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageBatchOperationsJobBucketListBucketsPrefixListStructOutputReference_Override(s StorageBatchOperationsJobBucketListBucketsPrefixListStructOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.storageBatchOperationsJob.StorageBatchOperationsJobBucketListBucketsPrefixListStructOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageBatchOperationsJobBucketListBucketsPrefixListStructOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageBatchOperationsJobBucketListBucketsPrefixListStructOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageBatchOperationsJobBucketListBucketsPrefixListStructOutputReference)SetIncludedObjectPrefixes(val *[]*string) {
	if err := j.validateSetIncludedObjectPrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includedObjectPrefixes",
		val,
	)
}

func (j *jsiiProxy_StorageBatchOperationsJobBucketListBucketsPrefixListStructOutputReference)SetInternalValue(val *StorageBatchOperationsJobBucketListBucketsPrefixListStruct) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageBatchOperationsJobBucketListBucketsPrefixListStructOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageBatchOperationsJobBucketListBucketsPrefixListStructOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StorageBatchOperationsJobBucketListBucketsPrefixListStructOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBatchOperationsJobBucketListBucketsPrefixListStructOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StorageBatchOperationsJobBucketListBucketsPrefixListStructOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StorageBatchOperationsJobBucketListBucketsPrefixListStructOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StorageBatchOperationsJobBucketListBucketsPrefixListStructOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StorageBatchOperationsJobBucketListBucketsPrefixListStructOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StorageBatchOperationsJobBucketListBucketsPrefixListStructOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StorageBatchOperationsJobBucketListBucketsPrefixListStructOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StorageBatchOperationsJobBucketListBucketsPrefixListStructOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StorageBatchOperationsJobBucketListBucketsPrefixListStructOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StorageBatchOperationsJobBucketListBucketsPrefixListStructOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBatchOperationsJobBucketListBucketsPrefixListStructOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StorageBatchOperationsJobBucketListBucketsPrefixListStructOutputReference) ResetIncludedObjectPrefixes() {
	_jsii_.InvokeVoid(
		s,
		"resetIncludedObjectPrefixes",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBatchOperationsJobBucketListBucketsPrefixListStructOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_StorageBatchOperationsJobBucketListBucketsPrefixListStructOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

