// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagoogleoracledatabaseautonomousdatabase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/datagoogleoracledatabaseautonomousdatabase/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference interface {
	cdktf.ComplexObject
	ActualUsedDataStorageSizeTb() *float64
	AllocatedStorageSizeTb() *float64
	ApexDetails() DataGoogleOracleDatabaseAutonomousDatabasePropertiesApexDetailsList
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
	ConnectionStrings() DataGoogleOracleDatabaseAutonomousDatabasePropertiesConnectionStringsList
	ConnectionUrls() DataGoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CustomerContacts() DataGoogleOracleDatabaseAutonomousDatabasePropertiesCustomerContactsList
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
	InternalValue() *DataGoogleOracleDatabaseAutonomousDatabaseProperties
	SetInternalValue(val *DataGoogleOracleDatabaseAutonomousDatabaseProperties)
	IsAutoScalingEnabled() cdktf.IResolvable
	IsLocalDataGuardEnabled() cdktf.IResolvable
	IsStorageAutoScalingEnabled() cdktf.IResolvable
	LicenseType() *string
	LifecycleDetails() *string
	LocalAdgAutoFailoverMaxDataLossLimit() *float64
	LocalDisasterRecoveryType() *string
	LocalStandbyDb() DataGoogleOracleDatabaseAutonomousDatabasePropertiesLocalStandbyDbList
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
	ScheduledOperationDetails() DataGoogleOracleDatabaseAutonomousDatabasePropertiesScheduledOperationDetailsList
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

// The jsii proxy struct for DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference
type jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) ActualUsedDataStorageSizeTb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"actualUsedDataStorageSizeTb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) AllocatedStorageSizeTb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocatedStorageSizeTb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) ApexDetails() DataGoogleOracleDatabaseAutonomousDatabasePropertiesApexDetailsList {
	var returns DataGoogleOracleDatabaseAutonomousDatabasePropertiesApexDetailsList
	_jsii_.Get(
		j,
		"apexDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) ArePrimaryAllowlistedIpsUsed() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"arePrimaryAllowlistedIpsUsed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) AutonomousContainerDatabaseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autonomousContainerDatabaseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) AvailableUpgradeVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availableUpgradeVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) BackupRetentionPeriodDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionPeriodDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) CharacterSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"characterSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) ComputeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"computeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) ConnectionStrings() DataGoogleOracleDatabaseAutonomousDatabasePropertiesConnectionStringsList {
	var returns DataGoogleOracleDatabaseAutonomousDatabasePropertiesConnectionStringsList
	_jsii_.Get(
		j,
		"connectionStrings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) ConnectionUrls() DataGoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsList {
	var returns DataGoogleOracleDatabaseAutonomousDatabasePropertiesConnectionUrlsList
	_jsii_.Get(
		j,
		"connectionUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) CustomerContacts() DataGoogleOracleDatabaseAutonomousDatabasePropertiesCustomerContactsList {
	var returns DataGoogleOracleDatabaseAutonomousDatabasePropertiesCustomerContactsList
	_jsii_.Get(
		j,
		"customerContacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) DatabaseManagementState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseManagementState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) DataSafeState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSafeState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) DataStorageSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) DataStorageSizeTb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeTb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) DbEdition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbEdition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) DbVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) DbWorkload() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbWorkload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) FailedDataRecoveryDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failedDataRecoveryDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) InternalValue() *DataGoogleOracleDatabaseAutonomousDatabaseProperties {
	var returns *DataGoogleOracleDatabaseAutonomousDatabaseProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) IsAutoScalingEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isAutoScalingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) IsLocalDataGuardEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isLocalDataGuardEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) IsStorageAutoScalingEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isStorageAutoScalingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) LicenseType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) LifecycleDetails() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lifecycleDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) LocalAdgAutoFailoverMaxDataLossLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"localAdgAutoFailoverMaxDataLossLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) LocalDisasterRecoveryType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localDisasterRecoveryType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) LocalStandbyDb() DataGoogleOracleDatabaseAutonomousDatabasePropertiesLocalStandbyDbList {
	var returns DataGoogleOracleDatabaseAutonomousDatabasePropertiesLocalStandbyDbList
	_jsii_.Get(
		j,
		"localStandbyDb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) MaintenanceBeginTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceBeginTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) MaintenanceEndTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceEndTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) MaintenanceScheduleType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceScheduleType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) MemoryPerOracleComputeUnitGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryPerOracleComputeUnitGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) MemoryTableGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryTableGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) MtlsConnectionRequired() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"mtlsConnectionRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) NCharacterSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nCharacterSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) NextLongTermBackupTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextLongTermBackupTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) Ocid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) OciUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) OpenMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) OperationsInsightsState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationsInsightsState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) PeerDbIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"peerDbIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) PermissionLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"permissionLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) PrivateEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) PrivateEndpointIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateEndpointIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) PrivateEndpointLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateEndpointLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) RefreshableMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refreshableMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) RefreshableState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"refreshableState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) ScheduledOperationDetails() DataGoogleOracleDatabaseAutonomousDatabasePropertiesScheduledOperationDetailsList {
	var returns DataGoogleOracleDatabaseAutonomousDatabasePropertiesScheduledOperationDetailsList
	_jsii_.Get(
		j,
		"scheduledOperationDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) SqlWebDeveloperUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlWebDeveloperUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) SupportedCloneRegions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedCloneRegions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) TotalAutoBackupStorageSizeGbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalAutoBackupStorageSizeGbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) UsedDataStorageSizeTbs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"usedDataStorageSizeTbs",
		&returns,
	)
	return returns
}


func NewDataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewDataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleOracleDatabaseAutonomousDatabase.DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference_Override(d DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleOracleDatabaseAutonomousDatabase.DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference)SetInternalValue(val *DataGoogleOracleDatabaseAutonomousDatabaseProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseAutonomousDatabasePropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

