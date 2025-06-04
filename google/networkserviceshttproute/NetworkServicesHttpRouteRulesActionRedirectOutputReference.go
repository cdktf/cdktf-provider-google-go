// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkserviceshttproute

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/networkserviceshttproute/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkServicesHttpRouteRulesActionRedirectOutputReference interface {
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
	HostRedirect() *string
	SetHostRedirect(val *string)
	HostRedirectInput() *string
	HttpsRedirect() interface{}
	SetHttpsRedirect(val interface{})
	HttpsRedirectInput() interface{}
	InternalValue() *NetworkServicesHttpRouteRulesActionRedirect
	SetInternalValue(val *NetworkServicesHttpRouteRulesActionRedirect)
	PathRedirect() *string
	SetPathRedirect(val *string)
	PathRedirectInput() *string
	PortRedirect() *float64
	SetPortRedirect(val *float64)
	PortRedirectInput() *float64
	PrefixRewrite() *string
	SetPrefixRewrite(val *string)
	PrefixRewriteInput() *string
	ResponseCode() *string
	SetResponseCode(val *string)
	ResponseCodeInput() *string
	StripQuery() interface{}
	SetStripQuery(val interface{})
	StripQueryInput() interface{}
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
	ResetHostRedirect()
	ResetHttpsRedirect()
	ResetPathRedirect()
	ResetPortRedirect()
	ResetPrefixRewrite()
	ResetResponseCode()
	ResetStripQuery()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkServicesHttpRouteRulesActionRedirectOutputReference
type jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference) HostRedirect() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostRedirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference) HostRedirectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostRedirectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference) HttpsRedirect() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpsRedirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference) HttpsRedirectInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpsRedirectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference) InternalValue() *NetworkServicesHttpRouteRulesActionRedirect {
	var returns *NetworkServicesHttpRouteRulesActionRedirect
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference) PathRedirect() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathRedirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference) PathRedirectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathRedirectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference) PortRedirect() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portRedirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference) PortRedirectInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portRedirectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference) PrefixRewrite() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixRewrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference) PrefixRewriteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixRewriteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference) ResponseCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference) ResponseCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference) StripQuery() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stripQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference) StripQueryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stripQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNetworkServicesHttpRouteRulesActionRedirectOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NetworkServicesHttpRouteRulesActionRedirectOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkServicesHttpRouteRulesActionRedirectOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.networkServicesHttpRoute.NetworkServicesHttpRouteRulesActionRedirectOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetworkServicesHttpRouteRulesActionRedirectOutputReference_Override(n NetworkServicesHttpRouteRulesActionRedirectOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.networkServicesHttpRoute.NetworkServicesHttpRouteRulesActionRedirectOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference)SetHostRedirect(val *string) {
	if err := j.validateSetHostRedirectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostRedirect",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference)SetHttpsRedirect(val interface{}) {
	if err := j.validateSetHttpsRedirectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpsRedirect",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference)SetInternalValue(val *NetworkServicesHttpRouteRulesActionRedirect) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference)SetPathRedirect(val *string) {
	if err := j.validateSetPathRedirectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pathRedirect",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference)SetPortRedirect(val *float64) {
	if err := j.validateSetPortRedirectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portRedirect",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference)SetPrefixRewrite(val *string) {
	if err := j.validateSetPrefixRewriteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefixRewrite",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference)SetResponseCode(val *string) {
	if err := j.validateSetResponseCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"responseCode",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference)SetStripQuery(val interface{}) {
	if err := j.validateSetStripQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stripQuery",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference) ResetHostRedirect() {
	_jsii_.InvokeVoid(
		n,
		"resetHostRedirect",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference) ResetHttpsRedirect() {
	_jsii_.InvokeVoid(
		n,
		"resetHttpsRedirect",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference) ResetPathRedirect() {
	_jsii_.InvokeVoid(
		n,
		"resetPathRedirect",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference) ResetPortRedirect() {
	_jsii_.InvokeVoid(
		n,
		"resetPortRedirect",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference) ResetPrefixRewrite() {
	_jsii_.InvokeVoid(
		n,
		"resetPrefixRewrite",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference) ResetResponseCode() {
	_jsii_.InvokeVoid(
		n,
		"resetResponseCode",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference) ResetStripQuery() {
	_jsii_.InvokeVoid(
		n,
		"resetStripQuery",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := n.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionRedirectOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

