// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageinsightsdatasetconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v16/storageinsightsdatasetconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/storage_insights_dataset_config google_storage_insights_dataset_config}.
type StorageInsightsDatasetConfig interface {
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
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreateTime() *string
	DatasetConfigId() *string
	SetDatasetConfigId(val *string)
	DatasetConfigIdInput() *string
	DatasetConfigState() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	ExcludeCloudStorageBuckets() StorageInsightsDatasetConfigExcludeCloudStorageBucketsOutputReference
	ExcludeCloudStorageBucketsInput() *StorageInsightsDatasetConfigExcludeCloudStorageBuckets
	ExcludeCloudStorageLocations() StorageInsightsDatasetConfigExcludeCloudStorageLocationsOutputReference
	ExcludeCloudStorageLocationsInput() *StorageInsightsDatasetConfigExcludeCloudStorageLocations
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
	Identity() StorageInsightsDatasetConfigIdentityOutputReference
	IdentityInput() *StorageInsightsDatasetConfigIdentity
	IdInput() *string
	IncludeCloudStorageBuckets() StorageInsightsDatasetConfigIncludeCloudStorageBucketsOutputReference
	IncludeCloudStorageBucketsInput() *StorageInsightsDatasetConfigIncludeCloudStorageBuckets
	IncludeCloudStorageLocations() StorageInsightsDatasetConfigIncludeCloudStorageLocationsOutputReference
	IncludeCloudStorageLocationsInput() *StorageInsightsDatasetConfigIncludeCloudStorageLocations
	IncludeNewlyCreatedBuckets() interface{}
	SetIncludeNewlyCreatedBuckets(val interface{})
	IncludeNewlyCreatedBucketsInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Link() StorageInsightsDatasetConfigLinkList
	LinkDataset() interface{}
	SetLinkDataset(val interface{})
	LinkDatasetInput() interface{}
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	// The tree node.
	Node() constructs.Node
	OrganizationNumber() *string
	SetOrganizationNumber(val *string)
	OrganizationNumberInput() *string
	OrganizationScope() interface{}
	SetOrganizationScope(val interface{})
	OrganizationScopeInput() interface{}
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
	RetentionPeriodDays() *float64
	SetRetentionPeriodDays(val *float64)
	RetentionPeriodDaysInput() *float64
	SourceFolders() StorageInsightsDatasetConfigSourceFoldersOutputReference
	SourceFoldersInput() *StorageInsightsDatasetConfigSourceFolders
	SourceProjects() StorageInsightsDatasetConfigSourceProjectsOutputReference
	SourceProjectsInput() *StorageInsightsDatasetConfigSourceProjects
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() StorageInsightsDatasetConfigTimeoutsOutputReference
	TimeoutsInput() interface{}
	Uid() *string
	UpdateTime() *string
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
	PutExcludeCloudStorageBuckets(value *StorageInsightsDatasetConfigExcludeCloudStorageBuckets)
	PutExcludeCloudStorageLocations(value *StorageInsightsDatasetConfigExcludeCloudStorageLocations)
	PutIdentity(value *StorageInsightsDatasetConfigIdentity)
	PutIncludeCloudStorageBuckets(value *StorageInsightsDatasetConfigIncludeCloudStorageBuckets)
	PutIncludeCloudStorageLocations(value *StorageInsightsDatasetConfigIncludeCloudStorageLocations)
	PutSourceFolders(value *StorageInsightsDatasetConfigSourceFolders)
	PutSourceProjects(value *StorageInsightsDatasetConfigSourceProjects)
	PutTimeouts(value *StorageInsightsDatasetConfigTimeouts)
	ResetDescription()
	ResetExcludeCloudStorageBuckets()
	ResetExcludeCloudStorageLocations()
	ResetId()
	ResetIncludeCloudStorageBuckets()
	ResetIncludeCloudStorageLocations()
	ResetIncludeNewlyCreatedBuckets()
	ResetLinkDataset()
	ResetOrganizationNumber()
	ResetOrganizationScope()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetSourceFolders()
	ResetSourceProjects()
	ResetTimeouts()
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

// The jsii proxy struct for StorageInsightsDatasetConfig
type jsiiProxy_StorageInsightsDatasetConfig struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) DatasetConfigId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetConfigId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) DatasetConfigIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetConfigIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) DatasetConfigState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datasetConfigState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) ExcludeCloudStorageBuckets() StorageInsightsDatasetConfigExcludeCloudStorageBucketsOutputReference {
	var returns StorageInsightsDatasetConfigExcludeCloudStorageBucketsOutputReference
	_jsii_.Get(
		j,
		"excludeCloudStorageBuckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) ExcludeCloudStorageBucketsInput() *StorageInsightsDatasetConfigExcludeCloudStorageBuckets {
	var returns *StorageInsightsDatasetConfigExcludeCloudStorageBuckets
	_jsii_.Get(
		j,
		"excludeCloudStorageBucketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) ExcludeCloudStorageLocations() StorageInsightsDatasetConfigExcludeCloudStorageLocationsOutputReference {
	var returns StorageInsightsDatasetConfigExcludeCloudStorageLocationsOutputReference
	_jsii_.Get(
		j,
		"excludeCloudStorageLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) ExcludeCloudStorageLocationsInput() *StorageInsightsDatasetConfigExcludeCloudStorageLocations {
	var returns *StorageInsightsDatasetConfigExcludeCloudStorageLocations
	_jsii_.Get(
		j,
		"excludeCloudStorageLocationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) Identity() StorageInsightsDatasetConfigIdentityOutputReference {
	var returns StorageInsightsDatasetConfigIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) IdentityInput() *StorageInsightsDatasetConfigIdentity {
	var returns *StorageInsightsDatasetConfigIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) IncludeCloudStorageBuckets() StorageInsightsDatasetConfigIncludeCloudStorageBucketsOutputReference {
	var returns StorageInsightsDatasetConfigIncludeCloudStorageBucketsOutputReference
	_jsii_.Get(
		j,
		"includeCloudStorageBuckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) IncludeCloudStorageBucketsInput() *StorageInsightsDatasetConfigIncludeCloudStorageBuckets {
	var returns *StorageInsightsDatasetConfigIncludeCloudStorageBuckets
	_jsii_.Get(
		j,
		"includeCloudStorageBucketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) IncludeCloudStorageLocations() StorageInsightsDatasetConfigIncludeCloudStorageLocationsOutputReference {
	var returns StorageInsightsDatasetConfigIncludeCloudStorageLocationsOutputReference
	_jsii_.Get(
		j,
		"includeCloudStorageLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) IncludeCloudStorageLocationsInput() *StorageInsightsDatasetConfigIncludeCloudStorageLocations {
	var returns *StorageInsightsDatasetConfigIncludeCloudStorageLocations
	_jsii_.Get(
		j,
		"includeCloudStorageLocationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) IncludeNewlyCreatedBuckets() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeNewlyCreatedBuckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) IncludeNewlyCreatedBucketsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeNewlyCreatedBucketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) Link() StorageInsightsDatasetConfigLinkList {
	var returns StorageInsightsDatasetConfigLinkList
	_jsii_.Get(
		j,
		"link",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) LinkDataset() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"linkDataset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) LinkDatasetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"linkDatasetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) OrganizationNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) OrganizationNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) OrganizationScope() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"organizationScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) OrganizationScopeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"organizationScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) RetentionPeriodDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionPeriodDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) RetentionPeriodDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionPeriodDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) SourceFolders() StorageInsightsDatasetConfigSourceFoldersOutputReference {
	var returns StorageInsightsDatasetConfigSourceFoldersOutputReference
	_jsii_.Get(
		j,
		"sourceFolders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) SourceFoldersInput() *StorageInsightsDatasetConfigSourceFolders {
	var returns *StorageInsightsDatasetConfigSourceFolders
	_jsii_.Get(
		j,
		"sourceFoldersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) SourceProjects() StorageInsightsDatasetConfigSourceProjectsOutputReference {
	var returns StorageInsightsDatasetConfigSourceProjectsOutputReference
	_jsii_.Get(
		j,
		"sourceProjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) SourceProjectsInput() *StorageInsightsDatasetConfigSourceProjects {
	var returns *StorageInsightsDatasetConfigSourceProjects
	_jsii_.Get(
		j,
		"sourceProjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) Timeouts() StorageInsightsDatasetConfigTimeoutsOutputReference {
	var returns StorageInsightsDatasetConfigTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) Uid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageInsightsDatasetConfig) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/storage_insights_dataset_config google_storage_insights_dataset_config} Resource.
func NewStorageInsightsDatasetConfig(scope constructs.Construct, id *string, config *StorageInsightsDatasetConfigConfig) StorageInsightsDatasetConfig {
	_init_.Initialize()

	if err := validateNewStorageInsightsDatasetConfigParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_StorageInsightsDatasetConfig{}

	_jsii_.Create(
		"@cdktf/provider-google.storageInsightsDatasetConfig.StorageInsightsDatasetConfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/storage_insights_dataset_config google_storage_insights_dataset_config} Resource.
func NewStorageInsightsDatasetConfig_Override(s StorageInsightsDatasetConfig, scope constructs.Construct, id *string, config *StorageInsightsDatasetConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.storageInsightsDatasetConfig.StorageInsightsDatasetConfig",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_StorageInsightsDatasetConfig)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_StorageInsightsDatasetConfig)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_StorageInsightsDatasetConfig)SetDatasetConfigId(val *string) {
	if err := j.validateSetDatasetConfigIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datasetConfigId",
		val,
	)
}

func (j *jsiiProxy_StorageInsightsDatasetConfig)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_StorageInsightsDatasetConfig)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_StorageInsightsDatasetConfig)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_StorageInsightsDatasetConfig)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_StorageInsightsDatasetConfig)SetIncludeNewlyCreatedBuckets(val interface{}) {
	if err := j.validateSetIncludeNewlyCreatedBucketsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeNewlyCreatedBuckets",
		val,
	)
}

func (j *jsiiProxy_StorageInsightsDatasetConfig)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_StorageInsightsDatasetConfig)SetLinkDataset(val interface{}) {
	if err := j.validateSetLinkDatasetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"linkDataset",
		val,
	)
}

