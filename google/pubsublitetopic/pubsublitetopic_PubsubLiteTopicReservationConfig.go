package pubsublitetopic


type PubsubLiteTopicReservationConfig struct {
	// The Reservation to use for this topic's throughput capacity.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/pubsub_lite_topic#throughput_reservation PubsubLiteTopic#throughput_reservation}
	ThroughputReservation *string `field:"optional" json:"throughputReservation" yaml:"throughputReservation"`
}

