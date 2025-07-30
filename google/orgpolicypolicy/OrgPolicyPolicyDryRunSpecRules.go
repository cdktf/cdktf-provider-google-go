// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package orgpolicypolicy


type OrgPolicyPolicyDryRunSpecRules struct {
	// Setting this to '"TRUE"' means that all values are allowed.
	//
	// This field can be set only in Policies for list constraints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/org_policy_policy#allow_all OrgPolicyPolicy#allow_all}
	AllowAll *string `field:"optional" json:"allowAll" yaml:"allowAll"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/org_policy_policy#condition OrgPolicyPolicy#condition}
	Condition *OrgPolicyPolicyDryRunSpecRulesCondition `field:"optional" json:"condition" yaml:"condition"`
	// Setting this to '"TRUE"' means that all values are denied.
	//
	// This field can be set only in Policies for list constraints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/org_policy_policy#deny_all OrgPolicyPolicy#deny_all}
	DenyAll *string `field:"optional" json:"denyAll" yaml:"denyAll"`
	// If '"TRUE"', then the 'Policy' is enforced.
	//
	// If '"FALSE"', then any configuration is acceptable. This field can be set only in Policies for boolean constraints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/org_policy_policy#enforce OrgPolicyPolicy#enforce}
	Enforce *string `field:"optional" json:"enforce" yaml:"enforce"`
	// Optional.
	//
	// Required for Managed Constraints if parameters defined in constraints. Pass parameter values when policy enforcement is enabled. Ensure that parameter value types match those defined in the constraint definition. For example: { \"allowedLocations\" : [\"us-east1\", \"us-west1\"], \"allowAll\" : true }
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/org_policy_policy#parameters OrgPolicyPolicy#parameters}
	Parameters *string `field:"optional" json:"parameters" yaml:"parameters"`
	// values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/org_policy_policy#values OrgPolicyPolicy#values}
	Values *OrgPolicyPolicyDryRunSpecRulesValues `field:"optional" json:"values" yaml:"values"`
}

