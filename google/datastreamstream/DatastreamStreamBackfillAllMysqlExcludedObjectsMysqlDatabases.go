package datastreamstream


type DatastreamStreamBackfillAllMysqlExcludedObjectsMysqlDatabases struct {
	// Database name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/datastream_stream#database DatastreamStream#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// mysql_tables block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/datastream_stream#mysql_tables DatastreamStream#mysql_tables}
	MysqlTables interface{} `field:"optional" json:"mysqlTables" yaml:"mysqlTables"`
}

