// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package artifactregistryrepository

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v14/artifactregistryrepository/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/artifact_registry_repository google_artifact_registry_repository}.
type ArtifactRegistryRepository interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CleanupPolicies() ArtifactRegistryRepositoryCleanupPoliciesList
	CleanupPoliciesInput() interface{}
	CleanupPolicyDryRun() interface{}
	SetCleanupPolicyDryRun(val interface{})
	CleanupPolicyDryRunInput() interface{}
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DockerConfig() ArtifactRegistryRepositoryDockerConfigOutputReference
	DockerConfigInput() *ArtifactRegistryRepositoryDockerConfig
	EffectiveLabels() cdktf.StringMap
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	Format() *string
	SetFormat(val *string)
	FormatInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	KmsKeyName() *string
	SetKmsKeyName(val *string)
	KmsKeyNameInput() *string
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
	MavenConfig() ArtifactRegistryRepositoryMavenConfigOutputReference
	MavenConfigInput() *ArtifactRegistryRepositoryMavenConfig
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	Name() *string
	// The tree node.
	Node() constructs.Node
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
	RemoteRepositoryConfig() ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference
	RemoteRepositoryConfigInput() *ArtifactRegistryRepositoryRemoteRepositoryConfig
	RepositoryId() *string
	SetRepositoryId(val *string)
	RepositoryIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ArtifactRegistryRepositoryTimeoutsOutputReference
	TimeoutsInput() interface{}
	UpdateTime() *string
	VirtualRepositoryConfig() ArtifactRegistryRepositoryVirtualRepositoryConfigOutputReference
	VirtualRepositoryConfigInput() *ArtifactRegistryRepositoryVirtualRepositoryConfig
	VulnerabilityScanningConfig() ArtifactRegistryRepositoryVulnerabilityScanningConfigOutputReference
	VulnerabilityScanningConfigInput() *ArtifactRegistryRepositoryVulnerabilityScanningConfig
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
	PutCleanupPolicies(value interface{})
	PutDockerConfig(value *ArtifactRegistryRepositoryDockerConfig)
	PutMavenConfig(value *ArtifactRegistryRepositoryMavenConfig)
	PutRemoteRepositoryConfig(value *ArtifactRegistryRepositoryRemoteRepositoryConfig)
	PutTimeouts(value *ArtifactRegistryRepositoryTimeouts)
	PutVirtualRepositoryConfig(value *ArtifactRegistryRepositoryVirtualRepositoryConfig)
	PutVulnerabilityScanningConfig(value *ArtifactRegistryRepositoryVulnerabilityScanningConfig)
	ResetCleanupPolicies()
	ResetCleanupPolicyDryRun()
	ResetDescription()
	ResetDockerConfig()
	ResetId()
	ResetKmsKeyName()
	ResetLabels()
	ResetLocation()
	ResetMavenConfig()
	ResetMode()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetRemoteRepositoryConfig()
	ResetTimeouts()
	ResetVirtualRepositoryConfig()
	ResetVulnerabilityScanningConfig()
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

