// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataplexassetiambinding


type DataplexAssetIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/dataplex_asset_iam_binding#expression DataplexAssetIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/dataplex_asset_iam_binding#title DataplexAssetIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/dataplex_asset_iam_binding#description DataplexAssetIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

