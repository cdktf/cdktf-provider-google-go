// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkconnectivityspoke

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v16/networkconnectivityspoke/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/network_connectivity_spoke google_network_connectivity_spoke}.
type NetworkConnectivitySpoke interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	CreateTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EffectiveLabels() cdktf.StringMap
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Group() *string
	SetGroup(val *string)
	GroupInput() *string
	Hub() *string
	SetHub(val *string)
	HubInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LinkedInterconnectAttachments() NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference
	LinkedInterconnectAttachmentsInput() *NetworkConnectivitySpokeLinkedInterconnectAttachments
	LinkedProducerVpcNetwork() NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference
	LinkedProducerVpcNetworkInput() *NetworkConnectivitySpokeLinkedProducerVpcNetwork
	LinkedRouterApplianceInstances() NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference
	LinkedRouterApplianceInstancesInput() *NetworkConnectivitySpokeLinkedRouterApplianceInstances
	LinkedVpcNetwork() NetworkConnectivitySpokeLinkedVpcNetworkOutputReference
	LinkedVpcNetworkInput() *NetworkConnectivitySpokeLinkedVpcNetwork
	LinkedVpnTunnels() NetworkConnectivitySpokeLinkedVpnTunnelsOutputReference
	LinkedVpnTunnelsInput() *NetworkConnectivitySpokeLinkedVpnTunnels
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
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
	Reasons() NetworkConnectivitySpokeReasonsList
	State() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() NetworkConnectivitySpokeTimeoutsOutputReference
	TimeoutsInput() interface{}
	UniqueId() *string
	UpdateTime() *string
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
	PutLinkedInterconnectAttachments(value *NetworkConnectivitySpokeLinkedInterconnectAttachments)
	PutLinkedProducerVpcNetwork(value *NetworkConnectivitySpokeLinkedProducerVpcNetwork)
	PutLinkedRouterApplianceInstances(value *NetworkConnectivitySpokeLinkedRouterApplianceInstances)
	PutLinkedVpcNetwork(value *NetworkConnectivitySpokeLinkedVpcNetwork)
	PutLinkedVpnTunnels(value *NetworkConnectivitySpokeLinkedVpnTunnels)
	PutTimeouts(value *NetworkConnectivitySpokeTimeouts)
	ResetDescription()
	ResetGroup()
	ResetId()
	ResetLabels()
	ResetLinkedInterconnectAttachments()
	ResetLinkedProducerVpcNetwork()
	ResetLinkedRouterApplianceInstances()
	ResetLinkedVpcNetwork()
	ResetLinkedVpnTunnels()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
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

