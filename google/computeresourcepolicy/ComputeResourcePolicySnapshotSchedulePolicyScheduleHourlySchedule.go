package computeresourcepolicy


type ComputeResourcePolicySnapshotSchedulePolicyScheduleHourlySchedule struct {
	// The number of hours between snapshots.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_resource_policy#hours_in_cycle ComputeResourcePolicy#hours_in_cycle}
	HoursInCycle *float64 `field:"required" json:"hoursInCycle" yaml:"hoursInCycle"`
	// Time within the window to start the operations.
	//
	// It must be in an hourly format "HH:MM",
	// where HH : [00-23] and MM : [00] GMT.
	// eg: 21:00
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_resource_policy#start_time ComputeResourcePolicy#start_time}
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
}

