// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeresourcepolicy


type ComputeResourcePolicyDiskConsistencyGroupPolicy struct {
	// Enable disk consistency on the resource policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_resource_policy#enabled ComputeResourcePolicy#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

