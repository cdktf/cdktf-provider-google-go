// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loggingorganizationsink


type LoggingOrganizationSinkBigqueryOptions struct {
	// Whether to use BigQuery's partition tables.
	//
	// By default, Logging creates dated tables based on the log entries' timestamps, e.g. syslog_20170523. With partitioned tables the date suffix is no longer present and special query syntax has to be used instead. In both cases, tables are sharded based on UTC timezone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.21.0/docs/resources/logging_organization_sink#use_partitioned_tables LoggingOrganizationSink#use_partitioned_tables}
	UsePartitionedTables interface{} `field:"required" json:"usePartitionedTables" yaml:"usePartitionedTables"`
}

