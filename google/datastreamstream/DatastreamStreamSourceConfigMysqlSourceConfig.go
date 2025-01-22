// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datastreamstream


type DatastreamStreamSourceConfigMysqlSourceConfig struct {
	// binary_log_position block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/datastream_stream#binary_log_position DatastreamStream#binary_log_position}
	BinaryLogPosition *DatastreamStreamSourceConfigMysqlSourceConfigBinaryLogPosition `field:"optional" json:"binaryLogPosition" yaml:"binaryLogPosition"`
	// exclude_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/datastream_stream#exclude_objects DatastreamStream#exclude_objects}
	ExcludeObjects *DatastreamStreamSourceConfigMysqlSourceConfigExcludeObjects `field:"optional" json:"excludeObjects" yaml:"excludeObjects"`
	// gtid block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/datastream_stream#gtid DatastreamStream#gtid}
	Gtid *DatastreamStreamSourceConfigMysqlSourceConfigGtid `field:"optional" json:"gtid" yaml:"gtid"`
	// include_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/datastream_stream#include_objects DatastreamStream#include_objects}
	IncludeObjects *DatastreamStreamSourceConfigMysqlSourceConfigIncludeObjects `field:"optional" json:"includeObjects" yaml:"includeObjects"`
	// Maximum number of concurrent backfill tasks.
	//
	// The number should be non negative.
	// If not set (or set to 0), the system's default value will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/datastream_stream#max_concurrent_backfill_tasks DatastreamStream#max_concurrent_backfill_tasks}
	MaxConcurrentBackfillTasks *float64 `field:"optional" json:"maxConcurrentBackfillTasks" yaml:"maxConcurrentBackfillTasks"`
	// Maximum number of concurrent CDC tasks.
	//
	// The number should be non negative.
	// If not set (or set to 0), the system's default value will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/datastream_stream#max_concurrent_cdc_tasks DatastreamStream#max_concurrent_cdc_tasks}
	MaxConcurrentCdcTasks *float64 `field:"optional" json:"maxConcurrentCdcTasks" yaml:"maxConcurrentCdcTasks"`
}

