// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package biglaketable


type BiglakeTableHiveOptions struct {
	// Stores user supplied Hive table parameters.
	//
	// An object containing a
	// list of "key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/biglake_table#parameters BiglakeTable#parameters}
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// storage_descriptor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/biglake_table#storage_descriptor BiglakeTable#storage_descriptor}
	StorageDescriptor *BiglakeTableHiveOptionsStorageDescriptor `field:"optional" json:"storageDescriptor" yaml:"storageDescriptor"`
	// Hive table type. For example, MANAGED_TABLE, EXTERNAL_TABLE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/biglake_table#table_type BiglakeTable#table_type}
	TableType *string `field:"optional" json:"tableType" yaml:"tableType"`
}

