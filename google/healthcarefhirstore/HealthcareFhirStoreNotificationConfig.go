package healthcarefhirstore


type HealthcareFhirStoreNotificationConfig struct {
	// The Cloud Pub/Sub topic that notifications of changes are published on.
	//
	// Supplied by the client.
	// PubsubMessage.Data will contain the resource name. PubsubMessage.MessageId is the ID of this message.
	// It is guaranteed to be unique within the topic. PubsubMessage.PublishTime is the time at which the message
	// was published. Notifications are only sent if the topic is non-empty. Topic names must be scoped to a
	// project. service-PROJECT_NUMBER@gcp-sa-healthcare.iam.gserviceaccount.com must have publisher permissions on the given
	// Cloud Pub/Sub topic. Not having adequate permissions will cause the calls that send notifications to fail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/healthcare_fhir_store#pubsub_topic HealthcareFhirStore#pubsub_topic}
	PubsubTopic *string `field:"required" json:"pubsubTopic" yaml:"pubsubTopic"`
}

