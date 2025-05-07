// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package developerconnectconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v14/developerconnectconnection/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/developer_connect_connection google_developer_connect_connection}.
type DeveloperConnectConnection interface {
	cdktf.TerraformResource
	Annotations() *map[string]*string
	SetAnnotations(val *map[string]*string)
	AnnotationsInput() *map[string]*string
	BitbucketCloudConfig() DeveloperConnectConnectionBitbucketCloudConfigOutputReference
	BitbucketCloudConfigInput() *DeveloperConnectConnectionBitbucketCloudConfig
	BitbucketDataCenterConfig() DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference
	BitbucketDataCenterConfigInput() *DeveloperConnectConnectionBitbucketDataCenterConfig
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectionId() *string
	SetConnectionId(val *string)
	ConnectionIdInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreateTime() *string
	CryptoKeyConfig() DeveloperConnectConnectionCryptoKeyConfigOutputReference
	CryptoKeyConfigInput() *DeveloperConnectConnectionCryptoKeyConfig
	DeleteTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Disabled() interface{}
	SetDisabled(val interface{})
	DisabledInput() interface{}
	EffectiveAnnotations() cdktf.StringMap
	EffectiveLabels() cdktf.StringMap
	Etag() *string
	SetEtag(val *string)
	EtagInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GithubConfig() DeveloperConnectConnectionGithubConfigOutputReference
	GithubConfigInput() *DeveloperConnectConnectionGithubConfig
	GithubEnterpriseConfig() DeveloperConnectConnectionGithubEnterpriseConfigOutputReference
	GithubEnterpriseConfigInput() *DeveloperConnectConnectionGithubEnterpriseConfig
	GitlabConfig() DeveloperConnectConnectionGitlabConfigOutputReference
	GitlabConfigInput() *DeveloperConnectConnectionGitlabConfig
	GitlabEnterpriseConfig() DeveloperConnectConnectionGitlabEnterpriseConfigOutputReference
	GitlabEnterpriseConfigInput() *DeveloperConnectConnectionGitlabEnterpriseConfig
	Id() *string
	SetId(val *string)
	IdInput() *string
	InstallationState() DeveloperConnectConnectionInstallationStateList
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
	Reconciling() cdktf.IResolvable
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DeveloperConnectConnectionTimeoutsOutputReference
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
	PutBitbucketCloudConfig(value *DeveloperConnectConnectionBitbucketCloudConfig)
	PutBitbucketDataCenterConfig(value *DeveloperConnectConnectionBitbucketDataCenterConfig)
	PutCryptoKeyConfig(value *DeveloperConnectConnectionCryptoKeyConfig)
	PutGithubConfig(value *DeveloperConnectConnectionGithubConfig)
	PutGithubEnterpriseConfig(value *DeveloperConnectConnectionGithubEnterpriseConfig)
	PutGitlabConfig(value *DeveloperConnectConnectionGitlabConfig)
	PutGitlabEnterpriseConfig(value *DeveloperConnectConnectionGitlabEnterpriseConfig)
	PutTimeouts(value *DeveloperConnectConnectionTimeouts)
	ResetAnnotations()
	ResetBitbucketCloudConfig()
	ResetBitbucketDataCenterConfig()
	ResetCryptoKeyConfig()
	ResetDisabled()
	ResetEtag()
	ResetGithubConfig()
	ResetGithubEnterpriseConfig()
	ResetGitlabConfig()
	ResetGitlabEnterpriseConfig()
	ResetId()
	ResetLabels()
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

// The jsii proxy struct for DeveloperConnectConnection
type jsiiProxy_DeveloperConnectConnection struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DeveloperConnectConnection) Annotations() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) AnnotationsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) BitbucketCloudConfig() DeveloperConnectConnectionBitbucketCloudConfigOutputReference {
	var returns DeveloperConnectConnectionBitbucketCloudConfigOutputReference
	_jsii_.Get(
		j,
		"bitbucketCloudConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) BitbucketCloudConfigInput() *DeveloperConnectConnectionBitbucketCloudConfig {
	var returns *DeveloperConnectConnectionBitbucketCloudConfig
	_jsii_.Get(
		j,
		"bitbucketCloudConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) BitbucketDataCenterConfig() DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference {
	var returns DeveloperConnectConnectionBitbucketDataCenterConfigOutputReference
	_jsii_.Get(
		j,
		"bitbucketDataCenterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) BitbucketDataCenterConfigInput() *DeveloperConnectConnectionBitbucketDataCenterConfig {
	var returns *DeveloperConnectConnectionBitbucketDataCenterConfig
	_jsii_.Get(
		j,
		"bitbucketDataCenterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) ConnectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) ConnectionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) CryptoKeyConfig() DeveloperConnectConnectionCryptoKeyConfigOutputReference {
	var returns DeveloperConnectConnectionCryptoKeyConfigOutputReference
	_jsii_.Get(
		j,
		"cryptoKeyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) CryptoKeyConfigInput() *DeveloperConnectConnectionCryptoKeyConfig {
	var returns *DeveloperConnectConnectionCryptoKeyConfig
	_jsii_.Get(
		j,
		"cryptoKeyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) DeleteTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) Disabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) DisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) EffectiveAnnotations() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) EtagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) GithubConfig() DeveloperConnectConnectionGithubConfigOutputReference {
	var returns DeveloperConnectConnectionGithubConfigOutputReference
	_jsii_.Get(
		j,
		"githubConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) GithubConfigInput() *DeveloperConnectConnectionGithubConfig {
	var returns *DeveloperConnectConnectionGithubConfig
	_jsii_.Get(
		j,
		"githubConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) GithubEnterpriseConfig() DeveloperConnectConnectionGithubEnterpriseConfigOutputReference {
	var returns DeveloperConnectConnectionGithubEnterpriseConfigOutputReference
	_jsii_.Get(
		j,
		"githubEnterpriseConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) GithubEnterpriseConfigInput() *DeveloperConnectConnectionGithubEnterpriseConfig {
	var returns *DeveloperConnectConnectionGithubEnterpriseConfig
	_jsii_.Get(
		j,
		"githubEnterpriseConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) GitlabConfig() DeveloperConnectConnectionGitlabConfigOutputReference {
	var returns DeveloperConnectConnectionGitlabConfigOutputReference
	_jsii_.Get(
		j,
		"gitlabConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) GitlabConfigInput() *DeveloperConnectConnectionGitlabConfig {
	var returns *DeveloperConnectConnectionGitlabConfig
	_jsii_.Get(
		j,
		"gitlabConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) GitlabEnterpriseConfig() DeveloperConnectConnectionGitlabEnterpriseConfigOutputReference {
	var returns DeveloperConnectConnectionGitlabEnterpriseConfigOutputReference
	_jsii_.Get(
		j,
		"gitlabEnterpriseConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) GitlabEnterpriseConfigInput() *DeveloperConnectConnectionGitlabEnterpriseConfig {
	var returns *DeveloperConnectConnectionGitlabEnterpriseConfig
	_jsii_.Get(
		j,
		"gitlabEnterpriseConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) InstallationState() DeveloperConnectConnectionInstallationStateList {
	var returns DeveloperConnectConnectionInstallationStateList
	_jsii_.Get(
		j,
		"installationState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) Reconciling() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"reconciling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) Timeouts() DeveloperConnectConnectionTimeoutsOutputReference {
	var returns DeveloperConnectConnectionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) Uid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeveloperConnectConnection) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/developer_connect_connection google_developer_connect_connection} Resource.
func NewDeveloperConnectConnection(scope constructs.Construct, id *string, config *DeveloperConnectConnectionConfig) DeveloperConnectConnection {
	_init_.Initialize()

	if err := validateNewDeveloperConnectConnectionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DeveloperConnectConnection{}

	_jsii_.Create(
		"@cdktf/provider-google.developerConnectConnection.DeveloperConnectConnection",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/developer_connect_connection google_developer_connect_connection} Resource.
func NewDeveloperConnectConnection_Override(d DeveloperConnectConnection, scope constructs.Construct, id *string, config *DeveloperConnectConnectionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.developerConnectConnection.DeveloperConnectConnection",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DeveloperConnectConnection)SetAnnotations(val *map[string]*string) {
	if err := j.validateSetAnnotationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotations",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectConnection)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectConnection)SetConnectionId(val *string) {
	if err := j.validateSetConnectionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionId",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectConnection)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectConnection)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectConnection)SetDisabled(val interface{}) {
	if err := j.validateSetDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disabled",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectConnection)SetEtag(val *string) {
	if err := j.validateSetEtagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"etag",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectConnection)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectConnection)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectConnection)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectConnection)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectConnection)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectConnection)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectConnection)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DeveloperConnectConnection)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a DeveloperConnectConnection resource upon running "cdktf plan <stack-name>".
func DeveloperConnectConnection_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDeveloperConnectConnection_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.developerConnectConnection.DeveloperConnectConnection",
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
func DeveloperConnectConnection_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDeveloperConnectConnection_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.developerConnectConnection.DeveloperConnectConnection",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DeveloperConnectConnection_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDeveloperConnectConnection_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.developerConnectConnection.DeveloperConnectConnection",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DeveloperConnectConnection_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDeveloperConnectConnection_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.developerConnectConnection.DeveloperConnectConnection",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DeveloperConnectConnection_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.developerConnectConnection.DeveloperConnectConnection",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DeveloperConnectConnection) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DeveloperConnectConnection) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DeveloperConnectConnection) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DeveloperConnectConnection) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DeveloperConnectConnection) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DeveloperConnectConnection) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DeveloperConnectConnection) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DeveloperConnectConnection) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DeveloperConnectConnection) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DeveloperConnectConnection) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DeveloperConnectConnection) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DeveloperConnectConnection) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectConnection) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DeveloperConnectConnection) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DeveloperConnectConnection) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DeveloperConnectConnection) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DeveloperConnectConnection) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DeveloperConnectConnection) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DeveloperConnectConnection) PutBitbucketCloudConfig(value *DeveloperConnectConnectionBitbucketCloudConfig) {
	if err := d.validatePutBitbucketCloudConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putBitbucketCloudConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeveloperConnectConnection) PutBitbucketDataCenterConfig(value *DeveloperConnectConnectionBitbucketDataCenterConfig) {
	if err := d.validatePutBitbucketDataCenterConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putBitbucketDataCenterConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeveloperConnectConnection) PutCryptoKeyConfig(value *DeveloperConnectConnectionCryptoKeyConfig) {
	if err := d.validatePutCryptoKeyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCryptoKeyConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeveloperConnectConnection) PutGithubConfig(value *DeveloperConnectConnectionGithubConfig) {
	if err := d.validatePutGithubConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGithubConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeveloperConnectConnection) PutGithubEnterpriseConfig(value *DeveloperConnectConnectionGithubEnterpriseConfig) {
	if err := d.validatePutGithubEnterpriseConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGithubEnterpriseConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeveloperConnectConnection) PutGitlabConfig(value *DeveloperConnectConnectionGitlabConfig) {
	if err := d.validatePutGitlabConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGitlabConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeveloperConnectConnection) PutGitlabEnterpriseConfig(value *DeveloperConnectConnectionGitlabEnterpriseConfig) {
	if err := d.validatePutGitlabEnterpriseConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGitlabEnterpriseConfig",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeveloperConnectConnection) PutTimeouts(value *DeveloperConnectConnectionTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DeveloperConnectConnection) ResetAnnotations() {
	_jsii_.InvokeVoid(
		d,
		"resetAnnotations",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeveloperConnectConnection) ResetBitbucketCloudConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetBitbucketCloudConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeveloperConnectConnection) ResetBitbucketDataCenterConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetBitbucketDataCenterConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeveloperConnectConnection) ResetCryptoKeyConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetCryptoKeyConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeveloperConnectConnection) ResetDisabled() {
	_jsii_.InvokeVoid(
		d,
		"resetDisabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeveloperConnectConnection) ResetEtag() {
	_jsii_.InvokeVoid(
		d,
		"resetEtag",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeveloperConnectConnection) ResetGithubConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetGithubConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeveloperConnectConnection) ResetGithubEnterpriseConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetGithubEnterpriseConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeveloperConnectConnection) ResetGitlabConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetGitlabConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeveloperConnectConnection) ResetGitlabEnterpriseConfig() {
	_jsii_.InvokeVoid(
		d,
		"resetGitlabEnterpriseConfig",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeveloperConnectConnection) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeveloperConnectConnection) ResetLabels() {
	_jsii_.InvokeVoid(
		d,
		"resetLabels",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeveloperConnectConnection) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeveloperConnectConnection) ResetProject() {
	_jsii_.InvokeVoid(
		d,
		"resetProject",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeveloperConnectConnection) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DeveloperConnectConnection) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectConnection) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectConnection) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectConnection) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectConnection) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeveloperConnectConnection) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

