// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oracledatabaseautonomousdatabase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/oracledatabaseautonomousdatabase/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OracleDatabaseAutonomousDatabasePropertiesOutputReference interface {
	cdktf.ComplexObject
	ActualUsedDataStorageSizeTb() *float64
	AllocatedStorageSizeTb() *float64
	ApexDetails() OracleDatabaseAutonomousDatabasePropertiesApexDetailsList
	ArePrimaryAllowlistedIpsUsed() cdktf.IResolvable
	AutonomousContainerDatabaseId() *string
	AvailableUpgradeVersions() *[]*string
	BackupRetentionPeriodDays() *float64
	SetBackupRetentionPeriodDays(val *float64)
	BackupRetentionPeriodDaysInput() *float64
	CharacterSet() *string
	SetCharacterSet(val *string)
	CharacterSetInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	ComputeCount() *float64
	SetComputeCount(val *float64)
	ComputeCountInput() *float64
	ConnectionStrings() OracleDatabaseAutonomousDatabasePropertiesConnectionStringsList
	ConnectionUrls() OracleDatabaseAutonomousDatabasePropertiesConnectionUrlsList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomerContacts() OracleDatabaseAutonomousDatabasePropertiesCustomerContactsList
	CustomerContactsInput() interface{}
	DatabaseManagementState() *string
	DataSafeState() *string
	DataStorageSizeGb() *float64
	SetDataStorageSizeGb(val *float64)
	DataStorageSizeGbInput() *float64
	DataStorageSizeTb() *float64
	SetDataStorageSizeTb(val *float64)
	DataStorageSizeTbInput() *float64
	DbEdition() *string
	SetDbEdition(val *string)
	DbEditionInput() *string
	DbVersion() *string
	SetDbVersion(val *string)
	DbVersionInput() *string
	DbWorkload() *string
	SetDbWorkload(val *string)
	DbWorkloadInput() *string
	FailedDataRecoveryDuration() *string
	// Experimental.
	Fqn() *string
	InternalValue() *OracleDatabaseAutonomousDatabaseProperties
	SetInternalValue(val *OracleDatabaseAutonomousDatabaseProperties)
	IsAutoScalingEnabled() interface{}
	SetIsAutoScalingEnabled(val interface{})
	IsAutoScalingEnabledInput() interface{}
	IsLocalDataGuardEnabled() cdktf.IResolvable
	IsStorageAutoScalingEnabled() interface{}
	SetIsStorageAutoScalingEnabled(val interface{})
	IsStorageAutoScalingEnabledInput() interface{}
	LicenseType() *string
	SetLicenseType(val *string)
	LicenseTypeInput() *string
	LifecycleDetails() *string
	LocalAdgAutoFailoverMaxDataLossLimit() *float64
	LocalDisasterRecoveryType() *string
	LocalStandbyDb() OracleDatabaseAutonomousDatabasePropertiesLocalStandbyDbList
	MaintenanceBeginTime() *string
	MaintenanceEndTime() *string
	MaintenanceScheduleType() *string
	SetMaintenanceScheduleType(val *string)
	MaintenanceScheduleTypeInput() *string
	MemoryPerOracleComputeUnitGbs() *float64
	MemoryTableGbs() *float64
	MtlsConnectionRequired() interface{}
	SetMtlsConnectionRequired(val interface{})
	MtlsConnectionRequiredInput() interface{}
	NCharacterSet() *string
	SetNCharacterSet(val *string)
	NCharacterSetInput() *string
	NextLongTermBackupTime() *string
	Ocid() *string
	OciUrl() *string
	OpenMode() *string
	OperationsInsightsState() *string
	SetOperationsInsightsState(val *string)
	OperationsInsightsStateInput() *string
	PeerDbIds() *[]*string
	PermissionLevel() *string
	PrivateEndpoint() *string
	PrivateEndpointIp() *string
	SetPrivateEndpointIp(val *string)
	PrivateEndpointIpInput() *string
	PrivateEndpointLabel() *string
	SetPrivateEndpointLabel(val *string)
	PrivateEndpointLabelInput() *string
	RefreshableMode() *string
	RefreshableState() *string
	Role() *string
	ScheduledOperationDetails() OracleDatabaseAutonomousDatabasePropertiesScheduledOperationDetailsList
	SqlWebDeveloperUrl() *string
	State() *string
	SupportedCloneRegions() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TotalAutoBackupStorageSizeGbs() *float64
	UsedDataStorageSizeTbs() *float64
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutCustomerContacts(value interface{})
	ResetBackupRetentionPeriodDays()
	ResetCharacterSet()
	ResetComputeCount()
	ResetCustomerContacts()
	ResetDataStorageSizeGb()
	ResetDataStorageSizeTb()
	ResetDbEdition()
	ResetDbVersion()
	ResetIsAutoScalingEnabled()
	ResetIsStorageAutoScalingEnabled()
	ResetMaintenanceScheduleType()
	ResetMtlsConnectionRequired()
	ResetNCharacterSet()
	ResetOperationsInsightsState()
	ResetPrivateEndpointIp()
	ResetPrivateEndpointLabel()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OracleDatabaseAutonomousDatabasePropertiesOutputReference
type jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) ActualUsedDataStorageSizeTb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"actualUsedDataStorageSizeTb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) AllocatedStorageSizeTb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocatedStorageSizeTb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) ApexDetails() OracleDatabaseAutonomousDatabasePropertiesApexDetailsList {
	var returns OracleDatabaseAutonomousDatabasePropertiesApexDetailsList
	_jsii_.Get(
		j,
		"apexDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) ArePrimaryAllowlistedIpsUsed() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"arePrimaryAllowlistedIpsUsed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) AutonomousContainerDatabaseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autonomousContainerDatabaseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) AvailableUpgradeVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availableUpgradeVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) BackupRetentionPeriodDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionPeriodDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) BackupRetentionPeriodDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionPeriodDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) CharacterSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"characterSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) CharacterSetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"characterSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) ComputeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) ComputeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) ConnectionStrings() OracleDatabaseAutonomousDatabasePropertiesConnectionStringsList {
	var returns OracleDatabaseAutonomousDatabasePropertiesConnectionStringsList
	_jsii_.Get(
		j,
		"connectionStrings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) ConnectionUrls() OracleDatabaseAutonomousDatabasePropertiesConnectionUrlsList {
	var returns OracleDatabaseAutonomousDatabasePropertiesConnectionUrlsList
	_jsii_.Get(
		j,
		"connectionUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) CustomerContacts() OracleDatabaseAutonomousDatabasePropertiesCustomerContactsList {
	var returns OracleDatabaseAutonomousDatabasePropertiesCustomerContactsList
	_jsii_.Get(
		j,
		"customerContacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) CustomerContactsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customerContactsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) DatabaseManagementState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseManagementState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) DataSafeState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSafeState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) DataStorageSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) DataStorageSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) DataStorageSizeTb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeTb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) DataStorageSizeTbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeTbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) DbEdition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbEdition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) DbEditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbEditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) DbVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) DbVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) DbWorkload() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbWorkload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) DbWorkloadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbWorkloadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) FailedDataRecoveryDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failedDataRecoveryDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) InternalValue() *OracleDatabaseAutonomousDatabaseProperties {
	var returns *OracleDatabaseAutonomousDatabaseProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) IsAutoScalingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isAutoScalingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) IsAutoScalingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isAutoScalingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) IsLocalDataGuardEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isLocalDataGuardEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) IsStorageAutoScalingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isStorageAutoScalingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) IsStorageAutoScalingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isStorageAutoScalingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) LicenseType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) LicenseTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) LifecycleDetails() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) LocalAdgAutoFailoverMaxDataLossLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"localAdgAutoFailoverMaxDataLossLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) LocalDisasterRecoveryType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localDisasterRecoveryType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) LocalStandbyDb() OracleDatabaseAutonomousDatabasePropertiesLocalStandbyDbList {
	var returns OracleDatabaseAutonomousDatabasePropertiesLocalStandbyDbList
	_jsii_.Get(
		j,
		"localStandbyDb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) MaintenanceBeginTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceBeginTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) MaintenanceEndTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceEndTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) MaintenanceScheduleType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceScheduleType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) MaintenanceScheduleTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceScheduleTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) MemoryPerOracleComputeUnitGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryPerOracleComputeUnitGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) MemoryTableGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryTableGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) MtlsConnectionRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mtlsConnectionRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) MtlsConnectionRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mtlsConnectionRequiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) NCharacterSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nCharacterSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) NCharacterSetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nCharacterSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) NextLongTermBackupTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextLongTermBackupTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) Ocid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) OciUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) OpenMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) OperationsInsightsState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationsInsightsState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) OperationsInsightsStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationsInsightsStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) PeerDbIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"peerDbIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) PermissionLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"permissionLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) PrivateEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) PrivateEndpointIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateEndpointIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) PrivateEndpointIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateEndpointIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) PrivateEndpointLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateEndpointLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) PrivateEndpointLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateEndpointLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) RefreshableMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refreshableMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) RefreshableState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refreshableState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) ScheduledOperationDetails() OracleDatabaseAutonomousDatabasePropertiesScheduledOperationDetailsList {
	var returns OracleDatabaseAutonomousDatabasePropertiesScheduledOperationDetailsList
	_jsii_.Get(
		j,
		"scheduledOperationDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) SqlWebDeveloperUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlWebDeveloperUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) SupportedCloneRegions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedCloneRegions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) TotalAutoBackupStorageSizeGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalAutoBackupStorageSizeGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) UsedDataStorageSizeTbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"usedDataStorageSizeTbs",
		&returns,
	)
	return returns
}


func NewOracleDatabaseAutonomousDatabasePropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OracleDatabaseAutonomousDatabasePropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewOracleDatabaseAutonomousDatabasePropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.oracleDatabaseAutonomousDatabase.OracleDatabaseAutonomousDatabasePropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOracleDatabaseAutonomousDatabasePropertiesOutputReference_Override(o OracleDatabaseAutonomousDatabasePropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.oracleDatabaseAutonomousDatabase.OracleDatabaseAutonomousDatabasePropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference)SetBackupRetentionPeriodDays(val *float64) {
	if err := j.validateSetBackupRetentionPeriodDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupRetentionPeriodDays",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference)SetCharacterSet(val *string) {
	if err := j.validateSetCharacterSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"characterSet",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference)SetComputeCount(val *float64) {
	if err := j.validateSetComputeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeCount",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference)SetDataStorageSizeGb(val *float64) {
	if err := j.validateSetDataStorageSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataStorageSizeGb",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference)SetDataStorageSizeTb(val *float64) {
	if err := j.validateSetDataStorageSizeTbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataStorageSizeTb",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference)SetDbEdition(val *string) {
	if err := j.validateSetDbEditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbEdition",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference)SetDbVersion(val *string) {
	if err := j.validateSetDbVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbVersion",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference)SetDbWorkload(val *string) {
	if err := j.validateSetDbWorkloadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbWorkload",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference)SetInternalValue(val *OracleDatabaseAutonomousDatabaseProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference)SetIsAutoScalingEnabled(val interface{}) {
	if err := j.validateSetIsAutoScalingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isAutoScalingEnabled",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference)SetIsStorageAutoScalingEnabled(val interface{}) {
	if err := j.validateSetIsStorageAutoScalingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isStorageAutoScalingEnabled",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference)SetLicenseType(val *string) {
	if err := j.validateSetLicenseTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseType",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference)SetMaintenanceScheduleType(val *string) {
	if err := j.validateSetMaintenanceScheduleTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintenanceScheduleType",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference)SetMtlsConnectionRequired(val interface{}) {
	if err := j.validateSetMtlsConnectionRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mtlsConnectionRequired",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference)SetNCharacterSet(val *string) {
	if err := j.validateSetNCharacterSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nCharacterSet",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference)SetOperationsInsightsState(val *string) {
	if err := j.validateSetOperationsInsightsStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operationsInsightsState",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference)SetPrivateEndpointIp(val *string) {
	if err := j.validateSetPrivateEndpointIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateEndpointIp",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference)SetPrivateEndpointLabel(val *string) {
	if err := j.validateSetPrivateEndpointLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateEndpointLabel",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) PutCustomerContacts(value interface{}) {
	if err := o.validatePutCustomerContactsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putCustomerContacts",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) ResetBackupRetentionPeriodDays() {
	_jsii_.InvokeVoid(
		o,
		"resetBackupRetentionPeriodDays",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) ResetCharacterSet() {
	_jsii_.InvokeVoid(
		o,
		"resetCharacterSet",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) ResetComputeCount() {
	_jsii_.InvokeVoid(
		o,
		"resetComputeCount",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) ResetCustomerContacts() {
	_jsii_.InvokeVoid(
		o,
		"resetCustomerContacts",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) ResetDataStorageSizeGb() {
	_jsii_.InvokeVoid(
		o,
		"resetDataStorageSizeGb",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) ResetDataStorageSizeTb() {
	_jsii_.InvokeVoid(
		o,
		"resetDataStorageSizeTb",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) ResetDbEdition() {
	_jsii_.InvokeVoid(
		o,
		"resetDbEdition",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) ResetDbVersion() {
	_jsii_.InvokeVoid(
		o,
		"resetDbVersion",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) ResetIsAutoScalingEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetIsAutoScalingEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) ResetIsStorageAutoScalingEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetIsStorageAutoScalingEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) ResetMaintenanceScheduleType() {
	_jsii_.InvokeVoid(
		o,
		"resetMaintenanceScheduleType",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) ResetMtlsConnectionRequired() {
	_jsii_.InvokeVoid(
		o,
		"resetMtlsConnectionRequired",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) ResetNCharacterSet() {
	_jsii_.InvokeVoid(
		o,
		"resetNCharacterSet",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) ResetOperationsInsightsState() {
	_jsii_.InvokeVoid(
		o,
		"resetOperationsInsightsState",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) ResetPrivateEndpointIp() {
	_jsii_.InvokeVoid(
		o,
		"resetPrivateEndpointIp",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) ResetPrivateEndpointLabel() {
	_jsii_.InvokeVoid(
		o,
		"resetPrivateEndpointLabel",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseAutonomousDatabasePropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

