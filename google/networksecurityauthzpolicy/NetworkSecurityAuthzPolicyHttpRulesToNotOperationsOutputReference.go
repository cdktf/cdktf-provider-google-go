// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networksecurityauthzpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v15/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v15/networksecurityauthzpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference interface {
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
	HeaderSet() NetworkSecurityAuthzPolicyHttpRulesToNotOperationsHeaderSetOutputReference
	HeaderSetInput() *NetworkSecurityAuthzPolicyHttpRulesToNotOperationsHeaderSet
	Hosts() NetworkSecurityAuthzPolicyHttpRulesToNotOperationsHostsList
	HostsInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Methods() *[]*string
	SetMethods(val *[]*string)
	MethodsInput() *[]*string
	Paths() NetworkSecurityAuthzPolicyHttpRulesToNotOperationsPathsList
	PathsInput() interface{}
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
	PutHeaderSet(value *NetworkSecurityAuthzPolicyHttpRulesToNotOperationsHeaderSet)
	PutHosts(value interface{})
	PutPaths(value interface{})
	ResetHeaderSet()
	ResetHosts()
	ResetMethods()
	ResetPaths()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference
type jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference) HeaderSet() NetworkSecurityAuthzPolicyHttpRulesToNotOperationsHeaderSetOutputReference {
	var returns NetworkSecurityAuthzPolicyHttpRulesToNotOperationsHeaderSetOutputReference
	_jsii_.Get(
		j,
		"headerSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference) HeaderSetInput() *NetworkSecurityAuthzPolicyHttpRulesToNotOperationsHeaderSet {
	var returns *NetworkSecurityAuthzPolicyHttpRulesToNotOperationsHeaderSet
	_jsii_.Get(
		j,
		"headerSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference) Hosts() NetworkSecurityAuthzPolicyHttpRulesToNotOperationsHostsList {
	var returns NetworkSecurityAuthzPolicyHttpRulesToNotOperationsHostsList
	_jsii_.Get(
		j,
		"hosts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference) HostsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference) Methods() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"methods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference) MethodsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"methodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference) Paths() NetworkSecurityAuthzPolicyHttpRulesToNotOperationsPathsList {
	var returns NetworkSecurityAuthzPolicyHttpRulesToNotOperationsPathsList
	_jsii_.Get(
		j,
		"paths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference) PathsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pathsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.networkSecurityAuthzPolicy.NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewNetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference_Override(n NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.networkSecurityAuthzPolicy.NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		n,
	)
}

func (j *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference)SetMethods(val *[]*string) {
	if err := j.validateSetMethodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"methods",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference) PutHeaderSet(value *NetworkSecurityAuthzPolicyHttpRulesToNotOperationsHeaderSet) {
	if err := n.validatePutHeaderSetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putHeaderSet",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference) PutHosts(value interface{}) {
	if err := n.validatePutHostsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putHosts",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference) PutPaths(value interface{}) {
	if err := n.validatePutPathsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putPaths",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference) ResetHeaderSet() {
	_jsii_.InvokeVoid(
		n,
		"resetHeaderSet",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference) ResetHosts() {
	_jsii_.InvokeVoid(
		n,
		"resetHosts",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference) ResetMethods() {
	_jsii_.InvokeVoid(
		n,
		"resetMethods",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference) ResetPaths() {
	_jsii_.InvokeVoid(
		n,
		"resetPaths",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (n *jsiiProxy_NetworkSecurityAuthzPolicyHttpRulesToNotOperationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

