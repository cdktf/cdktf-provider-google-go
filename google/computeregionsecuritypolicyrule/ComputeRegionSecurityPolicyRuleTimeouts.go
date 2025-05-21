// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregionsecuritypolicyrule


type ComputeRegionSecurityPolicyRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/compute_region_security_policy_rule#create ComputeRegionSecurityPolicyRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/compute_region_security_policy_rule#delete ComputeRegionSecurityPolicyRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/compute_region_security_policy_rule#update ComputeRegionSecurityPolicyRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

