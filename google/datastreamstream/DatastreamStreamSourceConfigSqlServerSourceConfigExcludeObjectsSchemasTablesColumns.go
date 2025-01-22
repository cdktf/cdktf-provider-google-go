// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datastreamstream


type DatastreamStreamSourceConfigSqlServerSourceConfigExcludeObjectsSchemasTablesColumns struct {
	// Column name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/datastream_stream#column DatastreamStream#column}
	Column *string `field:"optional" json:"column" yaml:"column"`
	// The SQL Server data type. Full data types list can be found here: https://learn.microsoft.com/en-us/sql/t-sql/data-types/data-types-transact-sql?view=sql-server-ver16.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/datastream_stream#data_type DatastreamStream#data_type}
	DataType *string `field:"optional" json:"dataType" yaml:"dataType"`
}

