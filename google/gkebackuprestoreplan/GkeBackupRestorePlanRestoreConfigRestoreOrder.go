// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkebackuprestoreplan


type GkeBackupRestorePlanRestoreConfigRestoreOrder struct {
	// group_kind_dependencies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/gke_backup_restore_plan#group_kind_dependencies GkeBackupRestorePlan#group_kind_dependencies}
	GroupKindDependencies interface{} `field:"required" json:"groupKindDependencies" yaml:"groupKindDependencies"`
}

