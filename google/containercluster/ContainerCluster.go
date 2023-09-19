// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v9/containercluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/4.83.0/docs/resources/container_cluster google_container_cluster}.
type ContainerCluster interface {
	cdktf.TerraformResource
	AddonsConfig() ContainerClusterAddonsConfigOutputReference
	AddonsConfigInput() *ContainerClusterAddonsConfig
	AllowNetAdmin() interface{}
	SetAllowNetAdmin(val interface{})
	AllowNetAdminInput() interface{}
	AuthenticatorGroupsConfig() ContainerClusterAuthenticatorGroupsConfigOutputReference
	AuthenticatorGroupsConfigInput() *ContainerClusterAuthenticatorGroupsConfig
	BinaryAuthorization() ContainerClusterBinaryAuthorizationOutputReference
	BinaryAuthorizationInput() *ContainerClusterBinaryAuthorization
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterAutoscaling() ContainerClusterClusterAutoscalingOutputReference
	ClusterAutoscalingInput() *ContainerClusterClusterAutoscaling
	ClusterIpv4Cidr() *string
	SetClusterIpv4Cidr(val *string)
	ClusterIpv4CidrInput() *string
	ConfidentialNodes() ContainerClusterConfidentialNodesOutputReference
	ConfidentialNodesInput() *ContainerClusterConfidentialNodes
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CostManagementConfig() ContainerClusterCostManagementConfigOutputReference
	CostManagementConfigInput() *ContainerClusterCostManagementConfig
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DatabaseEncryption() ContainerClusterDatabaseEncryptionOutputReference
	DatabaseEncryptionInput() *ContainerClusterDatabaseEncryption
	DatapathProvider() *string
	SetDatapathProvider(val *string)
	DatapathProviderInput() *string
	DefaultMaxPodsPerNode() *float64
	SetDefaultMaxPodsPerNode(val *float64)
	DefaultMaxPodsPerNodeInput() *float64
	DefaultSnatStatus() ContainerClusterDefaultSnatStatusOutputReference
	DefaultSnatStatusInput() *ContainerClusterDefaultSnatStatus
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DnsConfig() ContainerClusterDnsConfigOutputReference
	DnsConfigInput() *ContainerClusterDnsConfig
	EnableAutopilot() interface{}
	SetEnableAutopilot(val interface{})
	EnableAutopilotInput() interface{}
	EnableBinaryAuthorization() interface{}
	SetEnableBinaryAuthorization(val interface{})
	EnableBinaryAuthorizationInput() interface{}
	EnableIntranodeVisibility() interface{}
	SetEnableIntranodeVisibility(val interface{})
	EnableIntranodeVisibilityInput() interface{}
	EnableK8SBetaApis() ContainerClusterEnableK8SBetaApisOutputReference
	EnableK8SBetaApisInput() *ContainerClusterEnableK8SBetaApis
	EnableKubernetesAlpha() interface{}
	SetEnableKubernetesAlpha(val interface{})
	EnableKubernetesAlphaInput() interface{}
	EnableL4IlbSubsetting() interface{}
	SetEnableL4IlbSubsetting(val interface{})
	EnableL4IlbSubsettingInput() interface{}
	EnableLegacyAbac() interface{}
	SetEnableLegacyAbac(val interface{})
	EnableLegacyAbacInput() interface{}
	EnableShieldedNodes() interface{}
	SetEnableShieldedNodes(val interface{})
	EnableShieldedNodesInput() interface{}
	EnableTpu() interface{}
	SetEnableTpu(val interface{})
	EnableTpuInput() interface{}
	Endpoint() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GatewayApiConfig() ContainerClusterGatewayApiConfigOutputReference
	GatewayApiConfigInput() *ContainerClusterGatewayApiConfig
	Id() *string
	SetId(val *string)
	IdInput() *string
	InitialNodeCount() *float64
	SetInitialNodeCount(val *float64)
	InitialNodeCountInput() *float64
	IpAllocationPolicy() ContainerClusterIpAllocationPolicyOutputReference
	IpAllocationPolicyInput() *ContainerClusterIpAllocationPolicy
	LabelFingerprint() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	LoggingConfig() ContainerClusterLoggingConfigOutputReference
	LoggingConfigInput() *ContainerClusterLoggingConfig
	LoggingService() *string
	SetLoggingService(val *string)
	LoggingServiceInput() *string
	MaintenancePolicy() ContainerClusterMaintenancePolicyOutputReference
	MaintenancePolicyInput() *ContainerClusterMaintenancePolicy
	MasterAuth() ContainerClusterMasterAuthOutputReference
	MasterAuthInput() *ContainerClusterMasterAuth
	MasterAuthorizedNetworksConfig() ContainerClusterMasterAuthorizedNetworksConfigOutputReference
	MasterAuthorizedNetworksConfigInput() *ContainerClusterMasterAuthorizedNetworksConfig
	MasterVersion() *string
	MeshCertificates() ContainerClusterMeshCertificatesOutputReference
	MeshCertificatesInput() *ContainerClusterMeshCertificates
	MinMasterVersion() *string
	SetMinMasterVersion(val *string)
	MinMasterVersionInput() *string
	MonitoringConfig() ContainerClusterMonitoringConfigOutputReference
	MonitoringConfigInput() *ContainerClusterMonitoringConfig
	MonitoringService() *string
	SetMonitoringService(val *string)
	MonitoringServiceInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Network() *string
	SetNetwork(val *string)
	NetworkingMode() *string
	SetNetworkingMode(val *string)
	NetworkingModeInput() *string
	NetworkInput() *string
	NetworkPolicy() ContainerClusterNetworkPolicyOutputReference
	NetworkPolicyInput() *ContainerClusterNetworkPolicy
	// The tree node.
	Node() constructs.Node
	NodeConfig() ContainerClusterNodeConfigOutputReference
	NodeConfigInput() *ContainerClusterNodeConfig
	NodeLocations() *[]*string
	SetNodeLocations(val *[]*string)
	NodeLocationsInput() *[]*string
	NodePool() ContainerClusterNodePoolList
	NodePoolDefaults() ContainerClusterNodePoolDefaultsOutputReference
	NodePoolDefaultsInput() *ContainerClusterNodePoolDefaults
	NodePoolInput() interface{}
	NodeVersion() *string
	SetNodeVersion(val *string)
	NodeVersionInput() *string
	NotificationConfig() ContainerClusterNotificationConfigOutputReference
	NotificationConfigInput() *ContainerClusterNotificationConfig
	Operation() *string
	PrivateClusterConfig() ContainerClusterPrivateClusterConfigOutputReference
	PrivateClusterConfigInput() *ContainerClusterPrivateClusterConfig
	PrivateIpv6GoogleAccess() *string
	SetPrivateIpv6GoogleAccess(val *string)
	PrivateIpv6GoogleAccessInput() *string
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
	ReleaseChannel() ContainerClusterReleaseChannelOutputReference
	ReleaseChannelInput() *ContainerClusterReleaseChannel
	RemoveDefaultNodePool() interface{}
	SetRemoveDefaultNodePool(val interface{})
	RemoveDefaultNodePoolInput() interface{}
	ResourceLabels() *map[string]*string
	SetResourceLabels(val *map[string]*string)
	ResourceLabelsInput() *map[string]*string
	ResourceUsageExportConfig() ContainerClusterResourceUsageExportConfigOutputReference
	ResourceUsageExportConfigInput() *ContainerClusterResourceUsageExportConfig
	SecurityPostureConfig() ContainerClusterSecurityPostureConfigOutputReference
	SecurityPostureConfigInput() *ContainerClusterSecurityPostureConfig
	SelfLink() *string
	ServiceExternalIpsConfig() ContainerClusterServiceExternalIpsConfigOutputReference
	ServiceExternalIpsConfigInput() *ContainerClusterServiceExternalIpsConfig
	ServicesIpv4Cidr() *string
	Subnetwork() *string
	SetSubnetwork(val *string)
	SubnetworkInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ContainerClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
	TpuIpv4CidrBlock() *string
	VerticalPodAutoscaling() ContainerClusterVerticalPodAutoscalingOutputReference
	VerticalPodAutoscalingInput() *ContainerClusterVerticalPodAutoscaling
	WorkloadIdentityConfig() ContainerClusterWorkloadIdentityConfigOutputReference
	WorkloadIdentityConfigInput() *ContainerClusterWorkloadIdentityConfig
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
	PutAddonsConfig(value *ContainerClusterAddonsConfig)
	PutAuthenticatorGroupsConfig(value *ContainerClusterAuthenticatorGroupsConfig)
	PutBinaryAuthorization(value *ContainerClusterBinaryAuthorization)
	PutClusterAutoscaling(value *ContainerClusterClusterAutoscaling)
	PutConfidentialNodes(value *ContainerClusterConfidentialNodes)
	PutCostManagementConfig(value *ContainerClusterCostManagementConfig)
	PutDatabaseEncryption(value *ContainerClusterDatabaseEncryption)
	PutDefaultSnatStatus(value *ContainerClusterDefaultSnatStatus)
	PutDnsConfig(value *ContainerClusterDnsConfig)
	PutEnableK8SBetaApis(value *ContainerClusterEnableK8SBetaApis)
	PutGatewayApiConfig(value *ContainerClusterGatewayApiConfig)
	PutIpAllocationPolicy(value *ContainerClusterIpAllocationPolicy)
	PutLoggingConfig(value *ContainerClusterLoggingConfig)
	PutMaintenancePolicy(value *ContainerClusterMaintenancePolicy)
	PutMasterAuth(value *ContainerClusterMasterAuth)
	PutMasterAuthorizedNetworksConfig(value *ContainerClusterMasterAuthorizedNetworksConfig)
	PutMeshCertificates(value *ContainerClusterMeshCertificates)
	PutMonitoringConfig(value *ContainerClusterMonitoringConfig)
	PutNetworkPolicy(value *ContainerClusterNetworkPolicy)
	PutNodeConfig(value *ContainerClusterNodeConfig)
	PutNodePool(value interface{})
	PutNodePoolDefaults(value *ContainerClusterNodePoolDefaults)
	PutNotificationConfig(value *ContainerClusterNotificationConfig)
	PutPrivateClusterConfig(value *ContainerClusterPrivateClusterConfig)
	PutReleaseChannel(value *ContainerClusterReleaseChannel)
	PutResourceUsageExportConfig(value *ContainerClusterResourceUsageExportConfig)
	PutSecurityPostureConfig(value *ContainerClusterSecurityPostureConfig)
	PutServiceExternalIpsConfig(value *ContainerClusterServiceExternalIpsConfig)
	PutTimeouts(value *ContainerClusterTimeouts)
	PutVerticalPodAutoscaling(value *ContainerClusterVerticalPodAutoscaling)
	PutWorkloadIdentityConfig(value *ContainerClusterWorkloadIdentityConfig)
	ResetAddonsConfig()
	ResetAllowNetAdmin()
	ResetAuthenticatorGroupsConfig()
	ResetBinaryAuthorization()
	ResetClusterAutoscaling()
	ResetClusterIpv4Cidr()
	ResetConfidentialNodes()
	ResetCostManagementConfig()
	ResetDatabaseEncryption()
	ResetDatapathProvider()
	ResetDefaultMaxPodsPerNode()
	ResetDefaultSnatStatus()
	ResetDescription()
	ResetDnsConfig()
	ResetEnableAutopilot()
	ResetEnableBinaryAuthorization()
	ResetEnableIntranodeVisibility()
	ResetEnableK8SBetaApis()
	ResetEnableKubernetesAlpha()
	ResetEnableL4IlbSubsetting()
	ResetEnableLegacyAbac()
	ResetEnableShieldedNodes()
	ResetEnableTpu()
	ResetGatewayApiConfig()
	ResetId()
	ResetInitialNodeCount()
	ResetIpAllocationPolicy()
	ResetLocation()
	ResetLoggingConfig()
	ResetLoggingService()
	ResetMaintenancePolicy()
	ResetMasterAuth()
	ResetMasterAuthorizedNetworksConfig()
	ResetMeshCertificates()
	ResetMinMasterVersion()
	ResetMonitoringConfig()
	ResetMonitoringService()
	ResetNetwork()
	ResetNetworkingMode()
	ResetNetworkPolicy()
	ResetNodeConfig()
	ResetNodeLocations()
	ResetNodePool()
	ResetNodePoolDefaults()
	ResetNodeVersion()
	ResetNotificationConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrivateClusterConfig()
	ResetPrivateIpv6GoogleAccess()
	ResetProject()
	ResetReleaseChannel()
	ResetRemoveDefaultNodePool()
	ResetResourceLabels()
	ResetResourceUsageExportConfig()
	ResetSecurityPostureConfig()
	ResetServiceExternalIpsConfig()
	ResetSubnetwork()
	ResetTimeouts()
	ResetVerticalPodAutoscaling()
	ResetWorkloadIdentityConfig()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ContainerCluster
type jsiiProxy_ContainerCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ContainerCluster) AddonsConfig() ContainerClusterAddonsConfigOutputReference {
	var returns ContainerClusterAddonsConfigOutputReference
	_jsii_.Get(
		j,
		"addonsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) AddonsConfigInput() *ContainerClusterAddonsConfig {
	var returns *ContainerClusterAddonsConfig
	_jsii_.Get(
		j,
		"addonsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) AllowNetAdmin() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowNetAdmin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) AllowNetAdminInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowNetAdminInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) AuthenticatorGroupsConfig() ContainerClusterAuthenticatorGroupsConfigOutputReference {
	var returns ContainerClusterAuthenticatorGroupsConfigOutputReference
	_jsii_.Get(
		j,
		"authenticatorGroupsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) AuthenticatorGroupsConfigInput() *ContainerClusterAuthenticatorGroupsConfig {
	var returns *ContainerClusterAuthenticatorGroupsConfig
	_jsii_.Get(
		j,
		"authenticatorGroupsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) BinaryAuthorization() ContainerClusterBinaryAuthorizationOutputReference {
	var returns ContainerClusterBinaryAuthorizationOutputReference
	_jsii_.Get(
		j,
		"binaryAuthorization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) BinaryAuthorizationInput() *ContainerClusterBinaryAuthorization {
	var returns *ContainerClusterBinaryAuthorization
	_jsii_.Get(
		j,
		"binaryAuthorizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) ClusterAutoscaling() ContainerClusterClusterAutoscalingOutputReference {
	var returns ContainerClusterClusterAutoscalingOutputReference
	_jsii_.Get(
		j,
		"clusterAutoscaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) ClusterAutoscalingInput() *ContainerClusterClusterAutoscaling {
	var returns *ContainerClusterClusterAutoscaling
	_jsii_.Get(
		j,
		"clusterAutoscalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) ClusterIpv4Cidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIpv4Cidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) ClusterIpv4CidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIpv4CidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) ConfidentialNodes() ContainerClusterConfidentialNodesOutputReference {
	var returns ContainerClusterConfidentialNodesOutputReference
	_jsii_.Get(
		j,
		"confidentialNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) ConfidentialNodesInput() *ContainerClusterConfidentialNodes {
	var returns *ContainerClusterConfidentialNodes
	_jsii_.Get(
		j,
		"confidentialNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) CostManagementConfig() ContainerClusterCostManagementConfigOutputReference {
	var returns ContainerClusterCostManagementConfigOutputReference
	_jsii_.Get(
		j,
		"costManagementConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) CostManagementConfigInput() *ContainerClusterCostManagementConfig {
	var returns *ContainerClusterCostManagementConfig
	_jsii_.Get(
		j,
		"costManagementConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) DatabaseEncryption() ContainerClusterDatabaseEncryptionOutputReference {
	var returns ContainerClusterDatabaseEncryptionOutputReference
	_jsii_.Get(
		j,
		"databaseEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) DatabaseEncryptionInput() *ContainerClusterDatabaseEncryption {
	var returns *ContainerClusterDatabaseEncryption
	_jsii_.Get(
		j,
		"databaseEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) DatapathProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datapathProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) DatapathProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datapathProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) DefaultMaxPodsPerNode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultMaxPodsPerNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) DefaultMaxPodsPerNodeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultMaxPodsPerNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) DefaultSnatStatus() ContainerClusterDefaultSnatStatusOutputReference {
	var returns ContainerClusterDefaultSnatStatusOutputReference
	_jsii_.Get(
		j,
		"defaultSnatStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) DefaultSnatStatusInput() *ContainerClusterDefaultSnatStatus {
	var returns *ContainerClusterDefaultSnatStatus
	_jsii_.Get(
		j,
		"defaultSnatStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) DnsConfig() ContainerClusterDnsConfigOutputReference {
	var returns ContainerClusterDnsConfigOutputReference
	_jsii_.Get(
		j,
		"dnsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) DnsConfigInput() *ContainerClusterDnsConfig {
	var returns *ContainerClusterDnsConfig
	_jsii_.Get(
		j,
		"dnsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) EnableAutopilot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutopilot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) EnableAutopilotInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutopilotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) EnableBinaryAuthorization() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBinaryAuthorization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) EnableBinaryAuthorizationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableBinaryAuthorizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) EnableIntranodeVisibility() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableIntranodeVisibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) EnableIntranodeVisibilityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableIntranodeVisibilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) EnableK8SBetaApis() ContainerClusterEnableK8SBetaApisOutputReference {
	var returns ContainerClusterEnableK8SBetaApisOutputReference
	_jsii_.Get(
		j,
		"enableK8SBetaApis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) EnableK8SBetaApisInput() *ContainerClusterEnableK8SBetaApis {
	var returns *ContainerClusterEnableK8SBetaApis
	_jsii_.Get(
		j,
		"enableK8SBetaApisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) EnableKubernetesAlpha() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableKubernetesAlpha",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) EnableKubernetesAlphaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableKubernetesAlphaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) EnableL4IlbSubsetting() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableL4IlbSubsetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) EnableL4IlbSubsettingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableL4IlbSubsettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) EnableLegacyAbac() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLegacyAbac",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) EnableLegacyAbacInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLegacyAbacInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) EnableShieldedNodes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableShieldedNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) EnableShieldedNodesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableShieldedNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) EnableTpu() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableTpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) EnableTpuInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableTpuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) GatewayApiConfig() ContainerClusterGatewayApiConfigOutputReference {
	var returns ContainerClusterGatewayApiConfigOutputReference
	_jsii_.Get(
		j,
		"gatewayApiConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) GatewayApiConfigInput() *ContainerClusterGatewayApiConfig {
	var returns *ContainerClusterGatewayApiConfig
	_jsii_.Get(
		j,
		"gatewayApiConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) InitialNodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialNodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) InitialNodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialNodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) IpAllocationPolicy() ContainerClusterIpAllocationPolicyOutputReference {
	var returns ContainerClusterIpAllocationPolicyOutputReference
	_jsii_.Get(
		j,
		"ipAllocationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) IpAllocationPolicyInput() *ContainerClusterIpAllocationPolicy {
	var returns *ContainerClusterIpAllocationPolicy
	_jsii_.Get(
		j,
		"ipAllocationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) LabelFingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelFingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) LoggingConfig() ContainerClusterLoggingConfigOutputReference {
	var returns ContainerClusterLoggingConfigOutputReference
	_jsii_.Get(
		j,
		"loggingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) LoggingConfigInput() *ContainerClusterLoggingConfig {
	var returns *ContainerClusterLoggingConfig
	_jsii_.Get(
		j,
		"loggingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) LoggingService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) LoggingServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) MaintenancePolicy() ContainerClusterMaintenancePolicyOutputReference {
	var returns ContainerClusterMaintenancePolicyOutputReference
	_jsii_.Get(
		j,
		"maintenancePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) MaintenancePolicyInput() *ContainerClusterMaintenancePolicy {
	var returns *ContainerClusterMaintenancePolicy
	_jsii_.Get(
		j,
		"maintenancePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) MasterAuth() ContainerClusterMasterAuthOutputReference {
	var returns ContainerClusterMasterAuthOutputReference
	_jsii_.Get(
		j,
		"masterAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) MasterAuthInput() *ContainerClusterMasterAuth {
	var returns *ContainerClusterMasterAuth
	_jsii_.Get(
		j,
		"masterAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) MasterAuthorizedNetworksConfig() ContainerClusterMasterAuthorizedNetworksConfigOutputReference {
	var returns ContainerClusterMasterAuthorizedNetworksConfigOutputReference
	_jsii_.Get(
		j,
		"masterAuthorizedNetworksConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) MasterAuthorizedNetworksConfigInput() *ContainerClusterMasterAuthorizedNetworksConfig {
	var returns *ContainerClusterMasterAuthorizedNetworksConfig
	_jsii_.Get(
		j,
		"masterAuthorizedNetworksConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) MasterVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) MeshCertificates() ContainerClusterMeshCertificatesOutputReference {
	var returns ContainerClusterMeshCertificatesOutputReference
	_jsii_.Get(
		j,
		"meshCertificates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) MeshCertificatesInput() *ContainerClusterMeshCertificates {
	var returns *ContainerClusterMeshCertificates
	_jsii_.Get(
		j,
		"meshCertificatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) MinMasterVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minMasterVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) MinMasterVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minMasterVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) MonitoringConfig() ContainerClusterMonitoringConfigOutputReference {
	var returns ContainerClusterMonitoringConfigOutputReference
	_jsii_.Get(
		j,
		"monitoringConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) MonitoringConfigInput() *ContainerClusterMonitoringConfig {
	var returns *ContainerClusterMonitoringConfig
	_jsii_.Get(
		j,
		"monitoringConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) MonitoringService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) MonitoringServiceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) NetworkingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) NetworkingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) NetworkPolicy() ContainerClusterNetworkPolicyOutputReference {
	var returns ContainerClusterNetworkPolicyOutputReference
	_jsii_.Get(
		j,
		"networkPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) NetworkPolicyInput() *ContainerClusterNetworkPolicy {
	var returns *ContainerClusterNetworkPolicy
	_jsii_.Get(
		j,
		"networkPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) NodeConfig() ContainerClusterNodeConfigOutputReference {
	var returns ContainerClusterNodeConfigOutputReference
	_jsii_.Get(
		j,
		"nodeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) NodeConfigInput() *ContainerClusterNodeConfig {
	var returns *ContainerClusterNodeConfig
	_jsii_.Get(
		j,
		"nodeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) NodeLocations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nodeLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) NodeLocationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nodeLocationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) NodePool() ContainerClusterNodePoolList {
	var returns ContainerClusterNodePoolList
	_jsii_.Get(
		j,
		"nodePool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) NodePoolDefaults() ContainerClusterNodePoolDefaultsOutputReference {
	var returns ContainerClusterNodePoolDefaultsOutputReference
	_jsii_.Get(
		j,
		"nodePoolDefaults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) NodePoolDefaultsInput() *ContainerClusterNodePoolDefaults {
	var returns *ContainerClusterNodePoolDefaults
	_jsii_.Get(
		j,
		"nodePoolDefaultsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) NodePoolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nodePoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) NodeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) NodeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) NotificationConfig() ContainerClusterNotificationConfigOutputReference {
	var returns ContainerClusterNotificationConfigOutputReference
	_jsii_.Get(
		j,
		"notificationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) NotificationConfigInput() *ContainerClusterNotificationConfig {
	var returns *ContainerClusterNotificationConfig
	_jsii_.Get(
		j,
		"notificationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) Operation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) PrivateClusterConfig() ContainerClusterPrivateClusterConfigOutputReference {
	var returns ContainerClusterPrivateClusterConfigOutputReference
	_jsii_.Get(
		j,
		"privateClusterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) PrivateClusterConfigInput() *ContainerClusterPrivateClusterConfig {
	var returns *ContainerClusterPrivateClusterConfig
	_jsii_.Get(
		j,
		"privateClusterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) PrivateIpv6GoogleAccess() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpv6GoogleAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) PrivateIpv6GoogleAccessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpv6GoogleAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) ReleaseChannel() ContainerClusterReleaseChannelOutputReference {
	var returns ContainerClusterReleaseChannelOutputReference
	_jsii_.Get(
		j,
		"releaseChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) ReleaseChannelInput() *ContainerClusterReleaseChannel {
	var returns *ContainerClusterReleaseChannel
	_jsii_.Get(
		j,
		"releaseChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) RemoveDefaultNodePool() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removeDefaultNodePool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) RemoveDefaultNodePoolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removeDefaultNodePoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) ResourceLabels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"resourceLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) ResourceLabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"resourceLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) ResourceUsageExportConfig() ContainerClusterResourceUsageExportConfigOutputReference {
	var returns ContainerClusterResourceUsageExportConfigOutputReference
	_jsii_.Get(
		j,
		"resourceUsageExportConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) ResourceUsageExportConfigInput() *ContainerClusterResourceUsageExportConfig {
	var returns *ContainerClusterResourceUsageExportConfig
	_jsii_.Get(
		j,
		"resourceUsageExportConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) SecurityPostureConfig() ContainerClusterSecurityPostureConfigOutputReference {
	var returns ContainerClusterSecurityPostureConfigOutputReference
	_jsii_.Get(
		j,
		"securityPostureConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) SecurityPostureConfigInput() *ContainerClusterSecurityPostureConfig {
	var returns *ContainerClusterSecurityPostureConfig
	_jsii_.Get(
		j,
		"securityPostureConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) ServiceExternalIpsConfig() ContainerClusterServiceExternalIpsConfigOutputReference {
	var returns ContainerClusterServiceExternalIpsConfigOutputReference
	_jsii_.Get(
		j,
		"serviceExternalIpsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) ServiceExternalIpsConfigInput() *ContainerClusterServiceExternalIpsConfig {
	var returns *ContainerClusterServiceExternalIpsConfig
	_jsii_.Get(
		j,
		"serviceExternalIpsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) ServicesIpv4Cidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicesIpv4Cidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) Subnetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) SubnetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) Timeouts() ContainerClusterTimeoutsOutputReference {
	var returns ContainerClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) TpuIpv4CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpuIpv4CidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) VerticalPodAutoscaling() ContainerClusterVerticalPodAutoscalingOutputReference {
	var returns ContainerClusterVerticalPodAutoscalingOutputReference
	_jsii_.Get(
		j,
		"verticalPodAutoscaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) VerticalPodAutoscalingInput() *ContainerClusterVerticalPodAutoscaling {
	var returns *ContainerClusterVerticalPodAutoscaling
	_jsii_.Get(
		j,
		"verticalPodAutoscalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) WorkloadIdentityConfig() ContainerClusterWorkloadIdentityConfigOutputReference {
	var returns ContainerClusterWorkloadIdentityConfigOutputReference
	_jsii_.Get(
		j,
		"workloadIdentityConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ContainerCluster) WorkloadIdentityConfigInput() *ContainerClusterWorkloadIdentityConfig {
	var returns *ContainerClusterWorkloadIdentityConfig
	_jsii_.Get(
		j,
		"workloadIdentityConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/4.83.0/docs/resources/container_cluster google_container_cluster} Resource.
func NewContainerCluster(scope constructs.Construct, id *string, config *ContainerClusterConfig) ContainerCluster {
	_init_.Initialize()

	if err := validateNewContainerClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ContainerCluster{}

	_jsii_.Create(
		"@cdktf/provider-google.containerCluster.ContainerCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/4.83.0/docs/resources/container_cluster google_container_cluster} Resource.
func NewContainerCluster_Override(c ContainerCluster, scope constructs.Construct, id *string, config *ContainerClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.containerCluster.ContainerCluster",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ContainerCluster)SetAllowNetAdmin(val interface{}) {
	if err := j.validateSetAllowNetAdminParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowNetAdmin",
		val,
	)
}

func (j *jsiiProxy_ContainerCluster)SetClusterIpv4Cidr(val *string) {
	if err := j.validateSetClusterIpv4CidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterIpv4Cidr",
		val,
	)
}

func (j *jsiiProxy_ContainerCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ContainerCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ContainerCluster)SetDatapathProvider(val *string) {
	if err := j.validateSetDatapathProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"datapathProvider",
		val,
	)
}

func (j *jsiiProxy_ContainerCluster)SetDefaultMaxPodsPerNode(val *float64) {
	if err := j.validateSetDefaultMaxPodsPerNodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultMaxPodsPerNode",
		val,
	)
}

func (j *jsiiProxy_ContainerCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ContainerCluster)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ContainerCluster)SetEnableAutopilot(val interface{}) {
	if err := j.validateSetEnableAutopilotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAutopilot",
		val,
	)
}

func (j *jsiiProxy_ContainerCluster)SetEnableBinaryAuthorization(val interface{}) {
	if err := j.validateSetEnableBinaryAuthorizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableBinaryAuthorization",
		val,
	)
}

func (j *jsiiProxy_ContainerCluster)SetEnableIntranodeVisibility(val interface{}) {
	if err := j.validateSetEnableIntranodeVisibilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableIntranodeVisibility",
		val,
	)
}

func (j *jsiiProxy_ContainerCluster)SetEnableKubernetesAlpha(val interface{}) {
	if err := j.validateSetEnableKubernetesAlphaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableKubernetesAlpha",
		val,
	)
}

func (j *jsiiProxy_ContainerCluster)SetEnableL4IlbSubsetting(val interface{}) {
	if err := j.validateSetEnableL4IlbSubsettingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableL4IlbSubsetting",
		val,
	)
}

func (j *jsiiProxy_ContainerCluster)SetEnableLegacyAbac(val interface{}) {
	if err := j.validateSetEnableLegacyAbacParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableLegacyAbac",
		val,
	)
}

func (j *jsiiProxy_ContainerCluster)SetEnableShieldedNodes(val interface{}) {
	if err := j.validateSetEnableShieldedNodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableShieldedNodes",
		val,
	)
}

func (j *jsiiProxy_ContainerCluster)SetEnableTpu(val interface{}) {
	if err := j.validateSetEnableTpuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableTpu",
		val,
	)
}

func (j *jsiiProxy_ContainerCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ContainerCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ContainerCluster)SetInitialNodeCount(val *float64) {
	if err := j.validateSetInitialNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialNodeCount",
		val,
	)
}

func (j *jsiiProxy_ContainerCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ContainerCluster)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_ContainerCluster)SetLoggingService(val *string) {
	if err := j.validateSetLoggingServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loggingService",
		val,
	)
}

func (j *jsiiProxy_ContainerCluster)SetMinMasterVersion(val *string) {
	if err := j.validateSetMinMasterVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minMasterVersion",
		val,
	)
}

func (j *jsiiProxy_ContainerCluster)SetMonitoringService(val *string) {
	if err := j.validateSetMonitoringServiceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoringService",
		val,
	)
}

func (j *jsiiProxy_ContainerCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ContainerCluster)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_ContainerCluster)SetNetworkingMode(val *string) {
	if err := j.validateSetNetworkingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkingMode",
		val,
	)
}

func (j *jsiiProxy_ContainerCluster)SetNodeLocations(val *[]*string) {
	if err := j.validateSetNodeLocationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeLocations",
		val,
	)
}

func (j *jsiiProxy_ContainerCluster)SetNodeVersion(val *string) {
	if err := j.validateSetNodeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeVersion",
		val,
	)
}

func (j *jsiiProxy_ContainerCluster)SetPrivateIpv6GoogleAccess(val *string) {
	if err := j.validateSetPrivateIpv6GoogleAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateIpv6GoogleAccess",
		val,
	)
}

func (j *jsiiProxy_ContainerCluster)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_ContainerCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ContainerCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ContainerCluster)SetRemoveDefaultNodePool(val interface{}) {
	if err := j.validateSetRemoveDefaultNodePoolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"removeDefaultNodePool",
		val,
	)
}

func (j *jsiiProxy_ContainerCluster)SetResourceLabels(val *map[string]*string) {
	if err := j.validateSetResourceLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceLabels",
		val,
	)
}

func (j *jsiiProxy_ContainerCluster)SetSubnetwork(val *string) {
	if err := j.validateSetSubnetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetwork",
		val,
	)
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
func ContainerCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateContainerCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.containerCluster.ContainerCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ContainerCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateContainerCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.containerCluster.ContainerCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ContainerCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateContainerCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.containerCluster.ContainerCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ContainerCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.containerCluster.ContainerCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ContainerCluster) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ContainerCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_ContainerCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ContainerCluster) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ContainerCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ContainerCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ContainerCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ContainerCluster) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ContainerCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ContainerCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ContainerCluster) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ContainerCluster) PutAddonsConfig(value *ContainerClusterAddonsConfig) {
	if err := c.validatePutAddonsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAddonsConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerCluster) PutAuthenticatorGroupsConfig(value *ContainerClusterAuthenticatorGroupsConfig) {
	if err := c.validatePutAuthenticatorGroupsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAuthenticatorGroupsConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerCluster) PutBinaryAuthorization(value *ContainerClusterBinaryAuthorization) {
	if err := c.validatePutBinaryAuthorizationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBinaryAuthorization",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerCluster) PutClusterAutoscaling(value *ContainerClusterClusterAutoscaling) {
	if err := c.validatePutClusterAutoscalingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putClusterAutoscaling",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerCluster) PutConfidentialNodes(value *ContainerClusterConfidentialNodes) {
	if err := c.validatePutConfidentialNodesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putConfidentialNodes",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerCluster) PutCostManagementConfig(value *ContainerClusterCostManagementConfig) {
	if err := c.validatePutCostManagementConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCostManagementConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerCluster) PutDatabaseEncryption(value *ContainerClusterDatabaseEncryption) {
	if err := c.validatePutDatabaseEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDatabaseEncryption",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerCluster) PutDefaultSnatStatus(value *ContainerClusterDefaultSnatStatus) {
	if err := c.validatePutDefaultSnatStatusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDefaultSnatStatus",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerCluster) PutDnsConfig(value *ContainerClusterDnsConfig) {
	if err := c.validatePutDnsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDnsConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerCluster) PutEnableK8SBetaApis(value *ContainerClusterEnableK8SBetaApis) {
	if err := c.validatePutEnableK8SBetaApisParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEnableK8SBetaApis",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerCluster) PutGatewayApiConfig(value *ContainerClusterGatewayApiConfig) {
	if err := c.validatePutGatewayApiConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGatewayApiConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerCluster) PutIpAllocationPolicy(value *ContainerClusterIpAllocationPolicy) {
	if err := c.validatePutIpAllocationPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putIpAllocationPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerCluster) PutLoggingConfig(value *ContainerClusterLoggingConfig) {
	if err := c.validatePutLoggingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putLoggingConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerCluster) PutMaintenancePolicy(value *ContainerClusterMaintenancePolicy) {
	if err := c.validatePutMaintenancePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMaintenancePolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerCluster) PutMasterAuth(value *ContainerClusterMasterAuth) {
	if err := c.validatePutMasterAuthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMasterAuth",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerCluster) PutMasterAuthorizedNetworksConfig(value *ContainerClusterMasterAuthorizedNetworksConfig) {
	if err := c.validatePutMasterAuthorizedNetworksConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMasterAuthorizedNetworksConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerCluster) PutMeshCertificates(value *ContainerClusterMeshCertificates) {
	if err := c.validatePutMeshCertificatesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMeshCertificates",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerCluster) PutMonitoringConfig(value *ContainerClusterMonitoringConfig) {
	if err := c.validatePutMonitoringConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMonitoringConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerCluster) PutNetworkPolicy(value *ContainerClusterNetworkPolicy) {
	if err := c.validatePutNetworkPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNetworkPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerCluster) PutNodeConfig(value *ContainerClusterNodeConfig) {
	if err := c.validatePutNodeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNodeConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerCluster) PutNodePool(value interface{}) {
	if err := c.validatePutNodePoolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNodePool",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerCluster) PutNodePoolDefaults(value *ContainerClusterNodePoolDefaults) {
	if err := c.validatePutNodePoolDefaultsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNodePoolDefaults",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerCluster) PutNotificationConfig(value *ContainerClusterNotificationConfig) {
	if err := c.validatePutNotificationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNotificationConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerCluster) PutPrivateClusterConfig(value *ContainerClusterPrivateClusterConfig) {
	if err := c.validatePutPrivateClusterConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPrivateClusterConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerCluster) PutReleaseChannel(value *ContainerClusterReleaseChannel) {
	if err := c.validatePutReleaseChannelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putReleaseChannel",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerCluster) PutResourceUsageExportConfig(value *ContainerClusterResourceUsageExportConfig) {
	if err := c.validatePutResourceUsageExportConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putResourceUsageExportConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerCluster) PutSecurityPostureConfig(value *ContainerClusterSecurityPostureConfig) {
	if err := c.validatePutSecurityPostureConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSecurityPostureConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerCluster) PutServiceExternalIpsConfig(value *ContainerClusterServiceExternalIpsConfig) {
	if err := c.validatePutServiceExternalIpsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putServiceExternalIpsConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerCluster) PutTimeouts(value *ContainerClusterTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerCluster) PutVerticalPodAutoscaling(value *ContainerClusterVerticalPodAutoscaling) {
	if err := c.validatePutVerticalPodAutoscalingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putVerticalPodAutoscaling",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerCluster) PutWorkloadIdentityConfig(value *ContainerClusterWorkloadIdentityConfig) {
	if err := c.validatePutWorkloadIdentityConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putWorkloadIdentityConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ContainerCluster) ResetAddonsConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetAddonsConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetAllowNetAdmin() {
	_jsii_.InvokeVoid(
		c,
		"resetAllowNetAdmin",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetAuthenticatorGroupsConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetAuthenticatorGroupsConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetBinaryAuthorization() {
	_jsii_.InvokeVoid(
		c,
		"resetBinaryAuthorization",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetClusterAutoscaling() {
	_jsii_.InvokeVoid(
		c,
		"resetClusterAutoscaling",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetClusterIpv4Cidr() {
	_jsii_.InvokeVoid(
		c,
		"resetClusterIpv4Cidr",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetConfidentialNodes() {
	_jsii_.InvokeVoid(
		c,
		"resetConfidentialNodes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetCostManagementConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetCostManagementConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetDatabaseEncryption() {
	_jsii_.InvokeVoid(
		c,
		"resetDatabaseEncryption",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetDatapathProvider() {
	_jsii_.InvokeVoid(
		c,
		"resetDatapathProvider",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetDefaultMaxPodsPerNode() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultMaxPodsPerNode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetDefaultSnatStatus() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultSnatStatus",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetDnsConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetDnsConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetEnableAutopilot() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableAutopilot",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetEnableBinaryAuthorization() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableBinaryAuthorization",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetEnableIntranodeVisibility() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableIntranodeVisibility",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetEnableK8SBetaApis() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableK8SBetaApis",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetEnableKubernetesAlpha() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableKubernetesAlpha",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetEnableL4IlbSubsetting() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableL4IlbSubsetting",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetEnableLegacyAbac() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableLegacyAbac",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetEnableShieldedNodes() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableShieldedNodes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetEnableTpu() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableTpu",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetGatewayApiConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetGatewayApiConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetInitialNodeCount() {
	_jsii_.InvokeVoid(
		c,
		"resetInitialNodeCount",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetIpAllocationPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetIpAllocationPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetLocation() {
	_jsii_.InvokeVoid(
		c,
		"resetLocation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetLoggingConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetLoggingConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetLoggingService() {
	_jsii_.InvokeVoid(
		c,
		"resetLoggingService",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetMaintenancePolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetMaintenancePolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetMasterAuth() {
	_jsii_.InvokeVoid(
		c,
		"resetMasterAuth",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetMasterAuthorizedNetworksConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetMasterAuthorizedNetworksConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetMeshCertificates() {
	_jsii_.InvokeVoid(
		c,
		"resetMeshCertificates",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetMinMasterVersion() {
	_jsii_.InvokeVoid(
		c,
		"resetMinMasterVersion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetMonitoringConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetMonitoringConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetMonitoringService() {
	_jsii_.InvokeVoid(
		c,
		"resetMonitoringService",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetNetwork() {
	_jsii_.InvokeVoid(
		c,
		"resetNetwork",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetNetworkingMode() {
	_jsii_.InvokeVoid(
		c,
		"resetNetworkingMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetNetworkPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetNetworkPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetNodeConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetNodeConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetNodeLocations() {
	_jsii_.InvokeVoid(
		c,
		"resetNodeLocations",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetNodePool() {
	_jsii_.InvokeVoid(
		c,
		"resetNodePool",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetNodePoolDefaults() {
	_jsii_.InvokeVoid(
		c,
		"resetNodePoolDefaults",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetNodeVersion() {
	_jsii_.InvokeVoid(
		c,
		"resetNodeVersion",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetNotificationConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetNotificationConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetPrivateClusterConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetPrivateClusterConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetPrivateIpv6GoogleAccess() {
	_jsii_.InvokeVoid(
		c,
		"resetPrivateIpv6GoogleAccess",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetProject() {
	_jsii_.InvokeVoid(
		c,
		"resetProject",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetReleaseChannel() {
	_jsii_.InvokeVoid(
		c,
		"resetReleaseChannel",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetRemoveDefaultNodePool() {
	_jsii_.InvokeVoid(
		c,
		"resetRemoveDefaultNodePool",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetResourceLabels() {
	_jsii_.InvokeVoid(
		c,
		"resetResourceLabels",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetResourceUsageExportConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetResourceUsageExportConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetSecurityPostureConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetSecurityPostureConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetServiceExternalIpsConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetServiceExternalIpsConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetSubnetwork() {
	_jsii_.InvokeVoid(
		c,
		"resetSubnetwork",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetVerticalPodAutoscaling() {
	_jsii_.InvokeVoid(
		c,
		"resetVerticalPodAutoscaling",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) ResetWorkloadIdentityConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetWorkloadIdentityConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ContainerCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ContainerCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

