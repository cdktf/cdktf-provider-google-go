// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computesecuritypolicyrule


type ComputeSecurityPolicyRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.0/docs/resources/compute_security_policy_rule#create ComputeSecurityPolicyRuleA#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.0/docs/resources/compute_security_policy_rule#delete ComputeSecurityPolicyRuleA#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.0/docs/resources/compute_security_policy_rule#update ComputeSecurityPolicyRuleA#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

