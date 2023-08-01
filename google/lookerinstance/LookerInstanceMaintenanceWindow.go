package lookerinstance


type LookerInstanceMaintenanceWindow struct {
	// Required. Day of the week for this MaintenanceWindow (in UTC).
	//
	// - MONDAY: Monday
	// - TUESDAY: Tuesday
	// - WEDNESDAY: Wednesday
	// - THURSDAY: Thursday
	// - FRIDAY: Friday
	// - SATURDAY: Saturday
	// - SUNDAY: Sunday Possible values: ["MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/looker_instance#day_of_week LookerInstance#day_of_week}
	DayOfWeek *string `field:"required" json:"dayOfWeek" yaml:"dayOfWeek"`
	// start_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/looker_instance#start_time LookerInstance#start_time}
	StartTime *LookerInstanceMaintenanceWindowStartTime `field:"required" json:"startTime" yaml:"startTime"`
}

