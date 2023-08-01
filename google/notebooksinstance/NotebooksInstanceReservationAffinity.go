package notebooksinstance


type NotebooksInstanceReservationAffinity struct {
	// The type of Compute Reservation. Possible values: ["NO_RESERVATION", "ANY_RESERVATION", "SPECIFIC_RESERVATION"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/notebooks_instance#consume_reservation_type NotebooksInstance#consume_reservation_type}
	ConsumeReservationType *string `field:"required" json:"consumeReservationType" yaml:"consumeReservationType"`
	// Corresponds to the label key of reservation resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/notebooks_instance#key NotebooksInstance#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Corresponds to the label values of reservation resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/notebooks_instance#values NotebooksInstance#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

