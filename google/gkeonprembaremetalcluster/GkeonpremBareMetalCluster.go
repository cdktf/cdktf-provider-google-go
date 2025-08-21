// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonprembaremetalcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v16/gkeonprembaremetalcluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/gkeonprem_bare_metal_cluster google_gkeonprem_bare_metal_cluster}.
type GkeonpremBareMetalCluster interface {
	cdktf.TerraformResource
	AdminClusterMembership() *string
	SetAdminClusterMembership(val *string)
	AdminClusterMembershipInput() *string
	Annotations() *map[string]*string
	SetAnnotations(val *map[string]*string)
	AnnotationsInput() *map[string]*string
	BareMetalVersion() *string
	SetBareMetalVersion(val *string)
	BareMetalVersionInput() *string
	BinaryAuthorization() GkeonpremBareMetalClusterBinaryAuthorizationOutputReference
	BinaryAuthorizationInput() *GkeonpremBareMetalClusterBinaryAuthorization
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterOperations() GkeonpremBareMetalClusterClusterOperationsOutputReference
	ClusterOperationsInput() *GkeonpremBareMetalClusterClusterOperations
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ControlPlane() GkeonpremBareMetalClusterControlPlaneOutputReference
	ControlPlaneInput() *GkeonpremBareMetalClusterControlPlane
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreateTime() *string
	DeleteTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EffectiveAnnotations() cdktf.StringMap
	Endpoint() *string
	Etag() *string
	Fleet() GkeonpremBareMetalClusterFleetList
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoadBalancer() GkeonpremBareMetalClusterLoadBalancerOutputReference
	LoadBalancerInput() *GkeonpremBareMetalClusterLoadBalancer
	LocalName() *string
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MaintenanceConfig() GkeonpremBareMetalClusterMaintenanceConfigOutputReference
	MaintenanceConfigInput() *GkeonpremBareMetalClusterMaintenanceConfig
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkConfig() GkeonpremBareMetalClusterNetworkConfigOutputReference
	NetworkConfigInput() *GkeonpremBareMetalClusterNetworkConfig
	// The tree node.
	Node() constructs.Node
	NodeAccessConfig() GkeonpremBareMetalClusterNodeAccessConfigOutputReference
	NodeAccessConfigInput() *GkeonpremBareMetalClusterNodeAccessConfig
	NodeConfig() GkeonpremBareMetalClusterNodeConfigOutputReference
	NodeConfigInput() *GkeonpremBareMetalClusterNodeConfig
	OsEnvironmentConfig() GkeonpremBareMetalClusterOsEnvironmentConfigOutputReference
	OsEnvironmentConfigInput() *GkeonpremBareMetalClusterOsEnvironmentConfig
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
	Proxy() GkeonpremBareMetalClusterProxyOutputReference
	ProxyInput() *GkeonpremBareMetalClusterProxy
	// Experimental.
	RawOverrides() interface{}
	Reconciling() cdktf.IResolvable
	SecurityConfig() GkeonpremBareMetalClusterSecurityConfigOutputReference
	SecurityConfigInput() *GkeonpremBareMetalClusterSecurityConfig
	State() *string
	Status() GkeonpremBareMetalClusterStatusList
	Storage() GkeonpremBareMetalClusterStorageOutputReference
	StorageInput() *GkeonpremBareMetalClusterStorage
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GkeonpremBareMetalClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
	Uid() *string
	UpdateTime() *string
	UpgradePolicy() GkeonpremBareMetalClusterUpgradePolicyOutputReference
	UpgradePolicyInput() *GkeonpremBareMetalClusterUpgradePolicy
	ValidationCheck() GkeonpremBareMetalClusterValidationCheckList
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
	PutBinaryAuthorization(value *GkeonpremBareMetalClusterBinaryAuthorization)
	PutClusterOperations(value *GkeonpremBareMetalClusterClusterOperations)
	PutControlPlane(value *GkeonpremBareMetalClusterControlPlane)
	PutLoadBalancer(value *GkeonpremBareMetalClusterLoadBalancer)
	PutMaintenanceConfig(value *GkeonpremBareMetalClusterMaintenanceConfig)
	PutNetworkConfig(value *GkeonpremBareMetalClusterNetworkConfig)
	PutNodeAccessConfig(value *GkeonpremBareMetalClusterNodeAccessConfig)
	PutNodeConfig(value *GkeonpremBareMetalClusterNodeConfig)
	PutOsEnvironmentConfig(value *GkeonpremBareMetalClusterOsEnvironmentConfig)
	PutProxy(value *GkeonpremBareMetalClusterProxy)
	PutSecurityConfig(value *GkeonpremBareMetalClusterSecurityConfig)
	PutStorage(value *GkeonpremBareMetalClusterStorage)
	PutTimeouts(value *GkeonpremBareMetalClusterTimeouts)
	PutUpgradePolicy(value *GkeonpremBareMetalClusterUpgradePolicy)
	ResetAnnotations()
	ResetBinaryAuthorization()
	ResetClusterOperations()
	ResetDescription()
	ResetId()
	ResetMaintenanceConfig()
	ResetNodeAccessConfig()
	ResetNodeConfig()
	ResetOsEnvironmentConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetProxy()
	ResetSecurityConfig()
	ResetTimeouts()
	ResetUpgradePolicy()
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

// The jsii proxy struct for GkeonpremBareMetalCluster
type jsiiProxy_GkeonpremBareMetalCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) AdminClusterMembership() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminClusterMembership",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) AdminClusterMembershipInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminClusterMembershipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) Annotations() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) AnnotationsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) BareMetalVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bareMetalVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) BareMetalVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bareMetalVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) BinaryAuthorization() GkeonpremBareMetalClusterBinaryAuthorizationOutputReference {
	var returns GkeonpremBareMetalClusterBinaryAuthorizationOutputReference
	_jsii_.Get(
		j,
		"binaryAuthorization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) BinaryAuthorizationInput() *GkeonpremBareMetalClusterBinaryAuthorization {
	var returns *GkeonpremBareMetalClusterBinaryAuthorization
	_jsii_.Get(
		j,
		"binaryAuthorizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) ClusterOperations() GkeonpremBareMetalClusterClusterOperationsOutputReference {
	var returns GkeonpremBareMetalClusterClusterOperationsOutputReference
	_jsii_.Get(
		j,
		"clusterOperations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) ClusterOperationsInput() *GkeonpremBareMetalClusterClusterOperations {
	var returns *GkeonpremBareMetalClusterClusterOperations
	_jsii_.Get(
		j,
		"clusterOperationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) ControlPlane() GkeonpremBareMetalClusterControlPlaneOutputReference {
	var returns GkeonpremBareMetalClusterControlPlaneOutputReference
	_jsii_.Get(
		j,
		"controlPlane",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) ControlPlaneInput() *GkeonpremBareMetalClusterControlPlane {
	var returns *GkeonpremBareMetalClusterControlPlane
	_jsii_.Get(
		j,
		"controlPlaneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) DeleteTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) EffectiveAnnotations() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) Fleet() GkeonpremBareMetalClusterFleetList {
	var returns GkeonpremBareMetalClusterFleetList
	_jsii_.Get(
		j,
		"fleet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) LoadBalancer() GkeonpremBareMetalClusterLoadBalancerOutputReference {
	var returns GkeonpremBareMetalClusterLoadBalancerOutputReference
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) LoadBalancerInput() *GkeonpremBareMetalClusterLoadBalancer {
	var returns *GkeonpremBareMetalClusterLoadBalancer
	_jsii_.Get(
		j,
		"loadBalancerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) LocalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) MaintenanceConfig() GkeonpremBareMetalClusterMaintenanceConfigOutputReference {
	var returns GkeonpremBareMetalClusterMaintenanceConfigOutputReference
	_jsii_.Get(
		j,
		"maintenanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) MaintenanceConfigInput() *GkeonpremBareMetalClusterMaintenanceConfig {
	var returns *GkeonpremBareMetalClusterMaintenanceConfig
	_jsii_.Get(
		j,
		"maintenanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) NetworkConfig() GkeonpremBareMetalClusterNetworkConfigOutputReference {
	var returns GkeonpremBareMetalClusterNetworkConfigOutputReference
	_jsii_.Get(
		j,
		"networkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) NetworkConfigInput() *GkeonpremBareMetalClusterNetworkConfig {
	var returns *GkeonpremBareMetalClusterNetworkConfig
	_jsii_.Get(
		j,
		"networkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) NodeAccessConfig() GkeonpremBareMetalClusterNodeAccessConfigOutputReference {
	var returns GkeonpremBareMetalClusterNodeAccessConfigOutputReference
	_jsii_.Get(
		j,
		"nodeAccessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) NodeAccessConfigInput() *GkeonpremBareMetalClusterNodeAccessConfig {
	var returns *GkeonpremBareMetalClusterNodeAccessConfig
	_jsii_.Get(
		j,
		"nodeAccessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) NodeConfig() GkeonpremBareMetalClusterNodeConfigOutputReference {
	var returns GkeonpremBareMetalClusterNodeConfigOutputReference
	_jsii_.Get(
		j,
		"nodeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) NodeConfigInput() *GkeonpremBareMetalClusterNodeConfig {
	var returns *GkeonpremBareMetalClusterNodeConfig
	_jsii_.Get(
		j,
		"nodeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) OsEnvironmentConfig() GkeonpremBareMetalClusterOsEnvironmentConfigOutputReference {
	var returns GkeonpremBareMetalClusterOsEnvironmentConfigOutputReference
	_jsii_.Get(
		j,
		"osEnvironmentConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) OsEnvironmentConfigInput() *GkeonpremBareMetalClusterOsEnvironmentConfig {
	var returns *GkeonpremBareMetalClusterOsEnvironmentConfig
	_jsii_.Get(
		j,
		"osEnvironmentConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) Proxy() GkeonpremBareMetalClusterProxyOutputReference {
	var returns GkeonpremBareMetalClusterProxyOutputReference
	_jsii_.Get(
		j,
		"proxy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) ProxyInput() *GkeonpremBareMetalClusterProxy {
	var returns *GkeonpremBareMetalClusterProxy
	_jsii_.Get(
		j,
		"proxyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) Reconciling() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"reconciling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) SecurityConfig() GkeonpremBareMetalClusterSecurityConfigOutputReference {
	var returns GkeonpremBareMetalClusterSecurityConfigOutputReference
	_jsii_.Get(
		j,
		"securityConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) SecurityConfigInput() *GkeonpremBareMetalClusterSecurityConfig {
	var returns *GkeonpremBareMetalClusterSecurityConfig
	_jsii_.Get(
		j,
		"securityConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) Status() GkeonpremBareMetalClusterStatusList {
	var returns GkeonpremBareMetalClusterStatusList
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) Storage() GkeonpremBareMetalClusterStorageOutputReference {
	var returns GkeonpremBareMetalClusterStorageOutputReference
	_jsii_.Get(
		j,
		"storage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) StorageInput() *GkeonpremBareMetalClusterStorage {
	var returns *GkeonpremBareMetalClusterStorage
	_jsii_.Get(
		j,
		"storageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) Timeouts() GkeonpremBareMetalClusterTimeoutsOutputReference {
	var returns GkeonpremBareMetalClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) Uid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) UpgradePolicy() GkeonpremBareMetalClusterUpgradePolicyOutputReference {
	var returns GkeonpremBareMetalClusterUpgradePolicyOutputReference
	_jsii_.Get(
		j,
		"upgradePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) UpgradePolicyInput() *GkeonpremBareMetalClusterUpgradePolicy {
	var returns *GkeonpremBareMetalClusterUpgradePolicy
	_jsii_.Get(
		j,
		"upgradePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalCluster) ValidationCheck() GkeonpremBareMetalClusterValidationCheckList {
	var returns GkeonpremBareMetalClusterValidationCheckList
	_jsii_.Get(
		j,
		"validationCheck",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/gkeonprem_bare_metal_cluster google_gkeonprem_bare_metal_cluster} Resource.
func NewGkeonpremBareMetalCluster(scope constructs.Construct, id *string, config *GkeonpremBareMetalClusterConfig) GkeonpremBareMetalCluster {
	_init_.Initialize()

	if err := validateNewGkeonpremBareMetalClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GkeonpremBareMetalCluster{}

	_jsii_.Create(
		"@cdktf/provider-google.gkeonpremBareMetalCluster.GkeonpremBareMetalCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/gkeonprem_bare_metal_cluster google_gkeonprem_bare_metal_cluster} Resource.
func NewGkeonpremBareMetalCluster_Override(g GkeonpremBareMetalCluster, scope constructs.Construct, id *string, config *GkeonpremBareMetalClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.gkeonpremBareMetalCluster.GkeonpremBareMetalCluster",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalCluster)SetAdminClusterMembership(val *string) {
	if err := j.validateSetAdminClusterMembershipParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminClusterMembership",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalCluster)SetAnnotations(val *map[string]*string) {
	if err := j.validateSetAnnotationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotations",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalCluster)SetBareMetalVersion(val *string) {
	if err := j.validateSetBareMetalVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bareMetalVersion",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalCluster)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalCluster)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalCluster)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a GkeonpremBareMetalCluster resource upon running "cdktf plan <stack-name>".
func GkeonpremBareMetalCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGkeonpremBareMetalCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.gkeonpremBareMetalCluster.GkeonpremBareMetalCluster",
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
func GkeonpremBareMetalCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGkeonpremBareMetalCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.gkeonpremBareMetalCluster.GkeonpremBareMetalCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GkeonpremBareMetalCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGkeonpremBareMetalCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.gkeonpremBareMetalCluster.GkeonpremBareMetalCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GkeonpremBareMetalCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGkeonpremBareMetalCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.gkeonpremBareMetalCluster.GkeonpremBareMetalCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GkeonpremBareMetalCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.gkeonpremBareMetalCluster.GkeonpremBareMetalCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) PutBinaryAuthorization(value *GkeonpremBareMetalClusterBinaryAuthorization) {
	if err := g.validatePutBinaryAuthorizationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBinaryAuthorization",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) PutClusterOperations(value *GkeonpremBareMetalClusterClusterOperations) {
	if err := g.validatePutClusterOperationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putClusterOperations",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) PutControlPlane(value *GkeonpremBareMetalClusterControlPlane) {
	if err := g.validatePutControlPlaneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putControlPlane",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) PutLoadBalancer(value *GkeonpremBareMetalClusterLoadBalancer) {
	if err := g.validatePutLoadBalancerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLoadBalancer",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) PutMaintenanceConfig(value *GkeonpremBareMetalClusterMaintenanceConfig) {
	if err := g.validatePutMaintenanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMaintenanceConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) PutNetworkConfig(value *GkeonpremBareMetalClusterNetworkConfig) {
	if err := g.validatePutNetworkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNetworkConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) PutNodeAccessConfig(value *GkeonpremBareMetalClusterNodeAccessConfig) {
	if err := g.validatePutNodeAccessConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNodeAccessConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) PutNodeConfig(value *GkeonpremBareMetalClusterNodeConfig) {
	if err := g.validatePutNodeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNodeConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) PutOsEnvironmentConfig(value *GkeonpremBareMetalClusterOsEnvironmentConfig) {
	if err := g.validatePutOsEnvironmentConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOsEnvironmentConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) PutProxy(value *GkeonpremBareMetalClusterProxy) {
	if err := g.validatePutProxyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putProxy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) PutSecurityConfig(value *GkeonpremBareMetalClusterSecurityConfig) {
	if err := g.validatePutSecurityConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSecurityConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) PutStorage(value *GkeonpremBareMetalClusterStorage) {
	if err := g.validatePutStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStorage",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) PutTimeouts(value *GkeonpremBareMetalClusterTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) PutUpgradePolicy(value *GkeonpremBareMetalClusterUpgradePolicy) {
	if err := g.validatePutUpgradePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putUpgradePolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) ResetAnnotations() {
	_jsii_.InvokeVoid(
		g,
		"resetAnnotations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) ResetBinaryAuthorization() {
	_jsii_.InvokeVoid(
		g,
		"resetBinaryAuthorization",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) ResetClusterOperations() {
	_jsii_.InvokeVoid(
		g,
		"resetClusterOperations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) ResetMaintenanceConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetMaintenanceConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) ResetNodeAccessConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetNodeAccessConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) ResetNodeConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetNodeConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) ResetOsEnvironmentConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetOsEnvironmentConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) ResetProxy() {
	_jsii_.InvokeVoid(
		g,
		"resetProxy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) ResetSecurityConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetSecurityConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) ResetUpgradePolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetUpgradePolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

