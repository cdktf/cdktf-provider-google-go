// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkebackupbackupplan


type GkeBackupBackupPlanBackupScheduleRpoConfigExclusionWindowsStartTime struct {
	// Hours of day in 24 hour format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/gke_backup_backup_plan#hours GkeBackupBackupPlan#hours}
	Hours *float64 `field:"optional" json:"hours" yaml:"hours"`
	// Minutes of hour of day.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/gke_backup_backup_plan#minutes GkeBackupBackupPlan#minutes}
	Minutes *float64 `field:"optional" json:"minutes" yaml:"minutes"`
	// Fractions of seconds in nanoseconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/gke_backup_backup_plan#nanos GkeBackupBackupPlan#nanos}
	Nanos *float64 `field:"optional" json:"nanos" yaml:"nanos"`
	// Seconds of minutes of the time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/gke_backup_backup_plan#seconds GkeBackupBackupPlan#seconds}
	Seconds *float64 `field:"optional" json:"seconds" yaml:"seconds"`
}

