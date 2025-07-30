// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkconnectivityinternalrange

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/networkconnectivityinternalrange/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkConnectivityInternalRangeAllocationOptionsOutputReference interface {
	cdktf.ComplexObject
	AllocationStrategy() *string
	SetAllocationStrategy(val *string)
	AllocationStrategyInput() *string
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
	FirstAvailableRangesLookupSize() *float64
	SetFirstAvailableRangesLookupSize(val *float64)
	FirstAvailableRangesLookupSizeInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *NetworkConnectivityInternalRangeAllocationOptions
	SetInternalValue(val *NetworkConnectivityInternalRangeAllocationOptions)
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
	ResetAllocationStrategy()
	ResetFirstAvailableRangesLookupSize()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkConnectivityInternalRangeAllocationOptionsOutputReference
type jsiiProxy_NetworkConnectivityInternalRangeAllocationOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetworkConnectivityInternalRangeAllocationOptionsOutputReference) AllocationStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocationStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityInternalRangeAllocationOptionsOutputReference) AllocationStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocationStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityInternalRangeAllocationOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityInternalRangeAllocationOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityInternalRangeAllocationOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityInternalRangeAllocationOptionsOutputReference) FirstAvailableRangesLookupSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"firstAvailableRangesLookupSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityInternalRangeAllocationOptionsOutputReference) FirstAvailableRangesLookupSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"firstAvailableRangesLookupSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityInternalRangeAllocationOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityInternalRangeAllocationOptionsOutputReference) InternalValue() *NetworkConnectivityInternalRangeAllocationOptions {
	var returns *NetworkConnectivityInternalRangeAllocationOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityInternalRangeAllocationOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityInternalRangeAllocationOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNetworkConnectivityInternalRangeAllocationOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NetworkConnectivityInternalRangeAllocationOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkConnectivityInternalRangeAllocationOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkConnectivityInternalRangeAllocationOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.networkConnectivityInternalRange.NetworkConnectivityInternalRangeAllocationOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetworkConnectivityInternalRangeAllocationOptionsOutputReference_Override(n NetworkConnectivityInternalRangeAllocationOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.networkConnectivityInternalRange.NetworkConnectivityInternalRangeAllocationOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetworkConnectivityInternalRangeAllocationOptionsOutputReference)SetAllocationStrategy(val *string) {
	if err := j.validateSetAllocationStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocationStrategy",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivityInternalRangeAllocationOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivityInternalRangeAllocationOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivityInternalRangeAllocationOptionsOutputReference)SetFirstAvailableRangesLookupSize(val *float64) {
	if err := j.validateSetFirstAvailableRangesLookupSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firstAvailableRangesLookupSize",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivityInternalRangeAllocationOptionsOutputReference)SetInternalValue(val *NetworkConnectivityInternalRangeAllocationOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivityInternalRangeAllocationOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivityInternalRangeAllocationOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetworkConnectivityInternalRangeAllocationOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectivityInternalRangeAllocationOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NetworkConnectivityInternalRangeAllocationOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkConnectivityInternalRangeAllocationOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NetworkConnectivityInternalRangeAllocationOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NetworkConnectivityInternalRangeAllocationOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NetworkConnectivityInternalRangeAllocationOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NetworkConnectivityInternalRangeAllocationOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NetworkConnectivityInternalRangeAllocationOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NetworkConnectivityInternalRangeAllocationOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NetworkConnectivityInternalRangeAllocationOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectivityInternalRangeAllocationOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkConnectivityInternalRangeAllocationOptionsOutputReference) ResetAllocationStrategy() {
	_jsii_.InvokeVoid(
		n,
		"resetAllocationStrategy",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectivityInternalRangeAllocationOptionsOutputReference) ResetFirstAvailableRangesLookupSize() {
	_jsii_.InvokeVoid(
		n,
		"resetFirstAvailableRangesLookupSize",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectivityInternalRangeAllocationOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (n *jsiiProxy_NetworkConnectivityInternalRangeAllocationOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

