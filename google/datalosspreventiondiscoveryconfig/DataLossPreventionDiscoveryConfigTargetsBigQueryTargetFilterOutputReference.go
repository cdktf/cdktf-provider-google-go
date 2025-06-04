// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventiondiscoveryconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/datalosspreventiondiscoveryconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference interface {
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
	InternalValue() *DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilter
	SetInternalValue(val *DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilter)
	OtherTables() DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOtherTablesOutputReference
	OtherTablesInput() *DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOtherTables
	TableReference() DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterTableReferenceOutputReference
	TableReferenceInput() *DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterTableReference
	Tables() DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterTablesOutputReference
	TablesInput() *DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterTables
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
	PutOtherTables(value *DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOtherTables)
	PutTableReference(value *DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterTableReference)
	PutTables(value *DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterTables)
	ResetOtherTables()
	ResetTableReference()
	ResetTables()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference
type jsiiProxy_DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference) InternalValue() *DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilter {
	var returns *DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilter
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference) OtherTables() DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOtherTablesOutputReference {
	var returns DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOtherTablesOutputReference
	_jsii_.Get(
		j,
		"otherTables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference) OtherTablesInput() *DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOtherTables {
	var returns *DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOtherTables
	_jsii_.Get(
		j,
		"otherTablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference) TableReference() DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterTableReferenceOutputReference {
	var returns DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterTableReferenceOutputReference
	_jsii_.Get(
		j,
		"tableReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference) TableReferenceInput() *DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterTableReference {
	var returns *DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterTableReference
	_jsii_.Get(
		j,
		"tableReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference) Tables() DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterTablesOutputReference {
	var returns DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterTablesOutputReference
	_jsii_.Get(
		j,
		"tables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference) TablesInput() *DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterTables {
	var returns *DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterTables
	_jsii_.Get(
		j,
		"tablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference {
	_init_.Initialize()

	if err := validateNewDataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataLossPreventionDiscoveryConfig.DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference_Override(d DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataLossPreventionDiscoveryConfig.DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference)SetInternalValue(val *DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilter) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference) PutOtherTables(value *DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOtherTables) {
	if err := d.validatePutOtherTablesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOtherTables",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference) PutTableReference(value *DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterTableReference) {
	if err := d.validatePutTableReferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTableReference",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference) PutTables(value *DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterTables) {
	if err := d.validatePutTablesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTables",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference) ResetOtherTables() {
	_jsii_.InvokeVoid(
		d,
		"resetOtherTables",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference) ResetTableReference() {
	_jsii_.InvokeVoid(
		d,
		"resetTableReference",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference) ResetTables() {
	_jsii_.InvokeVoid(
		d,
		"resetTables",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

