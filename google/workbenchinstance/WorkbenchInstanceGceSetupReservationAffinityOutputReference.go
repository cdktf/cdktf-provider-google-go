// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workbenchinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/workbenchinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkbenchInstanceGceSetupReservationAffinityOutputReference interface {
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
	ConsumeReservationType() *string
	SetConsumeReservationType(val *string)
	ConsumeReservationTypeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *WorkbenchInstanceGceSetupReservationAffinity
	SetInternalValue(val *WorkbenchInstanceGceSetupReservationAffinity)
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Values() *[]*string
	SetValues(val *[]*string)
	ValuesInput() *[]*string
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
	ResetConsumeReservationType()
	ResetKey()
	ResetValues()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WorkbenchInstanceGceSetupReservationAffinityOutputReference
type jsiiProxy_WorkbenchInstanceGceSetupReservationAffinityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupReservationAffinityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupReservationAffinityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupReservationAffinityOutputReference) ConsumeReservationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumeReservationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupReservationAffinityOutputReference) ConsumeReservationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumeReservationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupReservationAffinityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupReservationAffinityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupReservationAffinityOutputReference) InternalValue() *WorkbenchInstanceGceSetupReservationAffinity {
	var returns *WorkbenchInstanceGceSetupReservationAffinity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupReservationAffinityOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupReservationAffinityOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupReservationAffinityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupReservationAffinityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupReservationAffinityOutputReference) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupReservationAffinityOutputReference) ValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewWorkbenchInstanceGceSetupReservationAffinityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WorkbenchInstanceGceSetupReservationAffinityOutputReference {
	_init_.Initialize()

	if err := validateNewWorkbenchInstanceGceSetupReservationAffinityOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkbenchInstanceGceSetupReservationAffinityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.workbenchInstance.WorkbenchInstanceGceSetupReservationAffinityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWorkbenchInstanceGceSetupReservationAffinityOutputReference_Override(w WorkbenchInstanceGceSetupReservationAffinityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.workbenchInstance.WorkbenchInstanceGceSetupReservationAffinityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupReservationAffinityOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupReservationAffinityOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupReservationAffinityOutputReference)SetConsumeReservationType(val *string) {
	if err := j.validateSetConsumeReservationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consumeReservationType",
		val,
	)
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupReservationAffinityOutputReference)SetInternalValue(val *WorkbenchInstanceGceSetupReservationAffinity) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupReservationAffinityOutputReference)SetKey(val *string) {
	if err := j.validateSetKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupReservationAffinityOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupReservationAffinityOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WorkbenchInstanceGceSetupReservationAffinityOutputReference)SetValues(val *[]*string) {
	if err := j.validateSetValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupReservationAffinityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupReservationAffinityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupReservationAffinityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupReservationAffinityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupReservationAffinityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupReservationAffinityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupReservationAffinityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupReservationAffinityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupReservationAffinityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupReservationAffinityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupReservationAffinityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupReservationAffinityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupReservationAffinityOutputReference) ResetConsumeReservationType() {
	_jsii_.InvokeVoid(
		w,
		"resetConsumeReservationType",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupReservationAffinityOutputReference) ResetKey() {
	_jsii_.InvokeVoid(
		w,
		"resetKey",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupReservationAffinityOutputReference) ResetValues() {
	_jsii_.InvokeVoid(
		w,
		"resetValues",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupReservationAffinityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := w.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkbenchInstanceGceSetupReservationAffinityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

