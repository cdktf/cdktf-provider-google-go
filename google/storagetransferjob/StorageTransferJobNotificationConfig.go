package storagetransferjob


type StorageTransferJobNotificationConfig struct {
	// The desired format of the notification message payloads. One of "NONE" or "JSON".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/storage_transfer_job#payload_format StorageTransferJob#payload_format}
	PayloadFormat *string `field:"required" json:"payloadFormat" yaml:"payloadFormat"`
	// The Topic.name of the Pub/Sub topic to which to publish notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/storage_transfer_job#pubsub_topic StorageTransferJob#pubsub_topic}
	PubsubTopic *string `field:"required" json:"pubsubTopic" yaml:"pubsubTopic"`
	// Event types for which a notification is desired.
	//
	// If empty, send notifications for all event types. The valid types are "TRANSFER_OPERATION_SUCCESS", "TRANSFER_OPERATION_FAILED", "TRANSFER_OPERATION_ABORTED".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/storage_transfer_job#event_types StorageTransferJob#event_types}
	EventTypes *[]*string `field:"optional" json:"eventTypes" yaml:"eventTypes"`
}

