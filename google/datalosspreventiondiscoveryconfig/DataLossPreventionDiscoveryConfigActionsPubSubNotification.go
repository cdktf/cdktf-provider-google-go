// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigActionsPubSubNotification struct {
	// How much data to include in the pub/sub message. Possible values: ["TABLE_PROFILE", "RESOURCE_NAME"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/data_loss_prevention_discovery_config#detail_of_message DataLossPreventionDiscoveryConfig#detail_of_message}
	DetailOfMessage *string `field:"optional" json:"detailOfMessage" yaml:"detailOfMessage"`
	// The type of event that triggers a Pub/Sub.
	//
	// At most one PubSubNotification per EventType is permitted. Possible values: ["NEW_PROFILE", "CHANGED_PROFILE", "SCORE_INCREASED", "ERROR_CHANGED"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/data_loss_prevention_discovery_config#event DataLossPreventionDiscoveryConfig#event}
	Event *string `field:"optional" json:"event" yaml:"event"`
	// pubsub_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/data_loss_prevention_discovery_config#pubsub_condition DataLossPreventionDiscoveryConfig#pubsub_condition}
	PubsubCondition *DataLossPreventionDiscoveryConfigActionsPubSubNotificationPubsubCondition `field:"optional" json:"pubsubCondition" yaml:"pubsubCondition"`
	// Cloud Pub/Sub topic to send notifications to. Format is projects/{project}/topics/{topic}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/data_loss_prevention_discovery_config#topic DataLossPreventionDiscoveryConfig#topic}
	Topic *string `field:"optional" json:"topic" yaml:"topic"`
}

