// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataplexaspecttypeiammember


type DataplexAspectTypeIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/dataplex_aspect_type_iam_member#expression DataplexAspectTypeIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/dataplex_aspect_type_iam_member#title DataplexAspectTypeIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/dataplex_aspect_type_iam_member#description DataplexAspectTypeIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

