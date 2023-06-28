package datagooglecontainercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-google-go/google/v8/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-google-go/google/v8/datagooglecontainercluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/4.71.0/docs/data-sources/container_cluster google_container_cluster}.
type DataGoogleContainerCluster interface {
	cdktf.TerraformDataSource
	AddonsConfig() DataGoogleContainerClusterAddonsConfigList
	AuthenticatorGroupsConfig() DataGoogleContainerClusterAuthenticatorGroupsConfigList
	BinaryAuthorization() DataGoogleContainerClusterBinaryAuthorizationList
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterAutoscaling() DataGoogleContainerClusterClusterAutoscalingList
	ClusterIpv4Cidr() *string
	ConfidentialNodes() DataGoogleContainerClusterConfidentialNodesList
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CostManagementConfig() DataGoogleContainerClusterCostManagementConfigList
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DatabaseEncryption() DataGoogleContainerClusterDatabaseEncryptionList
	DatapathProvider() *string
	DefaultMaxPodsPerNode() *float64
	DefaultSnatStatus() DataGoogleContainerClusterDefaultSnatStatusList
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	DnsConfig() DataGoogleContainerClusterDnsConfigList
	EnableAutopilot() cdktf.IResolvable
	EnableBinaryAuthorization() cdktf.IResolvable
	EnableIntranodeVisibility() cdktf.IResolvable
	EnableKubernetesAlpha() cdktf.IResolvable
	EnableL4IlbSubsetting() cdktf.IResolvable
	EnableLegacyAbac() cdktf.IResolvable
	EnableShieldedNodes() cdktf.IResolvable
	EnableTpu() cdktf.IResolvable
	Endpoint() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GatewayApiConfig() DataGoogleContainerClusterGatewayApiConfigList
	Id() *string
	SetId(val *string)
	IdInput() *string
	InitialNodeCount() *float64
	IpAllocationPolicy() DataGoogleContainerClusterIpAllocationPolicyList
	LabelFingerprint() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	LoggingConfig() DataGoogleContainerClusterLoggingConfigList
	LoggingService() *string
	MaintenancePolicy() DataGoogleContainerClusterMaintenancePolicyList
	MasterAuth() DataGoogleContainerClusterMasterAuthList
	MasterAuthorizedNetworksConfig() DataGoogleContainerClusterMasterAuthorizedNetworksConfigList
	MasterVersion() *string
	MeshCertificates() DataGoogleContainerClusterMeshCertificatesList
	MinMasterVersion() *string
	MonitoringConfig() DataGoogleContainerClusterMonitoringConfigList
	MonitoringService() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Network() *string
	NetworkingMode() *string
	NetworkPolicy() DataGoogleContainerClusterNetworkPolicyList
	// The tree node.
	Node() constructs.Node
	NodeConfig() DataGoogleContainerClusterNodeConfigList
	NodeLocations() *[]*string
	NodePool() DataGoogleContainerClusterNodePoolList
	NodePoolDefaults() DataGoogleContainerClusterNodePoolDefaultsList
	NodeVersion() *string
	NotificationConfig() DataGoogleContainerClusterNotificationConfigList
	Operation() *string
	PrivateClusterConfig() DataGoogleContainerClusterPrivateClusterConfigList
	PrivateIpv6GoogleAccess() *string
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	ReleaseChannel() DataGoogleContainerClusterReleaseChannelList
	RemoveDefaultNodePool() cdktf.IResolvable
	ResourceLabels() cdktf.StringMap
	ResourceUsageExportConfig() DataGoogleContainerClusterResourceUsageExportConfigList
	SelfLink() *string
	ServiceExternalIpsConfig() DataGoogleContainerClusterServiceExternalIpsConfigList
	ServicesIpv4Cidr() *string
	Subnetwork() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TpuIpv4CidrBlock() *string
	VerticalPodAutoscaling() DataGoogleContainerClusterVerticalPodAutoscalingList
	WorkloadIdentityConfig() DataGoogleContainerClusterWorkloadIdentityConfigList
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
	ResetLocation()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataGoogleContainerCluster
type jsiiProxy_DataGoogleContainerCluster struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataGoogleContainerCluster) AddonsConfig() DataGoogleContainerClusterAddonsConfigList {
	var returns DataGoogleContainerClusterAddonsConfigList
	_jsii_.Get(
		j,
		"addonsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) AuthenticatorGroupsConfig() DataGoogleContainerClusterAuthenticatorGroupsConfigList {
	var returns DataGoogleContainerClusterAuthenticatorGroupsConfigList
	_jsii_.Get(
		j,
		"authenticatorGroupsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) BinaryAuthorization() DataGoogleContainerClusterBinaryAuthorizationList {
	var returns DataGoogleContainerClusterBinaryAuthorizationList
	_jsii_.Get(
		j,
		"binaryAuthorization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) ClusterAutoscaling() DataGoogleContainerClusterClusterAutoscalingList {
	var returns DataGoogleContainerClusterClusterAutoscalingList
	_jsii_.Get(
		j,
		"clusterAutoscaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) ClusterIpv4Cidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIpv4Cidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) ConfidentialNodes() DataGoogleContainerClusterConfidentialNodesList {
	var returns DataGoogleContainerClusterConfidentialNodesList
	_jsii_.Get(
		j,
		"confidentialNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) CostManagementConfig() DataGoogleContainerClusterCostManagementConfigList {
	var returns DataGoogleContainerClusterCostManagementConfigList
	_jsii_.Get(
		j,
		"costManagementConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) DatabaseEncryption() DataGoogleContainerClusterDatabaseEncryptionList {
	var returns DataGoogleContainerClusterDatabaseEncryptionList
	_jsii_.Get(
		j,
		"databaseEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) DatapathProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datapathProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) DefaultMaxPodsPerNode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultMaxPodsPerNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) DefaultSnatStatus() DataGoogleContainerClusterDefaultSnatStatusList {
	var returns DataGoogleContainerClusterDefaultSnatStatusList
	_jsii_.Get(
		j,
		"defaultSnatStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) DnsConfig() DataGoogleContainerClusterDnsConfigList {
	var returns DataGoogleContainerClusterDnsConfigList
	_jsii_.Get(
		j,
		"dnsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) EnableAutopilot() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enableAutopilot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) EnableBinaryAuthorization() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enableBinaryAuthorization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) EnableIntranodeVisibility() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enableIntranodeVisibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) EnableKubernetesAlpha() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enableKubernetesAlpha",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) EnableL4IlbSubsetting() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enableL4IlbSubsetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) EnableLegacyAbac() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enableLegacyAbac",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) EnableShieldedNodes() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enableShieldedNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) EnableTpu() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enableTpu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) GatewayApiConfig() DataGoogleContainerClusterGatewayApiConfigList {
	var returns DataGoogleContainerClusterGatewayApiConfigList
	_jsii_.Get(
		j,
		"gatewayApiConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) InitialNodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialNodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) IpAllocationPolicy() DataGoogleContainerClusterIpAllocationPolicyList {
	var returns DataGoogleContainerClusterIpAllocationPolicyList
	_jsii_.Get(
		j,
		"ipAllocationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) LabelFingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"labelFingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) LoggingConfig() DataGoogleContainerClusterLoggingConfigList {
	var returns DataGoogleContainerClusterLoggingConfigList
	_jsii_.Get(
		j,
		"loggingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) LoggingService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loggingService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) MaintenancePolicy() DataGoogleContainerClusterMaintenancePolicyList {
	var returns DataGoogleContainerClusterMaintenancePolicyList
	_jsii_.Get(
		j,
		"maintenancePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) MasterAuth() DataGoogleContainerClusterMasterAuthList {
	var returns DataGoogleContainerClusterMasterAuthList
	_jsii_.Get(
		j,
		"masterAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) MasterAuthorizedNetworksConfig() DataGoogleContainerClusterMasterAuthorizedNetworksConfigList {
	var returns DataGoogleContainerClusterMasterAuthorizedNetworksConfigList
	_jsii_.Get(
		j,
		"masterAuthorizedNetworksConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) MasterVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) MeshCertificates() DataGoogleContainerClusterMeshCertificatesList {
	var returns DataGoogleContainerClusterMeshCertificatesList
	_jsii_.Get(
		j,
		"meshCertificates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) MinMasterVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minMasterVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) MonitoringConfig() DataGoogleContainerClusterMonitoringConfigList {
	var returns DataGoogleContainerClusterMonitoringConfigList
	_jsii_.Get(
		j,
		"monitoringConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) MonitoringService() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) NetworkingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) NetworkPolicy() DataGoogleContainerClusterNetworkPolicyList {
	var returns DataGoogleContainerClusterNetworkPolicyList
	_jsii_.Get(
		j,
		"networkPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) NodeConfig() DataGoogleContainerClusterNodeConfigList {
	var returns DataGoogleContainerClusterNodeConfigList
	_jsii_.Get(
		j,
		"nodeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) NodeLocations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nodeLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) NodePool() DataGoogleContainerClusterNodePoolList {
	var returns DataGoogleContainerClusterNodePoolList
	_jsii_.Get(
		j,
		"nodePool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) NodePoolDefaults() DataGoogleContainerClusterNodePoolDefaultsList {
	var returns DataGoogleContainerClusterNodePoolDefaultsList
	_jsii_.Get(
		j,
		"nodePoolDefaults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) NodeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) NotificationConfig() DataGoogleContainerClusterNotificationConfigList {
	var returns DataGoogleContainerClusterNotificationConfigList
	_jsii_.Get(
		j,
		"notificationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) Operation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) PrivateClusterConfig() DataGoogleContainerClusterPrivateClusterConfigList {
	var returns DataGoogleContainerClusterPrivateClusterConfigList
	_jsii_.Get(
		j,
		"privateClusterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) PrivateIpv6GoogleAccess() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpv6GoogleAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) ReleaseChannel() DataGoogleContainerClusterReleaseChannelList {
	var returns DataGoogleContainerClusterReleaseChannelList
	_jsii_.Get(
		j,
		"releaseChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) RemoveDefaultNodePool() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"removeDefaultNodePool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) ResourceLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"resourceLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) ResourceUsageExportConfig() DataGoogleContainerClusterResourceUsageExportConfigList {
	var returns DataGoogleContainerClusterResourceUsageExportConfigList
	_jsii_.Get(
		j,
		"resourceUsageExportConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) ServiceExternalIpsConfig() DataGoogleContainerClusterServiceExternalIpsConfigList {
	var returns DataGoogleContainerClusterServiceExternalIpsConfigList
	_jsii_.Get(
		j,
		"serviceExternalIpsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) ServicesIpv4Cidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"servicesIpv4Cidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) Subnetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) TpuIpv4CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tpuIpv4CidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) VerticalPodAutoscaling() DataGoogleContainerClusterVerticalPodAutoscalingList {
	var returns DataGoogleContainerClusterVerticalPodAutoscalingList
	_jsii_.Get(
		j,
		"verticalPodAutoscaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleContainerCluster) WorkloadIdentityConfig() DataGoogleContainerClusterWorkloadIdentityConfigList {
	var returns DataGoogleContainerClusterWorkloadIdentityConfigList
	_jsii_.Get(
		j,
		"workloadIdentityConfig",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/4.71.0/docs/data-sources/container_cluster google_container_cluster} Data Source.
func NewDataGoogleContainerCluster(scope constructs.Construct, id *string, config *DataGoogleContainerClusterConfig) DataGoogleContainerCluster {
	_init_.Initialize()

	if err := validateNewDataGoogleContainerClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGoogleContainerCluster{}

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleContainerCluster.DataGoogleContainerCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/4.71.0/docs/data-sources/container_cluster google_container_cluster} Data Source.
func NewDataGoogleContainerCluster_Override(d DataGoogleContainerCluster, scope constructs.Construct, id *string, config *DataGoogleContainerClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataGoogleContainerCluster.DataGoogleContainerCluster",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataGoogleContainerCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataGoogleContainerCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataGoogleContainerCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataGoogleContainerCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataGoogleContainerCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataGoogleContainerCluster)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_DataGoogleContainerCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataGoogleContainerCluster)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_DataGoogleContainerCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
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
func DataGoogleContainerCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGoogleContainerCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dataGoogleContainerCluster.DataGoogleContainerCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataGoogleContainerCluster_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGoogleContainerCluster_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dataGoogleContainerCluster.DataGoogleContainerCluster",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataGoogleContainerCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGoogleContainerCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.dataGoogleContainerCluster.DataGoogleContainerCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataGoogleContainerCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.dataGoogleContainerCluster.DataGoogleContainerCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataGoogleContainerCluster) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataGoogleContainerCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataGoogleContainerCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGoogleContainerCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataGoogleContainerCluster) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataGoogleContainerCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataGoogleContainerCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataGoogleContainerCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataGoogleContainerCluster) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataGoogleContainerCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataGoogleContainerCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGoogleContainerCluster) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataGoogleContainerCluster) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGoogleContainerCluster) ResetLocation() {
	_jsii_.InvokeVoid(
		d,
		"resetLocation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGoogleContainerCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGoogleContainerCluster) ResetProject() {
	_jsii_.InvokeVoid(
		d,
		"resetProject",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGoogleContainerCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleContainerCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleContainerCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleContainerCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

