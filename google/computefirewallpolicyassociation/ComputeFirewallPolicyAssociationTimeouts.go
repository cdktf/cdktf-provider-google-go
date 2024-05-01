// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computefirewallpolicyassociation


type ComputeFirewallPolicyAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.27.0/docs/resources/compute_firewall_policy_association#create ComputeFirewallPolicyAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.27.0/docs/resources/compute_firewall_policy_association#delete ComputeFirewallPolicyAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

