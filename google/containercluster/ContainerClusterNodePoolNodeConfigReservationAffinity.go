package containercluster


type ContainerClusterNodePoolNodeConfigReservationAffinity struct {
	// Corresponds to the type of reservation consumption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/container_cluster#consume_reservation_type ContainerCluster#consume_reservation_type}
	ConsumeReservationType *string `field:"required" json:"consumeReservationType" yaml:"consumeReservationType"`
	// The label key of a reservation resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/container_cluster#key ContainerCluster#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The label values of the reservation resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/container_cluster#values ContainerCluster#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

