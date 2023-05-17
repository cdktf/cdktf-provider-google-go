package containernodepool


type ContainerNodePoolNodeConfigReservationAffinity struct {
	// Corresponds to the type of reservation consumption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/container_node_pool#consume_reservation_type ContainerNodePool#consume_reservation_type}
	ConsumeReservationType *string `field:"required" json:"consumeReservationType" yaml:"consumeReservationType"`
	// The label key of a reservation resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/container_node_pool#key ContainerNodePool#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The label values of the reservation resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/container_node_pool#values ContainerNodePool#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

