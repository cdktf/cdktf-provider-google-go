// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeresizerequest

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/computeresizerequest/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoOutputReference interface {
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
	Dimensions() cdktf.StringMap
	// Experimental.
	Fqn() *string
	FutureLimit() *float64
	InternalValue() *ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfo
	SetInternalValue(val *ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfo)
	Limit() *float64
	LimitName() *string
	MetricName() *string
	RolloutStatus() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoOutputReference
type jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoOutputReference) Dimensions() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"dimensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoOutputReference) FutureLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"futureLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoOutputReference) InternalValue() *ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfo {
	var returns *ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfo
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoOutputReference) Limit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"limit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoOutputReference) LimitName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"limitName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoOutputReference) MetricName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoOutputReference) RolloutStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rolloutStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoOutputReference {
	_init_.Initialize()

	if err := validateNewComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeResizeRequest.ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoOutputReference_Override(c ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeResizeRequest.ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoOutputReference)SetInternalValue(val *ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfo) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

