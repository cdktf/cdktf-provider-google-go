// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagetransferjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/storagetransferjob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageTransferJobLoggingConfigOutputReference interface {
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
	EnableOnPremGcsTransferLogs() interface{}
	SetEnableOnPremGcsTransferLogs(val interface{})
	EnableOnPremGcsTransferLogsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *StorageTransferJobLoggingConfig
	SetInternalValue(val *StorageTransferJobLoggingConfig)
	LogActions() *[]*string
	SetLogActions(val *[]*string)
	LogActionsInput() *[]*string
	LogActionStates() *[]*string
	SetLogActionStates(val *[]*string)
	LogActionStatesInput() *[]*string
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
	ResetEnableOnPremGcsTransferLogs()
	ResetLogActions()
	ResetLogActionStates()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for StorageTransferJobLoggingConfigOutputReference
type jsiiProxy_StorageTransferJobLoggingConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_StorageTransferJobLoggingConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobLoggingConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobLoggingConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobLoggingConfigOutputReference) EnableOnPremGcsTransferLogs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableOnPremGcsTransferLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobLoggingConfigOutputReference) EnableOnPremGcsTransferLogsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableOnPremGcsTransferLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobLoggingConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobLoggingConfigOutputReference) InternalValue() *StorageTransferJobLoggingConfig {
	var returns *StorageTransferJobLoggingConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobLoggingConfigOutputReference) LogActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"logActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobLoggingConfigOutputReference) LogActionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"logActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobLoggingConfigOutputReference) LogActionStates() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"logActionStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobLoggingConfigOutputReference) LogActionStatesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"logActionStatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobLoggingConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageTransferJobLoggingConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewStorageTransferJobLoggingConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) StorageTransferJobLoggingConfigOutputReference {
	_init_.Initialize()

	if err := validateNewStorageTransferJobLoggingConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_StorageTransferJobLoggingConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.storageTransferJob.StorageTransferJobLoggingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewStorageTransferJobLoggingConfigOutputReference_Override(s StorageTransferJobLoggingConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.storageTransferJob.StorageTransferJobLoggingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_StorageTransferJobLoggingConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobLoggingConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobLoggingConfigOutputReference)SetEnableOnPremGcsTransferLogs(val interface{}) {
	if err := j.validateSetEnableOnPremGcsTransferLogsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableOnPremGcsTransferLogs",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobLoggingConfigOutputReference)SetInternalValue(val *StorageTransferJobLoggingConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobLoggingConfigOutputReference)SetLogActions(val *[]*string) {
	if err := j.validateSetLogActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logActions",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobLoggingConfigOutputReference)SetLogActionStates(val *[]*string) {
	if err := j.validateSetLogActionStatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logActionStates",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobLoggingConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_StorageTransferJobLoggingConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_StorageTransferJobLoggingConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageTransferJobLoggingConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StorageTransferJobLoggingConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StorageTransferJobLoggingConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StorageTransferJobLoggingConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StorageTransferJobLoggingConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StorageTransferJobLoggingConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StorageTransferJobLoggingConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StorageTransferJobLoggingConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StorageTransferJobLoggingConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StorageTransferJobLoggingConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageTransferJobLoggingConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StorageTransferJobLoggingConfigOutputReference) ResetEnableOnPremGcsTransferLogs() {
	_jsii_.InvokeVoid(
		s,
		"resetEnableOnPremGcsTransferLogs",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobLoggingConfigOutputReference) ResetLogActions() {
	_jsii_.InvokeVoid(
		s,
		"resetLogActions",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobLoggingConfigOutputReference) ResetLogActionStates() {
	_jsii_.InvokeVoid(
		s,
		"resetLogActionStates",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageTransferJobLoggingConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_StorageTransferJobLoggingConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

