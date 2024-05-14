// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigActionsExportDataProfileTable struct {
	// Dataset Id of the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.0/docs/resources/data_loss_prevention_discovery_config#dataset_id DataLossPreventionDiscoveryConfig#dataset_id}
	DatasetId *string `field:"optional" json:"datasetId" yaml:"datasetId"`
	// The Google Cloud Platform project ID of the project containing the table.
	//
	// If omitted, the project ID is inferred from the API call.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.0/docs/resources/data_loss_prevention_discovery_config#project_id DataLossPreventionDiscoveryConfig#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Name of the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.29.0/docs/resources/data_loss_prevention_discovery_config#table_id DataLossPreventionDiscoveryConfig#table_id}
	TableId *string `field:"optional" json:"tableId" yaml:"tableId"`
}

