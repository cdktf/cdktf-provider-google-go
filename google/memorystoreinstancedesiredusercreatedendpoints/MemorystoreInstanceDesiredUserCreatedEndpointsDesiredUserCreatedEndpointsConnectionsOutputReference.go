// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package memorystoreinstancedesiredusercreatedendpoints

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/memorystoreinstancedesiredusercreatedendpoints/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PscConnection() MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsPscConnectionOutputReference
	PscConnectionInput() *MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsPscConnection
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
	PutPscConnection(value *MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsPscConnection)
	ResetPscConnection()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsOutputReference
type jsiiProxy_MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsOutputReference) PscConnection() MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsPscConnectionOutputReference {
	var returns MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsPscConnectionOutputReference
	_jsii_.Get(
		j,
		"pscConnection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsOutputReference) PscConnectionInput() *MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsPscConnection {
	var returns *MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsPscConnection
	_jsii_.Get(
		j,
		"pscConnectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsOutputReference {
	_init_.Initialize()

	if err := validateNewMemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.memorystoreInstanceDesiredUserCreatedEndpoints.MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsOutputReference_Override(m MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.memorystoreInstanceDesiredUserCreatedEndpoints.MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsOutputReference) PutPscConnection(value *MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsPscConnection) {
	if err := m.validatePutPscConnectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putPscConnection",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsOutputReference) ResetPscConnection() {
	_jsii_.InvokeVoid(
		m,
		"resetPscConnection",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MemorystoreInstanceDesiredUserCreatedEndpointsDesiredUserCreatedEndpointsConnectionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

