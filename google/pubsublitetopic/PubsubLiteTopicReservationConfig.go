package pubsublitetopic


type PubsubLiteTopicReservationConfig struct {
	// The Reservation to use for this topic's throughput capacity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/pubsub_lite_topic#throughput_reservation PubsubLiteTopic#throughput_reservation}
	ThroughputReservation *string `field:"optional" json:"throughputReservation" yaml:"throughputReservation"`
}

