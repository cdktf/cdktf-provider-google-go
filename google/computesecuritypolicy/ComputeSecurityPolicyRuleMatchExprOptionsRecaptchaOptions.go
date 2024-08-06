// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computesecuritypolicy


type ComputeSecurityPolicyRuleMatchExprOptionsRecaptchaOptions struct {
	// A list of site keys to be used during the validation of reCAPTCHA action-tokens.
	//
	// The provided site keys need to be created from reCAPTCHA API under the same project where the security policy is created
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.40.0/docs/resources/compute_security_policy#action_token_site_keys ComputeSecurityPolicy#action_token_site_keys}
	ActionTokenSiteKeys *[]*string `field:"optional" json:"actionTokenSiteKeys" yaml:"actionTokenSiteKeys"`
	// A list of site keys to be used during the validation of reCAPTCHA session-tokens.
	//
	// The provided site keys need to be created from reCAPTCHA API under the same project where the security policy is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.40.0/docs/resources/compute_security_policy#session_token_site_keys ComputeSecurityPolicy#session_token_site_keys}
	SessionTokenSiteKeys *[]*string `field:"optional" json:"sessionTokenSiteKeys" yaml:"sessionTokenSiteKeys"`
}

