package pubsublitetopic


type PubsubLiteTopicRetentionConfig struct {
	// The provisioned storage, in bytes, per partition.
	//
	// If the number of bytes stored
	// in any of the topic's partitions grows beyond this value, older messages will be
	// dropped to make room for newer ones, regardless of the value of period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/pubsub_lite_topic#per_partition_bytes PubsubLiteTopic#per_partition_bytes}
	PerPartitionBytes *string `field:"required" json:"perPartitionBytes" yaml:"perPartitionBytes"`
	// How long a published message is retained.
	//
	// If unset, messages will be retained as
	// long as the bytes retained for each partition is below perPartitionBytes. A
	// duration in seconds with up to nine fractional digits, terminated by 's'.
	// Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/pubsub_lite_topic#period PubsubLiteTopic#period}
	Period *string `field:"optional" json:"period" yaml:"period"`
}

