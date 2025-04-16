// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iamdenypolicy


type IamDenyPolicyRules struct {
	// deny_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/iam_deny_policy#deny_rule IamDenyPolicy#deny_rule}
	DenyRule *IamDenyPolicyRulesDenyRule `field:"optional" json:"denyRule" yaml:"denyRule"`
	// The description of the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/iam_deny_policy#description IamDenyPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

