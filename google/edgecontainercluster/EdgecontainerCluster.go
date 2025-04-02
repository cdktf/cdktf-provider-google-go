// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package edgecontainercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v14/edgecontainercluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/edgecontainer_cluster google_edgecontainer_cluster}.
type EdgecontainerCluster interface {
	cdktf.TerraformResource
	Authorization() EdgecontainerClusterAuthorizationOutputReference
	AuthorizationInput() *EdgecontainerClusterAuthorization
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterCaCertificate() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ControlPlane() EdgecontainerClusterControlPlaneOutputReference
	ControlPlaneEncryption() EdgecontainerClusterControlPlaneEncryptionOutputReference
	ControlPlaneEncryptionInput() *EdgecontainerClusterControlPlaneEncryption
	ControlPlaneInput() *EdgecontainerClusterControlPlane
	ControlPlaneVersion() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreateTime() *string
	DefaultMaxPodsPerNode() *float64
	SetDefaultMaxPodsPerNode(val *float64)
	DefaultMaxPodsPerNodeInput() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EffectiveLabels() cdktf.StringMap
	Endpoint() *string
	ExternalLoadBalancerIpv4AddressPools() *[]*string
	SetExternalLoadBalancerIpv4AddressPools(val *[]*string)
	ExternalLoadBalancerIpv4AddressPoolsInput() *[]*string
	Fleet() EdgecontainerClusterFleetOutputReference
	FleetInput() *EdgecontainerClusterFleet
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
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MaintenanceEvents() EdgecontainerClusterMaintenanceEventsList
	MaintenancePolicy() EdgecontainerClusterMaintenancePolicyOutputReference
	MaintenancePolicyInput() *EdgecontainerClusterMaintenancePolicy
	Name() *string
	SetName(val *string)
	NameInput() *string
	Networking() EdgecontainerClusterNetworkingOutputReference
	NetworkingInput() *EdgecontainerClusterNetworking
	// The tree node.
	Node() constructs.Node
	NodeVersion() *string
	Port() *float64
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
	ReleaseChannel() *string
	SetReleaseChannel(val *string)
	ReleaseChannelInput() *string
	Status() *string
	SystemAddonsConfig() EdgecontainerClusterSystemAddonsConfigOutputReference
	SystemAddonsConfigInput() *EdgecontainerClusterSystemAddonsConfig
	TargetVersion() *string
	SetTargetVersion(val *string)
	TargetVersionInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() EdgecontainerClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutAuthorization(value *EdgecontainerClusterAuthorization)
	PutControlPlane(value *EdgecontainerClusterControlPlane)
	PutControlPlaneEncryption(value *EdgecontainerClusterControlPlaneEncryption)
	PutFleet(value *EdgecontainerClusterFleet)
	PutMaintenancePolicy(value *EdgecontainerClusterMaintenancePolicy)
	PutNetworking(value *EdgecontainerClusterNetworking)
	PutSystemAddonsConfig(value *EdgecontainerClusterSystemAddonsConfig)
	PutTimeouts(value *EdgecontainerClusterTimeouts)
	ResetControlPlane()
	ResetControlPlaneEncryption()
	ResetDefaultMaxPodsPerNode()
	ResetExternalLoadBalancerIpv4AddressPools()
	ResetId()
	ResetLabels()
	ResetMaintenancePolicy()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetReleaseChannel()
	ResetSystemAddonsConfig()
	ResetTargetVersion()
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

// The jsii proxy struct for EdgecontainerCluster
type jsiiProxy_EdgecontainerCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_EdgecontainerCluster) Authorization() EdgecontainerClusterAuthorizationOutputReference {
	var returns EdgecontainerClusterAuthorizationOutputReference
	_jsii_.Get(
		j,
		"authorization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) AuthorizationInput() *EdgecontainerClusterAuthorization {
	var returns *EdgecontainerClusterAuthorization
	_jsii_.Get(
		j,
		"authorizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) ClusterCaCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterCaCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) ControlPlane() EdgecontainerClusterControlPlaneOutputReference {
	var returns EdgecontainerClusterControlPlaneOutputReference
	_jsii_.Get(
		j,
		"controlPlane",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) ControlPlaneEncryption() EdgecontainerClusterControlPlaneEncryptionOutputReference {
	var returns EdgecontainerClusterControlPlaneEncryptionOutputReference
	_jsii_.Get(
		j,
		"controlPlaneEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) ControlPlaneEncryptionInput() *EdgecontainerClusterControlPlaneEncryption {
	var returns *EdgecontainerClusterControlPlaneEncryption
	_jsii_.Get(
		j,
		"controlPlaneEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) ControlPlaneInput() *EdgecontainerClusterControlPlane {
	var returns *EdgecontainerClusterControlPlane
	_jsii_.Get(
		j,
		"controlPlaneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) ControlPlaneVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"controlPlaneVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) DefaultMaxPodsPerNode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultMaxPodsPerNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) DefaultMaxPodsPerNodeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultMaxPodsPerNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) ExternalLoadBalancerIpv4AddressPools() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"externalLoadBalancerIpv4AddressPools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) ExternalLoadBalancerIpv4AddressPoolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"externalLoadBalancerIpv4AddressPoolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) Fleet() EdgecontainerClusterFleetOutputReference {
	var returns EdgecontainerClusterFleetOutputReference
	_jsii_.Get(
		j,
		"fleet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) FleetInput() *EdgecontainerClusterFleet {
	var returns *EdgecontainerClusterFleet
	_jsii_.Get(
		j,
		"fleetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) MaintenanceEvents() EdgecontainerClusterMaintenanceEventsList {
	var returns EdgecontainerClusterMaintenanceEventsList
	_jsii_.Get(
		j,
		"maintenanceEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) MaintenancePolicy() EdgecontainerClusterMaintenancePolicyOutputReference {
	var returns EdgecontainerClusterMaintenancePolicyOutputReference
	_jsii_.Get(
		j,
		"maintenancePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) MaintenancePolicyInput() *EdgecontainerClusterMaintenancePolicy {
	var returns *EdgecontainerClusterMaintenancePolicy
	_jsii_.Get(
		j,
		"maintenancePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) Networking() EdgecontainerClusterNetworkingOutputReference {
	var returns EdgecontainerClusterNetworkingOutputReference
	_jsii_.Get(
		j,
		"networking",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) NetworkingInput() *EdgecontainerClusterNetworking {
	var returns *EdgecontainerClusterNetworking
	_jsii_.Get(
		j,
		"networkingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) NodeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) ReleaseChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) ReleaseChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) SystemAddonsConfig() EdgecontainerClusterSystemAddonsConfigOutputReference {
	var returns EdgecontainerClusterSystemAddonsConfigOutputReference
	_jsii_.Get(
		j,
		"systemAddonsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) SystemAddonsConfigInput() *EdgecontainerClusterSystemAddonsConfig {
	var returns *EdgecontainerClusterSystemAddonsConfig
	_jsii_.Get(
		j,
		"systemAddonsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) TargetVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) TargetVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) Timeouts() EdgecontainerClusterTimeoutsOutputReference {
	var returns EdgecontainerClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EdgecontainerCluster) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/edgecontainer_cluster google_edgecontainer_cluster} Resource.
func NewEdgecontainerCluster(scope constructs.Construct, id *string, config *EdgecontainerClusterConfig) EdgecontainerCluster {
	_init_.Initialize()

	if err := validateNewEdgecontainerClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_EdgecontainerCluster{}

	_jsii_.Create(
		"@cdktf/provider-google.edgecontainerCluster.EdgecontainerCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/edgecontainer_cluster google_edgecontainer_cluster} Resource.
func NewEdgecontainerCluster_Override(e EdgecontainerCluster, scope constructs.Construct, id *string, config *EdgecontainerClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.edgecontainerCluster.EdgecontainerCluster",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EdgecontainerCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_EdgecontainerCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EdgecontainerCluster)SetDefaultMaxPodsPerNode(val *float64) {
	if err := j.validateSetDefaultMaxPodsPerNodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultMaxPodsPerNode",
		val,
	)
}

func (j *jsiiProxy_EdgecontainerCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EdgecontainerCluster)SetExternalLoadBalancerIpv4AddressPools(val *[]*string) {
	if err := j.validateSetExternalLoadBalancerIpv4AddressPoolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalLoadBalancerIpv4AddressPools",
		val,
	)
}

func (j *jsiiProxy_EdgecontainerCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_EdgecontainerCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_EdgecontainerCluster)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_EdgecontainerCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EdgecontainerCluster)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_EdgecontainerCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_EdgecontainerCluster)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_EdgecontainerCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EdgecontainerCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_EdgecontainerCluster)SetReleaseChannel(val *string) {
	if err := j.validateSetReleaseChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"releaseChannel",
		val,
	)
}

func (j *jsiiProxy_EdgecontainerCluster)SetTargetVersion(val *string) {
	if err := j.validateSetTargetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetVersion",
		val,
	)
}

// Generates CDKTF code for importing a EdgecontainerCluster resource upon running "cdktf plan <stack-name>".
func EdgecontainerCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateEdgecontainerCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.edgecontainerCluster.EdgecontainerCluster",
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
func EdgecontainerCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEdgecontainerCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.edgecontainerCluster.EdgecontainerCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EdgecontainerCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEdgecontainerCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.edgecontainerCluster.EdgecontainerCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EdgecontainerCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEdgecontainerCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.edgecontainerCluster.EdgecontainerCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EdgecontainerCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.edgecontainerCluster.EdgecontainerCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EdgecontainerCluster) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_EdgecontainerCluster) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_EdgecontainerCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgecontainerCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgecontainerCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgecontainerCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgecontainerCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgecontainerCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgecontainerCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgecontainerCluster) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgecontainerCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgecontainerCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgecontainerCluster) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_EdgecontainerCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgecontainerCluster) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EdgecontainerCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_EdgecontainerCluster) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EdgecontainerCluster) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EdgecontainerCluster) PutAuthorization(value *EdgecontainerClusterAuthorization) {
	if err := e.validatePutAuthorizationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAuthorization",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EdgecontainerCluster) PutControlPlane(value *EdgecontainerClusterControlPlane) {
	if err := e.validatePutControlPlaneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putControlPlane",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EdgecontainerCluster) PutControlPlaneEncryption(value *EdgecontainerClusterControlPlaneEncryption) {
	if err := e.validatePutControlPlaneEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putControlPlaneEncryption",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EdgecontainerCluster) PutFleet(value *EdgecontainerClusterFleet) {
	if err := e.validatePutFleetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putFleet",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EdgecontainerCluster) PutMaintenancePolicy(value *EdgecontainerClusterMaintenancePolicy) {
	if err := e.validatePutMaintenancePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putMaintenancePolicy",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EdgecontainerCluster) PutNetworking(value *EdgecontainerClusterNetworking) {
	if err := e.validatePutNetworkingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNetworking",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EdgecontainerCluster) PutSystemAddonsConfig(value *EdgecontainerClusterSystemAddonsConfig) {
	if err := e.validatePutSystemAddonsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putSystemAddonsConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EdgecontainerCluster) PutTimeouts(value *EdgecontainerClusterTimeouts) {
	if err := e.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EdgecontainerCluster) ResetControlPlane() {
	_jsii_.InvokeVoid(
		e,
		"resetControlPlane",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgecontainerCluster) ResetControlPlaneEncryption() {
	_jsii_.InvokeVoid(
		e,
		"resetControlPlaneEncryption",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgecontainerCluster) ResetDefaultMaxPodsPerNode() {
	_jsii_.InvokeVoid(
		e,
		"resetDefaultMaxPodsPerNode",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgecontainerCluster) ResetExternalLoadBalancerIpv4AddressPools() {
	_jsii_.InvokeVoid(
		e,
		"resetExternalLoadBalancerIpv4AddressPools",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgecontainerCluster) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgecontainerCluster) ResetLabels() {
	_jsii_.InvokeVoid(
		e,
		"resetLabels",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgecontainerCluster) ResetMaintenancePolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetMaintenancePolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgecontainerCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgecontainerCluster) ResetProject() {
	_jsii_.InvokeVoid(
		e,
		"resetProject",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgecontainerCluster) ResetReleaseChannel() {
	_jsii_.InvokeVoid(
		e,
		"resetReleaseChannel",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgecontainerCluster) ResetSystemAddonsConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetSystemAddonsConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgecontainerCluster) ResetTargetVersion() {
	_jsii_.InvokeVoid(
		e,
		"resetTargetVersion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgecontainerCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		e,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EdgecontainerCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgecontainerCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgecontainerCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgecontainerCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgecontainerCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EdgecontainerCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

