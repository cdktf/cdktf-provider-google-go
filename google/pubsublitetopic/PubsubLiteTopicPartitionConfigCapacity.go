package pubsublitetopic


type PubsubLiteTopicPartitionConfigCapacity struct {
	// Subscribe throughput capacity per partition in MiB/s. Must be >= 4 and <= 16.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/pubsub_lite_topic#publish_mib_per_sec PubsubLiteTopic#publish_mib_per_sec}
	PublishMibPerSec *float64 `field:"required" json:"publishMibPerSec" yaml:"publishMibPerSec"`
	// Publish throughput capacity per partition in MiB/s. Must be >= 4 and <= 16.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/pubsub_lite_topic#subscribe_mib_per_sec PubsubLiteTopic#subscribe_mib_per_sec}
	SubscribeMibPerSec *float64 `field:"required" json:"subscribeMibPerSec" yaml:"subscribeMibPerSec"`
}

