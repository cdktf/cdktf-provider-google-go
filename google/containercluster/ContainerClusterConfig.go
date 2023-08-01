package containercluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerClusterConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The name of the cluster, unique within the project and location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#name ContainerCluster#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// addons_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#addons_config ContainerCluster#addons_config}
	AddonsConfig *ContainerClusterAddonsConfig `field:"optional" json:"addonsConfig" yaml:"addonsConfig"`
	// Enable NET_ADMIN for this cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#allow_net_admin ContainerCluster#allow_net_admin}
	AllowNetAdmin interface{} `field:"optional" json:"allowNetAdmin" yaml:"allowNetAdmin"`
	// authenticator_groups_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#authenticator_groups_config ContainerCluster#authenticator_groups_config}
	AuthenticatorGroupsConfig *ContainerClusterAuthenticatorGroupsConfig `field:"optional" json:"authenticatorGroupsConfig" yaml:"authenticatorGroupsConfig"`
	// binary_authorization block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#binary_authorization ContainerCluster#binary_authorization}
	BinaryAuthorization *ContainerClusterBinaryAuthorization `field:"optional" json:"binaryAuthorization" yaml:"binaryAuthorization"`
	// cluster_autoscaling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#cluster_autoscaling ContainerCluster#cluster_autoscaling}
	ClusterAutoscaling *ContainerClusterClusterAutoscaling `field:"optional" json:"clusterAutoscaling" yaml:"clusterAutoscaling"`
	// The IP address range of the Kubernetes pods in this cluster in CIDR notation (e.g. 10.96.0.0/14). Leave blank to have one automatically chosen or specify a /14 block in 10.0.0.0/8. This field will only work for routes-based clusters, where ip_allocation_policy is not defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#cluster_ipv4_cidr ContainerCluster#cluster_ipv4_cidr}
	ClusterIpv4Cidr *string `field:"optional" json:"clusterIpv4Cidr" yaml:"clusterIpv4Cidr"`
	// confidential_nodes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#confidential_nodes ContainerCluster#confidential_nodes}
	ConfidentialNodes *ContainerClusterConfidentialNodes `field:"optional" json:"confidentialNodes" yaml:"confidentialNodes"`
	// cost_management_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#cost_management_config ContainerCluster#cost_management_config}
	CostManagementConfig *ContainerClusterCostManagementConfig `field:"optional" json:"costManagementConfig" yaml:"costManagementConfig"`
	// database_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#database_encryption ContainerCluster#database_encryption}
	DatabaseEncryption *ContainerClusterDatabaseEncryption `field:"optional" json:"databaseEncryption" yaml:"databaseEncryption"`
	// The desired datapath provider for this cluster. By default, uses the IPTables-based kube-proxy implementation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#datapath_provider ContainerCluster#datapath_provider}
	DatapathProvider *string `field:"optional" json:"datapathProvider" yaml:"datapathProvider"`
	// The default maximum number of pods per node in this cluster.
	//
	// This doesn't work on "routes-based" clusters, clusters that don't have IP Aliasing enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#default_max_pods_per_node ContainerCluster#default_max_pods_per_node}
	DefaultMaxPodsPerNode *float64 `field:"optional" json:"defaultMaxPodsPerNode" yaml:"defaultMaxPodsPerNode"`
	// default_snat_status block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#default_snat_status ContainerCluster#default_snat_status}
	DefaultSnatStatus *ContainerClusterDefaultSnatStatus `field:"optional" json:"defaultSnatStatus" yaml:"defaultSnatStatus"`
	// Description of the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#description ContainerCluster#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// dns_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#dns_config ContainerCluster#dns_config}
	DnsConfig *ContainerClusterDnsConfig `field:"optional" json:"dnsConfig" yaml:"dnsConfig"`
	// Enable Autopilot for this cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#enable_autopilot ContainerCluster#enable_autopilot}
	EnableAutopilot interface{} `field:"optional" json:"enableAutopilot" yaml:"enableAutopilot"`
	// Enable Binary Authorization for this cluster. If enabled, all container images will be validated by Google Binary Authorization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#enable_binary_authorization ContainerCluster#enable_binary_authorization}
	EnableBinaryAuthorization interface{} `field:"optional" json:"enableBinaryAuthorization" yaml:"enableBinaryAuthorization"`
	// Whether Intra-node visibility is enabled for this cluster.
	//
	// This makes same node pod to pod traffic visible for VPC network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#enable_intranode_visibility ContainerCluster#enable_intranode_visibility}
	EnableIntranodeVisibility interface{} `field:"optional" json:"enableIntranodeVisibility" yaml:"enableIntranodeVisibility"`
	// Whether to enable Kubernetes Alpha features for this cluster.
	//
	// Note that when this option is enabled, the cluster cannot be upgraded and will be automatically deleted after 30 days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#enable_kubernetes_alpha ContainerCluster#enable_kubernetes_alpha}
	EnableKubernetesAlpha interface{} `field:"optional" json:"enableKubernetesAlpha" yaml:"enableKubernetesAlpha"`
	// Whether L4ILB Subsetting is enabled for this cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#enable_l4_ilb_subsetting ContainerCluster#enable_l4_ilb_subsetting}
	EnableL4IlbSubsetting interface{} `field:"optional" json:"enableL4IlbSubsetting" yaml:"enableL4IlbSubsetting"`
	// Whether the ABAC authorizer is enabled for this cluster.
	//
	// When enabled, identities in the system, including service accounts, nodes, and controllers, will have statically granted permissions beyond those provided by the RBAC configuration or IAM. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#enable_legacy_abac ContainerCluster#enable_legacy_abac}
	EnableLegacyAbac interface{} `field:"optional" json:"enableLegacyAbac" yaml:"enableLegacyAbac"`
	// Enable Shielded Nodes features on all nodes in this cluster. Defaults to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#enable_shielded_nodes ContainerCluster#enable_shielded_nodes}
	EnableShieldedNodes interface{} `field:"optional" json:"enableShieldedNodes" yaml:"enableShieldedNodes"`
	// Whether to enable Cloud TPU resources in this cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#enable_tpu ContainerCluster#enable_tpu}
	EnableTpu interface{} `field:"optional" json:"enableTpu" yaml:"enableTpu"`
	// gateway_api_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#gateway_api_config ContainerCluster#gateway_api_config}
	GatewayApiConfig *ContainerClusterGatewayApiConfig `field:"optional" json:"gatewayApiConfig" yaml:"gatewayApiConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#id ContainerCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The number of nodes to create in this cluster's default node pool.
	//
	// In regional or multi-zonal clusters, this is the number of nodes per zone. Must be set if node_pool is not set. If you're using google_container_node_pool objects with no default node pool, you'll need to set this to a value of at least 1, alongside setting remove_default_node_pool to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#initial_node_count ContainerCluster#initial_node_count}
	InitialNodeCount *float64 `field:"optional" json:"initialNodeCount" yaml:"initialNodeCount"`
	// ip_allocation_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#ip_allocation_policy ContainerCluster#ip_allocation_policy}
	IpAllocationPolicy *ContainerClusterIpAllocationPolicy `field:"optional" json:"ipAllocationPolicy" yaml:"ipAllocationPolicy"`
	// The location (region or zone) in which the cluster master will be created, as well as the default node location.
	//
	// If you specify a zone (such as us-central1-a), the cluster will be a zonal cluster with a single cluster master. If you specify a region (such as us-west1), the cluster will be a regional cluster with multiple masters spread across zones in the region, and with default node locations in those zones as well.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#location ContainerCluster#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// logging_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#logging_config ContainerCluster#logging_config}
	LoggingConfig *ContainerClusterLoggingConfig `field:"optional" json:"loggingConfig" yaml:"loggingConfig"`
	// The logging service that the cluster should write logs to.
	//
	// Available options include logging.googleapis.com(Legacy Stackdriver), logging.googleapis.com/kubernetes(Stackdriver Kubernetes Engine Logging), and none. Defaults to logging.googleapis.com/kubernetes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#logging_service ContainerCluster#logging_service}
	LoggingService *string `field:"optional" json:"loggingService" yaml:"loggingService"`
	// maintenance_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#maintenance_policy ContainerCluster#maintenance_policy}
	MaintenancePolicy *ContainerClusterMaintenancePolicy `field:"optional" json:"maintenancePolicy" yaml:"maintenancePolicy"`
	// master_auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#master_auth ContainerCluster#master_auth}
	MasterAuth *ContainerClusterMasterAuth `field:"optional" json:"masterAuth" yaml:"masterAuth"`
	// master_authorized_networks_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#master_authorized_networks_config ContainerCluster#master_authorized_networks_config}
	MasterAuthorizedNetworksConfig *ContainerClusterMasterAuthorizedNetworksConfig `field:"optional" json:"masterAuthorizedNetworksConfig" yaml:"masterAuthorizedNetworksConfig"`
	// mesh_certificates block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#mesh_certificates ContainerCluster#mesh_certificates}
	MeshCertificates *ContainerClusterMeshCertificates `field:"optional" json:"meshCertificates" yaml:"meshCertificates"`
	// The minimum version of the master.
	//
	// GKE will auto-update the master to new versions, so this does not guarantee the current master version--use the read-only master_version field to obtain that. If unset, the cluster's version will be set by GKE to the version of the most recent official release (which is not necessarily the latest version).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#min_master_version ContainerCluster#min_master_version}
	MinMasterVersion *string `field:"optional" json:"minMasterVersion" yaml:"minMasterVersion"`
	// monitoring_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#monitoring_config ContainerCluster#monitoring_config}
	MonitoringConfig *ContainerClusterMonitoringConfig `field:"optional" json:"monitoringConfig" yaml:"monitoringConfig"`
	// The monitoring service that the cluster should write metrics to.
	//
	// Automatically send metrics from pods in the cluster to the Google Cloud Monitoring API. VM metrics will be collected by Google Compute Engine regardless of this setting Available options include monitoring.googleapis.com(Legacy Stackdriver), monitoring.googleapis.com/kubernetes(Stackdriver Kubernetes Engine Monitoring), and none. Defaults to monitoring.googleapis.com/kubernetes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#monitoring_service ContainerCluster#monitoring_service}
	MonitoringService *string `field:"optional" json:"monitoringService" yaml:"monitoringService"`
	// The name or self_link of the Google Compute Engine network to which the cluster is connected.
	//
	// For Shared VPC, set this to the self link of the shared network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#network ContainerCluster#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// Determines whether alias IPs or routes will be used for pod IPs in the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#networking_mode ContainerCluster#networking_mode}
	NetworkingMode *string `field:"optional" json:"networkingMode" yaml:"networkingMode"`
	// network_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#network_policy ContainerCluster#network_policy}
	NetworkPolicy *ContainerClusterNetworkPolicy `field:"optional" json:"networkPolicy" yaml:"networkPolicy"`
	// node_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#node_config ContainerCluster#node_config}
	NodeConfig *ContainerClusterNodeConfig `field:"optional" json:"nodeConfig" yaml:"nodeConfig"`
	// The list of zones in which the cluster's nodes are located.
	//
	// Nodes must be in the region of their regional cluster or in the same region as their cluster's zone for zonal clusters. If this is specified for a zonal cluster, omit the cluster's zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#node_locations ContainerCluster#node_locations}
	NodeLocations *[]*string `field:"optional" json:"nodeLocations" yaml:"nodeLocations"`
	// node_pool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#node_pool ContainerCluster#node_pool}
	NodePool interface{} `field:"optional" json:"nodePool" yaml:"nodePool"`
	// node_pool_defaults block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#node_pool_defaults ContainerCluster#node_pool_defaults}
	NodePoolDefaults *ContainerClusterNodePoolDefaults `field:"optional" json:"nodePoolDefaults" yaml:"nodePoolDefaults"`
	// The Kubernetes version on the nodes.
	//
	// Must either be unset or set to the same value as min_master_version on create. Defaults to the default version set by GKE which is not necessarily the latest version. This only affects nodes in the default node pool. While a fuzzy version can be specified, it's recommended that you specify explicit versions as Terraform will see spurious diffs when fuzzy versions are used. See the google_container_engine_versions data source's version_prefix field to approximate fuzzy versions in a Terraform-compatible way. To update nodes in other node pools, use the version attribute on the node pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#node_version ContainerCluster#node_version}
	NodeVersion *string `field:"optional" json:"nodeVersion" yaml:"nodeVersion"`
	// notification_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#notification_config ContainerCluster#notification_config}
	NotificationConfig *ContainerClusterNotificationConfig `field:"optional" json:"notificationConfig" yaml:"notificationConfig"`
	// private_cluster_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#private_cluster_config ContainerCluster#private_cluster_config}
	PrivateClusterConfig *ContainerClusterPrivateClusterConfig `field:"optional" json:"privateClusterConfig" yaml:"privateClusterConfig"`
	// The desired state of IPv6 connectivity to Google Services.
	//
	// By default, no private IPv6 access to or from Google Services (all access will be via IPv4).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#private_ipv6_google_access ContainerCluster#private_ipv6_google_access}
	PrivateIpv6GoogleAccess *string `field:"optional" json:"privateIpv6GoogleAccess" yaml:"privateIpv6GoogleAccess"`
	// The ID of the project in which the resource belongs.
	//
	// If it is not provided, the provider project is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#project ContainerCluster#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// release_channel block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#release_channel ContainerCluster#release_channel}
	ReleaseChannel *ContainerClusterReleaseChannel `field:"optional" json:"releaseChannel" yaml:"releaseChannel"`
	// If true, deletes the default node pool upon cluster creation.
	//
	// If you're using google_container_node_pool resources with no default node pool, this should be set to true, alongside setting initial_node_count to at least 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#remove_default_node_pool ContainerCluster#remove_default_node_pool}
	RemoveDefaultNodePool interface{} `field:"optional" json:"removeDefaultNodePool" yaml:"removeDefaultNodePool"`
	// The GCE resource labels (a map of key/value pairs) to be applied to the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#resource_labels ContainerCluster#resource_labels}
	ResourceLabels *map[string]*string `field:"optional" json:"resourceLabels" yaml:"resourceLabels"`
	// resource_usage_export_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#resource_usage_export_config ContainerCluster#resource_usage_export_config}
	ResourceUsageExportConfig *ContainerClusterResourceUsageExportConfig `field:"optional" json:"resourceUsageExportConfig" yaml:"resourceUsageExportConfig"`
	// security_posture_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#security_posture_config ContainerCluster#security_posture_config}
	SecurityPostureConfig *ContainerClusterSecurityPostureConfig `field:"optional" json:"securityPostureConfig" yaml:"securityPostureConfig"`
	// service_external_ips_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#service_external_ips_config ContainerCluster#service_external_ips_config}
	ServiceExternalIpsConfig *ContainerClusterServiceExternalIpsConfig `field:"optional" json:"serviceExternalIpsConfig" yaml:"serviceExternalIpsConfig"`
	// The name or self_link of the Google Compute Engine subnetwork in which the cluster's instances are launched.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#subnetwork ContainerCluster#subnetwork}
	Subnetwork *string `field:"optional" json:"subnetwork" yaml:"subnetwork"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#timeouts ContainerCluster#timeouts}
	Timeouts *ContainerClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// vertical_pod_autoscaling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#vertical_pod_autoscaling ContainerCluster#vertical_pod_autoscaling}
	VerticalPodAutoscaling *ContainerClusterVerticalPodAutoscaling `field:"optional" json:"verticalPodAutoscaling" yaml:"verticalPodAutoscaling"`
	// workload_identity_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/container_cluster#workload_identity_config ContainerCluster#workload_identity_config}
	WorkloadIdentityConfig *ContainerClusterWorkloadIdentityConfig `field:"optional" json:"workloadIdentityConfig" yaml:"workloadIdentityConfig"`
}

