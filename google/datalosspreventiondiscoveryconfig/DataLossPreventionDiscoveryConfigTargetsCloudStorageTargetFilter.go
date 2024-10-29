// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilter struct {
	// cloud_storage_resource_reference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/data_loss_prevention_discovery_config#cloud_storage_resource_reference DataLossPreventionDiscoveryConfig#cloud_storage_resource_reference}
	CloudStorageResourceReference *DataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterCloudStorageResourceReference `field:"optional" json:"cloudStorageResourceReference" yaml:"cloudStorageResourceReference"`
	// collection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/data_loss_prevention_discovery_config#collection DataLossPreventionDiscoveryConfig#collection}
	Collection *DataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterCollection `field:"optional" json:"collection" yaml:"collection"`
	// others block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/data_loss_prevention_discovery_config#others DataLossPreventionDiscoveryConfig#others}
	Others *DataLossPreventionDiscoveryConfigTargetsCloudStorageTargetFilterOthers `field:"optional" json:"others" yaml:"others"`
}

