// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregiontargettcpproxy


type ComputeRegionTargetTcpProxyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.16.0/docs/resources/compute_region_target_tcp_proxy#create ComputeRegionTargetTcpProxy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.16.0/docs/resources/compute_region_target_tcp_proxy#delete ComputeRegionTargetTcpProxy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

