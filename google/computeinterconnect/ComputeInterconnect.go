// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeinterconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v13/computeinterconnect/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/compute_interconnect google_compute_interconnect}.
type ComputeInterconnect interface {
	cdktf.TerraformResource
	AdminEnabled() interface{}
	SetAdminEnabled(val interface{})
	AdminEnabledInput() interface{}
	AvailableFeatures() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CircuitInfos() ComputeInterconnectCircuitInfosList
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
	CustomerName() *string
	SetCustomerName(val *string)
	CustomerNameInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EffectiveLabels() cdktf.StringMap
	ExpectedOutages() ComputeInterconnectExpectedOutagesList
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GoogleIpAddress() *string
	GoogleReferenceId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InterconnectAttachments() *[]*string
	InterconnectType() *string
	SetInterconnectType(val *string)
	InterconnectTypeInput() *string
	LabelFingerprint() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LinkType() *string
	SetLinkType(val *string)
	LinkTypeInput() *string
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Macsec() ComputeInterconnectMacsecOutputReference
	MacsecEnabled() interface{}
	SetMacsecEnabled(val interface{})
	MacsecEnabledInput() interface{}
	MacsecInput() *ComputeInterconnectMacsec
	Name() *string
	SetName(val *string)
	NameInput() *string
	NocContactEmail() *string
	SetNocContactEmail(val *string)
	NocContactEmailInput() *string
	// The tree node.
	Node() constructs.Node
	OperationalStatus() *string
	PeerIpAddress() *string
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	ProvisionedLinkCount() *float64
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	RemoteLocation() *string
	SetRemoteLocation(val *string)
	RemoteLocationInput() *string
	RequestedFeatures() *[]*string
	SetRequestedFeatures(val *[]*string)
	RequestedFeaturesInput() *[]*string
	RequestedLinkCount() *float64
	SetRequestedLinkCount(val *float64)
	RequestedLinkCountInput() *float64
	SatisfiesPzs() cdktf.IResolvable
	State() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ComputeInterconnectTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutMacsec(value *ComputeInterconnectMacsec)
	PutTimeouts(value *ComputeInterconnectTimeouts)
	ResetAdminEnabled()
	ResetDescription()
	ResetId()
	ResetLabels()
	ResetMacsec()
	ResetMacsecEnabled()
	ResetNocContactEmail()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetRemoteLocation()
	ResetRequestedFeatures()
	ResetTimeouts()
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

// The jsii proxy struct for ComputeInterconnect
type jsiiProxy_ComputeInterconnect struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ComputeInterconnect) AdminEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adminEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) AdminEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adminEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) AvailableFeatures() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availableFeatures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) CircuitInfos() ComputeInterconnectCircuitInfosList {
	var returns ComputeInterconnectCircuitInfosList
	_jsii_.Get(
		j,
		"circuitInfos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) CreationTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) CustomerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) CustomerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) ExpectedOutages() ComputeInterconnectExpectedOutagesList {
	var returns ComputeInterconnectExpectedOutagesList
	_jsii_.Get(
		j,
		"expectedOutages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) GoogleIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) GoogleReferenceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleReferenceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) InterconnectAttachments() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"interconnectAttachments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) InterconnectType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interconnectType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) InterconnectTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interconnectTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) LabelFingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelFingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) LinkType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linkType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) LinkTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linkTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) Macsec() ComputeInterconnectMacsecOutputReference {
	var returns ComputeInterconnectMacsecOutputReference
	_jsii_.Get(
		j,
		"macsec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) MacsecEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"macsecEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) MacsecEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"macsecEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) MacsecInput() *ComputeInterconnectMacsec {
	var returns *ComputeInterconnectMacsec
	_jsii_.Get(
		j,
		"macsecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) NocContactEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nocContactEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) NocContactEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nocContactEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) OperationalStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationalStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) PeerIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peerIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) ProvisionedLinkCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedLinkCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) RemoteLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) RemoteLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) RequestedFeatures() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requestedFeatures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) RequestedFeaturesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"requestedFeaturesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) RequestedLinkCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestedLinkCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) RequestedLinkCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestedLinkCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) SatisfiesPzs() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"satisfiesPzs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) Timeouts() ComputeInterconnectTimeoutsOutputReference {
	var returns ComputeInterconnectTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInterconnect) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/compute_interconnect google_compute_interconnect} Resource.
