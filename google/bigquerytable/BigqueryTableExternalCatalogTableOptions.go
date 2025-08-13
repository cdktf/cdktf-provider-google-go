// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigquerytable


type BigqueryTableExternalCatalogTableOptions struct {
	// The connection specifying the credentials to be used to read external storage, such as Azure Blob, Cloud Storage, or S3.
	//
	// The connection is needed to read the open source table from BigQuery Engine. The connection_id can have the form <project_id>.<location_id>.<connection_id> or projects/<project_id>/locations/<location_id>/connections/<connection_id>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/bigquery_table#connection_id BigqueryTable#connection_id}
	ConnectionId *string `field:"optional" json:"connectionId" yaml:"connectionId"`
	// A map of key value pairs defining the parameters and properties of the open source table.
	//
	// Corresponds with hive meta store table parameters. Maximum size of 4Mib.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/bigquery_table#parameters BigqueryTable#parameters}
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// storage_descriptor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/bigquery_table#storage_descriptor BigqueryTable#storage_descriptor}
	StorageDescriptor *BigqueryTableExternalCatalogTableOptionsStorageDescriptor `field:"optional" json:"storageDescriptor" yaml:"storageDescriptor"`
}

