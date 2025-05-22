// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterTables struct {
	// include_regexes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/data_loss_prevention_discovery_config#include_regexes DataLossPreventionDiscoveryConfig#include_regexes}
	IncludeRegexes *DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterTablesIncludeRegexes `field:"optional" json:"includeRegexes" yaml:"includeRegexes"`
}

