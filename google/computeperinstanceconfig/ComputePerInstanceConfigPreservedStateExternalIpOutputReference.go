// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeperinstanceconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v13/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v13/computeperinstanceconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputePerInstanceConfigPreservedStateExternalIpOutputReference interface {
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
	IpAddress() ComputePerInstanceConfigPreservedStateExternalIpIpAddressOutputReference
	IpAddressInput() *ComputePerInstanceConfigPreservedStateExternalIpIpAddress
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
	PutIpAddress(value *ComputePerInstanceConfigPreservedStateExternalIpIpAddress)
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

// The jsii proxy struct for ComputePerInstanceConfigPreservedStateExternalIpOutputReference
type jsiiProxy_ComputePerInstanceConfigPreservedStateExternalIpOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputePerInstanceConfigPreservedStateExternalIpOutputReference) AutoDelete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputePerInstanceConfigPreservedStateExternalIpOutputReference) AutoDeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputePerInstanceConfigPreservedStateExternalIpOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputePerInstanceConfigPreservedStateExternalIpOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputePerInstanceConfigPreservedStateExternalIpOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputePerInstanceConfigPreservedStateExternalIpOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputePerInstanceConfigPreservedStateExternalIpOutputReference) InterfaceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interfaceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputePerInstanceConfigPreservedStateExternalIpOutputReference) InterfaceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interfaceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputePerInstanceConfigPreservedStateExternalIpOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputePerInstanceConfigPreservedStateExternalIpOutputReference) IpAddress() ComputePerInstanceConfigPreservedStateExternalIpIpAddressOutputReference {
	var returns ComputePerInstanceConfigPreservedStateExternalIpIpAddressOutputReference
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputePerInstanceConfigPreservedStateExternalIpOutputReference) IpAddressInput() *ComputePerInstanceConfigPreservedStateExternalIpIpAddress {
	var returns *ComputePerInstanceConfigPreservedStateExternalIpIpAddress
	_jsii_.Get(
		j,
		"ipAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputePerInstanceConfigPreservedStateExternalIpOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputePerInstanceConfigPreservedStateExternalIpOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputePerInstanceConfigPreservedStateExternalIpOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ComputePerInstanceConfigPreservedStateExternalIpOutputReference {
	_init_.Initialize()

	if err := validateNewComputePerInstanceConfigPreservedStateExternalIpOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputePerInstanceConfigPreservedStateExternalIpOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computePerInstanceConfig.ComputePerInstanceConfigPreservedStateExternalIpOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewComputePerInstanceConfigPreservedStateExternalIpOutputReference_Override(c ComputePerInstanceConfigPreservedStateExternalIpOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computePerInstanceConfig.ComputePerInstanceConfigPreservedStateExternalIpOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ComputePerInstanceConfigPreservedStateExternalIpOutputReference)SetAutoDelete(val *string) {
	if err := j.validateSetAutoDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoDelete",
		val,
	)
}

func (j *jsiiProxy_ComputePerInstanceConfigPreservedStateExternalIpOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputePerInstanceConfigPreservedStateExternalIpOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputePerInstanceConfigPreservedStateExternalIpOutputReference)SetInterfaceName(val *string) {
	if err := j.validateSetInterfaceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interfaceName",
		val,
	)
}

func (j *jsiiProxy_ComputePerInstanceConfigPreservedStateExternalIpOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputePerInstanceConfigPreservedStateExternalIpOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputePerInstanceConfigPreservedStateExternalIpOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputePerInstanceConfigPreservedStateExternalIpOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputePerInstanceConfigPreservedStateExternalIpOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputePerInstanceConfigPreservedStateExternalIpOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputePerInstanceConfigPreservedStateExternalIpOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputePerInstanceConfigPreservedStateExternalIpOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputePerInstanceConfigPreservedStateExternalIpOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputePerInstanceConfigPreservedStateExternalIpOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputePerInstanceConfigPreservedStateExternalIpOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputePerInstanceConfigPreservedStateExternalIpOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputePerInstanceConfigPreservedStateExternalIpOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputePerInstanceConfigPreservedStateExternalIpOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputePerInstanceConfigPreservedStateExternalIpOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputePerInstanceConfigPreservedStateExternalIpOutputReference) PutIpAddress(value *ComputePerInstanceConfigPreservedStateExternalIpIpAddress) {
	if err := c.validatePutIpAddressParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putIpAddress",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputePerInstanceConfigPreservedStateExternalIpOutputReference) ResetAutoDelete() {
	_jsii_.InvokeVoid(
		c,
		"resetAutoDelete",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputePerInstanceConfigPreservedStateExternalIpOutputReference) ResetIpAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetIpAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputePerInstanceConfigPreservedStateExternalIpOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputePerInstanceConfigPreservedStateExternalIpOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

