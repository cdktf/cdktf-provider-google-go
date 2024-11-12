// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oracledatabasecloudexadatainfrastructure


type OracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindow struct {
	// Determines the amount of time the system will wait before the start of each database server patching operation.
	//
	// Custom action timeout is in minutes and
	// valid value is between 15 to 120 (inclusive).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.0/docs/resources/oracle_database_cloud_exadata_infrastructure#custom_action_timeout_mins OracleDatabaseCloudExadataInfrastructure#custom_action_timeout_mins}
	CustomActionTimeoutMins *float64 `field:"optional" json:"customActionTimeoutMins" yaml:"customActionTimeoutMins"`
	// Days during the week when maintenance should be performed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.0/docs/resources/oracle_database_cloud_exadata_infrastructure#days_of_week OracleDatabaseCloudExadataInfrastructure#days_of_week}
	DaysOfWeek *[]*string `field:"optional" json:"daysOfWeek" yaml:"daysOfWeek"`
	// The window of hours during the day when maintenance should be performed.
	//
	// The window is a 4 hour slot. Valid values are:
	//   0 - represents time slot 0:00 - 3:59 UTC
	//   4 - represents time slot 4:00 - 7:59 UTC
	//   8 - represents time slot 8:00 - 11:59 UTC
	//   12 - represents time slot 12:00 - 15:59 UTC
	//   16 - represents time slot 16:00 - 19:59 UTC
	//   20 - represents time slot 20:00 - 23:59 UTC
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.0/docs/resources/oracle_database_cloud_exadata_infrastructure#hours_of_day OracleDatabaseCloudExadataInfrastructure#hours_of_day}
	HoursOfDay *[]*float64 `field:"optional" json:"hoursOfDay" yaml:"hoursOfDay"`
	// If true, enables the configuration of a custom action timeout (waiting period) between database server patching operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.0/docs/resources/oracle_database_cloud_exadata_infrastructure#is_custom_action_timeout_enabled OracleDatabaseCloudExadataInfrastructure#is_custom_action_timeout_enabled}
	IsCustomActionTimeoutEnabled interface{} `field:"optional" json:"isCustomActionTimeoutEnabled" yaml:"isCustomActionTimeoutEnabled"`
	// Lead time window allows user to set a lead time to prepare for a down time.
	//
	// The lead time is in weeks and valid value is between 1 to 4.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.0/docs/resources/oracle_database_cloud_exadata_infrastructure#lead_time_week OracleDatabaseCloudExadataInfrastructure#lead_time_week}
	LeadTimeWeek *float64 `field:"optional" json:"leadTimeWeek" yaml:"leadTimeWeek"`
	// Months during the year when maintenance should be performed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.0/docs/resources/oracle_database_cloud_exadata_infrastructure#months OracleDatabaseCloudExadataInfrastructure#months}
	Months *[]*string `field:"optional" json:"months" yaml:"months"`
	// Cloud CloudExadataInfrastructure node patching method, either "ROLLING"  or "NONROLLING". Default value is ROLLING.   Possible values:  PATCHING_MODE_UNSPECIFIED ROLLING NON_ROLLING.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.0/docs/resources/oracle_database_cloud_exadata_infrastructure#patching_mode OracleDatabaseCloudExadataInfrastructure#patching_mode}
	PatchingMode *string `field:"optional" json:"patchingMode" yaml:"patchingMode"`
	// The maintenance window scheduling preference.   Possible values:  MAINTENANCE_WINDOW_PREFERENCE_UNSPECIFIED CUSTOM_PREFERENCE NO_PREFERENCE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.0/docs/resources/oracle_database_cloud_exadata_infrastructure#preference OracleDatabaseCloudExadataInfrastructure#preference}
	Preference *string `field:"optional" json:"preference" yaml:"preference"`
	// Weeks during the month when maintenance should be performed.
	//
	// Weeks start on
	// the 1st, 8th, 15th, and 22nd days of the month, and have a duration of 7
	// days. Weeks start and end based on calendar dates, not days of the week.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.0/docs/resources/oracle_database_cloud_exadata_infrastructure#weeks_of_month OracleDatabaseCloudExadataInfrastructure#weeks_of_month}
	WeeksOfMonth *[]*float64 `field:"optional" json:"weeksOfMonth" yaml:"weeksOfMonth"`
}

