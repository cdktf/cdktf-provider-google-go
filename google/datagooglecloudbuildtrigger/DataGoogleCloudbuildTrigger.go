// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagooglecloudbuildtrigger

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v14/datagooglecloudbuildtrigger/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/data-sources/cloudbuild_trigger google_cloudbuild_trigger}.
type DataGoogleCloudbuildTrigger interface {
	cdktf.TerraformDataSource
	ApprovalConfig() DataGoogleCloudbuildTriggerApprovalConfigList
	BitbucketServerTriggerConfig() DataGoogleCloudbuildTriggerBitbucketServerTriggerConfigList
	BuildAttribute() DataGoogleCloudbuildTriggerBuildList
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	Disabled() cdktf.IResolvable
	Filename() *string
	Filter() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GitFileSource() DataGoogleCloudbuildTriggerGitFileSourceList
	Github() DataGoogleCloudbuildTriggerGithubList
	Id() *string
	SetId(val *string)
	IdInput() *string
	IgnoredFiles() *[]*string
	IncludeBuildLogs() *string
	IncludedFiles() *[]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
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
	PubsubConfig() DataGoogleCloudbuildTriggerPubsubConfigList
	// Experimental.
	RawOverrides() interface{}
	RepositoryEventConfig() DataGoogleCloudbuildTriggerRepositoryEventConfigList
	ServiceAccount() *string
	SourceToBuild() DataGoogleCloudbuildTriggerSourceToBuildList
	Substitutions() cdktf.StringMap
	Tags() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TriggerId() *string
	SetTriggerId(val *string)
	TriggerIdInput() *string
	TriggerTemplate() DataGoogleCloudbuildTriggerTriggerTemplateList
	WebhookConfig() DataGoogleCloudbuildTriggerWebhookConfigList
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

// The jsii proxy struct for DataGoogleCloudbuildTrigger
type jsiiProxy_DataGoogleCloudbuildTrigger struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger) ApprovalConfig() DataGoogleCloudbuildTriggerApprovalConfigList {
	var returns DataGoogleCloudbuildTriggerApprovalConfigList
	_jsii_.Get(
		j,
		"approvalConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger) BitbucketServerTriggerConfig() DataGoogleCloudbuildTriggerBitbucketServerTriggerConfigList {
	var returns DataGoogleCloudbuildTriggerBitbucketServerTriggerConfigList
	_jsii_.Get(
		j,
		"bitbucketServerTriggerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger) BuildAttribute() DataGoogleCloudbuildTriggerBuildList {
	var returns DataGoogleCloudbuildTriggerBuildList
	_jsii_.Get(
		j,
		"buildAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger) Disabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"disabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger) Filename() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filename",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger) Filter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger) GitFileSource() DataGoogleCloudbuildTriggerGitFileSourceList {
	var returns DataGoogleCloudbuildTriggerGitFileSourceList
	_jsii_.Get(
		j,
		"gitFileSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger) Github() DataGoogleCloudbuildTriggerGithubList {
	var returns DataGoogleCloudbuildTriggerGithubList
	_jsii_.Get(
		j,
		"github",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger) IgnoredFiles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ignoredFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger) IncludeBuildLogs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"includeBuildLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger) IncludedFiles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger) PubsubConfig() DataGoogleCloudbuildTriggerPubsubConfigList {
	var returns DataGoogleCloudbuildTriggerPubsubConfigList
	_jsii_.Get(
		j,
		"pubsubConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger) RepositoryEventConfig() DataGoogleCloudbuildTriggerRepositoryEventConfigList {
	var returns DataGoogleCloudbuildTriggerRepositoryEventConfigList
	_jsii_.Get(
		j,
		"repositoryEventConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger) SourceToBuild() DataGoogleCloudbuildTriggerSourceToBuildList {
	var returns DataGoogleCloudbuildTriggerSourceToBuildList
	_jsii_.Get(
		j,
		"sourceToBuild",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger) Substitutions() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"substitutions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger) TriggerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"triggerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger) TriggerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"triggerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger) TriggerTemplate() DataGoogleCloudbuildTriggerTriggerTemplateList {
	var returns DataGoogleCloudbuildTriggerTriggerTemplateList
	_jsii_.Get(
		j,
		"triggerTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger) WebhookConfig() DataGoogleCloudbuildTriggerWebhookConfigList {
	var returns DataGoogleCloudbuildTriggerWebhookConfigList
	_jsii_.Get(
		j,
		"webhookConfig",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/data-sources/cloudbuild_trigger google_cloudbuild_trigger} Data Source.
func NewDataGoogleCloudbuildTrigger(scope constructs.Construct, id *string, config *DataGoogleCloudbuildTriggerConfig) DataGoogleCloudbuildTrigger {
	_init_.Initialize()

	if err := validateNewDataGoogleCloudbuildTriggerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGoogleCloudbuildTrigger{}

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleCloudbuildTrigger.DataGoogleCloudbuildTrigger",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/data-sources/cloudbuild_trigger google_cloudbuild_trigger} Data Source.
func NewDataGoogleCloudbuildTrigger_Override(d DataGoogleCloudbuildTrigger, scope constructs.Construct, id *string, config *DataGoogleCloudbuildTriggerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleCloudbuildTrigger.DataGoogleCloudbuildTrigger",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataGoogleCloudbuildTrigger)SetTriggerId(val *string) {
	if err := j.validateSetTriggerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"triggerId",
		val,
	)
}

// Generates CDKTF code for importing a DataGoogleCloudbuildTrigger resource upon running "cdktf plan <stack-name>".
func DataGoogleCloudbuildTrigger_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataGoogleCloudbuildTrigger_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dataGoogleCloudbuildTrigger.DataGoogleCloudbuildTrigger",
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
func DataGoogleCloudbuildTrigger_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGoogleCloudbuildTrigger_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dataGoogleCloudbuildTrigger.DataGoogleCloudbuildTrigger",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataGoogleCloudbuildTrigger_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGoogleCloudbuildTrigger_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dataGoogleCloudbuildTrigger.DataGoogleCloudbuildTrigger",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataGoogleCloudbuildTrigger_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGoogleCloudbuildTrigger_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dataGoogleCloudbuildTrigger.DataGoogleCloudbuildTrigger",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataGoogleCloudbuildTrigger_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.dataGoogleCloudbuildTrigger.DataGoogleCloudbuildTrigger",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataGoogleCloudbuildTrigger) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataGoogleCloudbuildTrigger) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataGoogleCloudbuildTrigger) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGoogleCloudbuildTrigger) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataGoogleCloudbuildTrigger) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataGoogleCloudbuildTrigger) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataGoogleCloudbuildTrigger) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataGoogleCloudbuildTrigger) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataGoogleCloudbuildTrigger) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataGoogleCloudbuildTrigger) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataGoogleCloudbuildTrigger) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGoogleCloudbuildTrigger) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataGoogleCloudbuildTrigger) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGoogleCloudbuildTrigger) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGoogleCloudbuildTrigger) ResetProject() {
	_jsii_.InvokeVoid(
		d,
		"resetProject",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGoogleCloudbuildTrigger) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleCloudbuildTrigger) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleCloudbuildTrigger) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleCloudbuildTrigger) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleCloudbuildTrigger) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleCloudbuildTrigger) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

