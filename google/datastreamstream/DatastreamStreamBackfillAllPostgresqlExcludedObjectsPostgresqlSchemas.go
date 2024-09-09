// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datastreamstream


type DatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemas struct {
	// Database name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.1.0/docs/resources/datastream_stream#schema DatastreamStream#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// postgresql_tables block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.1.0/docs/resources/datastream_stream#postgresql_tables DatastreamStream#postgresql_tables}
	PostgresqlTables interface{} `field:"optional" json:"postgresqlTables" yaml:"postgresqlTables"`
}

