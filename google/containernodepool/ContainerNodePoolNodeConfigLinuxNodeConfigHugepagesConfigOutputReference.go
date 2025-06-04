// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containernodepool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/containernodepool/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference interface {
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
	HugepageSize1G() *float64
	SetHugepageSize1G(val *float64)
	HugepageSize1GInput() *float64
	HugepageSize2M() *float64
	SetHugepageSize2M(val *float64)
	HugepageSize2MInput() *float64
	InternalValue() *ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfig
	SetInternalValue(val *ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfig)
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
	ResetHugepageSize1G()
	ResetHugepageSize2M()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference
type jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference) HugepageSize1G() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hugepageSize1G",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference) HugepageSize1GInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hugepageSize1GInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference) HugepageSize2M() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hugepageSize2M",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference) HugepageSize2MInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hugepageSize2MInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference) InternalValue() *ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfig {
	var returns *ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference {
	_init_.Initialize()

	if err := validateNewContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.containerNodePool.ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference_Override(c ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.containerNodePool.ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference)SetHugepageSize1G(val *float64) {
	if err := j.validateSetHugepageSize1GParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hugepageSize1G",
		val,
	)
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference)SetHugepageSize2M(val *float64) {
	if err := j.validateSetHugepageSize2MParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hugepageSize2M",
		val,
	)
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference)SetInternalValue(val *ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference) ResetHugepageSize1G() {
	_jsii_.InvokeVoid(
		c,
		"resetHugepageSize1G",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference) ResetHugepageSize2M() {
	_jsii_.InvokeVoid(
		c,
		"resetHugepageSize2M",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ContainerNodePoolNodeConfigLinuxNodeConfigHugepagesConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

