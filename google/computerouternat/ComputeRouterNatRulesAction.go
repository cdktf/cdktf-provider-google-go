// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computerouternat


type ComputeRouterNatRulesAction struct {
	// A list of URLs of the IP resources used for this NAT rule.
	//
	// These IP addresses must be valid static external IP addresses assigned to the project.
	// This field is used for public NAT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/compute_router_nat#source_nat_active_ips ComputeRouterNat#source_nat_active_ips}
	SourceNatActiveIps *[]*string `field:"optional" json:"sourceNatActiveIps" yaml:"sourceNatActiveIps"`
	// A list of URLs of the subnetworks used as source ranges for this NAT Rule.
	//
	// These subnetworks must have purpose set to PRIVATE_NAT.
	// This field is used for private NAT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/compute_router_nat#source_nat_active_ranges ComputeRouterNat#source_nat_active_ranges}
	SourceNatActiveRanges *[]*string `field:"optional" json:"sourceNatActiveRanges" yaml:"sourceNatActiveRanges"`
	// A list of URLs of the IP resources to be drained.
	//
	// These IPs must be valid static external IPs that have been assigned to the NAT.
	// These IPs should be used for updating/patching a NAT rule only.
	// This field is used for public NAT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/compute_router_nat#source_nat_drain_ips ComputeRouterNat#source_nat_drain_ips}
	SourceNatDrainIps *[]*string `field:"optional" json:"sourceNatDrainIps" yaml:"sourceNatDrainIps"`
	// A list of URLs of subnetworks representing source ranges to be drained.
	//
	// This is only supported on patch/update, and these subnetworks must have previously been used as active ranges in this NAT Rule.
	// This field is used for private NAT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/compute_router_nat#source_nat_drain_ranges ComputeRouterNat#source_nat_drain_ranges}
	SourceNatDrainRanges *[]*string `field:"optional" json:"sourceNatDrainRanges" yaml:"sourceNatDrainRanges"`
}

