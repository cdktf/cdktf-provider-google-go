package networkservicesedgecacheservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v4/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v4/networkservicesedgecacheservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference interface {
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
	ExcludedQueryParameters() *[]*string
	SetExcludedQueryParameters(val *[]*string)
	ExcludedQueryParametersInput() *[]*string
	ExcludeHost() interface{}
	SetExcludeHost(val interface{})
	ExcludeHostInput() interface{}
	ExcludeQueryString() interface{}
	SetExcludeQueryString(val interface{})
	ExcludeQueryStringInput() interface{}
	// Experimental.
	Fqn() *string
	IncludedCookieNames() *[]*string
	SetIncludedCookieNames(val *[]*string)
	IncludedCookieNamesInput() *[]*string
	IncludedHeaderNames() *[]*string
	SetIncludedHeaderNames(val *[]*string)
	IncludedHeaderNamesInput() *[]*string
	IncludedQueryParameters() *[]*string
	SetIncludedQueryParameters(val *[]*string)
	IncludedQueryParametersInput() *[]*string
	IncludeProtocol() interface{}
	SetIncludeProtocol(val interface{})
	IncludeProtocolInput() interface{}
	InternalValue() *NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicy
	SetInternalValue(val *NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicy)
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
	ResetExcludedQueryParameters()
	ResetExcludeHost()
	ResetExcludeQueryString()
	ResetIncludedCookieNames()
	ResetIncludedHeaderNames()
	ResetIncludedQueryParameters()
	ResetIncludeProtocol()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference
type jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference) ExcludedQueryParameters() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedQueryParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference) ExcludedQueryParametersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedQueryParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference) ExcludeHost() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference) ExcludeHostInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeHostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference) ExcludeQueryString() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeQueryString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference) ExcludeQueryStringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeQueryStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference) IncludedCookieNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedCookieNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference) IncludedCookieNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedCookieNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference) IncludedHeaderNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedHeaderNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference) IncludedHeaderNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedHeaderNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference) IncludedQueryParameters() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedQueryParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference) IncludedQueryParametersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedQueryParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference) IncludeProtocol() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference) IncludeProtocolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference) InternalValue() *NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicy {
	var returns *NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.networkServicesEdgeCacheService.NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference_Override(n NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.networkServicesEdgeCacheService.NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference)SetExcludedQueryParameters(val *[]*string) {
	if err := j.validateSetExcludedQueryParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedQueryParameters",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference)SetExcludeHost(val interface{}) {
	if err := j.validateSetExcludeHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeHost",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference)SetExcludeQueryString(val interface{}) {
	if err := j.validateSetExcludeQueryStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeQueryString",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference)SetIncludedCookieNames(val *[]*string) {
	if err := j.validateSetIncludedCookieNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includedCookieNames",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference)SetIncludedHeaderNames(val *[]*string) {
	if err := j.validateSetIncludedHeaderNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includedHeaderNames",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference)SetIncludedQueryParameters(val *[]*string) {
	if err := j.validateSetIncludedQueryParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includedQueryParameters",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference)SetIncludeProtocol(val interface{}) {
	if err := j.validateSetIncludeProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeProtocol",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference)SetInternalValue(val *NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference) ResetExcludedQueryParameters() {
	_jsii_.InvokeVoid(
		n,
		"resetExcludedQueryParameters",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference) ResetExcludeHost() {
	_jsii_.InvokeVoid(
		n,
		"resetExcludeHost",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference) ResetExcludeQueryString() {
	_jsii_.InvokeVoid(
		n,
		"resetExcludeQueryString",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference) ResetIncludedCookieNames() {
	_jsii_.InvokeVoid(
		n,
		"resetIncludedCookieNames",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference) ResetIncludedHeaderNames() {
	_jsii_.InvokeVoid(
		n,
		"resetIncludedHeaderNames",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference) ResetIncludedQueryParameters() {
	_jsii_.InvokeVoid(
		n,
		"resetIncludedQueryParameters",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference) ResetIncludeProtocol() {
	_jsii_.InvokeVoid(
		n,
		"resetIncludeProtocol",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (n *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleRouteActionCdnPolicyCacheKeyPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

