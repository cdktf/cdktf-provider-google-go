package cloudiotregistry


type CloudiotRegistryEventNotificationConfigs struct {
	// PubSub topic name to publish device events.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudiot_registry#pubsub_topic_name CloudiotRegistry#pubsub_topic_name}
	PubsubTopicName *string `field:"required" json:"pubsubTopicName" yaml:"pubsubTopicName"`
	// If the subfolder name matches this string exactly, this configuration will be used.
	//
	// The string must not include the
	// leading '/' character. If empty, all strings are matched. Empty
	// value can only be used for the last 'event_notification_configs'
	// item.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudiot_registry#subfolder_matches CloudiotRegistry#subfolder_matches}
	SubfolderMatches *string `field:"optional" json:"subfolderMatches" yaml:"subfolderMatches"`
}

