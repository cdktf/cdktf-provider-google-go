// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computenetworkfirewallpolicywithrules

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/computenetworkfirewallpolicywithrules/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeNetworkFirewallPolicyWithRulesRuleMatchSrcSecureTagList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) ComputeNetworkFirewallPolicyWithRulesRuleMatchSrcSecureTagOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeNetworkFirewallPolicyWithRulesRuleMatchSrcSecureTagList
type jsiiProxy_ComputeNetworkFirewallPolicyWithRulesRuleMatchSrcSecureTagList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesRuleMatchSrcSecureTagList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesRuleMatchSrcSecureTagList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesRuleMatchSrcSecureTagList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesRuleMatchSrcSecureTagList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesRuleMatchSrcSecureTagList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesRuleMatchSrcSecureTagList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewComputeNetworkFirewallPolicyWithRulesRuleMatchSrcSecureTagList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) ComputeNetworkFirewallPolicyWithRulesRuleMatchSrcSecureTagList {
	_init_.Initialize()

	if err := validateNewComputeNetworkFirewallPolicyWithRulesRuleMatchSrcSecureTagListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeNetworkFirewallPolicyWithRulesRuleMatchSrcSecureTagList{}

	_jsii_.Create(
		"@cdktf/provider-google.computeNetworkFirewallPolicyWithRules.ComputeNetworkFirewallPolicyWithRulesRuleMatchSrcSecureTagList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewComputeNetworkFirewallPolicyWithRulesRuleMatchSrcSecureTagList_Override(c ComputeNetworkFirewallPolicyWithRulesRuleMatchSrcSecureTagList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeNetworkFirewallPolicyWithRules.ComputeNetworkFirewallPolicyWithRulesRuleMatchSrcSecureTagList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		c,
	)
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesRuleMatchSrcSecureTagList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesRuleMatchSrcSecureTagList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesRuleMatchSrcSecureTagList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesRuleMatchSrcSecureTagList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (c *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesRuleMatchSrcSecureTagList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := c.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		c,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesRuleMatchSrcSecureTagList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesRuleMatchSrcSecureTagList) Get(index *float64) ComputeNetworkFirewallPolicyWithRulesRuleMatchSrcSecureTagOutputReference {
	if err := c.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns ComputeNetworkFirewallPolicyWithRulesRuleMatchSrcSecureTagOutputReference

	_jsii_.Invoke(
		c,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesRuleMatchSrcSecureTagList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeNetworkFirewallPolicyWithRulesRuleMatchSrcSecureTagList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

