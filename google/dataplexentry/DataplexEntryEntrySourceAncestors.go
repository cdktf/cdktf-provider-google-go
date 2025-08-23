// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataplexentry


type DataplexEntryEntrySourceAncestors struct {
	// The name of the ancestor resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/dataplex_entry#name DataplexEntry#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The type of the ancestor resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/dataplex_entry#type DataplexEntry#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

