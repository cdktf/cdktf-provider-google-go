// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitypostureposture


type SecurityposturePosturePolicySetsPoliciesComplianceStandards struct {
	// Mapping of security controls for the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/securityposture_posture#control SecurityposturePosture#control}
	Control *string `field:"optional" json:"control" yaml:"control"`
	// Mapping of compliance standards for the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/securityposture_posture#standard SecurityposturePosture#standard}
	Standard *string `field:"optional" json:"standard" yaml:"standard"`
}

