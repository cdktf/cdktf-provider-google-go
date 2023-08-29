// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagooglecontainercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v9/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v9/datagooglecontainercluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleContainerClusterNodeConfigOutputReference interface {
	cdktf.ComplexObject
	AdvancedMachineFeatures() DataGoogleContainerClusterNodeConfigAdvancedMachineFeaturesList
	BootDiskKmsKey() *string
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
	DiskSizeGb() *float64
	DiskType() *string
	EphemeralStorageLocalSsdConfig() DataGoogleContainerClusterNodeConfigEphemeralStorageLocalSsdConfigList
	// Experimental.
	Fqn() *string
	GcfsConfig() DataGoogleContainerClusterNodeConfigGcfsConfigList
	GuestAccelerator() DataGoogleContainerClusterNodeConfigGuestAcceleratorList
	Gvnic() DataGoogleContainerClusterNodeConfigGvnicList
	HostMaintenancePolicy() DataGoogleContainerClusterNodeConfigHostMaintenancePolicyList
	ImageType() *string
	InternalValue() *DataGoogleContainerClusterNodeConfig
	SetInternalValue(val *DataGoogleContainerClusterNodeConfig)
	KubeletConfig() DataGoogleContainerClusterNodeConfigKubeletConfigList
	Labels() cdktf.StringMap
	LinuxNodeConfig() DataGoogleContainerClusterNodeConfigLinuxNodeConfigList
	LocalNvmeSsdBlockConfig() DataGoogleContainerClusterNodeConfigLocalNvmeSsdBlockConfigList
	LocalSsdCount() *float64
	LoggingVariant() *string
	MachineType() *string
	Metadata() cdktf.StringMap
	MinCpuPlatform() *string
	NodeGroup() *string
	OauthScopes() *[]*string
	Preemptible() cdktf.IResolvable
	ReservationAffinity() DataGoogleContainerClusterNodeConfigReservationAffinityList
	ResourceLabels() cdktf.StringMap
	ServiceAccount() *string
	ShieldedInstanceConfig() DataGoogleContainerClusterNodeConfigShieldedInstanceConfigList
	SoleTenantConfig() DataGoogleContainerClusterNodeConfigSoleTenantConfigList
	Spot() cdktf.IResolvable
	Tags() *[]*string
	Taint() DataGoogleContainerClusterNodeConfigTaintList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WorkloadMetadataConfig() DataGoogleContainerClusterNodeConfigWorkloadMetadataConfigList
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

// The jsii proxy struct for DataGoogleContainerClusterNodeConfigOutputReference
type jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) AdvancedMachineFeatures() DataGoogleContainerClusterNodeConfigAdvancedMachineFeaturesList {
	var returns DataGoogleContainerClusterNodeConfigAdvancedMachineFeaturesList
	_jsii_.Get(
		j,
		"advancedMachineFeatures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) BootDiskKmsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootDiskKmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) DiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) DiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) EphemeralStorageLocalSsdConfig() DataGoogleContainerClusterNodeConfigEphemeralStorageLocalSsdConfigList {
	var returns DataGoogleContainerClusterNodeConfigEphemeralStorageLocalSsdConfigList
	_jsii_.Get(
		j,
		"ephemeralStorageLocalSsdConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) GcfsConfig() DataGoogleContainerClusterNodeConfigGcfsConfigList {
	var returns DataGoogleContainerClusterNodeConfigGcfsConfigList
	_jsii_.Get(
		j,
		"gcfsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) GuestAccelerator() DataGoogleContainerClusterNodeConfigGuestAcceleratorList {
	var returns DataGoogleContainerClusterNodeConfigGuestAcceleratorList
	_jsii_.Get(
		j,
		"guestAccelerator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) Gvnic() DataGoogleContainerClusterNodeConfigGvnicList {
	var returns DataGoogleContainerClusterNodeConfigGvnicList
	_jsii_.Get(
		j,
		"gvnic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) HostMaintenancePolicy() DataGoogleContainerClusterNodeConfigHostMaintenancePolicyList {
	var returns DataGoogleContainerClusterNodeConfigHostMaintenancePolicyList
	_jsii_.Get(
		j,
		"hostMaintenancePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) ImageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) InternalValue() *DataGoogleContainerClusterNodeConfig {
	var returns *DataGoogleContainerClusterNodeConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) KubeletConfig() DataGoogleContainerClusterNodeConfigKubeletConfigList {
	var returns DataGoogleContainerClusterNodeConfigKubeletConfigList
	_jsii_.Get(
		j,
		"kubeletConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) Labels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) LinuxNodeConfig() DataGoogleContainerClusterNodeConfigLinuxNodeConfigList {
	var returns DataGoogleContainerClusterNodeConfigLinuxNodeConfigList
	_jsii_.Get(
		j,
		"linuxNodeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) LocalNvmeSsdBlockConfig() DataGoogleContainerClusterNodeConfigLocalNvmeSsdBlockConfigList {
	var returns DataGoogleContainerClusterNodeConfigLocalNvmeSsdBlockConfigList
	_jsii_.Get(
		j,
		"localNvmeSsdBlockConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) LocalSsdCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"localSsdCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) LoggingVariant() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingVariant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) MachineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) Metadata() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) MinCpuPlatform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCpuPlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) NodeGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) OauthScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"oauthScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) Preemptible() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"preemptible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) ReservationAffinity() DataGoogleContainerClusterNodeConfigReservationAffinityList {
	var returns DataGoogleContainerClusterNodeConfigReservationAffinityList
	_jsii_.Get(
		j,
		"reservationAffinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) ResourceLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"resourceLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) ShieldedInstanceConfig() DataGoogleContainerClusterNodeConfigShieldedInstanceConfigList {
	var returns DataGoogleContainerClusterNodeConfigShieldedInstanceConfigList
	_jsii_.Get(
		j,
		"shieldedInstanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) SoleTenantConfig() DataGoogleContainerClusterNodeConfigSoleTenantConfigList {
	var returns DataGoogleContainerClusterNodeConfigSoleTenantConfigList
	_jsii_.Get(
		j,
		"soleTenantConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) Spot() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"spot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) Taint() DataGoogleContainerClusterNodeConfigTaintList {
	var returns DataGoogleContainerClusterNodeConfigTaintList
	_jsii_.Get(
		j,
		"taint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) WorkloadMetadataConfig() DataGoogleContainerClusterNodeConfigWorkloadMetadataConfigList {
	var returns DataGoogleContainerClusterNodeConfigWorkloadMetadataConfigList
	_jsii_.Get(
		j,
		"workloadMetadataConfig",
		&returns,
	)
	return returns
}


func NewDataGoogleContainerClusterNodeConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataGoogleContainerClusterNodeConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataGoogleContainerClusterNodeConfigOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleContainerCluster.DataGoogleContainerClusterNodeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataGoogleContainerClusterNodeConfigOutputReference_Override(d DataGoogleContainerClusterNodeConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleContainerCluster.DataGoogleContainerClusterNodeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference)SetInternalValue(val *DataGoogleContainerClusterNodeConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataGoogleContainerClusterNodeConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

