// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregionhealthcheck

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v9/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v9/computeregionhealthcheck/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeRegionHealthCheckSslHealthCheckOutputReference interface {
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
	InternalValue() *ComputeRegionHealthCheckSslHealthCheck
	SetInternalValue(val *ComputeRegionHealthCheckSslHealthCheck)
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

// The jsii proxy struct for ComputeRegionHealthCheckSslHealthCheckOutputReference
type jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference) InternalValue() *ComputeRegionHealthCheckSslHealthCheck {
	var returns *ComputeRegionHealthCheckSslHealthCheck
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference) PortName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference) PortNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference) PortSpecification() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference) PortSpecificationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference) ProxyHeader() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference) ProxyHeaderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference) Request() *string {
	var returns *string
	_jsii_.Get(
		j,
		"request",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference) RequestInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference) Response() *string {
	var returns *string
	_jsii_.Get(
		j,
		"response",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference) ResponseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeRegionHealthCheckSslHealthCheckOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeRegionHealthCheckSslHealthCheckOutputReference {
	_init_.Initialize()

	if err := validateNewComputeRegionHealthCheckSslHealthCheckOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionHealthCheck.ComputeRegionHealthCheckSslHealthCheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeRegionHealthCheckSslHealthCheckOutputReference_Override(c ComputeRegionHealthCheckSslHealthCheckOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionHealthCheck.ComputeRegionHealthCheckSslHealthCheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference)SetInternalValue(val *ComputeRegionHealthCheckSslHealthCheck) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference)SetPortName(val *string) {
	if err := j.validateSetPortNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portName",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference)SetPortSpecification(val *string) {
	if err := j.validateSetPortSpecificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portSpecification",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference)SetProxyHeader(val *string) {
	if err := j.validateSetProxyHeaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proxyHeader",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference)SetRequest(val *string) {
	if err := j.validateSetRequestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"request",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference)SetResponse(val *string) {
	if err := j.validateSetResponseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"response",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		c,
		"resetPort",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference) ResetPortName() {
	_jsii_.InvokeVoid(
		c,
		"resetPortName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference) ResetPortSpecification() {
	_jsii_.InvokeVoid(
		c,
		"resetPortSpecification",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference) ResetProxyHeader() {
	_jsii_.InvokeVoid(
		c,
		"resetProxyHeader",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference) ResetRequest() {
	_jsii_.InvokeVoid(
		c,
		"resetRequest",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference) ResetResponse() {
	_jsii_.InvokeVoid(
		c,
		"resetResponse",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeRegionHealthCheckSslHealthCheckOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

