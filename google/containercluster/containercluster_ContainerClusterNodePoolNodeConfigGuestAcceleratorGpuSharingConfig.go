package containercluster


type ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#gpu_sharing_strategy ContainerCluster#gpu_sharing_strategy}.
	GpuSharingStrategy *string `field:"optional" json:"gpuSharingStrategy" yaml:"gpuSharingStrategy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#max_shared_clients_per_gpu ContainerCluster#max_shared_clients_per_gpu}.
	MaxSharedClientsPerGpu *float64 `field:"optional" json:"maxSharedClientsPerGpu" yaml:"maxSharedClientsPerGpu"`
}

