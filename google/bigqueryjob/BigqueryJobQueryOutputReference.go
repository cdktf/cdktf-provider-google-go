// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigqueryjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v13/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v13/bigqueryjob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BigqueryJobQueryOutputReference interface {
	cdktf.ComplexObject
	AllowLargeResults() interface{}
	SetAllowLargeResults(val interface{})
	AllowLargeResultsInput() interface{}
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
	CreateDisposition() *string
	SetCreateDisposition(val *string)
	CreateDispositionInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DefaultDataset() BigqueryJobQueryDefaultDatasetOutputReference
	DefaultDatasetInput() *BigqueryJobQueryDefaultDataset
	DestinationEncryptionConfiguration() BigqueryJobQueryDestinationEncryptionConfigurationOutputReference
	DestinationEncryptionConfigurationInput() *BigqueryJobQueryDestinationEncryptionConfiguration
	DestinationTable() BigqueryJobQueryDestinationTableOutputReference
	DestinationTableInput() *BigqueryJobQueryDestinationTable
	FlattenResults() interface{}
	SetFlattenResults(val interface{})
	FlattenResultsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *BigqueryJobQuery
	SetInternalValue(val *BigqueryJobQuery)
	MaximumBillingTier() *float64
	SetMaximumBillingTier(val *float64)
	MaximumBillingTierInput() *float64
	MaximumBytesBilled() *string
	SetMaximumBytesBilled(val *string)
	MaximumBytesBilledInput() *string
	ParameterMode() *string
	SetParameterMode(val *string)
	ParameterModeInput() *string
	Priority() *string
	SetPriority(val *string)
	PriorityInput() *string
	Query() *string
	SetQuery(val *string)
	QueryInput() *string
	SchemaUpdateOptions() *[]*string
	SetSchemaUpdateOptions(val *[]*string)
	SchemaUpdateOptionsInput() *[]*string
	ScriptOptions() BigqueryJobQueryScriptOptionsOutputReference
	ScriptOptionsInput() *BigqueryJobQueryScriptOptions
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UseLegacySql() interface{}
	SetUseLegacySql(val interface{})
	UseLegacySqlInput() interface{}
	UseQueryCache() interface{}
	SetUseQueryCache(val interface{})
	UseQueryCacheInput() interface{}
	UserDefinedFunctionResources() BigqueryJobQueryUserDefinedFunctionResourcesList
	UserDefinedFunctionResourcesInput() interface{}
	WriteDisposition() *string
	SetWriteDisposition(val *string)
	WriteDispositionInput() *string
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
	PutDefaultDataset(value *BigqueryJobQueryDefaultDataset)
	PutDestinationEncryptionConfiguration(value *BigqueryJobQueryDestinationEncryptionConfiguration)
	PutDestinationTable(value *BigqueryJobQueryDestinationTable)
	PutScriptOptions(value *BigqueryJobQueryScriptOptions)
	PutUserDefinedFunctionResources(value interface{})
	ResetAllowLargeResults()
	ResetCreateDisposition()
	ResetDefaultDataset()
	ResetDestinationEncryptionConfiguration()
	ResetDestinationTable()
	ResetFlattenResults()
	ResetMaximumBillingTier()
	ResetMaximumBytesBilled()
	ResetParameterMode()
	ResetPriority()
	ResetSchemaUpdateOptions()
	ResetScriptOptions()
	ResetUseLegacySql()
	ResetUseQueryCache()
	ResetUserDefinedFunctionResources()
	ResetWriteDisposition()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BigqueryJobQueryOutputReference
