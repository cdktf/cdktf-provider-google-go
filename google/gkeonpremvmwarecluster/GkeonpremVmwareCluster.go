// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonpremvmwarecluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v13/gkeonpremvmwarecluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/5.19.0/docs/resources/gkeonprem_vmware_cluster google_gkeonprem_vmware_cluster}.
type GkeonpremVmwareCluster interface {
	cdktf.TerraformResource
	AdminClusterMembership() *string
	SetAdminClusterMembership(val *string)
	AdminClusterMembershipInput() *string
	Annotations() *map[string]*string
	SetAnnotations(val *map[string]*string)
	AnnotationsInput() *map[string]*string
	AntiAffinityGroups() GkeonpremVmwareClusterAntiAffinityGroupsOutputReference
	AntiAffinityGroupsInput() *GkeonpremVmwareClusterAntiAffinityGroups
	Authorization() GkeonpremVmwareClusterAuthorizationOutputReference
	AuthorizationInput() *GkeonpremVmwareClusterAuthorization
	AutoRepairConfig() GkeonpremVmwareClusterAutoRepairConfigOutputReference
	AutoRepairConfigInput() *GkeonpremVmwareClusterAutoRepairConfig
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ControlPlaneNode() GkeonpremVmwareClusterControlPlaneNodeOutputReference
	ControlPlaneNodeInput() *GkeonpremVmwareClusterControlPlaneNode
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreateTime() *string
	DataplaneV2() GkeonpremVmwareClusterDataplaneV2OutputReference
	DataplaneV2Input() *GkeonpremVmwareClusterDataplaneV2
	DeleteTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EffectiveAnnotations() cdktf.StringMap
	EnableControlPlaneV2() interface{}
	SetEnableControlPlaneV2(val interface{})
	EnableControlPlaneV2Input() interface{}
	Endpoint() *string
	Etag() *string
	Fleet() GkeonpremVmwareClusterFleetList
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoadBalancer() GkeonpremVmwareClusterLoadBalancerOutputReference
	LoadBalancerInput() *GkeonpremVmwareClusterLoadBalancer
	LocalName() *string
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkConfig() GkeonpremVmwareClusterNetworkConfigOutputReference
	NetworkConfigInput() *GkeonpremVmwareClusterNetworkConfig
	// The tree node.
	Node() constructs.Node
	OnPremVersion() *string
	SetOnPremVersion(val *string)
	OnPremVersionInput() *string
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
	Reconciling() cdktf.IResolvable
	State() *string
	Status() GkeonpremVmwareClusterStatusList
	Storage() GkeonpremVmwareClusterStorageOutputReference
	StorageInput() *GkeonpremVmwareClusterStorage
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GkeonpremVmwareClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
	Uid() *string
	UpdateTime() *string
	UpgradePolicy() GkeonpremVmwareClusterUpgradePolicyOutputReference
	UpgradePolicyInput() *GkeonpremVmwareClusterUpgradePolicy
	ValidationCheck() GkeonpremVmwareClusterValidationCheckList
	Vcenter() GkeonpremVmwareClusterVcenterOutputReference
	VcenterInput() *GkeonpremVmwareClusterVcenter
	VmTrackingEnabled() interface{}
	SetVmTrackingEnabled(val interface{})
	VmTrackingEnabledInput() interface{}
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
	PutAntiAffinityGroups(value *GkeonpremVmwareClusterAntiAffinityGroups)
	PutAuthorization(value *GkeonpremVmwareClusterAuthorization)
	PutAutoRepairConfig(value *GkeonpremVmwareClusterAutoRepairConfig)
	PutControlPlaneNode(value *GkeonpremVmwareClusterControlPlaneNode)
	PutDataplaneV2(value *GkeonpremVmwareClusterDataplaneV2)
	PutLoadBalancer(value *GkeonpremVmwareClusterLoadBalancer)
	PutNetworkConfig(value *GkeonpremVmwareClusterNetworkConfig)
	PutStorage(value *GkeonpremVmwareClusterStorage)
	PutTimeouts(value *GkeonpremVmwareClusterTimeouts)
	PutUpgradePolicy(value *GkeonpremVmwareClusterUpgradePolicy)
	PutVcenter(value *GkeonpremVmwareClusterVcenter)
	ResetAnnotations()
	ResetAntiAffinityGroups()
	ResetAuthorization()
	ResetAutoRepairConfig()
	ResetDataplaneV2()
	ResetDescription()
	ResetEnableControlPlaneV2()
	ResetId()
	ResetLoadBalancer()
	ResetNetworkConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetStorage()
	ResetTimeouts()
	ResetUpgradePolicy()
	ResetVcenter()
	ResetVmTrackingEnabled()
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

// The jsii proxy struct for GkeonpremVmwareCluster
type jsiiProxy_GkeonpremVmwareCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GkeonpremVmwareCluster) AdminClusterMembership() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminClusterMembership",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) AdminClusterMembershipInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminClusterMembershipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) Annotations() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) AnnotationsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) AntiAffinityGroups() GkeonpremVmwareClusterAntiAffinityGroupsOutputReference {
	var returns GkeonpremVmwareClusterAntiAffinityGroupsOutputReference
	_jsii_.Get(
		j,
		"antiAffinityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) AntiAffinityGroupsInput() *GkeonpremVmwareClusterAntiAffinityGroups {
	var returns *GkeonpremVmwareClusterAntiAffinityGroups
	_jsii_.Get(
		j,
		"antiAffinityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) Authorization() GkeonpremVmwareClusterAuthorizationOutputReference {
	var returns GkeonpremVmwareClusterAuthorizationOutputReference
	_jsii_.Get(
		j,
		"authorization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) AuthorizationInput() *GkeonpremVmwareClusterAuthorization {
	var returns *GkeonpremVmwareClusterAuthorization
	_jsii_.Get(
		j,
		"authorizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) AutoRepairConfig() GkeonpremVmwareClusterAutoRepairConfigOutputReference {
	var returns GkeonpremVmwareClusterAutoRepairConfigOutputReference
	_jsii_.Get(
		j,
		"autoRepairConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) AutoRepairConfigInput() *GkeonpremVmwareClusterAutoRepairConfig {
	var returns *GkeonpremVmwareClusterAutoRepairConfig
	_jsii_.Get(
		j,
		"autoRepairConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) ControlPlaneNode() GkeonpremVmwareClusterControlPlaneNodeOutputReference {
	var returns GkeonpremVmwareClusterControlPlaneNodeOutputReference
	_jsii_.Get(
		j,
		"controlPlaneNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) ControlPlaneNodeInput() *GkeonpremVmwareClusterControlPlaneNode {
	var returns *GkeonpremVmwareClusterControlPlaneNode
	_jsii_.Get(
		j,
		"controlPlaneNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) DataplaneV2() GkeonpremVmwareClusterDataplaneV2OutputReference {
	var returns GkeonpremVmwareClusterDataplaneV2OutputReference
	_jsii_.Get(
		j,
		"dataplaneV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) DataplaneV2Input() *GkeonpremVmwareClusterDataplaneV2 {
	var returns *GkeonpremVmwareClusterDataplaneV2
	_jsii_.Get(
		j,
		"dataplaneV2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) DeleteTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) EffectiveAnnotations() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) EnableControlPlaneV2() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableControlPlaneV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) EnableControlPlaneV2Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableControlPlaneV2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) Fleet() GkeonpremVmwareClusterFleetList {
	var returns GkeonpremVmwareClusterFleetList
	_jsii_.Get(
		j,
		"fleet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) LoadBalancer() GkeonpremVmwareClusterLoadBalancerOutputReference {
	var returns GkeonpremVmwareClusterLoadBalancerOutputReference
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) LoadBalancerInput() *GkeonpremVmwareClusterLoadBalancer {
	var returns *GkeonpremVmwareClusterLoadBalancer
	_jsii_.Get(
		j,
		"loadBalancerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) LocalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) NetworkConfig() GkeonpremVmwareClusterNetworkConfigOutputReference {
	var returns GkeonpremVmwareClusterNetworkConfigOutputReference
	_jsii_.Get(
		j,
		"networkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) NetworkConfigInput() *GkeonpremVmwareClusterNetworkConfig {
	var returns *GkeonpremVmwareClusterNetworkConfig
	_jsii_.Get(
		j,
		"networkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) OnPremVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onPremVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) OnPremVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onPremVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) Reconciling() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"reconciling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) Status() GkeonpremVmwareClusterStatusList {
	var returns GkeonpremVmwareClusterStatusList
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) Storage() GkeonpremVmwareClusterStorageOutputReference {
	var returns GkeonpremVmwareClusterStorageOutputReference
	_jsii_.Get(
		j,
		"storage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) StorageInput() *GkeonpremVmwareClusterStorage {
	var returns *GkeonpremVmwareClusterStorage
	_jsii_.Get(
		j,
		"storageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) Timeouts() GkeonpremVmwareClusterTimeoutsOutputReference {
	var returns GkeonpremVmwareClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) Uid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) UpgradePolicy() GkeonpremVmwareClusterUpgradePolicyOutputReference {
	var returns GkeonpremVmwareClusterUpgradePolicyOutputReference
	_jsii_.Get(
		j,
		"upgradePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) UpgradePolicyInput() *GkeonpremVmwareClusterUpgradePolicy {
	var returns *GkeonpremVmwareClusterUpgradePolicy
	_jsii_.Get(
		j,
		"upgradePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) ValidationCheck() GkeonpremVmwareClusterValidationCheckList {
	var returns GkeonpremVmwareClusterValidationCheckList
	_jsii_.Get(
		j,
		"validationCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) Vcenter() GkeonpremVmwareClusterVcenterOutputReference {
	var returns GkeonpremVmwareClusterVcenterOutputReference
	_jsii_.Get(
		j,
		"vcenter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) VcenterInput() *GkeonpremVmwareClusterVcenter {
	var returns *GkeonpremVmwareClusterVcenter
	_jsii_.Get(
		j,
		"vcenterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) VmTrackingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vmTrackingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremVmwareCluster) VmTrackingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vmTrackingEnabledInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/5.19.0/docs/resources/gkeonprem_vmware_cluster google_gkeonprem_vmware_cluster} Resource.
func NewGkeonpremVmwareCluster(scope constructs.Construct, id *string, config *GkeonpremVmwareClusterConfig) GkeonpremVmwareCluster {
	_init_.Initialize()

	if err := validateNewGkeonpremVmwareClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GkeonpremVmwareCluster{}

	_jsii_.Create(
		"@cdktf/provider-google.gkeonpremVmwareCluster.GkeonpremVmwareCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/5.19.0/docs/resources/gkeonprem_vmware_cluster google_gkeonprem_vmware_cluster} Resource.
func NewGkeonpremVmwareCluster_Override(g GkeonpremVmwareCluster, scope constructs.Construct, id *string, config *GkeonpremVmwareClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.gkeonpremVmwareCluster.GkeonpremVmwareCluster",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GkeonpremVmwareCluster)SetAdminClusterMembership(val *string) {
	if err := j.validateSetAdminClusterMembershipParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminClusterMembership",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareCluster)SetAnnotations(val *map[string]*string) {
	if err := j.validateSetAnnotationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotations",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareCluster)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareCluster)SetEnableControlPlaneV2(val interface{}) {
	if err := j.validateSetEnableControlPlaneV2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableControlPlaneV2",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareCluster)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareCluster)SetOnPremVersion(val *string) {
	if err := j.validateSetOnPremVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onPremVersion",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareCluster)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GkeonpremVmwareCluster)SetVmTrackingEnabled(val interface{}) {
	if err := j.validateSetVmTrackingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmTrackingEnabled",
		val,
	)
}

// Generates CDKTF code for importing a GkeonpremVmwareCluster resource upon running "cdktf plan <stack-name>".
func GkeonpremVmwareCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGkeonpremVmwareCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.gkeonpremVmwareCluster.GkeonpremVmwareCluster",
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
func GkeonpremVmwareCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGkeonpremVmwareCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.gkeonpremVmwareCluster.GkeonpremVmwareCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GkeonpremVmwareCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGkeonpremVmwareCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.gkeonpremVmwareCluster.GkeonpremVmwareCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GkeonpremVmwareCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGkeonpremVmwareCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.gkeonpremVmwareCluster.GkeonpremVmwareCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GkeonpremVmwareCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.gkeonpremVmwareCluster.GkeonpremVmwareCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GkeonpremVmwareCluster) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GkeonpremVmwareCluster) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GkeonpremVmwareCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareCluster) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareCluster) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GkeonpremVmwareCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareCluster) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GkeonpremVmwareCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GkeonpremVmwareCluster) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GkeonpremVmwareCluster) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GkeonpremVmwareCluster) PutAntiAffinityGroups(value *GkeonpremVmwareClusterAntiAffinityGroups) {
	if err := g.validatePutAntiAffinityGroupsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAntiAffinityGroups",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremVmwareCluster) PutAuthorization(value *GkeonpremVmwareClusterAuthorization) {
	if err := g.validatePutAuthorizationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAuthorization",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremVmwareCluster) PutAutoRepairConfig(value *GkeonpremVmwareClusterAutoRepairConfig) {
	if err := g.validatePutAutoRepairConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAutoRepairConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremVmwareCluster) PutControlPlaneNode(value *GkeonpremVmwareClusterControlPlaneNode) {
	if err := g.validatePutControlPlaneNodeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putControlPlaneNode",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremVmwareCluster) PutDataplaneV2(value *GkeonpremVmwareClusterDataplaneV2) {
	if err := g.validatePutDataplaneV2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDataplaneV2",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremVmwareCluster) PutLoadBalancer(value *GkeonpremVmwareClusterLoadBalancer) {
	if err := g.validatePutLoadBalancerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLoadBalancer",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremVmwareCluster) PutNetworkConfig(value *GkeonpremVmwareClusterNetworkConfig) {
	if err := g.validatePutNetworkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNetworkConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremVmwareCluster) PutStorage(value *GkeonpremVmwareClusterStorage) {
	if err := g.validatePutStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStorage",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremVmwareCluster) PutTimeouts(value *GkeonpremVmwareClusterTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremVmwareCluster) PutUpgradePolicy(value *GkeonpremVmwareClusterUpgradePolicy) {
	if err := g.validatePutUpgradePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putUpgradePolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremVmwareCluster) PutVcenter(value *GkeonpremVmwareClusterVcenter) {
	if err := g.validatePutVcenterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVcenter",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremVmwareCluster) ResetAnnotations() {
	_jsii_.InvokeVoid(
		g,
		"resetAnnotations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareCluster) ResetAntiAffinityGroups() {
	_jsii_.InvokeVoid(
		g,
		"resetAntiAffinityGroups",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareCluster) ResetAuthorization() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthorization",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareCluster) ResetAutoRepairConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetAutoRepairConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareCluster) ResetDataplaneV2() {
	_jsii_.InvokeVoid(
		g,
		"resetDataplaneV2",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareCluster) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareCluster) ResetEnableControlPlaneV2() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableControlPlaneV2",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareCluster) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareCluster) ResetLoadBalancer() {
	_jsii_.InvokeVoid(
		g,
		"resetLoadBalancer",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareCluster) ResetNetworkConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareCluster) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareCluster) ResetStorage() {
	_jsii_.InvokeVoid(
		g,
		"resetStorage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareCluster) ResetUpgradePolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetUpgradePolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareCluster) ResetVcenter() {
	_jsii_.InvokeVoid(
		g,
		"resetVcenter",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareCluster) ResetVmTrackingEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetVmTrackingEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremVmwareCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremVmwareCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

