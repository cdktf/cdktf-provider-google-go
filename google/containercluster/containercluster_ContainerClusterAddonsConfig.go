package containercluster


type ContainerClusterAddonsConfig struct {
	// cloudrun_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#cloudrun_config ContainerCluster#cloudrun_config}
	CloudrunConfig *ContainerClusterAddonsConfigCloudrunConfig `field:"optional" json:"cloudrunConfig" yaml:"cloudrunConfig"`
	// dns_cache_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#dns_cache_config ContainerCluster#dns_cache_config}
	DnsCacheConfig *ContainerClusterAddonsConfigDnsCacheConfig `field:"optional" json:"dnsCacheConfig" yaml:"dnsCacheConfig"`
	// gce_persistent_disk_csi_driver_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#gce_persistent_disk_csi_driver_config ContainerCluster#gce_persistent_disk_csi_driver_config}
	GcePersistentDiskCsiDriverConfig *ContainerClusterAddonsConfigGcePersistentDiskCsiDriverConfig `field:"optional" json:"gcePersistentDiskCsiDriverConfig" yaml:"gcePersistentDiskCsiDriverConfig"`
	// gcp_filestore_csi_driver_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#gcp_filestore_csi_driver_config ContainerCluster#gcp_filestore_csi_driver_config}
	GcpFilestoreCsiDriverConfig *ContainerClusterAddonsConfigGcpFilestoreCsiDriverConfig `field:"optional" json:"gcpFilestoreCsiDriverConfig" yaml:"gcpFilestoreCsiDriverConfig"`
	// horizontal_pod_autoscaling block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#horizontal_pod_autoscaling ContainerCluster#horizontal_pod_autoscaling}
	HorizontalPodAutoscaling *ContainerClusterAddonsConfigHorizontalPodAutoscaling `field:"optional" json:"horizontalPodAutoscaling" yaml:"horizontalPodAutoscaling"`
	// http_load_balancing block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#http_load_balancing ContainerCluster#http_load_balancing}
	HttpLoadBalancing *ContainerClusterAddonsConfigHttpLoadBalancing `field:"optional" json:"httpLoadBalancing" yaml:"httpLoadBalancing"`
	// network_policy_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#network_policy_config ContainerCluster#network_policy_config}
	NetworkPolicyConfig *ContainerClusterAddonsConfigNetworkPolicyConfig `field:"optional" json:"networkPolicyConfig" yaml:"networkPolicyConfig"`
}

