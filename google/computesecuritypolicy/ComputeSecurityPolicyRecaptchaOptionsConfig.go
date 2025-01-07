// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computesecuritypolicy


type ComputeSecurityPolicyRecaptchaOptionsConfig struct {
	// A field to supply a reCAPTCHA site key to be used for all the rules using the redirect action with the type of GOOGLE_RECAPTCHA under the security policy.
	//
	// The specified site key needs to be created from the reCAPTCHA API. The user is responsible for the validity of the specified site key. If not specified, a Google-managed site key is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/compute_security_policy#redirect_site_key ComputeSecurityPolicy#redirect_site_key}
	RedirectSiteKey *string `field:"required" json:"redirectSiteKey" yaml:"redirectSiteKey"`
}