// The jsii proxy struct for ArtifactRegistryRepository
type jsiiProxy_ArtifactRegistryRepository struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ArtifactRegistryRepository) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) CleanupPolicies() ArtifactRegistryRepositoryCleanupPoliciesList {
	var returns ArtifactRegistryRepositoryCleanupPoliciesList
	_jsii_.Get(
		j,
		"cleanupPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) CleanupPoliciesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cleanupPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) CleanupPolicyDryRun() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cleanupPolicyDryRun",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) CleanupPolicyDryRunInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cleanupPolicyDryRunInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) DockerConfig() ArtifactRegistryRepositoryDockerConfigOutputReference {
	var returns ArtifactRegistryRepositoryDockerConfigOutputReference
	_jsii_.Get(
		j,
		"dockerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) DockerConfigInput() *ArtifactRegistryRepositoryDockerConfig {
	var returns *ArtifactRegistryRepositoryDockerConfig
	_jsii_.Get(
		j,
		"dockerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) Format() *string {
	var returns *string
	_jsii_.Get(
		j,
		"format",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) FormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) KmsKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) KmsKeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) MavenConfig() ArtifactRegistryRepositoryMavenConfigOutputReference {
	var returns ArtifactRegistryRepositoryMavenConfigOutputReference
	_jsii_.Get(
		j,
		"mavenConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) MavenConfigInput() *ArtifactRegistryRepositoryMavenConfig {
	var returns *ArtifactRegistryRepositoryMavenConfig
	_jsii_.Get(
		j,
		"mavenConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) RemoteRepositoryConfig() ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference {
	var returns ArtifactRegistryRepositoryRemoteRepositoryConfigOutputReference
	_jsii_.Get(
		j,
		"remoteRepositoryConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) RemoteRepositoryConfigInput() *ArtifactRegistryRepositoryRemoteRepositoryConfig {
	var returns *ArtifactRegistryRepositoryRemoteRepositoryConfig
	_jsii_.Get(
		j,
		"remoteRepositoryConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) RepositoryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) RepositoryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) Timeouts() ArtifactRegistryRepositoryTimeoutsOutputReference {
	var returns ArtifactRegistryRepositoryTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) VirtualRepositoryConfig() ArtifactRegistryRepositoryVirtualRepositoryConfigOutputReference {
	var returns ArtifactRegistryRepositoryVirtualRepositoryConfigOutputReference
	_jsii_.Get(
		j,
		"virtualRepositoryConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) VirtualRepositoryConfigInput() *ArtifactRegistryRepositoryVirtualRepositoryConfig {
	var returns *ArtifactRegistryRepositoryVirtualRepositoryConfig
	_jsii_.Get(
		j,
		"virtualRepositoryConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) VulnerabilityScanningConfig() ArtifactRegistryRepositoryVulnerabilityScanningConfigOutputReference {
	var returns ArtifactRegistryRepositoryVulnerabilityScanningConfigOutputReference
	_jsii_.Get(
		j,
		"vulnerabilityScanningConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArtifactRegistryRepository) VulnerabilityScanningConfigInput() *ArtifactRegistryRepositoryVulnerabilityScanningConfig {
	var returns *ArtifactRegistryRepositoryVulnerabilityScanningConfig
	_jsii_.Get(
		j,
		"vulnerabilityScanningConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/artifact_registry_repository google_artifact_registry_repository} Resource.
func NewArtifactRegistryRepository(scope constructs.Construct, id *string, config *ArtifactRegistryRepositoryConfig) ArtifactRegistryRepository {
	_init_.Initialize()

	if err := validateNewArtifactRegistryRepositoryParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ArtifactRegistryRepository{}

	_jsii_.Create(
		"@cdktf/provider-google.artifactRegistryRepository.ArtifactRegistryRepository",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/artifact_registry_repository google_artifact_registry_repository} Resource.
func NewArtifactRegistryRepository_Override(a ArtifactRegistryRepository, scope constructs.Construct, id *string, config *ArtifactRegistryRepositoryConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.artifactRegistryRepository.ArtifactRegistryRepository",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ArtifactRegistryRepository)SetCleanupPolicyDryRun(val interface{}) {
	if err := j.validateSetCleanupPolicyDryRunParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cleanupPolicyDryRun",
		val,
	)
}

func (j *jsiiProxy_ArtifactRegistryRepository)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ArtifactRegistryRepository)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ArtifactRegistryRepository)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ArtifactRegistryRepository)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ArtifactRegistryRepository)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ArtifactRegistryRepository)SetFormat(val *string) {
	if err := j.validateSetFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"format",
		val,
	)
}

func (j *jsiiProxy_ArtifactRegistryRepository)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ArtifactRegistryRepository)SetKmsKeyName(val *string) {
	if err := j.validateSetKmsKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyName",
		val,
	)
}

func (j *jsiiProxy_ArtifactRegistryRepository)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_ArtifactRegistryRepository)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ArtifactRegistryRepository)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_ArtifactRegistryRepository)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_ArtifactRegistryRepository)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_ArtifactRegistryRepository)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ArtifactRegistryRepository)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ArtifactRegistryRepository)SetRepositoryId(val *string) {
	if err := j.validateSetRepositoryIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repositoryId",
		val,
	)
}

