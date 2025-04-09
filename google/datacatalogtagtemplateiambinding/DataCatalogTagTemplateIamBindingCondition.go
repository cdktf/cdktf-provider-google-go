// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacatalogtagtemplateiambinding


type DataCatalogTagTemplateIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/data_catalog_tag_template_iam_binding#expression DataCatalogTagTemplateIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/data_catalog_tag_template_iam_binding#title DataCatalogTagTemplateIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/data_catalog_tag_template_iam_binding#description DataCatalogTagTemplateIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

