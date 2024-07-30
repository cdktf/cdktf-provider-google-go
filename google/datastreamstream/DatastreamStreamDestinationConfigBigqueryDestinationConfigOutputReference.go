// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datastreamstream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v13/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v13/datastreamstream/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference interface {
	cdktf.ComplexObject
	AppendOnly() DatastreamStreamDestinationConfigBigqueryDestinationConfigAppendOnlyOutputReference
	AppendOnlyInput() *DatastreamStreamDestinationConfigBigqueryDestinationConfigAppendOnly
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
	DataFreshness() *string
	SetDataFreshness(val *string)
	DataFreshnessInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DatastreamStreamDestinationConfigBigqueryDestinationConfig
	SetInternalValue(val *DatastreamStreamDestinationConfigBigqueryDestinationConfig)
	Merge() DatastreamStreamDestinationConfigBigqueryDestinationConfigMergeOutputReference
	MergeInput() *DatastreamStreamDestinationConfigBigqueryDestinationConfigMerge
	SingleTargetDataset() DatastreamStreamDestinationConfigBigqueryDestinationConfigSingleTargetDatasetOutputReference
	SingleTargetDatasetInput() *DatastreamStreamDestinationConfigBigqueryDestinationConfigSingleTargetDataset
	SourceHierarchyDatasets() DatastreamStreamDestinationConfigBigqueryDestinationConfigSourceHierarchyDatasetsOutputReference
	SourceHierarchyDatasetsInput() *DatastreamStreamDestinationConfigBigqueryDestinationConfigSourceHierarchyDatasets
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
	PutAppendOnly(value *DatastreamStreamDestinationConfigBigqueryDestinationConfigAppendOnly)
	PutMerge(value *DatastreamStreamDestinationConfigBigqueryDestinationConfigMerge)
	PutSingleTargetDataset(value *DatastreamStreamDestinationConfigBigqueryDestinationConfigSingleTargetDataset)
	PutSourceHierarchyDatasets(value *DatastreamStreamDestinationConfigBigqueryDestinationConfigSourceHierarchyDatasets)
	ResetAppendOnly()
	ResetDataFreshness()
	ResetMerge()
	ResetSingleTargetDataset()
	ResetSourceHierarchyDatasets()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference
type jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) AppendOnly() DatastreamStreamDestinationConfigBigqueryDestinationConfigAppendOnlyOutputReference {
	var returns DatastreamStreamDestinationConfigBigqueryDestinationConfigAppendOnlyOutputReference
	_jsii_.Get(
		j,
		"appendOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) AppendOnlyInput() *DatastreamStreamDestinationConfigBigqueryDestinationConfigAppendOnly {
	var returns *DatastreamStreamDestinationConfigBigqueryDestinationConfigAppendOnly
	_jsii_.Get(
		j,
		"appendOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) DataFreshness() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFreshness",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) DataFreshnessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFreshnessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) InternalValue() *DatastreamStreamDestinationConfigBigqueryDestinationConfig {
	var returns *DatastreamStreamDestinationConfigBigqueryDestinationConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) Merge() DatastreamStreamDestinationConfigBigqueryDestinationConfigMergeOutputReference {
	var returns DatastreamStreamDestinationConfigBigqueryDestinationConfigMergeOutputReference
	_jsii_.Get(
		j,
		"merge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) MergeInput() *DatastreamStreamDestinationConfigBigqueryDestinationConfigMerge {
	var returns *DatastreamStreamDestinationConfigBigqueryDestinationConfigMerge
	_jsii_.Get(
		j,
		"mergeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) SingleTargetDataset() DatastreamStreamDestinationConfigBigqueryDestinationConfigSingleTargetDatasetOutputReference {
	var returns DatastreamStreamDestinationConfigBigqueryDestinationConfigSingleTargetDatasetOutputReference
	_jsii_.Get(
		j,
		"singleTargetDataset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) SingleTargetDatasetInput() *DatastreamStreamDestinationConfigBigqueryDestinationConfigSingleTargetDataset {
	var returns *DatastreamStreamDestinationConfigBigqueryDestinationConfigSingleTargetDataset
	_jsii_.Get(
		j,
		"singleTargetDatasetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) SourceHierarchyDatasets() DatastreamStreamDestinationConfigBigqueryDestinationConfigSourceHierarchyDatasetsOutputReference {
	var returns DatastreamStreamDestinationConfigBigqueryDestinationConfigSourceHierarchyDatasetsOutputReference
	_jsii_.Get(
		j,
		"sourceHierarchyDatasets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) SourceHierarchyDatasetsInput() *DatastreamStreamDestinationConfigBigqueryDestinationConfigSourceHierarchyDatasets {
	var returns *DatastreamStreamDestinationConfigBigqueryDestinationConfigSourceHierarchyDatasets
	_jsii_.Get(
		j,
		"sourceHierarchyDatasetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.datastreamStream.DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference_Override(d DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.datastreamStream.DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference)SetDataFreshness(val *string) {
	if err := j.validateSetDataFreshnessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataFreshness",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference)SetInternalValue(val *DatastreamStreamDestinationConfigBigqueryDestinationConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) PutAppendOnly(value *DatastreamStreamDestinationConfigBigqueryDestinationConfigAppendOnly) {
	if err := d.validatePutAppendOnlyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAppendOnly",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) PutMerge(value *DatastreamStreamDestinationConfigBigqueryDestinationConfigMerge) {
	if err := d.validatePutMergeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMerge",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) PutSingleTargetDataset(value *DatastreamStreamDestinationConfigBigqueryDestinationConfigSingleTargetDataset) {
	if err := d.validatePutSingleTargetDatasetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSingleTargetDataset",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) PutSourceHierarchyDatasets(value *DatastreamStreamDestinationConfigBigqueryDestinationConfigSourceHierarchyDatasets) {
	if err := d.validatePutSourceHierarchyDatasetsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSourceHierarchyDatasets",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) ResetAppendOnly() {
	_jsii_.InvokeVoid(
		d,
		"resetAppendOnly",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) ResetDataFreshness() {
	_jsii_.InvokeVoid(
		d,
		"resetDataFreshness",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) ResetMerge() {
	_jsii_.InvokeVoid(
		d,
		"resetMerge",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) ResetSingleTargetDataset() {
	_jsii_.InvokeVoid(
		d,
		"resetSingleTargetDataset",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) ResetSourceHierarchyDatasets() {
	_jsii_.InvokeVoid(
		d,
		"resetSourceHierarchyDatasets",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

