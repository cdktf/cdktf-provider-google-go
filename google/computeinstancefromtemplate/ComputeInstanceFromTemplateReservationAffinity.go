package computeinstancefromtemplate


type ComputeInstanceFromTemplateReservationAffinity struct {
	// The type of reservation from which this instance can consume resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_from_template#type ComputeInstanceFromTemplate#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// specific_reservation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_instance_from_template#specific_reservation ComputeInstanceFromTemplate#specific_reservation}
	SpecificReservation *ComputeInstanceFromTemplateReservationAffinitySpecificReservation `field:"optional" json:"specificReservation" yaml:"specificReservation"`
}

