// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notebooksinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v14/notebooksinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.2.0/docs/resources/notebooks_instance google_notebooks_instance}.
type NotebooksInstance interface {
	cdktf.TerraformResource
	AcceleratorConfig() NotebooksInstanceAcceleratorConfigOutputReference
	AcceleratorConfigInput() *NotebooksInstanceAcceleratorConfig
	BootDiskSizeGb() *float64
	SetBootDiskSizeGb(val *float64)
	BootDiskSizeGbInput() *float64
	BootDiskType() *string
	SetBootDiskType(val *string)
	BootDiskTypeInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContainerImage() NotebooksInstanceContainerImageOutputReference
	ContainerImageInput() *NotebooksInstanceContainerImage
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreateTime() *string
	SetCreateTime(val *string)
	CreateTimeInput() *string
	CustomGpuDriverPath() *string
	SetCustomGpuDriverPath(val *string)
	CustomGpuDriverPathInput() *string
	DataDiskSizeGb() *float64
	SetDataDiskSizeGb(val *float64)
	DataDiskSizeGbInput() *float64
	DataDiskType() *string
	SetDataDiskType(val *string)
	DataDiskTypeInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DesiredState() *string
	SetDesiredState(val *string)
	DesiredStateInput() *string
	DiskEncryption() *string
	SetDiskEncryption(val *string)
	DiskEncryptionInput() *string
	EffectiveLabels() cdktf.StringMap
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InstallGpuDriver() interface{}
	SetInstallGpuDriver(val interface{})
	InstallGpuDriverInput() interface{}
	InstanceOwners() *[]*string
	SetInstanceOwners(val *[]*string)
	InstanceOwnersInput() *[]*string
	KmsKey() *string
	SetKmsKey(val *string)
	KmsKeyInput() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MachineType() *string
	SetMachineType(val *string)
	MachineTypeInput() *string
	Metadata() *map[string]*string
	SetMetadata(val *map[string]*string)
	MetadataInput() *map[string]*string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Network() *string
	SetNetwork(val *string)
	NetworkInput() *string
	NicType() *string
	SetNicType(val *string)
	NicTypeInput() *string
	// The tree node.
	Node() constructs.Node
	NoProxyAccess() interface{}
	SetNoProxyAccess(val interface{})
	NoProxyAccessInput() interface{}
	NoPublicIp() interface{}
	SetNoPublicIp(val interface{})
	NoPublicIpInput() interface{}
	NoRemoveDataDisk() interface{}
	SetNoRemoveDataDisk(val interface{})
	NoRemoveDataDiskInput() interface{}
	PostStartupScript() *string
	SetPostStartupScript(val *string)
	PostStartupScriptInput() *string
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	ProxyUri() *string
	// Experimental.
	RawOverrides() interface{}
	ReservationAffinity() NotebooksInstanceReservationAffinityOutputReference
	ReservationAffinityInput() *NotebooksInstanceReservationAffinity
	ServiceAccount() *string
	SetServiceAccount(val *string)
	ServiceAccountInput() *string
	ServiceAccountScopes() *[]*string
	SetServiceAccountScopes(val *[]*string)
	ServiceAccountScopesInput() *[]*string
	ShieldedInstanceConfig() NotebooksInstanceShieldedInstanceConfigOutputReference
	ShieldedInstanceConfigInput() *NotebooksInstanceShieldedInstanceConfig
	State() *string
	Subnet() *string
	SetSubnet(val *string)
	SubnetInput() *string
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() NotebooksInstanceTimeoutsOutputReference
	TimeoutsInput() interface{}
	UpdateTime() *string
	SetUpdateTime(val *string)
	UpdateTimeInput() *string
	VmImage() NotebooksInstanceVmImageOutputReference
	VmImageInput() *NotebooksInstanceVmImage
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAcceleratorConfig(value *NotebooksInstanceAcceleratorConfig)
	PutContainerImage(value *NotebooksInstanceContainerImage)
	PutReservationAffinity(value *NotebooksInstanceReservationAffinity)
	PutShieldedInstanceConfig(value *NotebooksInstanceShieldedInstanceConfig)
	PutTimeouts(value *NotebooksInstanceTimeouts)
	PutVmImage(value *NotebooksInstanceVmImage)
	ResetAcceleratorConfig()
	ResetBootDiskSizeGb()
	ResetBootDiskType()
	ResetContainerImage()
	ResetCreateTime()
	ResetCustomGpuDriverPath()
	ResetDataDiskSizeGb()
	ResetDataDiskType()
	ResetDesiredState()
	ResetDiskEncryption()
	ResetId()
	ResetInstallGpuDriver()
	ResetInstanceOwners()
	ResetKmsKey()
	ResetLabels()
	ResetMetadata()
	ResetNetwork()
	ResetNicType()
	ResetNoProxyAccess()
	ResetNoPublicIp()
	ResetNoRemoveDataDisk()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPostStartupScript()
	ResetProject()
	ResetReservationAffinity()
	ResetServiceAccount()
	ResetServiceAccountScopes()
	ResetShieldedInstanceConfig()
	ResetSubnet()
	ResetTags()
	ResetTimeouts()
	ResetUpdateTime()
	ResetVmImage()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for NotebooksInstance
type jsiiProxy_NotebooksInstance struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_NotebooksInstance) AcceleratorConfig() NotebooksInstanceAcceleratorConfigOutputReference {
	var returns NotebooksInstanceAcceleratorConfigOutputReference
	_jsii_.Get(
		j,
		"acceleratorConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) AcceleratorConfigInput() *NotebooksInstanceAcceleratorConfig {
	var returns *NotebooksInstanceAcceleratorConfig
	_jsii_.Get(
		j,
		"acceleratorConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) BootDiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDiskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) BootDiskSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDiskSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) BootDiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootDiskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) BootDiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootDiskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) ContainerImage() NotebooksInstanceContainerImageOutputReference {
	var returns NotebooksInstanceContainerImageOutputReference
	_jsii_.Get(
		j,
		"containerImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) ContainerImageInput() *NotebooksInstanceContainerImage {
	var returns *NotebooksInstanceContainerImage
	_jsii_.Get(
		j,
		"containerImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) CreateTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) CustomGpuDriverPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customGpuDriverPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) CustomGpuDriverPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customGpuDriverPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) DataDiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataDiskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) DataDiskSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataDiskSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) DataDiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataDiskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) DataDiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataDiskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) DesiredState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) DesiredStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) DiskEncryption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) DiskEncryptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) InstallGpuDriver() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installGpuDriver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) InstallGpuDriverInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installGpuDriverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) InstanceOwners() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceOwners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) InstanceOwnersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceOwnersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) KmsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) KmsKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) MachineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) MachineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) Metadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) MetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) NicType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nicType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) NicTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nicTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) NoProxyAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noProxyAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) NoProxyAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noProxyAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) NoPublicIp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noPublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) NoPublicIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noPublicIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) NoRemoveDataDisk() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noRemoveDataDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) NoRemoveDataDiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noRemoveDataDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) PostStartupScript() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postStartupScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) PostStartupScriptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postStartupScriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) ProxyUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proxyUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) ReservationAffinity() NotebooksInstanceReservationAffinityOutputReference {
	var returns NotebooksInstanceReservationAffinityOutputReference
	_jsii_.Get(
		j,
		"reservationAffinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) ReservationAffinityInput() *NotebooksInstanceReservationAffinity {
	var returns *NotebooksInstanceReservationAffinity
	_jsii_.Get(
		j,
		"reservationAffinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) ServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) ServiceAccountScopes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serviceAccountScopes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) ServiceAccountScopesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serviceAccountScopesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) ShieldedInstanceConfig() NotebooksInstanceShieldedInstanceConfigOutputReference {
	var returns NotebooksInstanceShieldedInstanceConfigOutputReference
	_jsii_.Get(
		j,
		"shieldedInstanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) ShieldedInstanceConfigInput() *NotebooksInstanceShieldedInstanceConfig {
	var returns *NotebooksInstanceShieldedInstanceConfig
	_jsii_.Get(
		j,
		"shieldedInstanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) Subnet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) SubnetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) Timeouts() NotebooksInstanceTimeoutsOutputReference {
	var returns NotebooksInstanceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) UpdateTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) VmImage() NotebooksInstanceVmImageOutputReference {
	var returns NotebooksInstanceVmImageOutputReference
	_jsii_.Get(
		j,
		"vmImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NotebooksInstance) VmImageInput() *NotebooksInstanceVmImage {
	var returns *NotebooksInstanceVmImage
	_jsii_.Get(
		j,
		"vmImageInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.2.0/docs/resources/notebooks_instance google_notebooks_instance} Resource.
func NewNotebooksInstance(scope constructs.Construct, id *string, config *NotebooksInstanceConfig) NotebooksInstance {
	_init_.Initialize()

	if err := validateNewNotebooksInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_NotebooksInstance{}

	_jsii_.Create(
		"@cdktf/provider-google.notebooksInstance.NotebooksInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.2.0/docs/resources/notebooks_instance google_notebooks_instance} Resource.
func NewNotebooksInstance_Override(n NotebooksInstance, scope constructs.Construct, id *string, config *NotebooksInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.notebooksInstance.NotebooksInstance",
		[]interface{}{scope, id, config},
		n,
	)
}

func (j *jsiiProxy_NotebooksInstance)SetBootDiskSizeGb(val *float64) {
	if err := j.validateSetBootDiskSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootDiskSizeGb",
		val,
	)
}

func (j *jsiiProxy_NotebooksInstance)SetBootDiskType(val *string) {
	if err := j.validateSetBootDiskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bootDiskType",
		val,
	)
}

func (j *jsiiProxy_NotebooksInstance)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_NotebooksInstance)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_NotebooksInstance)SetCreateTime(val *string) {
	if err := j.validateSetCreateTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createTime",
		val,
	)
}

func (j *jsiiProxy_NotebooksInstance)SetCustomGpuDriverPath(val *string) {
	if err := j.validateSetCustomGpuDriverPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customGpuDriverPath",
		val,
	)
}

func (j *jsiiProxy_NotebooksInstance)SetDataDiskSizeGb(val *float64) {
	if err := j.validateSetDataDiskSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataDiskSizeGb",
		val,
	)
}

func (j *jsiiProxy_NotebooksInstance)SetDataDiskType(val *string) {
	if err := j.validateSetDataDiskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataDiskType",
		val,
	)
}

func (j *jsiiProxy_NotebooksInstance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_NotebooksInstance)SetDesiredState(val *string) {
	if err := j.validateSetDesiredStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredState",
		val,
	)
}

func (j *jsiiProxy_NotebooksInstance)SetDiskEncryption(val *string) {
	if err := j.validateSetDiskEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskEncryption",
		val,
	)
}

func (j *jsiiProxy_NotebooksInstance)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_NotebooksInstance)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_NotebooksInstance)SetInstallGpuDriver(val interface{}) {
	if err := j.validateSetInstallGpuDriverParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"installGpuDriver",
		val,
	)
}

