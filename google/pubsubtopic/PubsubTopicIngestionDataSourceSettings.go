// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pubsubtopic


type PubsubTopicIngestionDataSourceSettings struct {
	// aws_kinesis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/pubsub_topic#aws_kinesis PubsubTopic#aws_kinesis}
	AwsKinesis *PubsubTopicIngestionDataSourceSettingsAwsKinesis `field:"optional" json:"awsKinesis" yaml:"awsKinesis"`
	// cloud_storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/pubsub_topic#cloud_storage PubsubTopic#cloud_storage}
	CloudStorage *PubsubTopicIngestionDataSourceSettingsCloudStorage `field:"optional" json:"cloudStorage" yaml:"cloudStorage"`
	// platform_logs_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/pubsub_topic#platform_logs_settings PubsubTopic#platform_logs_settings}
	PlatformLogsSettings *PubsubTopicIngestionDataSourceSettingsPlatformLogsSettings `field:"optional" json:"platformLogsSettings" yaml:"platformLogsSettings"`
}

