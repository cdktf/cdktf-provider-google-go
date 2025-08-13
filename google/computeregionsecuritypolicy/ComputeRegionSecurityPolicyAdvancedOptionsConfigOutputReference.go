// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregionsecuritypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/computeregionsecuritypolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference interface {
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
	InternalValue() *ComputeRegionSecurityPolicyAdvancedOptionsConfig
	SetInternalValue(val *ComputeRegionSecurityPolicyAdvancedOptionsConfig)
	JsonCustomConfig() ComputeRegionSecurityPolicyAdvancedOptionsConfigJsonCustomConfigOutputReference
	JsonCustomConfigInput() *ComputeRegionSecurityPolicyAdvancedOptionsConfigJsonCustomConfig
	JsonParsing() *string
	SetJsonParsing(val *string)
	JsonParsingInput() *string
	LogLevel() *string
	SetLogLevel(val *string)
	LogLevelInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserIpRequestHeaders() *[]*string
	SetUserIpRequestHeaders(val *[]*string)
	UserIpRequestHeadersInput() *[]*string
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
	PutJsonCustomConfig(value *ComputeRegionSecurityPolicyAdvancedOptionsConfigJsonCustomConfig)
	ResetJsonCustomConfig()
	ResetJsonParsing()
	ResetLogLevel()
	ResetUserIpRequestHeaders()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference
type jsiiProxy_ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) InternalValue() *ComputeRegionSecurityPolicyAdvancedOptionsConfig {
	var returns *ComputeRegionSecurityPolicyAdvancedOptionsConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) JsonCustomConfig() ComputeRegionSecurityPolicyAdvancedOptionsConfigJsonCustomConfigOutputReference {
	var returns ComputeRegionSecurityPolicyAdvancedOptionsConfigJsonCustomConfigOutputReference
	_jsii_.Get(
		j,
		"jsonCustomConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) JsonCustomConfigInput() *ComputeRegionSecurityPolicyAdvancedOptionsConfigJsonCustomConfig {
	var returns *ComputeRegionSecurityPolicyAdvancedOptionsConfigJsonCustomConfig
	_jsii_.Get(
		j,
		"jsonCustomConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) JsonParsing() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jsonParsing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) JsonParsingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jsonParsingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) LogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) LogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) UserIpRequestHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userIpRequestHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) UserIpRequestHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userIpRequestHeadersInput",
		&returns,
	)
	return returns
}


func NewComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionSecurityPolicy.ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference_Override(c ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionSecurityPolicy.ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference)SetInternalValue(val *ComputeRegionSecurityPolicyAdvancedOptionsConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference)SetJsonParsing(val *string) {
	if err := j.validateSetJsonParsingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jsonParsing",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference)SetLogLevel(val *string) {
	if err := j.validateSetLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logLevel",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference)SetUserIpRequestHeaders(val *[]*string) {
	if err := j.validateSetUserIpRequestHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userIpRequestHeaders",
		val,
	)
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) PutJsonCustomConfig(value *ComputeRegionSecurityPolicyAdvancedOptionsConfigJsonCustomConfig) {
	if err := c.validatePutJsonCustomConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putJsonCustomConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) ResetJsonCustomConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetJsonCustomConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) ResetJsonParsing() {
	_jsii_.InvokeVoid(
		c,
		"resetJsonParsing",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) ResetLogLevel() {
	_jsii_.InvokeVoid(
		c,
		"resetLogLevel",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) ResetUserIpRequestHeaders() {
	_jsii_.InvokeVoid(
		c,
		"resetUserIpRequestHeaders",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeRegionSecurityPolicyAdvancedOptionsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

