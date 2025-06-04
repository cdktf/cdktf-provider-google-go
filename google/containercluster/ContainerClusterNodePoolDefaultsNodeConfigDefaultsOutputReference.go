// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/containercluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference interface {
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
	ContainerdConfig() ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigOutputReference
	ContainerdConfigInput() *ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfig
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	GcfsConfig() ContainerClusterNodePoolDefaultsNodeConfigDefaultsGcfsConfigOutputReference
	GcfsConfigInput() *ContainerClusterNodePoolDefaultsNodeConfigDefaultsGcfsConfig
	InsecureKubeletReadonlyPortEnabled() *string
	SetInsecureKubeletReadonlyPortEnabled(val *string)
	InsecureKubeletReadonlyPortEnabledInput() *string
	InternalValue() *ContainerClusterNodePoolDefaultsNodeConfigDefaults
	SetInternalValue(val *ContainerClusterNodePoolDefaultsNodeConfigDefaults)
	LoggingVariant() *string
	SetLoggingVariant(val *string)
	LoggingVariantInput() *string
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
	PutContainerdConfig(value *ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfig)
	PutGcfsConfig(value *ContainerClusterNodePoolDefaultsNodeConfigDefaultsGcfsConfig)
	ResetContainerdConfig()
	ResetGcfsConfig()
	ResetInsecureKubeletReadonlyPortEnabled()
	ResetLoggingVariant()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference
type jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) ContainerdConfig() ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigOutputReference {
	var returns ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfigOutputReference
	_jsii_.Get(
		j,
		"containerdConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) ContainerdConfigInput() *ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfig {
	var returns *ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfig
	_jsii_.Get(
		j,
		"containerdConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) GcfsConfig() ContainerClusterNodePoolDefaultsNodeConfigDefaultsGcfsConfigOutputReference {
	var returns ContainerClusterNodePoolDefaultsNodeConfigDefaultsGcfsConfigOutputReference
	_jsii_.Get(
		j,
		"gcfsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) GcfsConfigInput() *ContainerClusterNodePoolDefaultsNodeConfigDefaultsGcfsConfig {
	var returns *ContainerClusterNodePoolDefaultsNodeConfigDefaultsGcfsConfig
	_jsii_.Get(
		j,
		"gcfsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) InsecureKubeletReadonlyPortEnabled() *string {
	var returns *string
	_jsii_.Get(
		j,
		"insecureKubeletReadonlyPortEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) InsecureKubeletReadonlyPortEnabledInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"insecureKubeletReadonlyPortEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) InternalValue() *ContainerClusterNodePoolDefaultsNodeConfigDefaults {
	var returns *ContainerClusterNodePoolDefaultsNodeConfigDefaults
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) LoggingVariant() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingVariant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) LoggingVariantInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingVariantInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference {
	_init_.Initialize()

	if err := validateNewContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.containerCluster.ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference_Override(c ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.containerCluster.ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference)SetInsecureKubeletReadonlyPortEnabled(val *string) {
	if err := j.validateSetInsecureKubeletReadonlyPortEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insecureKubeletReadonlyPortEnabled",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference)SetInternalValue(val *ContainerClusterNodePoolDefaultsNodeConfigDefaults) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference)SetLoggingVariant(val *string) {
	if err := j.validateSetLoggingVariantParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loggingVariant",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) PutContainerdConfig(value *ContainerClusterNodePoolDefaultsNodeConfigDefaultsContainerdConfig) {
	if err := c.validatePutContainerdConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putContainerdConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) PutGcfsConfig(value *ContainerClusterNodePoolDefaultsNodeConfigDefaultsGcfsConfig) {
	if err := c.validatePutGcfsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGcfsConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) ResetContainerdConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetContainerdConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) ResetGcfsConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetGcfsConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) ResetInsecureKubeletReadonlyPortEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetInsecureKubeletReadonlyPortEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) ResetLoggingVariant() {
	_jsii_.InvokeVoid(
		c,
		"resetLoggingVariant",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ContainerClusterNodePoolDefaultsNodeConfigDefaultsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

