package computedisk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v8/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v8/computedisk/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_disk google_compute_disk}.
type ComputeDisk interface {
	cdktf.TerraformResource
	AsyncPrimaryDisk() ComputeDiskAsyncPrimaryDiskOutputReference
	AsyncPrimaryDiskInput() *ComputeDiskAsyncPrimaryDisk
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreationTimestamp() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DiskEncryptionKey() ComputeDiskDiskEncryptionKeyOutputReference
	DiskEncryptionKeyInput() *ComputeDiskDiskEncryptionKey
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GuestOsFeatures() ComputeDiskGuestOsFeaturesList
	GuestOsFeaturesInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	Image() *string
	SetImage(val *string)
	ImageInput() *string
	LabelFingerprint() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	LastAttachTimestamp() *string
	LastDetachTimestamp() *string
	Licenses() *[]*string
	SetLicenses(val *[]*string)
	LicensesInput() *[]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PhysicalBlockSizeBytes() *float64
	SetPhysicalBlockSizeBytes(val *float64)
	PhysicalBlockSizeBytesInput() *float64
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	ProvisionedIops() *float64
	SetProvisionedIops(val *float64)
	ProvisionedIopsInput() *float64
	ProvisionedThroughput() *float64
	SetProvisionedThroughput(val *float64)
	ProvisionedThroughputInput() *float64
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	SelfLink() *string
	Size() *float64
	SetSize(val *float64)
	SizeInput() *float64
	Snapshot() *string
	SetSnapshot(val *string)
	SnapshotInput() *string
	SourceDisk() *string
	SetSourceDisk(val *string)
	SourceDiskId() *string
	SourceDiskInput() *string
	SourceImageEncryptionKey() ComputeDiskSourceImageEncryptionKeyOutputReference
	SourceImageEncryptionKeyInput() *ComputeDiskSourceImageEncryptionKey
	SourceImageId() *string
	SourceSnapshotEncryptionKey() ComputeDiskSourceSnapshotEncryptionKeyOutputReference
	SourceSnapshotEncryptionKeyInput() *ComputeDiskSourceSnapshotEncryptionKey
	SourceSnapshotId() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ComputeDiskTimeoutsOutputReference
	TimeoutsInput() interface{}
	Type() *string
	SetType(val *string)
	TypeInput() *string
	Users() *[]*string
	Zone() *string
	SetZone(val *string)
	ZoneInput() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAsyncPrimaryDisk(value *ComputeDiskAsyncPrimaryDisk)
	PutDiskEncryptionKey(value *ComputeDiskDiskEncryptionKey)
	PutGuestOsFeatures(value interface{})
	PutSourceImageEncryptionKey(value *ComputeDiskSourceImageEncryptionKey)
	PutSourceSnapshotEncryptionKey(value *ComputeDiskSourceSnapshotEncryptionKey)
	PutTimeouts(value *ComputeDiskTimeouts)
	ResetAsyncPrimaryDisk()
	ResetDescription()
	ResetDiskEncryptionKey()
	ResetGuestOsFeatures()
	ResetId()
	ResetImage()
	ResetLabels()
	ResetLicenses()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPhysicalBlockSizeBytes()
	ResetProject()
	ResetProvisionedIops()
	ResetProvisionedThroughput()
	ResetSize()
	ResetSnapshot()
	ResetSourceDisk()
	ResetSourceImageEncryptionKey()
	ResetSourceSnapshotEncryptionKey()
	ResetTimeouts()
	ResetType()
	ResetZone()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ComputeDisk
type jsiiProxy_ComputeDisk struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ComputeDisk) AsyncPrimaryDisk() ComputeDiskAsyncPrimaryDiskOutputReference {
	var returns ComputeDiskAsyncPrimaryDiskOutputReference
	_jsii_.Get(
		j,
		"asyncPrimaryDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) AsyncPrimaryDiskInput() *ComputeDiskAsyncPrimaryDisk {
	var returns *ComputeDiskAsyncPrimaryDisk
	_jsii_.Get(
		j,
		"asyncPrimaryDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) CreationTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) DiskEncryptionKey() ComputeDiskDiskEncryptionKeyOutputReference {
	var returns ComputeDiskDiskEncryptionKeyOutputReference
	_jsii_.Get(
		j,
		"diskEncryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) DiskEncryptionKeyInput() *ComputeDiskDiskEncryptionKey {
	var returns *ComputeDiskDiskEncryptionKey
	_jsii_.Get(
		j,
		"diskEncryptionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) GuestOsFeatures() ComputeDiskGuestOsFeaturesList {
	var returns ComputeDiskGuestOsFeaturesList
	_jsii_.Get(
		j,
		"guestOsFeatures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) GuestOsFeaturesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"guestOsFeaturesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) LabelFingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelFingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) LastAttachTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastAttachTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) LastDetachTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastDetachTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) Licenses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"licenses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) LicensesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"licensesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) PhysicalBlockSizeBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"physicalBlockSizeBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) PhysicalBlockSizeBytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"physicalBlockSizeBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) ProvisionedIops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedIops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) ProvisionedIopsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedIopsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) ProvisionedThroughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) ProvisionedThroughputInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) Size() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"size",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) SizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) Snapshot() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) SnapshotInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) SourceDisk() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) SourceDiskId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDiskId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) SourceDiskInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) SourceImageEncryptionKey() ComputeDiskSourceImageEncryptionKeyOutputReference {
	var returns ComputeDiskSourceImageEncryptionKeyOutputReference
	_jsii_.Get(
		j,
		"sourceImageEncryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) SourceImageEncryptionKeyInput() *ComputeDiskSourceImageEncryptionKey {
	var returns *ComputeDiskSourceImageEncryptionKey
	_jsii_.Get(
		j,
		"sourceImageEncryptionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) SourceImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceImageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) SourceSnapshotEncryptionKey() ComputeDiskSourceSnapshotEncryptionKeyOutputReference {
	var returns ComputeDiskSourceSnapshotEncryptionKeyOutputReference
	_jsii_.Get(
		j,
		"sourceSnapshotEncryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) SourceSnapshotEncryptionKeyInput() *ComputeDiskSourceSnapshotEncryptionKey {
	var returns *ComputeDiskSourceSnapshotEncryptionKey
	_jsii_.Get(
		j,
		"sourceSnapshotEncryptionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) SourceSnapshotId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceSnapshotId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) Timeouts() ComputeDiskTimeoutsOutputReference {
	var returns ComputeDiskTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) Users() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"users",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeDisk) ZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_disk google_compute_disk} Resource.
func NewComputeDisk(scope constructs.Construct, id *string, config *ComputeDiskConfig) ComputeDisk {
	_init_.Initialize()

	if err := validateNewComputeDiskParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeDisk{}

	_jsii_.Create(
		"@cdktf/provider-google.computeDisk.ComputeDisk",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_disk google_compute_disk} Resource.
func NewComputeDisk_Override(c ComputeDisk, scope constructs.Construct, id *string, config *ComputeDiskConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeDisk.ComputeDisk",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ComputeDisk)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ComputeDisk)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ComputeDisk)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ComputeDisk)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ComputeDisk)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ComputeDisk)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ComputeDisk)SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_ComputeDisk)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_ComputeDisk)SetLicenses(val *[]*string) {
	if err := j.validateSetLicensesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenses",
		val,
	)
}

func (j *jsiiProxy_ComputeDisk)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ComputeDisk)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ComputeDisk)SetPhysicalBlockSizeBytes(val *float64) {
	if err := j.validateSetPhysicalBlockSizeBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"physicalBlockSizeBytes",
		val,
	)
}

