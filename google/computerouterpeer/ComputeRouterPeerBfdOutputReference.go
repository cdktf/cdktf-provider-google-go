// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computerouterpeer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/computerouterpeer/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeRouterPeerBfdOutputReference interface {
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
	InternalValue() *ComputeRouterPeerBfd
	SetInternalValue(val *ComputeRouterPeerBfd)
	MinReceiveInterval() *float64
	SetMinReceiveInterval(val *float64)
	MinReceiveIntervalInput() *float64
	MinTransmitInterval() *float64
	SetMinTransmitInterval(val *float64)
	MinTransmitIntervalInput() *float64
	Multiplier() *float64
	SetMultiplier(val *float64)
	MultiplierInput() *float64
	SessionInitializationMode() *string
	SetSessionInitializationMode(val *string)
	SessionInitializationModeInput() *string
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
	ResetMinReceiveInterval()
	ResetMinTransmitInterval()
	ResetMultiplier()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeRouterPeerBfdOutputReference
type jsiiProxy_ComputeRouterPeerBfdOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeRouterPeerBfdOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeerBfdOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeerBfdOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeerBfdOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeerBfdOutputReference) InternalValue() *ComputeRouterPeerBfd {
	var returns *ComputeRouterPeerBfd
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeerBfdOutputReference) MinReceiveInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minReceiveInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeerBfdOutputReference) MinReceiveIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minReceiveIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeerBfdOutputReference) MinTransmitInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minTransmitInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeerBfdOutputReference) MinTransmitIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minTransmitIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeerBfdOutputReference) Multiplier() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"multiplier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeerBfdOutputReference) MultiplierInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"multiplierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeerBfdOutputReference) SessionInitializationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionInitializationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeerBfdOutputReference) SessionInitializationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionInitializationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeerBfdOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRouterPeerBfdOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeRouterPeerBfdOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeRouterPeerBfdOutputReference {
	_init_.Initialize()

	if err := validateNewComputeRouterPeerBfdOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeRouterPeerBfdOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeRouterPeer.ComputeRouterPeerBfdOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeRouterPeerBfdOutputReference_Override(c ComputeRouterPeerBfdOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeRouterPeer.ComputeRouterPeerBfdOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeRouterPeerBfdOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterPeerBfdOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterPeerBfdOutputReference)SetInternalValue(val *ComputeRouterPeerBfd) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterPeerBfdOutputReference)SetMinReceiveInterval(val *float64) {
	if err := j.validateSetMinReceiveIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minReceiveInterval",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterPeerBfdOutputReference)SetMinTransmitInterval(val *float64) {
	if err := j.validateSetMinTransmitIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minTransmitInterval",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterPeerBfdOutputReference)SetMultiplier(val *float64) {
	if err := j.validateSetMultiplierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiplier",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterPeerBfdOutputReference)SetSessionInitializationMode(val *string) {
	if err := j.validateSetSessionInitializationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionInitializationMode",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterPeerBfdOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeRouterPeerBfdOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeRouterPeerBfdOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRouterPeerBfdOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeRouterPeerBfdOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeRouterPeerBfdOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeRouterPeerBfdOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeRouterPeerBfdOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeRouterPeerBfdOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeRouterPeerBfdOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeRouterPeerBfdOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeRouterPeerBfdOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeRouterPeerBfdOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRouterPeerBfdOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeRouterPeerBfdOutputReference) ResetMinReceiveInterval() {
	_jsii_.InvokeVoid(
		c,
		"resetMinReceiveInterval",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterPeerBfdOutputReference) ResetMinTransmitInterval() {
	_jsii_.InvokeVoid(
		c,
		"resetMinTransmitInterval",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterPeerBfdOutputReference) ResetMultiplier() {
	_jsii_.InvokeVoid(
		c,
		"resetMultiplier",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRouterPeerBfdOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeRouterPeerBfdOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

