// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregionhealthcheck

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/computeregionhealthcheck/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeRegionHealthCheckTcpHealthCheckOutputReference interface {
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
	InternalValue() *ComputeRegionHealthCheckTcpHealthCheck
	SetInternalValue(val *ComputeRegionHealthCheckTcpHealthCheck)
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	PortName() *string
	SetPortName(val *string)
	PortNameInput() *string
	PortSpecification() *string
	SetPortSpecification(val *string)
	PortSpecificationInput() *string
	ProxyHeader() *string
	SetProxyHeader(val *string)
	ProxyHeaderInput() *string
	Request() *string
	SetRequest(val *string)
	RequestInput() *string
	Response() *string
	SetResponse(val *string)
	ResponseInput() *string
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
	ResetPort()
	ResetPortName()
	ResetPortSpecification()
	ResetProxyHeader()
	ResetRequest()
	ResetResponse()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeRegionHealthCheckTcpHealthCheckOutputReference
type jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference) InternalValue() *ComputeRegionHealthCheckTcpHealthCheck {
	var returns *ComputeRegionHealthCheckTcpHealthCheck
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference) PortName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference) PortNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference) PortSpecification() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference) PortSpecificationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference) ProxyHeader() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference) ProxyHeaderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference) Request() *string {
	var returns *string
	_jsii_.Get(
		j,
		"request",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference) RequestInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference) Response() *string {
	var returns *string
	_jsii_.Get(
		j,
		"response",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference) ResponseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeRegionHealthCheckTcpHealthCheckOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeRegionHealthCheckTcpHealthCheckOutputReference {
	_init_.Initialize()

	if err := validateNewComputeRegionHealthCheckTcpHealthCheckOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionHealthCheck.ComputeRegionHealthCheckTcpHealthCheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeRegionHealthCheckTcpHealthCheckOutputReference_Override(c ComputeRegionHealthCheckTcpHealthCheckOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionHealthCheck.ComputeRegionHealthCheckTcpHealthCheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference)SetInternalValue(val *ComputeRegionHealthCheckTcpHealthCheck) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference)SetPortName(val *string) {
	if err := j.validateSetPortNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portName",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference)SetPortSpecification(val *string) {
	if err := j.validateSetPortSpecificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portSpecification",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference)SetProxyHeader(val *string) {
	if err := j.validateSetProxyHeaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proxyHeader",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference)SetRequest(val *string) {
	if err := j.validateSetRequestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"request",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference)SetResponse(val *string) {
	if err := j.validateSetResponseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"response",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		c,
		"resetPort",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference) ResetPortName() {
	_jsii_.InvokeVoid(
		c,
		"resetPortName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference) ResetPortSpecification() {
	_jsii_.InvokeVoid(
		c,
		"resetPortSpecification",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference) ResetProxyHeader() {
	_jsii_.InvokeVoid(
		c,
		"resetProxyHeader",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference) ResetRequest() {
	_jsii_.InvokeVoid(
		c,
		"resetRequest",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference) ResetResponse() {
	_jsii_.InvokeVoid(
		c,
		"resetResponse",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeRegionHealthCheckTcpHealthCheckOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

