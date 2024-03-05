// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterNodePoolNodeConfigEphemeralStorageLocalSsdConfig struct {
	// Number of local SSDs to use to back ephemeral storage.
	//
	// Uses NVMe interfaces. Each local SSD must be 375 or 3000 GB in size, and all local SSDs must share the same size.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.19.0/docs/resources/container_cluster#local_ssd_count ContainerCluster#local_ssd_count}
	LocalSsdCount *float64 `field:"required" json:"localSsdCount" yaml:"localSsdCount"`
}

