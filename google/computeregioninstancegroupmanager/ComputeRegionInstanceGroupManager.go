// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregioninstancegroupmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v14/computeregioninstancegroupmanager/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/compute_region_instance_group_manager google_compute_region_instance_group_manager}.
type ComputeRegionInstanceGroupManager interface {
	cdktf.TerraformResource
	AllInstancesConfig() ComputeRegionInstanceGroupManagerAllInstancesConfigOutputReference
	AllInstancesConfigInput() *ComputeRegionInstanceGroupManagerAllInstancesConfig
	AutoHealingPolicies() ComputeRegionInstanceGroupManagerAutoHealingPoliciesOutputReference
	AutoHealingPoliciesInput() *ComputeRegionInstanceGroupManagerAutoHealingPolicies
	BaseInstanceName() *string
	SetBaseInstanceName(val *string)
	BaseInstanceNameInput() *string
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
	CreationTimestamp() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DistributionPolicyTargetShape() *string
	SetDistributionPolicyTargetShape(val *string)
	DistributionPolicyTargetShapeInput() *string
	DistributionPolicyZones() *[]*string
	SetDistributionPolicyZones(val *[]*string)
	DistributionPolicyZonesInput() *[]*string
	Fingerprint() *string
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
	InstanceFlexibilityPolicy() ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyOutputReference
	InstanceFlexibilityPolicyInput() *ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicy
	InstanceGroup() *string
	InstanceGroupManagerId() *float64
	InstanceLifecyclePolicy() ComputeRegionInstanceGroupManagerInstanceLifecyclePolicyOutputReference
	InstanceLifecyclePolicyInput() *ComputeRegionInstanceGroupManagerInstanceLifecyclePolicy
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ListManagedInstancesResults() *string
	SetListManagedInstancesResults(val *string)
	ListManagedInstancesResultsInput() *string
	Name() *string
	SetName(val *string)
	NamedPort() ComputeRegionInstanceGroupManagerNamedPortList
	NamedPortInput() interface{}
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
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	SelfLink() *string
	StandbyPolicy() ComputeRegionInstanceGroupManagerStandbyPolicyOutputReference
	StandbyPolicyInput() *ComputeRegionInstanceGroupManagerStandbyPolicy
	StatefulDisk() ComputeRegionInstanceGroupManagerStatefulDiskList
	StatefulDiskInput() interface{}
	StatefulExternalIp() ComputeRegionInstanceGroupManagerStatefulExternalIpList
	StatefulExternalIpInput() interface{}
	StatefulInternalIp() ComputeRegionInstanceGroupManagerStatefulInternalIpList
	StatefulInternalIpInput() interface{}
	Status() ComputeRegionInstanceGroupManagerStatusList
	TargetPools() *[]*string
	SetTargetPools(val *[]*string)
	TargetPoolsInput() *[]*string
	TargetSize() *float64
	SetTargetSize(val *float64)
	TargetSizeInput() *float64
	TargetStoppedSize() *float64
	SetTargetStoppedSize(val *float64)
	TargetStoppedSizeInput() *float64
	TargetSuspendedSize() *float64
	SetTargetSuspendedSize(val *float64)
	TargetSuspendedSizeInput() *float64
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ComputeRegionInstanceGroupManagerTimeoutsOutputReference
	TimeoutsInput() interface{}
	UpdatePolicy() ComputeRegionInstanceGroupManagerUpdatePolicyOutputReference
	UpdatePolicyInput() *ComputeRegionInstanceGroupManagerUpdatePolicy
	Version() ComputeRegionInstanceGroupManagerVersionList
	VersionInput() interface{}
	WaitForInstances() interface{}
	SetWaitForInstances(val interface{})
	WaitForInstancesInput() interface{}
	WaitForInstancesStatus() *string
	SetWaitForInstancesStatus(val *string)
	WaitForInstancesStatusInput() *string
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
	PutAllInstancesConfig(value *ComputeRegionInstanceGroupManagerAllInstancesConfig)
	PutAutoHealingPolicies(value *ComputeRegionInstanceGroupManagerAutoHealingPolicies)
	PutInstanceFlexibilityPolicy(value *ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicy)
	PutInstanceLifecyclePolicy(value *ComputeRegionInstanceGroupManagerInstanceLifecyclePolicy)
	PutNamedPort(value interface{})
	PutStandbyPolicy(value *ComputeRegionInstanceGroupManagerStandbyPolicy)
	PutStatefulDisk(value interface{})
	PutStatefulExternalIp(value interface{})
	PutStatefulInternalIp(value interface{})
	PutTimeouts(value *ComputeRegionInstanceGroupManagerTimeouts)
	PutUpdatePolicy(value *ComputeRegionInstanceGroupManagerUpdatePolicy)
	PutVersion(value interface{})
	ResetAllInstancesConfig()
	ResetAutoHealingPolicies()
	ResetDescription()
	ResetDistributionPolicyTargetShape()
	ResetDistributionPolicyZones()
	ResetId()
	ResetInstanceFlexibilityPolicy()
	ResetInstanceLifecyclePolicy()
	ResetListManagedInstancesResults()
	ResetNamedPort()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetRegion()
	ResetStandbyPolicy()
	ResetStatefulDisk()
	ResetStatefulExternalIp()
	ResetStatefulInternalIp()
	ResetTargetPools()
	ResetTargetSize()
	ResetTargetStoppedSize()
	ResetTargetSuspendedSize()
	ResetTimeouts()
	ResetUpdatePolicy()
	ResetWaitForInstances()
	ResetWaitForInstancesStatus()
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

// The jsii proxy struct for ComputeRegionInstanceGroupManager
type jsiiProxy_ComputeRegionInstanceGroupManager struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) AllInstancesConfig() ComputeRegionInstanceGroupManagerAllInstancesConfigOutputReference {
	var returns ComputeRegionInstanceGroupManagerAllInstancesConfigOutputReference
	_jsii_.Get(
		j,
		"allInstancesConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) AllInstancesConfigInput() *ComputeRegionInstanceGroupManagerAllInstancesConfig {
	var returns *ComputeRegionInstanceGroupManagerAllInstancesConfig
	_jsii_.Get(
		j,
		"allInstancesConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) AutoHealingPolicies() ComputeRegionInstanceGroupManagerAutoHealingPoliciesOutputReference {
	var returns ComputeRegionInstanceGroupManagerAutoHealingPoliciesOutputReference
	_jsii_.Get(
		j,
		"autoHealingPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) AutoHealingPoliciesInput() *ComputeRegionInstanceGroupManagerAutoHealingPolicies {
	var returns *ComputeRegionInstanceGroupManagerAutoHealingPolicies
	_jsii_.Get(
		j,
		"autoHealingPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) BaseInstanceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseInstanceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) BaseInstanceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseInstanceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) CreationTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) DistributionPolicyTargetShape() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributionPolicyTargetShape",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) DistributionPolicyTargetShapeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributionPolicyTargetShapeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) DistributionPolicyZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"distributionPolicyZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) DistributionPolicyZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"distributionPolicyZonesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) Fingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) InstanceFlexibilityPolicy() ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyOutputReference {
	var returns ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicyOutputReference
	_jsii_.Get(
		j,
		"instanceFlexibilityPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) InstanceFlexibilityPolicyInput() *ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicy {
	var returns *ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicy
	_jsii_.Get(
		j,
		"instanceFlexibilityPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) InstanceGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) InstanceGroupManagerId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceGroupManagerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) InstanceLifecyclePolicy() ComputeRegionInstanceGroupManagerInstanceLifecyclePolicyOutputReference {
	var returns ComputeRegionInstanceGroupManagerInstanceLifecyclePolicyOutputReference
	_jsii_.Get(
		j,
		"instanceLifecyclePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) InstanceLifecyclePolicyInput() *ComputeRegionInstanceGroupManagerInstanceLifecyclePolicy {
	var returns *ComputeRegionInstanceGroupManagerInstanceLifecyclePolicy
	_jsii_.Get(
		j,
		"instanceLifecyclePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) ListManagedInstancesResults() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listManagedInstancesResults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) ListManagedInstancesResultsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listManagedInstancesResultsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) NamedPort() ComputeRegionInstanceGroupManagerNamedPortList {
	var returns ComputeRegionInstanceGroupManagerNamedPortList
	_jsii_.Get(
		j,
		"namedPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) NamedPortInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"namedPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) StandbyPolicy() ComputeRegionInstanceGroupManagerStandbyPolicyOutputReference {
	var returns ComputeRegionInstanceGroupManagerStandbyPolicyOutputReference
	_jsii_.Get(
		j,
		"standbyPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) StandbyPolicyInput() *ComputeRegionInstanceGroupManagerStandbyPolicy {
	var returns *ComputeRegionInstanceGroupManagerStandbyPolicy
	_jsii_.Get(
		j,
		"standbyPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) StatefulDisk() ComputeRegionInstanceGroupManagerStatefulDiskList {
	var returns ComputeRegionInstanceGroupManagerStatefulDiskList
	_jsii_.Get(
		j,
		"statefulDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) StatefulDiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statefulDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) StatefulExternalIp() ComputeRegionInstanceGroupManagerStatefulExternalIpList {
	var returns ComputeRegionInstanceGroupManagerStatefulExternalIpList
	_jsii_.Get(
		j,
		"statefulExternalIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) StatefulExternalIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statefulExternalIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) StatefulInternalIp() ComputeRegionInstanceGroupManagerStatefulInternalIpList {
	var returns ComputeRegionInstanceGroupManagerStatefulInternalIpList
	_jsii_.Get(
		j,
		"statefulInternalIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) StatefulInternalIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statefulInternalIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) Status() ComputeRegionInstanceGroupManagerStatusList {
	var returns ComputeRegionInstanceGroupManagerStatusList
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) TargetPools() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetPools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) TargetPoolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetPoolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) TargetSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) TargetSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) TargetStoppedSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetStoppedSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) TargetStoppedSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetStoppedSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) TargetSuspendedSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetSuspendedSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) TargetSuspendedSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetSuspendedSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) Timeouts() ComputeRegionInstanceGroupManagerTimeoutsOutputReference {
	var returns ComputeRegionInstanceGroupManagerTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) UpdatePolicy() ComputeRegionInstanceGroupManagerUpdatePolicyOutputReference {
	var returns ComputeRegionInstanceGroupManagerUpdatePolicyOutputReference
	_jsii_.Get(
		j,
		"updatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) UpdatePolicyInput() *ComputeRegionInstanceGroupManagerUpdatePolicy {
	var returns *ComputeRegionInstanceGroupManagerUpdatePolicy
	_jsii_.Get(
		j,
		"updatePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) Version() ComputeRegionInstanceGroupManagerVersionList {
	var returns ComputeRegionInstanceGroupManagerVersionList
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) VersionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) WaitForInstances() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) WaitForInstancesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) WaitForInstancesStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"waitForInstancesStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager) WaitForInstancesStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"waitForInstancesStatusInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/compute_region_instance_group_manager google_compute_region_instance_group_manager} Resource.
func NewComputeRegionInstanceGroupManager(scope constructs.Construct, id *string, config *ComputeRegionInstanceGroupManagerConfig) ComputeRegionInstanceGroupManager {
	_init_.Initialize()

	if err := validateNewComputeRegionInstanceGroupManagerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeRegionInstanceGroupManager{}

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionInstanceGroupManager.ComputeRegionInstanceGroupManager",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/compute_region_instance_group_manager google_compute_region_instance_group_manager} Resource.
func NewComputeRegionInstanceGroupManager_Override(c ComputeRegionInstanceGroupManager, scope constructs.Construct, id *string, config *ComputeRegionInstanceGroupManagerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionInstanceGroupManager.ComputeRegionInstanceGroupManager",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager)SetBaseInstanceName(val *string) {
	if err := j.validateSetBaseInstanceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseInstanceName",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager)SetDistributionPolicyTargetShape(val *string) {
	if err := j.validateSetDistributionPolicyTargetShapeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"distributionPolicyTargetShape",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager)SetDistributionPolicyZones(val *[]*string) {
	if err := j.validateSetDistributionPolicyZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"distributionPolicyZones",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager)SetListManagedInstancesResults(val *string) {
	if err := j.validateSetListManagedInstancesResultsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"listManagedInstancesResults",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager)SetTargetPools(val *[]*string) {
	if err := j.validateSetTargetPoolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetPools",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager)SetTargetSize(val *float64) {
	if err := j.validateSetTargetSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetSize",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager)SetTargetStoppedSize(val *float64) {
	if err := j.validateSetTargetStoppedSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetStoppedSize",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager)SetTargetSuspendedSize(val *float64) {
	if err := j.validateSetTargetSuspendedSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetSuspendedSize",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager)SetWaitForInstances(val interface{}) {
	if err := j.validateSetWaitForInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitForInstances",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionInstanceGroupManager)SetWaitForInstancesStatus(val *string) {
	if err := j.validateSetWaitForInstancesStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitForInstancesStatus",
		val,
	)
}

// Generates CDKTF code for importing a ComputeRegionInstanceGroupManager resource upon running "cdktf plan <stack-name>".
func ComputeRegionInstanceGroupManager_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateComputeRegionInstanceGroupManager_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeRegionInstanceGroupManager.ComputeRegionInstanceGroupManager",
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
func ComputeRegionInstanceGroupManager_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeRegionInstanceGroupManager_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeRegionInstanceGroupManager.ComputeRegionInstanceGroupManager",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeRegionInstanceGroupManager_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeRegionInstanceGroupManager_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeRegionInstanceGroupManager.ComputeRegionInstanceGroupManager",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeRegionInstanceGroupManager_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeRegionInstanceGroupManager_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeRegionInstanceGroupManager.ComputeRegionInstanceGroupManager",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ComputeRegionInstanceGroupManager_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.computeRegionInstanceGroupManager.ComputeRegionInstanceGroupManager",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) PutAllInstancesConfig(value *ComputeRegionInstanceGroupManagerAllInstancesConfig) {
	if err := c.validatePutAllInstancesConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAllInstancesConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) PutAutoHealingPolicies(value *ComputeRegionInstanceGroupManagerAutoHealingPolicies) {
	if err := c.validatePutAutoHealingPoliciesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAutoHealingPolicies",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) PutInstanceFlexibilityPolicy(value *ComputeRegionInstanceGroupManagerInstanceFlexibilityPolicy) {
	if err := c.validatePutInstanceFlexibilityPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putInstanceFlexibilityPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) PutInstanceLifecyclePolicy(value *ComputeRegionInstanceGroupManagerInstanceLifecyclePolicy) {
	if err := c.validatePutInstanceLifecyclePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putInstanceLifecyclePolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) PutNamedPort(value interface{}) {
	if err := c.validatePutNamedPortParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNamedPort",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) PutStandbyPolicy(value *ComputeRegionInstanceGroupManagerStandbyPolicy) {
	if err := c.validatePutStandbyPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putStandbyPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) PutStatefulDisk(value interface{}) {
	if err := c.validatePutStatefulDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putStatefulDisk",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) PutStatefulExternalIp(value interface{}) {
	if err := c.validatePutStatefulExternalIpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putStatefulExternalIp",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) PutStatefulInternalIp(value interface{}) {
	if err := c.validatePutStatefulInternalIpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putStatefulInternalIp",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) PutTimeouts(value *ComputeRegionInstanceGroupManagerTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) PutUpdatePolicy(value *ComputeRegionInstanceGroupManagerUpdatePolicy) {
	if err := c.validatePutUpdatePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putUpdatePolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) PutVersion(value interface{}) {
	if err := c.validatePutVersionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putVersion",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) ResetAllInstancesConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetAllInstancesConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) ResetAutoHealingPolicies() {
	_jsii_.InvokeVoid(
		c,
		"resetAutoHealingPolicies",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) ResetDistributionPolicyTargetShape() {
	_jsii_.InvokeVoid(
		c,
		"resetDistributionPolicyTargetShape",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) ResetDistributionPolicyZones() {
	_jsii_.InvokeVoid(
		c,
		"resetDistributionPolicyZones",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) ResetInstanceFlexibilityPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetInstanceFlexibilityPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) ResetInstanceLifecyclePolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetInstanceLifecyclePolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) ResetListManagedInstancesResults() {
	_jsii_.InvokeVoid(
		c,
		"resetListManagedInstancesResults",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) ResetNamedPort() {
	_jsii_.InvokeVoid(
		c,
		"resetNamedPort",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) ResetProject() {
	_jsii_.InvokeVoid(
		c,
		"resetProject",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) ResetRegion() {
	_jsii_.InvokeVoid(
		c,
		"resetRegion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) ResetStandbyPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetStandbyPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) ResetStatefulDisk() {
	_jsii_.InvokeVoid(
		c,
		"resetStatefulDisk",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) ResetStatefulExternalIp() {
	_jsii_.InvokeVoid(
		c,
		"resetStatefulExternalIp",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) ResetStatefulInternalIp() {
	_jsii_.InvokeVoid(
		c,
		"resetStatefulInternalIp",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) ResetTargetPools() {
	_jsii_.InvokeVoid(
		c,
		"resetTargetPools",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) ResetTargetSize() {
	_jsii_.InvokeVoid(
		c,
		"resetTargetSize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) ResetTargetStoppedSize() {
	_jsii_.InvokeVoid(
		c,
		"resetTargetStoppedSize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) ResetTargetSuspendedSize() {
	_jsii_.InvokeVoid(
		c,
		"resetTargetSuspendedSize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) ResetUpdatePolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetUpdatePolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) ResetWaitForInstances() {
	_jsii_.InvokeVoid(
		c,
		"resetWaitForInstances",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) ResetWaitForInstancesStatus() {
	_jsii_.InvokeVoid(
		c,
		"resetWaitForInstancesStatus",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionInstanceGroupManager) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

