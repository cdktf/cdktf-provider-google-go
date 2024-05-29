// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregionurlmap


type ComputeRegionUrlMapDefaultRouteActionFaultInjectionPolicy struct {
	// abort block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/compute_region_url_map#abort ComputeRegionUrlMap#abort}
	Abort *ComputeRegionUrlMapDefaultRouteActionFaultInjectionPolicyAbort `field:"optional" json:"abort" yaml:"abort"`
	// delay block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/compute_region_url_map#delay ComputeRegionUrlMap#delay}
	Delay *ComputeRegionUrlMapDefaultRouteActionFaultInjectionPolicyDelay `field:"optional" json:"delay" yaml:"delay"`
}

