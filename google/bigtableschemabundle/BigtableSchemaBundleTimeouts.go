// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigtableschemabundle


type BigtableSchemaBundleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/bigtable_schema_bundle#create BigtableSchemaBundle#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/bigtable_schema_bundle#delete BigtableSchemaBundle#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/bigtable_schema_bundle#update BigtableSchemaBundle#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

