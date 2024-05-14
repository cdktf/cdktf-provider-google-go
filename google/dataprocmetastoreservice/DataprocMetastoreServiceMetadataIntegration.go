// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprocmetastoreservice


type DataprocMetastoreServiceMetadataIntegration struct {
	// data_catalog_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.0/docs/resources/dataproc_metastore_service#data_catalog_config DataprocMetastoreService#data_catalog_config}
	DataCatalogConfig *DataprocMetastoreServiceMetadataIntegrationDataCatalogConfig `field:"required" json:"dataCatalogConfig" yaml:"dataCatalogConfig"`
}

