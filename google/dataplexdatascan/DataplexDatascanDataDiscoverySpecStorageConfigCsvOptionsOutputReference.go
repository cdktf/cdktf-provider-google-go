// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataplexdatascan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/dataplexdatascan/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference interface {
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
	Delimiter() *string
	SetDelimiter(val *string)
	DelimiterInput() *string
	Encoding() *string
	SetEncoding(val *string)
	EncodingInput() *string
	// Experimental.
	Fqn() *string
	HeaderRows() *float64
	SetHeaderRows(val *float64)
	HeaderRowsInput() *float64
	InternalValue() *DataplexDatascanDataDiscoverySpecStorageConfigCsvOptions
	SetInternalValue(val *DataplexDatascanDataDiscoverySpecStorageConfigCsvOptions)
	Quote() *string
	SetQuote(val *string)
	QuoteInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TypeInferenceDisabled() interface{}
	SetTypeInferenceDisabled(val interface{})
	TypeInferenceDisabledInput() interface{}
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	ResetDelimiter()
	ResetEncoding()
	ResetHeaderRows()
	ResetQuote()
	ResetTypeInferenceDisabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference
type jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference) Delimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference) DelimiterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delimiterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference) Encoding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encoding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference) EncodingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encodingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference) HeaderRows() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"headerRows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference) HeaderRowsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"headerRowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference) InternalValue() *DataplexDatascanDataDiscoverySpecStorageConfigCsvOptions {
	var returns *DataplexDatascanDataDiscoverySpecStorageConfigCsvOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference) Quote() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quote",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference) QuoteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quoteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference) TypeInferenceDisabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"typeInferenceDisabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference) TypeInferenceDisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"typeInferenceDisabledInput",
		&returns,
	)
	return returns
}


func NewDataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewDataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataplexDatascan.DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference_Override(d DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataplexDatascan.DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference)SetDelimiter(val *string) {
	if err := j.validateSetDelimiterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"delimiter",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference)SetEncoding(val *string) {
	if err := j.validateSetEncodingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encoding",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference)SetHeaderRows(val *float64) {
	if err := j.validateSetHeaderRowsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"headerRows",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference)SetInternalValue(val *DataplexDatascanDataDiscoverySpecStorageConfigCsvOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference)SetQuote(val *string) {
	if err := j.validateSetQuoteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"quote",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference)SetTypeInferenceDisabled(val interface{}) {
	if err := j.validateSetTypeInferenceDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"typeInferenceDisabled",
		val,
	)
}

func (d *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference) ResetDelimiter() {
	_jsii_.InvokeVoid(
		d,
		"resetDelimiter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference) ResetEncoding() {
	_jsii_.InvokeVoid(
		d,
		"resetEncoding",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference) ResetHeaderRows() {
	_jsii_.InvokeVoid(
		d,
		"resetHeaderRows",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference) ResetQuote() {
	_jsii_.InvokeVoid(
		d,
		"resetQuote",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference) ResetTypeInferenceDisabled() {
	_jsii_.InvokeVoid(
		d,
		"resetTypeInferenceDisabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

