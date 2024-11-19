// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package orgpolicypolicy


type OrgPolicyPolicySpecRules struct {
	// Setting this to '"TRUE"' means that all values are allowed.
	//
	// This field can be set only in Policies for list constraints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/org_policy_policy#allow_all OrgPolicyPolicy#allow_all}
	AllowAll *string `field:"optional" json:"allowAll" yaml:"allowAll"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/org_policy_policy#condition OrgPolicyPolicy#condition}
	Condition *OrgPolicyPolicySpecRulesCondition `field:"optional" json:"condition" yaml:"condition"`
	// Setting this to '"TRUE"' means that all values are denied.
	//
	// This field can be set only in Policies for list constraints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/org_policy_policy#deny_all OrgPolicyPolicy#deny_all}
	DenyAll *string `field:"optional" json:"denyAll" yaml:"denyAll"`
	// If '"TRUE"', then the 'Policy' is enforced.
	//
	// If '"FALSE"', then any configuration is acceptable. This field can be set only in Policies for boolean constraints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/org_policy_policy#enforce OrgPolicyPolicy#enforce}
	Enforce *string `field:"optional" json:"enforce" yaml:"enforce"`
	// values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/org_policy_policy#values OrgPolicyPolicy#values}
	Values *OrgPolicyPolicySpecRulesValues `field:"optional" json:"values" yaml:"values"`
}

