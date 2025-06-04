// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computenetworkfirewallpolicyrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/computenetworkfirewallpolicyrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeNetworkFirewallPolicyRuleMatchOutputReference interface {
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
	SetDestAddressGroups(val *[]*string)
	DestAddressGroupsInput() *[]*string
	DestFqdns() *[]*string
	SetDestFqdns(val *[]*string)
	DestFqdnsInput() *[]*string
	DestIpRanges() *[]*string
	SetDestIpRanges(val *[]*string)
	DestIpRangesInput() *[]*string
	DestRegionCodes() *[]*string
	SetDestRegionCodes(val *[]*string)
	DestRegionCodesInput() *[]*string
	DestThreatIntelligences() *[]*string
	SetDestThreatIntelligences(val *[]*string)
	DestThreatIntelligencesInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *ComputeNetworkFirewallPolicyRuleMatch
	SetInternalValue(val *ComputeNetworkFirewallPolicyRuleMatch)
	Layer4Configs() ComputeNetworkFirewallPolicyRuleMatchLayer4ConfigsList
	Layer4ConfigsInput() interface{}
	SrcAddressGroups() *[]*string
	SetSrcAddressGroups(val *[]*string)
	SrcAddressGroupsInput() *[]*string
	SrcFqdns() *[]*string
	SetSrcFqdns(val *[]*string)
	SrcFqdnsInput() *[]*string
	SrcIpRanges() *[]*string
	SetSrcIpRanges(val *[]*string)
	SrcIpRangesInput() *[]*string
	SrcRegionCodes() *[]*string
	SetSrcRegionCodes(val *[]*string)
	SrcRegionCodesInput() *[]*string
	SrcSecureTags() ComputeNetworkFirewallPolicyRuleMatchSrcSecureTagsList
	SrcSecureTagsInput() interface{}
	SrcThreatIntelligences() *[]*string
	SetSrcThreatIntelligences(val *[]*string)
	SrcThreatIntelligencesInput() *[]*string
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
	PutLayer4Configs(value interface{})
	PutSrcSecureTags(value interface{})
	ResetDestAddressGroups()
	ResetDestFqdns()
	ResetDestIpRanges()
	ResetDestRegionCodes()
	ResetDestThreatIntelligences()
	ResetSrcAddressGroups()
	ResetSrcFqdns()
	ResetSrcIpRanges()
	ResetSrcRegionCodes()
	ResetSrcSecureTags()
	ResetSrcThreatIntelligences()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComputeNetworkFirewallPolicyRuleMatchOutputReference
type jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) DestAddressGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destAddressGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) DestAddressGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destAddressGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) DestFqdns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destFqdns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) DestFqdnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destFqdnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) DestIpRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) DestIpRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) DestRegionCodes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destRegionCodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) DestRegionCodesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destRegionCodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) DestThreatIntelligences() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destThreatIntelligences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) DestThreatIntelligencesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destThreatIntelligencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) InternalValue() *ComputeNetworkFirewallPolicyRuleMatch {
	var returns *ComputeNetworkFirewallPolicyRuleMatch
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) Layer4Configs() ComputeNetworkFirewallPolicyRuleMatchLayer4ConfigsList {
	var returns ComputeNetworkFirewallPolicyRuleMatchLayer4ConfigsList
	_jsii_.Get(
		j,
		"layer4Configs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) Layer4ConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"layer4ConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) SrcAddressGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcAddressGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) SrcAddressGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcAddressGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) SrcFqdns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcFqdns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) SrcFqdnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcFqdnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) SrcIpRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) SrcIpRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) SrcRegionCodes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcRegionCodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) SrcRegionCodesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcRegionCodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) SrcSecureTags() ComputeNetworkFirewallPolicyRuleMatchSrcSecureTagsList {
	var returns ComputeNetworkFirewallPolicyRuleMatchSrcSecureTagsList
	_jsii_.Get(
		j,
		"srcSecureTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) SrcSecureTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"srcSecureTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) SrcThreatIntelligences() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcThreatIntelligences",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) SrcThreatIntelligencesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"srcThreatIntelligencesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewComputeNetworkFirewallPolicyRuleMatchOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComputeNetworkFirewallPolicyRuleMatchOutputReference {
	_init_.Initialize()

	if err := validateNewComputeNetworkFirewallPolicyRuleMatchOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.computeNetworkFirewallPolicyRule.ComputeNetworkFirewallPolicyRuleMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComputeNetworkFirewallPolicyRuleMatchOutputReference_Override(c ComputeNetworkFirewallPolicyRuleMatchOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeNetworkFirewallPolicyRule.ComputeNetworkFirewallPolicyRuleMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference)SetDestAddressGroups(val *[]*string) {
	if err := j.validateSetDestAddressGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destAddressGroups",
		val,
	)
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference)SetDestFqdns(val *[]*string) {
	if err := j.validateSetDestFqdnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destFqdns",
		val,
	)
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference)SetDestIpRanges(val *[]*string) {
	if err := j.validateSetDestIpRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destIpRanges",
		val,
	)
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference)SetDestRegionCodes(val *[]*string) {
	if err := j.validateSetDestRegionCodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destRegionCodes",
		val,
	)
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference)SetDestThreatIntelligences(val *[]*string) {
	if err := j.validateSetDestThreatIntelligencesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destThreatIntelligences",
		val,
	)
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference)SetInternalValue(val *ComputeNetworkFirewallPolicyRuleMatch) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference)SetSrcAddressGroups(val *[]*string) {
	if err := j.validateSetSrcAddressGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"srcAddressGroups",
		val,
	)
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference)SetSrcFqdns(val *[]*string) {
	if err := j.validateSetSrcFqdnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"srcFqdns",
		val,
	)
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference)SetSrcIpRanges(val *[]*string) {
	if err := j.validateSetSrcIpRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"srcIpRanges",
		val,
	)
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference)SetSrcRegionCodes(val *[]*string) {
	if err := j.validateSetSrcRegionCodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"srcRegionCodes",
		val,
	)
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference)SetSrcThreatIntelligences(val *[]*string) {
	if err := j.validateSetSrcThreatIntelligencesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"srcThreatIntelligences",
		val,
	)
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) PutLayer4Configs(value interface{}) {
	if err := c.validatePutLayer4ConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLayer4Configs",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) PutSrcSecureTags(value interface{}) {
	if err := c.validatePutSrcSecureTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSrcSecureTags",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) ResetDestAddressGroups() {
	_jsii_.InvokeVoid(
		c,
		"resetDestAddressGroups",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) ResetDestFqdns() {
	_jsii_.InvokeVoid(
		c,
		"resetDestFqdns",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) ResetDestIpRanges() {
	_jsii_.InvokeVoid(
		c,
		"resetDestIpRanges",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) ResetDestRegionCodes() {
	_jsii_.InvokeVoid(
		c,
		"resetDestRegionCodes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) ResetDestThreatIntelligences() {
	_jsii_.InvokeVoid(
		c,
		"resetDestThreatIntelligences",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) ResetSrcAddressGroups() {
	_jsii_.InvokeVoid(
		c,
		"resetSrcAddressGroups",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) ResetSrcFqdns() {
	_jsii_.InvokeVoid(
		c,
		"resetSrcFqdns",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) ResetSrcIpRanges() {
	_jsii_.InvokeVoid(
		c,
		"resetSrcIpRanges",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) ResetSrcRegionCodes() {
	_jsii_.InvokeVoid(
		c,
		"resetSrcRegionCodes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) ResetSrcSecureTags() {
	_jsii_.InvokeVoid(
		c,
		"resetSrcSecureTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) ResetSrcThreatIntelligences() {
	_jsii_.InvokeVoid(
		c,
		"resetSrcThreatIntelligences",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComputeNetworkFirewallPolicyRuleMatchOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

