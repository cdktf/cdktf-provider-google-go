// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigTargetsCloudSqlTargetFilter struct {
	// collection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.30.0/docs/resources/data_loss_prevention_discovery_config#collection DataLossPreventionDiscoveryConfig#collection}
	Collection *DataLossPreventionDiscoveryConfigTargetsCloudSqlTargetFilterCollection `field:"optional" json:"collection" yaml:"collection"`
	// others block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.30.0/docs/resources/data_loss_prevention_discovery_config#others DataLossPreventionDiscoveryConfig#others}
	Others *DataLossPreventionDiscoveryConfigTargetsCloudSqlTargetFilterOthers `field:"optional" json:"others" yaml:"others"`
}

