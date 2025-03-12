// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkservicesgrpcroute

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/networkservicesgrpcroute/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference interface {
	cdktf.ComplexObject
	Abort() NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyAbortOutputReference
	AbortInput() *NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyAbort
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
	Delay() NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyDelayOutputReference
	DelayInput() *NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyDelay
	// Experimental.
	Fqn() *string
	InternalValue() *NetworkServicesGrpcRouteRulesActionFaultInjectionPolicy
	SetInternalValue(val *NetworkServicesGrpcRouteRulesActionFaultInjectionPolicy)
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
	PutAbort(value *NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyAbort)
	PutDelay(value *NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyDelay)
	ResetAbort()
	ResetDelay()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference
type jsiiProxy_NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) Abort() NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyAbortOutputReference {
	var returns NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyAbortOutputReference
	_jsii_.Get(
		j,
		"abort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) AbortInput() *NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyAbort {
	var returns *NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyAbort
	_jsii_.Get(
		j,
		"abortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) Delay() NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyDelayOutputReference {
	var returns NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyDelayOutputReference
	_jsii_.Get(
		j,
		"delay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) DelayInput() *NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyDelay {
	var returns *NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyDelay
	_jsii_.Get(
		j,
		"delayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) InternalValue() *NetworkServicesGrpcRouteRulesActionFaultInjectionPolicy {
	var returns *NetworkServicesGrpcRouteRulesActionFaultInjectionPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.networkServicesGrpcRoute.NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference_Override(n NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.networkServicesGrpcRoute.NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference)SetInternalValue(val *NetworkServicesGrpcRouteRulesActionFaultInjectionPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) PutAbort(value *NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyAbort) {
	if err := n.validatePutAbortParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putAbort",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) PutDelay(value *NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyDelay) {
	if err := n.validatePutDelayParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putDelay",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) ResetAbort() {
	_jsii_.InvokeVoid(
		n,
		"resetAbort",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) ResetDelay() {
	_jsii_.InvokeVoid(
		n,
		"resetDelay",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (n *jsiiProxy_NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