func (j *jsiiProxy_NotebooksInstance)SetInstanceOwners(val *[]*string) {
	if err := j.validateSetInstanceOwnersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceOwners",
		val,
	)
}

func (j *jsiiProxy_NotebooksInstance)SetKmsKey(val *string) {
	if err := j.validateSetKmsKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKey",
		val,
	)
}

func (j *jsiiProxy_NotebooksInstance)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_NotebooksInstance)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_NotebooksInstance)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_NotebooksInstance)SetMachineType(val *string) {
	if err := j.validateSetMachineTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"machineType",
		val,
	)
}

func (j *jsiiProxy_NotebooksInstance)SetMetadata(val *map[string]*string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_NotebooksInstance)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_NotebooksInstance)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_NotebooksInstance)SetNicType(val *string) {
	if err := j.validateSetNicTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nicType",
		val,
	)
}

func (j *jsiiProxy_NotebooksInstance)SetNoProxyAccess(val interface{}) {
	if err := j.validateSetNoProxyAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noProxyAccess",
		val,
	)
}

func (j *jsiiProxy_NotebooksInstance)SetNoPublicIp(val interface{}) {
	if err := j.validateSetNoPublicIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noPublicIp",
		val,
	)
}

func (j *jsiiProxy_NotebooksInstance)SetNoRemoveDataDisk(val interface{}) {
	if err := j.validateSetNoRemoveDataDiskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noRemoveDataDisk",
		val,
	)
}

