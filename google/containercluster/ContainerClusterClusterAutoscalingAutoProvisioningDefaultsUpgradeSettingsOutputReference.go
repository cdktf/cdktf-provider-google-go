// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/containercluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference interface {
	cdktf.ComplexObject
	BlueGreenSettings() ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsBlueGreenSettingsOutputReference
	BlueGreenSettingsInput() *ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsBlueGreenSettings
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
	InternalValue() *ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettings
	SetInternalValue(val *ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettings)
	MaxSurge() *float64
	SetMaxSurge(val *float64)
	MaxSurgeInput() *float64
	MaxUnavailable() *float64
	SetMaxUnavailable(val *float64)
	MaxUnavailableInput() *float64
	Strategy() *string
	SetStrategy(val *string)
	StrategyInput() *string
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
	PutBlueGreenSettings(value *ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsBlueGreenSettings)
	ResetBlueGreenSettings()
	ResetMaxSurge()
	ResetMaxUnavailable()
	ResetStrategy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference
type jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) BlueGreenSettings() ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsBlueGreenSettingsOutputReference {
	var returns ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsBlueGreenSettingsOutputReference
	_jsii_.Get(
		j,
		"blueGreenSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) BlueGreenSettingsInput() *ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsBlueGreenSettings {
	var returns *ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsBlueGreenSettings
	_jsii_.Get(
		j,
		"blueGreenSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) InternalValue() *ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettings {
	var returns *ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) MaxSurge() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSurge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) MaxSurgeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSurgeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) MaxUnavailable() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnavailable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) MaxUnavailableInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnavailableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) Strategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) StrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.containerCluster.ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference_Override(c ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.containerCluster.ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference)SetInternalValue(val *ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference)SetMaxSurge(val *float64) {
	if err := j.validateSetMaxSurgeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSurge",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference)SetMaxUnavailable(val *float64) {
	if err := j.validateSetMaxUnavailableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxUnavailable",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference)SetStrategy(val *string) {
	if err := j.validateSetStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"strategy",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) PutBlueGreenSettings(value *ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsBlueGreenSettings) {
	if err := c.validatePutBlueGreenSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBlueGreenSettings",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) ResetBlueGreenSettings() {
	_jsii_.InvokeVoid(
		c,
		"resetBlueGreenSettings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) ResetMaxSurge() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxSurge",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) ResetMaxUnavailable() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxUnavailable",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) ResetStrategy() {
	_jsii_.InvokeVoid(
		c,
		"resetStrategy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterClusterAutoscalingAutoProvisioningDefaultsUpgradeSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

