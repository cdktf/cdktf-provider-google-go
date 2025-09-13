// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appengineflexibleappversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v16/appengineflexibleappversion/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/app_engine_flexible_app_version google_app_engine_flexible_app_version}.
type AppEngineFlexibleAppVersion interface {
	cdktf.TerraformResource
	ApiConfig() AppEngineFlexibleAppVersionApiConfigOutputReference
	ApiConfigInput() *AppEngineFlexibleAppVersionApiConfig
	AutomaticScaling() AppEngineFlexibleAppVersionAutomaticScalingOutputReference
	AutomaticScalingInput() *AppEngineFlexibleAppVersionAutomaticScaling
	BetaSettings() *map[string]*string
	SetBetaSettings(val *map[string]*string)
	BetaSettingsInput() *map[string]*string
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
	DefaultExpiration() *string
	SetDefaultExpiration(val *string)
	DefaultExpirationInput() *string
	DeleteServiceOnDestroy() interface{}
	SetDeleteServiceOnDestroy(val interface{})
	DeleteServiceOnDestroyInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Deployment() AppEngineFlexibleAppVersionDeploymentOutputReference
	DeploymentInput() *AppEngineFlexibleAppVersionDeployment
	EndpointsApiService() AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference
	EndpointsApiServiceInput() *AppEngineFlexibleAppVersionEndpointsApiService
	Entrypoint() AppEngineFlexibleAppVersionEntrypointOutputReference
	EntrypointInput() *AppEngineFlexibleAppVersionEntrypoint
	EnvVariables() *map[string]*string
	SetEnvVariables(val *map[string]*string)
	EnvVariablesInput() *map[string]*string
	FlexibleRuntimeSettings() AppEngineFlexibleAppVersionFlexibleRuntimeSettingsOutputReference
	FlexibleRuntimeSettingsInput() *AppEngineFlexibleAppVersionFlexibleRuntimeSettings
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Handlers() AppEngineFlexibleAppVersionHandlersList
	HandlersInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	InboundServices() *[]*string
	SetInboundServices(val *[]*string)
	InboundServicesInput() *[]*string
	InstanceClass() *string
	SetInstanceClass(val *string)
	InstanceClassInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LivenessCheck() AppEngineFlexibleAppVersionLivenessCheckOutputReference
	LivenessCheckInput() *AppEngineFlexibleAppVersionLivenessCheck
	ManualScaling() AppEngineFlexibleAppVersionManualScalingOutputReference
	ManualScalingInput() *AppEngineFlexibleAppVersionManualScaling
	Name() *string
	Network() AppEngineFlexibleAppVersionNetworkOutputReference
	NetworkInput() *AppEngineFlexibleAppVersionNetwork
	NobuildFilesRegex() *string
	SetNobuildFilesRegex(val *string)
	NobuildFilesRegexInput() *string
	// The tree node.
	Node() constructs.Node
	NoopOnDestroy() interface{}
	SetNoopOnDestroy(val interface{})
	NoopOnDestroyInput() interface{}
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
	ReadinessCheck() AppEngineFlexibleAppVersionReadinessCheckOutputReference
	ReadinessCheckInput() *AppEngineFlexibleAppVersionReadinessCheck
	Resources() AppEngineFlexibleAppVersionResourcesOutputReference
	ResourcesInput() *AppEngineFlexibleAppVersionResources
	Runtime() *string
	SetRuntime(val *string)
	RuntimeApiVersion() *string
	SetRuntimeApiVersion(val *string)
	RuntimeApiVersionInput() *string
	RuntimeChannel() *string
	SetRuntimeChannel(val *string)
	RuntimeChannelInput() *string
	RuntimeInput() *string
	RuntimeMainExecutablePath() *string
	SetRuntimeMainExecutablePath(val *string)
	RuntimeMainExecutablePathInput() *string
	Service() *string
	SetService(val *string)
	ServiceAccount() *string
	SetServiceAccount(val *string)
	ServiceAccountInput() *string
	ServiceInput() *string
	ServingStatus() *string
	SetServingStatus(val *string)
	ServingStatusInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() AppEngineFlexibleAppVersionTimeoutsOutputReference
	TimeoutsInput() interface{}
	VersionId() *string
	SetVersionId(val *string)
	VersionIdInput() *string
	VpcAccessConnector() AppEngineFlexibleAppVersionVpcAccessConnectorOutputReference
	VpcAccessConnectorInput() *AppEngineFlexibleAppVersionVpcAccessConnector
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
	PutApiConfig(value *AppEngineFlexibleAppVersionApiConfig)
	PutAutomaticScaling(value *AppEngineFlexibleAppVersionAutomaticScaling)
	PutDeployment(value *AppEngineFlexibleAppVersionDeployment)
	PutEndpointsApiService(value *AppEngineFlexibleAppVersionEndpointsApiService)
	PutEntrypoint(value *AppEngineFlexibleAppVersionEntrypoint)
	PutFlexibleRuntimeSettings(value *AppEngineFlexibleAppVersionFlexibleRuntimeSettings)
	PutHandlers(value interface{})
	PutLivenessCheck(value *AppEngineFlexibleAppVersionLivenessCheck)
	PutManualScaling(value *AppEngineFlexibleAppVersionManualScaling)
	PutNetwork(value *AppEngineFlexibleAppVersionNetwork)
	PutReadinessCheck(value *AppEngineFlexibleAppVersionReadinessCheck)
	PutResources(value *AppEngineFlexibleAppVersionResources)
	PutTimeouts(value *AppEngineFlexibleAppVersionTimeouts)
	PutVpcAccessConnector(value *AppEngineFlexibleAppVersionVpcAccessConnector)
	ResetApiConfig()
	ResetAutomaticScaling()
	ResetBetaSettings()
	ResetDefaultExpiration()
	ResetDeleteServiceOnDestroy()
	ResetDeployment()
	ResetEndpointsApiService()
	ResetEntrypoint()
	ResetEnvVariables()
	ResetFlexibleRuntimeSettings()
	ResetHandlers()
	ResetId()
	ResetInboundServices()
	ResetInstanceClass()
	ResetManualScaling()
	ResetNetwork()
	ResetNobuildFilesRegex()
	ResetNoopOnDestroy()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetResources()
	ResetRuntimeApiVersion()
	ResetRuntimeChannel()
	ResetRuntimeMainExecutablePath()
	ResetServiceAccount()
	ResetServingStatus()
	ResetTimeouts()
	ResetVersionId()
	ResetVpcAccessConnector()
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

// The jsii proxy struct for AppEngineFlexibleAppVersion
type jsiiProxy_AppEngineFlexibleAppVersion struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) ApiConfig() AppEngineFlexibleAppVersionApiConfigOutputReference {
	var returns AppEngineFlexibleAppVersionApiConfigOutputReference
	_jsii_.Get(
		j,
		"apiConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) ApiConfigInput() *AppEngineFlexibleAppVersionApiConfig {
	var returns *AppEngineFlexibleAppVersionApiConfig
	_jsii_.Get(
		j,
		"apiConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) AutomaticScaling() AppEngineFlexibleAppVersionAutomaticScalingOutputReference {
	var returns AppEngineFlexibleAppVersionAutomaticScalingOutputReference
	_jsii_.Get(
		j,
		"automaticScaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) AutomaticScalingInput() *AppEngineFlexibleAppVersionAutomaticScaling {
	var returns *AppEngineFlexibleAppVersionAutomaticScaling
	_jsii_.Get(
		j,
		"automaticScalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) BetaSettings() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"betaSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) BetaSettingsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"betaSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) DefaultExpiration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultExpiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) DefaultExpirationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultExpirationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) DeleteServiceOnDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteServiceOnDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) DeleteServiceOnDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteServiceOnDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) Deployment() AppEngineFlexibleAppVersionDeploymentOutputReference {
	var returns AppEngineFlexibleAppVersionDeploymentOutputReference
	_jsii_.Get(
		j,
		"deployment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) DeploymentInput() *AppEngineFlexibleAppVersionDeployment {
	var returns *AppEngineFlexibleAppVersionDeployment
	_jsii_.Get(
		j,
		"deploymentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) EndpointsApiService() AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference {
	var returns AppEngineFlexibleAppVersionEndpointsApiServiceOutputReference
	_jsii_.Get(
		j,
		"endpointsApiService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) EndpointsApiServiceInput() *AppEngineFlexibleAppVersionEndpointsApiService {
	var returns *AppEngineFlexibleAppVersionEndpointsApiService
	_jsii_.Get(
		j,
		"endpointsApiServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) Entrypoint() AppEngineFlexibleAppVersionEntrypointOutputReference {
	var returns AppEngineFlexibleAppVersionEntrypointOutputReference
	_jsii_.Get(
		j,
		"entrypoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) EntrypointInput() *AppEngineFlexibleAppVersionEntrypoint {
	var returns *AppEngineFlexibleAppVersionEntrypoint
	_jsii_.Get(
		j,
		"entrypointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) EnvVariables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"envVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) EnvVariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"envVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) FlexibleRuntimeSettings() AppEngineFlexibleAppVersionFlexibleRuntimeSettingsOutputReference {
	var returns AppEngineFlexibleAppVersionFlexibleRuntimeSettingsOutputReference
	_jsii_.Get(
		j,
		"flexibleRuntimeSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) FlexibleRuntimeSettingsInput() *AppEngineFlexibleAppVersionFlexibleRuntimeSettings {
	var returns *AppEngineFlexibleAppVersionFlexibleRuntimeSettings
	_jsii_.Get(
		j,
		"flexibleRuntimeSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) Handlers() AppEngineFlexibleAppVersionHandlersList {
	var returns AppEngineFlexibleAppVersionHandlersList
	_jsii_.Get(
		j,
		"handlers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) HandlersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"handlersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) InboundServices() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inboundServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) InboundServicesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inboundServicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) InstanceClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) InstanceClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) LivenessCheck() AppEngineFlexibleAppVersionLivenessCheckOutputReference {
	var returns AppEngineFlexibleAppVersionLivenessCheckOutputReference
	_jsii_.Get(
		j,
		"livenessCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) LivenessCheckInput() *AppEngineFlexibleAppVersionLivenessCheck {
	var returns *AppEngineFlexibleAppVersionLivenessCheck
	_jsii_.Get(
		j,
		"livenessCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) ManualScaling() AppEngineFlexibleAppVersionManualScalingOutputReference {
	var returns AppEngineFlexibleAppVersionManualScalingOutputReference
	_jsii_.Get(
		j,
		"manualScaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) ManualScalingInput() *AppEngineFlexibleAppVersionManualScaling {
	var returns *AppEngineFlexibleAppVersionManualScaling
	_jsii_.Get(
		j,
		"manualScalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) Network() AppEngineFlexibleAppVersionNetworkOutputReference {
	var returns AppEngineFlexibleAppVersionNetworkOutputReference
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) NetworkInput() *AppEngineFlexibleAppVersionNetwork {
	var returns *AppEngineFlexibleAppVersionNetwork
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) NobuildFilesRegex() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nobuildFilesRegex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) NobuildFilesRegexInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nobuildFilesRegexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) NoopOnDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noopOnDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) NoopOnDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noopOnDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) ReadinessCheck() AppEngineFlexibleAppVersionReadinessCheckOutputReference {
	var returns AppEngineFlexibleAppVersionReadinessCheckOutputReference
	_jsii_.Get(
		j,
		"readinessCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) ReadinessCheckInput() *AppEngineFlexibleAppVersionReadinessCheck {
	var returns *AppEngineFlexibleAppVersionReadinessCheck
	_jsii_.Get(
		j,
		"readinessCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) Resources() AppEngineFlexibleAppVersionResourcesOutputReference {
	var returns AppEngineFlexibleAppVersionResourcesOutputReference
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) ResourcesInput() *AppEngineFlexibleAppVersionResources {
	var returns *AppEngineFlexibleAppVersionResources
	_jsii_.Get(
		j,
		"resourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) Runtime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) RuntimeApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeApiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) RuntimeApiVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeApiVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) RuntimeChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) RuntimeChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) RuntimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) RuntimeMainExecutablePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeMainExecutablePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) RuntimeMainExecutablePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeMainExecutablePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) Service() *string {
	var returns *string
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) ServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) ServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) ServingStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servingStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) ServingStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servingStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) Timeouts() AppEngineFlexibleAppVersionTimeoutsOutputReference {
	var returns AppEngineFlexibleAppVersionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) VersionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) VersionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) VpcAccessConnector() AppEngineFlexibleAppVersionVpcAccessConnectorOutputReference {
	var returns AppEngineFlexibleAppVersionVpcAccessConnectorOutputReference
	_jsii_.Get(
		j,
		"vpcAccessConnector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion) VpcAccessConnectorInput() *AppEngineFlexibleAppVersionVpcAccessConnector {
	var returns *AppEngineFlexibleAppVersionVpcAccessConnector
	_jsii_.Get(
		j,
		"vpcAccessConnectorInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/app_engine_flexible_app_version google_app_engine_flexible_app_version} Resource.
func NewAppEngineFlexibleAppVersion(scope constructs.Construct, id *string, config *AppEngineFlexibleAppVersionConfig) AppEngineFlexibleAppVersion {
	_init_.Initialize()

	if err := validateNewAppEngineFlexibleAppVersionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppEngineFlexibleAppVersion{}

	_jsii_.Create(
		"@cdktf/provider-google.appEngineFlexibleAppVersion.AppEngineFlexibleAppVersion",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/app_engine_flexible_app_version google_app_engine_flexible_app_version} Resource.
func NewAppEngineFlexibleAppVersion_Override(a AppEngineFlexibleAppVersion, scope constructs.Construct, id *string, config *AppEngineFlexibleAppVersionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.appEngineFlexibleAppVersion.AppEngineFlexibleAppVersion",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion)SetBetaSettings(val *map[string]*string) {
	if err := j.validateSetBetaSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"betaSettings",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion)SetDefaultExpiration(val *string) {
	if err := j.validateSetDefaultExpirationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultExpiration",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion)SetDeleteServiceOnDestroy(val interface{}) {
	if err := j.validateSetDeleteServiceOnDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteServiceOnDestroy",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion)SetEnvVariables(val *map[string]*string) {
	if err := j.validateSetEnvVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"envVariables",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion)SetInboundServices(val *[]*string) {
	if err := j.validateSetInboundServicesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inboundServices",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion)SetInstanceClass(val *string) {
	if err := j.validateSetInstanceClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceClass",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion)SetNobuildFilesRegex(val *string) {
	if err := j.validateSetNobuildFilesRegexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nobuildFilesRegex",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion)SetNoopOnDestroy(val interface{}) {
	if err := j.validateSetNoopOnDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noopOnDestroy",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion)SetRuntime(val *string) {
	if err := j.validateSetRuntimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtime",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion)SetRuntimeApiVersion(val *string) {
	if err := j.validateSetRuntimeApiVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeApiVersion",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion)SetRuntimeChannel(val *string) {
	if err := j.validateSetRuntimeChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeChannel",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion)SetRuntimeMainExecutablePath(val *string) {
	if err := j.validateSetRuntimeMainExecutablePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeMainExecutablePath",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion)SetService(val *string) {
	if err := j.validateSetServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"service",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion)SetServiceAccount(val *string) {
	if err := j.validateSetServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccount",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion)SetServingStatus(val *string) {
	if err := j.validateSetServingStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"servingStatus",
		val,
	)
}

func (j *jsiiProxy_AppEngineFlexibleAppVersion)SetVersionId(val *string) {
	if err := j.validateSetVersionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"versionId",
		val,
	)
}

// Generates CDKTF code for importing a AppEngineFlexibleAppVersion resource upon running "cdktf plan <stack-name>".
func AppEngineFlexibleAppVersion_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateAppEngineFlexibleAppVersion_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.appEngineFlexibleAppVersion.AppEngineFlexibleAppVersion",
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
func AppEngineFlexibleAppVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppEngineFlexibleAppVersion_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.appEngineFlexibleAppVersion.AppEngineFlexibleAppVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AppEngineFlexibleAppVersion_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppEngineFlexibleAppVersion_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.appEngineFlexibleAppVersion.AppEngineFlexibleAppVersion",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AppEngineFlexibleAppVersion_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAppEngineFlexibleAppVersion_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.appEngineFlexibleAppVersion.AppEngineFlexibleAppVersion",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AppEngineFlexibleAppVersion_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.appEngineFlexibleAppVersion.AppEngineFlexibleAppVersion",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersion) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersion) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersion) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersion) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersion) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersion) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersion) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersion) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersion) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppEngineFlexibleAppVersion) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) PutApiConfig(value *AppEngineFlexibleAppVersionApiConfig) {
	if err := a.validatePutApiConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putApiConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) PutAutomaticScaling(value *AppEngineFlexibleAppVersionAutomaticScaling) {
	if err := a.validatePutAutomaticScalingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAutomaticScaling",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) PutDeployment(value *AppEngineFlexibleAppVersionDeployment) {
	if err := a.validatePutDeploymentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDeployment",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) PutEndpointsApiService(value *AppEngineFlexibleAppVersionEndpointsApiService) {
	if err := a.validatePutEndpointsApiServiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEndpointsApiService",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) PutEntrypoint(value *AppEngineFlexibleAppVersionEntrypoint) {
	if err := a.validatePutEntrypointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEntrypoint",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) PutFlexibleRuntimeSettings(value *AppEngineFlexibleAppVersionFlexibleRuntimeSettings) {
	if err := a.validatePutFlexibleRuntimeSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFlexibleRuntimeSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) PutHandlers(value interface{}) {
	if err := a.validatePutHandlersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHandlers",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) PutLivenessCheck(value *AppEngineFlexibleAppVersionLivenessCheck) {
	if err := a.validatePutLivenessCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLivenessCheck",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) PutManualScaling(value *AppEngineFlexibleAppVersionManualScaling) {
	if err := a.validatePutManualScalingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putManualScaling",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) PutNetwork(value *AppEngineFlexibleAppVersionNetwork) {
	if err := a.validatePutNetworkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNetwork",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) PutReadinessCheck(value *AppEngineFlexibleAppVersionReadinessCheck) {
	if err := a.validatePutReadinessCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putReadinessCheck",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) PutResources(value *AppEngineFlexibleAppVersionResources) {
	if err := a.validatePutResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResources",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) PutTimeouts(value *AppEngineFlexibleAppVersionTimeouts) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) PutVpcAccessConnector(value *AppEngineFlexibleAppVersionVpcAccessConnector) {
	if err := a.validatePutVpcAccessConnectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVpcAccessConnector",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) ResetApiConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetApiConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) ResetAutomaticScaling() {
	_jsii_.InvokeVoid(
		a,
		"resetAutomaticScaling",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) ResetBetaSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetBetaSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) ResetDefaultExpiration() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultExpiration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) ResetDeleteServiceOnDestroy() {
	_jsii_.InvokeVoid(
		a,
		"resetDeleteServiceOnDestroy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) ResetDeployment() {
	_jsii_.InvokeVoid(
		a,
		"resetDeployment",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) ResetEndpointsApiService() {
	_jsii_.InvokeVoid(
		a,
		"resetEndpointsApiService",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) ResetEntrypoint() {
	_jsii_.InvokeVoid(
		a,
		"resetEntrypoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) ResetEnvVariables() {
	_jsii_.InvokeVoid(
		a,
		"resetEnvVariables",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) ResetFlexibleRuntimeSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetFlexibleRuntimeSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) ResetHandlers() {
	_jsii_.InvokeVoid(
		a,
		"resetHandlers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) ResetInboundServices() {
	_jsii_.InvokeVoid(
		a,
		"resetInboundServices",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) ResetInstanceClass() {
	_jsii_.InvokeVoid(
		a,
		"resetInstanceClass",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) ResetManualScaling() {
	_jsii_.InvokeVoid(
		a,
		"resetManualScaling",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) ResetNetwork() {
	_jsii_.InvokeVoid(
		a,
		"resetNetwork",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) ResetNobuildFilesRegex() {
	_jsii_.InvokeVoid(
		a,
		"resetNobuildFilesRegex",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) ResetNoopOnDestroy() {
	_jsii_.InvokeVoid(
		a,
		"resetNoopOnDestroy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) ResetProject() {
	_jsii_.InvokeVoid(
		a,
		"resetProject",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) ResetResources() {
	_jsii_.InvokeVoid(
		a,
		"resetResources",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) ResetRuntimeApiVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetRuntimeApiVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) ResetRuntimeChannel() {
	_jsii_.InvokeVoid(
		a,
		"resetRuntimeChannel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) ResetRuntimeMainExecutablePath() {
	_jsii_.InvokeVoid(
		a,
		"resetRuntimeMainExecutablePath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) ResetServingStatus() {
	_jsii_.InvokeVoid(
		a,
		"resetServingStatus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) ResetVersionId() {
	_jsii_.InvokeVoid(
		a,
		"resetVersionId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) ResetVpcAccessConnector() {
	_jsii_.InvokeVoid(
		a,
		"resetVpcAccessConnector",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppEngineFlexibleAppVersion) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

