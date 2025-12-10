// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigeesecurityaction

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/apigeesecurityaction/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigeeSecurityActionConditionConfigOutputReference interface {
	cdktf.ComplexObject
	AccessTokens() *[]*string
	SetAccessTokens(val *[]*string)
	AccessTokensInput() *[]*string
	ApiKeys() *[]*string
	SetApiKeys(val *[]*string)
	ApiKeysInput() *[]*string
	ApiProducts() *[]*string
	SetApiProducts(val *[]*string)
	ApiProductsInput() *[]*string
	Asns() *[]*string
	SetAsns(val *[]*string)
	AsnsInput() *[]*string
	BotReasons() *[]*string
	SetBotReasons(val *[]*string)
	BotReasonsInput() *[]*string
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
	DeveloperApps() *[]*string
	SetDeveloperApps(val *[]*string)
	DeveloperAppsInput() *[]*string
	Developers() *[]*string
	SetDevelopers(val *[]*string)
	DevelopersInput() *[]*string
	// Experimental.
	Fqn() *string
	HttpMethods() *[]*string
	SetHttpMethods(val *[]*string)
	HttpMethodsInput() *[]*string
	InternalValue() *ApigeeSecurityActionConditionConfig
	SetInternalValue(val *ApigeeSecurityActionConditionConfig)
	IpAddressRanges() *[]*string
	SetIpAddressRanges(val *[]*string)
	IpAddressRangesInput() *[]*string
	RegionCodes() *[]*string
	SetRegionCodes(val *[]*string)
	RegionCodesInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserAgents() *[]*string
	SetUserAgents(val *[]*string)
	UserAgentsInput() *[]*string
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
	ResetAccessTokens()
	ResetApiKeys()
	ResetApiProducts()
	ResetAsns()
	ResetBotReasons()
	ResetDeveloperApps()
	ResetDevelopers()
	ResetHttpMethods()
	ResetIpAddressRanges()
	ResetRegionCodes()
	ResetUserAgents()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApigeeSecurityActionConditionConfigOutputReference
type jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) AccessTokens() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accessTokens",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) AccessTokensInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accessTokensInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) ApiKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"apiKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) ApiKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"apiKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) ApiProducts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"apiProducts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) ApiProductsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"apiProductsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) Asns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"asns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) AsnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"asnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) BotReasons() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"botReasons",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) BotReasonsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"botReasonsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) DeveloperApps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"developerApps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) DeveloperAppsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"developerAppsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) Developers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"developers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) DevelopersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"developersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) HttpMethods() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"httpMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) HttpMethodsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"httpMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) InternalValue() *ApigeeSecurityActionConditionConfig {
	var returns *ApigeeSecurityActionConditionConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) IpAddressRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipAddressRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) IpAddressRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipAddressRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) RegionCodes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regionCodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) RegionCodesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regionCodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) UserAgents() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userAgents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) UserAgentsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userAgentsInput",
		&returns,
	)
	return returns
}


func NewApigeeSecurityActionConditionConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApigeeSecurityActionConditionConfigOutputReference {
	_init_.Initialize()

	if err := validateNewApigeeSecurityActionConditionConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.apigeeSecurityAction.ApigeeSecurityActionConditionConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApigeeSecurityActionConditionConfigOutputReference_Override(a ApigeeSecurityActionConditionConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.apigeeSecurityAction.ApigeeSecurityActionConditionConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference)SetAccessTokens(val *[]*string) {
	if err := j.validateSetAccessTokensParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessTokens",
		val,
	)
}

func (j *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference)SetApiKeys(val *[]*string) {
	if err := j.validateSetApiKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiKeys",
		val,
	)
}

func (j *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference)SetApiProducts(val *[]*string) {
	if err := j.validateSetApiProductsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiProducts",
		val,
	)
}

func (j *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference)SetAsns(val *[]*string) {
	if err := j.validateSetAsnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"asns",
		val,
	)
}

func (j *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference)SetBotReasons(val *[]*string) {
	if err := j.validateSetBotReasonsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"botReasons",
		val,
	)
}

func (j *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference)SetDeveloperApps(val *[]*string) {
	if err := j.validateSetDeveloperAppsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"developerApps",
		val,
	)
}

func (j *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference)SetDevelopers(val *[]*string) {
	if err := j.validateSetDevelopersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"developers",
		val,
	)
}

func (j *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference)SetHttpMethods(val *[]*string) {
	if err := j.validateSetHttpMethodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpMethods",
		val,
	)
}

func (j *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference)SetInternalValue(val *ApigeeSecurityActionConditionConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference)SetIpAddressRanges(val *[]*string) {
	if err := j.validateSetIpAddressRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddressRanges",
		val,
	)
}

func (j *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference)SetRegionCodes(val *[]*string) {
	if err := j.validateSetRegionCodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regionCodes",
		val,
	)
}

func (j *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference)SetUserAgents(val *[]*string) {
	if err := j.validateSetUserAgentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userAgents",
		val,
	)
}

func (a *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) ResetAccessTokens() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessTokens",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) ResetApiKeys() {
	_jsii_.InvokeVoid(
		a,
		"resetApiKeys",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) ResetApiProducts() {
	_jsii_.InvokeVoid(
		a,
		"resetApiProducts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) ResetAsns() {
	_jsii_.InvokeVoid(
		a,
		"resetAsns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) ResetBotReasons() {
	_jsii_.InvokeVoid(
		a,
		"resetBotReasons",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) ResetDeveloperApps() {
	_jsii_.InvokeVoid(
		a,
		"resetDeveloperApps",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) ResetDevelopers() {
	_jsii_.InvokeVoid(
		a,
		"resetDevelopers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) ResetHttpMethods() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpMethods",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) ResetIpAddressRanges() {
	_jsii_.InvokeVoid(
		a,
		"resetIpAddressRanges",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) ResetRegionCodes() {
	_jsii_.InvokeVoid(
		a,
		"resetRegionCodes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) ResetUserAgents() {
	_jsii_.InvokeVoid(
		a,
		"resetUserAgents",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApigeeSecurityActionConditionConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

