package datastreamstream


type DatastreamStreamBackfillAllMysqlExcludedObjectsMysqlDatabasesMysqlTables struct {
	// Table name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/datastream_stream#table DatastreamStream#table}
	Table *string `field:"required" json:"table" yaml:"table"`
	// mysql_columns block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/datastream_stream#mysql_columns DatastreamStream#mysql_columns}
	MysqlColumns interface{} `field:"optional" json:"mysqlColumns" yaml:"mysqlColumns"`
}
