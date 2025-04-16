// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkebackupbackupplaniambinding


type GkeBackupBackupPlanIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/gke_backup_backup_plan_iam_binding#expression GkeBackupBackupPlanIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/gke_backup_backup_plan_iam_binding#title GkeBackupBackupPlanIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/gke_backup_backup_plan_iam_binding#description GkeBackupBackupPlanIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