func (j *jsiiProxy_NotebooksInstance)SetPostStartupScript(val *string) {
	if err := j.validateSetPostStartupScriptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postStartupScript",
		val,
	)
}

func (j *jsiiProxy_NotebooksInstance)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_NotebooksInstance)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_NotebooksInstance)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_NotebooksInstance)SetServiceAccount(val *string) {
	if err := j.validateSetServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccount",
		val,
	)
}

func (j *jsiiProxy_NotebooksInstance)SetServiceAccountScopes(val *[]*string) {
	if err := j.validateSetServiceAccountScopesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccountScopes",
		val,
	)
}

func (j *jsiiProxy_NotebooksInstance)SetSubnet(val *string) {
	if err := j.validateSetSubnetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnet",
		val,
	)
}

func (j *jsiiProxy_NotebooksInstance)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_NotebooksInstance)SetUpdateTime(val *string) {
	if err := j.validateSetUpdateTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateTime",
		val,
	)
}

// Generates CDKTF code for importing a NotebooksInstance resource upon running "cdktf plan <stack-name>".
func NotebooksInstance_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateNotebooksInstance_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.notebooksInstance.NotebooksInstance",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func NotebooksInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNotebooksInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.notebooksInstance.NotebooksInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NotebooksInstance_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNotebooksInstance_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.notebooksInstance.NotebooksInstance",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func NotebooksInstance_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateNotebooksInstance_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.notebooksInstance.NotebooksInstance",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func NotebooksInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.notebooksInstance.NotebooksInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (n *jsiiProxy_NotebooksInstance) AddMoveTarget(moveTarget *string) {
	if err := n.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (n *jsiiProxy_NotebooksInstance) AddOverride(path *string, value interface{}) {
	if err := n.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (n *jsiiProxy_NotebooksInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksInstance) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksInstance) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksInstance) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksInstance) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksInstance) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := n.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (n *jsiiProxy_NotebooksInstance) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksInstance) MoveFromId(id *string) {
	if err := n.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveFromId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NotebooksInstance) MoveTo(moveTarget *string, index interface{}) {
	if err := n.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (n *jsiiProxy_NotebooksInstance) MoveToId(id *string) {
	if err := n.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"moveToId",
		[]interface{}{id},
	)
}

