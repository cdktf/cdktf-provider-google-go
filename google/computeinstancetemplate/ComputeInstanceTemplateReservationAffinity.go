package computeinstancetemplate


type ComputeInstanceTemplateReservationAffinity struct {
	// The type of reservation from which this instance can consume resources.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_instance_template#type ComputeInstanceTemplate#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// specific_reservation block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_instance_template#specific_reservation ComputeInstanceTemplate#specific_reservation}
	SpecificReservation *ComputeInstanceTemplateReservationAffinitySpecificReservation `field:"optional" json:"specificReservation" yaml:"specificReservation"`
}