type jsiiProxy_BigqueryJobQueryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference) AllowLargeResults() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowLargeResults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference) AllowLargeResultsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowLargeResultsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference) CreateDisposition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createDisposition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference) CreateDispositionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createDispositionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference) DefaultDataset() BigqueryJobQueryDefaultDatasetOutputReference {
	var returns BigqueryJobQueryDefaultDatasetOutputReference
	_jsii_.Get(
		j,
		"defaultDataset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference) DefaultDatasetInput() *BigqueryJobQueryDefaultDataset {
	var returns *BigqueryJobQueryDefaultDataset
	_jsii_.Get(
		j,
		"defaultDatasetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference) DestinationEncryptionConfiguration() BigqueryJobQueryDestinationEncryptionConfigurationOutputReference {
	var returns BigqueryJobQueryDestinationEncryptionConfigurationOutputReference
	_jsii_.Get(
		j,
		"destinationEncryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference) DestinationEncryptionConfigurationInput() *BigqueryJobQueryDestinationEncryptionConfiguration {
	var returns *BigqueryJobQueryDestinationEncryptionConfiguration
	_jsii_.Get(
		j,
		"destinationEncryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference) DestinationTable() BigqueryJobQueryDestinationTableOutputReference {
	var returns BigqueryJobQueryDestinationTableOutputReference
	_jsii_.Get(
		j,
		"destinationTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference) DestinationTableInput() *BigqueryJobQueryDestinationTable {
	var returns *BigqueryJobQueryDestinationTable
	_jsii_.Get(
		j,
		"destinationTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference) FlattenResults() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"flattenResults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference) FlattenResultsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"flattenResultsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference) InternalValue() *BigqueryJobQuery {
	var returns *BigqueryJobQuery
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference) MaximumBillingTier() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBillingTier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference) MaximumBillingTierInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBillingTierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference) MaximumBytesBilled() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maximumBytesBilled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference) MaximumBytesBilledInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maximumBytesBilledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference) ParameterMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference) ParameterModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference) Priority() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference) PriorityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference) Query() *string {
	var returns *string
	_jsii_.Get(
		j,
		"query",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference) QueryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference) SchemaUpdateOptions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"schemaUpdateOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference) SchemaUpdateOptionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"schemaUpdateOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference) ScriptOptions() BigqueryJobQueryScriptOptionsOutputReference {
	var returns BigqueryJobQueryScriptOptionsOutputReference
	_jsii_.Get(
		j,
		"scriptOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference) ScriptOptionsInput() *BigqueryJobQueryScriptOptions {
	var returns *BigqueryJobQueryScriptOptions
	_jsii_.Get(
		j,
		"scriptOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference) UseLegacySql() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useLegacySql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference) UseLegacySqlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useLegacySqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference) UseQueryCache() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useQueryCache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference) UseQueryCacheInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useQueryCacheInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference) UserDefinedFunctionResources() BigqueryJobQueryUserDefinedFunctionResourcesList {
	var returns BigqueryJobQueryUserDefinedFunctionResourcesList
	_jsii_.Get(
		j,
		"userDefinedFunctionResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference) UserDefinedFunctionResourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userDefinedFunctionResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference) WriteDisposition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"writeDisposition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference) WriteDispositionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"writeDispositionInput",
		&returns,
	)
	return returns
}


func NewBigqueryJobQueryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BigqueryJobQueryOutputReference {
	_init_.Initialize()

	if err := validateNewBigqueryJobQueryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BigqueryJobQueryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.bigqueryJob.BigqueryJobQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBigqueryJobQueryOutputReference_Override(b BigqueryJobQueryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.bigqueryJob.BigqueryJobQueryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference)SetAllowLargeResults(val interface{}) {
	if err := j.validateSetAllowLargeResultsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowLargeResults",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference)SetCreateDisposition(val *string) {
	if err := j.validateSetCreateDispositionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createDisposition",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference)SetFlattenResults(val interface{}) {
	if err := j.validateSetFlattenResultsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flattenResults",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference)SetInternalValue(val *BigqueryJobQuery) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference)SetMaximumBillingTier(val *float64) {
	if err := j.validateSetMaximumBillingTierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumBillingTier",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference)SetMaximumBytesBilled(val *string) {
	if err := j.validateSetMaximumBytesBilledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumBytesBilled",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference)SetParameterMode(val *string) {
	if err := j.validateSetParameterModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameterMode",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference)SetPriority(val *string) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference)SetQuery(val *string) {
	if err := j.validateSetQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"query",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference)SetSchemaUpdateOptions(val *[]*string) {
	if err := j.validateSetSchemaUpdateOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schemaUpdateOptions",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference)SetUseLegacySql(val interface{}) {
	if err := j.validateSetUseLegacySqlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useLegacySql",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference)SetUseQueryCache(val interface{}) {
	if err := j.validateSetUseQueryCacheParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useQueryCache",
		val,
	)
}

func (j *jsiiProxy_BigqueryJobQueryOutputReference)SetWriteDisposition(val *string) {
	if err := j.validateSetWriteDispositionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"writeDisposition",
		val,
	)
}

