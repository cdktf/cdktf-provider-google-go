// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilter struct {
	// other_tables block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/data_loss_prevention_discovery_config#other_tables DataLossPreventionDiscoveryConfig#other_tables}
	OtherTables *DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterOtherTables `field:"optional" json:"otherTables" yaml:"otherTables"`
	// table_reference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/data_loss_prevention_discovery_config#table_reference DataLossPreventionDiscoveryConfig#table_reference}
	TableReference *DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterTableReference `field:"optional" json:"tableReference" yaml:"tableReference"`
	// tables block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/data_loss_prevention_discovery_config#tables DataLossPreventionDiscoveryConfig#tables}
	Tables *DataLossPreventionDiscoveryConfigTargetsBigQueryTargetFilterTables `field:"optional" json:"tables" yaml:"tables"`
}

