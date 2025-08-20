// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigtablelogicalview


type BigtableLogicalViewTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/bigtable_logical_view#create BigtableLogicalView#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/bigtable_logical_view#delete BigtableLogicalView#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/bigtable_logical_view#update BigtableLogicalView#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

