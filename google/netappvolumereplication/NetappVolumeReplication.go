// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package netappvolumereplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v16/netappvolumereplication/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/netapp_volume_replication google_netapp_volume_replication}.
type NetappVolumeReplication interface {
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
	DeleteDestinationVolume() interface{}
	SetDeleteDestinationVolume(val interface{})
	DeleteDestinationVolumeInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DestinationVolume() *string
	DestinationVolumeParameters() NetappVolumeReplicationDestinationVolumeParametersOutputReference
	DestinationVolumeParametersInput() *NetappVolumeReplicationDestinationVolumeParameters
	EffectiveLabels() cdktf.StringMap
	ForceStopping() interface{}
	SetForceStopping(val interface{})
	ForceStoppingInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Healthy() cdktf.IResolvable
	HybridPeeringDetails() NetappVolumeReplicationHybridPeeringDetailsList
	HybridReplicationType() *string
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
	MirrorState() *string
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
	ReplicationEnabled() interface{}
	SetReplicationEnabled(val interface{})
	ReplicationEnabledInput() interface{}
	ReplicationSchedule() *string
	SetReplicationSchedule(val *string)
	ReplicationScheduleInput() *string
	Role() *string
	SourceVolume() *string
	State() *string
	StateDetails() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() NetappVolumeReplicationTimeoutsOutputReference
	TimeoutsInput() interface{}
	TransferStats() NetappVolumeReplicationTransferStatsList
	VolumeName() *string
	SetVolumeName(val *string)
	VolumeNameInput() *string
	WaitForMirror() interface{}
	SetWaitForMirror(val interface{})
	WaitForMirrorInput() interface{}
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
	PutDestinationVolumeParameters(value *NetappVolumeReplicationDestinationVolumeParameters)
	PutTimeouts(value *NetappVolumeReplicationTimeouts)
	ResetDeleteDestinationVolume()
	ResetDescription()
	ResetDestinationVolumeParameters()
	ResetForceStopping()
	ResetId()
	ResetLabels()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetReplicationEnabled()
	ResetTimeouts()
	ResetWaitForMirror()
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

// The jsii proxy struct for NetappVolumeReplication
type jsiiProxy_NetappVolumeReplication struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_NetappVolumeReplication) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) DeleteDestinationVolume() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteDestinationVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) DeleteDestinationVolumeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteDestinationVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) DestinationVolume() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) DestinationVolumeParameters() NetappVolumeReplicationDestinationVolumeParametersOutputReference {
	var returns NetappVolumeReplicationDestinationVolumeParametersOutputReference
	_jsii_.Get(
		j,
		"destinationVolumeParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) DestinationVolumeParametersInput() *NetappVolumeReplicationDestinationVolumeParameters {
	var returns *NetappVolumeReplicationDestinationVolumeParameters
	_jsii_.Get(
		j,
		"destinationVolumeParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) ForceStopping() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceStopping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) ForceStoppingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceStoppingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) Healthy() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"healthy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) HybridPeeringDetails() NetappVolumeReplicationHybridPeeringDetailsList {
	var returns NetappVolumeReplicationHybridPeeringDetailsList
	_jsii_.Get(
		j,
		"hybridPeeringDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) HybridReplicationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hybridReplicationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) MirrorState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mirrorState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) ReplicationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replicationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) ReplicationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replicationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) ReplicationSchedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) ReplicationScheduleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) SourceVolume() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) StateDetails() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) Timeouts() NetappVolumeReplicationTimeoutsOutputReference {
	var returns NetappVolumeReplicationTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) TransferStats() NetappVolumeReplicationTransferStatsList {
	var returns NetappVolumeReplicationTransferStatsList
	_jsii_.Get(
		j,
		"transferStats",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) VolumeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) VolumeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) WaitForMirror() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForMirror",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeReplication) WaitForMirrorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForMirrorInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/netapp_volume_replication google_netapp_volume_replication} Resource.
func NewNetappVolumeReplication(scope constructs.Construct, id *string, config *NetappVolumeReplicationConfig) NetappVolumeReplication {
	_init_.Initialize()

	if err := validateNewNetappVolumeReplicationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetappVolumeReplication{}

	_jsii_.Create(
		"@cdktf/provider-google.netappVolumeReplication.NetappVolumeReplication",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/netapp_volume_replication google_netapp_volume_replication} Resource.
func NewNetappVolumeReplication_Override(n NetappVolumeReplication, scope constructs.Construct, id *string, config *NetappVolumeReplicationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.netappVolumeReplication.NetappVolumeReplication",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NetappVolumeReplication)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeReplication)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeReplication)SetDeleteDestinationVolume(val interface{}) {
	if err := j.validateSetDeleteDestinationVolumeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteDestinationVolume",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeReplication)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeReplication)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeReplication)SetForceStopping(val interface{}) {
	if err := j.validateSetForceStoppingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceStopping",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeReplication)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeReplication)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeReplication)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeReplication)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeReplication)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeReplication)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeReplication)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeReplication)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeReplication)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeReplication)SetReplicationEnabled(val interface{}) {
	if err := j.validateSetReplicationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationEnabled",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeReplication)SetReplicationSchedule(val *string) {
	if err := j.validateSetReplicationScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationSchedule",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeReplication)SetVolumeName(val *string) {
	if err := j.validateSetVolumeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeName",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeReplication)SetWaitForMirror(val interface{}) {
	if err := j.validateSetWaitForMirrorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitForMirror",
		val,
	)
}

// Generates CDKTF code for importing a NetappVolumeReplication resource upon running "cdktf plan <stack-name>".
func NetappVolumeReplication_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateNetappVolumeReplication_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.netappVolumeReplication.NetappVolumeReplication",
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
func NetappVolumeReplication_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetappVolumeReplication_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.netappVolumeReplication.NetappVolumeReplication",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetappVolumeReplication_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetappVolumeReplication_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.netappVolumeReplication.NetappVolumeReplication",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NetappVolumeReplication_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNetappVolumeReplication_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.netappVolumeReplication.NetappVolumeReplication",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NetappVolumeReplication_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.netappVolumeReplication.NetappVolumeReplication",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NetappVolumeReplication) AddMoveTarget(moveTarget *string) {
	if err := n.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (n *jsiiProxy_NetappVolumeReplication) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NetappVolumeReplication) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NetappVolumeReplication) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetappVolumeReplication) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NetappVolumeReplication) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NetappVolumeReplication) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NetappVolumeReplication) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NetappVolumeReplication) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NetappVolumeReplication) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NetappVolumeReplication) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NetappVolumeReplication) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeReplication) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := n.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (n *jsiiProxy_NetappVolumeReplication) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetappVolumeReplication) MoveFromId(id *string) {
	if err := n.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveFromId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NetappVolumeReplication) MoveTo(moveTarget *string, index interface{}) {
	if err := n.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (n *jsiiProxy_NetappVolumeReplication) MoveToId(id *string) {
	if err := n.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveToId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NetappVolumeReplication) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NetappVolumeReplication) PutDestinationVolumeParameters(value *NetappVolumeReplicationDestinationVolumeParameters) {
	if err := n.validatePutDestinationVolumeParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putDestinationVolumeParameters",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetappVolumeReplication) PutTimeouts(value *NetappVolumeReplicationTimeouts) {
	if err := n.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NetappVolumeReplication) ResetDeleteDestinationVolume() {
	_jsii_.InvokeVoid(
		n,
		"resetDeleteDestinationVolume",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeReplication) ResetDescription() {
	_jsii_.InvokeVoid(
		n,
		"resetDescription",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeReplication) ResetDestinationVolumeParameters() {
	_jsii_.InvokeVoid(
		n,
		"resetDestinationVolumeParameters",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeReplication) ResetForceStopping() {
	_jsii_.InvokeVoid(
		n,
		"resetForceStopping",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeReplication) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeReplication) ResetLabels() {
	_jsii_.InvokeVoid(
		n,
		"resetLabels",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeReplication) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeReplication) ResetProject() {
	_jsii_.InvokeVoid(
		n,
		"resetProject",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeReplication) ResetReplicationEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetReplicationEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeReplication) ResetTimeouts() {
	_jsii_.InvokeVoid(
		n,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeReplication) ResetWaitForMirror() {
	_jsii_.InvokeVoid(
		n,
		"resetWaitForMirror",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeReplication) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeReplication) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeReplication) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeReplication) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeReplication) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeReplication) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

