// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computesecuritypolicyrule


type ComputeSecurityPolicyRuleMatchExprOptionsA struct {
	// recaptcha_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/compute_security_policy_rule#recaptcha_options ComputeSecurityPolicyRuleA#recaptcha_options}
	RecaptchaOptions *ComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptionsA `field:"required" json:"recaptchaOptions" yaml:"recaptchaOptions"`
}

