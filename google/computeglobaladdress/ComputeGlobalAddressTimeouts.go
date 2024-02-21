// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeglobaladdress


type ComputeGlobalAddressTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.17.0/docs/resources/compute_global_address#create ComputeGlobalAddress#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.17.0/docs/resources/compute_global_address#delete ComputeGlobalAddress#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

