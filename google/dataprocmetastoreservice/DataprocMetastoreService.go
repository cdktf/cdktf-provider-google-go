// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocmetastoreservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v16/dataprocmetastoreservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/dataproc_metastore_service google_dataproc_metastore_service}.
type DataprocMetastoreService interface {
	cdktf.TerraformResource
	ArtifactGcsUri() *string
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
	DatabaseType() *string
	SetDatabaseType(val *string)
	DatabaseTypeInput() *string
	DeletionProtection() interface{}
	SetDeletionProtection(val interface{})
	DeletionProtectionInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EffectiveLabels() cdktf.StringMap
	EncryptionConfig() DataprocMetastoreServiceEncryptionConfigOutputReference
	EncryptionConfigInput() *DataprocMetastoreServiceEncryptionConfig
	EndpointUri() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HiveMetastoreConfig() DataprocMetastoreServiceHiveMetastoreConfigOutputReference
	HiveMetastoreConfigInput() *DataprocMetastoreServiceHiveMetastoreConfig
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
	MaintenanceWindow() DataprocMetastoreServiceMaintenanceWindowOutputReference
	MaintenanceWindowInput() *DataprocMetastoreServiceMaintenanceWindow
	MetadataIntegration() DataprocMetastoreServiceMetadataIntegrationOutputReference
	MetadataIntegrationInput() *DataprocMetastoreServiceMetadataIntegration
	Name() *string
	Network() *string
	SetNetwork(val *string)
	NetworkConfig() DataprocMetastoreServiceNetworkConfigOutputReference
	NetworkConfigInput() *DataprocMetastoreServiceNetworkConfig
	NetworkInput() *string
	// The tree node.
	Node() constructs.Node
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
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
	ScalingConfig() DataprocMetastoreServiceScalingConfigOutputReference
	ScalingConfigInput() *DataprocMetastoreServiceScalingConfig
	ScheduledBackup() DataprocMetastoreServiceScheduledBackupOutputReference
	ScheduledBackupInput() *DataprocMetastoreServiceScheduledBackup
	ServiceId() *string
	SetServiceId(val *string)
	ServiceIdInput() *string
	State() *string
	StateMessage() *string
	TelemetryConfig() DataprocMetastoreServiceTelemetryConfigOutputReference
	TelemetryConfigInput() *DataprocMetastoreServiceTelemetryConfig
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Tier() *string
	SetTier(val *string)
	TierInput() *string
	Timeouts() DataprocMetastoreServiceTimeoutsOutputReference
	TimeoutsInput() interface{}
	Uid() *string
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
	PutEncryptionConfig(value *DataprocMetastoreServiceEncryptionConfig)
	PutHiveMetastoreConfig(value *DataprocMetastoreServiceHiveMetastoreConfig)
	PutMaintenanceWindow(value *DataprocMetastoreServiceMaintenanceWindow)
	PutMetadataIntegration(value *DataprocMetastoreServiceMetadataIntegration)
	PutNetworkConfig(value *DataprocMetastoreServiceNetworkConfig)
	PutScalingConfig(value *DataprocMetastoreServiceScalingConfig)
	PutScheduledBackup(value *DataprocMetastoreServiceScheduledBackup)
	PutTelemetryConfig(value *DataprocMetastoreServiceTelemetryConfig)
	PutTimeouts(value *DataprocMetastoreServiceTimeouts)
	ResetDatabaseType()
	ResetDeletionProtection()
	ResetEncryptionConfig()
	ResetHiveMetastoreConfig()
	ResetId()
	ResetLabels()
	ResetLocation()
	ResetMaintenanceWindow()
	ResetMetadataIntegration()
	ResetNetwork()
	ResetNetworkConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPort()
	ResetProject()
	ResetReleaseChannel()
	ResetScalingConfig()
	ResetScheduledBackup()
	ResetTelemetryConfig()
	ResetTier()
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

// The jsii proxy struct for DataprocMetastoreService
type jsiiProxy_DataprocMetastoreService struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DataprocMetastoreService) ArtifactGcsUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactGcsUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) DatabaseType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) DatabaseTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) DeletionProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) DeletionProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) EncryptionConfig() DataprocMetastoreServiceEncryptionConfigOutputReference {
	var returns DataprocMetastoreServiceEncryptionConfigOutputReference
	_jsii_.Get(
		j,
		"encryptionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) EncryptionConfigInput() *DataprocMetastoreServiceEncryptionConfig {
	var returns *DataprocMetastoreServiceEncryptionConfig
	_jsii_.Get(
		j,
		"encryptionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) EndpointUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) HiveMetastoreConfig() DataprocMetastoreServiceHiveMetastoreConfigOutputReference {
	var returns DataprocMetastoreServiceHiveMetastoreConfigOutputReference
	_jsii_.Get(
		j,
		"hiveMetastoreConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) HiveMetastoreConfigInput() *DataprocMetastoreServiceHiveMetastoreConfig {
	var returns *DataprocMetastoreServiceHiveMetastoreConfig
	_jsii_.Get(
		j,
		"hiveMetastoreConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) MaintenanceWindow() DataprocMetastoreServiceMaintenanceWindowOutputReference {
	var returns DataprocMetastoreServiceMaintenanceWindowOutputReference
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) MaintenanceWindowInput() *DataprocMetastoreServiceMaintenanceWindow {
	var returns *DataprocMetastoreServiceMaintenanceWindow
	_jsii_.Get(
		j,
		"maintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) MetadataIntegration() DataprocMetastoreServiceMetadataIntegrationOutputReference {
	var returns DataprocMetastoreServiceMetadataIntegrationOutputReference
	_jsii_.Get(
		j,
		"metadataIntegration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) MetadataIntegrationInput() *DataprocMetastoreServiceMetadataIntegration {
	var returns *DataprocMetastoreServiceMetadataIntegration
	_jsii_.Get(
		j,
		"metadataIntegrationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) NetworkConfig() DataprocMetastoreServiceNetworkConfigOutputReference {
	var returns DataprocMetastoreServiceNetworkConfigOutputReference
	_jsii_.Get(
		j,
		"networkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) NetworkConfigInput() *DataprocMetastoreServiceNetworkConfig {
	var returns *DataprocMetastoreServiceNetworkConfig
	_jsii_.Get(
		j,
		"networkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) ReleaseChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) ReleaseChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) ScalingConfig() DataprocMetastoreServiceScalingConfigOutputReference {
	var returns DataprocMetastoreServiceScalingConfigOutputReference
	_jsii_.Get(
		j,
		"scalingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) ScalingConfigInput() *DataprocMetastoreServiceScalingConfig {
	var returns *DataprocMetastoreServiceScalingConfig
	_jsii_.Get(
		j,
		"scalingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) ScheduledBackup() DataprocMetastoreServiceScheduledBackupOutputReference {
	var returns DataprocMetastoreServiceScheduledBackupOutputReference
	_jsii_.Get(
		j,
		"scheduledBackup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) ScheduledBackupInput() *DataprocMetastoreServiceScheduledBackup {
	var returns *DataprocMetastoreServiceScheduledBackup
	_jsii_.Get(
		j,
		"scheduledBackupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) ServiceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) ServiceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) StateMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) TelemetryConfig() DataprocMetastoreServiceTelemetryConfigOutputReference {
	var returns DataprocMetastoreServiceTelemetryConfigOutputReference
	_jsii_.Get(
		j,
		"telemetryConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) TelemetryConfigInput() *DataprocMetastoreServiceTelemetryConfig {
	var returns *DataprocMetastoreServiceTelemetryConfig
	_jsii_.Get(
		j,
		"telemetryConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) Tier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) TierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) Timeouts() DataprocMetastoreServiceTimeoutsOutputReference {
	var returns DataprocMetastoreServiceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) Uid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocMetastoreService) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/dataproc_metastore_service google_dataproc_metastore_service} Resource.
func NewDataprocMetastoreService(scope constructs.Construct, id *string, config *DataprocMetastoreServiceConfig) DataprocMetastoreService {
	_init_.Initialize()

	if err := validateNewDataprocMetastoreServiceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataprocMetastoreService{}

	_jsii_.Create(
		"@cdktf/provider-google.dataprocMetastoreService.DataprocMetastoreService",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/dataproc_metastore_service google_dataproc_metastore_service} Resource.
func NewDataprocMetastoreService_Override(d DataprocMetastoreService, scope constructs.Construct, id *string, config *DataprocMetastoreServiceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataprocMetastoreService.DataprocMetastoreService",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataprocMetastoreService)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DataprocMetastoreService)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataprocMetastoreService)SetDatabaseType(val *string) {
	if err := j.validateSetDatabaseTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseType",
		val,
	)
}

func (j *jsiiProxy_DataprocMetastoreService)SetDeletionProtection(val interface{}) {
	if err := j.validateSetDeletionProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_DataprocMetastoreService)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataprocMetastoreService)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataprocMetastoreService)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataprocMetastoreService)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_DataprocMetastoreService)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataprocMetastoreService)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_DataprocMetastoreService)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_DataprocMetastoreService)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_DataprocMetastoreService)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_DataprocMetastoreService)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataprocMetastoreService)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DataprocMetastoreService)SetReleaseChannel(val *string) {
	if err := j.validateSetReleaseChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"releaseChannel",
		val,
	)
}

func (j *jsiiProxy_DataprocMetastoreService)SetServiceId(val *string) {
	if err := j.validateSetServiceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceId",
		val,
	)
}

