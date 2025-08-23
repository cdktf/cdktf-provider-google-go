// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigTargetsCloudSqlTargetFilter struct {
	// collection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/data_loss_prevention_discovery_config#collection DataLossPreventionDiscoveryConfig#collection}
	Collection *DataLossPreventionDiscoveryConfigTargetsCloudSqlTargetFilterCollection `field:"optional" json:"collection" yaml:"collection"`
	// database_resource_reference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/data_loss_prevention_discovery_config#database_resource_reference DataLossPreventionDiscoveryConfig#database_resource_reference}
	DatabaseResourceReference *DataLossPreventionDiscoveryConfigTargetsCloudSqlTargetFilterDatabaseResourceReference `field:"optional" json:"databaseResourceReference" yaml:"databaseResourceReference"`
	// others block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/data_loss_prevention_discovery_config#others DataLossPreventionDiscoveryConfig#others}
	Others *DataLossPreventionDiscoveryConfigTargetsCloudSqlTargetFilterOthers `field:"optional" json:"others" yaml:"others"`
}

