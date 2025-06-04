// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datastreamstream


type DatastreamStreamSourceConfig struct {
	// Source connection profile resource. Format: projects/{project}/locations/{location}/connectionProfiles/{name}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/datastream_stream#source_connection_profile DatastreamStream#source_connection_profile}
	SourceConnectionProfile *string `field:"required" json:"sourceConnectionProfile" yaml:"sourceConnectionProfile"`
	// mysql_source_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/datastream_stream#mysql_source_config DatastreamStream#mysql_source_config}
	MysqlSourceConfig *DatastreamStreamSourceConfigMysqlSourceConfig `field:"optional" json:"mysqlSourceConfig" yaml:"mysqlSourceConfig"`
	// oracle_source_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/datastream_stream#oracle_source_config DatastreamStream#oracle_source_config}
	OracleSourceConfig *DatastreamStreamSourceConfigOracleSourceConfig `field:"optional" json:"oracleSourceConfig" yaml:"oracleSourceConfig"`
	// postgresql_source_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/datastream_stream#postgresql_source_config DatastreamStream#postgresql_source_config}
	PostgresqlSourceConfig *DatastreamStreamSourceConfigPostgresqlSourceConfig `field:"optional" json:"postgresqlSourceConfig" yaml:"postgresqlSourceConfig"`
	// salesforce_source_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/datastream_stream#salesforce_source_config DatastreamStream#salesforce_source_config}
	SalesforceSourceConfig *DatastreamStreamSourceConfigSalesforceSourceConfig `field:"optional" json:"salesforceSourceConfig" yaml:"salesforceSourceConfig"`
	// sql_server_source_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/datastream_stream#sql_server_source_config DatastreamStream#sql_server_source_config}
	SqlServerSourceConfig *DatastreamStreamSourceConfigSqlServerSourceConfig `field:"optional" json:"sqlServerSourceConfig" yaml:"sqlServerSourceConfig"`
}

