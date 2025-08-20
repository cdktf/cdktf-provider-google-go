// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterAddonsConfig struct {
	// cloudrun_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#cloudrun_config ContainerCluster#cloudrun_config}
	CloudrunConfig *ContainerClusterAddonsConfigCloudrunConfig `field:"optional" json:"cloudrunConfig" yaml:"cloudrunConfig"`
	// config_connector_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#config_connector_config ContainerCluster#config_connector_config}
	ConfigConnectorConfig *ContainerClusterAddonsConfigConfigConnectorConfig `field:"optional" json:"configConnectorConfig" yaml:"configConnectorConfig"`
	// dns_cache_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#dns_cache_config ContainerCluster#dns_cache_config}
	DnsCacheConfig *ContainerClusterAddonsConfigDnsCacheConfig `field:"optional" json:"dnsCacheConfig" yaml:"dnsCacheConfig"`
	// gce_persistent_disk_csi_driver_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#gce_persistent_disk_csi_driver_config ContainerCluster#gce_persistent_disk_csi_driver_config}
	GcePersistentDiskCsiDriverConfig *ContainerClusterAddonsConfigGcePersistentDiskCsiDriverConfig `field:"optional" json:"gcePersistentDiskCsiDriverConfig" yaml:"gcePersistentDiskCsiDriverConfig"`
	// gcp_filestore_csi_driver_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#gcp_filestore_csi_driver_config ContainerCluster#gcp_filestore_csi_driver_config}
	GcpFilestoreCsiDriverConfig *ContainerClusterAddonsConfigGcpFilestoreCsiDriverConfig `field:"optional" json:"gcpFilestoreCsiDriverConfig" yaml:"gcpFilestoreCsiDriverConfig"`
	// gcs_fuse_csi_driver_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#gcs_fuse_csi_driver_config ContainerCluster#gcs_fuse_csi_driver_config}
	GcsFuseCsiDriverConfig *ContainerClusterAddonsConfigGcsFuseCsiDriverConfig `field:"optional" json:"gcsFuseCsiDriverConfig" yaml:"gcsFuseCsiDriverConfig"`
	// gke_backup_agent_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#gke_backup_agent_config ContainerCluster#gke_backup_agent_config}
	GkeBackupAgentConfig *ContainerClusterAddonsConfigGkeBackupAgentConfig `field:"optional" json:"gkeBackupAgentConfig" yaml:"gkeBackupAgentConfig"`
	// horizontal_pod_autoscaling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#horizontal_pod_autoscaling ContainerCluster#horizontal_pod_autoscaling}
	HorizontalPodAutoscaling *ContainerClusterAddonsConfigHorizontalPodAutoscaling `field:"optional" json:"horizontalPodAutoscaling" yaml:"horizontalPodAutoscaling"`
	// http_load_balancing block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#http_load_balancing ContainerCluster#http_load_balancing}
	HttpLoadBalancing *ContainerClusterAddonsConfigHttpLoadBalancing `field:"optional" json:"httpLoadBalancing" yaml:"httpLoadBalancing"`
	// lustre_csi_driver_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#lustre_csi_driver_config ContainerCluster#lustre_csi_driver_config}
	LustreCsiDriverConfig *ContainerClusterAddonsConfigLustreCsiDriverConfig `field:"optional" json:"lustreCsiDriverConfig" yaml:"lustreCsiDriverConfig"`
	// network_policy_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#network_policy_config ContainerCluster#network_policy_config}
	NetworkPolicyConfig *ContainerClusterAddonsConfigNetworkPolicyConfig `field:"optional" json:"networkPolicyConfig" yaml:"networkPolicyConfig"`
	// parallelstore_csi_driver_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#parallelstore_csi_driver_config ContainerCluster#parallelstore_csi_driver_config}
	ParallelstoreCsiDriverConfig *ContainerClusterAddonsConfigParallelstoreCsiDriverConfig `field:"optional" json:"parallelstoreCsiDriverConfig" yaml:"parallelstoreCsiDriverConfig"`
	// ray_operator_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#ray_operator_config ContainerCluster#ray_operator_config}
	RayOperatorConfig interface{} `field:"optional" json:"rayOperatorConfig" yaml:"rayOperatorConfig"`
	// stateful_ha_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_cluster#stateful_ha_config ContainerCluster#stateful_ha_config}
	StatefulHaConfig *ContainerClusterAddonsConfigStatefulHaConfig `field:"optional" json:"statefulHaConfig" yaml:"statefulHaConfig"`
}

