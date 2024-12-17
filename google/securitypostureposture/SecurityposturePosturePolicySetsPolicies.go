// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitypostureposture


type SecurityposturePosturePolicySetsPolicies struct {
	// constraint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.0/docs/resources/securityposture_posture#constraint SecurityposturePosture#constraint}
	Constraint *SecurityposturePosturePolicySetsPoliciesConstraint `field:"required" json:"constraint" yaml:"constraint"`
	// ID of the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.0/docs/resources/securityposture_posture#policy_id SecurityposturePosture#policy_id}
	PolicyId *string `field:"required" json:"policyId" yaml:"policyId"`
	// compliance_standards block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.0/docs/resources/securityposture_posture#compliance_standards SecurityposturePosture#compliance_standards}
	ComplianceStandards interface{} `field:"optional" json:"complianceStandards" yaml:"complianceStandards"`
	// Description of the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.0/docs/resources/securityposture_posture#description SecurityposturePosture#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

