// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterNodeConfigGuestAcceleratorGpuDriverInstallationConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.22.0/docs/resources/container_cluster#gpu_driver_version ContainerCluster#gpu_driver_version}.
	GpuDriverVersion *string `field:"optional" json:"gpuDriverVersion" yaml:"gpuDriverVersion"`
}

