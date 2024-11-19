// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeinterconnectattachment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v14/computeinterconnectattachment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/compute_interconnect_attachment google_compute_interconnect_attachment}.
type ComputeInterconnectAttachment interface {
	cdktf.TerraformResource
	AdminEnabled() interface{}
	SetAdminEnabled(val interface{})
	AdminEnabledInput() interface{}
	Bandwidth() *string
	SetBandwidth(val *string)
	BandwidthInput() *string
	CandidateSubnets() *[]*string
	SetCandidateSubnets(val *[]*string)
	CandidateSubnetsInput() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CloudRouterIpAddress() *string
	CloudRouterIpv6Address() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreationTimestamp() *string
	CustomerRouterIpAddress() *string
	CustomerRouterIpv6Address() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EdgeAvailabilityDomain() *string
	SetEdgeAvailabilityDomain(val *string)
	EdgeAvailabilityDomainInput() *string
	Encryption() *string
	SetEncryption(val *string)
	EncryptionInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GoogleReferenceId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Interconnect() *string
	SetInterconnect(val *string)
	InterconnectInput() *string
	IpsecInternalAddresses() *[]*string
	SetIpsecInternalAddresses(val *[]*string)
	IpsecInternalAddressesInput() *[]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Mtu() *string
	SetMtu(val *string)
	MtuInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PairingKey() *string
	PartnerAsn() *string
	PrivateInterconnectInfo() ComputeInterconnectAttachmentPrivateInterconnectInfoList
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	Router() *string
	SetRouter(val *string)
	RouterInput() *string
	SelfLink() *string
	StackType() *string
	SetStackType(val *string)
	StackTypeInput() *string
	State() *string
	SubnetLength() *float64
	SetSubnetLength(val *float64)
	SubnetLengthInput() *float64
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ComputeInterconnectAttachmentTimeoutsOutputReference
	TimeoutsInput() interface{}
	Type() *string
	SetType(val *string)
	TypeInput() *string
	VlanTag8021Q() *float64
	SetVlanTag8021Q(val *float64)
	VlanTag8021QInput() *float64
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutTimeouts(value *ComputeInterconnectAttachmentTimeouts)
	ResetAdminEnabled()
	ResetBandwidth()
	ResetCandidateSubnets()
	ResetDescription()
	ResetEdgeAvailabilityDomain()
	ResetEncryption()
	ResetId()
	ResetInterconnect()
	ResetIpsecInternalAddresses()
	ResetMtu()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetRegion()
	ResetStackType()
	ResetSubnetLength()
	ResetTimeouts()
	ResetType()
	ResetVlanTag8021Q()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ComputeInterconnectAttachment
type jsiiProxy_ComputeInterconnectAttachment struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ComputeInterconnectAttachment) AdminEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adminEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) AdminEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adminEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) Bandwidth() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bandwidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) BandwidthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bandwidthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) CandidateSubnets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"candidateSubnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) CandidateSubnetsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"candidateSubnetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) CloudRouterIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudRouterIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) CloudRouterIpv6Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudRouterIpv6Address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) CreationTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) CustomerRouterIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerRouterIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) CustomerRouterIpv6Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerRouterIpv6Address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) EdgeAvailabilityDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeAvailabilityDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) EdgeAvailabilityDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeAvailabilityDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) Encryption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) EncryptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) GoogleReferenceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleReferenceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) Interconnect() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interconnect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) InterconnectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interconnectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) IpsecInternalAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipsecInternalAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) IpsecInternalAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipsecInternalAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) Mtu() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mtu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) MtuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mtuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) PairingKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pairingKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) PartnerAsn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partnerAsn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) PrivateInterconnectInfo() ComputeInterconnectAttachmentPrivateInterconnectInfoList {
	var returns ComputeInterconnectAttachmentPrivateInterconnectInfoList
	_jsii_.Get(
		j,
		"privateInterconnectInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) Router() *string {
	var returns *string
	_jsii_.Get(
		j,
		"router",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) RouterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) StackType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) StackTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) SubnetLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"subnetLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) SubnetLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"subnetLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) Timeouts() ComputeInterconnectAttachmentTimeoutsOutputReference {
	var returns ComputeInterconnectAttachmentTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) VlanTag8021Q() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vlanTag8021Q",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnectAttachment) VlanTag8021QInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vlanTag8021QInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/compute_interconnect_attachment google_compute_interconnect_attachment} Resource.
