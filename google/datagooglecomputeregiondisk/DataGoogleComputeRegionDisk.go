// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagooglecomputeregiondisk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v14/datagooglecomputeregiondisk/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/data-sources/compute_region_disk google_compute_region_disk}.
type DataGoogleComputeRegionDisk interface {
	cdktf.TerraformDataSource
	AsyncPrimaryDisk() DataGoogleComputeRegionDiskAsyncPrimaryDiskList
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreateSnapshotBeforeDestroy() cdktf.IResolvable
	CreateSnapshotBeforeDestroyPrefix() *string
	CreationTimestamp() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	DiskEncryptionKey() DataGoogleComputeRegionDiskDiskEncryptionKeyList
	DiskId() *string
	EffectiveLabels() cdktf.StringMap
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GuestOsFeatures() DataGoogleComputeRegionDiskGuestOsFeaturesList
	Id() *string
	SetId(val *string)
	IdInput() *string
	LabelFingerprint() *string
	Labels() cdktf.StringMap
	LastAttachTimestamp() *string
	LastDetachTimestamp() *string
	Licenses() *[]*string
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
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	ReplicaZones() *[]*string
	SelfLink() *string
	Size() *float64
	Snapshot() *string
	SourceDisk() *string
	SourceDiskId() *string
	SourceSnapshotEncryptionKey() DataGoogleComputeRegionDiskSourceSnapshotEncryptionKeyList
	SourceSnapshotId() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Type() *string
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
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetRegion()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
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

// The jsii proxy struct for DataGoogleComputeRegionDisk
type jsiiProxy_DataGoogleComputeRegionDisk struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) AsyncPrimaryDisk() DataGoogleComputeRegionDiskAsyncPrimaryDiskList {
	var returns DataGoogleComputeRegionDiskAsyncPrimaryDiskList
	_jsii_.Get(
		j,
		"asyncPrimaryDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) CreateSnapshotBeforeDestroy() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"createSnapshotBeforeDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) CreateSnapshotBeforeDestroyPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createSnapshotBeforeDestroyPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) CreationTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) DiskEncryptionKey() DataGoogleComputeRegionDiskDiskEncryptionKeyList {
	var returns DataGoogleComputeRegionDiskDiskEncryptionKeyList
	_jsii_.Get(
		j,
		"diskEncryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) DiskId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) GuestOsFeatures() DataGoogleComputeRegionDiskGuestOsFeaturesList {
	var returns DataGoogleComputeRegionDiskGuestOsFeaturesList
	_jsii_.Get(
		j,
		"guestOsFeatures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) LabelFingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelFingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) Labels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) LastAttachTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastAttachTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) LastDetachTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastDetachTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) Licenses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"licenses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) PhysicalBlockSizeBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"physicalBlockSizeBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) ReplicaZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replicaZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) Size() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"size",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) Snapshot() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) SourceDisk() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) SourceDiskId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDiskId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) SourceSnapshotEncryptionKey() DataGoogleComputeRegionDiskSourceSnapshotEncryptionKeyList {
	var returns DataGoogleComputeRegionDiskSourceSnapshotEncryptionKeyList
	_jsii_.Get(
		j,
		"sourceSnapshotEncryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) SourceSnapshotId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceSnapshotId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk) Users() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"users",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/data-sources/compute_region_disk google_compute_region_disk} Data Source.
func NewDataGoogleComputeRegionDisk(scope constructs.Construct, id *string, config *DataGoogleComputeRegionDiskConfig) DataGoogleComputeRegionDisk {
	_init_.Initialize()

	if err := validateNewDataGoogleComputeRegionDiskParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGoogleComputeRegionDisk{}

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleComputeRegionDisk.DataGoogleComputeRegionDisk",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/data-sources/compute_region_disk google_compute_region_disk} Data Source.
func NewDataGoogleComputeRegionDisk_Override(d DataGoogleComputeRegionDisk, scope constructs.Construct, id *string, config *DataGoogleComputeRegionDiskConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleComputeRegionDisk.DataGoogleComputeRegionDisk",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataGoogleComputeRegionDisk)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

// Generates CDKTF code for importing a DataGoogleComputeRegionDisk resource upon running "cdktf plan <stack-name>".
func DataGoogleComputeRegionDisk_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataGoogleComputeRegionDisk_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dataGoogleComputeRegionDisk.DataGoogleComputeRegionDisk",
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
func DataGoogleComputeRegionDisk_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGoogleComputeRegionDisk_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dataGoogleComputeRegionDisk.DataGoogleComputeRegionDisk",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataGoogleComputeRegionDisk_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGoogleComputeRegionDisk_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dataGoogleComputeRegionDisk.DataGoogleComputeRegionDisk",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataGoogleComputeRegionDisk_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGoogleComputeRegionDisk_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dataGoogleComputeRegionDisk.DataGoogleComputeRegionDisk",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataGoogleComputeRegionDisk_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.dataGoogleComputeRegionDisk.DataGoogleComputeRegionDisk",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataGoogleComputeRegionDisk) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataGoogleComputeRegionDisk) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataGoogleComputeRegionDisk) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGoogleComputeRegionDisk) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataGoogleComputeRegionDisk) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataGoogleComputeRegionDisk) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataGoogleComputeRegionDisk) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataGoogleComputeRegionDisk) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataGoogleComputeRegionDisk) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataGoogleComputeRegionDisk) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataGoogleComputeRegionDisk) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeRegionDisk) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataGoogleComputeRegionDisk) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGoogleComputeRegionDisk) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGoogleComputeRegionDisk) ResetProject() {
	_jsii_.InvokeVoid(
		d,
		"resetProject",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGoogleComputeRegionDisk) ResetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGoogleComputeRegionDisk) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeRegionDisk) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeRegionDisk) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeRegionDisk) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeRegionDisk) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleComputeRegionDisk) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

