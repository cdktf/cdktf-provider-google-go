package computeresourcepolicy


type ComputeResourcePolicySnapshotSchedulePolicySchedule struct {
	// daily_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_resource_policy#daily_schedule ComputeResourcePolicy#daily_schedule}
	DailySchedule *ComputeResourcePolicySnapshotSchedulePolicyScheduleDailySchedule `field:"optional" json:"dailySchedule" yaml:"dailySchedule"`
	// hourly_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_resource_policy#hourly_schedule ComputeResourcePolicy#hourly_schedule}
	HourlySchedule *ComputeResourcePolicySnapshotSchedulePolicyScheduleHourlySchedule `field:"optional" json:"hourlySchedule" yaml:"hourlySchedule"`
	// weekly_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_resource_policy#weekly_schedule ComputeResourcePolicy#weekly_schedule}
	WeeklySchedule *ComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklySchedule `field:"optional" json:"weeklySchedule" yaml:"weeklySchedule"`
}

