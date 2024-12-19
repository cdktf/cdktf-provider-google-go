// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package integrationconnectorsconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v14/integrationconnectorsconnection/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/integration_connectors_connection google_integration_connectors_connection}.
type IntegrationConnectorsConnection interface {
	cdktf.TerraformResource
	AuthConfig() IntegrationConnectorsConnectionAuthConfigOutputReference
	AuthConfigInput() *IntegrationConnectorsConnectionAuthConfig
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ConfigVariable() IntegrationConnectorsConnectionConfigVariableList
	ConfigVariableInput() interface{}
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectionRevision() *string
	ConnectorVersion() *string
	SetConnectorVersion(val *string)
	ConnectorVersionInfraConfig() IntegrationConnectorsConnectionConnectorVersionInfraConfigList
	ConnectorVersionInput() *string
	ConnectorVersionLaunchStage() *string
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
	DestinationConfig() IntegrationConnectorsConnectionDestinationConfigList
	DestinationConfigInput() interface{}
	EffectiveLabels() cdktf.StringMap
	EventingConfig() IntegrationConnectorsConnectionEventingConfigOutputReference
	EventingConfigInput() *IntegrationConnectorsConnectionEventingConfig
	EventingEnablementType() *string
	SetEventingEnablementType(val *string)
	EventingEnablementTypeInput() *string
	EventingRuntimeData() IntegrationConnectorsConnectionEventingRuntimeDataList
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
	LockConfig() IntegrationConnectorsConnectionLockConfigOutputReference
	LockConfigInput() *IntegrationConnectorsConnectionLockConfig
	LogConfig() IntegrationConnectorsConnectionLogConfigOutputReference
	LogConfigInput() *IntegrationConnectorsConnectionLogConfig
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NodeConfig() IntegrationConnectorsConnectionNodeConfigOutputReference
	NodeConfigInput() *IntegrationConnectorsConnectionNodeConfig
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
	ServiceAccount() *string
	SetServiceAccount(val *string)
	ServiceAccountInput() *string
	ServiceDirectory() *string
	SslConfig() IntegrationConnectorsConnectionSslConfigOutputReference
	SslConfigInput() *IntegrationConnectorsConnectionSslConfig
	Status() IntegrationConnectorsConnectionStatusList
	SubscriptionType() *string
	Suspended() interface{}
	SetSuspended(val interface{})
	SuspendedInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() IntegrationConnectorsConnectionTimeoutsOutputReference
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
	PutAuthConfig(value *IntegrationConnectorsConnectionAuthConfig)
	PutConfigVariable(value interface{})
	PutDestinationConfig(value interface{})
	PutEventingConfig(value *IntegrationConnectorsConnectionEventingConfig)
	PutLockConfig(value *IntegrationConnectorsConnectionLockConfig)
	PutLogConfig(value *IntegrationConnectorsConnectionLogConfig)
	PutNodeConfig(value *IntegrationConnectorsConnectionNodeConfig)
	PutSslConfig(value *IntegrationConnectorsConnectionSslConfig)
	PutTimeouts(value *IntegrationConnectorsConnectionTimeouts)
	ResetAuthConfig()
	ResetConfigVariable()
	ResetDescription()
	ResetDestinationConfig()
	ResetEventingConfig()
	ResetEventingEnablementType()
	ResetId()
	ResetLabels()
	ResetLockConfig()
	ResetLogConfig()
	ResetNodeConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetServiceAccount()
	ResetSslConfig()
	ResetSuspended()
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

// The jsii proxy struct for IntegrationConnectorsConnection
type jsiiProxy_IntegrationConnectorsConnection struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_IntegrationConnectorsConnection) AuthConfig() IntegrationConnectorsConnectionAuthConfigOutputReference {
	var returns IntegrationConnectorsConnectionAuthConfigOutputReference
	_jsii_.Get(
		j,
		"authConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) AuthConfigInput() *IntegrationConnectorsConnectionAuthConfig {
	var returns *IntegrationConnectorsConnectionAuthConfig
	_jsii_.Get(
		j,
		"authConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) ConfigVariable() IntegrationConnectorsConnectionConfigVariableList {
	var returns IntegrationConnectorsConnectionConfigVariableList
	_jsii_.Get(
		j,
		"configVariable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) ConfigVariableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"configVariableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) ConnectionRevision() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionRevision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) ConnectorVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) ConnectorVersionInfraConfig() IntegrationConnectorsConnectionConnectorVersionInfraConfigList {
	var returns IntegrationConnectorsConnectionConnectorVersionInfraConfigList
	_jsii_.Get(
		j,
		"connectorVersionInfraConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) ConnectorVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) ConnectorVersionLaunchStage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorVersionLaunchStage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) DestinationConfig() IntegrationConnectorsConnectionDestinationConfigList {
	var returns IntegrationConnectorsConnectionDestinationConfigList
	_jsii_.Get(
		j,
		"destinationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) DestinationConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) EventingConfig() IntegrationConnectorsConnectionEventingConfigOutputReference {
	var returns IntegrationConnectorsConnectionEventingConfigOutputReference
	_jsii_.Get(
		j,
		"eventingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) EventingConfigInput() *IntegrationConnectorsConnectionEventingConfig {
	var returns *IntegrationConnectorsConnectionEventingConfig
	_jsii_.Get(
		j,
		"eventingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) EventingEnablementType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventingEnablementType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) EventingEnablementTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventingEnablementTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) EventingRuntimeData() IntegrationConnectorsConnectionEventingRuntimeDataList {
	var returns IntegrationConnectorsConnectionEventingRuntimeDataList
	_jsii_.Get(
		j,
		"eventingRuntimeData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) LockConfig() IntegrationConnectorsConnectionLockConfigOutputReference {
	var returns IntegrationConnectorsConnectionLockConfigOutputReference
	_jsii_.Get(
		j,
		"lockConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) LockConfigInput() *IntegrationConnectorsConnectionLockConfig {
	var returns *IntegrationConnectorsConnectionLockConfig
	_jsii_.Get(
		j,
		"lockConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) LogConfig() IntegrationConnectorsConnectionLogConfigOutputReference {
	var returns IntegrationConnectorsConnectionLogConfigOutputReference
	_jsii_.Get(
		j,
		"logConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) LogConfigInput() *IntegrationConnectorsConnectionLogConfig {
	var returns *IntegrationConnectorsConnectionLogConfig
	_jsii_.Get(
		j,
		"logConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) NodeConfig() IntegrationConnectorsConnectionNodeConfigOutputReference {
	var returns IntegrationConnectorsConnectionNodeConfigOutputReference
	_jsii_.Get(
		j,
		"nodeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) NodeConfigInput() *IntegrationConnectorsConnectionNodeConfig {
	var returns *IntegrationConnectorsConnectionNodeConfig
	_jsii_.Get(
		j,
		"nodeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) ServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) ServiceDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) SslConfig() IntegrationConnectorsConnectionSslConfigOutputReference {
	var returns IntegrationConnectorsConnectionSslConfigOutputReference
	_jsii_.Get(
		j,
		"sslConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) SslConfigInput() *IntegrationConnectorsConnectionSslConfig {
	var returns *IntegrationConnectorsConnectionSslConfig
	_jsii_.Get(
		j,
		"sslConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) Status() IntegrationConnectorsConnectionStatusList {
	var returns IntegrationConnectorsConnectionStatusList
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) SubscriptionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) Suspended() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suspended",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) SuspendedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suspendedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) Timeouts() IntegrationConnectorsConnectionTimeoutsOutputReference {
	var returns IntegrationConnectorsConnectionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationConnectorsConnection) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/integration_connectors_connection google_integration_connectors_connection} Resource.
func NewIntegrationConnectorsConnection(scope constructs.Construct, id *string, config *IntegrationConnectorsConnectionConfig) IntegrationConnectorsConnection {
	_init_.Initialize()

	if err := validateNewIntegrationConnectorsConnectionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_IntegrationConnectorsConnection{}

	_jsii_.Create(
		"@cdktf/provider-google.integrationConnectorsConnection.IntegrationConnectorsConnection",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/integration_connectors_connection google_integration_connectors_connection} Resource.
func NewIntegrationConnectorsConnection_Override(i IntegrationConnectorsConnection, scope constructs.Construct, id *string, config *IntegrationConnectorsConnectionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.integrationConnectorsConnection.IntegrationConnectorsConnection",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnection)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnection)SetConnectorVersion(val *string) {
	if err := j.validateSetConnectorVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectorVersion",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnection)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnection)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnection)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnection)SetEventingEnablementType(val *string) {
	if err := j.validateSetEventingEnablementTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventingEnablementType",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnection)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnection)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnection)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnection)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnection)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnection)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnection)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnection)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnection)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnection)SetServiceAccount(val *string) {
	if err := j.validateSetServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccount",
		val,
	)
}

func (j *jsiiProxy_IntegrationConnectorsConnection)SetSuspended(val interface{}) {
	if err := j.validateSetSuspendedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suspended",
		val,
	)
}

// Generates CDKTF code for importing a IntegrationConnectorsConnection resource upon running "cdktf plan <stack-name>".
func IntegrationConnectorsConnection_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateIntegrationConnectorsConnection_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.integrationConnectorsConnection.IntegrationConnectorsConnection",
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
func IntegrationConnectorsConnection_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIntegrationConnectorsConnection_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.integrationConnectorsConnection.IntegrationConnectorsConnection",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IntegrationConnectorsConnection_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIntegrationConnectorsConnection_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.integrationConnectorsConnection.IntegrationConnectorsConnection",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IntegrationConnectorsConnection_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIntegrationConnectorsConnection_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.integrationConnectorsConnection.IntegrationConnectorsConnection",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IntegrationConnectorsConnection_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.integrationConnectorsConnection.IntegrationConnectorsConnection",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IntegrationConnectorsConnection) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnection) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnection) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationConnectorsConnection) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationConnectorsConnection) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationConnectorsConnection) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationConnectorsConnection) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationConnectorsConnection) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationConnectorsConnection) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationConnectorsConnection) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationConnectorsConnection) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationConnectorsConnection) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationConnectorsConnection) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnection) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationConnectorsConnection) MoveFromId(id *string) {
	if err := i.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveFromId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnection) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnection) MoveToId(id *string) {
	if err := i.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveToId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnection) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnection) PutAuthConfig(value *IntegrationConnectorsConnectionAuthConfig) {
	if err := i.validatePutAuthConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAuthConfig",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnection) PutConfigVariable(value interface{}) {
	if err := i.validatePutConfigVariableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putConfigVariable",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnection) PutDestinationConfig(value interface{}) {
	if err := i.validatePutDestinationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putDestinationConfig",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnection) PutEventingConfig(value *IntegrationConnectorsConnectionEventingConfig) {
	if err := i.validatePutEventingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putEventingConfig",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnection) PutLockConfig(value *IntegrationConnectorsConnectionLockConfig) {
	if err := i.validatePutLockConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putLockConfig",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnection) PutLogConfig(value *IntegrationConnectorsConnectionLogConfig) {
	if err := i.validatePutLogConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putLogConfig",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnection) PutNodeConfig(value *IntegrationConnectorsConnectionNodeConfig) {
	if err := i.validatePutNodeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putNodeConfig",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnection) PutSslConfig(value *IntegrationConnectorsConnectionSslConfig) {
	if err := i.validatePutSslConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSslConfig",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnection) PutTimeouts(value *IntegrationConnectorsConnectionTimeouts) {
	if err := i.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnection) ResetAuthConfig() {
	_jsii_.InvokeVoid(
		i,
		"resetAuthConfig",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnection) ResetConfigVariable() {
	_jsii_.InvokeVoid(
		i,
		"resetConfigVariable",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnection) ResetDescription() {
	_jsii_.InvokeVoid(
		i,
		"resetDescription",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnection) ResetDestinationConfig() {
	_jsii_.InvokeVoid(
		i,
		"resetDestinationConfig",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnection) ResetEventingConfig() {
	_jsii_.InvokeVoid(
		i,
		"resetEventingConfig",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnection) ResetEventingEnablementType() {
	_jsii_.InvokeVoid(
		i,
		"resetEventingEnablementType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnection) ResetId() {
	_jsii_.InvokeVoid(
		i,
		"resetId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnection) ResetLabels() {
	_jsii_.InvokeVoid(
		i,
		"resetLabels",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnection) ResetLockConfig() {
	_jsii_.InvokeVoid(
		i,
		"resetLockConfig",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnection) ResetLogConfig() {
	_jsii_.InvokeVoid(
		i,
		"resetLogConfig",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnection) ResetNodeConfig() {
	_jsii_.InvokeVoid(
		i,
		"resetNodeConfig",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnection) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnection) ResetProject() {
	_jsii_.InvokeVoid(
		i,
		"resetProject",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnection) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		i,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnection) ResetSslConfig() {
	_jsii_.InvokeVoid(
		i,
		"resetSslConfig",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnection) ResetSuspended() {
	_jsii_.InvokeVoid(
		i,
		"resetSuspended",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnection) ResetTimeouts() {
	_jsii_.InvokeVoid(
		i,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationConnectorsConnection) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationConnectorsConnection) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationConnectorsConnection) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationConnectorsConnection) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationConnectorsConnection) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationConnectorsConnection) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

