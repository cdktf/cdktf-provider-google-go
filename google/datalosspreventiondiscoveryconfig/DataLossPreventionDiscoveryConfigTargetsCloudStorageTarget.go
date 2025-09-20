// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigTargetsCloudStorageTarget struct {
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/data_loss_prevention_discovery_config#filter DataLossPreventionDiscoveryConfig#filter}
	Filter *DataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilter `field:"required" json:"filter" yaml:"filter"`
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/data_loss_prevention_discovery_config#conditions DataLossPreventionDiscoveryConfig#conditions}
	Conditions *DataLossPreventionDiscoveryConfigTargetsCloudStorageTargetConditions `field:"optional" json:"conditions" yaml:"conditions"`
	// disabled block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/data_loss_prevention_discovery_config#disabled DataLossPreventionDiscoveryConfig#disabled}
	Disabled *DataLossPreventionDiscoveryConfigTargetsCloudStorageTargetDisabled `field:"optional" json:"disabled" yaml:"disabled"`
	// generation_cadence block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/data_loss_prevention_discovery_config#generation_cadence DataLossPreventionDiscoveryConfig#generation_cadence}
	GenerationCadence *DataLossPreventionDiscoveryConfigTargetsCloudStorageTargetGenerationCadence `field:"optional" json:"generationCadence" yaml:"generationCadence"`
}

