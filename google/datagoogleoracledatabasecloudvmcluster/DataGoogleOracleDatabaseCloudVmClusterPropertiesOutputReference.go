// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagoogleoracledatabasecloudvmcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v16/datagoogleoracledatabasecloudvmcluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference interface {
	cdktf.ComplexObject
	ClusterName() *string
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DataStorageSizeTb() *float64
	DbNodeStorageSizeGb() *float64
	DbServerOcids() *[]*string
	DiagnosticsDataCollectionOptions() DataGoogleOracleDatabaseCloudVmClusterPropertiesDiagnosticsDataCollectionOptionsList
	DiskRedundancy() *string
	DnsListenerIp() *string
	Domain() *string
	// Experimental.
	Fqn() *string
	GiVersion() *string
	Hostname() *string
	HostnamePrefix() *string
	InternalValue() *DataGoogleOracleDatabaseCloudVmClusterProperties
	SetInternalValue(val *DataGoogleOracleDatabaseCloudVmClusterProperties)
	LicenseType() *string
	LocalBackupEnabled() cdktf.IResolvable
	MemorySizeGb() *float64
	NodeCount() *float64
	Ocid() *string
	OciUrl() *string
	OcpuCount() *float64
	ScanDns() *string
	ScanDnsRecordId() *string
	ScanIpIds() *[]*string
	ScanListenerPortTcp() *float64
	ScanListenerPortTcpSsl() *float64
	Shape() *string
	SparseDiskgroupEnabled() cdktf.IResolvable
	SshPublicKeys() *[]*string
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
	TimeZone() DataGoogleOracleDatabaseCloudVmClusterPropertiesTimeZoneList
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

// The jsii proxy struct for DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference
type jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) CompartmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compartmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) CpuCoreCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cpuCoreCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) DataStorageSizeTb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataStorageSizeTb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) DbNodeStorageSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dbNodeStorageSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) DbServerOcids() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dbServerOcids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) DiagnosticsDataCollectionOptions() DataGoogleOracleDatabaseCloudVmClusterPropertiesDiagnosticsDataCollectionOptionsList {
	var returns DataGoogleOracleDatabaseCloudVmClusterPropertiesDiagnosticsDataCollectionOptionsList
	_jsii_.Get(
		j,
		"diagnosticsDataCollectionOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) DiskRedundancy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskRedundancy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) DnsListenerIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsListenerIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) GiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"giVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) HostnamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnamePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) InternalValue() *DataGoogleOracleDatabaseCloudVmClusterProperties {
	var returns *DataGoogleOracleDatabaseCloudVmClusterProperties
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) LicenseType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) LocalBackupEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"localBackupEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) MemorySizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) NodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) Ocid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) OciUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) OcpuCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ocpuCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) ScanDns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scanDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) ScanDnsRecordId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scanDnsRecordId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) ScanIpIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"scanIpIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) ScanListenerPortTcp() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scanListenerPortTcp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) ScanListenerPortTcpSsl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"scanListenerPortTcpSsl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) Shape() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shape",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) SparseDiskgroupEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"sparseDiskgroupEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) SshPublicKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sshPublicKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) StorageSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) SystemVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"systemVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) TimeZone() DataGoogleOracleDatabaseCloudVmClusterPropertiesTimeZoneList {
	var returns DataGoogleOracleDatabaseCloudVmClusterPropertiesTimeZoneList
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}


func NewDataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference {
	_init_.Initialize()

	if err := validateNewDataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleOracleDatabaseCloudVmCluster.DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference_Override(d DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleOracleDatabaseCloudVmCluster.DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference)SetInternalValue(val *DataGoogleOracleDatabaseCloudVmClusterProperties) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataGoogleOracleDatabaseCloudVmClusterPropertiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

