// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregiontargethttpproxy


type ComputeRegionTargetHttpProxyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.18.0/docs/resources/compute_region_target_http_proxy#create ComputeRegionTargetHttpProxy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.18.0/docs/resources/compute_region_target_http_proxy#delete ComputeRegionTargetHttpProxy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.18.0/docs/resources/compute_region_target_http_proxy#update ComputeRegionTargetHttpProxy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

