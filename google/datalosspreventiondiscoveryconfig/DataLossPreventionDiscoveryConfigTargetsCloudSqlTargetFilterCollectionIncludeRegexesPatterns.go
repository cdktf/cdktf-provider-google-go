// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigTargetsCloudSqlTargetFilterCollectionIncludeRegexesPatterns struct {
	// Regex to test the database name against. If empty, all databases match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.1/docs/resources/data_loss_prevention_discovery_config#database_regex DataLossPreventionDiscoveryConfig#database_regex}
	DatabaseRegex *string `field:"optional" json:"databaseRegex" yaml:"databaseRegex"`
	// Regex to test the database resource's name against.
	//
	// An example of a database resource name is a table's name. Other database resource names like view names could be included in the future. If empty, all database resources match.'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.1/docs/resources/data_loss_prevention_discovery_config#database_resource_name_regex DataLossPreventionDiscoveryConfig#database_resource_name_regex}
	DatabaseResourceNameRegex *string `field:"optional" json:"databaseResourceNameRegex" yaml:"databaseResourceNameRegex"`
	// Regex to test the instance name against. If empty, all instances match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.1/docs/resources/data_loss_prevention_discovery_config#instance_regex DataLossPreventionDiscoveryConfig#instance_regex}
	InstanceRegex *string `field:"optional" json:"instanceRegex" yaml:"instanceRegex"`
	// For organizations, if unset, will match all projects. Has no effect for data profile configurations created within a project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.1/docs/resources/data_loss_prevention_discovery_config#project_id_regex DataLossPreventionDiscoveryConfig#project_id_regex}
	ProjectIdRegex *string `field:"optional" json:"projectIdRegex" yaml:"projectIdRegex"`
}

