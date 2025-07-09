// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package backupdrbackupplan

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BackupDrBackupPlanConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The ID of the backup plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/backup_dr_backup_plan#backup_plan_id BackupDrBackupPlan#backup_plan_id}
	BackupPlanId *string `field:"required" json:"backupPlanId" yaml:"backupPlanId"`
	// backup_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/backup_dr_backup_plan#backup_rules BackupDrBackupPlan#backup_rules}
	BackupRules interface{} `field:"required" json:"backupRules" yaml:"backupRules"`
	// Backup vault where the backups gets stored using this Backup plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/backup_dr_backup_plan#backup_vault BackupDrBackupPlan#backup_vault}
	BackupVault *string `field:"required" json:"backupVault" yaml:"backupVault"`
	// The location for the backup plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/backup_dr_backup_plan#location BackupDrBackupPlan#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The resource type to which the 'BackupPlan' will be applied. Examples include, "compute.googleapis.com/Instance", "compute.googleapis.com/Disk", and "storage.googleapis.com/Bucket".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/backup_dr_backup_plan#resource_type BackupDrBackupPlan#resource_type}
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// The description allows for additional details about 'BackupPlan' and its use cases to be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/backup_dr_backup_plan#description BackupDrBackupPlan#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/backup_dr_backup_plan#id BackupDrBackupPlan#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/backup_dr_backup_plan#project BackupDrBackupPlan#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/backup_dr_backup_plan#timeouts BackupDrBackupPlan#timeouts}
	Timeouts *BackupDrBackupPlanTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

