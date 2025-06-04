// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagebatchoperationsjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/storagebatchoperationsjob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageBatchOperationsJobPutObjectHoldOutputReference interface {
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
	EventBasedHold() *string
	SetEventBasedHold(val *string)
	EventBasedHoldInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *StorageBatchOperationsJobPutObjectHold
	SetInternalValue(val *StorageBatchOperationsJobPutObjectHold)
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
	ResetEventBasedHold()
	ResetTemporaryHold()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageBatchOperationsJobPutObjectHoldOutputReference
type jsiiProxy_StorageBatchOperationsJobPutObjectHoldOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageBatchOperationsJobPutObjectHoldOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJobPutObjectHoldOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJobPutObjectHoldOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJobPutObjectHoldOutputReference) EventBasedHold() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventBasedHold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJobPutObjectHoldOutputReference) EventBasedHoldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventBasedHoldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJobPutObjectHoldOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJobPutObjectHoldOutputReference) InternalValue() *StorageBatchOperationsJobPutObjectHold {
	var returns *StorageBatchOperationsJobPutObjectHold
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJobPutObjectHoldOutputReference) TemporaryHold() *string {
	var returns *string
	_jsii_.Get(
		j,
		"temporaryHold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJobPutObjectHoldOutputReference) TemporaryHoldInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"temporaryHoldInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJobPutObjectHoldOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageBatchOperationsJobPutObjectHoldOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewStorageBatchOperationsJobPutObjectHoldOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageBatchOperationsJobPutObjectHoldOutputReference {
	_init_.Initialize()

	if err := validateNewStorageBatchOperationsJobPutObjectHoldOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StorageBatchOperationsJobPutObjectHoldOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.storageBatchOperationsJob.StorageBatchOperationsJobPutObjectHoldOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageBatchOperationsJobPutObjectHoldOutputReference_Override(s StorageBatchOperationsJobPutObjectHoldOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.storageBatchOperationsJob.StorageBatchOperationsJobPutObjectHoldOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageBatchOperationsJobPutObjectHoldOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageBatchOperationsJobPutObjectHoldOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageBatchOperationsJobPutObjectHoldOutputReference)SetEventBasedHold(val *string) {
	if err := j.validateSetEventBasedHoldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventBasedHold",
		val,
	)
}

func (j *jsiiProxy_StorageBatchOperationsJobPutObjectHoldOutputReference)SetInternalValue(val *StorageBatchOperationsJobPutObjectHold) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageBatchOperationsJobPutObjectHoldOutputReference)SetTemporaryHold(val *string) {
	if err := j.validateSetTemporaryHoldParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"temporaryHold",
		val,
	)
}

func (j *jsiiProxy_StorageBatchOperationsJobPutObjectHoldOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageBatchOperationsJobPutObjectHoldOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StorageBatchOperationsJobPutObjectHoldOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBatchOperationsJobPutObjectHoldOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StorageBatchOperationsJobPutObjectHoldOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StorageBatchOperationsJobPutObjectHoldOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StorageBatchOperationsJobPutObjectHoldOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StorageBatchOperationsJobPutObjectHoldOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StorageBatchOperationsJobPutObjectHoldOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StorageBatchOperationsJobPutObjectHoldOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StorageBatchOperationsJobPutObjectHoldOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StorageBatchOperationsJobPutObjectHoldOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StorageBatchOperationsJobPutObjectHoldOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageBatchOperationsJobPutObjectHoldOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StorageBatchOperationsJobPutObjectHoldOutputReference) ResetEventBasedHold() {
	_jsii_.InvokeVoid(
		s,
		"resetEventBasedHold",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBatchOperationsJobPutObjectHoldOutputReference) ResetTemporaryHold() {
	_jsii_.InvokeVoid(
		s,
		"resetTemporaryHold",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageBatchOperationsJobPutObjectHoldOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_StorageBatchOperationsJobPutObjectHoldOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

