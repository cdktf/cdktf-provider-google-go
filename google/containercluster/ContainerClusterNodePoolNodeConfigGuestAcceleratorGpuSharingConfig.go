// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterNodePoolNodeConfigGuestAcceleratorGpuSharingConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.1.0/docs/resources/container_cluster#gpu_sharing_strategy ContainerCluster#gpu_sharing_strategy}.
	GpuSharingStrategy *string `field:"optional" json:"gpuSharingStrategy" yaml:"gpuSharingStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.1.0/docs/resources/container_cluster#max_shared_clients_per_gpu ContainerCluster#max_shared_clients_per_gpu}.
	MaxSharedClientsPerGpu *float64 `field:"optional" json:"maxSharedClientsPerGpu" yaml:"maxSharedClientsPerGpu"`
}

