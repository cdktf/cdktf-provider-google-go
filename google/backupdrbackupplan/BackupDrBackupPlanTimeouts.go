// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package backupdrbackupplan


type BackupDrBackupPlanTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/backup_dr_backup_plan#create BackupDrBackupPlan#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/backup_dr_backup_plan#delete BackupDrBackupPlan#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

