// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package discoveryenginechatengine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/discoveryenginechatengine/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DiscoveryEngineChatEngineChatEngineConfigOutputReference interface {
	cdktf.ComplexObject
	AgentCreationConfig() DiscoveryEngineChatEngineChatEngineConfigAgentCreationConfigOutputReference
	AgentCreationConfigInput() *DiscoveryEngineChatEngineChatEngineConfigAgentCreationConfig
	AllowCrossRegion() interface{}
	SetAllowCrossRegion(val interface{})
	AllowCrossRegionInput() interface{}
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
	DialogflowAgentToLink() *string
	SetDialogflowAgentToLink(val *string)
	DialogflowAgentToLinkInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DiscoveryEngineChatEngineChatEngineConfig
	SetInternalValue(val *DiscoveryEngineChatEngineChatEngineConfig)
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
	PutAgentCreationConfig(value *DiscoveryEngineChatEngineChatEngineConfigAgentCreationConfig)
	ResetAgentCreationConfig()
	ResetAllowCrossRegion()
	ResetDialogflowAgentToLink()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DiscoveryEngineChatEngineChatEngineConfigOutputReference
type jsiiProxy_DiscoveryEngineChatEngineChatEngineConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DiscoveryEngineChatEngineChatEngineConfigOutputReference) AgentCreationConfig() DiscoveryEngineChatEngineChatEngineConfigAgentCreationConfigOutputReference {
	var returns DiscoveryEngineChatEngineChatEngineConfigAgentCreationConfigOutputReference
	_jsii_.Get(
		j,
		"agentCreationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineChatEngineChatEngineConfigOutputReference) AgentCreationConfigInput() *DiscoveryEngineChatEngineChatEngineConfigAgentCreationConfig {
	var returns *DiscoveryEngineChatEngineChatEngineConfigAgentCreationConfig
	_jsii_.Get(
		j,
		"agentCreationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineChatEngineChatEngineConfigOutputReference) AllowCrossRegion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowCrossRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineChatEngineChatEngineConfigOutputReference) AllowCrossRegionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowCrossRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineChatEngineChatEngineConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineChatEngineChatEngineConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineChatEngineChatEngineConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineChatEngineChatEngineConfigOutputReference) DialogflowAgentToLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dialogflowAgentToLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineChatEngineChatEngineConfigOutputReference) DialogflowAgentToLinkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dialogflowAgentToLinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineChatEngineChatEngineConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineChatEngineChatEngineConfigOutputReference) InternalValue() *DiscoveryEngineChatEngineChatEngineConfig {
	var returns *DiscoveryEngineChatEngineChatEngineConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineChatEngineChatEngineConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineChatEngineChatEngineConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDiscoveryEngineChatEngineChatEngineConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DiscoveryEngineChatEngineChatEngineConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDiscoveryEngineChatEngineChatEngineConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DiscoveryEngineChatEngineChatEngineConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.discoveryEngineChatEngine.DiscoveryEngineChatEngineChatEngineConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDiscoveryEngineChatEngineChatEngineConfigOutputReference_Override(d DiscoveryEngineChatEngineChatEngineConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.discoveryEngineChatEngine.DiscoveryEngineChatEngineChatEngineConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DiscoveryEngineChatEngineChatEngineConfigOutputReference)SetAllowCrossRegion(val interface{}) {
	if err := j.validateSetAllowCrossRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowCrossRegion",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineChatEngineChatEngineConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineChatEngineChatEngineConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineChatEngineChatEngineConfigOutputReference)SetDialogflowAgentToLink(val *string) {
	if err := j.validateSetDialogflowAgentToLinkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dialogflowAgentToLink",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineChatEngineChatEngineConfigOutputReference)SetInternalValue(val *DiscoveryEngineChatEngineChatEngineConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineChatEngineChatEngineConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineChatEngineChatEngineConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DiscoveryEngineChatEngineChatEngineConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineChatEngineChatEngineConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DiscoveryEngineChatEngineChatEngineConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DiscoveryEngineChatEngineChatEngineConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DiscoveryEngineChatEngineChatEngineConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DiscoveryEngineChatEngineChatEngineConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DiscoveryEngineChatEngineChatEngineConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DiscoveryEngineChatEngineChatEngineConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DiscoveryEngineChatEngineChatEngineConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DiscoveryEngineChatEngineChatEngineConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DiscoveryEngineChatEngineChatEngineConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineChatEngineChatEngineConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineChatEngineChatEngineConfigOutputReference) PutAgentCreationConfig(value *DiscoveryEngineChatEngineChatEngineConfigAgentCreationConfig) {
	if err := d.validatePutAgentCreationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAgentCreationConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DiscoveryEngineChatEngineChatEngineConfigOutputReference) ResetAgentCreationConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetAgentCreationConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineChatEngineChatEngineConfigOutputReference) ResetAllowCrossRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowCrossRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineChatEngineChatEngineConfigOutputReference) ResetDialogflowAgentToLink() {
	_jsii_.InvokeVoid(
		d,
		"resetDialogflowAgentToLink",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineChatEngineChatEngineConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineChatEngineChatEngineConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

