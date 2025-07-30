// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregionsecuritypolicy


type ComputeRegionSecurityPolicyRulesPreconfiguredWafConfig struct {
	// exclusion block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/compute_region_security_policy#exclusion ComputeRegionSecurityPolicy#exclusion}
	Exclusion interface{} `field:"optional" json:"exclusion" yaml:"exclusion"`
}

