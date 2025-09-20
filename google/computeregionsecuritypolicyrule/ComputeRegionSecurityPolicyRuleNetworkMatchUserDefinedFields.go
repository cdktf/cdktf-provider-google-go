// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregionsecuritypolicyrule


type ComputeRegionSecurityPolicyRuleNetworkMatchUserDefinedFields struct {
	// Name of the user-defined field, as given in the definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/compute_region_security_policy_rule#name ComputeRegionSecurityPolicyRule#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Matching values of the field.
	//
	// Each element can be a 32-bit unsigned decimal or hexadecimal (starting with "0x") number (e.g. "64") or range (e.g. "0x400-0x7ff").
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/compute_region_security_policy_rule#values ComputeRegionSecurityPolicyRule#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

