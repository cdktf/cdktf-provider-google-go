// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sqldatabaseinstance


type SqlDatabaseInstanceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.44.0/docs/resources/sql_database_instance#create SqlDatabaseInstance#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.44.0/docs/resources/sql_database_instance#delete SqlDatabaseInstance#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.44.0/docs/resources/sql_database_instance#update SqlDatabaseInstance#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

