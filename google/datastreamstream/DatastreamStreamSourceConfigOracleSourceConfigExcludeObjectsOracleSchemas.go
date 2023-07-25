package datastreamstream


type DatastreamStreamSourceConfigOracleSourceConfigExcludeObjectsOracleSchemas struct {
	// Schema name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.75.0/docs/resources/datastream_stream#schema DatastreamStream#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// oracle_tables block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.75.0/docs/resources/datastream_stream#oracle_tables DatastreamStream#oracle_tables}
	OracleTables interface{} `field:"optional" json:"oracleTables" yaml:"oracleTables"`
}

