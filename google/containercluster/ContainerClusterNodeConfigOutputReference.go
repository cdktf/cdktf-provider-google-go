package containercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v5/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v5/containercluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerClusterNodeConfigOutputReference interface {
	cdktf.ComplexObject
	BootDiskKmsKey() *string
	SetBootDiskKmsKey(val *string)
	BootDiskKmsKeyInput() *string
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
	SetDiskSizeGb(val *float64)
	DiskSizeGbInput() *float64
	DiskType() *string
	SetDiskType(val *string)
	DiskTypeInput() *string
	// Experimental.
	Fqn() *string
	GcfsConfig() ContainerClusterNodeConfigGcfsConfigOutputReference
	GcfsConfigInput() *ContainerClusterNodeConfigGcfsConfig
	GuestAccelerator() ContainerClusterNodeConfigGuestAcceleratorList
	GuestAcceleratorInput() interface{}
	Gvnic() ContainerClusterNodeConfigGvnicOutputReference
	GvnicInput() *ContainerClusterNodeConfigGvnic
	ImageType() *string
	SetImageType(val *string)
	ImageTypeInput() *string
	InternalValue() *ContainerClusterNodeConfig
	SetInternalValue(val *ContainerClusterNodeConfig)
	KubeletConfig() ContainerClusterNodeConfigKubeletConfigOutputReference
	KubeletConfigInput() *ContainerClusterNodeConfigKubeletConfig
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	LinuxNodeConfig() ContainerClusterNodeConfigLinuxNodeConfigOutputReference
	LinuxNodeConfigInput() *ContainerClusterNodeConfigLinuxNodeConfig
	LocalNvmeSsdBlockConfig() ContainerClusterNodeConfigLocalNvmeSsdBlockConfigOutputReference
	LocalNvmeSsdBlockConfigInput() *ContainerClusterNodeConfigLocalNvmeSsdBlockConfig
	LocalSsdCount() *float64
	SetLocalSsdCount(val *float64)
	LocalSsdCountInput() *float64
	LoggingVariant() *string
	SetLoggingVariant(val *string)
	LoggingVariantInput() *string
	MachineType() *string
	SetMachineType(val *string)
	MachineTypeInput() *string
	Metadata() *map[string]*string
	SetMetadata(val *map[string]*string)
	MetadataInput() *map[string]*string
	MinCpuPlatform() *string
	SetMinCpuPlatform(val *string)
	MinCpuPlatformInput() *string
	NodeGroup() *string
	SetNodeGroup(val *string)
	NodeGroupInput() *string
	OauthScopes() *[]*string
	SetOauthScopes(val *[]*string)
	OauthScopesInput() *[]*string
	Preemptible() interface{}
	SetPreemptible(val interface{})
	PreemptibleInput() interface{}
	ReservationAffinity() ContainerClusterNodeConfigReservationAffinityOutputReference
	ReservationAffinityInput() *ContainerClusterNodeConfigReservationAffinity
	ResourceLabels() *map[string]*string
	SetResourceLabels(val *map[string]*string)
	ResourceLabelsInput() *map[string]*string
	ServiceAccount() *string
	SetServiceAccount(val *string)
	ServiceAccountInput() *string
	ShieldedInstanceConfig() ContainerClusterNodeConfigShieldedInstanceConfigOutputReference
	ShieldedInstanceConfigInput() *ContainerClusterNodeConfigShieldedInstanceConfig
	Spot() interface{}
	SetSpot(val interface{})
	SpotInput() interface{}
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	Taint() ContainerClusterNodeConfigTaintList
	TaintInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WorkloadMetadataConfig() ContainerClusterNodeConfigWorkloadMetadataConfigOutputReference
	WorkloadMetadataConfigInput() *ContainerClusterNodeConfigWorkloadMetadataConfig
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
	PutGcfsConfig(value *ContainerClusterNodeConfigGcfsConfig)
	PutGuestAccelerator(value interface{})
	PutGvnic(value *ContainerClusterNodeConfigGvnic)
	PutKubeletConfig(value *ContainerClusterNodeConfigKubeletConfig)
	PutLinuxNodeConfig(value *ContainerClusterNodeConfigLinuxNodeConfig)
	PutLocalNvmeSsdBlockConfig(value *ContainerClusterNodeConfigLocalNvmeSsdBlockConfig)
	PutReservationAffinity(value *ContainerClusterNodeConfigReservationAffinity)
	PutShieldedInstanceConfig(value *ContainerClusterNodeConfigShieldedInstanceConfig)
	PutTaint(value interface{})
	PutWorkloadMetadataConfig(value *ContainerClusterNodeConfigWorkloadMetadataConfig)
	ResetBootDiskKmsKey()
	ResetDiskSizeGb()
	ResetDiskType()
	ResetGcfsConfig()
	ResetGuestAccelerator()
	ResetGvnic()
	ResetImageType()
	ResetKubeletConfig()
	ResetLabels()
	ResetLinuxNodeConfig()
	ResetLocalNvmeSsdBlockConfig()
	ResetLocalSsdCount()
	ResetLoggingVariant()
	ResetMachineType()
	ResetMetadata()
	ResetMinCpuPlatform()
	ResetNodeGroup()
	ResetOauthScopes()
	ResetPreemptible()
	ResetReservationAffinity()
	ResetResourceLabels()
	ResetServiceAccount()
	ResetShieldedInstanceConfig()
	ResetSpot()
	ResetTags()
	ResetTaint()
	ResetWorkloadMetadataConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ContainerClusterNodeConfigOutputReference
type jsiiProxy_ContainerClusterNodeConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) BootDiskKmsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootDiskKmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) BootDiskKmsKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootDiskKmsKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) DiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) DiskSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) DiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) DiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) GcfsConfig() ContainerClusterNodeConfigGcfsConfigOutputReference {
	var returns ContainerClusterNodeConfigGcfsConfigOutputReference
	_jsii_.Get(
		j,
		"gcfsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) GcfsConfigInput() *ContainerClusterNodeConfigGcfsConfig {
	var returns *ContainerClusterNodeConfigGcfsConfig
	_jsii_.Get(
		j,
		"gcfsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) GuestAccelerator() ContainerClusterNodeConfigGuestAcceleratorList {
	var returns ContainerClusterNodeConfigGuestAcceleratorList
	_jsii_.Get(
		j,
		"guestAccelerator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) GuestAcceleratorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"guestAcceleratorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) Gvnic() ContainerClusterNodeConfigGvnicOutputReference {
	var returns ContainerClusterNodeConfigGvnicOutputReference
	_jsii_.Get(
		j,
		"gvnic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) GvnicInput() *ContainerClusterNodeConfigGvnic {
	var returns *ContainerClusterNodeConfigGvnic
	_jsii_.Get(
		j,
		"gvnicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) ImageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) ImageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) InternalValue() *ContainerClusterNodeConfig {
	var returns *ContainerClusterNodeConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) KubeletConfig() ContainerClusterNodeConfigKubeletConfigOutputReference {
	var returns ContainerClusterNodeConfigKubeletConfigOutputReference
	_jsii_.Get(
		j,
		"kubeletConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) KubeletConfigInput() *ContainerClusterNodeConfigKubeletConfig {
	var returns *ContainerClusterNodeConfigKubeletConfig
	_jsii_.Get(
		j,
		"kubeletConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) LinuxNodeConfig() ContainerClusterNodeConfigLinuxNodeConfigOutputReference {
	var returns ContainerClusterNodeConfigLinuxNodeConfigOutputReference
	_jsii_.Get(
		j,
		"linuxNodeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) LinuxNodeConfigInput() *ContainerClusterNodeConfigLinuxNodeConfig {
	var returns *ContainerClusterNodeConfigLinuxNodeConfig
	_jsii_.Get(
		j,
		"linuxNodeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) LocalNvmeSsdBlockConfig() ContainerClusterNodeConfigLocalNvmeSsdBlockConfigOutputReference {
	var returns ContainerClusterNodeConfigLocalNvmeSsdBlockConfigOutputReference
	_jsii_.Get(
		j,
		"localNvmeSsdBlockConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) LocalNvmeSsdBlockConfigInput() *ContainerClusterNodeConfigLocalNvmeSsdBlockConfig {
	var returns *ContainerClusterNodeConfigLocalNvmeSsdBlockConfig
	_jsii_.Get(
		j,
		"localNvmeSsdBlockConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) LocalSsdCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"localSsdCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) LocalSsdCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"localSsdCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) LoggingVariant() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingVariant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) LoggingVariantInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingVariantInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) MachineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) MachineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) Metadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) MetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) MinCpuPlatform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCpuPlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) MinCpuPlatformInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCpuPlatformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) NodeGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) NodeGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) OauthScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"oauthScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) OauthScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"oauthScopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) Preemptible() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preemptible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) PreemptibleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preemptibleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) ReservationAffinity() ContainerClusterNodeConfigReservationAffinityOutputReference {
	var returns ContainerClusterNodeConfigReservationAffinityOutputReference
	_jsii_.Get(
		j,
		"reservationAffinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) ReservationAffinityInput() *ContainerClusterNodeConfigReservationAffinity {
	var returns *ContainerClusterNodeConfigReservationAffinity
	_jsii_.Get(
		j,
		"reservationAffinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) ResourceLabels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"resourceLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) ResourceLabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"resourceLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) ServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) ShieldedInstanceConfig() ContainerClusterNodeConfigShieldedInstanceConfigOutputReference {
	var returns ContainerClusterNodeConfigShieldedInstanceConfigOutputReference
	_jsii_.Get(
		j,
		"shieldedInstanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) ShieldedInstanceConfigInput() *ContainerClusterNodeConfigShieldedInstanceConfig {
	var returns *ContainerClusterNodeConfigShieldedInstanceConfig
	_jsii_.Get(
		j,
		"shieldedInstanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) Spot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"spot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) SpotInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"spotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) Taint() ContainerClusterNodeConfigTaintList {
	var returns ContainerClusterNodeConfigTaintList
	_jsii_.Get(
		j,
		"taint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) TaintInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"taintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) WorkloadMetadataConfig() ContainerClusterNodeConfigWorkloadMetadataConfigOutputReference {
	var returns ContainerClusterNodeConfigWorkloadMetadataConfigOutputReference
	_jsii_.Get(
		j,
		"workloadMetadataConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference) WorkloadMetadataConfigInput() *ContainerClusterNodeConfigWorkloadMetadataConfig {
	var returns *ContainerClusterNodeConfigWorkloadMetadataConfig
	_jsii_.Get(
		j,
		"workloadMetadataConfigInput",
		&returns,
	)
	return returns
}


func NewContainerClusterNodeConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContainerClusterNodeConfigOutputReference {
	_init_.Initialize()

	if err := validateNewContainerClusterNodeConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerClusterNodeConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.containerCluster.ContainerClusterNodeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerClusterNodeConfigOutputReference_Override(c ContainerClusterNodeConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.containerCluster.ContainerClusterNodeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference)SetBootDiskKmsKey(val *string) {
	if err := j.validateSetBootDiskKmsKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootDiskKmsKey",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference)SetDiskSizeGb(val *float64) {
	if err := j.validateSetDiskSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskSizeGb",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference)SetDiskType(val *string) {
	if err := j.validateSetDiskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskType",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference)SetImageType(val *string) {
	if err := j.validateSetImageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageType",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference)SetInternalValue(val *ContainerClusterNodeConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference)SetLocalSsdCount(val *float64) {
	if err := j.validateSetLocalSsdCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localSsdCount",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference)SetLoggingVariant(val *string) {
	if err := j.validateSetLoggingVariantParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loggingVariant",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference)SetMachineType(val *string) {
	if err := j.validateSetMachineTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"machineType",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference)SetMetadata(val *map[string]*string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference)SetMinCpuPlatform(val *string) {
	if err := j.validateSetMinCpuPlatformParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minCpuPlatform",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference)SetNodeGroup(val *string) {
	if err := j.validateSetNodeGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeGroup",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference)SetOauthScopes(val *[]*string) {
	if err := j.validateSetOauthScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthScopes",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference)SetPreemptible(val interface{}) {
	if err := j.validateSetPreemptibleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preemptible",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference)SetResourceLabels(val *map[string]*string) {
	if err := j.validateSetResourceLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceLabels",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference)SetServiceAccount(val *string) {
	if err := j.validateSetServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccount",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference)SetSpot(val interface{}) {
	if err := j.validateSetSpotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spot",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodeConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) PutGcfsConfig(value *ContainerClusterNodeConfigGcfsConfig) {
	if err := c.validatePutGcfsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGcfsConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) PutGuestAccelerator(value interface{}) {
	if err := c.validatePutGuestAcceleratorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGuestAccelerator",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) PutGvnic(value *ContainerClusterNodeConfigGvnic) {
	if err := c.validatePutGvnicParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGvnic",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) PutKubeletConfig(value *ContainerClusterNodeConfigKubeletConfig) {
	if err := c.validatePutKubeletConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putKubeletConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) PutLinuxNodeConfig(value *ContainerClusterNodeConfigLinuxNodeConfig) {
	if err := c.validatePutLinuxNodeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLinuxNodeConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) PutLocalNvmeSsdBlockConfig(value *ContainerClusterNodeConfigLocalNvmeSsdBlockConfig) {
	if err := c.validatePutLocalNvmeSsdBlockConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLocalNvmeSsdBlockConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) PutReservationAffinity(value *ContainerClusterNodeConfigReservationAffinity) {
	if err := c.validatePutReservationAffinityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putReservationAffinity",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) PutShieldedInstanceConfig(value *ContainerClusterNodeConfigShieldedInstanceConfig) {
	if err := c.validatePutShieldedInstanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putShieldedInstanceConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) PutTaint(value interface{}) {
	if err := c.validatePutTaintParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTaint",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) PutWorkloadMetadataConfig(value *ContainerClusterNodeConfigWorkloadMetadataConfig) {
	if err := c.validatePutWorkloadMetadataConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putWorkloadMetadataConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) ResetBootDiskKmsKey() {
	_jsii_.InvokeVoid(
		c,
		"resetBootDiskKmsKey",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) ResetDiskSizeGb() {
	_jsii_.InvokeVoid(
		c,
		"resetDiskSizeGb",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) ResetDiskType() {
	_jsii_.InvokeVoid(
		c,
		"resetDiskType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) ResetGcfsConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetGcfsConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) ResetGuestAccelerator() {
	_jsii_.InvokeVoid(
		c,
		"resetGuestAccelerator",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) ResetGvnic() {
	_jsii_.InvokeVoid(
		c,
		"resetGvnic",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) ResetImageType() {
	_jsii_.InvokeVoid(
		c,
		"resetImageType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) ResetKubeletConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetKubeletConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		c,
		"resetLabels",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) ResetLinuxNodeConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetLinuxNodeConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) ResetLocalNvmeSsdBlockConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetLocalNvmeSsdBlockConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) ResetLocalSsdCount() {
	_jsii_.InvokeVoid(
		c,
		"resetLocalSsdCount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) ResetLoggingVariant() {
	_jsii_.InvokeVoid(
		c,
		"resetLoggingVariant",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) ResetMachineType() {
	_jsii_.InvokeVoid(
		c,
		"resetMachineType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) ResetMetadata() {
	_jsii_.InvokeVoid(
		c,
		"resetMetadata",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) ResetMinCpuPlatform() {
	_jsii_.InvokeVoid(
		c,
		"resetMinCpuPlatform",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) ResetNodeGroup() {
	_jsii_.InvokeVoid(
		c,
		"resetNodeGroup",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) ResetOauthScopes() {
	_jsii_.InvokeVoid(
		c,
		"resetOauthScopes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) ResetPreemptible() {
	_jsii_.InvokeVoid(
		c,
		"resetPreemptible",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) ResetReservationAffinity() {
	_jsii_.InvokeVoid(
		c,
		"resetReservationAffinity",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) ResetResourceLabels() {
	_jsii_.InvokeVoid(
		c,
		"resetResourceLabels",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		c,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) ResetShieldedInstanceConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetShieldedInstanceConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) ResetSpot() {
	_jsii_.InvokeVoid(
		c,
		"resetSpot",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) ResetTaint() {
	_jsii_.InvokeVoid(
		c,
		"resetTaint",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) ResetWorkloadMetadataConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetWorkloadMetadataConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodeConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

