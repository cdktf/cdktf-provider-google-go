// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigTargetsCloudSqlTargetGenerationCadenceInspectTemplateModifiedCadence struct {
	// How frequently data profiles can be updated when the template is modified.
	//
	// Defaults to never. Possible values: ["UPDATE_FREQUENCY_NEVER", "UPDATE_FREQUENCY_DAILY", "UPDATE_FREQUENCY_MONTHLY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/data_loss_prevention_discovery_config#frequency DataLossPreventionDiscoveryConfig#frequency}
	Frequency *string `field:"required" json:"frequency" yaml:"frequency"`
}