func (j *jsiiProxy_DataprocMetastoreService)SetTier(val *string) {
	if err := j.validateSetTierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tier",
		val,
	)
}

// Generates CDKTF code for importing a DataprocMetastoreService resource upon running "cdktf plan <stack-name>".
func DataprocMetastoreService_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataprocMetastoreService_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dataprocMetastoreService.DataprocMetastoreService",
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
func DataprocMetastoreService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataprocMetastoreService_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dataprocMetastoreService.DataprocMetastoreService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataprocMetastoreService_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataprocMetastoreService_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dataprocMetastoreService.DataprocMetastoreService",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataprocMetastoreService_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataprocMetastoreService_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dataprocMetastoreService.DataprocMetastoreService",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataprocMetastoreService_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.dataprocMetastoreService.DataprocMetastoreService",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataprocMetastoreService) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DataprocMetastoreService) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataprocMetastoreService) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocMetastoreService) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocMetastoreService) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocMetastoreService) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocMetastoreService) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocMetastoreService) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocMetastoreService) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocMetastoreService) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocMetastoreService) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocMetastoreService) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocMetastoreService) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DataprocMetastoreService) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocMetastoreService) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DataprocMetastoreService) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DataprocMetastoreService) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DataprocMetastoreService) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataprocMetastoreService) PutEncryptionConfig(value *DataprocMetastoreServiceEncryptionConfig) {
	if err := d.validatePutEncryptionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putEncryptionConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocMetastoreService) PutHiveMetastoreConfig(value *DataprocMetastoreServiceHiveMetastoreConfig) {
	if err := d.validatePutHiveMetastoreConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putHiveMetastoreConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocMetastoreService) PutMaintenanceWindow(value *DataprocMetastoreServiceMaintenanceWindow) {
	if err := d.validatePutMaintenanceWindowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMaintenanceWindow",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocMetastoreService) PutMetadataIntegration(value *DataprocMetastoreServiceMetadataIntegration) {
	if err := d.validatePutMetadataIntegrationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMetadataIntegration",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocMetastoreService) PutNetworkConfig(value *DataprocMetastoreServiceNetworkConfig) {
	if err := d.validatePutNetworkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNetworkConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocMetastoreService) PutScalingConfig(value *DataprocMetastoreServiceScalingConfig) {
	if err := d.validatePutScalingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putScalingConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocMetastoreService) PutScheduledBackup(value *DataprocMetastoreServiceScheduledBackup) {
	if err := d.validatePutScheduledBackupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putScheduledBackup",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocMetastoreService) PutTelemetryConfig(value *DataprocMetastoreServiceTelemetryConfig) {
	if err := d.validatePutTelemetryConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTelemetryConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocMetastoreService) PutTimeouts(value *DataprocMetastoreServiceTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataprocMetastoreService) ResetDatabaseType() {
	_jsii_.InvokeVoid(
		d,
		"resetDatabaseType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocMetastoreService) ResetDeletionProtection() {
	_jsii_.InvokeVoid(
		d,
		"resetDeletionProtection",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocMetastoreService) ResetEncryptionConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetEncryptionConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocMetastoreService) ResetHiveMetastoreConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetHiveMetastoreConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocMetastoreService) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocMetastoreService) ResetLabels() {
	_jsii_.InvokeVoid(
		d,
		"resetLabels",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocMetastoreService) ResetLocation() {
	_jsii_.InvokeVoid(
		d,
		"resetLocation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocMetastoreService) ResetMaintenanceWindow() {
	_jsii_.InvokeVoid(
		d,
		"resetMaintenanceWindow",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocMetastoreService) ResetMetadataIntegration() {
	_jsii_.InvokeVoid(
		d,
		"resetMetadataIntegration",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocMetastoreService) ResetNetwork() {
	_jsii_.InvokeVoid(
		d,
		"resetNetwork",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocMetastoreService) ResetNetworkConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetNetworkConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocMetastoreService) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocMetastoreService) ResetPort() {
	_jsii_.InvokeVoid(
		d,
		"resetPort",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocMetastoreService) ResetProject() {
	_jsii_.InvokeVoid(
		d,
		"resetProject",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocMetastoreService) ResetReleaseChannel() {
	_jsii_.InvokeVoid(
		d,
		"resetReleaseChannel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocMetastoreService) ResetScalingConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetScalingConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocMetastoreService) ResetScheduledBackup() {
	_jsii_.InvokeVoid(
		d,
		"resetScheduledBackup",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocMetastoreService) ResetTelemetryConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetTelemetryConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocMetastoreService) ResetTier() {
	_jsii_.InvokeVoid(
		d,
		"resetTier",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocMetastoreService) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocMetastoreService) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocMetastoreService) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocMetastoreService) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocMetastoreService) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocMetastoreService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocMetastoreService) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