func (b *jsiiProxy_BigqueryJobQueryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryJobQueryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryJobQueryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryJobQueryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryJobQueryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryJobQueryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryJobQueryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryJobQueryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryJobQueryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryJobQueryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryJobQueryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryJobQueryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryJobQueryOutputReference) PutDefaultDataset(value *BigqueryJobQueryDefaultDataset) {
	if err := b.validatePutDefaultDatasetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putDefaultDataset",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BigqueryJobQueryOutputReference) PutDestinationEncryptionConfiguration(value *BigqueryJobQueryDestinationEncryptionConfiguration) {
	if err := b.validatePutDestinationEncryptionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putDestinationEncryptionConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BigqueryJobQueryOutputReference) PutDestinationTable(value *BigqueryJobQueryDestinationTable) {
	if err := b.validatePutDestinationTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putDestinationTable",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BigqueryJobQueryOutputReference) PutScriptOptions(value *BigqueryJobQueryScriptOptions) {
	if err := b.validatePutScriptOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putScriptOptions",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BigqueryJobQueryOutputReference) PutUserDefinedFunctionResources(value interface{}) {
	if err := b.validatePutUserDefinedFunctionResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putUserDefinedFunctionResources",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BigqueryJobQueryOutputReference) ResetAllowLargeResults() {
	_jsii_.InvokeVoid(
		b,
		"resetAllowLargeResults",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryJobQueryOutputReference) ResetCreateDisposition() {
	_jsii_.InvokeVoid(
		b,
		"resetCreateDisposition",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryJobQueryOutputReference) ResetDefaultDataset() {
	_jsii_.InvokeVoid(
		b,
		"resetDefaultDataset",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryJobQueryOutputReference) ResetDestinationEncryptionConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetDestinationEncryptionConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryJobQueryOutputReference) ResetDestinationTable() {
	_jsii_.InvokeVoid(
		b,
		"resetDestinationTable",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryJobQueryOutputReference) ResetFlattenResults() {
	_jsii_.InvokeVoid(
		b,
		"resetFlattenResults",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryJobQueryOutputReference) ResetMaximumBillingTier() {
	_jsii_.InvokeVoid(
		b,
		"resetMaximumBillingTier",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryJobQueryOutputReference) ResetMaximumBytesBilled() {
	_jsii_.InvokeVoid(
		b,
		"resetMaximumBytesBilled",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryJobQueryOutputReference) ResetParameterMode() {
	_jsii_.InvokeVoid(
		b,
		"resetParameterMode",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryJobQueryOutputReference) ResetPriority() {
	_jsii_.InvokeVoid(
		b,
		"resetPriority",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryJobQueryOutputReference) ResetSchemaUpdateOptions() {
	_jsii_.InvokeVoid(
		b,
		"resetSchemaUpdateOptions",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryJobQueryOutputReference) ResetScriptOptions() {
	_jsii_.InvokeVoid(
		b,
		"resetScriptOptions",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryJobQueryOutputReference) ResetUseLegacySql() {
	_jsii_.InvokeVoid(
		b,
		"resetUseLegacySql",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryJobQueryOutputReference) ResetUseQueryCache() {
	_jsii_.InvokeVoid(
		b,
		"resetUseQueryCache",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryJobQueryOutputReference) ResetUserDefinedFunctionResources() {
	_jsii_.InvokeVoid(
		b,
		"resetUserDefinedFunctionResources",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryJobQueryOutputReference) ResetWriteDisposition() {
	_jsii_.InvokeVoid(
		b,
		"resetWriteDisposition",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryJobQueryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := b.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryJobQueryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

