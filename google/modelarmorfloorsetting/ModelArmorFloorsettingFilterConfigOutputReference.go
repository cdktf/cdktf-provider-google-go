// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelarmorfloorsetting

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/modelarmorfloorsetting/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ModelArmorFloorsettingFilterConfigOutputReference interface {
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
	InternalValue() *ModelArmorFloorsettingFilterConfig
	SetInternalValue(val *ModelArmorFloorsettingFilterConfig)
	MaliciousUriFilterSettings() ModelArmorFloorsettingFilterConfigMaliciousUriFilterSettingsOutputReference
	MaliciousUriFilterSettingsInput() *ModelArmorFloorsettingFilterConfigMaliciousUriFilterSettings
	PiAndJailbreakFilterSettings() ModelArmorFloorsettingFilterConfigPiAndJailbreakFilterSettingsOutputReference
	PiAndJailbreakFilterSettingsInput() *ModelArmorFloorsettingFilterConfigPiAndJailbreakFilterSettings
	RaiSettings() ModelArmorFloorsettingFilterConfigRaiSettingsOutputReference
	RaiSettingsInput() *ModelArmorFloorsettingFilterConfigRaiSettings
	SdpSettings() ModelArmorFloorsettingFilterConfigSdpSettingsOutputReference
	SdpSettingsInput() *ModelArmorFloorsettingFilterConfigSdpSettings
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
	PutMaliciousUriFilterSettings(value *ModelArmorFloorsettingFilterConfigMaliciousUriFilterSettings)
	PutPiAndJailbreakFilterSettings(value *ModelArmorFloorsettingFilterConfigPiAndJailbreakFilterSettings)
	PutRaiSettings(value *ModelArmorFloorsettingFilterConfigRaiSettings)
	PutSdpSettings(value *ModelArmorFloorsettingFilterConfigSdpSettings)
	ResetMaliciousUriFilterSettings()
	ResetPiAndJailbreakFilterSettings()
	ResetRaiSettings()
	ResetSdpSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ModelArmorFloorsettingFilterConfigOutputReference
type jsiiProxy_ModelArmorFloorsettingFilterConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ModelArmorFloorsettingFilterConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingFilterConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingFilterConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingFilterConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingFilterConfigOutputReference) InternalValue() *ModelArmorFloorsettingFilterConfig {
	var returns *ModelArmorFloorsettingFilterConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingFilterConfigOutputReference) MaliciousUriFilterSettings() ModelArmorFloorsettingFilterConfigMaliciousUriFilterSettingsOutputReference {
	var returns ModelArmorFloorsettingFilterConfigMaliciousUriFilterSettingsOutputReference
	_jsii_.Get(
		j,
		"maliciousUriFilterSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingFilterConfigOutputReference) MaliciousUriFilterSettingsInput() *ModelArmorFloorsettingFilterConfigMaliciousUriFilterSettings {
	var returns *ModelArmorFloorsettingFilterConfigMaliciousUriFilterSettings
	_jsii_.Get(
		j,
		"maliciousUriFilterSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingFilterConfigOutputReference) PiAndJailbreakFilterSettings() ModelArmorFloorsettingFilterConfigPiAndJailbreakFilterSettingsOutputReference {
	var returns ModelArmorFloorsettingFilterConfigPiAndJailbreakFilterSettingsOutputReference
	_jsii_.Get(
		j,
		"piAndJailbreakFilterSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingFilterConfigOutputReference) PiAndJailbreakFilterSettingsInput() *ModelArmorFloorsettingFilterConfigPiAndJailbreakFilterSettings {
	var returns *ModelArmorFloorsettingFilterConfigPiAndJailbreakFilterSettings
	_jsii_.Get(
		j,
		"piAndJailbreakFilterSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingFilterConfigOutputReference) RaiSettings() ModelArmorFloorsettingFilterConfigRaiSettingsOutputReference {
	var returns ModelArmorFloorsettingFilterConfigRaiSettingsOutputReference
	_jsii_.Get(
		j,
		"raiSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingFilterConfigOutputReference) RaiSettingsInput() *ModelArmorFloorsettingFilterConfigRaiSettings {
	var returns *ModelArmorFloorsettingFilterConfigRaiSettings
	_jsii_.Get(
		j,
		"raiSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingFilterConfigOutputReference) SdpSettings() ModelArmorFloorsettingFilterConfigSdpSettingsOutputReference {
	var returns ModelArmorFloorsettingFilterConfigSdpSettingsOutputReference
	_jsii_.Get(
		j,
		"sdpSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingFilterConfigOutputReference) SdpSettingsInput() *ModelArmorFloorsettingFilterConfigSdpSettings {
	var returns *ModelArmorFloorsettingFilterConfigSdpSettings
	_jsii_.Get(
		j,
		"sdpSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingFilterConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorFloorsettingFilterConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewModelArmorFloorsettingFilterConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ModelArmorFloorsettingFilterConfigOutputReference {
	_init_.Initialize()

	if err := validateNewModelArmorFloorsettingFilterConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ModelArmorFloorsettingFilterConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.modelArmorFloorsetting.ModelArmorFloorsettingFilterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewModelArmorFloorsettingFilterConfigOutputReference_Override(m ModelArmorFloorsettingFilterConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.modelArmorFloorsetting.ModelArmorFloorsettingFilterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ModelArmorFloorsettingFilterConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ModelArmorFloorsettingFilterConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ModelArmorFloorsettingFilterConfigOutputReference)SetInternalValue(val *ModelArmorFloorsettingFilterConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ModelArmorFloorsettingFilterConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ModelArmorFloorsettingFilterConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_ModelArmorFloorsettingFilterConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsettingFilterConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_ModelArmorFloorsettingFilterConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ModelArmorFloorsettingFilterConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_ModelArmorFloorsettingFilterConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_ModelArmorFloorsettingFilterConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_ModelArmorFloorsettingFilterConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_ModelArmorFloorsettingFilterConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_ModelArmorFloorsettingFilterConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_ModelArmorFloorsettingFilterConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_ModelArmorFloorsettingFilterConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsettingFilterConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsettingFilterConfigOutputReference) PutMaliciousUriFilterSettings(value *ModelArmorFloorsettingFilterConfigMaliciousUriFilterSettings) {
	if err := m.validatePutMaliciousUriFilterSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putMaliciousUriFilterSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ModelArmorFloorsettingFilterConfigOutputReference) PutPiAndJailbreakFilterSettings(value *ModelArmorFloorsettingFilterConfigPiAndJailbreakFilterSettings) {
	if err := m.validatePutPiAndJailbreakFilterSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putPiAndJailbreakFilterSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ModelArmorFloorsettingFilterConfigOutputReference) PutRaiSettings(value *ModelArmorFloorsettingFilterConfigRaiSettings) {
	if err := m.validatePutRaiSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putRaiSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ModelArmorFloorsettingFilterConfigOutputReference) PutSdpSettings(value *ModelArmorFloorsettingFilterConfigSdpSettings) {
	if err := m.validatePutSdpSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putSdpSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ModelArmorFloorsettingFilterConfigOutputReference) ResetMaliciousUriFilterSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetMaliciousUriFilterSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelArmorFloorsettingFilterConfigOutputReference) ResetPiAndJailbreakFilterSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetPiAndJailbreakFilterSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelArmorFloorsettingFilterConfigOutputReference) ResetRaiSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetRaiSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelArmorFloorsettingFilterConfigOutputReference) ResetSdpSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetSdpSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelArmorFloorsettingFilterConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorFloorsettingFilterConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

