// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iamaccessboundarypolicy


type IamAccessBoundaryPolicyRules struct {
	// access_boundary_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.44.0/docs/resources/iam_access_boundary_policy#access_boundary_rule IamAccessBoundaryPolicy#access_boundary_rule}
	AccessBoundaryRule *IamAccessBoundaryPolicyRulesAccessBoundaryRule `field:"optional" json:"accessBoundaryRule" yaml:"accessBoundaryRule"`
	// The description of the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.44.0/docs/resources/iam_access_boundary_policy#description IamAccessBoundaryPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

