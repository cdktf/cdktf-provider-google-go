// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregionsecuritypolicy


type ComputeRegionSecurityPolicyRulesMatch struct {
	// config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/compute_region_security_policy#config ComputeRegionSecurityPolicy#config}
	Config *ComputeRegionSecurityPolicyRulesMatchConfig `field:"optional" json:"config" yaml:"config"`
	// expr block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/compute_region_security_policy#expr ComputeRegionSecurityPolicy#expr}
	Expr *ComputeRegionSecurityPolicyRulesMatchExpr `field:"optional" json:"expr" yaml:"expr"`
	// Preconfigured versioned expression.
	//
	// If this field is specified, config must also be specified.
	// Available preconfigured expressions along with their requirements are: SRC_IPS_V1 - must specify the corresponding srcIpRange field in config. Possible values: ["SRC_IPS_V1"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/compute_region_security_policy#versioned_expr ComputeRegionSecurityPolicy#versioned_expr}
	VersionedExpr *string `field:"optional" json:"versionedExpr" yaml:"versionedExpr"`
}

