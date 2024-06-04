// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataplexaspecttypeiambinding


type DataplexAspectTypeIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.32.0/docs/resources/dataplex_aspect_type_iam_binding#expression DataplexAspectTypeIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.32.0/docs/resources/dataplex_aspect_type_iam_binding#title DataplexAspectTypeIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.32.0/docs/resources/dataplex_aspect_type_iam_binding#description DataplexAspectTypeIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

