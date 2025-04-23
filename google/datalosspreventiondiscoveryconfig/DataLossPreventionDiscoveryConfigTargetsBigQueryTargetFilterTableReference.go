// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterTableReference struct {
	// Dataset ID of the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/data_loss_prevention_discovery_config#dataset_id DataLossPreventionDiscoveryConfig#dataset_id}
	DatasetId *string `field:"required" json:"datasetId" yaml:"datasetId"`
	// Name of the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/data_loss_prevention_discovery_config#table_id DataLossPreventionDiscoveryConfig#table_id}
	TableId *string `field:"required" json:"tableId" yaml:"tableId"`
}

