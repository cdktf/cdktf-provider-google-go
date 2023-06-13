package pubsublitereservation


type PubsubLiteReservationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/pubsub_lite_reservation#create PubsubLiteReservation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/pubsub_lite_reservation#delete PubsubLiteReservation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/pubsub_lite_reservation#update PubsubLiteReservation#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

