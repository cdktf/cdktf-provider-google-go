// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkebackuprestoreplaniambinding


type GkeBackupRestorePlanIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.4.0/docs/resources/gke_backup_restore_plan_iam_binding#expression GkeBackupRestorePlanIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.4.0/docs/resources/gke_backup_restore_plan_iam_binding#title GkeBackupRestorePlanIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.4.0/docs/resources/gke_backup_restore_plan_iam_binding#description GkeBackupRestorePlanIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

