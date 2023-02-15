package computereservation


type ComputeReservationSpecificReservation struct {
	// The number of resources that are allocated.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_reservation#count ComputeReservation#count}
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// instance_properties block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_reservation#instance_properties ComputeReservation#instance_properties}
	InstanceProperties *ComputeReservationSpecificReservationInstanceProperties `field:"required" json:"instanceProperties" yaml:"instanceProperties"`
}

