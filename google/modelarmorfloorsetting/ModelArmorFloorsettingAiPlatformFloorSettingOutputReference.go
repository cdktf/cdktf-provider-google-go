// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelarmorfloorsetting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/modelarmorfloorsetting/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ModelArmorFloorsettingAiPlatformFloorSettingOutputReference interface {
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
	EnableCloudLogging() interface{}
	SetEnableCloudLogging(val interface{})
	EnableCloudLoggingInput() interface{}
	// Experimental.
	Fqn() *string
	InspectAndBlock() interface{}
	SetInspectAndBlock(val interface{})
	InspectAndBlockInput() interface{}
	InspectOnly() interface{}
	SetInspectOnly(val interface{})
	InspectOnlyInput() interface{}
	InternalValue() *ModelArmorFloorsettingAiPlatformFloorSetting
	SetInternalValue(val *ModelArmorFloorsettingAiPlatformFloorSetting)
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	ResetEnableCloudLogging()
	ResetInspectAndBlock()
	ResetInspectOnly()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ModelArmorFloorsettingAiPlatformFloorSettingOutputReference
type jsiiProxy_ModelArmorFloorsettingAiPlatformFloorSettingOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ModelArmorFloorsettingAiPlatformFloorSettingOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingAiPlatformFloorSettingOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingAiPlatformFloorSettingOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingAiPlatformFloorSettingOutputReference) EnableCloudLogging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableCloudLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingAiPlatformFloorSettingOutputReference) EnableCloudLoggingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableCloudLoggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingAiPlatformFloorSettingOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingAiPlatformFloorSettingOutputReference) InspectAndBlock() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inspectAndBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingAiPlatformFloorSettingOutputReference) InspectAndBlockInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inspectAndBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingAiPlatformFloorSettingOutputReference) InspectOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inspectOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingAiPlatformFloorSettingOutputReference) InspectOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inspectOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingAiPlatformFloorSettingOutputReference) InternalValue() *ModelArmorFloorsettingAiPlatformFloorSetting {
	var returns *ModelArmorFloorsettingAiPlatformFloorSetting
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingAiPlatformFloorSettingOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingAiPlatformFloorSettingOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewModelArmorFloorsettingAiPlatformFloorSettingOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ModelArmorFloorsettingAiPlatformFloorSettingOutputReference {
	_init_.Initialize()

	if err := validateNewModelArmorFloorsettingAiPlatformFloorSettingOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ModelArmorFloorsettingAiPlatformFloorSettingOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.modelArmorFloorsetting.ModelArmorFloorsettingAiPlatformFloorSettingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewModelArmorFloorsettingAiPlatformFloorSettingOutputReference_Override(m ModelArmorFloorsettingAiPlatformFloorSettingOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.modelArmorFloorsetting.ModelArmorFloorsettingAiPlatformFloorSettingOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ModelArmorFloorsettingAiPlatformFloorSettingOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ModelArmorFloorsettingAiPlatformFloorSettingOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ModelArmorFloorsettingAiPlatformFloorSettingOutputReference)SetEnableCloudLogging(val interface{}) {
	if err := j.validateSetEnableCloudLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableCloudLogging",
		val,
	)
}

func (j *jsiiProxy_ModelArmorFloorsettingAiPlatformFloorSettingOutputReference)SetInspectAndBlock(val interface{}) {
	if err := j.validateSetInspectAndBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inspectAndBlock",
		val,
	)
}

func (j *jsiiProxy_ModelArmorFloorsettingAiPlatformFloorSettingOutputReference)SetInspectOnly(val interface{}) {
	if err := j.validateSetInspectOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inspectOnly",
		val,
	)
}

func (j *jsiiProxy_ModelArmorFloorsettingAiPlatformFloorSettingOutputReference)SetInternalValue(val *ModelArmorFloorsettingAiPlatformFloorSetting) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ModelArmorFloorsettingAiPlatformFloorSettingOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ModelArmorFloorsettingAiPlatformFloorSettingOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_ModelArmorFloorsettingAiPlatformFloorSettingOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsettingAiPlatformFloorSettingOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsettingAiPlatformFloorSettingOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsettingAiPlatformFloorSettingOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsettingAiPlatformFloorSettingOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsettingAiPlatformFloorSettingOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsettingAiPlatformFloorSettingOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsettingAiPlatformFloorSettingOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsettingAiPlatformFloorSettingOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsettingAiPlatformFloorSettingOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsettingAiPlatformFloorSettingOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsettingAiPlatformFloorSettingOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsettingAiPlatformFloorSettingOutputReference) ResetEnableCloudLogging() {
	_jsii_.InvokeVoid(
		m,
		"resetEnableCloudLogging",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelArmorFloorsettingAiPlatformFloorSettingOutputReference) ResetInspectAndBlock() {
	_jsii_.InvokeVoid(
		m,
		"resetInspectAndBlock",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelArmorFloorsettingAiPlatformFloorSettingOutputReference) ResetInspectOnly() {
	_jsii_.InvokeVoid(
		m,
		"resetInspectOnly",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelArmorFloorsettingAiPlatformFloorSettingOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsettingAiPlatformFloorSettingOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

