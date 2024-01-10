// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudrunservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v13/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v13/cloudrunservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference interface {
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
	FailureThreshold() *float64
	SetFailureThreshold(val *float64)
	FailureThresholdInput() *float64
	// Experimental.
	Fqn() *string
	Grpc() CloudRunServiceTemplateSpecContainersLivenessProbeGrpcOutputReference
	GrpcInput() *CloudRunServiceTemplateSpecContainersLivenessProbeGrpc
	HttpGet() CloudRunServiceTemplateSpecContainersLivenessProbeHttpGetOutputReference
	HttpGetInput() *CloudRunServiceTemplateSpecContainersLivenessProbeHttpGet
	InitialDelaySeconds() *float64
	SetInitialDelaySeconds(val *float64)
	InitialDelaySecondsInput() *float64
	InternalValue() *CloudRunServiceTemplateSpecContainersLivenessProbe
	SetInternalValue(val *CloudRunServiceTemplateSpecContainersLivenessProbe)
	PeriodSeconds() *float64
	SetPeriodSeconds(val *float64)
	PeriodSecondsInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeoutSeconds() *float64
	SetTimeoutSeconds(val *float64)
	TimeoutSecondsInput() *float64
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
	PutGrpc(value *CloudRunServiceTemplateSpecContainersLivenessProbeGrpc)
	PutHttpGet(value *CloudRunServiceTemplateSpecContainersLivenessProbeHttpGet)
	ResetFailureThreshold()
	ResetGrpc()
	ResetHttpGet()
	ResetInitialDelaySeconds()
	ResetPeriodSeconds()
	ResetTimeoutSeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference
type jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference) FailureThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference) FailureThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference) Grpc() CloudRunServiceTemplateSpecContainersLivenessProbeGrpcOutputReference {
	var returns CloudRunServiceTemplateSpecContainersLivenessProbeGrpcOutputReference
	_jsii_.Get(
		j,
		"grpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference) GrpcInput() *CloudRunServiceTemplateSpecContainersLivenessProbeGrpc {
	var returns *CloudRunServiceTemplateSpecContainersLivenessProbeGrpc
	_jsii_.Get(
		j,
		"grpcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference) HttpGet() CloudRunServiceTemplateSpecContainersLivenessProbeHttpGetOutputReference {
	var returns CloudRunServiceTemplateSpecContainersLivenessProbeHttpGetOutputReference
	_jsii_.Get(
		j,
		"httpGet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference) HttpGetInput() *CloudRunServiceTemplateSpecContainersLivenessProbeHttpGet {
	var returns *CloudRunServiceTemplateSpecContainersLivenessProbeHttpGet
	_jsii_.Get(
		j,
		"httpGetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference) InitialDelaySeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialDelaySeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference) InitialDelaySecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialDelaySecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference) InternalValue() *CloudRunServiceTemplateSpecContainersLivenessProbe {
	var returns *CloudRunServiceTemplateSpecContainersLivenessProbe
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference) PeriodSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"periodSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference) PeriodSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"periodSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference) TimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference) TimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSecondsInput",
		&returns,
	)
	return returns
}


func NewCloudRunServiceTemplateSpecContainersLivenessProbeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference {
	_init_.Initialize()

	if err := validateNewCloudRunServiceTemplateSpecContainersLivenessProbeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cloudRunService.CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudRunServiceTemplateSpecContainersLivenessProbeOutputReference_Override(c CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cloudRunService.CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference)SetFailureThreshold(val *float64) {
	if err := j.validateSetFailureThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failureThreshold",
		val,
	)
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference)SetInitialDelaySeconds(val *float64) {
	if err := j.validateSetInitialDelaySecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialDelaySeconds",
		val,
	)
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference)SetInternalValue(val *CloudRunServiceTemplateSpecContainersLivenessProbe) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference)SetPeriodSeconds(val *float64) {
	if err := j.validateSetPeriodSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"periodSeconds",
		val,
	)
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference)SetTimeoutSeconds(val *float64) {
	if err := j.validateSetTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutSeconds",
		val,
	)
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference) PutGrpc(value *CloudRunServiceTemplateSpecContainersLivenessProbeGrpc) {
	if err := c.validatePutGrpcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGrpc",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference) PutHttpGet(value *CloudRunServiceTemplateSpecContainersLivenessProbeHttpGet) {
	if err := c.validatePutHttpGetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHttpGet",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference) ResetFailureThreshold() {
	_jsii_.InvokeVoid(
		c,
		"resetFailureThreshold",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference) ResetGrpc() {
	_jsii_.InvokeVoid(
		c,
		"resetGrpc",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference) ResetHttpGet() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpGet",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference) ResetInitialDelaySeconds() {
	_jsii_.InvokeVoid(
		c,
		"resetInitialDelaySeconds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference) ResetPeriodSeconds() {
	_jsii_.InvokeVoid(
		c,
		"resetPeriodSeconds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference) ResetTimeoutSeconds() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeoutSeconds",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CloudRunServiceTemplateSpecContainersLivenessProbeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

