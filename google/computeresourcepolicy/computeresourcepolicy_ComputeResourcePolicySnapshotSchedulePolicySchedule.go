package computeresourcepolicy


type ComputeResourcePolicySnapshotSchedulePolicySchedule struct {
	// daily_schedule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_resource_policy#daily_schedule ComputeResourcePolicy#daily_schedule}
	DailySchedule *ComputeResourcePolicySnapshotSchedulePolicyScheduleDailySchedule `field:"optional" json:"dailySchedule" yaml:"dailySchedule"`
	// hourly_schedule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_resource_policy#hourly_schedule ComputeResourcePolicy#hourly_schedule}
	HourlySchedule *ComputeResourcePolicySnapshotSchedulePolicyScheduleHourlySchedule `field:"optional" json:"hourlySchedule" yaml:"hourlySchedule"`
	// weekly_schedule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_resource_policy#weekly_schedule ComputeResourcePolicy#weekly_schedule}
	WeeklySchedule *ComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklySchedule `field:"optional" json:"weeklySchedule" yaml:"weeklySchedule"`
}

