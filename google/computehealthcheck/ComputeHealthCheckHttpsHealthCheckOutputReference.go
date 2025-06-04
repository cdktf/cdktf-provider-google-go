// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computehealthcheck

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/computehealthcheck/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeHealthCheckHttpsHealthCheckOutputReference interface {
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
	Host() *string
	SetHost(val *string)
	HostInput() *string
	InternalValue() *ComputeHealthCheckHttpsHealthCheck
	SetInternalValue(val *ComputeHealthCheckHttpsHealthCheck)
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
	RequestPath() *string
	SetRequestPath(val *string)
	RequestPathInput() *string
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
	ResetHost()
	ResetPort()
	ResetPortName()
	ResetPortSpecification()
	ResetProxyHeader()
	ResetRequestPath()
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

// The jsii proxy struct for ComputeHealthCheckHttpsHealthCheckOutputReference
type jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference) Host() *string {
	var returns *string
	_jsii_.Get(
		j,
		"host",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference) HostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference) InternalValue() *ComputeHealthCheckHttpsHealthCheck {
	var returns *ComputeHealthCheckHttpsHealthCheck
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference) PortName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference) PortNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference) PortSpecification() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference) PortSpecificationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference) ProxyHeader() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference) ProxyHeaderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference) RequestPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference) RequestPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference) Response() *string {
	var returns *string
	_jsii_.Get(
		j,
		"response",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference) ResponseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeHealthCheckHttpsHealthCheckOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeHealthCheckHttpsHealthCheckOutputReference {
	_init_.Initialize()

	if err := validateNewComputeHealthCheckHttpsHealthCheckOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeHealthCheck.ComputeHealthCheckHttpsHealthCheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeHealthCheckHttpsHealthCheckOutputReference_Override(c ComputeHealthCheckHttpsHealthCheckOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeHealthCheck.ComputeHealthCheckHttpsHealthCheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference)SetHost(val *string) {
	if err := j.validateSetHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"host",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference)SetInternalValue(val *ComputeHealthCheckHttpsHealthCheck) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference)SetPortName(val *string) {
	if err := j.validateSetPortNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portName",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference)SetPortSpecification(val *string) {
	if err := j.validateSetPortSpecificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portSpecification",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference)SetProxyHeader(val *string) {
	if err := j.validateSetProxyHeaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proxyHeader",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference)SetRequestPath(val *string) {
	if err := j.validateSetRequestPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestPath",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference)SetResponse(val *string) {
	if err := j.validateSetResponseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"response",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference) ResetHost() {
	_jsii_.InvokeVoid(
		c,
		"resetHost",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		c,
		"resetPort",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference) ResetPortName() {
	_jsii_.InvokeVoid(
		c,
		"resetPortName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference) ResetPortSpecification() {
	_jsii_.InvokeVoid(
		c,
		"resetPortSpecification",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference) ResetProxyHeader() {
	_jsii_.InvokeVoid(
		c,
		"resetProxyHeader",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference) ResetRequestPath() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestPath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference) ResetResponse() {
	_jsii_.InvokeVoid(
		c,
		"resetResponse",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeHealthCheckHttpsHealthCheckOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

