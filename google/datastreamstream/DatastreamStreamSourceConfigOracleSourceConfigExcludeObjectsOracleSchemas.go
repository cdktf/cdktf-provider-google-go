package datastreamstream


type DatastreamStreamSourceConfigOracleSourceConfigExcludeObjectsOracleSchemas struct {
	// Schema name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/datastream_stream#schema DatastreamStream#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// oracle_tables block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/datastream_stream#oracle_tables DatastreamStream#oracle_tables}
	OracleTables interface{} `field:"optional" json:"oracleTables" yaml:"oracleTables"`
}
