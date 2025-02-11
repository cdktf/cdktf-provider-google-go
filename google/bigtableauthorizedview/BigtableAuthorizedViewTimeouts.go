// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigtableauthorizedview


type BigtableAuthorizedViewTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/bigtable_authorized_view#create BigtableAuthorizedView#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/bigtable_authorized_view#update BigtableAuthorizedView#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

