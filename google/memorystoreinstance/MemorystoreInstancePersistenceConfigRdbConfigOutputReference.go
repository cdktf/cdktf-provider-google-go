// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package memorystoreinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/memorystoreinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MemorystoreInstancePersistenceConfigRdbConfigOutputReference interface {
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
	InternalValue() *MemorystoreInstancePersistenceConfigRdbConfig
	SetInternalValue(val *MemorystoreInstancePersistenceConfigRdbConfig)
	RdbSnapshotPeriod() *string
	SetRdbSnapshotPeriod(val *string)
	RdbSnapshotPeriodInput() *string
	RdbSnapshotStartTime() *string
	SetRdbSnapshotStartTime(val *string)
	RdbSnapshotStartTimeInput() *string
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
	ResetRdbSnapshotPeriod()
	ResetRdbSnapshotStartTime()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MemorystoreInstancePersistenceConfigRdbConfigOutputReference
type jsiiProxy_MemorystoreInstancePersistenceConfigRdbConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MemorystoreInstancePersistenceConfigRdbConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstancePersistenceConfigRdbConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstancePersistenceConfigRdbConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstancePersistenceConfigRdbConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstancePersistenceConfigRdbConfigOutputReference) InternalValue() *MemorystoreInstancePersistenceConfigRdbConfig {
	var returns *MemorystoreInstancePersistenceConfigRdbConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstancePersistenceConfigRdbConfigOutputReference) RdbSnapshotPeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdbSnapshotPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstancePersistenceConfigRdbConfigOutputReference) RdbSnapshotPeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdbSnapshotPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstancePersistenceConfigRdbConfigOutputReference) RdbSnapshotStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdbSnapshotStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstancePersistenceConfigRdbConfigOutputReference) RdbSnapshotStartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdbSnapshotStartTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstancePersistenceConfigRdbConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstancePersistenceConfigRdbConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMemorystoreInstancePersistenceConfigRdbConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MemorystoreInstancePersistenceConfigRdbConfigOutputReference {
	_init_.Initialize()

	if err := validateNewMemorystoreInstancePersistenceConfigRdbConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MemorystoreInstancePersistenceConfigRdbConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.memorystoreInstance.MemorystoreInstancePersistenceConfigRdbConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMemorystoreInstancePersistenceConfigRdbConfigOutputReference_Override(m MemorystoreInstancePersistenceConfigRdbConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.memorystoreInstance.MemorystoreInstancePersistenceConfigRdbConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MemorystoreInstancePersistenceConfigRdbConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MemorystoreInstancePersistenceConfigRdbConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MemorystoreInstancePersistenceConfigRdbConfigOutputReference)SetInternalValue(val *MemorystoreInstancePersistenceConfigRdbConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MemorystoreInstancePersistenceConfigRdbConfigOutputReference)SetRdbSnapshotPeriod(val *string) {
	if err := j.validateSetRdbSnapshotPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rdbSnapshotPeriod",
		val,
	)
}

func (j *jsiiProxy_MemorystoreInstancePersistenceConfigRdbConfigOutputReference)SetRdbSnapshotStartTime(val *string) {
	if err := j.validateSetRdbSnapshotStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rdbSnapshotStartTime",
		val,
	)
}

func (j *jsiiProxy_MemorystoreInstancePersistenceConfigRdbConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MemorystoreInstancePersistenceConfigRdbConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MemorystoreInstancePersistenceConfigRdbConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorystoreInstancePersistenceConfigRdbConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MemorystoreInstancePersistenceConfigRdbConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MemorystoreInstancePersistenceConfigRdbConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MemorystoreInstancePersistenceConfigRdbConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MemorystoreInstancePersistenceConfigRdbConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MemorystoreInstancePersistenceConfigRdbConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MemorystoreInstancePersistenceConfigRdbConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MemorystoreInstancePersistenceConfigRdbConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MemorystoreInstancePersistenceConfigRdbConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MemorystoreInstancePersistenceConfigRdbConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorystoreInstancePersistenceConfigRdbConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MemorystoreInstancePersistenceConfigRdbConfigOutputReference) ResetRdbSnapshotPeriod() {
	_jsii_.InvokeVoid(
		m,
		"resetRdbSnapshotPeriod",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorystoreInstancePersistenceConfigRdbConfigOutputReference) ResetRdbSnapshotStartTime() {
	_jsii_.InvokeVoid(
		m,
		"resetRdbSnapshotStartTime",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorystoreInstancePersistenceConfigRdbConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MemorystoreInstancePersistenceConfigRdbConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

