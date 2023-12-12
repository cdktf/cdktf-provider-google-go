// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkebackupbackupplan


type GkeBackupBackupPlanBackupSchedule struct {
	// A standard cron string that defines a repeating schedule for creating Backups via this BackupPlan.
	//
	// If this is defined, then backupRetainDays must also be defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.9.0/docs/resources/gke_backup_backup_plan#cron_schedule GkeBackupBackupPlan#cron_schedule}
	CronSchedule *string `field:"optional" json:"cronSchedule" yaml:"cronSchedule"`
	// This flag denotes whether automatic Backup creation is paused for this BackupPlan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.9.0/docs/resources/gke_backup_backup_plan#paused GkeBackupBackupPlan#paused}
	Paused interface{} `field:"optional" json:"paused" yaml:"paused"`
}

