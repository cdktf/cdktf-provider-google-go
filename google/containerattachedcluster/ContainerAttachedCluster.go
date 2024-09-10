// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerattachedcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v14/containerattachedcluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.2.0/docs/resources/container_attached_cluster google_container_attached_cluster}.
type ContainerAttachedCluster interface {
	cdktf.TerraformResource
	Annotations() *map[string]*string
	SetAnnotations(val *map[string]*string)
	AnnotationsInput() *map[string]*string
	Authorization() ContainerAttachedClusterAuthorizationOutputReference
	AuthorizationInput() *ContainerAttachedClusterAuthorization
	BinaryAuthorization() ContainerAttachedClusterBinaryAuthorizationOutputReference
	BinaryAuthorizationInput() *ContainerAttachedClusterBinaryAuthorization
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterRegion() *string
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
	DeletionPolicy() *string
	SetDeletionPolicy(val *string)
	DeletionPolicyInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Distribution() *string
	SetDistribution(val *string)
	DistributionInput() *string
	EffectiveAnnotations() cdktf.StringMap
	Errors() ContainerAttachedClusterErrorsList
	Fleet() ContainerAttachedClusterFleetOutputReference
	FleetInput() *ContainerAttachedClusterFleet
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
	KubernetesVersion() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	LoggingConfig() ContainerAttachedClusterLoggingConfigOutputReference
	LoggingConfigInput() *ContainerAttachedClusterLoggingConfig
	MonitoringConfig() ContainerAttachedClusterMonitoringConfigOutputReference
	MonitoringConfigInput() *ContainerAttachedClusterMonitoringConfig
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OidcConfig() ContainerAttachedClusterOidcConfigOutputReference
	OidcConfigInput() *ContainerAttachedClusterOidcConfig
	PlatformVersion() *string
	SetPlatformVersion(val *string)
	PlatformVersionInput() *string
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
	ProxyConfig() ContainerAttachedClusterProxyConfigOutputReference
	ProxyConfigInput() *ContainerAttachedClusterProxyConfig
	// Experimental.
	RawOverrides() interface{}
	Reconciling() cdktf.IResolvable
	State() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ContainerAttachedClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
	Uid() *string
	UpdateTime() *string
	WorkloadIdentityConfig() ContainerAttachedClusterWorkloadIdentityConfigList
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
	PutAuthorization(value *ContainerAttachedClusterAuthorization)
	PutBinaryAuthorization(value *ContainerAttachedClusterBinaryAuthorization)
	PutFleet(value *ContainerAttachedClusterFleet)
	PutLoggingConfig(value *ContainerAttachedClusterLoggingConfig)
	PutMonitoringConfig(value *ContainerAttachedClusterMonitoringConfig)
	PutOidcConfig(value *ContainerAttachedClusterOidcConfig)
	PutProxyConfig(value *ContainerAttachedClusterProxyConfig)
	PutTimeouts(value *ContainerAttachedClusterTimeouts)
	ResetAnnotations()
	ResetAuthorization()
	ResetBinaryAuthorization()
	ResetDeletionPolicy()
	ResetDescription()
	ResetId()
	ResetLoggingConfig()
	ResetMonitoringConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetProxyConfig()
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

// The jsii proxy struct for ContainerAttachedCluster
type jsiiProxy_ContainerAttachedCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ContainerAttachedCluster) Annotations() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) AnnotationsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) Authorization() ContainerAttachedClusterAuthorizationOutputReference {
	var returns ContainerAttachedClusterAuthorizationOutputReference
	_jsii_.Get(
		j,
		"authorization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) AuthorizationInput() *ContainerAttachedClusterAuthorization {
	var returns *ContainerAttachedClusterAuthorization
	_jsii_.Get(
		j,
		"authorizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) BinaryAuthorization() ContainerAttachedClusterBinaryAuthorizationOutputReference {
	var returns ContainerAttachedClusterBinaryAuthorizationOutputReference
	_jsii_.Get(
		j,
		"binaryAuthorization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) BinaryAuthorizationInput() *ContainerAttachedClusterBinaryAuthorization {
	var returns *ContainerAttachedClusterBinaryAuthorization
	_jsii_.Get(
		j,
		"binaryAuthorizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) ClusterRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) DeletionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) DeletionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) Distribution() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distribution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) DistributionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) EffectiveAnnotations() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) Errors() ContainerAttachedClusterErrorsList {
	var returns ContainerAttachedClusterErrorsList
	_jsii_.Get(
		j,
		"errors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) Fleet() ContainerAttachedClusterFleetOutputReference {
	var returns ContainerAttachedClusterFleetOutputReference
	_jsii_.Get(
		j,
		"fleet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) FleetInput() *ContainerAttachedClusterFleet {
	var returns *ContainerAttachedClusterFleet
	_jsii_.Get(
		j,
		"fleetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) KubernetesVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) LoggingConfig() ContainerAttachedClusterLoggingConfigOutputReference {
	var returns ContainerAttachedClusterLoggingConfigOutputReference
	_jsii_.Get(
		j,
		"loggingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) LoggingConfigInput() *ContainerAttachedClusterLoggingConfig {
	var returns *ContainerAttachedClusterLoggingConfig
	_jsii_.Get(
		j,
		"loggingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) MonitoringConfig() ContainerAttachedClusterMonitoringConfigOutputReference {
	var returns ContainerAttachedClusterMonitoringConfigOutputReference
	_jsii_.Get(
		j,
		"monitoringConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) MonitoringConfigInput() *ContainerAttachedClusterMonitoringConfig {
	var returns *ContainerAttachedClusterMonitoringConfig
	_jsii_.Get(
		j,
		"monitoringConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) OidcConfig() ContainerAttachedClusterOidcConfigOutputReference {
	var returns ContainerAttachedClusterOidcConfigOutputReference
	_jsii_.Get(
		j,
		"oidcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) OidcConfigInput() *ContainerAttachedClusterOidcConfig {
	var returns *ContainerAttachedClusterOidcConfig
	_jsii_.Get(
		j,
		"oidcConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) PlatformVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) PlatformVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) ProxyConfig() ContainerAttachedClusterProxyConfigOutputReference {
	var returns ContainerAttachedClusterProxyConfigOutputReference
	_jsii_.Get(
		j,
		"proxyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) ProxyConfigInput() *ContainerAttachedClusterProxyConfig {
	var returns *ContainerAttachedClusterProxyConfig
	_jsii_.Get(
		j,
		"proxyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) Reconciling() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"reconciling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) Timeouts() ContainerAttachedClusterTimeoutsOutputReference {
	var returns ContainerAttachedClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) Uid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerAttachedCluster) WorkloadIdentityConfig() ContainerAttachedClusterWorkloadIdentityConfigList {
	var returns ContainerAttachedClusterWorkloadIdentityConfigList
	_jsii_.Get(
		j,
		"workloadIdentityConfig",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.2.0/docs/resources/container_attached_cluster google_container_attached_cluster} Resource.
func NewContainerAttachedCluster(scope constructs.Construct, id *string, config *ContainerAttachedClusterConfig) ContainerAttachedCluster {
	_init_.Initialize()

	if err := validateNewContainerAttachedClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerAttachedCluster{}

	_jsii_.Create(
		"@cdktf/provider-google.containerAttachedCluster.ContainerAttachedCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.2.0/docs/resources/container_attached_cluster google_container_attached_cluster} Resource.
func NewContainerAttachedCluster_Override(c ContainerAttachedCluster, scope constructs.Construct, id *string, config *ContainerAttachedClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.containerAttachedCluster.ContainerAttachedCluster",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ContainerAttachedCluster)SetAnnotations(val *map[string]*string) {
	if err := j.validateSetAnnotationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotations",
		val,
	)
}

func (j *jsiiProxy_ContainerAttachedCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ContainerAttachedCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ContainerAttachedCluster)SetDeletionPolicy(val *string) {
	if err := j.validateSetDeletionPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionPolicy",
		val,
	)
}

func (j *jsiiProxy_ContainerAttachedCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ContainerAttachedCluster)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ContainerAttachedCluster)SetDistribution(val *string) {
	if err := j.validateSetDistributionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"distribution",
		val,
	)
}

func (j *jsiiProxy_ContainerAttachedCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ContainerAttachedCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ContainerAttachedCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ContainerAttachedCluster)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_ContainerAttachedCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ContainerAttachedCluster)SetPlatformVersion(val *string) {
	if err := j.validateSetPlatformVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"platformVersion",
		val,
	)
}

func (j *jsiiProxy_ContainerAttachedCluster)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_ContainerAttachedCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ContainerAttachedCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a ContainerAttachedCluster resource upon running "cdktf plan <stack-name>".
func ContainerAttachedCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateContainerAttachedCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.containerAttachedCluster.ContainerAttachedCluster",
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
func ContainerAttachedCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateContainerAttachedCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.containerAttachedCluster.ContainerAttachedCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ContainerAttachedCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateContainerAttachedCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.containerAttachedCluster.ContainerAttachedCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ContainerAttachedCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateContainerAttachedCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.containerAttachedCluster.ContainerAttachedCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ContainerAttachedCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.containerAttachedCluster.ContainerAttachedCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ContainerAttachedCluster) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_ContainerAttachedCluster) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ContainerAttachedCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ContainerAttachedCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerAttachedCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ContainerAttachedCluster) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ContainerAttachedCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ContainerAttachedCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ContainerAttachedCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ContainerAttachedCluster) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ContainerAttachedCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ContainerAttachedCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAttachedCluster) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_ContainerAttachedCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerAttachedCluster) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ContainerAttachedCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_ContainerAttachedCluster) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ContainerAttachedCluster) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ContainerAttachedCluster) PutAuthorization(value *ContainerAttachedClusterAuthorization) {
	if err := c.validatePutAuthorizationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAuthorization",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAttachedCluster) PutBinaryAuthorization(value *ContainerAttachedClusterBinaryAuthorization) {
	if err := c.validatePutBinaryAuthorizationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBinaryAuthorization",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAttachedCluster) PutFleet(value *ContainerAttachedClusterFleet) {
	if err := c.validatePutFleetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putFleet",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAttachedCluster) PutLoggingConfig(value *ContainerAttachedClusterLoggingConfig) {
	if err := c.validatePutLoggingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLoggingConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAttachedCluster) PutMonitoringConfig(value *ContainerAttachedClusterMonitoringConfig) {
	if err := c.validatePutMonitoringConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMonitoringConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAttachedCluster) PutOidcConfig(value *ContainerAttachedClusterOidcConfig) {
	if err := c.validatePutOidcConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putOidcConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAttachedCluster) PutProxyConfig(value *ContainerAttachedClusterProxyConfig) {
	if err := c.validatePutProxyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putProxyConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAttachedCluster) PutTimeouts(value *ContainerAttachedClusterTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerAttachedCluster) ResetAnnotations() {
	_jsii_.InvokeVoid(
		c,
		"resetAnnotations",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAttachedCluster) ResetAuthorization() {
	_jsii_.InvokeVoid(
		c,
		"resetAuthorization",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAttachedCluster) ResetBinaryAuthorization() {
	_jsii_.InvokeVoid(
		c,
		"resetBinaryAuthorization",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAttachedCluster) ResetDeletionPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetDeletionPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAttachedCluster) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAttachedCluster) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAttachedCluster) ResetLoggingConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetLoggingConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAttachedCluster) ResetMonitoringConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetMonitoringConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAttachedCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAttachedCluster) ResetProject() {
	_jsii_.InvokeVoid(
		c,
		"resetProject",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAttachedCluster) ResetProxyConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetProxyConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAttachedCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerAttachedCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAttachedCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAttachedCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAttachedCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAttachedCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerAttachedCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

