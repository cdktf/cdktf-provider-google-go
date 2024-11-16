// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pubsubtopic


type PubsubTopicIngestionDataSourceSettingsCloudStorage struct {
	// Cloud Storage bucket. The bucket name must be without any prefix like "gs://". See the bucket naming requirements: https://cloud.google.com/storage/docs/buckets#naming.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/pubsub_topic#bucket PubsubTopic#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// avro_format block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/pubsub_topic#avro_format PubsubTopic#avro_format}
	AvroFormat *PubsubTopicIngestionDataSourceSettingsCloudStorageAvroFormat `field:"optional" json:"avroFormat" yaml:"avroFormat"`
	// Glob pattern used to match objects that will be ingested.
	//
	// If unset, all
	// objects will be ingested. See the supported patterns:
	// https://cloud.google.com/storage/docs/json_api/v1/objects/list#list-objects-and-prefixes-using-glob
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/pubsub_topic#match_glob PubsubTopic#match_glob}
	MatchGlob *string `field:"optional" json:"matchGlob" yaml:"matchGlob"`
	// The timestamp set in RFC3339 text format.
	//
	// If set, only objects with a
	// larger or equal timestamp will be ingested. Unset by default, meaning
	// all objects will be ingested.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/pubsub_topic#minimum_object_create_time PubsubTopic#minimum_object_create_time}
	MinimumObjectCreateTime *string `field:"optional" json:"minimumObjectCreateTime" yaml:"minimumObjectCreateTime"`
	// pubsub_avro_format block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/pubsub_topic#pubsub_avro_format PubsubTopic#pubsub_avro_format}
	PubsubAvroFormat *PubsubTopicIngestionDataSourceSettingsCloudStoragePubsubAvroFormat `field:"optional" json:"pubsubAvroFormat" yaml:"pubsubAvroFormat"`
	// text_format block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/pubsub_topic#text_format PubsubTopic#text_format}
	TextFormat *PubsubTopicIngestionDataSourceSettingsCloudStorageTextFormat `field:"optional" json:"textFormat" yaml:"textFormat"`
}

