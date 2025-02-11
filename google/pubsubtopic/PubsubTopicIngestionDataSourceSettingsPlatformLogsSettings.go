// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pubsubtopic


type PubsubTopicIngestionDataSourceSettingsPlatformLogsSettings struct {
	// The minimum severity level of Platform Logs that will be written.
	//
	// If unspecified,
	// no Platform Logs will be written. Default value: "SEVERITY_UNSPECIFIED" Possible values: ["SEVERITY_UNSPECIFIED", "DISABLED", "DEBUG", "INFO", "WARNING", "ERROR"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/pubsub_topic#severity PubsubTopic#severity}
	Severity *string `field:"optional" json:"severity" yaml:"severity"`
}