func NewComputeInterconnectAttachment(scope constructs.Construct, id *string, config *ComputeInterconnectAttachmentConfig) ComputeInterconnectAttachment {
	_init_.Initialize()

	if err := validateNewComputeInterconnectAttachmentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeInterconnectAttachment{}

	_jsii_.Create(
		"@cdktf/provider-google.computeInterconnectAttachment.ComputeInterconnectAttachment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/compute_interconnect_attachment google_compute_interconnect_attachment} Resource.
func NewComputeInterconnectAttachment_Override(c ComputeInterconnectAttachment, scope constructs.Construct, id *string, config *ComputeInterconnectAttachmentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeInterconnectAttachment.ComputeInterconnectAttachment",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachment)SetAdminEnabled(val interface{}) {
	if err := j.validateSetAdminEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminEnabled",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachment)SetBandwidth(val *string) {
	if err := j.validateSetBandwidthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bandwidth",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachment)SetCandidateSubnets(val *[]*string) {
	if err := j.validateSetCandidateSubnetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"candidateSubnets",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachment)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachment)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachment)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachment)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachment)SetEdgeAvailabilityDomain(val *string) {
	if err := j.validateSetEdgeAvailabilityDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edgeAvailabilityDomain",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachment)SetEncryption(val *string) {
	if err := j.validateSetEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryption",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachment)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachment)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachment)SetInterconnect(val *string) {
	if err := j.validateSetInterconnectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interconnect",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachment)SetIpsecInternalAddresses(val *[]*string) {
	if err := j.validateSetIpsecInternalAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipsecInternalAddresses",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachment)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachment)SetMtu(val *string) {
	if err := j.validateSetMtuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mtu",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachment)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachment)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachment)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachment)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachment)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachment)SetRouter(val *string) {
	if err := j.validateSetRouterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"router",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachment)SetStackType(val *string) {
	if err := j.validateSetStackTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stackType",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachment)SetSubnetLength(val *float64) {
	if err := j.validateSetSubnetLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetLength",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachment)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnectAttachment)SetVlanTag8021Q(val *float64) {
	if err := j.validateSetVlanTag8021QParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vlanTag8021Q",
		val,
	)
}

// Generates CDKTF code for importing a ComputeInterconnectAttachment resource upon running "cdktf plan <stack-name>".
func ComputeInterconnectAttachment_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateComputeInterconnectAttachment_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeInterconnectAttachment.ComputeInterconnectAttachment",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func ComputeInterconnectAttachment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeInterconnectAttachment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeInterconnectAttachment.ComputeInterconnectAttachment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeInterconnectAttachment_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeInterconnectAttachment_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeInterconnectAttachment.ComputeInterconnectAttachment",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeInterconnectAttachment_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeInterconnectAttachment_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeInterconnectAttachment.ComputeInterconnectAttachment",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ComputeInterconnectAttachment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.computeInterconnectAttachment.ComputeInterconnectAttachment",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ComputeInterconnectAttachment) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_ComputeInterconnectAttachment) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ComputeInterconnectAttachment) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeInterconnectAttachment) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeInterconnectAttachment) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeInterconnectAttachment) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeInterconnectAttachment) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeInterconnectAttachment) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeInterconnectAttachment) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeInterconnectAttachment) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeInterconnectAttachment) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeInterconnectAttachment) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInterconnectAttachment) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_ComputeInterconnectAttachment) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInterconnectAttachment) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ComputeInterconnectAttachment) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_ComputeInterconnectAttachment) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ComputeInterconnectAttachment) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ComputeInterconnectAttachment) PutTimeouts(value *ComputeInterconnectAttachmentTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInterconnectAttachment) ResetAdminEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetAdminEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInterconnectAttachment) ResetBandwidth() {
	_jsii_.InvokeVoid(
		c,
		"resetBandwidth",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInterconnectAttachment) ResetCandidateSubnets() {
	_jsii_.InvokeVoid(
		c,
		"resetCandidateSubnets",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInterconnectAttachment) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInterconnectAttachment) ResetEdgeAvailabilityDomain() {
	_jsii_.InvokeVoid(
		c,
		"resetEdgeAvailabilityDomain",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInterconnectAttachment) ResetEncryption() {
	_jsii_.InvokeVoid(
		c,
		"resetEncryption",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInterconnectAttachment) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInterconnectAttachment) ResetInterconnect() {
	_jsii_.InvokeVoid(
		c,
		"resetInterconnect",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInterconnectAttachment) ResetIpsecInternalAddresses() {
	_jsii_.InvokeVoid(
		c,
		"resetIpsecInternalAddresses",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInterconnectAttachment) ResetMtu() {
	_jsii_.InvokeVoid(
		c,
		"resetMtu",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInterconnectAttachment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInterconnectAttachment) ResetProject() {
	_jsii_.InvokeVoid(
		c,
		"resetProject",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInterconnectAttachment) ResetRegion() {
	_jsii_.InvokeVoid(
		c,
		"resetRegion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInterconnectAttachment) ResetStackType() {
	_jsii_.InvokeVoid(
		c,
		"resetStackType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInterconnectAttachment) ResetSubnetLength() {
	_jsii_.InvokeVoid(
		c,
		"resetSubnetLength",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInterconnectAttachment) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInterconnectAttachment) ResetType() {
	_jsii_.InvokeVoid(
		c,
		"resetType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInterconnectAttachment) ResetVlanTag8021Q() {
	_jsii_.InvokeVoid(
		c,
		"resetVlanTag8021Q",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInterconnectAttachment) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInterconnectAttachment) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInterconnectAttachment) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInterconnectAttachment) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInterconnectAttachment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInterconnectAttachment) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

