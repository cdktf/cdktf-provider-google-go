// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigtablegcpolicy


type BigtableGcPolicyMaxVersion struct {
	// Number of version before applying the GC policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.16.0/docs/resources/bigtable_gc_policy#number BigtableGcPolicy#number}
	Number *float64 `field:"required" json:"number" yaml:"number"`
}

