// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package memorystoreinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/memorystoreinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MemorystoreInstanceCrossInstanceReplicationConfigOutputReference interface {
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
	InstanceRole() *string
	SetInstanceRole(val *string)
	InstanceRoleInput() *string
	InternalValue() *MemorystoreInstanceCrossInstanceReplicationConfig
	SetInternalValue(val *MemorystoreInstanceCrossInstanceReplicationConfig)
	Membership() MemorystoreInstanceCrossInstanceReplicationConfigMembershipList
	PrimaryInstance() MemorystoreInstanceCrossInstanceReplicationConfigPrimaryInstanceOutputReference
	PrimaryInstanceInput() *MemorystoreInstanceCrossInstanceReplicationConfigPrimaryInstance
	SecondaryInstances() MemorystoreInstanceCrossInstanceReplicationConfigSecondaryInstancesList
	SecondaryInstancesInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UpdateTime() *string
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
	PutPrimaryInstance(value *MemorystoreInstanceCrossInstanceReplicationConfigPrimaryInstance)
	PutSecondaryInstances(value interface{})
	ResetInstanceRole()
	ResetPrimaryInstance()
	ResetSecondaryInstances()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MemorystoreInstanceCrossInstanceReplicationConfigOutputReference
type jsiiProxy_MemorystoreInstanceCrossInstanceReplicationConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MemorystoreInstanceCrossInstanceReplicationConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstanceCrossInstanceReplicationConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstanceCrossInstanceReplicationConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstanceCrossInstanceReplicationConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstanceCrossInstanceReplicationConfigOutputReference) InstanceRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstanceCrossInstanceReplicationConfigOutputReference) InstanceRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstanceCrossInstanceReplicationConfigOutputReference) InternalValue() *MemorystoreInstanceCrossInstanceReplicationConfig {
	var returns *MemorystoreInstanceCrossInstanceReplicationConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstanceCrossInstanceReplicationConfigOutputReference) Membership() MemorystoreInstanceCrossInstanceReplicationConfigMembershipList {
	var returns MemorystoreInstanceCrossInstanceReplicationConfigMembershipList
	_jsii_.Get(
		j,
		"membership",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstanceCrossInstanceReplicationConfigOutputReference) PrimaryInstance() MemorystoreInstanceCrossInstanceReplicationConfigPrimaryInstanceOutputReference {
	var returns MemorystoreInstanceCrossInstanceReplicationConfigPrimaryInstanceOutputReference
	_jsii_.Get(
		j,
		"primaryInstance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstanceCrossInstanceReplicationConfigOutputReference) PrimaryInstanceInput() *MemorystoreInstanceCrossInstanceReplicationConfigPrimaryInstance {
	var returns *MemorystoreInstanceCrossInstanceReplicationConfigPrimaryInstance
	_jsii_.Get(
		j,
		"primaryInstanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstanceCrossInstanceReplicationConfigOutputReference) SecondaryInstances() MemorystoreInstanceCrossInstanceReplicationConfigSecondaryInstancesList {
	var returns MemorystoreInstanceCrossInstanceReplicationConfigSecondaryInstancesList
	_jsii_.Get(
		j,
		"secondaryInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstanceCrossInstanceReplicationConfigOutputReference) SecondaryInstancesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secondaryInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstanceCrossInstanceReplicationConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstanceCrossInstanceReplicationConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstanceCrossInstanceReplicationConfigOutputReference) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


func NewMemorystoreInstanceCrossInstanceReplicationConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MemorystoreInstanceCrossInstanceReplicationConfigOutputReference {
	_init_.Initialize()

	if err := validateNewMemorystoreInstanceCrossInstanceReplicationConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MemorystoreInstanceCrossInstanceReplicationConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.memorystoreInstance.MemorystoreInstanceCrossInstanceReplicationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMemorystoreInstanceCrossInstanceReplicationConfigOutputReference_Override(m MemorystoreInstanceCrossInstanceReplicationConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.memorystoreInstance.MemorystoreInstanceCrossInstanceReplicationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MemorystoreInstanceCrossInstanceReplicationConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MemorystoreInstanceCrossInstanceReplicationConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MemorystoreInstanceCrossInstanceReplicationConfigOutputReference)SetInstanceRole(val *string) {
	if err := j.validateSetInstanceRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceRole",
		val,
	)
}

func (j *jsiiProxy_MemorystoreInstanceCrossInstanceReplicationConfigOutputReference)SetInternalValue(val *MemorystoreInstanceCrossInstanceReplicationConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MemorystoreInstanceCrossInstanceReplicationConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MemorystoreInstanceCrossInstanceReplicationConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MemorystoreInstanceCrossInstanceReplicationConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorystoreInstanceCrossInstanceReplicationConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorystoreInstanceCrossInstanceReplicationConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorystoreInstanceCrossInstanceReplicationConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorystoreInstanceCrossInstanceReplicationConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorystoreInstanceCrossInstanceReplicationConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorystoreInstanceCrossInstanceReplicationConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorystoreInstanceCrossInstanceReplicationConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorystoreInstanceCrossInstanceReplicationConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorystoreInstanceCrossInstanceReplicationConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorystoreInstanceCrossInstanceReplicationConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorystoreInstanceCrossInstanceReplicationConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorystoreInstanceCrossInstanceReplicationConfigOutputReference) PutPrimaryInstance(value *MemorystoreInstanceCrossInstanceReplicationConfigPrimaryInstance) {
	if err := m.validatePutPrimaryInstanceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putPrimaryInstance",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MemorystoreInstanceCrossInstanceReplicationConfigOutputReference) PutSecondaryInstances(value interface{}) {
	if err := m.validatePutSecondaryInstancesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putSecondaryInstances",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MemorystoreInstanceCrossInstanceReplicationConfigOutputReference) ResetInstanceRole() {
	_jsii_.InvokeVoid(
		m,
		"resetInstanceRole",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorystoreInstanceCrossInstanceReplicationConfigOutputReference) ResetPrimaryInstance() {
	_jsii_.InvokeVoid(
		m,
		"resetPrimaryInstance",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorystoreInstanceCrossInstanceReplicationConfigOutputReference) ResetSecondaryInstances() {
	_jsii_.InvokeVoid(
		m,
		"resetSecondaryInstances",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorystoreInstanceCrossInstanceReplicationConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorystoreInstanceCrossInstanceReplicationConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

