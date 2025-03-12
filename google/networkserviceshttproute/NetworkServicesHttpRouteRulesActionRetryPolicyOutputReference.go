// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkserviceshttproute

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/networkserviceshttproute/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference interface {
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
	InternalValue() *NetworkServicesHttpRouteRulesActionRetryPolicy
	SetInternalValue(val *NetworkServicesHttpRouteRulesActionRetryPolicy)
	NumRetries() *float64
	SetNumRetries(val *float64)
	NumRetriesInput() *float64
	PerTryTimeout() *string
	SetPerTryTimeout(val *string)
	PerTryTimeoutInput() *string
	RetryConditions() *[]*string
	SetRetryConditions(val *[]*string)
	RetryConditionsInput() *[]*string
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
	ResetNumRetries()
	ResetPerTryTimeout()
	ResetRetryConditions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference
type jsiiProxy_NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference) InternalValue() *NetworkServicesHttpRouteRulesActionRetryPolicy {
	var returns *NetworkServicesHttpRouteRulesActionRetryPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference) NumRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference) NumRetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numRetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference) PerTryTimeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"perTryTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference) PerTryTimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"perTryTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference) RetryConditions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"retryConditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference) RetryConditionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"retryConditionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNetworkServicesHttpRouteRulesActionRetryPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkServicesHttpRouteRulesActionRetryPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.networkServicesHttpRoute.NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetworkServicesHttpRouteRulesActionRetryPolicyOutputReference_Override(n NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.networkServicesHttpRoute.NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference)SetInternalValue(val *NetworkServicesHttpRouteRulesActionRetryPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference)SetNumRetries(val *float64) {
	if err := j.validateSetNumRetriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numRetries",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference)SetPerTryTimeout(val *string) {
	if err := j.validateSetPerTryTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"perTryTimeout",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference)SetRetryConditions(val *[]*string) {
	if err := j.validateSetRetryConditionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryConditions",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference) ResetNumRetries() {
	_jsii_.InvokeVoid(
		n,
		"resetNumRetries",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference) ResetPerTryTimeout() {
	_jsii_.InvokeVoid(
		n,
		"resetPerTryTimeout",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference) ResetRetryConditions() {
	_jsii_.InvokeVoid(
		n,
		"resetRetryConditions",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (n *jsiiProxy_NetworkServicesHttpRouteRulesActionRetryPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

