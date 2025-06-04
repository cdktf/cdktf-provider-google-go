// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregionnetworkfirewallpolicywithrules

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/computeregionnetworkfirewallpolicywithrules/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference interface {
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
	DestAddressGroups() *[]*string
	DestFqdns() *[]*string
	DestIpRanges() *[]*string
	DestRegionCodes() *[]*string
	DestThreatIntelligences() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatch
	SetInternalValue(val *ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatch)
	Layer4Config() ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchLayer4ConfigList
	SrcAddressGroups() *[]*string
	SrcFqdns() *[]*string
	SrcIpRanges() *[]*string
	SrcRegionCodes() *[]*string
	SrcSecureTag() ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchSrcSecureTagList
	SrcThreatIntelligences() *[]*string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference
type jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) DestAddressGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destAddressGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) DestFqdns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destFqdns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) DestIpRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) DestRegionCodes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destRegionCodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) DestThreatIntelligences() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destThreatIntelligences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) InternalValue() *ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatch {
	var returns *ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatch
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) Layer4Config() ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchLayer4ConfigList {
	var returns ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchLayer4ConfigList
	_jsii_.Get(
		j,
		"layer4Config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) SrcAddressGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcAddressGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) SrcFqdns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcFqdns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) SrcIpRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) SrcRegionCodes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcRegionCodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) SrcSecureTag() ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchSrcSecureTagList {
	var returns ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchSrcSecureTagList
	_jsii_.Get(
		j,
		"srcSecureTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) SrcThreatIntelligences() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcThreatIntelligences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference {
	_init_.Initialize()

	if err := validateNewComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionNetworkFirewallPolicyWithRules.ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference_Override(c ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionNetworkFirewallPolicyWithRules.ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference)SetInternalValue(val *ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatch) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeRegionNetworkFirewallPolicyWithRulesPredefinedRulesMatchOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

