// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudbuildv2connection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v14/cloudbuildv2connection/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/cloudbuildv2_connection google_cloudbuildv2_connection}.
type Cloudbuildv2Connection interface {
	cdktf.TerraformResource
	Annotations() *map[string]*string
	SetAnnotations(val *map[string]*string)
	AnnotationsInput() *map[string]*string
	BitbucketCloudConfig() Cloudbuildv2ConnectionBitbucketCloudConfigOutputReference
	BitbucketCloudConfigInput() *Cloudbuildv2ConnectionBitbucketCloudConfig
	BitbucketDataCenterConfig() Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference
	BitbucketDataCenterConfigInput() *Cloudbuildv2ConnectionBitbucketDataCenterConfig
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
	Disabled() interface{}
	SetDisabled(val interface{})
	DisabledInput() interface{}
	EffectiveAnnotations() cdktf.StringMap
	Etag() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GithubConfig() Cloudbuildv2ConnectionGithubConfigOutputReference
	GithubConfigInput() *Cloudbuildv2ConnectionGithubConfig
	GithubEnterpriseConfig() Cloudbuildv2ConnectionGithubEnterpriseConfigOutputReference
	GithubEnterpriseConfigInput() *Cloudbuildv2ConnectionGithubEnterpriseConfig
	GitlabConfig() Cloudbuildv2ConnectionGitlabConfigOutputReference
	GitlabConfigInput() *Cloudbuildv2ConnectionGitlabConfig
	Id() *string
	SetId(val *string)
	IdInput() *string
	InstallationState() Cloudbuildv2ConnectionInstallationStateList
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
	// Experimental.
	RawOverrides() interface{}
	Reconciling() cdktf.IResolvable
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() Cloudbuildv2ConnectionTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutBitbucketCloudConfig(value *Cloudbuildv2ConnectionBitbucketCloudConfig)
	PutBitbucketDataCenterConfig(value *Cloudbuildv2ConnectionBitbucketDataCenterConfig)
	PutGithubConfig(value *Cloudbuildv2ConnectionGithubConfig)
	PutGithubEnterpriseConfig(value *Cloudbuildv2ConnectionGithubEnterpriseConfig)
	PutGitlabConfig(value *Cloudbuildv2ConnectionGitlabConfig)
	PutTimeouts(value *Cloudbuildv2ConnectionTimeouts)
	ResetAnnotations()
	ResetBitbucketCloudConfig()
	ResetBitbucketDataCenterConfig()
	ResetDisabled()
	ResetGithubConfig()
	ResetGithubEnterpriseConfig()
	ResetGitlabConfig()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
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

// The jsii proxy struct for Cloudbuildv2Connection
type jsiiProxy_Cloudbuildv2Connection struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Cloudbuildv2Connection) Annotations() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2Connection) AnnotationsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2Connection) BitbucketCloudConfig() Cloudbuildv2ConnectionBitbucketCloudConfigOutputReference {
	var returns Cloudbuildv2ConnectionBitbucketCloudConfigOutputReference
	_jsii_.Get(
		j,
		"bitbucketCloudConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2Connection) BitbucketCloudConfigInput() *Cloudbuildv2ConnectionBitbucketCloudConfig {
	var returns *Cloudbuildv2ConnectionBitbucketCloudConfig
	_jsii_.Get(
		j,
		"bitbucketCloudConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2Connection) BitbucketDataCenterConfig() Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference {
	var returns Cloudbuildv2ConnectionBitbucketDataCenterConfigOutputReference
	_jsii_.Get(
		j,
		"bitbucketDataCenterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2Connection) BitbucketDataCenterConfigInput() *Cloudbuildv2ConnectionBitbucketDataCenterConfig {
	var returns *Cloudbuildv2ConnectionBitbucketDataCenterConfig
	_jsii_.Get(
		j,
		"bitbucketDataCenterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2Connection) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2Connection) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2Connection) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2Connection) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2Connection) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2Connection) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2Connection) Disabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2Connection) DisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2Connection) EffectiveAnnotations() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2Connection) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2Connection) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2Connection) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2Connection) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2Connection) GithubConfig() Cloudbuildv2ConnectionGithubConfigOutputReference {
	var returns Cloudbuildv2ConnectionGithubConfigOutputReference
	_jsii_.Get(
		j,
		"githubConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2Connection) GithubConfigInput() *Cloudbuildv2ConnectionGithubConfig {
	var returns *Cloudbuildv2ConnectionGithubConfig
	_jsii_.Get(
		j,
		"githubConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2Connection) GithubEnterpriseConfig() Cloudbuildv2ConnectionGithubEnterpriseConfigOutputReference {
	var returns Cloudbuildv2ConnectionGithubEnterpriseConfigOutputReference
	_jsii_.Get(
		j,
		"githubEnterpriseConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2Connection) GithubEnterpriseConfigInput() *Cloudbuildv2ConnectionGithubEnterpriseConfig {
	var returns *Cloudbuildv2ConnectionGithubEnterpriseConfig
	_jsii_.Get(
		j,
		"githubEnterpriseConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2Connection) GitlabConfig() Cloudbuildv2ConnectionGitlabConfigOutputReference {
	var returns Cloudbuildv2ConnectionGitlabConfigOutputReference
	_jsii_.Get(
		j,
		"gitlabConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2Connection) GitlabConfigInput() *Cloudbuildv2ConnectionGitlabConfig {
	var returns *Cloudbuildv2ConnectionGitlabConfig
	_jsii_.Get(
		j,
		"gitlabConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2Connection) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2Connection) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2Connection) InstallationState() Cloudbuildv2ConnectionInstallationStateList {
	var returns Cloudbuildv2ConnectionInstallationStateList
	_jsii_.Get(
		j,
		"installationState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2Connection) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2Connection) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2Connection) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2Connection) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2Connection) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2Connection) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2Connection) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2Connection) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2Connection) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2Connection) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2Connection) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2Connection) Reconciling() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"reconciling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2Connection) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2Connection) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2Connection) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2Connection) Timeouts() Cloudbuildv2ConnectionTimeoutsOutputReference {
	var returns Cloudbuildv2ConnectionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2Connection) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Cloudbuildv2Connection) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/cloudbuildv2_connection google_cloudbuildv2_connection} Resource.
func NewCloudbuildv2Connection(scope constructs.Construct, id *string, config *Cloudbuildv2ConnectionConfig) Cloudbuildv2Connection {
	_init_.Initialize()

	if err := validateNewCloudbuildv2ConnectionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Cloudbuildv2Connection{}

	_jsii_.Create(
		"@cdktf/provider-google.cloudbuildv2Connection.Cloudbuildv2Connection",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.22.0/docs/resources/cloudbuildv2_connection google_cloudbuildv2_connection} Resource.
func NewCloudbuildv2Connection_Override(c Cloudbuildv2Connection, scope constructs.Construct, id *string, config *Cloudbuildv2ConnectionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cloudbuildv2Connection.Cloudbuildv2Connection",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_Cloudbuildv2Connection)SetAnnotations(val *map[string]*string) {
	if err := j.validateSetAnnotationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotations",
		val,
	)
}

func (j *jsiiProxy_Cloudbuildv2Connection)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Cloudbuildv2Connection)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Cloudbuildv2Connection)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Cloudbuildv2Connection)SetDisabled(val interface{}) {
	if err := j.validateSetDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disabled",
		val,
	)
}

func (j *jsiiProxy_Cloudbuildv2Connection)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Cloudbuildv2Connection)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Cloudbuildv2Connection)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Cloudbuildv2Connection)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_Cloudbuildv2Connection)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Cloudbuildv2Connection)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_Cloudbuildv2Connection)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Cloudbuildv2Connection)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a Cloudbuildv2Connection resource upon running "cdktf plan <stack-name>".
func Cloudbuildv2Connection_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCloudbuildv2Connection_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.cloudbuildv2Connection.Cloudbuildv2Connection",
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
func Cloudbuildv2Connection_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudbuildv2Connection_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.cloudbuildv2Connection.Cloudbuildv2Connection",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Cloudbuildv2Connection_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudbuildv2Connection_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.cloudbuildv2Connection.Cloudbuildv2Connection",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Cloudbuildv2Connection_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCloudbuildv2Connection_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.cloudbuildv2Connection.Cloudbuildv2Connection",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Cloudbuildv2Connection_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.cloudbuildv2Connection.Cloudbuildv2Connection",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_Cloudbuildv2Connection) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_Cloudbuildv2Connection) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_Cloudbuildv2Connection) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_Cloudbuildv2Connection) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_Cloudbuildv2Connection) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_Cloudbuildv2Connection) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_Cloudbuildv2Connection) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_Cloudbuildv2Connection) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_Cloudbuildv2Connection) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_Cloudbuildv2Connection) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_Cloudbuildv2Connection) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_Cloudbuildv2Connection) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cloudbuildv2Connection) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_Cloudbuildv2Connection) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_Cloudbuildv2Connection) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_Cloudbuildv2Connection) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_Cloudbuildv2Connection) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_Cloudbuildv2Connection) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_Cloudbuildv2Connection) PutBitbucketCloudConfig(value *Cloudbuildv2ConnectionBitbucketCloudConfig) {
	if err := c.validatePutBitbucketCloudConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBitbucketCloudConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Cloudbuildv2Connection) PutBitbucketDataCenterConfig(value *Cloudbuildv2ConnectionBitbucketDataCenterConfig) {
	if err := c.validatePutBitbucketDataCenterConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBitbucketDataCenterConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Cloudbuildv2Connection) PutGithubConfig(value *Cloudbuildv2ConnectionGithubConfig) {
	if err := c.validatePutGithubConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGithubConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Cloudbuildv2Connection) PutGithubEnterpriseConfig(value *Cloudbuildv2ConnectionGithubEnterpriseConfig) {
	if err := c.validatePutGithubEnterpriseConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGithubEnterpriseConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Cloudbuildv2Connection) PutGitlabConfig(value *Cloudbuildv2ConnectionGitlabConfig) {
	if err := c.validatePutGitlabConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGitlabConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Cloudbuildv2Connection) PutTimeouts(value *Cloudbuildv2ConnectionTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_Cloudbuildv2Connection) ResetAnnotations() {
	_jsii_.InvokeVoid(
		c,
		"resetAnnotations",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudbuildv2Connection) ResetBitbucketCloudConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetBitbucketCloudConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudbuildv2Connection) ResetBitbucketDataCenterConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetBitbucketDataCenterConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudbuildv2Connection) ResetDisabled() {
	_jsii_.InvokeVoid(
		c,
		"resetDisabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudbuildv2Connection) ResetGithubConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetGithubConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudbuildv2Connection) ResetGithubEnterpriseConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetGithubEnterpriseConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudbuildv2Connection) ResetGitlabConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetGitlabConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudbuildv2Connection) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudbuildv2Connection) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudbuildv2Connection) ResetProject() {
	_jsii_.InvokeVoid(
		c,
		"resetProject",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudbuildv2Connection) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_Cloudbuildv2Connection) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cloudbuildv2Connection) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cloudbuildv2Connection) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cloudbuildv2Connection) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cloudbuildv2Connection) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Cloudbuildv2Connection) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

