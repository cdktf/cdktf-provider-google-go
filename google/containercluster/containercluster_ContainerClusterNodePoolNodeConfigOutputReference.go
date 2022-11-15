package containercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v4/jsii"

	"github.com/cdktf/cdktf-provider-google-go/google/v4/containercluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerClusterNodePoolNodeConfigOutputReference interface {
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
	GcfsConfig() ContainerClusterNodePoolNodeConfigGcfsConfigOutputReference
	GcfsConfigInput() *ContainerClusterNodePoolNodeConfigGcfsConfig
	GuestAccelerator() ContainerClusterNodePoolNodeConfigGuestAcceleratorList
	GuestAcceleratorInput() interface{}
	Gvnic() ContainerClusterNodePoolNodeConfigGvnicOutputReference
	GvnicInput() *ContainerClusterNodePoolNodeConfigGvnic
	ImageType() *string
	SetImageType(val *string)
	ImageTypeInput() *string
	InternalValue() *ContainerClusterNodePoolNodeConfig
	SetInternalValue(val *ContainerClusterNodePoolNodeConfig)
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	LocalSsdCount() *float64
	SetLocalSsdCount(val *float64)
	LocalSsdCountInput() *float64
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
	ReservationAffinity() ContainerClusterNodePoolNodeConfigReservationAffinityOutputReference
	ReservationAffinityInput() *ContainerClusterNodePoolNodeConfigReservationAffinity
	ServiceAccount() *string
	SetServiceAccount(val *string)
	ServiceAccountInput() *string
	ShieldedInstanceConfig() ContainerClusterNodePoolNodeConfigShieldedInstanceConfigOutputReference
	ShieldedInstanceConfigInput() *ContainerClusterNodePoolNodeConfigShieldedInstanceConfig
	Spot() interface{}
	SetSpot(val interface{})
	SpotInput() interface{}
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	Taint() ContainerClusterNodePoolNodeConfigTaintList
	TaintInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WorkloadMetadataConfig() ContainerClusterNodePoolNodeConfigWorkloadMetadataConfigOutputReference
	WorkloadMetadataConfigInput() *ContainerClusterNodePoolNodeConfigWorkloadMetadataConfig
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
	PutGcfsConfig(value *ContainerClusterNodePoolNodeConfigGcfsConfig)
	PutGuestAccelerator(value interface{})
	PutGvnic(value *ContainerClusterNodePoolNodeConfigGvnic)
	PutReservationAffinity(value *ContainerClusterNodePoolNodeConfigReservationAffinity)
	PutShieldedInstanceConfig(value *ContainerClusterNodePoolNodeConfigShieldedInstanceConfig)
	PutTaint(value interface{})
	PutWorkloadMetadataConfig(value *ContainerClusterNodePoolNodeConfigWorkloadMetadataConfig)
	ResetBootDiskKmsKey()
	ResetDiskSizeGb()
	ResetDiskType()
	ResetGcfsConfig()
	ResetGuestAccelerator()
	ResetGvnic()
	ResetImageType()
	ResetLabels()
	ResetLocalSsdCount()
	ResetMachineType()
	ResetMetadata()
	ResetMinCpuPlatform()
	ResetNodeGroup()
	ResetOauthScopes()
	ResetPreemptible()
	ResetReservationAffinity()
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

