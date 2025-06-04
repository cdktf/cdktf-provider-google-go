// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregionbackendservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/computeregionbackendservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeRegionBackendServiceCircuitBreakersOutputReference interface {
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
	InternalValue() *ComputeRegionBackendServiceCircuitBreakers
	SetInternalValue(val *ComputeRegionBackendServiceCircuitBreakers)
	MaxConnections() *float64
	SetMaxConnections(val *float64)
	MaxConnectionsInput() *float64
	MaxPendingRequests() *float64
	SetMaxPendingRequests(val *float64)
	MaxPendingRequestsInput() *float64
	MaxRequests() *float64
	SetMaxRequests(val *float64)
	MaxRequestsInput() *float64
	MaxRequestsPerConnection() *float64
	SetMaxRequestsPerConnection(val *float64)
	MaxRequestsPerConnectionInput() *float64
	MaxRetries() *float64
	SetMaxRetries(val *float64)
	MaxRetriesInput() *float64
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
	ResetMaxConnections()
	ResetMaxPendingRequests()
	ResetMaxRequests()
	ResetMaxRequestsPerConnection()
	ResetMaxRetries()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeRegionBackendServiceCircuitBreakersOutputReference
type jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference) InternalValue() *ComputeRegionBackendServiceCircuitBreakers {
	var returns *ComputeRegionBackendServiceCircuitBreakers
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference) MaxConnections() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConnections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference) MaxConnectionsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConnectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference) MaxPendingRequests() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPendingRequests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference) MaxPendingRequestsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPendingRequestsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference) MaxRequests() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRequests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference) MaxRequestsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRequestsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference) MaxRequestsPerConnection() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRequestsPerConnection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference) MaxRequestsPerConnectionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRequestsPerConnectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference) MaxRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference) MaxRetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeRegionBackendServiceCircuitBreakersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeRegionBackendServiceCircuitBreakersOutputReference {
	_init_.Initialize()

	if err := validateNewComputeRegionBackendServiceCircuitBreakersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionBackendService.ComputeRegionBackendServiceCircuitBreakersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeRegionBackendServiceCircuitBreakersOutputReference_Override(c ComputeRegionBackendServiceCircuitBreakersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionBackendService.ComputeRegionBackendServiceCircuitBreakersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference)SetInternalValue(val *ComputeRegionBackendServiceCircuitBreakers) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference)SetMaxConnections(val *float64) {
	if err := j.validateSetMaxConnectionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConnections",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference)SetMaxPendingRequests(val *float64) {
	if err := j.validateSetMaxPendingRequestsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxPendingRequests",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference)SetMaxRequests(val *float64) {
	if err := j.validateSetMaxRequestsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRequests",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference)SetMaxRequestsPerConnection(val *float64) {
	if err := j.validateSetMaxRequestsPerConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRequestsPerConnection",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference)SetMaxRetries(val *float64) {
	if err := j.validateSetMaxRetriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRetries",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference) ResetMaxConnections() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxConnections",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference) ResetMaxPendingRequests() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxPendingRequests",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference) ResetMaxRequests() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxRequests",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference) ResetMaxRequestsPerConnection() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxRequestsPerConnection",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference) ResetMaxRetries() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxRetries",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeRegionBackendServiceCircuitBreakersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

