package datastreamstream


type DatastreamStreamBackfillAllPostgresqlExcludedObjects struct {
	// postgresql_schemas block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/datastream_stream#postgresql_schemas DatastreamStream#postgresql_schemas}
	PostgresqlSchemas interface{} `field:"required" json:"postgresqlSchemas" yaml:"postgresqlSchemas"`
}