func (j *jsiiProxy_ComputeDisk)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_ComputeDisk)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ComputeDisk)SetProvisionedIops(val *float64) {
	if err := j.validateSetProvisionedIopsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisionedIops",
		val,
	)
}

func (j *jsiiProxy_ComputeDisk)SetProvisionedThroughput(val *float64) {
	if err := j.validateSetProvisionedThroughputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisionedThroughput",
		val,
	)
}

func (j *jsiiProxy_ComputeDisk)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ComputeDisk)SetSize(val *float64) {
	if err := j.validateSetSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"size",
		val,
	)
}

func (j *jsiiProxy_ComputeDisk)SetSnapshot(val *string) {
	if err := j.validateSetSnapshotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshot",
		val,
	)
}

func (j *jsiiProxy_ComputeDisk)SetSourceDisk(val *string) {
	if err := j.validateSetSourceDiskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDisk",
		val,
	)
}

func (j *jsiiProxy_ComputeDisk)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_ComputeDisk)SetZone(val *string) {
	if err := j.validateSetZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zone",
		val,
	)
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
func ComputeDisk_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeDisk_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeDisk.ComputeDisk",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeDisk_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeDisk_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeDisk.ComputeDisk",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeDisk_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeDisk_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeDisk.ComputeDisk",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ComputeDisk_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.computeDisk.ComputeDisk",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ComputeDisk) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ComputeDisk) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeDisk) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeDisk) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeDisk) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeDisk) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeDisk) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeDisk) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeDisk) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeDisk) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeDisk) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeDisk) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ComputeDisk) PutAsyncPrimaryDisk(value *ComputeDiskAsyncPrimaryDisk) {
	if err := c.validatePutAsyncPrimaryDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAsyncPrimaryDisk",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeDisk) PutDiskEncryptionKey(value *ComputeDiskDiskEncryptionKey) {
	if err := c.validatePutDiskEncryptionKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDiskEncryptionKey",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeDisk) PutGuestOsFeatures(value interface{}) {
	if err := c.validatePutGuestOsFeaturesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGuestOsFeatures",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeDisk) PutSourceImageEncryptionKey(value *ComputeDiskSourceImageEncryptionKey) {
	if err := c.validatePutSourceImageEncryptionKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSourceImageEncryptionKey",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeDisk) PutSourceSnapshotEncryptionKey(value *ComputeDiskSourceSnapshotEncryptionKey) {
	if err := c.validatePutSourceSnapshotEncryptionKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSourceSnapshotEncryptionKey",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeDisk) PutTimeouts(value *ComputeDiskTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeDisk) ResetAsyncPrimaryDisk() {
	_jsii_.InvokeVoid(
		c,
		"resetAsyncPrimaryDisk",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeDisk) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeDisk) ResetDiskEncryptionKey() {
	_jsii_.InvokeVoid(
		c,
		"resetDiskEncryptionKey",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeDisk) ResetGuestOsFeatures() {
	_jsii_.InvokeVoid(
		c,
		"resetGuestOsFeatures",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeDisk) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeDisk) ResetImage() {
	_jsii_.InvokeVoid(
		c,
		"resetImage",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeDisk) ResetLabels() {
	_jsii_.InvokeVoid(
		c,
		"resetLabels",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeDisk) ResetLicenses() {
	_jsii_.InvokeVoid(
		c,
		"resetLicenses",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeDisk) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeDisk) ResetPhysicalBlockSizeBytes() {
	_jsii_.InvokeVoid(
		c,
		"resetPhysicalBlockSizeBytes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeDisk) ResetProject() {
	_jsii_.InvokeVoid(
		c,
		"resetProject",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeDisk) ResetProvisionedIops() {
	_jsii_.InvokeVoid(
		c,
		"resetProvisionedIops",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeDisk) ResetProvisionedThroughput() {
	_jsii_.InvokeVoid(
		c,
		"resetProvisionedThroughput",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeDisk) ResetSize() {
	_jsii_.InvokeVoid(
		c,
		"resetSize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeDisk) ResetSnapshot() {
	_jsii_.InvokeVoid(
		c,
		"resetSnapshot",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeDisk) ResetSourceDisk() {
	_jsii_.InvokeVoid(
		c,
		"resetSourceDisk",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeDisk) ResetSourceImageEncryptionKey() {
	_jsii_.InvokeVoid(
		c,
		"resetSourceImageEncryptionKey",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeDisk) ResetSourceSnapshotEncryptionKey() {
	_jsii_.InvokeVoid(
		c,
		"resetSourceSnapshotEncryptionKey",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeDisk) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeDisk) ResetType() {
	_jsii_.InvokeVoid(
		c,
		"resetType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeDisk) ResetZone() {
	_jsii_.InvokeVoid(
		c,
		"resetZone",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeDisk) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeDisk) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeDisk) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeDisk) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

