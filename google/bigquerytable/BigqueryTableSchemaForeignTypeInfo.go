// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigquerytable


type BigqueryTableSchemaForeignTypeInfo struct {
	// Specifies the system which defines the foreign data type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/bigquery_table#type_system BigqueryTable#type_system}
	TypeSystem *string `field:"required" json:"typeSystem" yaml:"typeSystem"`
}

