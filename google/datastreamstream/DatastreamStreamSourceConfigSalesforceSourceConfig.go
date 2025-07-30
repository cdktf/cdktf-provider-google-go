// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datastreamstream


type DatastreamStreamSourceConfigSalesforceSourceConfig struct {
	// Salesforce objects polling interval.
	//
	// The interval at which new changes will be polled for each object. The duration must be between 5 minutes and 24 hours.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/datastream_stream#polling_interval DatastreamStream#polling_interval}
	PollingInterval *string `field:"required" json:"pollingInterval" yaml:"pollingInterval"`
	// exclude_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/datastream_stream#exclude_objects DatastreamStream#exclude_objects}
	ExcludeObjects *DatastreamStreamSourceConfigSalesforceSourceConfigExcludeObjects `field:"optional" json:"excludeObjects" yaml:"excludeObjects"`
	// include_objects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/datastream_stream#include_objects DatastreamStream#include_objects}
	IncludeObjects *DatastreamStreamSourceConfigSalesforceSourceConfigIncludeObjects `field:"optional" json:"includeObjects" yaml:"includeObjects"`
}

