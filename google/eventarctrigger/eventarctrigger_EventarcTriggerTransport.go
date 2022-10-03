package eventarctrigger


type EventarcTriggerTransport struct {
	// pubsub block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/eventarc_trigger#pubsub EventarcTrigger#pubsub}
	Pubsub *EventarcTriggerTransportPubsub `field:"optional" json:"pubsub" yaml:"pubsub"`
}