func (j *jsiiProxy_StorageInsightsDatasetConfig)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_StorageInsightsDatasetConfig)SetOrganizationNumber(val *string) {
	if err := j.validateSetOrganizationNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organizationNumber",
		val,
	)
}

func (j *jsiiProxy_StorageInsightsDatasetConfig)SetOrganizationScope(val interface{}) {
	if err := j.validateSetOrganizationScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organizationScope",
		val,
	)
}

func (j *jsiiProxy_StorageInsightsDatasetConfig)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_StorageInsightsDatasetConfig)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_StorageInsightsDatasetConfig)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_StorageInsightsDatasetConfig)SetRetentionPeriodDays(val *float64) {
	if err := j.validateSetRetentionPeriodDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionPeriodDays",
		val,
	)
}

// Generates CDKTF code for importing a StorageInsightsDatasetConfig resource upon running "cdktf plan <stack-name>".
func StorageInsightsDatasetConfig_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateStorageInsightsDatasetConfig_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.storageInsightsDatasetConfig.StorageInsightsDatasetConfig",
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
func StorageInsightsDatasetConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStorageInsightsDatasetConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.storageInsightsDatasetConfig.StorageInsightsDatasetConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func StorageInsightsDatasetConfig_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStorageInsightsDatasetConfig_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.storageInsightsDatasetConfig.StorageInsightsDatasetConfig",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func StorageInsightsDatasetConfig_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStorageInsightsDatasetConfig_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.storageInsightsDatasetConfig.StorageInsightsDatasetConfig",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func StorageInsightsDatasetConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.storageInsightsDatasetConfig.StorageInsightsDatasetConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) PutExcludeCloudStorageBuckets(value *StorageInsightsDatasetConfigExcludeCloudStorageBuckets) {
	if err := s.validatePutExcludeCloudStorageBucketsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putExcludeCloudStorageBuckets",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) PutExcludeCloudStorageLocations(value *StorageInsightsDatasetConfigExcludeCloudStorageLocations) {
	if err := s.validatePutExcludeCloudStorageLocationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putExcludeCloudStorageLocations",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) PutIdentity(value *StorageInsightsDatasetConfigIdentity) {
	if err := s.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putIdentity",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) PutIncludeCloudStorageBuckets(value *StorageInsightsDatasetConfigIncludeCloudStorageBuckets) {
	if err := s.validatePutIncludeCloudStorageBucketsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putIncludeCloudStorageBuckets",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) PutIncludeCloudStorageLocations(value *StorageInsightsDatasetConfigIncludeCloudStorageLocations) {
	if err := s.validatePutIncludeCloudStorageLocationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putIncludeCloudStorageLocations",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) PutSourceFolders(value *StorageInsightsDatasetConfigSourceFolders) {
	if err := s.validatePutSourceFoldersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSourceFolders",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) PutSourceProjects(value *StorageInsightsDatasetConfigSourceProjects) {
	if err := s.validatePutSourceProjectsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSourceProjects",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) PutTimeouts(value *StorageInsightsDatasetConfigTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) ResetDescription() {
	_jsii_.InvokeVoid(
		s,
		"resetDescription",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) ResetExcludeCloudStorageBuckets() {
	_jsii_.InvokeVoid(
		s,
		"resetExcludeCloudStorageBuckets",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) ResetExcludeCloudStorageLocations() {
	_jsii_.InvokeVoid(
		s,
		"resetExcludeCloudStorageLocations",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) ResetIncludeCloudStorageBuckets() {
	_jsii_.InvokeVoid(
		s,
		"resetIncludeCloudStorageBuckets",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) ResetIncludeCloudStorageLocations() {
	_jsii_.InvokeVoid(
		s,
		"resetIncludeCloudStorageLocations",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) ResetIncludeNewlyCreatedBuckets() {
	_jsii_.InvokeVoid(
		s,
		"resetIncludeNewlyCreatedBuckets",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) ResetLinkDataset() {
	_jsii_.InvokeVoid(
		s,
		"resetLinkDataset",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) ResetOrganizationNumber() {
	_jsii_.InvokeVoid(
		s,
		"resetOrganizationNumber",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) ResetOrganizationScope() {
	_jsii_.InvokeVoid(
		s,
		"resetOrganizationScope",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) ResetProject() {
	_jsii_.InvokeVoid(
		s,
		"resetProject",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) ResetSourceFolders() {
	_jsii_.InvokeVoid(
		s,
		"resetSourceFolders",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) ResetSourceProjects() {
	_jsii_.InvokeVoid(
		s,
		"resetSourceProjects",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageInsightsDatasetConfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

