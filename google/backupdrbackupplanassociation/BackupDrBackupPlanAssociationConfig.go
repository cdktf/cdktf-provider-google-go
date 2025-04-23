// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package backupdrbackupplanassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BackupDrBackupPlanAssociationConfig struct {
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
	// The BP with which resource needs to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/backup_dr_backup_plan_association#backup_plan BackupDrBackupPlanAssociation#backup_plan}
	BackupPlan *string `field:"required" json:"backupPlan" yaml:"backupPlan"`
	// The id of backupplan association.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/backup_dr_backup_plan_association#backup_plan_association_id BackupDrBackupPlanAssociation#backup_plan_association_id}
	BackupPlanAssociationId *string `field:"required" json:"backupPlanAssociationId" yaml:"backupPlanAssociationId"`
	// The location for the backupplan association.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/backup_dr_backup_plan_association#location BackupDrBackupPlanAssociation#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The resource for which BPA needs to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/backup_dr_backup_plan_association#resource BackupDrBackupPlanAssociation#resource}
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// The resource type of workload on which backupplan is applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/backup_dr_backup_plan_association#resource_type BackupDrBackupPlanAssociation#resource_type}
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/backup_dr_backup_plan_association#id BackupDrBackupPlanAssociation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/backup_dr_backup_plan_association#project BackupDrBackupPlanAssociation#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/backup_dr_backup_plan_association#timeouts BackupDrBackupPlanAssociation#timeouts}
	Timeouts *BackupDrBackupPlanAssociationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

