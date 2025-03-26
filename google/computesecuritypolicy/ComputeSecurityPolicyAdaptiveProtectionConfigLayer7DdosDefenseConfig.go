// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computesecuritypolicy


type ComputeSecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfig struct {
	// If set to true, enables CAAP for L7 DDoS detection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/compute_security_policy#enable ComputeSecurityPolicy#enable}
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
	// Rule visibility. Supported values include: "STANDARD", "PREMIUM".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/compute_security_policy#rule_visibility ComputeSecurityPolicy#rule_visibility}
	RuleVisibility *string `field:"optional" json:"ruleVisibility" yaml:"ruleVisibility"`
	// threshold_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/compute_security_policy#threshold_configs ComputeSecurityPolicy#threshold_configs}
	ThresholdConfigs interface{} `field:"optional" json:"thresholdConfigs" yaml:"thresholdConfigs"`
}

