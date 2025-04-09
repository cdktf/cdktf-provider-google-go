// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pubsubtopic


type PubsubTopicIngestionDataSourceSettings struct {
	// aws_kinesis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/pubsub_topic#aws_kinesis PubsubTopic#aws_kinesis}
	AwsKinesis *PubsubTopicIngestionDataSourceSettingsAwsKinesis `field:"optional" json:"awsKinesis" yaml:"awsKinesis"`
	// aws_msk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/pubsub_topic#aws_msk PubsubTopic#aws_msk}
	AwsMsk *PubsubTopicIngestionDataSourceSettingsAwsMsk `field:"optional" json:"awsMsk" yaml:"awsMsk"`
	// azure_event_hubs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/pubsub_topic#azure_event_hubs PubsubTopic#azure_event_hubs}
	AzureEventHubs *PubsubTopicIngestionDataSourceSettingsAzureEventHubs `field:"optional" json:"azureEventHubs" yaml:"azureEventHubs"`
	// cloud_storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/pubsub_topic#cloud_storage PubsubTopic#cloud_storage}
	CloudStorage *PubsubTopicIngestionDataSourceSettingsCloudStorage `field:"optional" json:"cloudStorage" yaml:"cloudStorage"`
	// confluent_cloud block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/pubsub_topic#confluent_cloud PubsubTopic#confluent_cloud}
	ConfluentCloud *PubsubTopicIngestionDataSourceSettingsConfluentCloud `field:"optional" json:"confluentCloud" yaml:"confluentCloud"`
	// platform_logs_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/pubsub_topic#platform_logs_settings PubsubTopic#platform_logs_settings}
	PlatformLogsSettings *PubsubTopicIngestionDataSourceSettingsPlatformLogsSettings `field:"optional" json:"platformLogsSettings" yaml:"platformLogsSettings"`
}

