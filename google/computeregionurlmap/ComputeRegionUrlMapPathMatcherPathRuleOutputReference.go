// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregionurlmap

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/computeregionurlmap/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeRegionUrlMapPathMatcherPathRuleOutputReference interface {
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Paths() *[]*string
	SetPaths(val *[]*string)
	PathsInput() *[]*string
	RouteAction() ComputeRegionUrlMapPathMatcherPathRuleRouteActionOutputReference
	RouteActionInput() *ComputeRegionUrlMapPathMatcherPathRuleRouteAction
	Service() *string
	SetService(val *string)
	ServiceInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UrlRedirect() ComputeRegionUrlMapPathMatcherPathRuleUrlRedirectOutputReference
	UrlRedirectInput() *ComputeRegionUrlMapPathMatcherPathRuleUrlRedirect
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
	PutRouteAction(value *ComputeRegionUrlMapPathMatcherPathRuleRouteAction)
	PutUrlRedirect(value *ComputeRegionUrlMapPathMatcherPathRuleUrlRedirect)
	ResetRouteAction()
	ResetService()
	ResetUrlRedirect()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeRegionUrlMapPathMatcherPathRuleOutputReference
type jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleOutputReference) Paths() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"paths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleOutputReference) PathsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pathsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleOutputReference) RouteAction() ComputeRegionUrlMapPathMatcherPathRuleRouteActionOutputReference {
	var returns ComputeRegionUrlMapPathMatcherPathRuleRouteActionOutputReference
	_jsii_.Get(
		j,
		"routeAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleOutputReference) RouteActionInput() *ComputeRegionUrlMapPathMatcherPathRuleRouteAction {
	var returns *ComputeRegionUrlMapPathMatcherPathRuleRouteAction
	_jsii_.Get(
		j,
		"routeActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleOutputReference) Service() *string {
	var returns *string
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleOutputReference) ServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleOutputReference) UrlRedirect() ComputeRegionUrlMapPathMatcherPathRuleUrlRedirectOutputReference {
	var returns ComputeRegionUrlMapPathMatcherPathRuleUrlRedirectOutputReference
	_jsii_.Get(
		j,
		"urlRedirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleOutputReference) UrlRedirectInput() *ComputeRegionUrlMapPathMatcherPathRuleUrlRedirect {
	var returns *ComputeRegionUrlMapPathMatcherPathRuleUrlRedirect
	_jsii_.Get(
		j,
		"urlRedirectInput",
		&returns,
	)
	return returns
}


func NewComputeRegionUrlMapPathMatcherPathRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ComputeRegionUrlMapPathMatcherPathRuleOutputReference {
	_init_.Initialize()

	if err := validateNewComputeRegionUrlMapPathMatcherPathRuleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherPathRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewComputeRegionUrlMapPathMatcherPathRuleOutputReference_Override(c ComputeRegionUrlMapPathMatcherPathRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionUrlMap.ComputeRegionUrlMapPathMatcherPathRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleOutputReference)SetPaths(val *[]*string) {
	if err := j.validateSetPathsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"paths",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleOutputReference)SetService(val *string) {
	if err := j.validateSetServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"service",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleOutputReference) PutRouteAction(value *ComputeRegionUrlMapPathMatcherPathRuleRouteAction) {
	if err := c.validatePutRouteActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRouteAction",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleOutputReference) PutUrlRedirect(value *ComputeRegionUrlMapPathMatcherPathRuleUrlRedirect) {
	if err := c.validatePutUrlRedirectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putUrlRedirect",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleOutputReference) ResetRouteAction() {
	_jsii_.InvokeVoid(
		c,
		"resetRouteAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleOutputReference) ResetService() {
	_jsii_.InvokeVoid(
		c,
		"resetService",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleOutputReference) ResetUrlRedirect() {
	_jsii_.InvokeVoid(
		c,
		"resetUrlRedirect",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeRegionUrlMapPathMatcherPathRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

