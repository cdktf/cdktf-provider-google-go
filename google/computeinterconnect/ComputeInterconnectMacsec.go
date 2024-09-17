// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeinterconnect


type ComputeInterconnectMacsec struct {
	// pre_shared_keys block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.3.0/docs/resources/compute_interconnect#pre_shared_keys ComputeInterconnect#pre_shared_keys}
	PreSharedKeys interface{} `field:"required" json:"preSharedKeys" yaml:"preSharedKeys"`
}

