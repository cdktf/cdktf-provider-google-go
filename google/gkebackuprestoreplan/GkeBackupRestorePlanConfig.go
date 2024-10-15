// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkebackuprestoreplan

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GkeBackupRestorePlanConfig struct {
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
	// A reference to the BackupPlan from which Backups may be used as the source for Restores created via this RestorePlan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/gke_backup_restore_plan#backup_plan GkeBackupRestorePlan#backup_plan}
	BackupPlan *string `field:"required" json:"backupPlan" yaml:"backupPlan"`
	// The source cluster from which Restores will be created via this RestorePlan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/gke_backup_restore_plan#cluster GkeBackupRestorePlan#cluster}
	Cluster *string `field:"required" json:"cluster" yaml:"cluster"`
	// The region of the Restore Plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/gke_backup_restore_plan#location GkeBackupRestorePlan#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The full name of the BackupPlan Resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/gke_backup_restore_plan#name GkeBackupRestorePlan#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// restore_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/gke_backup_restore_plan#restore_config GkeBackupRestorePlan#restore_config}
	RestoreConfig *GkeBackupRestorePlanRestoreConfig `field:"required" json:"restoreConfig" yaml:"restoreConfig"`
	// User specified descriptive string for this RestorePlan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/gke_backup_restore_plan#description GkeBackupRestorePlan#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/gke_backup_restore_plan#id GkeBackupRestorePlan#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Description: A set of custom labels supplied by the user.
	//
	// A list of key->value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/gke_backup_restore_plan#labels GkeBackupRestorePlan#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/gke_backup_restore_plan#project GkeBackupRestorePlan#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/gke_backup_restore_plan#timeouts GkeBackupRestorePlan#timeouts}
	Timeouts *GkeBackupRestorePlanTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

