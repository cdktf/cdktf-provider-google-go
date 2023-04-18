package datastreamstream


type DatastreamStreamSourceConfigPostgresqlSourceConfigIncludeObjectsPostgresqlSchemas struct {
	// Database name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.62.0/docs/resources/datastream_stream#schema DatastreamStream#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// postgresql_tables block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.62.0/docs/resources/datastream_stream#postgresql_tables DatastreamStream#postgresql_tables}
	PostgresqlTables interface{} `field:"optional" json:"postgresqlTables" yaml:"postgresqlTables"`
}

