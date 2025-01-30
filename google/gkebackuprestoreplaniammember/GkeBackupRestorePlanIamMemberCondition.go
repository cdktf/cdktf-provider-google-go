// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkebackuprestoreplaniammember


type GkeBackupRestorePlanIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.1/docs/resources/gke_backup_restore_plan_iam_member#expression GkeBackupRestorePlanIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.1/docs/resources/gke_backup_restore_plan_iam_member#title GkeBackupRestorePlanIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.18.1/docs/resources/gke_backup_restore_plan_iam_member#description GkeBackupRestorePlanIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

