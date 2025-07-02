// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sccsourceiammember


type SccSourceIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/scc_source_iam_member#expression SccSourceIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/scc_source_iam_member#title SccSourceIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/scc_source_iam_member#description SccSourceIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