// The jsii proxy struct for NetworkConnectivitySpoke
type jsiiProxy_NetworkConnectivitySpoke struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_NetworkConnectivitySpoke) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) Group() *string {
	var returns *string
	_jsii_.Get(
		j,
		"group",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) GroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) Hub() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hub",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) HubInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hubInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) LinkedInterconnectAttachments() NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference {
	var returns NetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference
	_jsii_.Get(
		j,
		"linkedInterconnectAttachments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) LinkedInterconnectAttachmentsInput() *NetworkConnectivitySpokeLinkedInterconnectAttachments {
	var returns *NetworkConnectivitySpokeLinkedInterconnectAttachments
	_jsii_.Get(
		j,
		"linkedInterconnectAttachmentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) LinkedProducerVpcNetwork() NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference {
	var returns NetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference
	_jsii_.Get(
		j,
		"linkedProducerVpcNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) LinkedProducerVpcNetworkInput() *NetworkConnectivitySpokeLinkedProducerVpcNetwork {
	var returns *NetworkConnectivitySpokeLinkedProducerVpcNetwork
	_jsii_.Get(
		j,
		"linkedProducerVpcNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) LinkedRouterApplianceInstances() NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference {
	var returns NetworkConnectivitySpokeLinkedRouterApplianceInstancesOutputReference
	_jsii_.Get(
		j,
		"linkedRouterApplianceInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) LinkedRouterApplianceInstancesInput() *NetworkConnectivitySpokeLinkedRouterApplianceInstances {
	var returns *NetworkConnectivitySpokeLinkedRouterApplianceInstances
	_jsii_.Get(
		j,
		"linkedRouterApplianceInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) LinkedVpcNetwork() NetworkConnectivitySpokeLinkedVpcNetworkOutputReference {
	var returns NetworkConnectivitySpokeLinkedVpcNetworkOutputReference
	_jsii_.Get(
		j,
		"linkedVpcNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) LinkedVpcNetworkInput() *NetworkConnectivitySpokeLinkedVpcNetwork {
	var returns *NetworkConnectivitySpokeLinkedVpcNetwork
	_jsii_.Get(
		j,
		"linkedVpcNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) LinkedVpnTunnels() NetworkConnectivitySpokeLinkedVpnTunnelsOutputReference {
	var returns NetworkConnectivitySpokeLinkedVpnTunnelsOutputReference
	_jsii_.Get(
		j,
		"linkedVpnTunnels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) LinkedVpnTunnelsInput() *NetworkConnectivitySpokeLinkedVpnTunnels {
	var returns *NetworkConnectivitySpokeLinkedVpnTunnels
	_jsii_.Get(
		j,
		"linkedVpnTunnelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) Reasons() NetworkConnectivitySpokeReasonsList {
	var returns NetworkConnectivitySpokeReasonsList
	_jsii_.Get(
		j,
		"reasons",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) Timeouts() NetworkConnectivitySpokeTimeoutsOutputReference {
	var returns NetworkConnectivitySpokeTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) UniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivitySpoke) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/network_connectivity_spoke google_network_connectivity_spoke} Resource.
func NewNetworkConnectivitySpoke(scope constructs.Construct, id *string, config *NetworkConnectivitySpokeConfig) NetworkConnectivitySpoke {
	_init_.Initialize()

	if err := validateNewNetworkConnectivitySpokeParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkConnectivitySpoke{}

	_jsii_.Create(
		"@cdktf/provider-google.networkConnectivitySpoke.NetworkConnectivitySpoke",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/network_connectivity_spoke google_network_connectivity_spoke} Resource.
func NewNetworkConnectivitySpoke_Override(n NetworkConnectivitySpoke, scope constructs.Construct, id *string, config *NetworkConnectivitySpokeConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.networkConnectivitySpoke.NetworkConnectivitySpoke",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NetworkConnectivitySpoke)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivitySpoke)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivitySpoke)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivitySpoke)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivitySpoke)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivitySpoke)SetGroup(val *string) {
	if err := j.validateSetGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"group",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivitySpoke)SetHub(val *string) {
	if err := j.validateSetHubParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hub",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivitySpoke)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivitySpoke)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivitySpoke)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivitySpoke)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivitySpoke)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivitySpoke)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivitySpoke)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivitySpoke)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a NetworkConnectivitySpoke resource upon running "cdktf plan <stack-name>".
func NetworkConnectivitySpoke_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateNetworkConnectivitySpoke_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.networkConnectivitySpoke.NetworkConnectivitySpoke",
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
func NetworkConnectivitySpoke_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkConnectivitySpoke_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.networkConnectivitySpoke.NetworkConnectivitySpoke",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetworkConnectivitySpoke_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkConnectivitySpoke_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.networkConnectivitySpoke.NetworkConnectivitySpoke",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetworkConnectivitySpoke_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkConnectivitySpoke_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.networkConnectivitySpoke.NetworkConnectivitySpoke",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NetworkConnectivitySpoke_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.networkConnectivitySpoke.NetworkConnectivitySpoke",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NetworkConnectivitySpoke) AddMoveTarget(moveTarget *string) {
	if err := n.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (n *jsiiProxy_NetworkConnectivitySpoke) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NetworkConnectivitySpoke) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NetworkConnectivitySpoke) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkConnectivitySpoke) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NetworkConnectivitySpoke) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NetworkConnectivitySpoke) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NetworkConnectivitySpoke) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NetworkConnectivitySpoke) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NetworkConnectivitySpoke) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NetworkConnectivitySpoke) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NetworkConnectivitySpoke) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectivitySpoke) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := n.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (n *jsiiProxy_NetworkConnectivitySpoke) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectivitySpoke) MoveFromId(id *string) {
	if err := n.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveFromId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NetworkConnectivitySpoke) MoveTo(moveTarget *string, index interface{}) {
	if err := n.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (n *jsiiProxy_NetworkConnectivitySpoke) MoveToId(id *string) {
	if err := n.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveToId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NetworkConnectivitySpoke) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NetworkConnectivitySpoke) PutLinkedInterconnectAttachments(value *NetworkConnectivitySpokeLinkedInterconnectAttachments) {
	if err := n.validatePutLinkedInterconnectAttachmentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putLinkedInterconnectAttachments",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkConnectivitySpoke) PutLinkedProducerVpcNetwork(value *NetworkConnectivitySpokeLinkedProducerVpcNetwork) {
	if err := n.validatePutLinkedProducerVpcNetworkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putLinkedProducerVpcNetwork",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkConnectivitySpoke) PutLinkedRouterApplianceInstances(value *NetworkConnectivitySpokeLinkedRouterApplianceInstances) {
	if err := n.validatePutLinkedRouterApplianceInstancesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putLinkedRouterApplianceInstances",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkConnectivitySpoke) PutLinkedVpcNetwork(value *NetworkConnectivitySpokeLinkedVpcNetwork) {
	if err := n.validatePutLinkedVpcNetworkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putLinkedVpcNetwork",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkConnectivitySpoke) PutLinkedVpnTunnels(value *NetworkConnectivitySpokeLinkedVpnTunnels) {
	if err := n.validatePutLinkedVpnTunnelsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putLinkedVpnTunnels",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkConnectivitySpoke) PutTimeouts(value *NetworkConnectivitySpokeTimeouts) {
	if err := n.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkConnectivitySpoke) ResetDescription() {
	_jsii_.InvokeVoid(
		n,
		"resetDescription",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectivitySpoke) ResetGroup() {
	_jsii_.InvokeVoid(
		n,
		"resetGroup",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectivitySpoke) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectivitySpoke) ResetLabels() {
	_jsii_.InvokeVoid(
		n,
		"resetLabels",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectivitySpoke) ResetLinkedInterconnectAttachments() {
	_jsii_.InvokeVoid(
		n,
		"resetLinkedInterconnectAttachments",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectivitySpoke) ResetLinkedProducerVpcNetwork() {
	_jsii_.InvokeVoid(
		n,
		"resetLinkedProducerVpcNetwork",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectivitySpoke) ResetLinkedRouterApplianceInstances() {
	_jsii_.InvokeVoid(
		n,
		"resetLinkedRouterApplianceInstances",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectivitySpoke) ResetLinkedVpcNetwork() {
	_jsii_.InvokeVoid(
		n,
		"resetLinkedVpcNetwork",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectivitySpoke) ResetLinkedVpnTunnels() {
	_jsii_.InvokeVoid(
		n,
		"resetLinkedVpnTunnels",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectivitySpoke) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectivitySpoke) ResetProject() {
	_jsii_.InvokeVoid(
		n,
		"resetProject",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectivitySpoke) ResetTimeouts() {
	_jsii_.InvokeVoid(
		n,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectivitySpoke) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectivitySpoke) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectivitySpoke) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectivitySpoke) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectivitySpoke) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectivitySpoke) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

