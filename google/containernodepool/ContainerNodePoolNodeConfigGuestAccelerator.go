// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containernodepool


type ContainerNodePoolNodeConfigGuestAccelerator struct {
	// The number of the accelerator cards exposed to an instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/container_node_pool#count ContainerNodePool#count}
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// The accelerator type resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/container_node_pool#type ContainerNodePool#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// gpu_driver_installation_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/container_node_pool#gpu_driver_installation_config ContainerNodePool#gpu_driver_installation_config}
	GpuDriverInstallationConfig *ContainerNodePoolNodeConfigGuestAcceleratorGpuDriverInstallationConfig `field:"optional" json:"gpuDriverInstallationConfig" yaml:"gpuDriverInstallationConfig"`
	// Size of partitions to create on the GPU. Valid values are described in the NVIDIA mig user guide (https://docs.nvidia.com/datacenter/tesla/mig-user-guide/#partitioning).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/container_node_pool#gpu_partition_size ContainerNodePool#gpu_partition_size}
	GpuPartitionSize *string `field:"optional" json:"gpuPartitionSize" yaml:"gpuPartitionSize"`
	// gpu_sharing_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/container_node_pool#gpu_sharing_config ContainerNodePool#gpu_sharing_config}
	GpuSharingConfig *ContainerNodePoolNodeConfigGuestAcceleratorGpuSharingConfig `field:"optional" json:"gpuSharingConfig" yaml:"gpuSharingConfig"`
}

