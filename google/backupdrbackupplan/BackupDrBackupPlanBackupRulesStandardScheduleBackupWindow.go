// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package backupdrbackupplan


type BackupDrBackupPlanBackupRulesStandardScheduleBackupWindow struct {
	// The hour of the day (0-23) when the window starts, for example, if the value of the start hour of the day is 6, that means the backup window starts at 6:00.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/backup_dr_backup_plan#start_hour_of_day BackupDrBackupPlan#start_hour_of_day}
	StartHourOfDay *float64 `field:"required" json:"startHourOfDay" yaml:"startHourOfDay"`
	// The hour of the day (1-24) when the window ends, for example, if the value of end hour of the day is 10, that means the backup window end time is 10:00.
	//
	// The end hour of the day should be greater than the start
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/backup_dr_backup_plan#end_hour_of_day BackupDrBackupPlan#end_hour_of_day}
	EndHourOfDay *float64 `field:"optional" json:"endHourOfDay" yaml:"endHourOfDay"`
}

