// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package modelarmortemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/modelarmortemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ModelArmorTemplateFilterConfigSdpSettingsOutputReference interface {
	cdktf.ComplexObject
	AdvancedConfig() ModelArmorTemplateFilterConfigSdpSettingsAdvancedConfigOutputReference
	AdvancedConfigInput() *ModelArmorTemplateFilterConfigSdpSettingsAdvancedConfig
	BasicConfig() ModelArmorTemplateFilterConfigSdpSettingsBasicConfigOutputReference
	BasicConfigInput() *ModelArmorTemplateFilterConfigSdpSettingsBasicConfig
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
	InternalValue() *ModelArmorTemplateFilterConfigSdpSettings
	SetInternalValue(val *ModelArmorTemplateFilterConfigSdpSettings)
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
	PutAdvancedConfig(value *ModelArmorTemplateFilterConfigSdpSettingsAdvancedConfig)
	PutBasicConfig(value *ModelArmorTemplateFilterConfigSdpSettingsBasicConfig)
	ResetAdvancedConfig()
	ResetBasicConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ModelArmorTemplateFilterConfigSdpSettingsOutputReference
type jsiiProxy_ModelArmorTemplateFilterConfigSdpSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ModelArmorTemplateFilterConfigSdpSettingsOutputReference) AdvancedConfig() ModelArmorTemplateFilterConfigSdpSettingsAdvancedConfigOutputReference {
	var returns ModelArmorTemplateFilterConfigSdpSettingsAdvancedConfigOutputReference
	_jsii_.Get(
		j,
		"advancedConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateFilterConfigSdpSettingsOutputReference) AdvancedConfigInput() *ModelArmorTemplateFilterConfigSdpSettingsAdvancedConfig {
	var returns *ModelArmorTemplateFilterConfigSdpSettingsAdvancedConfig
	_jsii_.Get(
		j,
		"advancedConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateFilterConfigSdpSettingsOutputReference) BasicConfig() ModelArmorTemplateFilterConfigSdpSettingsBasicConfigOutputReference {
	var returns ModelArmorTemplateFilterConfigSdpSettingsBasicConfigOutputReference
	_jsii_.Get(
		j,
		"basicConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateFilterConfigSdpSettingsOutputReference) BasicConfigInput() *ModelArmorTemplateFilterConfigSdpSettingsBasicConfig {
	var returns *ModelArmorTemplateFilterConfigSdpSettingsBasicConfig
	_jsii_.Get(
		j,
		"basicConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateFilterConfigSdpSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateFilterConfigSdpSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateFilterConfigSdpSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateFilterConfigSdpSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateFilterConfigSdpSettingsOutputReference) InternalValue() *ModelArmorTemplateFilterConfigSdpSettings {
	var returns *ModelArmorTemplateFilterConfigSdpSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateFilterConfigSdpSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ModelArmorTemplateFilterConfigSdpSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewModelArmorTemplateFilterConfigSdpSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ModelArmorTemplateFilterConfigSdpSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewModelArmorTemplateFilterConfigSdpSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ModelArmorTemplateFilterConfigSdpSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.modelArmorTemplate.ModelArmorTemplateFilterConfigSdpSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewModelArmorTemplateFilterConfigSdpSettingsOutputReference_Override(m ModelArmorTemplateFilterConfigSdpSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.modelArmorTemplate.ModelArmorTemplateFilterConfigSdpSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ModelArmorTemplateFilterConfigSdpSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ModelArmorTemplateFilterConfigSdpSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ModelArmorTemplateFilterConfigSdpSettingsOutputReference)SetInternalValue(val *ModelArmorTemplateFilterConfigSdpSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ModelArmorTemplateFilterConfigSdpSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ModelArmorTemplateFilterConfigSdpSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_ModelArmorTemplateFilterConfigSdpSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorTemplateFilterConfigSdpSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_ModelArmorTemplateFilterConfigSdpSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ModelArmorTemplateFilterConfigSdpSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_ModelArmorTemplateFilterConfigSdpSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_ModelArmorTemplateFilterConfigSdpSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_ModelArmorTemplateFilterConfigSdpSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_ModelArmorTemplateFilterConfigSdpSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_ModelArmorTemplateFilterConfigSdpSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_ModelArmorTemplateFilterConfigSdpSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_ModelArmorTemplateFilterConfigSdpSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ModelArmorTemplateFilterConfigSdpSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_ModelArmorTemplateFilterConfigSdpSettingsOutputReference) PutAdvancedConfig(value *ModelArmorTemplateFilterConfigSdpSettingsAdvancedConfig) {
	if err := m.validatePutAdvancedConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAdvancedConfig",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ModelArmorTemplateFilterConfigSdpSettingsOutputReference) PutBasicConfig(value *ModelArmorTemplateFilterConfigSdpSettingsBasicConfig) {
	if err := m.validatePutBasicConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putBasicConfig",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_ModelArmorTemplateFilterConfigSdpSettingsOutputReference) ResetAdvancedConfig() {
	_jsii_.InvokeVoid(
		m,
		"resetAdvancedConfig",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelArmorTemplateFilterConfigSdpSettingsOutputReference) ResetBasicConfig() {
	_jsii_.InvokeVoid(
		m,
		"resetBasicConfig",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ModelArmorTemplateFilterConfigSdpSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_ModelArmorTemplateFilterConfigSdpSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

