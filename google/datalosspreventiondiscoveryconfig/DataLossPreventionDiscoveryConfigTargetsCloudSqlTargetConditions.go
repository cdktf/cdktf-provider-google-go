// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigTargetsCloudSqlTargetConditions struct {
	// Database engines that should be profiled. Optional. Defaults to ALL_SUPPORTED_DATABASE_ENGINES if unspecified. Possible values: ["ALL_SUPPORTED_DATABASE_ENGINES", "MYSQL", "POSTGRES"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/data_loss_prevention_discovery_config#database_engines DataLossPreventionDiscoveryConfig#database_engines}
	DatabaseEngines *[]*string `field:"optional" json:"databaseEngines" yaml:"databaseEngines"`
	// Data profiles will only be generated for the database resource types specified in this field.
	//
	// If not specified, defaults to [DATABASE_RESOURCE_TYPE_ALL_SUPPORTED_TYPES]. Possible values: ["DATABASE_RESOURCE_TYPE_ALL_SUPPORTED_TYPES", "DATABASE_RESOURCE_TYPE_TABLE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/data_loss_prevention_discovery_config#types DataLossPreventionDiscoveryConfig#types}
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

