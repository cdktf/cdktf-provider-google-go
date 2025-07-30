// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigTargetsBigQueryTargetCadenceSchemaModifiedCadence struct {
	// How frequently profiles may be updated when schemas are modified. Default to monthly Possible values: ["UPDATE_FREQUENCY_NEVER", "UPDATE_FREQUENCY_DAILY", "UPDATE_FREQUENCY_MONTHLY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/data_loss_prevention_discovery_config#frequency DataLossPreventionDiscoveryConfig#frequency}
	Frequency *string `field:"optional" json:"frequency" yaml:"frequency"`
	// The type of events to consider when deciding if the table's schema has been modified and should have the profile updated.
	//
	// Defaults to NEW_COLUMN. Possible values: ["SCHEMA_NEW_COLUMNS", "SCHEMA_REMOVED_COLUMNS"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/data_loss_prevention_discovery_config#types DataLossPreventionDiscoveryConfig#types}
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

