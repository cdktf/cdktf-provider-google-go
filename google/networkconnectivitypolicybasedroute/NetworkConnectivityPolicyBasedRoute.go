// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkconnectivitypolicybasedroute

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v13/networkconnectivitypolicybasedroute/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/5.32.0/docs/resources/network_connectivity_policy_based_route google_network_connectivity_policy_based_route}.
type NetworkConnectivityPolicyBasedRoute interface {
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
	Filter() NetworkConnectivityPolicyBasedRouteFilterOutputReference
	FilterInput() *NetworkConnectivityPolicyBasedRouteFilter
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InterconnectAttachment() NetworkConnectivityPolicyBasedRouteInterconnectAttachmentOutputReference
	InterconnectAttachmentInput() *NetworkConnectivityPolicyBasedRouteInterconnectAttachment
	Kind() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Network() *string
	SetNetwork(val *string)
	NetworkInput() *string
	NextHopIlbIp() *string
	SetNextHopIlbIp(val *string)
	NextHopIlbIpInput() *string
	NextHopOtherRoutes() *string
	SetNextHopOtherRoutes(val *string)
	NextHopOtherRoutesInput() *string
	// The tree node.
	Node() constructs.Node
	Priority() *float64
	SetPriority(val *float64)
	PriorityInput() *float64
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() NetworkConnectivityPolicyBasedRouteTimeoutsOutputReference
	TimeoutsInput() interface{}
	UpdateTime() *string
	VirtualMachine() NetworkConnectivityPolicyBasedRouteVirtualMachineOutputReference
	VirtualMachineInput() *NetworkConnectivityPolicyBasedRouteVirtualMachine
	Warnings() NetworkConnectivityPolicyBasedRouteWarningsList
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
	PutFilter(value *NetworkConnectivityPolicyBasedRouteFilter)
	PutInterconnectAttachment(value *NetworkConnectivityPolicyBasedRouteInterconnectAttachment)
	PutTimeouts(value *NetworkConnectivityPolicyBasedRouteTimeouts)
	PutVirtualMachine(value *NetworkConnectivityPolicyBasedRouteVirtualMachine)
	ResetDescription()
	ResetId()
	ResetInterconnectAttachment()
	ResetLabels()
	ResetNextHopIlbIp()
	ResetNextHopOtherRoutes()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPriority()
	ResetProject()
	ResetTimeouts()
	ResetVirtualMachine()
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

// The jsii proxy struct for NetworkConnectivityPolicyBasedRoute
type jsiiProxy_NetworkConnectivityPolicyBasedRoute struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) Filter() NetworkConnectivityPolicyBasedRouteFilterOutputReference {
	var returns NetworkConnectivityPolicyBasedRouteFilterOutputReference
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) FilterInput() *NetworkConnectivityPolicyBasedRouteFilter {
	var returns *NetworkConnectivityPolicyBasedRouteFilter
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) InterconnectAttachment() NetworkConnectivityPolicyBasedRouteInterconnectAttachmentOutputReference {
	var returns NetworkConnectivityPolicyBasedRouteInterconnectAttachmentOutputReference
	_jsii_.Get(
		j,
		"interconnectAttachment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) InterconnectAttachmentInput() *NetworkConnectivityPolicyBasedRouteInterconnectAttachment {
	var returns *NetworkConnectivityPolicyBasedRouteInterconnectAttachment
	_jsii_.Get(
		j,
		"interconnectAttachmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) NextHopIlbIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextHopIlbIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) NextHopIlbIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextHopIlbIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) NextHopOtherRoutes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextHopOtherRoutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) NextHopOtherRoutesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextHopOtherRoutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) Timeouts() NetworkConnectivityPolicyBasedRouteTimeoutsOutputReference {
	var returns NetworkConnectivityPolicyBasedRouteTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) VirtualMachine() NetworkConnectivityPolicyBasedRouteVirtualMachineOutputReference {
	var returns NetworkConnectivityPolicyBasedRouteVirtualMachineOutputReference
	_jsii_.Get(
		j,
		"virtualMachine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) VirtualMachineInput() *NetworkConnectivityPolicyBasedRouteVirtualMachine {
	var returns *NetworkConnectivityPolicyBasedRouteVirtualMachine
	_jsii_.Get(
		j,
		"virtualMachineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute) Warnings() NetworkConnectivityPolicyBasedRouteWarningsList {
	var returns NetworkConnectivityPolicyBasedRouteWarningsList
	_jsii_.Get(
		j,
		"warnings",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/5.32.0/docs/resources/network_connectivity_policy_based_route google_network_connectivity_policy_based_route} Resource.
func NewNetworkConnectivityPolicyBasedRoute(scope constructs.Construct, id *string, config *NetworkConnectivityPolicyBasedRouteConfig) NetworkConnectivityPolicyBasedRoute {
	_init_.Initialize()

	if err := validateNewNetworkConnectivityPolicyBasedRouteParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkConnectivityPolicyBasedRoute{}

	_jsii_.Create(
		"@cdktf/provider-google.networkConnectivityPolicyBasedRoute.NetworkConnectivityPolicyBasedRoute",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/5.32.0/docs/resources/network_connectivity_policy_based_route google_network_connectivity_policy_based_route} Resource.
func NewNetworkConnectivityPolicyBasedRoute_Override(n NetworkConnectivityPolicyBasedRoute, scope constructs.Construct, id *string, config *NetworkConnectivityPolicyBasedRouteConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.networkConnectivityPolicyBasedRoute.NetworkConnectivityPolicyBasedRoute",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute)SetNextHopIlbIp(val *string) {
	if err := j.validateSetNextHopIlbIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nextHopIlbIp",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute)SetNextHopOtherRoutes(val *string) {
	if err := j.validateSetNextHopOtherRoutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nextHopOtherRoutes",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NetworkConnectivityPolicyBasedRoute)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a NetworkConnectivityPolicyBasedRoute resource upon running "cdktf plan <stack-name>".
func NetworkConnectivityPolicyBasedRoute_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateNetworkConnectivityPolicyBasedRoute_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.networkConnectivityPolicyBasedRoute.NetworkConnectivityPolicyBasedRoute",
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
func NetworkConnectivityPolicyBasedRoute_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkConnectivityPolicyBasedRoute_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.networkConnectivityPolicyBasedRoute.NetworkConnectivityPolicyBasedRoute",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetworkConnectivityPolicyBasedRoute_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkConnectivityPolicyBasedRoute_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.networkConnectivityPolicyBasedRoute.NetworkConnectivityPolicyBasedRoute",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetworkConnectivityPolicyBasedRoute_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetworkConnectivityPolicyBasedRoute_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.networkConnectivityPolicyBasedRoute.NetworkConnectivityPolicyBasedRoute",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NetworkConnectivityPolicyBasedRoute_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.networkConnectivityPolicyBasedRoute.NetworkConnectivityPolicyBasedRoute",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRoute) AddMoveTarget(moveTarget *string) {
	if err := n.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRoute) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRoute) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRoute) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRoute) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRoute) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRoute) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRoute) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRoute) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRoute) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRoute) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRoute) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRoute) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := n.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRoute) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRoute) MoveFromId(id *string) {
	if err := n.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveFromId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRoute) MoveTo(moveTarget *string, index interface{}) {
	if err := n.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRoute) MoveToId(id *string) {
	if err := n.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveToId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRoute) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRoute) PutFilter(value *NetworkConnectivityPolicyBasedRouteFilter) {
	if err := n.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putFilter",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRoute) PutInterconnectAttachment(value *NetworkConnectivityPolicyBasedRouteInterconnectAttachment) {
	if err := n.validatePutInterconnectAttachmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putInterconnectAttachment",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRoute) PutTimeouts(value *NetworkConnectivityPolicyBasedRouteTimeouts) {
	if err := n.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRoute) PutVirtualMachine(value *NetworkConnectivityPolicyBasedRouteVirtualMachine) {
	if err := n.validatePutVirtualMachineParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putVirtualMachine",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRoute) ResetDescription() {
	_jsii_.InvokeVoid(
		n,
		"resetDescription",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRoute) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRoute) ResetInterconnectAttachment() {
	_jsii_.InvokeVoid(
		n,
		"resetInterconnectAttachment",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRoute) ResetLabels() {
	_jsii_.InvokeVoid(
		n,
		"resetLabels",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRoute) ResetNextHopIlbIp() {
	_jsii_.InvokeVoid(
		n,
		"resetNextHopIlbIp",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRoute) ResetNextHopOtherRoutes() {
	_jsii_.InvokeVoid(
		n,
		"resetNextHopOtherRoutes",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRoute) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRoute) ResetPriority() {
	_jsii_.InvokeVoid(
		n,
		"resetPriority",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRoute) ResetProject() {
	_jsii_.InvokeVoid(
		n,
		"resetProject",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRoute) ResetTimeouts() {
	_jsii_.InvokeVoid(
		n,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRoute) ResetVirtualMachine() {
	_jsii_.InvokeVoid(
		n,
		"resetVirtualMachine",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRoute) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRoute) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRoute) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRoute) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRoute) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkConnectivityPolicyBasedRoute) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

