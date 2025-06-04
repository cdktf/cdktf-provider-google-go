// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataplexentry


type DataplexEntryAspects struct {
	// aspect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/dataplex_entry#aspect DataplexEntry#aspect}
	Aspect *DataplexEntryAspectsAspect `field:"required" json:"aspect" yaml:"aspect"`
	// Depending on how the aspect is attached to the entry, the format of the aspect key can be one of the following:.
	//
	// If the aspect is attached directly to the entry: {project_number}.{locationId}.{aspectTypeId}
	// If the aspect is attached to an entry's path: {project_number}.{locationId}.{aspectTypeId}@{path}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/dataplex_entry#aspect_key DataplexEntry#aspect_key}
	AspectKey *string `field:"required" json:"aspectKey" yaml:"aspectKey"`
}

