// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagoogleoracledatabaseautonomousdatabases

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/datagoogleoracledatabaseautonomousdatabases/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference interface {
	cdktf.ComplexObject
	ActualUsedDataStorageSizeTb() *float64
	AllocatedStorageSizeTb() *float64
	ApexDetails() DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesApexDetailsList
	ArePrimaryAllowlistedIpsUsed() cdktf.IResolvable
	AutonomousContainerDatabaseId() *string
	AvailableUpgradeVersions() *[]*string
	BackupRetentionPeriodDays() *float64
	CharacterSet() *string
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
	ConnectionStrings() DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesConnectionStringsList
	ConnectionUrls() DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesConnectionUrlsList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomerContacts() DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesCustomerContactsList
	DatabaseManagementState() *string
	DataSafeState() *string
	DataStorageSizeGb() *float64
	DataStorageSizeTb() *float64
	DbEdition() *string
	DbVersion() *string
	DbWorkload() *string
	FailedDataRecoveryDuration() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesProperties
	SetInternalValue(val *DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesProperties)
	IsAutoScalingEnabled() cdktf.IResolvable
	IsLocalDataGuardEnabled() cdktf.IResolvable
	IsStorageAutoScalingEnabled() cdktf.IResolvable
	LicenseType() *string
	LifecycleDetails() *string
	LocalAdgAutoFailoverMaxDataLossLimit() *float64
	LocalDisasterRecoveryType() *string
	LocalStandbyDb() DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesLocalStandbyDbList
	MaintenanceBeginTime() *string
	MaintenanceEndTime() *string
	MaintenanceScheduleType() *string
	MemoryPerOracleComputeUnitGbs() *float64
	MemoryTableGbs() *float64
	MtlsConnectionRequired() cdktf.IResolvable
	NCharacterSet() *string
	NextLongTermBackupTime() *string
	Ocid() *string
	OciUrl() *string
	OpenMode() *string
	OperationsInsightsState() *string
	PeerDbIds() *[]*string
	PermissionLevel() *string
	PrivateEndpoint() *string
	PrivateEndpointIp() *string
	PrivateEndpointLabel() *string
	RefreshableMode() *string
	RefreshableState() *string
	Role() *string
	ScheduledOperationDetails() DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesScheduledOperationDetailsList
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference
type jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) ActualUsedDataStorageSizeTb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"actualUsedDataStorageSizeTb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) AllocatedStorageSizeTb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocatedStorageSizeTb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) ApexDetails() DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesApexDetailsList {
	var returns DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesApexDetailsList
	_jsii_.Get(
		j,
		"apexDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) ArePrimaryAllowlistedIpsUsed() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"arePrimaryAllowlistedIpsUsed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) AutonomousContainerDatabaseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autonomousContainerDatabaseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) AvailableUpgradeVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availableUpgradeVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) BackupRetentionPeriodDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionPeriodDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) CharacterSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"characterSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) ComputeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) ConnectionStrings() DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesConnectionStringsList {
	var returns DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesConnectionStringsList
	_jsii_.Get(
		j,
		"connectionStrings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) ConnectionUrls() DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesConnectionUrlsList {
	var returns DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesConnectionUrlsList
	_jsii_.Get(
		j,
		"connectionUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) CustomerContacts() DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesCustomerContactsList {
	var returns DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesCustomerContactsList
	_jsii_.Get(
		j,
		"customerContacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) DatabaseManagementState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseManagementState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) DataSafeState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSafeState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) DataStorageSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) DataStorageSizeTb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeTb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) DbEdition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbEdition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) DbVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) DbWorkload() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbWorkload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) FailedDataRecoveryDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failedDataRecoveryDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) InternalValue() *DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesProperties {
	var returns *DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) IsAutoScalingEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isAutoScalingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) IsLocalDataGuardEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isLocalDataGuardEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) IsStorageAutoScalingEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isStorageAutoScalingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) LicenseType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) LifecycleDetails() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) LocalAdgAutoFailoverMaxDataLossLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"localAdgAutoFailoverMaxDataLossLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) LocalDisasterRecoveryType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localDisasterRecoveryType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) LocalStandbyDb() DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesLocalStandbyDbList {
	var returns DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesLocalStandbyDbList
	_jsii_.Get(
		j,
		"localStandbyDb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) MaintenanceBeginTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceBeginTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) MaintenanceEndTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceEndTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) MaintenanceScheduleType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceScheduleType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) MemoryPerOracleComputeUnitGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryPerOracleComputeUnitGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) MemoryTableGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryTableGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) MtlsConnectionRequired() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"mtlsConnectionRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) NCharacterSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nCharacterSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) NextLongTermBackupTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextLongTermBackupTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) Ocid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) OciUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) OpenMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) OperationsInsightsState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationsInsightsState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) PeerDbIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"peerDbIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) PermissionLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"permissionLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) PrivateEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) PrivateEndpointIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateEndpointIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) PrivateEndpointLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateEndpointLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) RefreshableMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refreshableMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) RefreshableState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refreshableState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) ScheduledOperationDetails() DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesScheduledOperationDetailsList {
	var returns DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesScheduledOperationDetailsList
	_jsii_.Get(
		j,
		"scheduledOperationDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) SqlWebDeveloperUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlWebDeveloperUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) SupportedCloneRegions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedCloneRegions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) TotalAutoBackupStorageSizeGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalAutoBackupStorageSizeGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) UsedDataStorageSizeTbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"usedDataStorageSizeTbs",
		&returns,
	)
	return returns
}


func NewDataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewDataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleOracleDatabaseAutonomousDatabases.DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference_Override(d DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleOracleDatabaseAutonomousDatabases.DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference)SetInternalValue(val *DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasesAutonomousDatabasesPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

