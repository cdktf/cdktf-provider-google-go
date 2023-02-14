package datastreamstream


type DatastreamStreamSourceConfigPostgresqlSourceConfigExcludeObjectsPostgresqlSchemas struct {
	// Database name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/datastream_stream#schema DatastreamStream#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// postgresql_tables block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/datastream_stream#postgresql_tables DatastreamStream#postgresql_tables}
	PostgresqlTables interface{} `field:"optional" json:"postgresqlTables" yaml:"postgresqlTables"`
}

