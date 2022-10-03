package computereservation


type ComputeReservationShareSettings struct {
	// project_map block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_reservation#project_map ComputeReservation#project_map}
	ProjectMap interface{} `field:"optional" json:"projectMap" yaml:"projectMap"`
	// Type of sharing for this shared-reservation Possible values: ["LOCAL", "SPECIFIC_PROJECTS"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_reservation#share_type ComputeReservation#share_type}
	ShareType *string `field:"optional" json:"shareType" yaml:"shareType"`
}