func (n *jsiiProxy_NotebooksInstance) OverrideLogicalId(newLogicalId *string) {
	if err := n.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (n *jsiiProxy_NotebooksInstance) PutAcceleratorConfig(value *NotebooksInstanceAcceleratorConfig) {
	if err := n.validatePutAcceleratorConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putAcceleratorConfig",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NotebooksInstance) PutContainerImage(value *NotebooksInstanceContainerImage) {
	if err := n.validatePutContainerImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putContainerImage",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NotebooksInstance) PutReservationAffinity(value *NotebooksInstanceReservationAffinity) {
	if err := n.validatePutReservationAffinityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putReservationAffinity",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NotebooksInstance) PutShieldedInstanceConfig(value *NotebooksInstanceShieldedInstanceConfig) {
	if err := n.validatePutShieldedInstanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putShieldedInstanceConfig",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NotebooksInstance) PutTimeouts(value *NotebooksInstanceTimeouts) {
	if err := n.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NotebooksInstance) PutVmImage(value *NotebooksInstanceVmImage) {
	if err := n.validatePutVmImageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"putVmImage",
		[]interface{}{value},
	)
}

func (n *jsiiProxy_NotebooksInstance) ResetAcceleratorConfig() {
	_jsii_.InvokeVoid(
		n,
		"resetAcceleratorConfig",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksInstance) ResetBootDiskSizeGb() {
	_jsii_.InvokeVoid(
		n,
		"resetBootDiskSizeGb",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksInstance) ResetBootDiskType() {
	_jsii_.InvokeVoid(
		n,
		"resetBootDiskType",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksInstance) ResetContainerImage() {
	_jsii_.InvokeVoid(
		n,
		"resetContainerImage",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksInstance) ResetCreateTime() {
	_jsii_.InvokeVoid(
		n,
		"resetCreateTime",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksInstance) ResetCustomGpuDriverPath() {
	_jsii_.InvokeVoid(
		n,
		"resetCustomGpuDriverPath",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksInstance) ResetDataDiskSizeGb() {
	_jsii_.InvokeVoid(
		n,
		"resetDataDiskSizeGb",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksInstance) ResetDataDiskType() {
	_jsii_.InvokeVoid(
		n,
		"resetDataDiskType",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksInstance) ResetDesiredState() {
	_jsii_.InvokeVoid(
		n,
		"resetDesiredState",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksInstance) ResetDiskEncryption() {
	_jsii_.InvokeVoid(
		n,
		"resetDiskEncryption",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksInstance) ResetId() {
	_jsii_.InvokeVoid(
		n,
		"resetId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksInstance) ResetInstallGpuDriver() {
	_jsii_.InvokeVoid(
		n,
		"resetInstallGpuDriver",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksInstance) ResetInstanceOwners() {
	_jsii_.InvokeVoid(
		n,
		"resetInstanceOwners",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksInstance) ResetKmsKey() {
	_jsii_.InvokeVoid(
		n,
		"resetKmsKey",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksInstance) ResetLabels() {
	_jsii_.InvokeVoid(
		n,
		"resetLabels",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksInstance) ResetMetadata() {
	_jsii_.InvokeVoid(
		n,
		"resetMetadata",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksInstance) ResetNetwork() {
	_jsii_.InvokeVoid(
		n,
		"resetNetwork",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksInstance) ResetNicType() {
	_jsii_.InvokeVoid(
		n,
		"resetNicType",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksInstance) ResetNoProxyAccess() {
	_jsii_.InvokeVoid(
		n,
		"resetNoProxyAccess",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksInstance) ResetNoPublicIp() {
	_jsii_.InvokeVoid(
		n,
		"resetNoPublicIp",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksInstance) ResetNoRemoveDataDisk() {
	_jsii_.InvokeVoid(
		n,
		"resetNoRemoveDataDisk",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		n,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksInstance) ResetPostStartupScript() {
	_jsii_.InvokeVoid(
		n,
		"resetPostStartupScript",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksInstance) ResetProject() {
	_jsii_.InvokeVoid(
		n,
		"resetProject",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksInstance) ResetReservationAffinity() {
	_jsii_.InvokeVoid(
		n,
		"resetReservationAffinity",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksInstance) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		n,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksInstance) ResetServiceAccountScopes() {
	_jsii_.InvokeVoid(
		n,
		"resetServiceAccountScopes",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksInstance) ResetShieldedInstanceConfig() {
	_jsii_.InvokeVoid(
		n,
		"resetShieldedInstanceConfig",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksInstance) ResetSubnet() {
	_jsii_.InvokeVoid(
		n,
		"resetSubnet",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksInstance) ResetTags() {
	_jsii_.InvokeVoid(
		n,
		"resetTags",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksInstance) ResetTimeouts() {
	_jsii_.InvokeVoid(
		n,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksInstance) ResetUpdateTime() {
	_jsii_.InvokeVoid(
		n,
		"resetUpdateTime",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksInstance) ResetVmImage() {
	_jsii_.InvokeVoid(
		n,
		"resetVmImage",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NotebooksInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksInstance) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksInstance) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NotebooksInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

