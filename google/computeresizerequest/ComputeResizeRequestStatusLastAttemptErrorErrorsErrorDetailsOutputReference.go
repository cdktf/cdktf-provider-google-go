// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeresizerequest

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/computeresizerequest/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference interface {
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
	ErrorInfo() ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsErrorInfoList
	// Experimental.
	Fqn() *string
	Help() ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsHelpList
	InternalValue() *ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetails
	SetInternalValue(val *ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetails)
	LocalizedMessage() ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsLocalizedMessageList
	QuotaInfo() ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoList
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

// The jsii proxy struct for ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference
type jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) ErrorInfo() ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsErrorInfoList {
	var returns ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsErrorInfoList
	_jsii_.Get(
		j,
		"errorInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) Help() ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsHelpList {
	var returns ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsHelpList
	_jsii_.Get(
		j,
		"help",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) InternalValue() *ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetails {
	var returns *ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetails
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) LocalizedMessage() ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsLocalizedMessageList {
	var returns ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsLocalizedMessageList
	_jsii_.Get(
		j,
		"localizedMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) QuotaInfo() ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoList {
	var returns ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsQuotaInfoList
	_jsii_.Get(
		j,
		"quotaInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference {
	_init_.Initialize()

	if err := validateNewComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeResizeRequest.ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference_Override(c ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeResizeRequest.ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference)SetInternalValue(val *ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetails) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeResizeRequestStatusLastAttemptErrorErrorsErrorDetailsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

