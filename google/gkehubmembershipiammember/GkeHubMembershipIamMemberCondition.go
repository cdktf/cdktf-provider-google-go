// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkehubmembershipiammember


type GkeHubMembershipIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/gke_hub_membership_iam_member#expression GkeHubMembershipIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/gke_hub_membership_iam_member#title GkeHubMembershipIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/gke_hub_membership_iam_member#description GkeHubMembershipIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

