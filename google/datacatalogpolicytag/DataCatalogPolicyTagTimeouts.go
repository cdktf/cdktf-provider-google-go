// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datacatalogpolicytag


type DataCatalogPolicyTagTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/data_catalog_policy_tag#create DataCatalogPolicyTag#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/data_catalog_policy_tag#delete DataCatalogPolicyTag#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/data_catalog_policy_tag#update DataCatalogPolicyTag#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

