// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigtabletable


type BigtableTableColumnFamily struct {
	// The name of the column family.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/bigtable_table#family BigtableTable#family}
	Family *string `field:"required" json:"family" yaml:"family"`
}

