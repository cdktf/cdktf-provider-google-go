package memcacheinstance


type MemcacheInstanceMaintenancePolicyWeeklyMaintenanceWindow struct {
	// Required.
	//
	// The day of week that maintenance updates occur.
	// - DAY_OF_WEEK_UNSPECIFIED: The day of the week is unspecified.
	// - MONDAY: Monday
	// - TUESDAY: Tuesday
	// - WEDNESDAY: Wednesday
	// - THURSDAY: Thursday
	// - FRIDAY: Friday
	// - SATURDAY: Saturday
	// - SUNDAY: Sunday Possible values: ["DAY_OF_WEEK_UNSPECIFIED", "MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"]
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/memcache_instance#day MemcacheInstance#day}
	Day *string `field:"required" json:"day" yaml:"day"`
	// Required.
	//
	// The length of the maintenance window, ranging from 3 hours to 8 hours.
	// A duration in seconds with up to nine fractional digits,
	// terminated by 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/memcache_instance#duration MemcacheInstance#duration}
	Duration *string `field:"required" json:"duration" yaml:"duration"`
	// start_time block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/memcache_instance#start_time MemcacheInstance#start_time}
	StartTime *MemcacheInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime `field:"required" json:"startTime" yaml:"startTime"`
}

