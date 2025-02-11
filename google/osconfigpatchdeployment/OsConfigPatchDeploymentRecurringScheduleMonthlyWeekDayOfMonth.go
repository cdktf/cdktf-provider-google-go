// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osconfigpatchdeployment


type OsConfigPatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth struct {
	// A day of the week. Possible values: ["MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/os_config_patch_deployment#day_of_week OsConfigPatchDeployment#day_of_week}
	DayOfWeek *string `field:"required" json:"dayOfWeek" yaml:"dayOfWeek"`
	// Week number in a month.
	//
	// 1-4 indicates the 1st to 4th week of the month. -1 indicates the last week of the month.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/os_config_patch_deployment#week_ordinal OsConfigPatchDeployment#week_ordinal}
	WeekOrdinal *float64 `field:"required" json:"weekOrdinal" yaml:"weekOrdinal"`
	// Represents the number of days before or after the given week day of month that the patch deployment is scheduled for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/os_config_patch_deployment#day_offset OsConfigPatchDeployment#day_offset}
	DayOffset *float64 `field:"optional" json:"dayOffset" yaml:"dayOffset"`
}

