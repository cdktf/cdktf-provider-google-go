// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocmetastoreservice


type DataprocMetastoreServiceMetadataIntegrationDataCatalogConfig struct {
	// Defines whether the metastore metadata should be synced to Data Catalog.
	//
	// The default value is to disable syncing metastore metadata to Data Catalog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/dataproc_metastore_service#enabled DataprocMetastoreService#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

