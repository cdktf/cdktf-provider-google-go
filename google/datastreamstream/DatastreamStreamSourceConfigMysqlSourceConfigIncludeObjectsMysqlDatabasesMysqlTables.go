package datastreamstream


type DatastreamStreamSourceConfigMysqlSourceConfigIncludeObjectsMysqlDatabasesMysqlTables struct {
	// Table name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.66.0/docs/resources/datastream_stream#table DatastreamStream#table}
	Table *string `field:"required" json:"table" yaml:"table"`
	// mysql_columns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.66.0/docs/resources/datastream_stream#mysql_columns DatastreamStream#mysql_columns}
	MysqlColumns interface{} `field:"optional" json:"mysqlColumns" yaml:"mysqlColumns"`
}

