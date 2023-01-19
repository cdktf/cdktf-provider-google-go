package datastreamstream


type DatastreamStreamSourceConfigMysqlSourceConfig struct {
	// exclude_objects block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/datastream_stream#exclude_objects DatastreamStream#exclude_objects}
	ExcludeObjects *DatastreamStreamSourceConfigMysqlSourceConfigExcludeObjects `field:"optional" json:"excludeObjects" yaml:"excludeObjects"`
	// include_objects block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/datastream_stream#include_objects DatastreamStream#include_objects}
	IncludeObjects *DatastreamStreamSourceConfigMysqlSourceConfigIncludeObjects `field:"optional" json:"includeObjects" yaml:"includeObjects"`
	// Maximum number of concurrent CDC tasks.
	//
	// The number should be non negative.
	// If not set (or set to 0), the system's default value will be used.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/datastream_stream#max_concurrent_cdc_tasks DatastreamStream#max_concurrent_cdc_tasks}
	MaxConcurrentCdcTasks *float64 `field:"optional" json:"maxConcurrentCdcTasks" yaml:"maxConcurrentCdcTasks"`
}

