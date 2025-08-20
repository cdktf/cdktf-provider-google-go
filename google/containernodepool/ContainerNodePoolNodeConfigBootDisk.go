// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containernodepool


type ContainerNodePoolNodeConfigBootDisk struct {
	// Type of the disk attached to each node. Such as pd-standard, pd-balanced or pd-ssd.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_node_pool#disk_type ContainerNodePool#disk_type}
	DiskType *string `field:"optional" json:"diskType" yaml:"diskType"`
	// Configured IOPs provisioning. Only valid with disk type hyperdisk-balanced.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_node_pool#provisioned_iops ContainerNodePool#provisioned_iops}
	ProvisionedIops *float64 `field:"optional" json:"provisionedIops" yaml:"provisionedIops"`
	// Configured throughput provisioning. Only valid with disk type hyperdisk-balanced.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_node_pool#provisioned_throughput ContainerNodePool#provisioned_throughput}
	ProvisionedThroughput *float64 `field:"optional" json:"provisionedThroughput" yaml:"provisionedThroughput"`
	// Size of the disk attached to each node, specified in GB. The smallest allowed disk size is 10GB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_node_pool#size_gb ContainerNodePool#size_gb}
	SizeGb *float64 `field:"optional" json:"sizeGb" yaml:"sizeGb"`
}

