// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonprembaremetaladmincluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v14/gkeonprembaremetaladmincluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/gkeonprem_bare_metal_admin_cluster google_gkeonprem_bare_metal_admin_cluster}.
type GkeonpremBareMetalAdminCluster interface {
	cdktf.TerraformResource
	Annotations() *map[string]*string
	SetAnnotations(val *map[string]*string)
	AnnotationsInput() *map[string]*string
	BareMetalVersion() *string
	SetBareMetalVersion(val *string)
	BareMetalVersionInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterOperations() GkeonpremBareMetalAdminClusterClusterOperationsOutputReference
	ClusterOperationsInput() *GkeonpremBareMetalAdminClusterClusterOperations
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ControlPlane() GkeonpremBareMetalAdminClusterControlPlaneOutputReference
	ControlPlaneInput() *GkeonpremBareMetalAdminClusterControlPlane
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
	Fleet() GkeonpremBareMetalAdminClusterFleetList
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
	LoadBalancer() GkeonpremBareMetalAdminClusterLoadBalancerOutputReference
	LoadBalancerInput() *GkeonpremBareMetalAdminClusterLoadBalancer
	LocalName() *string
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MaintenanceConfig() GkeonpremBareMetalAdminClusterMaintenanceConfigOutputReference
	MaintenanceConfigInput() *GkeonpremBareMetalAdminClusterMaintenanceConfig
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkConfig() GkeonpremBareMetalAdminClusterNetworkConfigOutputReference
	NetworkConfigInput() *GkeonpremBareMetalAdminClusterNetworkConfig
	// The tree node.
	Node() constructs.Node
	NodeAccessConfig() GkeonpremBareMetalAdminClusterNodeAccessConfigOutputReference
	NodeAccessConfigInput() *GkeonpremBareMetalAdminClusterNodeAccessConfig
	NodeConfig() GkeonpremBareMetalAdminClusterNodeConfigOutputReference
	NodeConfigInput() *GkeonpremBareMetalAdminClusterNodeConfig
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
	Proxy() GkeonpremBareMetalAdminClusterProxyOutputReference
	ProxyInput() *GkeonpremBareMetalAdminClusterProxy
	// Experimental.
	RawOverrides() interface{}
	Reconciling() cdktf.IResolvable
	SecurityConfig() GkeonpremBareMetalAdminClusterSecurityConfigOutputReference
	SecurityConfigInput() *GkeonpremBareMetalAdminClusterSecurityConfig
	State() *string
	Status() GkeonpremBareMetalAdminClusterStatusList
	Storage() GkeonpremBareMetalAdminClusterStorageOutputReference
	StorageInput() *GkeonpremBareMetalAdminClusterStorage
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GkeonpremBareMetalAdminClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
	Uid() *string
	UpdateTime() *string
	ValidationCheck() GkeonpremBareMetalAdminClusterValidationCheckList
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
	PutClusterOperations(value *GkeonpremBareMetalAdminClusterClusterOperations)
	PutControlPlane(value *GkeonpremBareMetalAdminClusterControlPlane)
	PutLoadBalancer(value *GkeonpremBareMetalAdminClusterLoadBalancer)
	PutMaintenanceConfig(value *GkeonpremBareMetalAdminClusterMaintenanceConfig)
	PutNetworkConfig(value *GkeonpremBareMetalAdminClusterNetworkConfig)
	PutNodeAccessConfig(value *GkeonpremBareMetalAdminClusterNodeAccessConfig)
	PutNodeConfig(value *GkeonpremBareMetalAdminClusterNodeConfig)
	PutProxy(value *GkeonpremBareMetalAdminClusterProxy)
	PutSecurityConfig(value *GkeonpremBareMetalAdminClusterSecurityConfig)
	PutStorage(value *GkeonpremBareMetalAdminClusterStorage)
	PutTimeouts(value *GkeonpremBareMetalAdminClusterTimeouts)
	ResetAnnotations()
	ResetBareMetalVersion()
	ResetClusterOperations()
	ResetControlPlane()
	ResetDescription()
	ResetId()
	ResetLoadBalancer()
	ResetMaintenanceConfig()
	ResetNetworkConfig()
	ResetNodeAccessConfig()
	ResetNodeConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetProxy()
	ResetSecurityConfig()
	ResetStorage()
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

// The jsii proxy struct for GkeonpremBareMetalAdminCluster
type jsiiProxy_GkeonpremBareMetalAdminCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) Annotations() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) AnnotationsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) BareMetalVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bareMetalVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) BareMetalVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bareMetalVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) ClusterOperations() GkeonpremBareMetalAdminClusterClusterOperationsOutputReference {
	var returns GkeonpremBareMetalAdminClusterClusterOperationsOutputReference
	_jsii_.Get(
		j,
		"clusterOperations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) ClusterOperationsInput() *GkeonpremBareMetalAdminClusterClusterOperations {
	var returns *GkeonpremBareMetalAdminClusterClusterOperations
	_jsii_.Get(
		j,
		"clusterOperationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) ControlPlane() GkeonpremBareMetalAdminClusterControlPlaneOutputReference {
	var returns GkeonpremBareMetalAdminClusterControlPlaneOutputReference
	_jsii_.Get(
		j,
		"controlPlane",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) ControlPlaneInput() *GkeonpremBareMetalAdminClusterControlPlane {
	var returns *GkeonpremBareMetalAdminClusterControlPlane
	_jsii_.Get(
		j,
		"controlPlaneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) DeleteTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) EffectiveAnnotations() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) Fleet() GkeonpremBareMetalAdminClusterFleetList {
	var returns GkeonpremBareMetalAdminClusterFleetList
	_jsii_.Get(
		j,
		"fleet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) LoadBalancer() GkeonpremBareMetalAdminClusterLoadBalancerOutputReference {
	var returns GkeonpremBareMetalAdminClusterLoadBalancerOutputReference
	_jsii_.Get(
		j,
		"loadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) LoadBalancerInput() *GkeonpremBareMetalAdminClusterLoadBalancer {
	var returns *GkeonpremBareMetalAdminClusterLoadBalancer
	_jsii_.Get(
		j,
		"loadBalancerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) LocalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) MaintenanceConfig() GkeonpremBareMetalAdminClusterMaintenanceConfigOutputReference {
	var returns GkeonpremBareMetalAdminClusterMaintenanceConfigOutputReference
	_jsii_.Get(
		j,
		"maintenanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) MaintenanceConfigInput() *GkeonpremBareMetalAdminClusterMaintenanceConfig {
	var returns *GkeonpremBareMetalAdminClusterMaintenanceConfig
	_jsii_.Get(
		j,
		"maintenanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) NetworkConfig() GkeonpremBareMetalAdminClusterNetworkConfigOutputReference {
	var returns GkeonpremBareMetalAdminClusterNetworkConfigOutputReference
	_jsii_.Get(
		j,
		"networkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) NetworkConfigInput() *GkeonpremBareMetalAdminClusterNetworkConfig {
	var returns *GkeonpremBareMetalAdminClusterNetworkConfig
	_jsii_.Get(
		j,
		"networkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) NodeAccessConfig() GkeonpremBareMetalAdminClusterNodeAccessConfigOutputReference {
	var returns GkeonpremBareMetalAdminClusterNodeAccessConfigOutputReference
	_jsii_.Get(
		j,
		"nodeAccessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) NodeAccessConfigInput() *GkeonpremBareMetalAdminClusterNodeAccessConfig {
	var returns *GkeonpremBareMetalAdminClusterNodeAccessConfig
	_jsii_.Get(
		j,
		"nodeAccessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) NodeConfig() GkeonpremBareMetalAdminClusterNodeConfigOutputReference {
	var returns GkeonpremBareMetalAdminClusterNodeConfigOutputReference
	_jsii_.Get(
		j,
		"nodeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) NodeConfigInput() *GkeonpremBareMetalAdminClusterNodeConfig {
	var returns *GkeonpremBareMetalAdminClusterNodeConfig
	_jsii_.Get(
		j,
		"nodeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) Proxy() GkeonpremBareMetalAdminClusterProxyOutputReference {
	var returns GkeonpremBareMetalAdminClusterProxyOutputReference
	_jsii_.Get(
		j,
		"proxy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) ProxyInput() *GkeonpremBareMetalAdminClusterProxy {
	var returns *GkeonpremBareMetalAdminClusterProxy
	_jsii_.Get(
		j,
		"proxyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) Reconciling() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"reconciling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) SecurityConfig() GkeonpremBareMetalAdminClusterSecurityConfigOutputReference {
	var returns GkeonpremBareMetalAdminClusterSecurityConfigOutputReference
	_jsii_.Get(
		j,
		"securityConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) SecurityConfigInput() *GkeonpremBareMetalAdminClusterSecurityConfig {
	var returns *GkeonpremBareMetalAdminClusterSecurityConfig
	_jsii_.Get(
		j,
		"securityConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) Status() GkeonpremBareMetalAdminClusterStatusList {
	var returns GkeonpremBareMetalAdminClusterStatusList
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) Storage() GkeonpremBareMetalAdminClusterStorageOutputReference {
	var returns GkeonpremBareMetalAdminClusterStorageOutputReference
	_jsii_.Get(
		j,
		"storage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) StorageInput() *GkeonpremBareMetalAdminClusterStorage {
	var returns *GkeonpremBareMetalAdminClusterStorage
	_jsii_.Get(
		j,
		"storageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) Timeouts() GkeonpremBareMetalAdminClusterTimeoutsOutputReference {
	var returns GkeonpremBareMetalAdminClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) Uid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster) ValidationCheck() GkeonpremBareMetalAdminClusterValidationCheckList {
	var returns GkeonpremBareMetalAdminClusterValidationCheckList
	_jsii_.Get(
		j,
		"validationCheck",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/gkeonprem_bare_metal_admin_cluster google_gkeonprem_bare_metal_admin_cluster} Resource.
func NewGkeonpremBareMetalAdminCluster(scope constructs.Construct, id *string, config *GkeonpremBareMetalAdminClusterConfig) GkeonpremBareMetalAdminCluster {
	_init_.Initialize()

	if err := validateNewGkeonpremBareMetalAdminClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GkeonpremBareMetalAdminCluster{}

	_jsii_.Create(
		"@cdktf/provider-google.gkeonpremBareMetalAdminCluster.GkeonpremBareMetalAdminCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/gkeonprem_bare_metal_admin_cluster google_gkeonprem_bare_metal_admin_cluster} Resource.
func NewGkeonpremBareMetalAdminCluster_Override(g GkeonpremBareMetalAdminCluster, scope constructs.Construct, id *string, config *GkeonpremBareMetalAdminClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.gkeonpremBareMetalAdminCluster.GkeonpremBareMetalAdminCluster",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster)SetAnnotations(val *map[string]*string) {
	if err := j.validateSetAnnotationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotations",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster)SetBareMetalVersion(val *string) {
	if err := j.validateSetBareMetalVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bareMetalVersion",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GkeonpremBareMetalAdminCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a GkeonpremBareMetalAdminCluster resource upon running "cdktf plan <stack-name>".
func GkeonpremBareMetalAdminCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGkeonpremBareMetalAdminCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.gkeonpremBareMetalAdminCluster.GkeonpremBareMetalAdminCluster",
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
func GkeonpremBareMetalAdminCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGkeonpremBareMetalAdminCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.gkeonpremBareMetalAdminCluster.GkeonpremBareMetalAdminCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GkeonpremBareMetalAdminCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGkeonpremBareMetalAdminCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.gkeonpremBareMetalAdminCluster.GkeonpremBareMetalAdminCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GkeonpremBareMetalAdminCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGkeonpremBareMetalAdminCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.gkeonpremBareMetalAdminCluster.GkeonpremBareMetalAdminCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GkeonpremBareMetalAdminCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.gkeonpremBareMetalAdminCluster.GkeonpremBareMetalAdminCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) PutClusterOperations(value *GkeonpremBareMetalAdminClusterClusterOperations) {
	if err := g.validatePutClusterOperationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putClusterOperations",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) PutControlPlane(value *GkeonpremBareMetalAdminClusterControlPlane) {
	if err := g.validatePutControlPlaneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putControlPlane",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) PutLoadBalancer(value *GkeonpremBareMetalAdminClusterLoadBalancer) {
	if err := g.validatePutLoadBalancerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLoadBalancer",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) PutMaintenanceConfig(value *GkeonpremBareMetalAdminClusterMaintenanceConfig) {
	if err := g.validatePutMaintenanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMaintenanceConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) PutNetworkConfig(value *GkeonpremBareMetalAdminClusterNetworkConfig) {
	if err := g.validatePutNetworkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNetworkConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) PutNodeAccessConfig(value *GkeonpremBareMetalAdminClusterNodeAccessConfig) {
	if err := g.validatePutNodeAccessConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNodeAccessConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) PutNodeConfig(value *GkeonpremBareMetalAdminClusterNodeConfig) {
	if err := g.validatePutNodeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNodeConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) PutProxy(value *GkeonpremBareMetalAdminClusterProxy) {
	if err := g.validatePutProxyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putProxy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) PutSecurityConfig(value *GkeonpremBareMetalAdminClusterSecurityConfig) {
	if err := g.validatePutSecurityConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSecurityConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) PutStorage(value *GkeonpremBareMetalAdminClusterStorage) {
	if err := g.validatePutStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStorage",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) PutTimeouts(value *GkeonpremBareMetalAdminClusterTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) ResetAnnotations() {
	_jsii_.InvokeVoid(
		g,
		"resetAnnotations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) ResetBareMetalVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetBareMetalVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) ResetClusterOperations() {
	_jsii_.InvokeVoid(
		g,
		"resetClusterOperations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) ResetControlPlane() {
	_jsii_.InvokeVoid(
		g,
		"resetControlPlane",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) ResetLoadBalancer() {
	_jsii_.InvokeVoid(
		g,
		"resetLoadBalancer",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) ResetMaintenanceConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetMaintenanceConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) ResetNetworkConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) ResetNodeAccessConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetNodeAccessConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) ResetNodeConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetNodeConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) ResetProxy() {
	_jsii_.InvokeVoid(
		g,
		"resetProxy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) ResetSecurityConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetSecurityConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) ResetStorage() {
	_jsii_.InvokeVoid(
		g,
		"resetStorage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GkeonpremBareMetalAdminCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

