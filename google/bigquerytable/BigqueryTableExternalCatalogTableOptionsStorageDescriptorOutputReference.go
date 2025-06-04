// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigquerytable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/bigquerytable/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference interface {
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
	InputFormat() *string
	SetInputFormat(val *string)
	InputFormatInput() *string
	InternalValue() *BigqueryTableExternalCatalogTableOptionsStorageDescriptor
	SetInternalValue(val *BigqueryTableExternalCatalogTableOptionsStorageDescriptor)
	LocationUri() *string
	SetLocationUri(val *string)
	LocationUriInput() *string
	OutputFormat() *string
	SetOutputFormat(val *string)
	OutputFormatInput() *string
	SerdeInfo() BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference
	SerdeInfoInput() *BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfo
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
	PutSerdeInfo(value *BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfo)
	ResetInputFormat()
	ResetLocationUri()
	ResetOutputFormat()
	ResetSerdeInfo()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference
type jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) InputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) InputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) InternalValue() *BigqueryTableExternalCatalogTableOptionsStorageDescriptor {
	var returns *BigqueryTableExternalCatalogTableOptionsStorageDescriptor
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) LocationUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) LocationUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) OutputFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) OutputFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) SerdeInfo() BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference {
	var returns BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfoOutputReference
	_jsii_.Get(
		j,
		"serdeInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) SerdeInfoInput() *BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfo {
	var returns *BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfo
	_jsii_.Get(
		j,
		"serdeInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference {
	_init_.Initialize()

	if err := validateNewBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.bigqueryTable.BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference_Override(b BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.bigqueryTable.BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference)SetInputFormat(val *string) {
	if err := j.validateSetInputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputFormat",
		val,
	)
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference)SetInternalValue(val *BigqueryTableExternalCatalogTableOptionsStorageDescriptor) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference)SetLocationUri(val *string) {
	if err := j.validateSetLocationUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locationUri",
		val,
	)
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference)SetOutputFormat(val *string) {
	if err := j.validateSetOutputFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputFormat",
		val,
	)
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) PutSerdeInfo(value *BigqueryTableExternalCatalogTableOptionsStorageDescriptorSerdeInfo) {
	if err := b.validatePutSerdeInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putSerdeInfo",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) ResetInputFormat() {
	_jsii_.InvokeVoid(
		b,
		"resetInputFormat",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) ResetLocationUri() {
	_jsii_.InvokeVoid(
		b,
		"resetLocationUri",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) ResetOutputFormat() {
	_jsii_.InvokeVoid(
		b,
		"resetOutputFormat",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) ResetSerdeInfo() {
	_jsii_.InvokeVoid(
		b,
		"resetSerdeInfo",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (b *jsiiProxy_BigqueryTableExternalCatalogTableOptionsStorageDescriptorOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

