package datastreamstream


type DatastreamStreamSourceConfigMysqlSourceConfigExcludeObjects struct {
	// mysql_databases block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/datastream_stream#mysql_databases DatastreamStream#mysql_databases}
	MysqlDatabases interface{} `field:"required" json:"mysqlDatabases" yaml:"mysqlDatabases"`
}

