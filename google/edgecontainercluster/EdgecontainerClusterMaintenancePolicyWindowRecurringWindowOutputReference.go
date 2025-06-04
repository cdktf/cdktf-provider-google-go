// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package edgecontainercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/edgecontainercluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EdgecontainerClusterMaintenancePolicyWindowRecurringWindowOutputReference interface {
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
	InternalValue() *EdgecontainerClusterMaintenancePolicyWindowRecurringWindow
	SetInternalValue(val *EdgecontainerClusterMaintenancePolicyWindowRecurringWindow)
	Recurrence() *string
	SetRecurrence(val *string)
	RecurrenceInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Window() EdgecontainerClusterMaintenancePolicyWindowRecurringWindowWindowOutputReference
	WindowInput() *EdgecontainerClusterMaintenancePolicyWindowRecurringWindowWindow
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
	PutWindow(value *EdgecontainerClusterMaintenancePolicyWindowRecurringWindowWindow)
	ResetRecurrence()
	ResetWindow()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EdgecontainerClusterMaintenancePolicyWindowRecurringWindowOutputReference
type jsiiProxy_EdgecontainerClusterMaintenancePolicyWindowRecurringWindowOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EdgecontainerClusterMaintenancePolicyWindowRecurringWindowOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerClusterMaintenancePolicyWindowRecurringWindowOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerClusterMaintenancePolicyWindowRecurringWindowOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerClusterMaintenancePolicyWindowRecurringWindowOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerClusterMaintenancePolicyWindowRecurringWindowOutputReference) InternalValue() *EdgecontainerClusterMaintenancePolicyWindowRecurringWindow {
	var returns *EdgecontainerClusterMaintenancePolicyWindowRecurringWindow
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerClusterMaintenancePolicyWindowRecurringWindowOutputReference) Recurrence() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recurrence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerClusterMaintenancePolicyWindowRecurringWindowOutputReference) RecurrenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recurrenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerClusterMaintenancePolicyWindowRecurringWindowOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerClusterMaintenancePolicyWindowRecurringWindowOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerClusterMaintenancePolicyWindowRecurringWindowOutputReference) Window() EdgecontainerClusterMaintenancePolicyWindowRecurringWindowWindowOutputReference {
	var returns EdgecontainerClusterMaintenancePolicyWindowRecurringWindowWindowOutputReference
	_jsii_.Get(
		j,
		"window",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerClusterMaintenancePolicyWindowRecurringWindowOutputReference) WindowInput() *EdgecontainerClusterMaintenancePolicyWindowRecurringWindowWindow {
	var returns *EdgecontainerClusterMaintenancePolicyWindowRecurringWindowWindow
	_jsii_.Get(
		j,
		"windowInput",
		&returns,
	)
	return returns
}


func NewEdgecontainerClusterMaintenancePolicyWindowRecurringWindowOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EdgecontainerClusterMaintenancePolicyWindowRecurringWindowOutputReference {
	_init_.Initialize()

	if err := validateNewEdgecontainerClusterMaintenancePolicyWindowRecurringWindowOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EdgecontainerClusterMaintenancePolicyWindowRecurringWindowOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.edgecontainerCluster.EdgecontainerClusterMaintenancePolicyWindowRecurringWindowOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEdgecontainerClusterMaintenancePolicyWindowRecurringWindowOutputReference_Override(e EdgecontainerClusterMaintenancePolicyWindowRecurringWindowOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.edgecontainerCluster.EdgecontainerClusterMaintenancePolicyWindowRecurringWindowOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EdgecontainerClusterMaintenancePolicyWindowRecurringWindowOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EdgecontainerClusterMaintenancePolicyWindowRecurringWindowOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EdgecontainerClusterMaintenancePolicyWindowRecurringWindowOutputReference)SetInternalValue(val *EdgecontainerClusterMaintenancePolicyWindowRecurringWindow) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EdgecontainerClusterMaintenancePolicyWindowRecurringWindowOutputReference)SetRecurrence(val *string) {
	if err := j.validateSetRecurrenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recurrence",
		val,
	)
}

func (j *jsiiProxy_EdgecontainerClusterMaintenancePolicyWindowRecurringWindowOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EdgecontainerClusterMaintenancePolicyWindowRecurringWindowOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EdgecontainerClusterMaintenancePolicyWindowRecurringWindowOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgecontainerClusterMaintenancePolicyWindowRecurringWindowOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgecontainerClusterMaintenancePolicyWindowRecurringWindowOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgecontainerClusterMaintenancePolicyWindowRecurringWindowOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgecontainerClusterMaintenancePolicyWindowRecurringWindowOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgecontainerClusterMaintenancePolicyWindowRecurringWindowOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgecontainerClusterMaintenancePolicyWindowRecurringWindowOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgecontainerClusterMaintenancePolicyWindowRecurringWindowOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgecontainerClusterMaintenancePolicyWindowRecurringWindowOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgecontainerClusterMaintenancePolicyWindowRecurringWindowOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgecontainerClusterMaintenancePolicyWindowRecurringWindowOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgecontainerClusterMaintenancePolicyWindowRecurringWindowOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgecontainerClusterMaintenancePolicyWindowRecurringWindowOutputReference) PutWindow(value *EdgecontainerClusterMaintenancePolicyWindowRecurringWindowWindow) {
	if err := e.validatePutWindowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putWindow",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EdgecontainerClusterMaintenancePolicyWindowRecurringWindowOutputReference) ResetRecurrence() {
	_jsii_.InvokeVoid(
		e,
		"resetRecurrence",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgecontainerClusterMaintenancePolicyWindowRecurringWindowOutputReference) ResetWindow() {
	_jsii_.InvokeVoid(
		e,
		"resetWindow",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgecontainerClusterMaintenancePolicyWindowRecurringWindowOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgecontainerClusterMaintenancePolicyWindowRecurringWindowOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

