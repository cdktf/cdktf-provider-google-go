package pubsublitetopic


type PubsubLiteTopicTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/pubsub_lite_topic#create PubsubLiteTopic#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/pubsub_lite_topic#delete PubsubLiteTopic#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/pubsub_lite_topic#update PubsubLiteTopic#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

