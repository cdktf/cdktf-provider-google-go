// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alloydbcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v14/alloydbcluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.19.0/docs/resources/alloydb_cluster google_alloydb_cluster}.
type AlloydbCluster interface {
	cdktf.TerraformResource
	Annotations() *map[string]*string
	SetAnnotations(val *map[string]*string)
	AnnotationsInput() *map[string]*string
	AutomatedBackupPolicy() AlloydbClusterAutomatedBackupPolicyOutputReference
	AutomatedBackupPolicyInput() *AlloydbClusterAutomatedBackupPolicy
	BackupSource() AlloydbClusterBackupSourceList
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterId() *string
	SetClusterId(val *string)
	ClusterIdInput() *string
	ClusterType() *string
	SetClusterType(val *string)
	ClusterTypeInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContinuousBackupConfig() AlloydbClusterContinuousBackupConfigOutputReference
	ContinuousBackupConfigInput() *AlloydbClusterContinuousBackupConfig
	ContinuousBackupInfo() AlloydbClusterContinuousBackupInfoList
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DatabaseVersion() *string
	SetDatabaseVersion(val *string)
	DatabaseVersionInput() *string
	DeletionPolicy() *string
	SetDeletionPolicy(val *string)
	DeletionPolicyInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EffectiveAnnotations() cdktf.StringMap
	EffectiveLabels() cdktf.StringMap
	EncryptionConfig() AlloydbClusterEncryptionConfigOutputReference
	EncryptionConfigInput() *AlloydbClusterEncryptionConfig
	EncryptionInfo() AlloydbClusterEncryptionInfoList
	Etag() *string
	SetEtag(val *string)
	EtagInput() *string
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
	InitialUser() AlloydbClusterInitialUserOutputReference
	InitialUserInput() *AlloydbClusterInitialUser
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
	MaintenanceUpdatePolicy() AlloydbClusterMaintenanceUpdatePolicyOutputReference
	MaintenanceUpdatePolicyInput() *AlloydbClusterMaintenanceUpdatePolicy
	MigrationSource() AlloydbClusterMigrationSourceList
	Name() *string
	NetworkConfig() AlloydbClusterNetworkConfigOutputReference
	NetworkConfigInput() *AlloydbClusterNetworkConfig
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
	PscConfig() AlloydbClusterPscConfigOutputReference
	PscConfigInput() *AlloydbClusterPscConfig
	// Experimental.
	RawOverrides() interface{}
	Reconciling() cdktf.IResolvable
	RestoreBackupSource() AlloydbClusterRestoreBackupSourceOutputReference
	RestoreBackupSourceInput() *AlloydbClusterRestoreBackupSource
	RestoreContinuousBackupSource() AlloydbClusterRestoreContinuousBackupSourceOutputReference
	RestoreContinuousBackupSourceInput() *AlloydbClusterRestoreContinuousBackupSource
	SecondaryConfig() AlloydbClusterSecondaryConfigOutputReference
	SecondaryConfigInput() *AlloydbClusterSecondaryConfig
	SkipAwaitMajorVersionUpgrade() interface{}
	SetSkipAwaitMajorVersionUpgrade(val interface{})
	SkipAwaitMajorVersionUpgradeInput() interface{}
	State() *string
	SubscriptionType() *string
	SetSubscriptionType(val *string)
	SubscriptionTypeInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() AlloydbClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
	TrialMetadata() AlloydbClusterTrialMetadataList
	Uid() *string
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
	PutAutomatedBackupPolicy(value *AlloydbClusterAutomatedBackupPolicy)
	PutContinuousBackupConfig(value *AlloydbClusterContinuousBackupConfig)
	PutEncryptionConfig(value *AlloydbClusterEncryptionConfig)
	PutInitialUser(value *AlloydbClusterInitialUser)
	PutMaintenanceUpdatePolicy(value *AlloydbClusterMaintenanceUpdatePolicy)
	PutNetworkConfig(value *AlloydbClusterNetworkConfig)
	PutPscConfig(value *AlloydbClusterPscConfig)
	PutRestoreBackupSource(value *AlloydbClusterRestoreBackupSource)
	PutRestoreContinuousBackupSource(value *AlloydbClusterRestoreContinuousBackupSource)
	PutSecondaryConfig(value *AlloydbClusterSecondaryConfig)
	PutTimeouts(value *AlloydbClusterTimeouts)
	ResetAnnotations()
	ResetAutomatedBackupPolicy()
	ResetClusterType()
	ResetContinuousBackupConfig()
	ResetDatabaseVersion()
	ResetDeletionPolicy()
	ResetDisplayName()
	ResetEncryptionConfig()
	ResetEtag()
	ResetId()
	ResetInitialUser()
	ResetLabels()
	ResetMaintenanceUpdatePolicy()
	ResetNetworkConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetPscConfig()
	ResetRestoreBackupSource()
	ResetRestoreContinuousBackupSource()
	ResetSecondaryConfig()
	ResetSkipAwaitMajorVersionUpgrade()
	ResetSubscriptionType()
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

// The jsii proxy struct for AlloydbCluster
type jsiiProxy_AlloydbCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AlloydbCluster) Annotations() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) AnnotationsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) AutomatedBackupPolicy() AlloydbClusterAutomatedBackupPolicyOutputReference {
	var returns AlloydbClusterAutomatedBackupPolicyOutputReference
	_jsii_.Get(
		j,
		"automatedBackupPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) AutomatedBackupPolicyInput() *AlloydbClusterAutomatedBackupPolicy {
	var returns *AlloydbClusterAutomatedBackupPolicy
	_jsii_.Get(
		j,
		"automatedBackupPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) BackupSource() AlloydbClusterBackupSourceList {
	var returns AlloydbClusterBackupSourceList
	_jsii_.Get(
		j,
		"backupSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) ClusterType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) ClusterTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) ContinuousBackupConfig() AlloydbClusterContinuousBackupConfigOutputReference {
	var returns AlloydbClusterContinuousBackupConfigOutputReference
	_jsii_.Get(
		j,
		"continuousBackupConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) ContinuousBackupConfigInput() *AlloydbClusterContinuousBackupConfig {
	var returns *AlloydbClusterContinuousBackupConfig
	_jsii_.Get(
		j,
		"continuousBackupConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) ContinuousBackupInfo() AlloydbClusterContinuousBackupInfoList {
	var returns AlloydbClusterContinuousBackupInfoList
	_jsii_.Get(
		j,
		"continuousBackupInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) DatabaseVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) DatabaseVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) DeletionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) DeletionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) EffectiveAnnotations() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) EncryptionConfig() AlloydbClusterEncryptionConfigOutputReference {
	var returns AlloydbClusterEncryptionConfigOutputReference
	_jsii_.Get(
		j,
		"encryptionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) EncryptionConfigInput() *AlloydbClusterEncryptionConfig {
	var returns *AlloydbClusterEncryptionConfig
	_jsii_.Get(
		j,
		"encryptionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) EncryptionInfo() AlloydbClusterEncryptionInfoList {
	var returns AlloydbClusterEncryptionInfoList
	_jsii_.Get(
		j,
		"encryptionInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) EtagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) InitialUser() AlloydbClusterInitialUserOutputReference {
	var returns AlloydbClusterInitialUserOutputReference
	_jsii_.Get(
		j,
		"initialUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) InitialUserInput() *AlloydbClusterInitialUser {
	var returns *AlloydbClusterInitialUser
	_jsii_.Get(
		j,
		"initialUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) MaintenanceUpdatePolicy() AlloydbClusterMaintenanceUpdatePolicyOutputReference {
	var returns AlloydbClusterMaintenanceUpdatePolicyOutputReference
	_jsii_.Get(
		j,
		"maintenanceUpdatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) MaintenanceUpdatePolicyInput() *AlloydbClusterMaintenanceUpdatePolicy {
	var returns *AlloydbClusterMaintenanceUpdatePolicy
	_jsii_.Get(
		j,
		"maintenanceUpdatePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) MigrationSource() AlloydbClusterMigrationSourceList {
	var returns AlloydbClusterMigrationSourceList
	_jsii_.Get(
		j,
		"migrationSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) NetworkConfig() AlloydbClusterNetworkConfigOutputReference {
	var returns AlloydbClusterNetworkConfigOutputReference
	_jsii_.Get(
		j,
		"networkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) NetworkConfigInput() *AlloydbClusterNetworkConfig {
	var returns *AlloydbClusterNetworkConfig
	_jsii_.Get(
		j,
		"networkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) PscConfig() AlloydbClusterPscConfigOutputReference {
	var returns AlloydbClusterPscConfigOutputReference
	_jsii_.Get(
		j,
		"pscConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) PscConfigInput() *AlloydbClusterPscConfig {
	var returns *AlloydbClusterPscConfig
	_jsii_.Get(
		j,
		"pscConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) Reconciling() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"reconciling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) RestoreBackupSource() AlloydbClusterRestoreBackupSourceOutputReference {
	var returns AlloydbClusterRestoreBackupSourceOutputReference
	_jsii_.Get(
		j,
		"restoreBackupSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) RestoreBackupSourceInput() *AlloydbClusterRestoreBackupSource {
	var returns *AlloydbClusterRestoreBackupSource
	_jsii_.Get(
		j,
		"restoreBackupSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) RestoreContinuousBackupSource() AlloydbClusterRestoreContinuousBackupSourceOutputReference {
	var returns AlloydbClusterRestoreContinuousBackupSourceOutputReference
	_jsii_.Get(
		j,
		"restoreContinuousBackupSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) RestoreContinuousBackupSourceInput() *AlloydbClusterRestoreContinuousBackupSource {
	var returns *AlloydbClusterRestoreContinuousBackupSource
	_jsii_.Get(
		j,
		"restoreContinuousBackupSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) SecondaryConfig() AlloydbClusterSecondaryConfigOutputReference {
	var returns AlloydbClusterSecondaryConfigOutputReference
	_jsii_.Get(
		j,
		"secondaryConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) SecondaryConfigInput() *AlloydbClusterSecondaryConfig {
	var returns *AlloydbClusterSecondaryConfig
	_jsii_.Get(
		j,
		"secondaryConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) SkipAwaitMajorVersionUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipAwaitMajorVersionUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) SkipAwaitMajorVersionUpgradeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipAwaitMajorVersionUpgradeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) SubscriptionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) SubscriptionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) Timeouts() AlloydbClusterTimeoutsOutputReference {
	var returns AlloydbClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) TrialMetadata() AlloydbClusterTrialMetadataList {
	var returns AlloydbClusterTrialMetadataList
	_jsii_.Get(
		j,
		"trialMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbCluster) Uid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.19.0/docs/resources/alloydb_cluster google_alloydb_cluster} Resource.
func NewAlloydbCluster(scope constructs.Construct, id *string, config *AlloydbClusterConfig) AlloydbCluster {
	_init_.Initialize()

	if err := validateNewAlloydbClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlloydbCluster{}

	_jsii_.Create(
		"@cdktf/provider-google.alloydbCluster.AlloydbCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.19.0/docs/resources/alloydb_cluster google_alloydb_cluster} Resource.
func NewAlloydbCluster_Override(a AlloydbCluster, scope constructs.Construct, id *string, config *AlloydbClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.alloydbCluster.AlloydbCluster",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AlloydbCluster)SetAnnotations(val *map[string]*string) {
	if err := j.validateSetAnnotationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotations",
		val,
	)
}

func (j *jsiiProxy_AlloydbCluster)SetClusterId(val *string) {
	if err := j.validateSetClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_AlloydbCluster)SetClusterType(val *string) {
	if err := j.validateSetClusterTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterType",
		val,
	)
}

func (j *jsiiProxy_AlloydbCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AlloydbCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AlloydbCluster)SetDatabaseVersion(val *string) {
	if err := j.validateSetDatabaseVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseVersion",
		val,
	)
}

func (j *jsiiProxy_AlloydbCluster)SetDeletionPolicy(val *string) {
	if err := j.validateSetDeletionPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionPolicy",
		val,
	)
}

func (j *jsiiProxy_AlloydbCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AlloydbCluster)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_AlloydbCluster)SetEtag(val *string) {
	if err := j.validateSetEtagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"etag",
		val,
	)
}

func (j *jsiiProxy_AlloydbCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AlloydbCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AlloydbCluster)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_AlloydbCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AlloydbCluster)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_AlloydbCluster)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_AlloydbCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AlloydbCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AlloydbCluster)SetSkipAwaitMajorVersionUpgrade(val interface{}) {
	if err := j.validateSetSkipAwaitMajorVersionUpgradeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipAwaitMajorVersionUpgrade",
		val,
	)
}

func (j *jsiiProxy_AlloydbCluster)SetSubscriptionType(val *string) {
	if err := j.validateSetSubscriptionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subscriptionType",
		val,
	)
}

// Generates CDKTF code for importing a AlloydbCluster resource upon running "cdktf plan <stack-name>".
func AlloydbCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateAlloydbCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.alloydbCluster.AlloydbCluster",
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
func AlloydbCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAlloydbCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.alloydbCluster.AlloydbCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AlloydbCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAlloydbCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.alloydbCluster.AlloydbCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AlloydbCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAlloydbCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.alloydbCluster.AlloydbCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AlloydbCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.alloydbCluster.AlloydbCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AlloydbCluster) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AlloydbCluster) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AlloydbCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlloydbCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlloydbCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlloydbCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlloydbCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlloydbCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlloydbCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlloydbCluster) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlloydbCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlloydbCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlloydbCluster) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AlloydbCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlloydbCluster) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AlloydbCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AlloydbCluster) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AlloydbCluster) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AlloydbCluster) PutAutomatedBackupPolicy(value *AlloydbClusterAutomatedBackupPolicy) {
	if err := a.validatePutAutomatedBackupPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAutomatedBackupPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlloydbCluster) PutContinuousBackupConfig(value *AlloydbClusterContinuousBackupConfig) {
	if err := a.validatePutContinuousBackupConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putContinuousBackupConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlloydbCluster) PutEncryptionConfig(value *AlloydbClusterEncryptionConfig) {
	if err := a.validatePutEncryptionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEncryptionConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlloydbCluster) PutInitialUser(value *AlloydbClusterInitialUser) {
	if err := a.validatePutInitialUserParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInitialUser",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlloydbCluster) PutMaintenanceUpdatePolicy(value *AlloydbClusterMaintenanceUpdatePolicy) {
	if err := a.validatePutMaintenanceUpdatePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMaintenanceUpdatePolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlloydbCluster) PutNetworkConfig(value *AlloydbClusterNetworkConfig) {
	if err := a.validatePutNetworkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNetworkConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlloydbCluster) PutPscConfig(value *AlloydbClusterPscConfig) {
	if err := a.validatePutPscConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPscConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlloydbCluster) PutRestoreBackupSource(value *AlloydbClusterRestoreBackupSource) {
	if err := a.validatePutRestoreBackupSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRestoreBackupSource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlloydbCluster) PutRestoreContinuousBackupSource(value *AlloydbClusterRestoreContinuousBackupSource) {
	if err := a.validatePutRestoreContinuousBackupSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRestoreContinuousBackupSource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlloydbCluster) PutSecondaryConfig(value *AlloydbClusterSecondaryConfig) {
	if err := a.validatePutSecondaryConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSecondaryConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlloydbCluster) PutTimeouts(value *AlloydbClusterTimeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlloydbCluster) ResetAnnotations() {
	_jsii_.InvokeVoid(
		a,
		"resetAnnotations",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlloydbCluster) ResetAutomatedBackupPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetAutomatedBackupPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlloydbCluster) ResetClusterType() {
	_jsii_.InvokeVoid(
		a,
		"resetClusterType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlloydbCluster) ResetContinuousBackupConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetContinuousBackupConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlloydbCluster) ResetDatabaseVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetDatabaseVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlloydbCluster) ResetDeletionPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetDeletionPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlloydbCluster) ResetDisplayName() {
	_jsii_.InvokeVoid(
		a,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlloydbCluster) ResetEncryptionConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetEncryptionConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlloydbCluster) ResetEtag() {
	_jsii_.InvokeVoid(
		a,
		"resetEtag",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlloydbCluster) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlloydbCluster) ResetInitialUser() {
	_jsii_.InvokeVoid(
		a,
		"resetInitialUser",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlloydbCluster) ResetLabels() {
	_jsii_.InvokeVoid(
		a,
		"resetLabels",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlloydbCluster) ResetMaintenanceUpdatePolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetMaintenanceUpdatePolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlloydbCluster) ResetNetworkConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlloydbCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlloydbCluster) ResetProject() {
	_jsii_.InvokeVoid(
		a,
		"resetProject",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlloydbCluster) ResetPscConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetPscConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlloydbCluster) ResetRestoreBackupSource() {
	_jsii_.InvokeVoid(
		a,
		"resetRestoreBackupSource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlloydbCluster) ResetRestoreContinuousBackupSource() {
	_jsii_.InvokeVoid(
		a,
		"resetRestoreContinuousBackupSource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlloydbCluster) ResetSecondaryConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetSecondaryConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlloydbCluster) ResetSkipAwaitMajorVersionUpgrade() {
	_jsii_.InvokeVoid(
		a,
		"resetSkipAwaitMajorVersionUpgrade",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlloydbCluster) ResetSubscriptionType() {
	_jsii_.InvokeVoid(
		a,
		"resetSubscriptionType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlloydbCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlloydbCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlloydbCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlloydbCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlloydbCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlloydbCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlloydbCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

