// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package memorystoreinstance


type MemorystoreInstanceMaintenancePolicyWeeklyMaintenanceWindow struct {
	// The day of week that maintenance updates occur.
	//
	// - DAY_OF_WEEK_UNSPECIFIED: The day of the week is unspecified.
	// - MONDAY: Monday
	// - TUESDAY: Tuesday
	// - WEDNESDAY: Wednesday
	// - THURSDAY: Thursday
	// - FRIDAY: Friday
	// - SATURDAY: Saturday
	// - SUNDAY: Sunday Possible values: ["DAY_OF_WEEK_UNSPECIFIED", "MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/memorystore_instance#day MemorystoreInstance#day}
	Day *string `field:"required" json:"day" yaml:"day"`
	// start_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/memorystore_instance#start_time MemorystoreInstance#start_time}
	StartTime *MemorystoreInstanceMaintenancePolicyWeeklyMaintenanceWindowStartTime `field:"required" json:"startTime" yaml:"startTime"`
}

