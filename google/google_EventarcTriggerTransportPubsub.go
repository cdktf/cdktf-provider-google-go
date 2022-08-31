// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type EventarcTriggerTransportPubsub struct {
	// Optional.
	//
	// The name of the Pub/Sub topic created and managed by Eventarc system as a transport for the event delivery. Format: `projects/{PROJECT_ID}/topics/{TOPIC_NAME}. You may set an existing topic for triggers of the type google.cloud.pubsub.topic.v1.messagePublished` only. The topic you provide here will not be deleted by Eventarc at trigger deletion.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/eventarc_trigger#topic EventarcTrigger#topic}
	Topic *string `field:"optional" json:"topic" yaml:"topic"`
}

