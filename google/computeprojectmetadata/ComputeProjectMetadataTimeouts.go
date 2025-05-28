// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeprojectmetadata


type ComputeProjectMetadataTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/compute_project_metadata#create ComputeProjectMetadata#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/compute_project_metadata#delete ComputeProjectMetadata#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

