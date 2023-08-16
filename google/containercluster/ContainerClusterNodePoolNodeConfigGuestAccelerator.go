package containercluster


type ContainerClusterNodePoolNodeConfigGuestAccelerator struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_cluster#count ContainerCluster#count}.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_cluster#gpu_driver_installation_config ContainerCluster#gpu_driver_installation_config}.
	GpuDriverInstallationConfig interface{} `field:"optional" json:"gpuDriverInstallationConfig" yaml:"gpuDriverInstallationConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_cluster#gpu_partition_size ContainerCluster#gpu_partition_size}.
	GpuPartitionSize *string `field:"optional" json:"gpuPartitionSize" yaml:"gpuPartitionSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_cluster#gpu_sharing_config ContainerCluster#gpu_sharing_config}.
	GpuSharingConfig interface{} `field:"optional" json:"gpuSharingConfig" yaml:"gpuSharingConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_cluster#type ContainerCluster#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

