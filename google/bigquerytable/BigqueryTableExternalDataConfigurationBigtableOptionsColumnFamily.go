// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigquerytable


type BigqueryTableExternalDataConfigurationBigtableOptionsColumnFamily struct {
	// column block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/bigquery_table#column BigqueryTable#column}
	Column interface{} `field:"optional" json:"column" yaml:"column"`
	// The encoding of the values when the type is not STRING.
	//
	// Acceptable encoding values are: TEXT - indicates values are alphanumeric text strings. BINARY - indicates values are encoded using HBase Bytes.toBytes family of functions. This can be overridden for a specific column by listing that column in 'columns' and specifying an encoding for it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/bigquery_table#encoding BigqueryTable#encoding}
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// Identifier of the column family.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/bigquery_table#family_id BigqueryTable#family_id}
	FamilyId *string `field:"optional" json:"familyId" yaml:"familyId"`
	// If this is set only the latest version of value are exposed for all columns in this column family.
	//
	// This can be overridden for a specific column by listing that column in 'columns' and specifying a different setting for that column.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/bigquery_table#only_read_latest BigqueryTable#only_read_latest}
	OnlyReadLatest interface{} `field:"optional" json:"onlyReadLatest" yaml:"onlyReadLatest"`
	// The type to convert the value in cells of this column family.
	//
	// The values are expected to be encoded using HBase Bytes.toBytes function when using the BINARY encoding value. Following BigQuery types are allowed (case-sensitive): "BYTES", "STRING", "INTEGER", "FLOAT", "BOOLEAN", "JSON". Default type is BYTES. This can be overridden for a specific column by listing that column in 'columns' and specifying a type for it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/bigquery_table#type BigqueryTable#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

