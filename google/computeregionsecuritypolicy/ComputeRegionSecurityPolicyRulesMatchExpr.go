// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregionsecuritypolicy


type ComputeRegionSecurityPolicyRulesMatchExpr struct {
	// Textual representation of an expression in Common Expression Language syntax.
	//
	// The application context of the containing message determines which well-known feature set of CEL is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/compute_region_security_policy#expression ComputeRegionSecurityPolicy#expression}
	Expression *string `field:"required" json:"expression" yaml:"expression"`
}

