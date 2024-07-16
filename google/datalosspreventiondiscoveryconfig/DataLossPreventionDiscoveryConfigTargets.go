// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigTargets struct {
	// big_query_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/data_loss_prevention_discovery_config#big_query_target DataLossPreventionDiscoveryConfig#big_query_target}
	BigQueryTarget *DataLossPreventionDiscoveryConfigTargetsBigQueryTarget `field:"optional" json:"bigQueryTarget" yaml:"bigQueryTarget"`
	// cloud_sql_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/data_loss_prevention_discovery_config#cloud_sql_target DataLossPreventionDiscoveryConfig#cloud_sql_target}
	CloudSqlTarget *DataLossPreventionDiscoveryConfigTargetsCloudSqlTarget `field:"optional" json:"cloudSqlTarget" yaml:"cloudSqlTarget"`
	// secrets_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.38.0/docs/resources/data_loss_prevention_discovery_config#secrets_target DataLossPreventionDiscoveryConfig#secrets_target}
	SecretsTarget *DataLossPreventionDiscoveryConfigTargetsSecretsTarget `field:"optional" json:"secretsTarget" yaml:"secretsTarget"`
}

