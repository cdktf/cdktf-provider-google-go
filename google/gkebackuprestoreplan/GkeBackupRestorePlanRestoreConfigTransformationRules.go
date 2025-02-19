// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkebackuprestoreplan


type GkeBackupRestorePlanRestoreConfigTransformationRules struct {
	// field_actions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.21.0/docs/resources/gke_backup_restore_plan#field_actions GkeBackupRestorePlan#field_actions}
	FieldActions interface{} `field:"required" json:"fieldActions" yaml:"fieldActions"`
	// The description is a user specified string description of the transformation rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.21.0/docs/resources/gke_backup_restore_plan#description GkeBackupRestorePlan#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// resource_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.21.0/docs/resources/gke_backup_restore_plan#resource_filter GkeBackupRestorePlan#resource_filter}
	ResourceFilter *GkeBackupRestorePlanRestoreConfigTransformationRulesResourceFilter `field:"optional" json:"resourceFilter" yaml:"resourceFilter"`
}

