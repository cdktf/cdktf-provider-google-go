// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterNodePoolNodeConfigBootDisk struct {
	// Type of the disk attached to each node. Such as pd-standard, pd-balanced or pd-ssd.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/container_cluster#disk_type ContainerCluster#disk_type}
	DiskType *string `field:"optional" json:"diskType" yaml:"diskType"`
	// Configured IOPs provisioning. Only valid with disk type hyperdisk-balanced.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/container_cluster#provisioned_iops ContainerCluster#provisioned_iops}
	ProvisionedIops *float64 `field:"optional" json:"provisionedIops" yaml:"provisionedIops"`
	// Configured throughput provisioning. Only valid with disk type hyperdisk-balanced.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/container_cluster#provisioned_throughput ContainerCluster#provisioned_throughput}
	ProvisionedThroughput *float64 `field:"optional" json:"provisionedThroughput" yaml:"provisionedThroughput"`
	// Size of the disk attached to each node, specified in GB. The smallest allowed disk size is 10GB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/container_cluster#size_gb ContainerCluster#size_gb}
	SizeGb *float64 `field:"optional" json:"sizeGb" yaml:"sizeGb"`
}

