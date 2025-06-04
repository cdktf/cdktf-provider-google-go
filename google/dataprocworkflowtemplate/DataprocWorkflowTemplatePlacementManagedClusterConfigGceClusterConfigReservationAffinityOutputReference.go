// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocworkflowtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/dataprocworkflowtemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference interface {
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
	InternalValue() *DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity
	SetInternalValue(val *DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity)
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

// The jsii proxy struct for DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference
type jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference) ConsumeReservationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumeReservationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference) ConsumeReservationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumeReservationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference) InternalValue() *DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity {
	var returns *DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference) ValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewDataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference {
	_init_.Initialize()

	if err := validateNewDataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataprocWorkflowTemplate.DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference_Override(d DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataprocWorkflowTemplate.DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference)SetConsumeReservationType(val *string) {
	if err := j.validateSetConsumeReservationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consumeReservationType",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference)SetInternalValue(val *DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference)SetKey(val *string) {
	if err := j.validateSetKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference)SetValues(val *[]*string) {
	if err := j.validateSetValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference) ResetConsumeReservationType() {
	_jsii_.InvokeVoid(
		d,
		"resetConsumeReservationType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference) ResetKey() {
	_jsii_.InvokeVoid(
		d,
		"resetKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference) ResetValues() {
	_jsii_.InvokeVoid(
		d,
		"resetValues",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

