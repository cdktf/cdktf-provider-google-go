// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computenodetemplate


type ComputeNodeTemplateDisks struct {
	// Specifies the number of such disks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.0/docs/resources/compute_node_template#disk_count ComputeNodeTemplate#disk_count}
	DiskCount *float64 `field:"optional" json:"diskCount" yaml:"diskCount"`
	// Specifies the size of the disk in base-2 GB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.0/docs/resources/compute_node_template#disk_size_gb ComputeNodeTemplate#disk_size_gb}
	DiskSizeGb *float64 `field:"optional" json:"diskSizeGb" yaml:"diskSizeGb"`
	// Specifies the desired disk type on the node.
	//
	// This disk type must be a local storage type (e.g.: local-ssd). Note that for nodeTemplates, this should be the name of the disk type and not its URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.0/docs/resources/compute_node_template#disk_type ComputeNodeTemplate#disk_type}
	DiskType *string `field:"optional" json:"diskType" yaml:"diskType"`
}

