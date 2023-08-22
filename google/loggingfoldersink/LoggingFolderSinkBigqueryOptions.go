package loggingfoldersink


type LoggingFolderSinkBigqueryOptions struct {
	// Whether to use BigQuery's partition tables.
	//
	// By default, Logging creates dated tables based on the log entries' timestamps, e.g. syslog_20170523. With partitioned tables the date suffix is no longer present and special query syntax has to be used instead. In both cases, tables are sharded based on UTC timezone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/logging_folder_sink#use_partitioned_tables LoggingFolderSink#use_partitioned_tables}
	UsePartitionedTables interface{} `field:"required" json:"usePartitionedTables" yaml:"usePartitionedTables"`
}

