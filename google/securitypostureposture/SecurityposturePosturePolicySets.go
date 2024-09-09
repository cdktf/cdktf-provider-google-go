// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitypostureposture


type SecurityposturePosturePolicySets struct {
	// policies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.1.0/docs/resources/securityposture_posture#policies SecurityposturePosture#policies}
	Policies interface{} `field:"required" json:"policies" yaml:"policies"`
	// ID of the policy set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.1.0/docs/resources/securityposture_posture#policy_set_id SecurityposturePosture#policy_set_id}
	PolicySetId *string `field:"required" json:"policySetId" yaml:"policySetId"`
	// Description of the policy set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.1.0/docs/resources/securityposture_posture#description SecurityposturePosture#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

