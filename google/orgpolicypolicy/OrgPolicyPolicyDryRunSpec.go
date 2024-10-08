// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package orgpolicypolicy


type OrgPolicyPolicyDryRunSpec struct {
	// Determines the inheritance behavior for this policy.
	//
	// If 'inherit_from_parent' is true, policy rules set higher up in the hierarchy (up to the closest root) are inherited and present in the effective policy. If it is false, then no rules are inherited, and this policy becomes the new root for evaluation. This field can be set only for policies which configure list constraints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/org_policy_policy#inherit_from_parent OrgPolicyPolicy#inherit_from_parent}
	InheritFromParent interface{} `field:"optional" json:"inheritFromParent" yaml:"inheritFromParent"`
	// Ignores policies set above this resource and restores the 'constraint_default' enforcement behavior of the specific constraint at this resource.
	//
	// This field can be set in policies for either list or boolean constraints. If set, 'rules' must be empty and 'inherit_from_parent' must be set to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/org_policy_policy#reset OrgPolicyPolicy#reset}
	Reset interface{} `field:"optional" json:"reset" yaml:"reset"`
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/org_policy_policy#rules OrgPolicyPolicy#rules}
	Rules interface{} `field:"optional" json:"rules" yaml:"rules"`
}

