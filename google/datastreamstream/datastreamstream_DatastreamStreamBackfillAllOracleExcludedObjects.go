package datastreamstream


type DatastreamStreamBackfillAllOracleExcludedObjects struct {
	// oracle_schemas block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/datastream_stream#oracle_schemas DatastreamStream#oracle_schemas}
	OracleSchemas interface{} `field:"required" json:"oracleSchemas" yaml:"oracleSchemas"`
}

