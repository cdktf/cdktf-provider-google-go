// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bigquerydataset


type BigqueryDatasetExternalCatalogDatasetOptions struct {
	// The storage location URI for all tables in the dataset.
	//
	// Equivalent to hive metastore's
	// database locationUri. Maximum length of 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/bigquery_dataset#default_storage_location_uri BigqueryDataset#default_storage_location_uri}
	DefaultStorageLocationUri *string `field:"optional" json:"defaultStorageLocationUri" yaml:"defaultStorageLocationUri"`
	// A map of key value pairs defining the parameters and properties of the open source schema. Maximum size of 2Mib.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/bigquery_dataset#parameters BigqueryDataset#parameters}
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
}

