// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigTargetsBigQueryTargetConditionsOrConditions struct {
	// Duration format. The minimum age a table must have before Cloud DLP can profile it. Value greater than 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/data_loss_prevention_discovery_config#min_age DataLossPreventionDiscoveryConfig#min_age}
	MinAge *string `field:"optional" json:"minAge" yaml:"minAge"`
	// Minimum number of rows that should be present before Cloud DLP profiles as a table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.35.0/docs/resources/data_loss_prevention_discovery_config#min_row_count DataLossPreventionDiscoveryConfig#min_row_count}
	MinRowCount *float64 `field:"optional" json:"minRowCount" yaml:"minRowCount"`
}

