// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iapsettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/iapsettings/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference interface {
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
	Enable() interface{}
	SetEnable(val interface{})
	EnableInput() interface{}
	Expression() *string
	SetExpression(val *string)
	ExpressionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *IapSettingsApplicationSettingsAttributePropagationSettings
	SetInternalValue(val *IapSettingsApplicationSettingsAttributePropagationSettings)
	OutputCredentials() *[]*string
	SetOutputCredentials(val *[]*string)
	OutputCredentialsInput() *[]*string
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
	ResetEnable()
	ResetExpression()
	ResetOutputCredentials()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference
type jsiiProxy_IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference) Enable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference) EnableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference) Expression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference) ExpressionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference) InternalValue() *IapSettingsApplicationSettingsAttributePropagationSettings {
	var returns *IapSettingsApplicationSettingsAttributePropagationSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference) OutputCredentials() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outputCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference) OutputCredentialsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outputCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIapSettingsApplicationSettingsAttributePropagationSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewIapSettingsApplicationSettingsAttributePropagationSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.iapSettings.IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIapSettingsApplicationSettingsAttributePropagationSettingsOutputReference_Override(i IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.iapSettings.IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference)SetEnable(val interface{}) {
	if err := j.validateSetEnableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enable",
		val,
	)
}

func (j *jsiiProxy_IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference)SetExpression(val *string) {
	if err := j.validateSetExpressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expression",
		val,
	)
}

func (j *jsiiProxy_IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference)SetInternalValue(val *IapSettingsApplicationSettingsAttributePropagationSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference)SetOutputCredentials(val *[]*string) {
	if err := j.validateSetOutputCredentialsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputCredentials",
		val,
	)
}

func (j *jsiiProxy_IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference) ResetEnable() {
	_jsii_.InvokeVoid(
		i,
		"resetEnable",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference) ResetExpression() {
	_jsii_.InvokeVoid(
		i,
		"resetExpression",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference) ResetOutputCredentials() {
	_jsii_.InvokeVoid(
		i,
		"resetOutputCredentials",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := i.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IapSettingsApplicationSettingsAttributePropagationSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

