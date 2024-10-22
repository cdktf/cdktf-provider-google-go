// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oracledatabasecloudvmcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v14/oracledatabasecloudvmcluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OracleDatabaseCloudVmClusterPropertiesOutputReference interface {
	cdktf.ComplexObject
	ClusterName() *string
	SetClusterName(val *string)
	ClusterNameInput() *string
	CompartmentId() *string
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
	CpuCoreCount() *float64
	SetCpuCoreCount(val *float64)
	CpuCoreCountInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DataStorageSizeTb() *float64
	SetDataStorageSizeTb(val *float64)
	DataStorageSizeTbInput() *float64
	DbNodeStorageSizeGb() *float64
	SetDbNodeStorageSizeGb(val *float64)
	DbNodeStorageSizeGbInput() *float64
	DbServerOcids() *[]*string
	SetDbServerOcids(val *[]*string)
	DbServerOcidsInput() *[]*string
	DiagnosticsDataCollectionOptions() OracleDatabaseCloudVmClusterPropertiesDiagnosticsDataCollectionOptionsOutputReference
	DiagnosticsDataCollectionOptionsInput() *OracleDatabaseCloudVmClusterPropertiesDiagnosticsDataCollectionOptions
	DiskRedundancy() *string
	SetDiskRedundancy(val *string)
	DiskRedundancyInput() *string
	DnsListenerIp() *string
	Domain() *string
	// Experimental.
	Fqn() *string
	GiVersion() *string
	SetGiVersion(val *string)
	GiVersionInput() *string
	Hostname() *string
	HostnamePrefix() *string
	SetHostnamePrefix(val *string)
	HostnamePrefixInput() *string
	InternalValue() *OracleDatabaseCloudVmClusterProperties
	SetInternalValue(val *OracleDatabaseCloudVmClusterProperties)
	LicenseType() *string
	SetLicenseType(val *string)
	LicenseTypeInput() *string
	LocalBackupEnabled() interface{}
	SetLocalBackupEnabled(val interface{})
	LocalBackupEnabledInput() interface{}
	MemorySizeGb() *float64
	SetMemorySizeGb(val *float64)
	MemorySizeGbInput() *float64
	NodeCount() *float64
	SetNodeCount(val *float64)
	NodeCountInput() *float64
	Ocid() *string
	OciUrl() *string
	OcpuCount() *float64
	SetOcpuCount(val *float64)
	OcpuCountInput() *float64
	ScanDns() *string
	ScanDnsRecordId() *string
	ScanIpIds() *[]*string
	ScanListenerPortTcp() *float64
	ScanListenerPortTcpSsl() *float64
	Shape() *string
	SparseDiskgroupEnabled() interface{}
	SetSparseDiskgroupEnabled(val interface{})
	SparseDiskgroupEnabledInput() interface{}
	SshPublicKeys() *[]*string
	SetSshPublicKeys(val *[]*string)
	SshPublicKeysInput() *[]*string
	State() *string
	StorageSizeGb() *float64
	SystemVersion() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeZone() OracleDatabaseCloudVmClusterPropertiesTimeZoneOutputReference
	TimeZoneInput() *OracleDatabaseCloudVmClusterPropertiesTimeZone
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
	PutDiagnosticsDataCollectionOptions(value *OracleDatabaseCloudVmClusterPropertiesDiagnosticsDataCollectionOptions)
	PutTimeZone(value *OracleDatabaseCloudVmClusterPropertiesTimeZone)
	ResetClusterName()
	ResetDataStorageSizeTb()
	ResetDbNodeStorageSizeGb()
	ResetDbServerOcids()
	ResetDiagnosticsDataCollectionOptions()
	ResetDiskRedundancy()
	ResetGiVersion()
	ResetHostnamePrefix()
	ResetLocalBackupEnabled()
	ResetMemorySizeGb()
	ResetNodeCount()
	ResetOcpuCount()
	ResetSparseDiskgroupEnabled()
	ResetSshPublicKeys()
	ResetTimeZone()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OracleDatabaseCloudVmClusterPropertiesOutputReference
type jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) ClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) CompartmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compartmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) CpuCoreCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCoreCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) CpuCoreCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCoreCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) DataStorageSizeTb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeTb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) DataStorageSizeTbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeTbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) DbNodeStorageSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dbNodeStorageSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) DbNodeStorageSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dbNodeStorageSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) DbServerOcids() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dbServerOcids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) DbServerOcidsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dbServerOcidsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) DiagnosticsDataCollectionOptions() OracleDatabaseCloudVmClusterPropertiesDiagnosticsDataCollectionOptionsOutputReference {
	var returns OracleDatabaseCloudVmClusterPropertiesDiagnosticsDataCollectionOptionsOutputReference
	_jsii_.Get(
		j,
		"diagnosticsDataCollectionOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) DiagnosticsDataCollectionOptionsInput() *OracleDatabaseCloudVmClusterPropertiesDiagnosticsDataCollectionOptions {
	var returns *OracleDatabaseCloudVmClusterPropertiesDiagnosticsDataCollectionOptions
	_jsii_.Get(
		j,
		"diagnosticsDataCollectionOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) DiskRedundancy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskRedundancy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) DiskRedundancyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskRedundancyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) DnsListenerIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsListenerIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) GiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"giVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) GiVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"giVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) HostnamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnamePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) HostnamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnamePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) InternalValue() *OracleDatabaseCloudVmClusterProperties {
	var returns *OracleDatabaseCloudVmClusterProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) LicenseType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) LicenseTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) LocalBackupEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localBackupEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) LocalBackupEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localBackupEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) MemorySizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) MemorySizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) NodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) NodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) Ocid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) OciUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) OcpuCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ocpuCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) OcpuCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ocpuCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) ScanDns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scanDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) ScanDnsRecordId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scanDnsRecordId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) ScanIpIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scanIpIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) ScanListenerPortTcp() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scanListenerPortTcp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) ScanListenerPortTcpSsl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scanListenerPortTcpSsl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) Shape() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shape",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) SparseDiskgroupEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sparseDiskgroupEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) SparseDiskgroupEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sparseDiskgroupEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) SshPublicKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshPublicKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) SshPublicKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshPublicKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) StorageSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) SystemVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"systemVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) TimeZone() OracleDatabaseCloudVmClusterPropertiesTimeZoneOutputReference {
	var returns OracleDatabaseCloudVmClusterPropertiesTimeZoneOutputReference
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) TimeZoneInput() *OracleDatabaseCloudVmClusterPropertiesTimeZone {
	var returns *OracleDatabaseCloudVmClusterPropertiesTimeZone
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}


func NewOracleDatabaseCloudVmClusterPropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OracleDatabaseCloudVmClusterPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewOracleDatabaseCloudVmClusterPropertiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.oracleDatabaseCloudVmCluster.OracleDatabaseCloudVmClusterPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOracleDatabaseCloudVmClusterPropertiesOutputReference_Override(o OracleDatabaseCloudVmClusterPropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.oracleDatabaseCloudVmCluster.OracleDatabaseCloudVmClusterPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference)SetClusterName(val *string) {
	if err := j.validateSetClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference)SetCpuCoreCount(val *float64) {
	if err := j.validateSetCpuCoreCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuCoreCount",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference)SetDataStorageSizeTb(val *float64) {
	if err := j.validateSetDataStorageSizeTbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataStorageSizeTb",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference)SetDbNodeStorageSizeGb(val *float64) {
	if err := j.validateSetDbNodeStorageSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbNodeStorageSizeGb",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference)SetDbServerOcids(val *[]*string) {
	if err := j.validateSetDbServerOcidsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbServerOcids",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference)SetDiskRedundancy(val *string) {
	if err := j.validateSetDiskRedundancyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskRedundancy",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference)SetGiVersion(val *string) {
	if err := j.validateSetGiVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"giVersion",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference)SetHostnamePrefix(val *string) {
	if err := j.validateSetHostnamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostnamePrefix",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference)SetInternalValue(val *OracleDatabaseCloudVmClusterProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference)SetLicenseType(val *string) {
	if err := j.validateSetLicenseTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseType",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference)SetLocalBackupEnabled(val interface{}) {
	if err := j.validateSetLocalBackupEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localBackupEnabled",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference)SetMemorySizeGb(val *float64) {
	if err := j.validateSetMemorySizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memorySizeGb",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference)SetNodeCount(val *float64) {
	if err := j.validateSetNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeCount",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference)SetOcpuCount(val *float64) {
	if err := j.validateSetOcpuCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ocpuCount",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference)SetSparseDiskgroupEnabled(val interface{}) {
	if err := j.validateSetSparseDiskgroupEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparseDiskgroupEnabled",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference)SetSshPublicKeys(val *[]*string) {
	if err := j.validateSetSshPublicKeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sshPublicKeys",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) PutDiagnosticsDataCollectionOptions(value *OracleDatabaseCloudVmClusterPropertiesDiagnosticsDataCollectionOptions) {
	if err := o.validatePutDiagnosticsDataCollectionOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putDiagnosticsDataCollectionOptions",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) PutTimeZone(value *OracleDatabaseCloudVmClusterPropertiesTimeZone) {
	if err := o.validatePutTimeZoneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putTimeZone",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) ResetClusterName() {
	_jsii_.InvokeVoid(
		o,
		"resetClusterName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) ResetDataStorageSizeTb() {
	_jsii_.InvokeVoid(
		o,
		"resetDataStorageSizeTb",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) ResetDbNodeStorageSizeGb() {
	_jsii_.InvokeVoid(
		o,
		"resetDbNodeStorageSizeGb",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) ResetDbServerOcids() {
	_jsii_.InvokeVoid(
		o,
		"resetDbServerOcids",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) ResetDiagnosticsDataCollectionOptions() {
	_jsii_.InvokeVoid(
		o,
		"resetDiagnosticsDataCollectionOptions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) ResetDiskRedundancy() {
	_jsii_.InvokeVoid(
		o,
		"resetDiskRedundancy",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) ResetGiVersion() {
	_jsii_.InvokeVoid(
		o,
		"resetGiVersion",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) ResetHostnamePrefix() {
	_jsii_.InvokeVoid(
		o,
		"resetHostnamePrefix",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) ResetLocalBackupEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetLocalBackupEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) ResetMemorySizeGb() {
	_jsii_.InvokeVoid(
		o,
		"resetMemorySizeGb",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) ResetNodeCount() {
	_jsii_.InvokeVoid(
		o,
		"resetNodeCount",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) ResetOcpuCount() {
	_jsii_.InvokeVoid(
		o,
		"resetOcpuCount",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) ResetSparseDiskgroupEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetSparseDiskgroupEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) ResetSshPublicKeys() {
	_jsii_.InvokeVoid(
		o,
		"resetSshPublicKeys",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) ResetTimeZone() {
	_jsii_.InvokeVoid(
		o,
		"resetTimeZone",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OracleDatabaseCloudVmClusterPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

