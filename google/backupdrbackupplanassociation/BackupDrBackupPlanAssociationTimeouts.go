// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package backupdrbackupplanassociation


type BackupDrBackupPlanAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/backup_dr_backup_plan_association#create BackupDrBackupPlanAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/backup_dr_backup_plan_association#delete BackupDrBackupPlanAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/backup_dr_backup_plan_association#update BackupDrBackupPlanAssociation#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

