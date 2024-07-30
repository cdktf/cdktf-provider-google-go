// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pubsubsubscription


type PubsubSubscriptionCloudStorageConfigAvroConfig struct {
	// When true, write the subscription name, messageId, publishTime, attributes, and orderingKey as additional fields in the output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/pubsub_subscription#write_metadata PubsubSubscription#write_metadata}
	WriteMetadata interface{} `field:"optional" json:"writeMetadata" yaml:"writeMetadata"`
}

