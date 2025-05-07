// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pubsubtopic


type PubsubTopicIngestionDataSourceSettingsConfluentCloud struct {
	// The Confluent Cloud bootstrap server. The format is url:port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/pubsub_topic#bootstrap_server PubsubTopic#bootstrap_server}
	BootstrapServer *string `field:"required" json:"bootstrapServer" yaml:"bootstrapServer"`
	// The GCP service account to be used for Federated Identity authentication with Confluent Cloud.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/pubsub_topic#gcp_service_account PubsubTopic#gcp_service_account}
	GcpServiceAccount *string `field:"required" json:"gcpServiceAccount" yaml:"gcpServiceAccount"`
	// Identity pool ID to be used for Federated Identity authentication with Confluent Cloud.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/pubsub_topic#identity_pool_id PubsubTopic#identity_pool_id}
	IdentityPoolId *string `field:"required" json:"identityPoolId" yaml:"identityPoolId"`
	// Name of the Confluent Cloud topic that Pub/Sub will import from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/pubsub_topic#topic PubsubTopic#topic}
	Topic *string `field:"required" json:"topic" yaml:"topic"`
	// The Confluent Cloud cluster ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/pubsub_topic#cluster_id PubsubTopic#cluster_id}
	ClusterId *string `field:"optional" json:"clusterId" yaml:"clusterId"`
}

