package computeresourcepolicy


type ComputeResourcePolicySnapshotSchedulePolicyScheduleDailySchedule struct {
	// Defines a schedule with units measured in days.
	//
	// The value determines how many days pass between the start of each cycle. Days in cycle for snapshot schedule policy must be 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_resource_policy#days_in_cycle ComputeResourcePolicy#days_in_cycle}
	DaysInCycle *float64 `field:"required" json:"daysInCycle" yaml:"daysInCycle"`
	// This must be in UTC format that resolves to one of 00:00, 04:00, 08:00, 12:00, 16:00, or 20:00.
	//
	// For example,
	// both 13:00-5 and 08:00 are valid.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_resource_policy#start_time ComputeResourcePolicy#start_time}
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
}

