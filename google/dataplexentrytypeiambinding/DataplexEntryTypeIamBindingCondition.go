// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataplexentrytypeiambinding


type DataplexEntryTypeIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/dataplex_entry_type_iam_binding#expression DataplexEntryTypeIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/dataplex_entry_type_iam_binding#title DataplexEntryTypeIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/dataplex_entry_type_iam_binding#description DataplexEntryTypeIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

