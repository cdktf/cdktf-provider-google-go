// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alloydbinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v13/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v13/alloydbinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlloydbInstanceQueryInsightsConfigOutputReference interface {
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
	InternalValue() *AlloydbInstanceQueryInsightsConfig
	SetInternalValue(val *AlloydbInstanceQueryInsightsConfig)
	QueryPlansPerMinute() *float64
	SetQueryPlansPerMinute(val *float64)
	QueryPlansPerMinuteInput() *float64
	QueryStringLength() *float64
	SetQueryStringLength(val *float64)
	QueryStringLengthInput() *float64
	RecordApplicationTags() interface{}
	SetRecordApplicationTags(val interface{})
	RecordApplicationTagsInput() interface{}
	RecordClientAddress() interface{}
	SetRecordClientAddress(val interface{})
	RecordClientAddressInput() interface{}
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
	ResetQueryPlansPerMinute()
	ResetQueryStringLength()
	ResetRecordApplicationTags()
	ResetRecordClientAddress()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AlloydbInstanceQueryInsightsConfigOutputReference
type jsiiProxy_AlloydbInstanceQueryInsightsConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlloydbInstanceQueryInsightsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbInstanceQueryInsightsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbInstanceQueryInsightsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbInstanceQueryInsightsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbInstanceQueryInsightsConfigOutputReference) InternalValue() *AlloydbInstanceQueryInsightsConfig {
	var returns *AlloydbInstanceQueryInsightsConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbInstanceQueryInsightsConfigOutputReference) QueryPlansPerMinute() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queryPlansPerMinute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbInstanceQueryInsightsConfigOutputReference) QueryPlansPerMinuteInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queryPlansPerMinuteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbInstanceQueryInsightsConfigOutputReference) QueryStringLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queryStringLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbInstanceQueryInsightsConfigOutputReference) QueryStringLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queryStringLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbInstanceQueryInsightsConfigOutputReference) RecordApplicationTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recordApplicationTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbInstanceQueryInsightsConfigOutputReference) RecordApplicationTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recordApplicationTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbInstanceQueryInsightsConfigOutputReference) RecordClientAddress() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recordClientAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbInstanceQueryInsightsConfigOutputReference) RecordClientAddressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"recordClientAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbInstanceQueryInsightsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbInstanceQueryInsightsConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAlloydbInstanceQueryInsightsConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AlloydbInstanceQueryInsightsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewAlloydbInstanceQueryInsightsConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlloydbInstanceQueryInsightsConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.alloydbInstance.AlloydbInstanceQueryInsightsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAlloydbInstanceQueryInsightsConfigOutputReference_Override(a AlloydbInstanceQueryInsightsConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.alloydbInstance.AlloydbInstanceQueryInsightsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AlloydbInstanceQueryInsightsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlloydbInstanceQueryInsightsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlloydbInstanceQueryInsightsConfigOutputReference)SetInternalValue(val *AlloydbInstanceQueryInsightsConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlloydbInstanceQueryInsightsConfigOutputReference)SetQueryPlansPerMinute(val *float64) {
	if err := j.validateSetQueryPlansPerMinuteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryPlansPerMinute",
		val,
	)
}

func (j *jsiiProxy_AlloydbInstanceQueryInsightsConfigOutputReference)SetQueryStringLength(val *float64) {
	if err := j.validateSetQueryStringLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryStringLength",
		val,
	)
}

func (j *jsiiProxy_AlloydbInstanceQueryInsightsConfigOutputReference)SetRecordApplicationTags(val interface{}) {
	if err := j.validateSetRecordApplicationTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recordApplicationTags",
		val,
	)
}

func (j *jsiiProxy_AlloydbInstanceQueryInsightsConfigOutputReference)SetRecordClientAddress(val interface{}) {
	if err := j.validateSetRecordClientAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recordClientAddress",
		val,
	)
}

func (j *jsiiProxy_AlloydbInstanceQueryInsightsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlloydbInstanceQueryInsightsConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AlloydbInstanceQueryInsightsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlloydbInstanceQueryInsightsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlloydbInstanceQueryInsightsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlloydbInstanceQueryInsightsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlloydbInstanceQueryInsightsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlloydbInstanceQueryInsightsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlloydbInstanceQueryInsightsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlloydbInstanceQueryInsightsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlloydbInstanceQueryInsightsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlloydbInstanceQueryInsightsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlloydbInstanceQueryInsightsConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlloydbInstanceQueryInsightsConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlloydbInstanceQueryInsightsConfigOutputReference) ResetQueryPlansPerMinute() {
	_jsii_.InvokeVoid(
		a,
		"resetQueryPlansPerMinute",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlloydbInstanceQueryInsightsConfigOutputReference) ResetQueryStringLength() {
	_jsii_.InvokeVoid(
		a,
		"resetQueryStringLength",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlloydbInstanceQueryInsightsConfigOutputReference) ResetRecordApplicationTags() {
	_jsii_.InvokeVoid(
		a,
		"resetRecordApplicationTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlloydbInstanceQueryInsightsConfigOutputReference) ResetRecordClientAddress() {
	_jsii_.InvokeVoid(
		a,
		"resetRecordClientAddress",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlloydbInstanceQueryInsightsConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlloydbInstanceQueryInsightsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

