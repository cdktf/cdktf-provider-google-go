// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitypostureposture


type SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustom struct {
	// policy_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.18.0/docs/resources/securityposture_posture#policy_rules SecurityposturePosture#policy_rules}
	PolicyRules interface{} `field:"required" json:"policyRules" yaml:"policyRules"`
	// custom_constraint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.18.0/docs/resources/securityposture_posture#custom_constraint SecurityposturePosture#custom_constraint}
	CustomConstraint *SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomCustomConstraint `field:"optional" json:"customConstraint" yaml:"customConstraint"`
}

