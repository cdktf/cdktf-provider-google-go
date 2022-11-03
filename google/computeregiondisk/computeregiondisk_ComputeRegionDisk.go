package computeregiondisk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v3/computeregiondisk/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/google/r/compute_region_disk google_compute_region_disk}.
type ComputeRegionDisk interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	CreationTimestamp() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DiskEncryptionKey() ComputeRegionDiskDiskEncryptionKeyOutputReference
	DiskEncryptionKeyInput() *ComputeRegionDiskDiskEncryptionKey
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
	LabelFingerprint() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	LastAttachTimestamp() *string
	LastDetachTimestamp() *string
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
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	ReplicaZones() *[]*string
	SetReplicaZones(val *[]*string)
	ReplicaZonesInput() *[]*string
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
	SourceSnapshotEncryptionKey() ComputeRegionDiskSourceSnapshotEncryptionKeyOutputReference
	SourceSnapshotEncryptionKeyInput() *ComputeRegionDiskSourceSnapshotEncryptionKey
	SourceSnapshotId() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ComputeRegionDiskTimeoutsOutputReference
	TimeoutsInput() interface{}
	Type() *string
	SetType(val *string)
	TypeInput() *string
	Users() *[]*string
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
	PutDiskEncryptionKey(value *ComputeRegionDiskDiskEncryptionKey)
	PutSourceSnapshotEncryptionKey(value *ComputeRegionDiskSourceSnapshotEncryptionKey)
	PutTimeouts(value *ComputeRegionDiskTimeouts)
	ResetDescription()
	ResetDiskEncryptionKey()
	ResetId()
	ResetLabels()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPhysicalBlockSizeBytes()
	ResetProject()
	ResetRegion()
	ResetSize()
	ResetSnapshot()
	ResetSourceDisk()
	ResetSourceSnapshotEncryptionKey()
	ResetTimeouts()
	ResetType()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ComputeRegionDisk
type jsiiProxy_ComputeRegionDisk struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ComputeRegionDisk) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) CreationTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) DiskEncryptionKey() ComputeRegionDiskDiskEncryptionKeyOutputReference {
	var returns ComputeRegionDiskDiskEncryptionKeyOutputReference
	_jsii_.Get(
		j,
		"diskEncryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) DiskEncryptionKeyInput() *ComputeRegionDiskDiskEncryptionKey {
	var returns *ComputeRegionDiskDiskEncryptionKey
	_jsii_.Get(
		j,
		"diskEncryptionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) LabelFingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelFingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) LastAttachTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastAttachTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) LastDetachTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastDetachTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) PhysicalBlockSizeBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"physicalBlockSizeBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) PhysicalBlockSizeBytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"physicalBlockSizeBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) ReplicaZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replicaZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) ReplicaZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replicaZonesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) Size() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"size",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) SizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) Snapshot() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) SnapshotInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) SourceDisk() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) SourceDiskId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDiskId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) SourceDiskInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) SourceSnapshotEncryptionKey() ComputeRegionDiskSourceSnapshotEncryptionKeyOutputReference {
	var returns ComputeRegionDiskSourceSnapshotEncryptionKeyOutputReference
	_jsii_.Get(
		j,
		"sourceSnapshotEncryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) SourceSnapshotEncryptionKeyInput() *ComputeRegionDiskSourceSnapshotEncryptionKey {
	var returns *ComputeRegionDiskSourceSnapshotEncryptionKey
	_jsii_.Get(
		j,
		"sourceSnapshotEncryptionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) SourceSnapshotId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceSnapshotId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) Timeouts() ComputeRegionDiskTimeoutsOutputReference {
	var returns ComputeRegionDiskTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeRegionDisk) Users() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"users",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/google/r/compute_region_disk google_compute_region_disk} Resource.
func NewComputeRegionDisk(scope constructs.Construct, id *string, config *ComputeRegionDiskConfig) ComputeRegionDisk {
	_init_.Initialize()

	if err := validateNewComputeRegionDiskParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeRegionDisk{}

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionDisk.ComputeRegionDisk",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/google/r/compute_region_disk google_compute_region_disk} Resource.
func NewComputeRegionDisk_Override(c ComputeRegionDisk, scope constructs.Construct, id *string, config *ComputeRegionDiskConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeRegionDisk.ComputeRegionDisk",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ComputeRegionDisk)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionDisk)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionDisk)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionDisk)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionDisk)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionDisk)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionDisk)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionDisk)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionDisk)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionDisk)SetPhysicalBlockSizeBytes(val *float64) {
	if err := j.validateSetPhysicalBlockSizeBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"physicalBlockSizeBytes",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionDisk)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionDisk)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionDisk)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionDisk)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionDisk)SetReplicaZones(val *[]*string) {
	if err := j.validateSetReplicaZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicaZones",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionDisk)SetSize(val *float64) {
	if err := j.validateSetSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"size",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionDisk)SetSnapshot(val *string) {
	if err := j.validateSetSnapshotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshot",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionDisk)SetSourceDisk(val *string) {
	if err := j.validateSetSourceDiskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDisk",
		val,
	)
}

func (j *jsiiProxy_ComputeRegionDisk)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
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
func ComputeRegionDisk_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeRegionDisk_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeRegionDisk.ComputeRegionDisk",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ComputeRegionDisk_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.computeRegionDisk.ComputeRegionDisk",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ComputeRegionDisk) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ComputeRegionDisk) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ComputeRegionDisk) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeRegionDisk) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ComputeRegionDisk) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ComputeRegionDisk) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ComputeRegionDisk) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ComputeRegionDisk) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ComputeRegionDisk) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ComputeRegionDisk) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ComputeRegionDisk) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ComputeRegionDisk) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ComputeRegionDisk) PutDiskEncryptionKey(value *ComputeRegionDiskDiskEncryptionKey) {
	if err := c.validatePutDiskEncryptionKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDiskEncryptionKey",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionDisk) PutSourceSnapshotEncryptionKey(value *ComputeRegionDiskSourceSnapshotEncryptionKey) {
	if err := c.validatePutSourceSnapshotEncryptionKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSourceSnapshotEncryptionKey",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionDisk) PutTimeouts(value *ComputeRegionDiskTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeRegionDisk) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionDisk) ResetDiskEncryptionKey() {
	_jsii_.InvokeVoid(
		c,
		"resetDiskEncryptionKey",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionDisk) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionDisk) ResetLabels() {
	_jsii_.InvokeVoid(
		c,
		"resetLabels",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionDisk) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionDisk) ResetPhysicalBlockSizeBytes() {
	_jsii_.InvokeVoid(
		c,
		"resetPhysicalBlockSizeBytes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionDisk) ResetProject() {
	_jsii_.InvokeVoid(
		c,
		"resetProject",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionDisk) ResetRegion() {
	_jsii_.InvokeVoid(
		c,
		"resetRegion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionDisk) ResetSize() {
	_jsii_.InvokeVoid(
		c,
		"resetSize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionDisk) ResetSnapshot() {
	_jsii_.InvokeVoid(
		c,
		"resetSnapshot",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionDisk) ResetSourceDisk() {
	_jsii_.InvokeVoid(
		c,
		"resetSourceDisk",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionDisk) ResetSourceSnapshotEncryptionKey() {
	_jsii_.InvokeVoid(
		c,
		"resetSourceSnapshotEncryptionKey",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionDisk) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionDisk) ResetType() {
	_jsii_.InvokeVoid(
		c,
		"resetType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeRegionDisk) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionDisk) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionDisk) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeRegionDisk) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

