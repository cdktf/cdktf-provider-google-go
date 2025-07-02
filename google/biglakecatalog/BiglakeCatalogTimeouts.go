// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package biglakecatalog


type BiglakeCatalogTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/biglake_catalog#create BiglakeCatalog#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/biglake_catalog#delete BiglakeCatalog#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

