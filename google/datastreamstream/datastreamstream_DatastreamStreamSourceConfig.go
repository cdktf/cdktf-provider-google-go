package datastreamstream


type DatastreamStreamSourceConfig struct {
	// mysql_source_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/datastream_stream#mysql_source_config DatastreamStream#mysql_source_config}
	MysqlSourceConfig *DatastreamStreamSourceConfigMysqlSourceConfig `field:"required" json:"mysqlSourceConfig" yaml:"mysqlSourceConfig"`
	// Source connection profile resource. Format: projects/{project}/locations/{location}/connectionProfiles/{name}.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/datastream_stream#source_connection_profile DatastreamStream#source_connection_profile}
	SourceConnectionProfile *string `field:"required" json:"sourceConnectionProfile" yaml:"sourceConnectionProfile"`
}

