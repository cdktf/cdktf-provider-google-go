// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computesecuritypolicy


type ComputeSecurityPolicyRulePreconfiguredWafConfigExclusion struct {
	// Target WAF rule set to apply the preconfigured WAF exclusion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/compute_security_policy#target_rule_set ComputeSecurityPolicy#target_rule_set}
	TargetRuleSet *string `field:"required" json:"targetRuleSet" yaml:"targetRuleSet"`
	// request_cookie block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/compute_security_policy#request_cookie ComputeSecurityPolicy#request_cookie}
	RequestCookie interface{} `field:"optional" json:"requestCookie" yaml:"requestCookie"`
	// request_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/compute_security_policy#request_header ComputeSecurityPolicy#request_header}
	RequestHeader interface{} `field:"optional" json:"requestHeader" yaml:"requestHeader"`
	// request_query_param block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/compute_security_policy#request_query_param ComputeSecurityPolicy#request_query_param}
	RequestQueryParam interface{} `field:"optional" json:"requestQueryParam" yaml:"requestQueryParam"`
	// request_uri block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/compute_security_policy#request_uri ComputeSecurityPolicy#request_uri}
	RequestUri interface{} `field:"optional" json:"requestUri" yaml:"requestUri"`
	// A list of target rule IDs under the WAF rule set to apply the preconfigured WAF exclusion.
	//
	// If omitted, it refers to all the rule IDs under the WAF rule set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/compute_security_policy#target_rule_ids ComputeSecurityPolicy#target_rule_ids}
	TargetRuleIds *[]*string `field:"optional" json:"targetRuleIds" yaml:"targetRuleIds"`
}

