// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitypostureposture


type SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraint struct {
	// Organization policy canned constraint Id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.19.0/docs/resources/securityposture_posture#canned_constraint_id SecurityposturePosture#canned_constraint_id}
	CannedConstraintId *string `field:"required" json:"cannedConstraintId" yaml:"cannedConstraintId"`
	// policy_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.19.0/docs/resources/securityposture_posture#policy_rules SecurityposturePosture#policy_rules}
	PolicyRules interface{} `field:"required" json:"policyRules" yaml:"policyRules"`
}

