// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkebackupbackupplan


type GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsSingleOccurrenceDate struct {
	// Day of a month.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/gke_backup_backup_plan#day GkeBackupBackupPlan#day}
	Day *float64 `field:"optional" json:"day" yaml:"day"`
	// Month of a year.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/gke_backup_backup_plan#month GkeBackupBackupPlan#month}
	Month *float64 `field:"optional" json:"month" yaml:"month"`
	// Year of the date.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/gke_backup_backup_plan#year GkeBackupBackupPlan#year}
	Year *float64 `field:"optional" json:"year" yaml:"year"`
}

