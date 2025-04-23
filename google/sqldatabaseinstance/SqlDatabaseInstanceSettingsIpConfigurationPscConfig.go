// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sqldatabaseinstance


type SqlDatabaseInstanceSettingsIpConfigurationPscConfig struct {
	// List of consumer projects that are allow-listed for PSC connections to this instance.
	//
	// This instance can be connected to with PSC from any network in these projects. Each consumer project in this list may be represented by a project number (numeric) or by a project id (alphanumeric).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/sql_database_instance#allowed_consumer_projects SqlDatabaseInstance#allowed_consumer_projects}
	AllowedConsumerProjects *[]*string `field:"optional" json:"allowedConsumerProjects" yaml:"allowedConsumerProjects"`
	// psc_auto_connections block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/sql_database_instance#psc_auto_connections SqlDatabaseInstance#psc_auto_connections}
	PscAutoConnections interface{} `field:"optional" json:"pscAutoConnections" yaml:"pscAutoConnections"`
	// Whether PSC connectivity is enabled for this instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/sql_database_instance#psc_enabled SqlDatabaseInstance#psc_enabled}
	PscEnabled interface{} `field:"optional" json:"pscEnabled" yaml:"pscEnabled"`
}

