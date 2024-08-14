// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigquerytable


type BigqueryTableExternalDataConfigurationBigtableOptions struct {
	// column_family block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.41.0/docs/resources/bigquery_table#column_family BigqueryTable#column_family}
	ColumnFamily interface{} `field:"optional" json:"columnFamily" yaml:"columnFamily"`
	// If field is true, then the column families that are not specified in columnFamilies list are not exposed in the table schema.
	//
	// Otherwise, they are read with BYTES type values. The default value is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.41.0/docs/resources/bigquery_table#ignore_unspecified_column_families BigqueryTable#ignore_unspecified_column_families}
	IgnoreUnspecifiedColumnFamilies interface{} `field:"optional" json:"ignoreUnspecifiedColumnFamilies" yaml:"ignoreUnspecifiedColumnFamilies"`
	// If field is true, then each column family will be read as a single JSON column.
	//
	// Otherwise they are read as a repeated cell structure containing timestamp/value tuples. The default value is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.41.0/docs/resources/bigquery_table#output_column_families_as_json BigqueryTable#output_column_families_as_json}
	OutputColumnFamiliesAsJson interface{} `field:"optional" json:"outputColumnFamiliesAsJson" yaml:"outputColumnFamiliesAsJson"`
	// If field is true, then the rowkey column families will be read and converted to string.
	//
	// Otherwise they are read with BYTES type values and users need to manually cast them with CAST if necessary. The default value is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.41.0/docs/resources/bigquery_table#read_rowkey_as_string BigqueryTable#read_rowkey_as_string}
	ReadRowkeyAsString interface{} `field:"optional" json:"readRowkeyAsString" yaml:"readRowkeyAsString"`
}

