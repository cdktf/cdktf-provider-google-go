// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigqueryanalyticshubdataexchange

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/bigqueryanalyticshubdataexchange/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigOutputReference interface {
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
	DcrExchangeConfig() BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigDcrExchangeConfigOutputReference
	DcrExchangeConfigInput() *BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigDcrExchangeConfig
	DefaultExchangeConfig() BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigDefaultExchangeConfigOutputReference
	DefaultExchangeConfigInput() *BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigDefaultExchangeConfig
	// Experimental.
	Fqn() *string
	InternalValue() *BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfig
	SetInternalValue(val *BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfig)
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
	PutDcrExchangeConfig(value *BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigDcrExchangeConfig)
	PutDefaultExchangeConfig(value *BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigDefaultExchangeConfig)
	ResetDcrExchangeConfig()
	ResetDefaultExchangeConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigOutputReference
type jsiiProxy_BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigOutputReference) DcrExchangeConfig() BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigDcrExchangeConfigOutputReference {
	var returns BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigDcrExchangeConfigOutputReference
	_jsii_.Get(
		j,
		"dcrExchangeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigOutputReference) DcrExchangeConfigInput() *BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigDcrExchangeConfig {
	var returns *BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigDcrExchangeConfig
	_jsii_.Get(
		j,
		"dcrExchangeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigOutputReference) DefaultExchangeConfig() BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigDefaultExchangeConfigOutputReference {
	var returns BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigDefaultExchangeConfigOutputReference
	_jsii_.Get(
		j,
		"defaultExchangeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigOutputReference) DefaultExchangeConfigInput() *BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigDefaultExchangeConfig {
	var returns *BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigDefaultExchangeConfig
	_jsii_.Get(
		j,
		"defaultExchangeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigOutputReference) InternalValue() *BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfig {
	var returns *BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigOutputReference {
	_init_.Initialize()

	if err := validateNewBigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.bigqueryAnalyticsHubDataExchange.BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigOutputReference_Override(b BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.bigqueryAnalyticsHubDataExchange.BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigOutputReference)SetInternalValue(val *BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigOutputReference) PutDcrExchangeConfig(value *BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigDcrExchangeConfig) {
	if err := b.validatePutDcrExchangeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putDcrExchangeConfig",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigOutputReference) PutDefaultExchangeConfig(value *BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigDefaultExchangeConfig) {
	if err := b.validatePutDefaultExchangeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putDefaultExchangeConfig",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigOutputReference) ResetDcrExchangeConfig() {
	_jsii_.InvokeVoid(
		b,
		"resetDcrExchangeConfig",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigOutputReference) ResetDefaultExchangeConfig() {
	_jsii_.InvokeVoid(
		b,
		"resetDefaultExchangeConfig",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (b *jsiiProxy_BigqueryAnalyticsHubDataExchangeSharingEnvironmentConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

