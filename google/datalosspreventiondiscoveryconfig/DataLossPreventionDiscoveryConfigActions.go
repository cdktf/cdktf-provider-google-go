// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigActions struct {
	// export_data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/data_loss_prevention_discovery_config#export_data DataLossPreventionDiscoveryConfig#export_data}
	ExportData *DataLossPreventionDiscoveryConfigActionsExportData `field:"optional" json:"exportData" yaml:"exportData"`
	// pub_sub_notification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/data_loss_prevention_discovery_config#pub_sub_notification DataLossPreventionDiscoveryConfig#pub_sub_notification}
	PubSubNotification *DataLossPreventionDiscoveryConfigActionsPubSubNotification `field:"optional" json:"pubSubNotification" yaml:"pubSubNotification"`
	// tag_resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.1/docs/resources/data_loss_prevention_discovery_config#tag_resources DataLossPreventionDiscoveryConfig#tag_resources}
	TagResources *DataLossPreventionDiscoveryConfigActionsTagResources `field:"optional" json:"tagResources" yaml:"tagResources"`
}

