// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudbuildtrigger

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v14/cloudbuildtrigger/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/cloudbuild_trigger google_cloudbuild_trigger}.
type CloudbuildTrigger interface {
	cdktf.TerraformResource
	ApprovalConfig() CloudbuildTriggerApprovalConfigOutputReference
	ApprovalConfigInput() *CloudbuildTriggerApprovalConfig
	BitbucketServerTriggerConfig() CloudbuildTriggerBitbucketServerTriggerConfigOutputReference
	BitbucketServerTriggerConfigInput() *CloudbuildTriggerBitbucketServerTriggerConfig
	BuildAttribute() CloudbuildTriggerBuildOutputReference
	BuildAttributeInput() *CloudbuildTriggerBuild
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Disabled() interface{}
	SetDisabled(val interface{})
	DisabledInput() interface{}
	Filename() *string
	SetFilename(val *string)
	FilenameInput() *string
	Filter() *string
	SetFilter(val *string)
	FilterInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GitFileSource() CloudbuildTriggerGitFileSourceOutputReference
	GitFileSourceInput() *CloudbuildTriggerGitFileSource
	Github() CloudbuildTriggerGithubOutputReference
	GithubInput() *CloudbuildTriggerGithub
	Id() *string
	SetId(val *string)
	IdInput() *string
	IgnoredFiles() *[]*string
	SetIgnoredFiles(val *[]*string)
	IgnoredFilesInput() *[]*string
	IncludeBuildLogs() *string
	SetIncludeBuildLogs(val *string)
	IncludeBuildLogsInput() *string
	IncludedFiles() *[]*string
	SetIncludedFiles(val *[]*string)
	IncludedFilesInput() *[]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	PubsubConfig() CloudbuildTriggerPubsubConfigOutputReference
	PubsubConfigInput() *CloudbuildTriggerPubsubConfig
	// Experimental.
	RawOverrides() interface{}
	RepositoryEventConfig() CloudbuildTriggerRepositoryEventConfigOutputReference
	RepositoryEventConfigInput() *CloudbuildTriggerRepositoryEventConfig
	ServiceAccount() *string
	SetServiceAccount(val *string)
	ServiceAccountInput() *string
	SourceToBuild() CloudbuildTriggerSourceToBuildOutputReference
	SourceToBuildInput() *CloudbuildTriggerSourceToBuild
	Substitutions() *map[string]*string
	SetSubstitutions(val *map[string]*string)
	SubstitutionsInput() *map[string]*string
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() CloudbuildTriggerTimeoutsOutputReference
	TimeoutsInput() interface{}
	TriggerId() *string
	TriggerTemplate() CloudbuildTriggerTriggerTemplateOutputReference
	TriggerTemplateInput() *CloudbuildTriggerTriggerTemplate
	WebhookConfig() CloudbuildTriggerWebhookConfigOutputReference
	WebhookConfigInput() *CloudbuildTriggerWebhookConfig
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
	PutApprovalConfig(value *CloudbuildTriggerApprovalConfig)
	PutBitbucketServerTriggerConfig(value *CloudbuildTriggerBitbucketServerTriggerConfig)
	PutBuildAttribute(value *CloudbuildTriggerBuild)
	PutGitFileSource(value *CloudbuildTriggerGitFileSource)
	PutGithub(value *CloudbuildTriggerGithub)
	PutPubsubConfig(value *CloudbuildTriggerPubsubConfig)
	PutRepositoryEventConfig(value *CloudbuildTriggerRepositoryEventConfig)
	PutSourceToBuild(value *CloudbuildTriggerSourceToBuild)
	PutTimeouts(value *CloudbuildTriggerTimeouts)
	PutTriggerTemplate(value *CloudbuildTriggerTriggerTemplate)
	PutWebhookConfig(value *CloudbuildTriggerWebhookConfig)
	ResetApprovalConfig()
	ResetBitbucketServerTriggerConfig()
	ResetBuildAttribute()
	ResetDescription()
	ResetDisabled()
	ResetFilename()
	ResetFilter()
	ResetGitFileSource()
	ResetGithub()
	ResetId()
	ResetIgnoredFiles()
	ResetIncludeBuildLogs()
	ResetIncludedFiles()
	ResetLocation()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetPubsubConfig()
	ResetRepositoryEventConfig()
	ResetServiceAccount()
	ResetSourceToBuild()
	ResetSubstitutions()
	ResetTags()
	ResetTimeouts()
	ResetTriggerTemplate()
	ResetWebhookConfig()
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

// The jsii proxy struct for CloudbuildTrigger
type jsiiProxy_CloudbuildTrigger struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CloudbuildTrigger) ApprovalConfig() CloudbuildTriggerApprovalConfigOutputReference {
	var returns CloudbuildTriggerApprovalConfigOutputReference
	_jsii_.Get(
		j,
		"approvalConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) ApprovalConfigInput() *CloudbuildTriggerApprovalConfig {
	var returns *CloudbuildTriggerApprovalConfig
	_jsii_.Get(
		j,
		"approvalConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) BitbucketServerTriggerConfig() CloudbuildTriggerBitbucketServerTriggerConfigOutputReference {
	var returns CloudbuildTriggerBitbucketServerTriggerConfigOutputReference
	_jsii_.Get(
		j,
		"bitbucketServerTriggerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) BitbucketServerTriggerConfigInput() *CloudbuildTriggerBitbucketServerTriggerConfig {
	var returns *CloudbuildTriggerBitbucketServerTriggerConfig
	_jsii_.Get(
		j,
		"bitbucketServerTriggerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) BuildAttribute() CloudbuildTriggerBuildOutputReference {
	var returns CloudbuildTriggerBuildOutputReference
	_jsii_.Get(
		j,
		"buildAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) BuildAttributeInput() *CloudbuildTriggerBuild {
	var returns *CloudbuildTriggerBuild
	_jsii_.Get(
		j,
		"buildAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) Disabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) DisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) Filename() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filename",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) FilenameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filenameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) Filter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) FilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) GitFileSource() CloudbuildTriggerGitFileSourceOutputReference {
	var returns CloudbuildTriggerGitFileSourceOutputReference
	_jsii_.Get(
		j,
		"gitFileSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) GitFileSourceInput() *CloudbuildTriggerGitFileSource {
	var returns *CloudbuildTriggerGitFileSource
	_jsii_.Get(
		j,
		"gitFileSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) Github() CloudbuildTriggerGithubOutputReference {
	var returns CloudbuildTriggerGithubOutputReference
	_jsii_.Get(
		j,
		"github",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) GithubInput() *CloudbuildTriggerGithub {
	var returns *CloudbuildTriggerGithub
	_jsii_.Get(
		j,
		"githubInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) IgnoredFiles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ignoredFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) IgnoredFilesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ignoredFilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) IncludeBuildLogs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"includeBuildLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) IncludeBuildLogsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"includeBuildLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) IncludedFiles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) IncludedFilesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedFilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) PubsubConfig() CloudbuildTriggerPubsubConfigOutputReference {
	var returns CloudbuildTriggerPubsubConfigOutputReference
	_jsii_.Get(
		j,
		"pubsubConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) PubsubConfigInput() *CloudbuildTriggerPubsubConfig {
	var returns *CloudbuildTriggerPubsubConfig
	_jsii_.Get(
		j,
		"pubsubConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) RepositoryEventConfig() CloudbuildTriggerRepositoryEventConfigOutputReference {
	var returns CloudbuildTriggerRepositoryEventConfigOutputReference
	_jsii_.Get(
		j,
		"repositoryEventConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) RepositoryEventConfigInput() *CloudbuildTriggerRepositoryEventConfig {
	var returns *CloudbuildTriggerRepositoryEventConfig
	_jsii_.Get(
		j,
		"repositoryEventConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) ServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) SourceToBuild() CloudbuildTriggerSourceToBuildOutputReference {
	var returns CloudbuildTriggerSourceToBuildOutputReference
	_jsii_.Get(
		j,
		"sourceToBuild",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) SourceToBuildInput() *CloudbuildTriggerSourceToBuild {
	var returns *CloudbuildTriggerSourceToBuild
	_jsii_.Get(
		j,
		"sourceToBuildInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) Substitutions() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"substitutions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) SubstitutionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"substitutionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) Timeouts() CloudbuildTriggerTimeoutsOutputReference {
	var returns CloudbuildTriggerTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) TriggerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"triggerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) TriggerTemplate() CloudbuildTriggerTriggerTemplateOutputReference {
	var returns CloudbuildTriggerTriggerTemplateOutputReference
	_jsii_.Get(
		j,
		"triggerTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) TriggerTemplateInput() *CloudbuildTriggerTriggerTemplate {
	var returns *CloudbuildTriggerTriggerTemplate
	_jsii_.Get(
		j,
		"triggerTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) WebhookConfig() CloudbuildTriggerWebhookConfigOutputReference {
	var returns CloudbuildTriggerWebhookConfigOutputReference
	_jsii_.Get(
		j,
		"webhookConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudbuildTrigger) WebhookConfigInput() *CloudbuildTriggerWebhookConfig {
	var returns *CloudbuildTriggerWebhookConfig
	_jsii_.Get(
		j,
		"webhookConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/cloudbuild_trigger google_cloudbuild_trigger} Resource.
func NewCloudbuildTrigger(scope constructs.Construct, id *string, config *CloudbuildTriggerConfig) CloudbuildTrigger {
	_init_.Initialize()

	if err := validateNewCloudbuildTriggerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudbuildTrigger{}

	_jsii_.Create(
		"@cdktf/provider-google.cloudbuildTrigger.CloudbuildTrigger",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/cloudbuild_trigger google_cloudbuild_trigger} Resource.
func NewCloudbuildTrigger_Override(c CloudbuildTrigger, scope constructs.Construct, id *string, config *CloudbuildTriggerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cloudbuildTrigger.CloudbuildTrigger",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CloudbuildTrigger)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTrigger)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTrigger)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTrigger)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTrigger)SetDisabled(val interface{}) {
	if err := j.validateSetDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disabled",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTrigger)SetFilename(val *string) {
	if err := j.validateSetFilenameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filename",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTrigger)SetFilter(val *string) {
	if err := j.validateSetFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filter",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTrigger)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTrigger)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTrigger)SetIgnoredFiles(val *[]*string) {
	if err := j.validateSetIgnoredFilesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoredFiles",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTrigger)SetIncludeBuildLogs(val *string) {
	if err := j.validateSetIncludeBuildLogsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeBuildLogs",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTrigger)SetIncludedFiles(val *[]*string) {
	if err := j.validateSetIncludedFilesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includedFiles",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTrigger)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTrigger)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTrigger)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTrigger)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTrigger)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTrigger)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTrigger)SetServiceAccount(val *string) {
	if err := j.validateSetServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccount",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTrigger)SetSubstitutions(val *map[string]*string) {
	if err := j.validateSetSubstitutionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"substitutions",
		val,
	)
}

func (j *jsiiProxy_CloudbuildTrigger)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a CloudbuildTrigger resource upon running "cdktf plan <stack-name>".
func CloudbuildTrigger_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCloudbuildTrigger_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.cloudbuildTrigger.CloudbuildTrigger",
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
func CloudbuildTrigger_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudbuildTrigger_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.cloudbuildTrigger.CloudbuildTrigger",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudbuildTrigger_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudbuildTrigger_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.cloudbuildTrigger.CloudbuildTrigger",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CloudbuildTrigger_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudbuildTrigger_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.cloudbuildTrigger.CloudbuildTrigger",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CloudbuildTrigger_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.cloudbuildTrigger.CloudbuildTrigger",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CloudbuildTrigger) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_CloudbuildTrigger) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CloudbuildTrigger) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CloudbuildTrigger) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudbuildTrigger) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudbuildTrigger) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudbuildTrigger) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudbuildTrigger) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudbuildTrigger) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudbuildTrigger) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudbuildTrigger) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudbuildTrigger) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTrigger) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_CloudbuildTrigger) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudbuildTrigger) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CloudbuildTrigger) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_CloudbuildTrigger) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CloudbuildTrigger) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CloudbuildTrigger) PutApprovalConfig(value *CloudbuildTriggerApprovalConfig) {
	if err := c.validatePutApprovalConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putApprovalConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudbuildTrigger) PutBitbucketServerTriggerConfig(value *CloudbuildTriggerBitbucketServerTriggerConfig) {
	if err := c.validatePutBitbucketServerTriggerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBitbucketServerTriggerConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudbuildTrigger) PutBuildAttribute(value *CloudbuildTriggerBuild) {
	if err := c.validatePutBuildAttributeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBuildAttribute",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudbuildTrigger) PutGitFileSource(value *CloudbuildTriggerGitFileSource) {
	if err := c.validatePutGitFileSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGitFileSource",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudbuildTrigger) PutGithub(value *CloudbuildTriggerGithub) {
	if err := c.validatePutGithubParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGithub",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudbuildTrigger) PutPubsubConfig(value *CloudbuildTriggerPubsubConfig) {
	if err := c.validatePutPubsubConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPubsubConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudbuildTrigger) PutRepositoryEventConfig(value *CloudbuildTriggerRepositoryEventConfig) {
	if err := c.validatePutRepositoryEventConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRepositoryEventConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudbuildTrigger) PutSourceToBuild(value *CloudbuildTriggerSourceToBuild) {
	if err := c.validatePutSourceToBuildParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSourceToBuild",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudbuildTrigger) PutTimeouts(value *CloudbuildTriggerTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudbuildTrigger) PutTriggerTemplate(value *CloudbuildTriggerTriggerTemplate) {
	if err := c.validatePutTriggerTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTriggerTemplate",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudbuildTrigger) PutWebhookConfig(value *CloudbuildTriggerWebhookConfig) {
	if err := c.validatePutWebhookConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putWebhookConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudbuildTrigger) ResetApprovalConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetApprovalConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTrigger) ResetBitbucketServerTriggerConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetBitbucketServerTriggerConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTrigger) ResetBuildAttribute() {
	_jsii_.InvokeVoid(
		c,
		"resetBuildAttribute",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTrigger) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTrigger) ResetDisabled() {
	_jsii_.InvokeVoid(
		c,
		"resetDisabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTrigger) ResetFilename() {
	_jsii_.InvokeVoid(
		c,
		"resetFilename",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTrigger) ResetFilter() {
	_jsii_.InvokeVoid(
		c,
		"resetFilter",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTrigger) ResetGitFileSource() {
	_jsii_.InvokeVoid(
		c,
		"resetGitFileSource",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTrigger) ResetGithub() {
	_jsii_.InvokeVoid(
		c,
		"resetGithub",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTrigger) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTrigger) ResetIgnoredFiles() {
	_jsii_.InvokeVoid(
		c,
		"resetIgnoredFiles",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTrigger) ResetIncludeBuildLogs() {
	_jsii_.InvokeVoid(
		c,
		"resetIncludeBuildLogs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTrigger) ResetIncludedFiles() {
	_jsii_.InvokeVoid(
		c,
		"resetIncludedFiles",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTrigger) ResetLocation() {
	_jsii_.InvokeVoid(
		c,
		"resetLocation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTrigger) ResetName() {
	_jsii_.InvokeVoid(
		c,
		"resetName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTrigger) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTrigger) ResetProject() {
	_jsii_.InvokeVoid(
		c,
		"resetProject",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTrigger) ResetPubsubConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetPubsubConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTrigger) ResetRepositoryEventConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetRepositoryEventConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTrigger) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		c,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTrigger) ResetSourceToBuild() {
	_jsii_.InvokeVoid(
		c,
		"resetSourceToBuild",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTrigger) ResetSubstitutions() {
	_jsii_.InvokeVoid(
		c,
		"resetSubstitutions",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTrigger) ResetTags() {
	_jsii_.InvokeVoid(
		c,
		"resetTags",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTrigger) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTrigger) ResetTriggerTemplate() {
	_jsii_.InvokeVoid(
		c,
		"resetTriggerTemplate",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTrigger) ResetWebhookConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetWebhookConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudbuildTrigger) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTrigger) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTrigger) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTrigger) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTrigger) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudbuildTrigger) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

