package datagooglesqldatabaseinstances

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v8/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v8/datagooglesqldatabaseinstances/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference interface {
	cdktf.ComplexObject
	ActivationPolicy() *string
	ActiveDirectoryConfig() DataGoogleSqlDatabaseInstancesInstancesSettingsActiveDirectoryConfigList
	AdvancedMachineFeatures() DataGoogleSqlDatabaseInstancesInstancesSettingsAdvancedMachineFeaturesList
	AvailabilityType() *string
	BackupConfiguration() DataGoogleSqlDatabaseInstancesInstancesSettingsBackupConfigurationList
	Collation() *string
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
	ConnectorEnforcement() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DatabaseFlags() DataGoogleSqlDatabaseInstancesInstancesSettingsDatabaseFlagsList
	DataCacheConfig() DataGoogleSqlDatabaseInstancesInstancesSettingsDataCacheConfigList
	DeletionProtectionEnabled() cdktf.IResolvable
	DenyMaintenancePeriod() DataGoogleSqlDatabaseInstancesInstancesSettingsDenyMaintenancePeriodList
	DiskAutoresize() cdktf.IResolvable
	DiskAutoresizeLimit() *float64
	DiskSize() *float64
	DiskType() *string
	Edition() *string
	// Experimental.
	Fqn() *string
	InsightsConfig() DataGoogleSqlDatabaseInstancesInstancesSettingsInsightsConfigList
	InternalValue() *DataGoogleSqlDatabaseInstancesInstancesSettings
	SetInternalValue(val *DataGoogleSqlDatabaseInstancesInstancesSettings)
	IpConfiguration() DataGoogleSqlDatabaseInstancesInstancesSettingsIpConfigurationList
	LocationPreference() DataGoogleSqlDatabaseInstancesInstancesSettingsLocationPreferenceList
	MaintenanceWindow() DataGoogleSqlDatabaseInstancesInstancesSettingsMaintenanceWindowList
	PasswordValidationPolicy() DataGoogleSqlDatabaseInstancesInstancesSettingsPasswordValidationPolicyList
	PricingPlan() *string
	SqlServerAuditConfig() DataGoogleSqlDatabaseInstancesInstancesSettingsSqlServerAuditConfigList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Tier() *string
	TimeZone() *string
	UserLabels() cdktf.StringMap
	Version() *float64
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

// The jsii proxy struct for DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference
type jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) ActivationPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) ActiveDirectoryConfig() DataGoogleSqlDatabaseInstancesInstancesSettingsActiveDirectoryConfigList {
	var returns DataGoogleSqlDatabaseInstancesInstancesSettingsActiveDirectoryConfigList
	_jsii_.Get(
		j,
		"activeDirectoryConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) AdvancedMachineFeatures() DataGoogleSqlDatabaseInstancesInstancesSettingsAdvancedMachineFeaturesList {
	var returns DataGoogleSqlDatabaseInstancesInstancesSettingsAdvancedMachineFeaturesList
	_jsii_.Get(
		j,
		"advancedMachineFeatures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) AvailabilityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) BackupConfiguration() DataGoogleSqlDatabaseInstancesInstancesSettingsBackupConfigurationList {
	var returns DataGoogleSqlDatabaseInstancesInstancesSettingsBackupConfigurationList
	_jsii_.Get(
		j,
		"backupConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) Collation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) ConnectorEnforcement() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorEnforcement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) DatabaseFlags() DataGoogleSqlDatabaseInstancesInstancesSettingsDatabaseFlagsList {
	var returns DataGoogleSqlDatabaseInstancesInstancesSettingsDatabaseFlagsList
	_jsii_.Get(
		j,
		"databaseFlags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) DataCacheConfig() DataGoogleSqlDatabaseInstancesInstancesSettingsDataCacheConfigList {
	var returns DataGoogleSqlDatabaseInstancesInstancesSettingsDataCacheConfigList
	_jsii_.Get(
		j,
		"dataCacheConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) DeletionProtectionEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"deletionProtectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) DenyMaintenancePeriod() DataGoogleSqlDatabaseInstancesInstancesSettingsDenyMaintenancePeriodList {
	var returns DataGoogleSqlDatabaseInstancesInstancesSettingsDenyMaintenancePeriodList
	_jsii_.Get(
		j,
		"denyMaintenancePeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) DiskAutoresize() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"diskAutoresize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) DiskAutoresizeLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskAutoresizeLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) DiskSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) DiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) Edition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) InsightsConfig() DataGoogleSqlDatabaseInstancesInstancesSettingsInsightsConfigList {
	var returns DataGoogleSqlDatabaseInstancesInstancesSettingsInsightsConfigList
	_jsii_.Get(
		j,
		"insightsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) InternalValue() *DataGoogleSqlDatabaseInstancesInstancesSettings {
	var returns *DataGoogleSqlDatabaseInstancesInstancesSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) IpConfiguration() DataGoogleSqlDatabaseInstancesInstancesSettingsIpConfigurationList {
	var returns DataGoogleSqlDatabaseInstancesInstancesSettingsIpConfigurationList
	_jsii_.Get(
		j,
		"ipConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) LocationPreference() DataGoogleSqlDatabaseInstancesInstancesSettingsLocationPreferenceList {
	var returns DataGoogleSqlDatabaseInstancesInstancesSettingsLocationPreferenceList
	_jsii_.Get(
		j,
		"locationPreference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) MaintenanceWindow() DataGoogleSqlDatabaseInstancesInstancesSettingsMaintenanceWindowList {
	var returns DataGoogleSqlDatabaseInstancesInstancesSettingsMaintenanceWindowList
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) PasswordValidationPolicy() DataGoogleSqlDatabaseInstancesInstancesSettingsPasswordValidationPolicyList {
	var returns DataGoogleSqlDatabaseInstancesInstancesSettingsPasswordValidationPolicyList
	_jsii_.Get(
		j,
		"passwordValidationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) PricingPlan() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pricingPlan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) SqlServerAuditConfig() DataGoogleSqlDatabaseInstancesInstancesSettingsSqlServerAuditConfigList {
	var returns DataGoogleSqlDatabaseInstancesInstancesSettingsSqlServerAuditConfigList
	_jsii_.Get(
		j,
		"sqlServerAuditConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) Tier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) UserLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"userLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) Version() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


func NewDataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDataGoogleSqlDatabaseInstancesInstancesSettingsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleSqlDatabaseInstances.DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference_Override(d DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleSqlDatabaseInstances.DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference)SetInternalValue(val *DataGoogleSqlDatabaseInstancesInstancesSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataGoogleSqlDatabaseInstancesInstancesSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

