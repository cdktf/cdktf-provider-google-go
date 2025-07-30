// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacatalogtaxonomy


type DataCatalogTaxonomyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/data_catalog_taxonomy#create DataCatalogTaxonomy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/data_catalog_taxonomy#delete DataCatalogTaxonomy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/data_catalog_taxonomy#update DataCatalogTaxonomy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

