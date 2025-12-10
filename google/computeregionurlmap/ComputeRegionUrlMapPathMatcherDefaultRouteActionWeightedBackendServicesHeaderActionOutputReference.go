// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregionurlmap

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/computeregionurlmap/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference interface {
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
	InternalValue() *ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderAction
	SetInternalValue(val *ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderAction)
	RequestHeadersToAdd() ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddList
	RequestHeadersToAddInput() interface{}
	RequestHeadersToRemove() *[]*string
	SetRequestHeadersToRemove(val *[]*string)
	RequestHeadersToRemoveInput() *[]*string
	ResponseHeadersToAdd() ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionResponseHeadersToAddList
	ResponseHeadersToAddInput() interface{}
	ResponseHeadersToRemove() *[]*string
	SetResponseHeadersToRemove(val *[]*string)
	ResponseHeadersToRemoveInput() *[]*string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutRequestHeadersToAdd(value interface{})
	PutResponseHeadersToAdd(value interface{})
	ResetRequestHeadersToAdd()
	ResetRequestHeadersToRemove()
	ResetResponseHeadersToAdd()
	ResetResponseHeadersToRemove()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference
type jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference) InternalValue() *ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderAction {
	var returns *ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderAction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference) RequestHeadersToAdd() ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddList {
	var returns ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionRequestHeadersToAddList
	_jsii_.Get(
		j,
		"requestHeadersToAdd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference) RequestHeadersToAddInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestHeadersToAddInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference) RequestHeadersToRemove() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requestHeadersToRemove",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference) RequestHeadersToRemoveInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requestHeadersToRemoveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference) ResponseHeadersToAdd() ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionResponseHeadersToAddList {
	var returns ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionResponseHeadersToAddList
	_jsii_.Get(
		j,
		"responseHeadersToAdd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference) ResponseHeadersToAddInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"responseHeadersToAddInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference) ResponseHeadersToRemove() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"responseHeadersToRemove",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference) ResponseHeadersToRemoveInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"responseHeadersToRemoveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference {
	_init_.Initialize()

	if err := validateNewComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference_Override(c ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference)SetInternalValue(val *ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderAction) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference)SetRequestHeadersToRemove(val *[]*string) {
	if err := j.validateSetRequestHeadersToRemoveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestHeadersToRemove",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference)SetResponseHeadersToRemove(val *[]*string) {
	if err := j.validateSetResponseHeadersToRemoveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseHeadersToRemove",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference) PutRequestHeadersToAdd(value interface{}) {
	if err := c.validatePutRequestHeadersToAddParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRequestHeadersToAdd",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference) PutResponseHeadersToAdd(value interface{}) {
	if err := c.validatePutResponseHeadersToAddParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putResponseHeadersToAdd",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference) ResetRequestHeadersToAdd() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestHeadersToAdd",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference) ResetRequestHeadersToRemove() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestHeadersToRemove",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference) ResetResponseHeadersToAdd() {
	_jsii_.InvokeVoid(
		c,
		"resetResponseHeadersToAdd",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference) ResetResponseHeadersToRemove() {
	_jsii_.InvokeVoid(
		c,
		"resetResponseHeadersToRemove",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherDefaultRouteActionWeightedBackendServicesHeaderActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

