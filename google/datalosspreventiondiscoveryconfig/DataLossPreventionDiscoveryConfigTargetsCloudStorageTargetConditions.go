// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigTargetsCloudStorageTargetConditions struct {
	// cloud_storage_conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/data_loss_prevention_discovery_config#cloud_storage_conditions DataLossPreventionDiscoveryConfig#cloud_storage_conditions}
	CloudStorageConditions *DataLossPreventionDiscoveryConfigTargetsCloudStorageTargetConditionsCloudStorageConditions `field:"optional" json:"cloudStorageConditions" yaml:"cloudStorageConditions"`
	// File store must have been created after this date.
	//
	// Used to avoid backfilling. A timestamp in RFC3339 UTC "Zulu" format with nanosecond resolution and upto nine fractional digits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/data_loss_prevention_discovery_config#created_after DataLossPreventionDiscoveryConfig#created_after}
	CreatedAfter *string `field:"optional" json:"createdAfter" yaml:"createdAfter"`
	// Duration format. Minimum age a file store must have. If set, the value must be 1 hour or greater.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/data_loss_prevention_discovery_config#min_age DataLossPreventionDiscoveryConfig#min_age}
	MinAge *string `field:"optional" json:"minAge" yaml:"minAge"`
}

