package computeresourcepolicy


type ComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklySchedule struct {
	// day_of_weeks block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_resource_policy#day_of_weeks ComputeResourcePolicy#day_of_weeks}
	DayOfWeeks interface{} `field:"required" json:"dayOfWeeks" yaml:"dayOfWeeks"`
}

