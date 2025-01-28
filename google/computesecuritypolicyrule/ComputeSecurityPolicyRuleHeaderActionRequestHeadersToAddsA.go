// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computesecuritypolicyrule


type ComputeSecurityPolicyRuleHeaderActionRequestHeadersToAddsA struct {
	// The name of the header to set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/compute_security_policy_rule#header_name ComputeSecurityPolicyRuleA#header_name}
	HeaderName *string `field:"optional" json:"headerName" yaml:"headerName"`
	// The value to set the named header to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/compute_security_policy_rule#header_value ComputeSecurityPolicyRuleA#header_value}
	HeaderValue *string `field:"optional" json:"headerValue" yaml:"headerValue"`
}

