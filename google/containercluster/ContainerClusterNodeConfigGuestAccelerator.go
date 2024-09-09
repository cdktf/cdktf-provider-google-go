// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterNodeConfigGuestAccelerator struct {
	// The number of the accelerator cards exposed to an instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.1.0/docs/resources/container_cluster#count ContainerCluster#count}
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// The accelerator type resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.1.0/docs/resources/container_cluster#type ContainerCluster#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// gpu_driver_installation_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.1.0/docs/resources/container_cluster#gpu_driver_installation_config ContainerCluster#gpu_driver_installation_config}
	GpuDriverInstallationConfig *ContainerClusterNodeConfigGuestAcceleratorGpuDriverInstallationConfig `field:"optional" json:"gpuDriverInstallationConfig" yaml:"gpuDriverInstallationConfig"`
	// Size of partitions to create on the GPU. Valid values are described in the NVIDIA mig user guide (https://docs.nvidia.com/datacenter/tesla/mig-user-guide/#partitioning).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.1.0/docs/resources/container_cluster#gpu_partition_size ContainerCluster#gpu_partition_size}
	GpuPartitionSize *string `field:"optional" json:"gpuPartitionSize" yaml:"gpuPartitionSize"`
	// gpu_sharing_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.1.0/docs/resources/container_cluster#gpu_sharing_config ContainerCluster#gpu_sharing_config}
	GpuSharingConfig *ContainerClusterNodeConfigGuestAcceleratorGpuSharingConfig `field:"optional" json:"gpuSharingConfig" yaml:"gpuSharingConfig"`
}

