// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computesecuritypolicyrule


type ComputeSecurityPolicyRuleRateLimitOptionsExceedRedirectOptionsA struct {
	// Target for the redirect action. This is required if the type is EXTERNAL_302 and cannot be specified for GOOGLE_RECAPTCHA.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/compute_security_policy_rule#target ComputeSecurityPolicyRuleA#target}
	Target *string `field:"optional" json:"target" yaml:"target"`
	// Type of the redirect action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/compute_security_policy_rule#type ComputeSecurityPolicyRuleA#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

