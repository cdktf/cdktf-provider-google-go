// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computesecuritypolicyrule


type ComputeSecurityPolicyRuleMatchExprA struct {
	// Textual representation of an expression in Common Expression Language syntax.
	//
	// The application context of the containing message determines which well-known feature set of CEL is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.2.0/docs/resources/compute_security_policy_rule#expression ComputeSecurityPolicyRuleA#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
}

