// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pubsubtopic


type PubsubTopicIngestionDataSourceSettingsAzureEventHubs struct {
	// The Azure event hub client ID to use for ingestion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/pubsub_topic#client_id PubsubTopic#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The Azure event hub to ingest data from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/pubsub_topic#event_hub PubsubTopic#event_hub}
	EventHub *string `field:"optional" json:"eventHub" yaml:"eventHub"`
	// The GCP service account to be used for Federated Identity authentication with Azure (via a 'AssumeRoleWithWebIdentity' call for the provided role).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/pubsub_topic#gcp_service_account PubsubTopic#gcp_service_account}
	GcpServiceAccount *string `field:"optional" json:"gcpServiceAccount" yaml:"gcpServiceAccount"`
	// The Azure event hub namespace to ingest data from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/pubsub_topic#namespace PubsubTopic#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The name of the resource group within an Azure subscription.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/pubsub_topic#resource_group PubsubTopic#resource_group}
	ResourceGroup *string `field:"optional" json:"resourceGroup" yaml:"resourceGroup"`
	// The Azure event hub subscription ID to use for ingestion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/pubsub_topic#subscription_id PubsubTopic#subscription_id}
	SubscriptionId *string `field:"optional" json:"subscriptionId" yaml:"subscriptionId"`
	// The Azure event hub tenant ID to use for ingestion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/pubsub_topic#tenant_id PubsubTopic#tenant_id}
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

