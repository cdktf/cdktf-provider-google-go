// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigOrgConfig struct {
	// location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.0/docs/resources/data_loss_prevention_discovery_config#location DataLossPreventionDiscoveryConfig#location}
	Location *DataLossPreventionDiscoveryConfigOrgConfigLocation `field:"optional" json:"location" yaml:"location"`
	// The project that will run the scan.
	//
	// The DLP service account that exists within this project must have access to all resources that are profiled, and the cloud DLP API must be enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.43.0/docs/resources/data_loss_prevention_discovery_config#project_id DataLossPreventionDiscoveryConfig#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

