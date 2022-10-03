package containercluster


type ContainerClusterNodePoolNodeConfigGuestAccelerator struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#count ContainerCluster#count}.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#gpu_partition_size ContainerCluster#gpu_partition_size}.
	GpuPartitionSize *string `field:"optional" json:"gpuPartitionSize" yaml:"gpuPartitionSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#type ContainerCluster#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

