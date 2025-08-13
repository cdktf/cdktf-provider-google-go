// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pubsubtopic


type PubsubTopicIngestionDataSourceSettingsCloudStorageTextFormat struct {
	// The delimiter to use when using the 'text' format.
	//
	// Each line of text as
	// specified by the delimiter will be set to the 'data' field of a Pub/Sub
	// message. When unset, '\n' is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/pubsub_topic#delimiter PubsubTopic#delimiter}
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
}

