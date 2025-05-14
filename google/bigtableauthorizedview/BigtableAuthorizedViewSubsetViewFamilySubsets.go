// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigtableauthorizedview


type BigtableAuthorizedViewSubsetViewFamilySubsets struct {
	// Name of the column family to be included in the authorized view.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/bigtable_authorized_view#family_name BigtableAuthorizedView#family_name}
	FamilyName *string `field:"required" json:"familyName" yaml:"familyName"`
	// Base64-encoded prefixes for qualifiers of the column family to be included in the authorized view.
	//
	// Every qualifier starting with one of these prefixes is included in the authorized view. To provide access to all qualifiers, include the empty string as a prefix ("").
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/bigtable_authorized_view#qualifier_prefixes BigtableAuthorizedView#qualifier_prefixes}
	QualifierPrefixes *[]*string `field:"optional" json:"qualifierPrefixes" yaml:"qualifierPrefixes"`
	// Base64-encoded individual exact column qualifiers of the column family to be included in the authorized view.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/bigtable_authorized_view#qualifiers BigtableAuthorizedView#qualifiers}
	Qualifiers *[]*string `field:"optional" json:"qualifiers" yaml:"qualifiers"`
}

