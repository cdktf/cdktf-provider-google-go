// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitypostureposture


type SecurityposturePosturePolicySetsPoliciesConstraint struct {
	// org_policy_constraint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/securityposture_posture#org_policy_constraint SecurityposturePosture#org_policy_constraint}
	OrgPolicyConstraint *SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraint `field:"optional" json:"orgPolicyConstraint" yaml:"orgPolicyConstraint"`
	// org_policy_constraint_custom block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/securityposture_posture#org_policy_constraint_custom SecurityposturePosture#org_policy_constraint_custom}
	OrgPolicyConstraintCustom *SecurityposturePosturePolicySetsPoliciesConstraintOrgPolicyConstraintCustom `field:"optional" json:"orgPolicyConstraintCustom" yaml:"orgPolicyConstraintCustom"`
	// security_health_analytics_custom_module block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/securityposture_posture#security_health_analytics_custom_module SecurityposturePosture#security_health_analytics_custom_module}
	SecurityHealthAnalyticsCustomModule *SecurityposturePosturePolicySetsPoliciesConstraintSecurityHealthAnalyticsCustomModule `field:"optional" json:"securityHealthAnalyticsCustomModule" yaml:"securityHealthAnalyticsCustomModule"`
	// security_health_analytics_module block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/securityposture_posture#security_health_analytics_module SecurityposturePosture#security_health_analytics_module}
	SecurityHealthAnalyticsModule *SecurityposturePosturePolicySetsPoliciesConstraintSecurityHealthAnalyticsModule `field:"optional" json:"securityHealthAnalyticsModule" yaml:"securityHealthAnalyticsModule"`
}

