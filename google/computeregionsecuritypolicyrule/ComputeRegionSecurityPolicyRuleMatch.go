// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregionsecuritypolicyrule


type ComputeRegionSecurityPolicyRuleMatch struct {
	// config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/compute_region_security_policy_rule#config ComputeRegionSecurityPolicyRule#config}
	Config *ComputeRegionSecurityPolicyRuleMatchConfig `field:"optional" json:"config" yaml:"config"`
	// expr block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/compute_region_security_policy_rule#expr ComputeRegionSecurityPolicyRule#expr}
	Expr *ComputeRegionSecurityPolicyRuleMatchExpr `field:"optional" json:"expr" yaml:"expr"`
	// Preconfigured versioned expression.
	//
	// If this field is specified, config must also be specified.
	// Available preconfigured expressions along with their requirements are: SRC_IPS_V1 - must specify the corresponding srcIpRange field in config. Possible values: ["SRC_IPS_V1"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/compute_region_security_policy_rule#versioned_expr ComputeRegionSecurityPolicyRule#versioned_expr}
	VersionedExpr *string `field:"optional" json:"versionedExpr" yaml:"versionedExpr"`
}

