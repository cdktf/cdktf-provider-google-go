// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeperinstanceconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v12/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v12/computeperinstanceconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputePerInstanceConfigPreservedStateInternalIpOutputReference interface {
	cdktf.ComplexObject
	AutoDelete() *string
	SetAutoDelete(val *string)
	AutoDeleteInput() *string
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
	InterfaceName() *string
	SetInterfaceName(val *string)
	InterfaceNameInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IpAddress() ComputePerInstanceConfigPreservedStateInternalIpIpAddressOutputReference
	IpAddressInput() *ComputePerInstanceConfigPreservedStateInternalIpIpAddress
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
	PutIpAddress(value *ComputePerInstanceConfigPreservedStateInternalIpIpAddress)
	ResetAutoDelete()
	ResetIpAddress()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputePerInstanceConfigPreservedStateInternalIpOutputReference
type jsiiProxy_ComputePerInstanceConfigPreservedStateInternalIpOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputePerInstanceConfigPreservedStateInternalIpOutputReference) AutoDelete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputePerInstanceConfigPreservedStateInternalIpOutputReference) AutoDeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputePerInstanceConfigPreservedStateInternalIpOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputePerInstanceConfigPreservedStateInternalIpOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputePerInstanceConfigPreservedStateInternalIpOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputePerInstanceConfigPreservedStateInternalIpOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputePerInstanceConfigPreservedStateInternalIpOutputReference) InterfaceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interfaceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputePerInstanceConfigPreservedStateInternalIpOutputReference) InterfaceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interfaceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputePerInstanceConfigPreservedStateInternalIpOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputePerInstanceConfigPreservedStateInternalIpOutputReference) IpAddress() ComputePerInstanceConfigPreservedStateInternalIpIpAddressOutputReference {
	var returns ComputePerInstanceConfigPreservedStateInternalIpIpAddressOutputReference
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputePerInstanceConfigPreservedStateInternalIpOutputReference) IpAddressInput() *ComputePerInstanceConfigPreservedStateInternalIpIpAddress {
	var returns *ComputePerInstanceConfigPreservedStateInternalIpIpAddress
	_jsii_.Get(
		j,
		"ipAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputePerInstanceConfigPreservedStateInternalIpOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputePerInstanceConfigPreservedStateInternalIpOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputePerInstanceConfigPreservedStateInternalIpOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ComputePerInstanceConfigPreservedStateInternalIpOutputReference {
	_init_.Initialize()

	if err := validateNewComputePerInstanceConfigPreservedStateInternalIpOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputePerInstanceConfigPreservedStateInternalIpOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computePerInstanceConfig.ComputePerInstanceConfigPreservedStateInternalIpOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewComputePerInstanceConfigPreservedStateInternalIpOutputReference_Override(c ComputePerInstanceConfigPreservedStateInternalIpOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computePerInstanceConfig.ComputePerInstanceConfigPreservedStateInternalIpOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ComputePerInstanceConfigPreservedStateInternalIpOutputReference)SetAutoDelete(val *string) {
	if err := j.validateSetAutoDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoDelete",
		val,
	)
}

func (j *jsiiProxy_ComputePerInstanceConfigPreservedStateInternalIpOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputePerInstanceConfigPreservedStateInternalIpOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputePerInstanceConfigPreservedStateInternalIpOutputReference)SetInterfaceName(val *string) {
	if err := j.validateSetInterfaceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interfaceName",
		val,
	)
}

func (j *jsiiProxy_ComputePerInstanceConfigPreservedStateInternalIpOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputePerInstanceConfigPreservedStateInternalIpOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputePerInstanceConfigPreservedStateInternalIpOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputePerInstanceConfigPreservedStateInternalIpOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputePerInstanceConfigPreservedStateInternalIpOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputePerInstanceConfigPreservedStateInternalIpOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputePerInstanceConfigPreservedStateInternalIpOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputePerInstanceConfigPreservedStateInternalIpOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputePerInstanceConfigPreservedStateInternalIpOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputePerInstanceConfigPreservedStateInternalIpOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputePerInstanceConfigPreservedStateInternalIpOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputePerInstanceConfigPreservedStateInternalIpOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputePerInstanceConfigPreservedStateInternalIpOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputePerInstanceConfigPreservedStateInternalIpOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputePerInstanceConfigPreservedStateInternalIpOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputePerInstanceConfigPreservedStateInternalIpOutputReference) PutIpAddress(value *ComputePerInstanceConfigPreservedStateInternalIpIpAddress) {
	if err := c.validatePutIpAddressParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putIpAddress",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputePerInstanceConfigPreservedStateInternalIpOutputReference) ResetAutoDelete() {
	_jsii_.InvokeVoid(
		c,
		"resetAutoDelete",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputePerInstanceConfigPreservedStateInternalIpOutputReference) ResetIpAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetIpAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputePerInstanceConfigPreservedStateInternalIpOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputePerInstanceConfigPreservedStateInternalIpOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

