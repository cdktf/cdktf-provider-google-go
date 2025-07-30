// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sqldatabase


type SqlDatabaseTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/sql_database#create SqlDatabase#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/sql_database#delete SqlDatabase#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/sql_database#update SqlDatabase#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

