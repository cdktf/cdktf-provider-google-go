package pubsublitetopic


type PubsubLiteTopicPartitionConfig struct {
	// The number of partitions in the topic. Must be at least 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/pubsub_lite_topic#count PubsubLiteTopic#count}
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// capacity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/pubsub_lite_topic#capacity PubsubLiteTopic#capacity}
	Capacity *PubsubLiteTopicPartitionConfigCapacity `field:"optional" json:"capacity" yaml:"capacity"`
}

