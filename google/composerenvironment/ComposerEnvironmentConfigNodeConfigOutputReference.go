// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package composerenvironment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/composerenvironment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComposerEnvironmentConfigNodeConfigOutputReference interface {
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
	ComposerInternalIpv4CidrBlock() *string
	SetComposerInternalIpv4CidrBlock(val *string)
	ComposerInternalIpv4CidrBlockInput() *string
	ComposerNetworkAttachment() *string
	SetComposerNetworkAttachment(val *string)
	ComposerNetworkAttachmentInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DiskSizeGb() *float64
	SetDiskSizeGb(val *float64)
	DiskSizeGbInput() *float64
	EnableIpMasqAgent() interface{}
	SetEnableIpMasqAgent(val interface{})
	EnableIpMasqAgentInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *ComposerEnvironmentConfigNodeConfig
	SetInternalValue(val *ComposerEnvironmentConfigNodeConfig)
	IpAllocationPolicy() ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference
	IpAllocationPolicyInput() *ComposerEnvironmentConfigNodeConfigIpAllocationPolicy
	MachineType() *string
	SetMachineType(val *string)
	MachineTypeInput() *string
	Network() *string
	SetNetwork(val *string)
	NetworkInput() *string
	OauthScopes() *[]*string
	SetOauthScopes(val *[]*string)
	OauthScopesInput() *[]*string
	ServiceAccount() *string
	SetServiceAccount(val *string)
	ServiceAccountInput() *string
	Subnetwork() *string
	SetSubnetwork(val *string)
	SubnetworkInput() *string
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Zone() *string
	SetZone(val *string)
	ZoneInput() *string
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
	PutIpAllocationPolicy(value *ComposerEnvironmentConfigNodeConfigIpAllocationPolicy)
	ResetComposerInternalIpv4CidrBlock()
	ResetComposerNetworkAttachment()
	ResetDiskSizeGb()
	ResetEnableIpMasqAgent()
	ResetIpAllocationPolicy()
	ResetMachineType()
	ResetNetwork()
	ResetOauthScopes()
	ResetServiceAccount()
	ResetSubnetwork()
	ResetTags()
	ResetZone()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ComposerEnvironmentConfigNodeConfigOutputReference
type jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) ComposerInternalIpv4CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"composerInternalIpv4CidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) ComposerInternalIpv4CidrBlockInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"composerInternalIpv4CidrBlockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) ComposerNetworkAttachment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"composerNetworkAttachment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) ComposerNetworkAttachmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"composerNetworkAttachmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) DiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) DiskSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) EnableIpMasqAgent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableIpMasqAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) EnableIpMasqAgentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableIpMasqAgentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) InternalValue() *ComposerEnvironmentConfigNodeConfig {
	var returns *ComposerEnvironmentConfigNodeConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) IpAllocationPolicy() ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference {
	var returns ComposerEnvironmentConfigNodeConfigIpAllocationPolicyOutputReference
	_jsii_.Get(
		j,
		"ipAllocationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) IpAllocationPolicyInput() *ComposerEnvironmentConfigNodeConfigIpAllocationPolicy {
	var returns *ComposerEnvironmentConfigNodeConfigIpAllocationPolicy
	_jsii_.Get(
		j,
		"ipAllocationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) MachineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) MachineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) OauthScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"oauthScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) OauthScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"oauthScopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) ServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) Subnetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) SubnetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) ZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneInput",
		&returns,
	)
	return returns
}


func NewComposerEnvironmentConfigNodeConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ComposerEnvironmentConfigNodeConfigOutputReference {
	_init_.Initialize()

	if err := validateNewComposerEnvironmentConfigNodeConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.composerEnvironment.ComposerEnvironmentConfigNodeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewComposerEnvironmentConfigNodeConfigOutputReference_Override(c ComposerEnvironmentConfigNodeConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.composerEnvironment.ComposerEnvironmentConfigNodeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference)SetComposerInternalIpv4CidrBlock(val *string) {
	if err := j.validateSetComposerInternalIpv4CidrBlockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"composerInternalIpv4CidrBlock",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference)SetComposerNetworkAttachment(val *string) {
	if err := j.validateSetComposerNetworkAttachmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"composerNetworkAttachment",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference)SetDiskSizeGb(val *float64) {
	if err := j.validateSetDiskSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskSizeGb",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference)SetEnableIpMasqAgent(val interface{}) {
	if err := j.validateSetEnableIpMasqAgentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableIpMasqAgent",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference)SetInternalValue(val *ComposerEnvironmentConfigNodeConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference)SetMachineType(val *string) {
	if err := j.validateSetMachineTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"machineType",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference)SetOauthScopes(val *[]*string) {
	if err := j.validateSetOauthScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthScopes",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference)SetServiceAccount(val *string) {
	if err := j.validateSetServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccount",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference)SetSubnetwork(val *string) {
	if err := j.validateSetSubnetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetwork",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference)SetZone(val *string) {
	if err := j.validateSetZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zone",
		val,
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) PutIpAllocationPolicy(value *ComposerEnvironmentConfigNodeConfigIpAllocationPolicy) {
	if err := c.validatePutIpAllocationPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putIpAllocationPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) ResetComposerInternalIpv4CidrBlock() {
	_jsii_.InvokeVoid(
		c,
		"resetComposerInternalIpv4CidrBlock",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) ResetComposerNetworkAttachment() {
	_jsii_.InvokeVoid(
		c,
		"resetComposerNetworkAttachment",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) ResetDiskSizeGb() {
	_jsii_.InvokeVoid(
		c,
		"resetDiskSizeGb",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) ResetEnableIpMasqAgent() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableIpMasqAgent",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) ResetIpAllocationPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetIpAllocationPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) ResetMachineType() {
	_jsii_.InvokeVoid(
		c,
		"resetMachineType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) ResetNetwork() {
	_jsii_.InvokeVoid(
		c,
		"resetNetwork",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) ResetOauthScopes() {
	_jsii_.InvokeVoid(
		c,
		"resetOauthScopes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		c,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) ResetSubnetwork() {
	_jsii_.InvokeVoid(
		c,
		"resetSubnetwork",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) ResetZone() {
	_jsii_.InvokeVoid(
		c,
		"resetZone",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ComposerEnvironmentConfigNodeConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

