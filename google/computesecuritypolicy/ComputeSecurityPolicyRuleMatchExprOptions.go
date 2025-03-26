// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computesecuritypolicy


type ComputeSecurityPolicyRuleMatchExprOptions struct {
	// recaptcha_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/compute_security_policy#recaptcha_options ComputeSecurityPolicy#recaptcha_options}
	RecaptchaOptions *ComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptions `field:"required" json:"recaptchaOptions" yaml:"recaptchaOptions"`
}

