// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databasemigrationserviceconnectionprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v12/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v12/databasemigrationserviceconnectionprofile/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference interface {
	cdktf.ComplexObject
	ActivationPolicy() *string
	SetActivationPolicy(val *string)
	ActivationPolicyInput() *string
	AutoStorageIncrease() interface{}
	SetAutoStorageIncrease(val interface{})
	AutoStorageIncreaseInput() interface{}
	CmekKeyName() *string
	SetCmekKeyName(val *string)
	CmekKeyNameInput() *string
	Collation() *string
	SetCollation(val *string)
	CollationInput() *string
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DatabaseFlags() *map[string]*string
	SetDatabaseFlags(val *map[string]*string)
	DatabaseFlagsInput() *map[string]*string
	DatabaseVersion() *string
	SetDatabaseVersion(val *string)
	DatabaseVersionInput() *string
	DataDiskSizeGb() *string
	SetDataDiskSizeGb(val *string)
	DataDiskSizeGbInput() *string
	DataDiskType() *string
	SetDataDiskType(val *string)
	DataDiskTypeInput() *string
	Edition() *string
	SetEdition(val *string)
	EditionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DatabaseMigrationServiceConnectionProfileCloudsqlSettings
	SetInternalValue(val *DatabaseMigrationServiceConnectionProfileCloudsqlSettings)
	IpConfig() DatabaseMigrationServiceConnectionProfileCloudsqlSettingsIpConfigOutputReference
	IpConfigInput() *DatabaseMigrationServiceConnectionProfileCloudsqlSettingsIpConfig
	RootPassword() *string
	SetRootPassword(val *string)
	RootPasswordInput() *string
	RootPasswordSet() cdktf.IResolvable
	SourceId() *string
	SetSourceId(val *string)
	SourceIdInput() *string
	StorageAutoResizeLimit() *string
	SetStorageAutoResizeLimit(val *string)
	StorageAutoResizeLimitInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Tier() *string
	SetTier(val *string)
	TierInput() *string
	UserLabels() *map[string]*string
	SetUserLabels(val *map[string]*string)
	UserLabelsInput() *map[string]*string
	Zone() *string
	SetZone(val *string)
	ZoneInput() *string
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
	PutIpConfig(value *DatabaseMigrationServiceConnectionProfileCloudsqlSettingsIpConfig)
	ResetActivationPolicy()
	ResetAutoStorageIncrease()
	ResetCmekKeyName()
	ResetCollation()
	ResetDatabaseFlags()
	ResetDatabaseVersion()
	ResetDataDiskSizeGb()
	ResetDataDiskType()
	ResetEdition()
	ResetIpConfig()
	ResetRootPassword()
	ResetStorageAutoResizeLimit()
	ResetTier()
	ResetUserLabels()
	ResetZone()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference
type jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) ActivationPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) ActivationPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) AutoStorageIncrease() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoStorageIncrease",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) AutoStorageIncreaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoStorageIncreaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) CmekKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cmekKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) CmekKeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cmekKeyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) Collation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) CollationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) DatabaseFlags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"databaseFlags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) DatabaseFlagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"databaseFlagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) DatabaseVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) DatabaseVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) DataDiskSizeGb() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataDiskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) DataDiskSizeGbInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataDiskSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) DataDiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataDiskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) DataDiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataDiskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) Edition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) EditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"editionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) InternalValue() *DatabaseMigrationServiceConnectionProfileCloudsqlSettings {
	var returns *DatabaseMigrationServiceConnectionProfileCloudsqlSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) IpConfig() DatabaseMigrationServiceConnectionProfileCloudsqlSettingsIpConfigOutputReference {
	var returns DatabaseMigrationServiceConnectionProfileCloudsqlSettingsIpConfigOutputReference
	_jsii_.Get(
		j,
		"ipConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) IpConfigInput() *DatabaseMigrationServiceConnectionProfileCloudsqlSettingsIpConfig {
	var returns *DatabaseMigrationServiceConnectionProfileCloudsqlSettingsIpConfig
	_jsii_.Get(
		j,
		"ipConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) RootPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) RootPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) RootPasswordSet() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"rootPasswordSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) SourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) SourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) StorageAutoResizeLimit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAutoResizeLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) StorageAutoResizeLimitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAutoResizeLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) Tier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) TierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) UserLabels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"userLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) UserLabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"userLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) ZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneInput",
		&returns,
	)
	return returns
}


func NewDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.databaseMigrationServiceConnectionProfile.DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference_Override(d DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.databaseMigrationServiceConnectionProfile.DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference)SetActivationPolicy(val *string) {
	if err := j.validateSetActivationPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activationPolicy",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference)SetAutoStorageIncrease(val interface{}) {
	if err := j.validateSetAutoStorageIncreaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoStorageIncrease",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference)SetCmekKeyName(val *string) {
	if err := j.validateSetCmekKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cmekKeyName",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference)SetCollation(val *string) {
	if err := j.validateSetCollationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"collation",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference)SetDatabaseFlags(val *map[string]*string) {
	if err := j.validateSetDatabaseFlagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseFlags",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference)SetDatabaseVersion(val *string) {
	if err := j.validateSetDatabaseVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseVersion",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference)SetDataDiskSizeGb(val *string) {
	if err := j.validateSetDataDiskSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataDiskSizeGb",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference)SetDataDiskType(val *string) {
	if err := j.validateSetDataDiskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataDiskType",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference)SetEdition(val *string) {
	if err := j.validateSetEditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edition",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference)SetInternalValue(val *DatabaseMigrationServiceConnectionProfileCloudsqlSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference)SetRootPassword(val *string) {
	if err := j.validateSetRootPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootPassword",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference)SetSourceId(val *string) {
	if err := j.validateSetSourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceId",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference)SetStorageAutoResizeLimit(val *string) {
	if err := j.validateSetStorageAutoResizeLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAutoResizeLimit",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference)SetTier(val *string) {
	if err := j.validateSetTierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tier",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference)SetUserLabels(val *map[string]*string) {
	if err := j.validateSetUserLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userLabels",
		val,
	)
}

func (j *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference)SetZone(val *string) {
	if err := j.validateSetZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zone",
		val,
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) PutIpConfig(value *DatabaseMigrationServiceConnectionProfileCloudsqlSettingsIpConfig) {
	if err := d.validatePutIpConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putIpConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) ResetActivationPolicy() {
	_jsii_.InvokeVoid(
		d,
		"resetActivationPolicy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) ResetAutoStorageIncrease() {
	_jsii_.InvokeVoid(
		d,
		"resetAutoStorageIncrease",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) ResetCmekKeyName() {
	_jsii_.InvokeVoid(
		d,
		"resetCmekKeyName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) ResetCollation() {
	_jsii_.InvokeVoid(
		d,
		"resetCollation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) ResetDatabaseFlags() {
	_jsii_.InvokeVoid(
		d,
		"resetDatabaseFlags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) ResetDatabaseVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetDatabaseVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) ResetDataDiskSizeGb() {
	_jsii_.InvokeVoid(
		d,
		"resetDataDiskSizeGb",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) ResetDataDiskType() {
	_jsii_.InvokeVoid(
		d,
		"resetDataDiskType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) ResetEdition() {
	_jsii_.InvokeVoid(
		d,
		"resetEdition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) ResetIpConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetIpConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) ResetRootPassword() {
	_jsii_.InvokeVoid(
		d,
		"resetRootPassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) ResetStorageAutoResizeLimit() {
	_jsii_.InvokeVoid(
		d,
		"resetStorageAutoResizeLimit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) ResetTier() {
	_jsii_.InvokeVoid(
		d,
		"resetTier",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) ResetUserLabels() {
	_jsii_.InvokeVoid(
		d,
		"resetUserLabels",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) ResetZone() {
	_jsii_.InvokeVoid(
		d,
		"resetZone",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DatabaseMigrationServiceConnectionProfileCloudsqlSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

