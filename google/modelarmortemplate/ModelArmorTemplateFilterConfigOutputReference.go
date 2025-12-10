// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelarmortemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/modelarmortemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ModelArmorTemplateFilterConfigOutputReference interface {
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
	InternalValue() *ModelArmorTemplateFilterConfig
	SetInternalValue(val *ModelArmorTemplateFilterConfig)
	MaliciousUriFilterSettings() ModelArmorTemplateFilterConfigMaliciousUriFilterSettingsOutputReference
	MaliciousUriFilterSettingsInput() *ModelArmorTemplateFilterConfigMaliciousUriFilterSettings
	PiAndJailbreakFilterSettings() ModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference
	PiAndJailbreakFilterSettingsInput() *ModelArmorTemplateFilterConfigPiAndJailbreakFilterSettings
	RaiSettings() ModelArmorTemplateFilterConfigRaiSettingsOutputReference
	RaiSettingsInput() *ModelArmorTemplateFilterConfigRaiSettings
	SdpSettings() ModelArmorTemplateFilterConfigSdpSettingsOutputReference
	SdpSettingsInput() *ModelArmorTemplateFilterConfigSdpSettings
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
	PutMaliciousUriFilterSettings(value *ModelArmorTemplateFilterConfigMaliciousUriFilterSettings)
	PutPiAndJailbreakFilterSettings(value *ModelArmorTemplateFilterConfigPiAndJailbreakFilterSettings)
	PutRaiSettings(value *ModelArmorTemplateFilterConfigRaiSettings)
	PutSdpSettings(value *ModelArmorTemplateFilterConfigSdpSettings)
	ResetMaliciousUriFilterSettings()
	ResetPiAndJailbreakFilterSettings()
	ResetRaiSettings()
	ResetSdpSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ModelArmorTemplateFilterConfigOutputReference
type jsiiProxy_ModelArmorTemplateFilterConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ModelArmorTemplateFilterConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateFilterConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateFilterConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateFilterConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateFilterConfigOutputReference) InternalValue() *ModelArmorTemplateFilterConfig {
	var returns *ModelArmorTemplateFilterConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateFilterConfigOutputReference) MaliciousUriFilterSettings() ModelArmorTemplateFilterConfigMaliciousUriFilterSettingsOutputReference {
	var returns ModelArmorTemplateFilterConfigMaliciousUriFilterSettingsOutputReference
	_jsii_.Get(
		j,
		"maliciousUriFilterSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateFilterConfigOutputReference) MaliciousUriFilterSettingsInput() *ModelArmorTemplateFilterConfigMaliciousUriFilterSettings {
	var returns *ModelArmorTemplateFilterConfigMaliciousUriFilterSettings
	_jsii_.Get(
		j,
		"maliciousUriFilterSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateFilterConfigOutputReference) PiAndJailbreakFilterSettings() ModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference {
	var returns ModelArmorTemplateFilterConfigPiAndJailbreakFilterSettingsOutputReference
	_jsii_.Get(
		j,
		"piAndJailbreakFilterSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateFilterConfigOutputReference) PiAndJailbreakFilterSettingsInput() *ModelArmorTemplateFilterConfigPiAndJailbreakFilterSettings {
	var returns *ModelArmorTemplateFilterConfigPiAndJailbreakFilterSettings
	_jsii_.Get(
		j,
		"piAndJailbreakFilterSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateFilterConfigOutputReference) RaiSettings() ModelArmorTemplateFilterConfigRaiSettingsOutputReference {
	var returns ModelArmorTemplateFilterConfigRaiSettingsOutputReference
	_jsii_.Get(
		j,
		"raiSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateFilterConfigOutputReference) RaiSettingsInput() *ModelArmorTemplateFilterConfigRaiSettings {
	var returns *ModelArmorTemplateFilterConfigRaiSettings
	_jsii_.Get(
		j,
		"raiSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateFilterConfigOutputReference) SdpSettings() ModelArmorTemplateFilterConfigSdpSettingsOutputReference {
	var returns ModelArmorTemplateFilterConfigSdpSettingsOutputReference
	_jsii_.Get(
		j,
		"sdpSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateFilterConfigOutputReference) SdpSettingsInput() *ModelArmorTemplateFilterConfigSdpSettings {
	var returns *ModelArmorTemplateFilterConfigSdpSettings
	_jsii_.Get(
		j,
		"sdpSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateFilterConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateFilterConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewModelArmorTemplateFilterConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ModelArmorTemplateFilterConfigOutputReference {
	_init_.Initialize()

	if err := validateNewModelArmorTemplateFilterConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ModelArmorTemplateFilterConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.modelArmorTemplate.ModelArmorTemplateFilterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewModelArmorTemplateFilterConfigOutputReference_Override(m ModelArmorTemplateFilterConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.modelArmorTemplate.ModelArmorTemplateFilterConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ModelArmorTemplateFilterConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ModelArmorTemplateFilterConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ModelArmorTemplateFilterConfigOutputReference)SetInternalValue(val *ModelArmorTemplateFilterConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ModelArmorTemplateFilterConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ModelArmorTemplateFilterConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_ModelArmorTemplateFilterConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorTemplateFilterConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_ModelArmorTemplateFilterConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ModelArmorTemplateFilterConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_ModelArmorTemplateFilterConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_ModelArmorTemplateFilterConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_ModelArmorTemplateFilterConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_ModelArmorTemplateFilterConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_ModelArmorTemplateFilterConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_ModelArmorTemplateFilterConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_ModelArmorTemplateFilterConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorTemplateFilterConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ModelArmorTemplateFilterConfigOutputReference) PutMaliciousUriFilterSettings(value *ModelArmorTemplateFilterConfigMaliciousUriFilterSettings) {
	if err := m.validatePutMaliciousUriFilterSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putMaliciousUriFilterSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ModelArmorTemplateFilterConfigOutputReference) PutPiAndJailbreakFilterSettings(value *ModelArmorTemplateFilterConfigPiAndJailbreakFilterSettings) {
	if err := m.validatePutPiAndJailbreakFilterSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putPiAndJailbreakFilterSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ModelArmorTemplateFilterConfigOutputReference) PutRaiSettings(value *ModelArmorTemplateFilterConfigRaiSettings) {
	if err := m.validatePutRaiSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putRaiSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ModelArmorTemplateFilterConfigOutputReference) PutSdpSettings(value *ModelArmorTemplateFilterConfigSdpSettings) {
	if err := m.validatePutSdpSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putSdpSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ModelArmorTemplateFilterConfigOutputReference) ResetMaliciousUriFilterSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetMaliciousUriFilterSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelArmorTemplateFilterConfigOutputReference) ResetPiAndJailbreakFilterSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetPiAndJailbreakFilterSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelArmorTemplateFilterConfigOutputReference) ResetRaiSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetRaiSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelArmorTemplateFilterConfigOutputReference) ResetSdpSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetSdpSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelArmorTemplateFilterConfigOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_ModelArmorTemplateFilterConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

