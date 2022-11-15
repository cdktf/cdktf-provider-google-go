package networkservicesedgecacheservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v4/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v4/networkservicesedgecacheservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference interface {
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
	FullPathMatch() *string
	SetFullPathMatch(val *string)
	FullPathMatchInput() *string
	HeaderMatch() NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleHeaderMatchList
	HeaderMatchInput() interface{}
	IgnoreCase() interface{}
	SetIgnoreCase(val interface{})
	IgnoreCaseInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PathTemplateMatch() *string
	SetPathTemplateMatch(val *string)
	PathTemplateMatchInput() *string
	PrefixMatch() *string
	SetPrefixMatch(val *string)
	PrefixMatchInput() *string
	QueryParameterMatch() NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleQueryParameterMatchList
	QueryParameterMatchInput() interface{}
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
	PutHeaderMatch(value interface{})
	PutQueryParameterMatch(value interface{})
	ResetFullPathMatch()
	ResetHeaderMatch()
	ResetIgnoreCase()
	ResetPathTemplateMatch()
	ResetPrefixMatch()
	ResetQueryParameterMatch()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference
type jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference) FullPathMatch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullPathMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference) FullPathMatchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullPathMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference) HeaderMatch() NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleHeaderMatchList {
	var returns NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleHeaderMatchList
	_jsii_.Get(
		j,
		"headerMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference) HeaderMatchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headerMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference) IgnoreCase() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreCase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference) IgnoreCaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreCaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference) PathTemplateMatch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathTemplateMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference) PathTemplateMatchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathTemplateMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference) PrefixMatch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference) PrefixMatchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference) QueryParameterMatch() NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleQueryParameterMatchList {
	var returns NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleQueryParameterMatchList
	_jsii_.Get(
		j,
		"queryParameterMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference) QueryParameterMatchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryParameterMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.networkServicesEdgeCacheService.NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference_Override(n NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.networkServicesEdgeCacheService.NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		n,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference)SetFullPathMatch(val *string) {
	if err := j.validateSetFullPathMatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fullPathMatch",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference)SetIgnoreCase(val interface{}) {
	if err := j.validateSetIgnoreCaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreCase",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference)SetPathTemplateMatch(val *string) {
	if err := j.validateSetPathTemplateMatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pathTemplateMatch",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference)SetPrefixMatch(val *string) {
	if err := j.validateSetPrefixMatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefixMatch",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference) PutHeaderMatch(value interface{}) {
	if err := n.validatePutHeaderMatchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putHeaderMatch",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference) PutQueryParameterMatch(value interface{}) {
	if err := n.validatePutQueryParameterMatchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putQueryParameterMatch",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference) ResetFullPathMatch() {
	_jsii_.InvokeVoid(
		n,
		"resetFullPathMatch",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference) ResetHeaderMatch() {
	_jsii_.InvokeVoid(
		n,
		"resetHeaderMatch",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference) ResetIgnoreCase() {
	_jsii_.InvokeVoid(
		n,
		"resetIgnoreCase",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference) ResetPathTemplateMatch() {
	_jsii_.InvokeVoid(
		n,
		"resetPathTemplateMatch",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference) ResetPrefixMatch() {
	_jsii_.InvokeVoid(
		n,
		"resetPrefixMatch",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference) ResetQueryParameterMatch() {
	_jsii_.InvokeVoid(
		n,
		"resetQueryParameterMatch",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (n *jsiiProxy_NetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleMatchRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

