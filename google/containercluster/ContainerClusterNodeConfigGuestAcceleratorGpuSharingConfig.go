package containercluster


type ContainerClusterNodeConfigGuestAcceleratorGpuSharingConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/container_cluster#gpu_sharing_strategy ContainerCluster#gpu_sharing_strategy}.
	GpuSharingStrategy *string `field:"optional" json:"gpuSharingStrategy" yaml:"gpuSharingStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/container_cluster#max_shared_clients_per_gpu ContainerCluster#max_shared_clients_per_gpu}.
	MaxSharedClientsPerGpu *float64 `field:"optional" json:"maxSharedClientsPerGpu" yaml:"maxSharedClientsPerGpu"`
}