// The jsii proxy struct for ContainerClusterNodePoolNodeConfigOutputReference
type jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) BootDiskKmsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootDiskKmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) BootDiskKmsKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootDiskKmsKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) DiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) DiskSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) DiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) DiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) GcfsConfig() ContainerClusterNodePoolNodeConfigGcfsConfigOutputReference {
	var returns ContainerClusterNodePoolNodeConfigGcfsConfigOutputReference
	_jsii_.Get(
		j,
		"gcfsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) GcfsConfigInput() *ContainerClusterNodePoolNodeConfigGcfsConfig {
	var returns *ContainerClusterNodePoolNodeConfigGcfsConfig
	_jsii_.Get(
		j,
		"gcfsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) GuestAccelerator() ContainerClusterNodePoolNodeConfigGuestAcceleratorList {
	var returns ContainerClusterNodePoolNodeConfigGuestAcceleratorList
	_jsii_.Get(
		j,
		"guestAccelerator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) GuestAcceleratorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"guestAcceleratorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) Gvnic() ContainerClusterNodePoolNodeConfigGvnicOutputReference {
	var returns ContainerClusterNodePoolNodeConfigGvnicOutputReference
	_jsii_.Get(
		j,
		"gvnic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) GvnicInput() *ContainerClusterNodePoolNodeConfigGvnic {
	var returns *ContainerClusterNodePoolNodeConfigGvnic
	_jsii_.Get(
		j,
		"gvnicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) ImageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) ImageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) InternalValue() *ContainerClusterNodePoolNodeConfig {
	var returns *ContainerClusterNodePoolNodeConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) LocalSsdCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"localSsdCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) LocalSsdCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"localSsdCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) MachineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) MachineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) Metadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) MetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) MinCpuPlatform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCpuPlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) MinCpuPlatformInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCpuPlatformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) NodeGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) NodeGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) OauthScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"oauthScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) OauthScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"oauthScopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) Preemptible() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preemptible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) PreemptibleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preemptibleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) ReservationAffinity() ContainerClusterNodePoolNodeConfigReservationAffinityOutputReference {
	var returns ContainerClusterNodePoolNodeConfigReservationAffinityOutputReference
	_jsii_.Get(
		j,
		"reservationAffinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) ReservationAffinityInput() *ContainerClusterNodePoolNodeConfigReservationAffinity {
	var returns *ContainerClusterNodePoolNodeConfigReservationAffinity
	_jsii_.Get(
		j,
		"reservationAffinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) ServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) ShieldedInstanceConfig() ContainerClusterNodePoolNodeConfigShieldedInstanceConfigOutputReference {
	var returns ContainerClusterNodePoolNodeConfigShieldedInstanceConfigOutputReference
	_jsii_.Get(
		j,
		"shieldedInstanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) ShieldedInstanceConfigInput() *ContainerClusterNodePoolNodeConfigShieldedInstanceConfig {
	var returns *ContainerClusterNodePoolNodeConfigShieldedInstanceConfig
	_jsii_.Get(
		j,
		"shieldedInstanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) Spot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"spot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) SpotInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"spotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) Taint() ContainerClusterNodePoolNodeConfigTaintList {
	var returns ContainerClusterNodePoolNodeConfigTaintList
	_jsii_.Get(
		j,
		"taint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) TaintInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"taintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) WorkloadMetadataConfig() ContainerClusterNodePoolNodeConfigWorkloadMetadataConfigOutputReference {
	var returns ContainerClusterNodePoolNodeConfigWorkloadMetadataConfigOutputReference
	_jsii_.Get(
		j,
		"workloadMetadataConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) WorkloadMetadataConfigInput() *ContainerClusterNodePoolNodeConfigWorkloadMetadataConfig {
	var returns *ContainerClusterNodePoolNodeConfigWorkloadMetadataConfig
	_jsii_.Get(
		j,
		"workloadMetadataConfigInput",
		&returns,
	)
	return returns
}


func NewContainerClusterNodePoolNodeConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ContainerClusterNodePoolNodeConfigOutputReference {
	_init_.Initialize()

	if err := validateNewContainerClusterNodePoolNodeConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.containerCluster.ContainerClusterNodePoolNodeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewContainerClusterNodePoolNodeConfigOutputReference_Override(c ContainerClusterNodePoolNodeConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.containerCluster.ContainerClusterNodePoolNodeConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference)SetBootDiskKmsKey(val *string) {
	if err := j.validateSetBootDiskKmsKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootDiskKmsKey",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference)SetDiskSizeGb(val *float64) {
	if err := j.validateSetDiskSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskSizeGb",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference)SetDiskType(val *string) {
	if err := j.validateSetDiskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskType",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference)SetImageType(val *string) {
	if err := j.validateSetImageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageType",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference)SetInternalValue(val *ContainerClusterNodePoolNodeConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference)SetLocalSsdCount(val *float64) {
	if err := j.validateSetLocalSsdCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localSsdCount",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference)SetMachineType(val *string) {
	if err := j.validateSetMachineTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"machineType",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference)SetMetadata(val *map[string]*string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference)SetMinCpuPlatform(val *string) {
	if err := j.validateSetMinCpuPlatformParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minCpuPlatform",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference)SetNodeGroup(val *string) {
	if err := j.validateSetNodeGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeGroup",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference)SetOauthScopes(val *[]*string) {
	if err := j.validateSetOauthScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauthScopes",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference)SetPreemptible(val interface{}) {
	if err := j.validateSetPreemptibleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preemptible",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference)SetServiceAccount(val *string) {
	if err := j.validateSetServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccount",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference)SetSpot(val interface{}) {
	if err := j.validateSetSpotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spot",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) PutGcfsConfig(value *ContainerClusterNodePoolNodeConfigGcfsConfig) {
	if err := c.validatePutGcfsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGcfsConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) PutGuestAccelerator(value interface{}) {
	if err := c.validatePutGuestAcceleratorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGuestAccelerator",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) PutGvnic(value *ContainerClusterNodePoolNodeConfigGvnic) {
	if err := c.validatePutGvnicParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGvnic",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) PutReservationAffinity(value *ContainerClusterNodePoolNodeConfigReservationAffinity) {
	if err := c.validatePutReservationAffinityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putReservationAffinity",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) PutShieldedInstanceConfig(value *ContainerClusterNodePoolNodeConfigShieldedInstanceConfig) {
	if err := c.validatePutShieldedInstanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putShieldedInstanceConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) PutTaint(value interface{}) {
	if err := c.validatePutTaintParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTaint",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) PutWorkloadMetadataConfig(value *ContainerClusterNodePoolNodeConfigWorkloadMetadataConfig) {
	if err := c.validatePutWorkloadMetadataConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putWorkloadMetadataConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) ResetBootDiskKmsKey() {
	_jsii_.InvokeVoid(
		c,
		"resetBootDiskKmsKey",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) ResetDiskSizeGb() {
	_jsii_.InvokeVoid(
		c,
		"resetDiskSizeGb",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) ResetDiskType() {
	_jsii_.InvokeVoid(
		c,
		"resetDiskType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) ResetGcfsConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetGcfsConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) ResetGuestAccelerator() {
	_jsii_.InvokeVoid(
		c,
		"resetGuestAccelerator",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) ResetGvnic() {
	_jsii_.InvokeVoid(
		c,
		"resetGvnic",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) ResetImageType() {
	_jsii_.InvokeVoid(
		c,
		"resetImageType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		c,
		"resetLabels",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) ResetLocalSsdCount() {
	_jsii_.InvokeVoid(
		c,
		"resetLocalSsdCount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) ResetMachineType() {
	_jsii_.InvokeVoid(
		c,
		"resetMachineType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) ResetMetadata() {
	_jsii_.InvokeVoid(
		c,
		"resetMetadata",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) ResetMinCpuPlatform() {
	_jsii_.InvokeVoid(
		c,
		"resetMinCpuPlatform",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) ResetNodeGroup() {
	_jsii_.InvokeVoid(
		c,
		"resetNodeGroup",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) ResetOauthScopes() {
	_jsii_.InvokeVoid(
		c,
		"resetOauthScopes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) ResetPreemptible() {
	_jsii_.InvokeVoid(
		c,
		"resetPreemptible",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) ResetReservationAffinity() {
	_jsii_.InvokeVoid(
		c,
		"resetReservationAffinity",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		c,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) ResetShieldedInstanceConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetShieldedInstanceConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) ResetSpot() {
	_jsii_.InvokeVoid(
		c,
		"resetSpot",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) ResetTaint() {
	_jsii_.InvokeVoid(
		c,
		"resetTaint",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) ResetWorkloadMetadataConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetWorkloadMetadataConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_ContainerClusterNodePoolNodeConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

