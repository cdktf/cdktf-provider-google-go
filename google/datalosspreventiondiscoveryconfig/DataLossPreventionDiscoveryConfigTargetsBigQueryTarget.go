// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigTargetsBigQueryTarget struct {
	// cadence block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/data_loss_prevention_discovery_config#cadence DataLossPreventionDiscoveryConfig#cadence}
	Cadence *DataLossPreventionDiscoveryConfigTargetsBigQueryTargetCadence `field:"optional" json:"cadence" yaml:"cadence"`
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/data_loss_prevention_discovery_config#conditions DataLossPreventionDiscoveryConfig#conditions}
	Conditions *DataLossPreventionDiscoveryConfigTargetsBigQueryTargetConditions `field:"optional" json:"conditions" yaml:"conditions"`
	// disabled block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/data_loss_prevention_discovery_config#disabled DataLossPreventionDiscoveryConfig#disabled}
	Disabled *DataLossPreventionDiscoveryConfigTargetsBigQueryTargetDisabled `field:"optional" json:"disabled" yaml:"disabled"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/data_loss_prevention_discovery_config#filter DataLossPreventionDiscoveryConfig#filter}
	Filter *DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilter `field:"optional" json:"filter" yaml:"filter"`
}

