// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computesecuritypolicy


type ComputeSecurityPolicyRuleRateLimitOptionsEnforceOnKeyConfigs struct {
	// Rate limit key name applicable only for the following key types: HTTP_HEADER -- Name of the HTTP header whose value is taken as the key value.
	//
	// HTTP_COOKIE -- Name of the HTTP cookie whose value is taken as the key value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/compute_security_policy#enforce_on_key_name ComputeSecurityPolicy#enforce_on_key_name}
	EnforceOnKeyName *string `field:"optional" json:"enforceOnKeyName" yaml:"enforceOnKeyName"`
	// Determines the key to enforce the rate_limit_threshold on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/compute_security_policy#enforce_on_key_type ComputeSecurityPolicy#enforce_on_key_type}
	EnforceOnKeyType *string `field:"optional" json:"enforceOnKeyType" yaml:"enforceOnKeyType"`
}

