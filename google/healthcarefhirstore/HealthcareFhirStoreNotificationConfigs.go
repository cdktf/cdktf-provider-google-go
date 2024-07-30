// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package healthcarefhirstore


type HealthcareFhirStoreNotificationConfigs struct {
	// The Cloud Pub/Sub topic that notifications of changes are published on.
	//
	// Supplied by the client.
	// PubsubMessage.Data will contain the resource name. PubsubMessage.MessageId is the ID of this message.
	// It is guaranteed to be unique within the topic. PubsubMessage.PublishTime is the time at which the message
	// was published. Notifications are only sent if the topic is non-empty. Topic names must be scoped to a
	// project. service-PROJECT_NUMBER@gcp-sa-healthcare.iam.gserviceaccount.com must have publisher permissions on the given
	// Cloud Pub/Sub topic. Not having adequate permissions will cause the calls that send notifications to fail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/healthcare_fhir_store#pubsub_topic HealthcareFhirStore#pubsub_topic}
	PubsubTopic *string `field:"required" json:"pubsubTopic" yaml:"pubsubTopic"`
	// Whether to send full FHIR resource to this Pub/Sub topic for Create and Update operation.
	//
	// Note that setting this to true does not guarantee that all resources will be sent in the format of
	// full FHIR resource. When a resource change is too large or during heavy traffic, only the resource name will be
	// sent. Clients should always check the "payloadType" label from a Pub/Sub message to determine whether
	// it needs to fetch the full resource as a separate operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/healthcare_fhir_store#send_full_resource HealthcareFhirStore#send_full_resource}
	SendFullResource interface{} `field:"optional" json:"sendFullResource" yaml:"sendFullResource"`
	// Whether to send full FHIR resource to this Pub/Sub topic for deleting FHIR resource.
	//
	// Note that setting this to
	// true does not guarantee that all previous resources will be sent in the format of full FHIR resource. When a
	// resource change is too large or during heavy traffic, only the resource name will be sent. Clients should always
	// check the "payloadType" label from a Pub/Sub message to determine whether it needs to fetch the full previous
	// resource as a separate operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/healthcare_fhir_store#send_previous_resource_on_delete HealthcareFhirStore#send_previous_resource_on_delete}
	SendPreviousResourceOnDelete interface{} `field:"optional" json:"sendPreviousResourceOnDelete" yaml:"sendPreviousResourceOnDelete"`
}

