// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitypostureposture


type SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRules struct {
	// Setting this to true means that all values are allowed.
	//
	// This field can be set only in policies for list constraints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/securityposture_posture#allow_all SecurityposturePosture#allow_all}
	AllowAll interface{} `field:"optional" json:"allowAll" yaml:"allowAll"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/securityposture_posture#condition SecurityposturePosture#condition}
	Condition *SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesCondition `field:"optional" json:"condition" yaml:"condition"`
	// Setting this to true means that all values are denied.
	//
	// This field can be set only in policies for list constraints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/securityposture_posture#deny_all SecurityposturePosture#deny_all}
	DenyAll interface{} `field:"optional" json:"denyAll" yaml:"denyAll"`
	// If 'true', then the policy is enforced.
	//
	// If 'false', then any configuration is acceptable.
	// This field can be set only in policies for boolean constraints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/securityposture_posture#enforce SecurityposturePosture#enforce}
	Enforce interface{} `field:"optional" json:"enforce" yaml:"enforce"`
	// values block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.0/docs/resources/securityposture_posture#values SecurityposturePosture#values}
	Values *SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustomPolicyRulesValues `field:"optional" json:"values" yaml:"values"`
}