func NewComputeInterconnect(scope constructs.Construct, id *string, config *ComputeInterconnectConfig) ComputeInterconnect {
	_init_.Initialize()

	if err := validateNewComputeInterconnectParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeInterconnect{}

	_jsii_.Create(
		"@cdktf/provider-google.computeInterconnect.ComputeInterconnect",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/compute_interconnect google_compute_interconnect} Resource.
func NewComputeInterconnect_Override(c ComputeInterconnect, scope constructs.Construct, id *string, config *ComputeInterconnectConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeInterconnect.ComputeInterconnect",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ComputeInterconnect)SetAdminEnabled(val interface{}) {
	if err := j.validateSetAdminEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminEnabled",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnect)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnect)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnect)SetCustomerName(val *string) {
	if err := j.validateSetCustomerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerName",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnect)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnect)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnect)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnect)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnect)SetInterconnectType(val *string) {
	if err := j.validateSetInterconnectTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interconnectType",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnect)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnect)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnect)SetLinkType(val *string) {
	if err := j.validateSetLinkTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"linkType",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnect)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnect)SetMacsecEnabled(val interface{}) {
	if err := j.validateSetMacsecEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"macsecEnabled",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnect)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnect)SetNocContactEmail(val *string) {
	if err := j.validateSetNocContactEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nocContactEmail",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnect)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnect)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnect)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnect)SetRemoteLocation(val *string) {
	if err := j.validateSetRemoteLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteLocation",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnect)SetRequestedFeatures(val *[]*string) {
	if err := j.validateSetRequestedFeaturesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestedFeatures",
		val,
	)
}

func (j *jsiiProxy_ComputeInterconnect)SetRequestedLinkCount(val *float64) {
	if err := j.validateSetRequestedLinkCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestedLinkCount",
		val,
	)
}

// Generates CDKTF code for importing a ComputeInterconnect resource upon running "cdktf plan <stack-name>".
func ComputeInterconnect_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateComputeInterconnect_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeInterconnect.ComputeInterconnect",
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
func ComputeInterconnect_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeInterconnect_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeInterconnect.ComputeInterconnect",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeInterconnect_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeInterconnect_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeInterconnect.ComputeInterconnect",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeInterconnect_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeInterconnect_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeInterconnect.ComputeInterconnect",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ComputeInterconnect_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.computeInterconnect.ComputeInterconnect",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ComputeInterconnect) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_ComputeInterconnect) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ComputeInterconnect) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeInterconnect) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeInterconnect) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeInterconnect) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeInterconnect) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeInterconnect) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeInterconnect) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeInterconnect) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeInterconnect) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeInterconnect) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInterconnect) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_ComputeInterconnect) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeInterconnect) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ComputeInterconnect) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_ComputeInterconnect) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ComputeInterconnect) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ComputeInterconnect) PutMacsec(value *ComputeInterconnectMacsec) {
	if err := c.validatePutMacsecParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMacsec",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInterconnect) PutTimeouts(value *ComputeInterconnectTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInterconnect) ResetAdminEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetAdminEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInterconnect) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInterconnect) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInterconnect) ResetLabels() {
	_jsii_.InvokeVoid(
		c,
		"resetLabels",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInterconnect) ResetMacsec() {
	_jsii_.InvokeVoid(
		c,
		"resetMacsec",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInterconnect) ResetMacsecEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetMacsecEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInterconnect) ResetNocContactEmail() {
	_jsii_.InvokeVoid(
		c,
		"resetNocContactEmail",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInterconnect) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInterconnect) ResetProject() {
	_jsii_.InvokeVoid(
		c,
		"resetProject",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInterconnect) ResetRemoteLocation() {
	_jsii_.InvokeVoid(
		c,
		"resetRemoteLocation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInterconnect) ResetRequestedFeatures() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestedFeatures",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInterconnect) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInterconnect) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInterconnect) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInterconnect) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInterconnect) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInterconnect) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInterconnect) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

