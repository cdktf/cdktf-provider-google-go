package datastreamstream


type DatastreamStreamBackfillAllMysqlExcludedObjects struct {
	// mysql_databases block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.74.0/docs/resources/datastream_stream#mysql_databases DatastreamStream#mysql_databases}
	MysqlDatabases interface{} `field:"required" json:"mysqlDatabases" yaml:"mysqlDatabases"`
}

