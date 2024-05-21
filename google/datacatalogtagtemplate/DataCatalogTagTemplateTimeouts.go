// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacatalogtagtemplate


type DataCatalogTagTemplateTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.30.0/docs/resources/data_catalog_tag_template#create DataCatalogTagTemplate#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.30.0/docs/resources/data_catalog_tag_template#delete DataCatalogTagTemplate#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.30.0/docs/resources/data_catalog_tag_template#update DataCatalogTagTemplate#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

