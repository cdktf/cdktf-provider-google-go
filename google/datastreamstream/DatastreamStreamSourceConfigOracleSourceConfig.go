package datastreamstream


type DatastreamStreamSourceConfigOracleSourceConfig struct {
	// drop_large_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/datastream_stream#drop_large_objects DatastreamStream#drop_large_objects}
	DropLargeObjects *DatastreamStreamSourceConfigOracleSourceConfigDropLargeObjects `field:"optional" json:"dropLargeObjects" yaml:"dropLargeObjects"`
	// exclude_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/datastream_stream#exclude_objects DatastreamStream#exclude_objects}
	ExcludeObjects *DatastreamStreamSourceConfigOracleSourceConfigExcludeObjects `field:"optional" json:"excludeObjects" yaml:"excludeObjects"`
	// include_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/datastream_stream#include_objects DatastreamStream#include_objects}
	IncludeObjects *DatastreamStreamSourceConfigOracleSourceConfigIncludeObjects `field:"optional" json:"includeObjects" yaml:"includeObjects"`
	// Maximum number of concurrent backfill tasks.
	//
	// The number should be non negative.
	// If not set (or set to 0), the system's default value will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/datastream_stream#max_concurrent_backfill_tasks DatastreamStream#max_concurrent_backfill_tasks}
	MaxConcurrentBackfillTasks *float64 `field:"optional" json:"maxConcurrentBackfillTasks" yaml:"maxConcurrentBackfillTasks"`
	// Maximum number of concurrent CDC tasks.
	//
	// The number should be non negative.
	// If not set (or set to 0), the system's default value will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/datastream_stream#max_concurrent_cdc_tasks DatastreamStream#max_concurrent_cdc_tasks}
	MaxConcurrentCdcTasks *float64 `field:"optional" json:"maxConcurrentCdcTasks" yaml:"maxConcurrentCdcTasks"`
	// stream_large_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/datastream_stream#stream_large_objects DatastreamStream#stream_large_objects}
	StreamLargeObjects *DatastreamStreamSourceConfigOracleSourceConfigStreamLargeObjects `field:"optional" json:"streamLargeObjects" yaml:"streamLargeObjects"`
}

