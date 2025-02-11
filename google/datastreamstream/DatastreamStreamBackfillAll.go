// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datastreamstream


type DatastreamStreamBackfillAll struct {
	// mysql_excluded_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/datastream_stream#mysql_excluded_objects DatastreamStream#mysql_excluded_objects}
	MysqlExcludedObjects *DatastreamStreamBackfillAllMysqlExcludedObjects `field:"optional" json:"mysqlExcludedObjects" yaml:"mysqlExcludedObjects"`
	// oracle_excluded_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/datastream_stream#oracle_excluded_objects DatastreamStream#oracle_excluded_objects}
	OracleExcludedObjects *DatastreamStreamBackfillAllOracleExcludedObjects `field:"optional" json:"oracleExcludedObjects" yaml:"oracleExcludedObjects"`
	// postgresql_excluded_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/datastream_stream#postgresql_excluded_objects DatastreamStream#postgresql_excluded_objects}
	PostgresqlExcludedObjects *DatastreamStreamBackfillAllPostgresqlExcludedObjects `field:"optional" json:"postgresqlExcludedObjects" yaml:"postgresqlExcludedObjects"`
	// sql_server_excluded_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/datastream_stream#sql_server_excluded_objects DatastreamStream#sql_server_excluded_objects}
	SqlServerExcludedObjects *DatastreamStreamBackfillAllSqlServerExcludedObjects `field:"optional" json:"sqlServerExcludedObjects" yaml:"sqlServerExcludedObjects"`
}

