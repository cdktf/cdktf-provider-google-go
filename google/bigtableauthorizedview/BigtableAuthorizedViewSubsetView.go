// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigtableauthorizedview


type BigtableAuthorizedViewSubsetView struct {
	// family_subsets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.36.0/docs/resources/bigtable_authorized_view#family_subsets BigtableAuthorizedView#family_subsets}
	FamilySubsets interface{} `field:"optional" json:"familySubsets" yaml:"familySubsets"`
	// Base64-encoded row prefixes to be included in the authorized view.
	//
	// To provide access to all rows, include the empty string as a prefix ("").
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.36.0/docs/resources/bigtable_authorized_view#row_prefixes BigtableAuthorizedView#row_prefixes}
	RowPrefixes *[]*string `field:"optional" json:"rowPrefixes" yaml:"rowPrefixes"`
}

