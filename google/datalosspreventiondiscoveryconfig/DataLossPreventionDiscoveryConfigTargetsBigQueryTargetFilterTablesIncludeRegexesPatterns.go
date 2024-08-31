// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterTablesIncludeRegexesPatterns struct {
	// if unset, this property matches all datasets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/data_loss_prevention_discovery_config#dataset_id_regex DataLossPreventionDiscoveryConfig#dataset_id_regex}
	DatasetIdRegex *string `field:"optional" json:"datasetIdRegex" yaml:"datasetIdRegex"`
	// For organizations, if unset, will match all projects. Has no effect for data profile configurations created within a project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/data_loss_prevention_discovery_config#project_id_regex DataLossPreventionDiscoveryConfig#project_id_regex}
	ProjectIdRegex *string `field:"optional" json:"projectIdRegex" yaml:"projectIdRegex"`
	// if unset, this property matches all tables.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.1/docs/resources/data_loss_prevention_discovery_config#table_id_regex DataLossPreventionDiscoveryConfig#table_id_regex}
	TableIdRegex *string `field:"optional" json:"tableIdRegex" yaml:"tableIdRegex"`
}

