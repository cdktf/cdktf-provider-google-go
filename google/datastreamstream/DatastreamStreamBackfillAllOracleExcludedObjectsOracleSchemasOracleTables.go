package datastreamstream


type DatastreamStreamBackfillAllOracleExcludedObjectsOracleSchemasOracleTables struct {
	// Table name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/datastream_stream#table DatastreamStream#table}
	Table *string `field:"required" json:"table" yaml:"table"`
	// oracle_columns block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/datastream_stream#oracle_columns DatastreamStream#oracle_columns}
	OracleColumns interface{} `field:"optional" json:"oracleColumns" yaml:"oracleColumns"`
}

