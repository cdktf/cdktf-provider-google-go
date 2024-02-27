// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computedisk


type ComputeDiskTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.18.0/docs/resources/compute_disk#create ComputeDisk#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.18.0/docs/resources/compute_disk#delete ComputeDisk#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.18.0/docs/resources/compute_disk#update ComputeDisk#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