// Generates CDKTF code for importing a ArtifactRegistryRepository resource upon running "cdktf plan <stack-name>".
func ArtifactRegistryRepository_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateArtifactRegistryRepository_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.artifactRegistryRepository.ArtifactRegistryRepository",
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
func ArtifactRegistryRepository_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateArtifactRegistryRepository_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.artifactRegistryRepository.ArtifactRegistryRepository",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ArtifactRegistryRepository_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateArtifactRegistryRepository_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.artifactRegistryRepository.ArtifactRegistryRepository",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ArtifactRegistryRepository_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateArtifactRegistryRepository_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.artifactRegistryRepository.ArtifactRegistryRepository",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ArtifactRegistryRepository_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.artifactRegistryRepository.ArtifactRegistryRepository",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ArtifactRegistryRepository) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_ArtifactRegistryRepository) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ArtifactRegistryRepository) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArtifactRegistryRepository) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArtifactRegistryRepository) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArtifactRegistryRepository) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArtifactRegistryRepository) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArtifactRegistryRepository) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArtifactRegistryRepository) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArtifactRegistryRepository) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArtifactRegistryRepository) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArtifactRegistryRepository) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArtifactRegistryRepository) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_ArtifactRegistryRepository) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArtifactRegistryRepository) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ArtifactRegistryRepository) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_ArtifactRegistryRepository) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_ArtifactRegistryRepository) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ArtifactRegistryRepository) PutCleanupPolicies(value interface{}) {
	if err := a.validatePutCleanupPoliciesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCleanupPolicies",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ArtifactRegistryRepository) PutDockerConfig(value *ArtifactRegistryRepositoryDockerConfig) {
	if err := a.validatePutDockerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDockerConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ArtifactRegistryRepository) PutMavenConfig(value *ArtifactRegistryRepositoryMavenConfig) {
	if err := a.validatePutMavenConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMavenConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ArtifactRegistryRepository) PutRemoteRepositoryConfig(value *ArtifactRegistryRepositoryRemoteRepositoryConfig) {
	if err := a.validatePutRemoteRepositoryConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRemoteRepositoryConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ArtifactRegistryRepository) PutTimeouts(value *ArtifactRegistryRepositoryTimeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ArtifactRegistryRepository) PutVirtualRepositoryConfig(value *ArtifactRegistryRepositoryVirtualRepositoryConfig) {
	if err := a.validatePutVirtualRepositoryConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVirtualRepositoryConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ArtifactRegistryRepository) PutVulnerabilityScanningConfig(value *ArtifactRegistryRepositoryVulnerabilityScanningConfig) {
	if err := a.validatePutVulnerabilityScanningConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVulnerabilityScanningConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ArtifactRegistryRepository) ResetCleanupPolicies() {
	_jsii_.InvokeVoid(
		a,
		"resetCleanupPolicies",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArtifactRegistryRepository) ResetCleanupPolicyDryRun() {
	_jsii_.InvokeVoid(
		a,
		"resetCleanupPolicyDryRun",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArtifactRegistryRepository) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArtifactRegistryRepository) ResetDockerConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetDockerConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArtifactRegistryRepository) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArtifactRegistryRepository) ResetKmsKeyName() {
	_jsii_.InvokeVoid(
		a,
		"resetKmsKeyName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArtifactRegistryRepository) ResetLabels() {
	_jsii_.InvokeVoid(
		a,
		"resetLabels",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArtifactRegistryRepository) ResetLocation() {
	_jsii_.InvokeVoid(
		a,
		"resetLocation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArtifactRegistryRepository) ResetMavenConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetMavenConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArtifactRegistryRepository) ResetMode() {
	_jsii_.InvokeVoid(
		a,
		"resetMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArtifactRegistryRepository) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArtifactRegistryRepository) ResetProject() {
	_jsii_.InvokeVoid(
		a,
		"resetProject",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArtifactRegistryRepository) ResetRemoteRepositoryConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetRemoteRepositoryConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArtifactRegistryRepository) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArtifactRegistryRepository) ResetVirtualRepositoryConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetVirtualRepositoryConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArtifactRegistryRepository) ResetVulnerabilityScanningConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetVulnerabilityScanningConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArtifactRegistryRepository) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArtifactRegistryRepository) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArtifactRegistryRepository) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArtifactRegistryRepository) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArtifactRegistryRepository) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArtifactRegistryRepository) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

