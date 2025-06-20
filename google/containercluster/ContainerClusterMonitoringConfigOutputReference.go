// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/containercluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerClusterMonitoringConfigOutputReference interface {
	cdktf.ComplexObject
	AdvancedDatapathObservabilityConfig() ContainerClusterMonitoringConfigAdvancedDatapathObservabilityConfigOutputReference
	AdvancedDatapathObservabilityConfigInput() *ContainerClusterMonitoringConfigAdvancedDatapathObservabilityConfig
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
	EnableComponents() *[]*string
	SetEnableComponents(val *[]*string)
	EnableComponentsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *ContainerClusterMonitoringConfig
	SetInternalValue(val *ContainerClusterMonitoringConfig)
	ManagedPrometheus() ContainerClusterMonitoringConfigManagedPrometheusOutputReference
	ManagedPrometheusInput() *ContainerClusterMonitoringConfigManagedPrometheus
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
	PutAdvancedDatapathObservabilityConfig(value *ContainerClusterMonitoringConfigAdvancedDatapathObservabilityConfig)
	PutManagedPrometheus(value *ContainerClusterMonitoringConfigManagedPrometheus)
	ResetAdvancedDatapathObservabilityConfig()
	ResetEnableComponents()
	ResetManagedPrometheus()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerClusterMonitoringConfigOutputReference
type jsiiProxy_ContainerClusterMonitoringConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerClusterMonitoringConfigOutputReference) AdvancedDatapathObservabilityConfig() ContainerClusterMonitoringConfigAdvancedDatapathObservabilityConfigOutputReference {
	var returns ContainerClusterMonitoringConfigAdvancedDatapathObservabilityConfigOutputReference
	_jsii_.Get(
		j,
		"advancedDatapathObservabilityConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterMonitoringConfigOutputReference) AdvancedDatapathObservabilityConfigInput() *ContainerClusterMonitoringConfigAdvancedDatapathObservabilityConfig {
	var returns *ContainerClusterMonitoringConfigAdvancedDatapathObservabilityConfig
	_jsii_.Get(
		j,
		"advancedDatapathObservabilityConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterMonitoringConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterMonitoringConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterMonitoringConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterMonitoringConfigOutputReference) EnableComponents() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enableComponents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterMonitoringConfigOutputReference) EnableComponentsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enableComponentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterMonitoringConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterMonitoringConfigOutputReference) InternalValue() *ContainerClusterMonitoringConfig {
	var returns *ContainerClusterMonitoringConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterMonitoringConfigOutputReference) ManagedPrometheus() ContainerClusterMonitoringConfigManagedPrometheusOutputReference {
	var returns ContainerClusterMonitoringConfigManagedPrometheusOutputReference
	_jsii_.Get(
		j,
		"managedPrometheus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterMonitoringConfigOutputReference) ManagedPrometheusInput() *ContainerClusterMonitoringConfigManagedPrometheus {
	var returns *ContainerClusterMonitoringConfigManagedPrometheus
	_jsii_.Get(
		j,
		"managedPrometheusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterMonitoringConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterMonitoringConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewContainerClusterMonitoringConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContainerClusterMonitoringConfigOutputReference {
	_init_.Initialize()

	if err := validateNewContainerClusterMonitoringConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerClusterMonitoringConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.containerCluster.ContainerClusterMonitoringConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerClusterMonitoringConfigOutputReference_Override(c ContainerClusterMonitoringConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.containerCluster.ContainerClusterMonitoringConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerClusterMonitoringConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterMonitoringConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterMonitoringConfigOutputReference)SetEnableComponents(val *[]*string) {
	if err := j.validateSetEnableComponentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableComponents",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterMonitoringConfigOutputReference)SetInternalValue(val *ContainerClusterMonitoringConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterMonitoringConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterMonitoringConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ContainerClusterMonitoringConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterMonitoringConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ContainerClusterMonitoringConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerClusterMonitoringConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ContainerClusterMonitoringConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ContainerClusterMonitoringConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ContainerClusterMonitoringConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ContainerClusterMonitoringConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ContainerClusterMonitoringConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ContainerClusterMonitoringConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ContainerClusterMonitoringConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterMonitoringConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerClusterMonitoringConfigOutputReference) PutAdvancedDatapathObservabilityConfig(value *ContainerClusterMonitoringConfigAdvancedDatapathObservabilityConfig) {
	if err := c.validatePutAdvancedDatapathObservabilityConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAdvancedDatapathObservabilityConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterMonitoringConfigOutputReference) PutManagedPrometheus(value *ContainerClusterMonitoringConfigManagedPrometheus) {
	if err := c.validatePutManagedPrometheusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putManagedPrometheus",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterMonitoringConfigOutputReference) ResetAdvancedDatapathObservabilityConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetAdvancedDatapathObservabilityConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterMonitoringConfigOutputReference) ResetEnableComponents() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableComponents",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterMonitoringConfigOutputReference) ResetManagedPrometheus() {
	_jsii_.InvokeVoid(
		c,
		"resetManagedPrometheus",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterMonitoringConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ContainerClusterMonitoringConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

