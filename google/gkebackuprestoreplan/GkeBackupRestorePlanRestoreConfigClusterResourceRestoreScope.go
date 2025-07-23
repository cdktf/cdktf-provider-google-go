// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkebackuprestoreplan


type GkeBackupRestorePlanRestoreConfigClusterResourceRestoreScope struct {
	// If True, all valid cluster-scoped resources will be restored. Mutually exclusive to any other field in 'clusterResourceRestoreScope'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gke_backup_restore_plan#all_group_kinds GkeBackupRestorePlan#all_group_kinds}
	AllGroupKinds interface{} `field:"optional" json:"allGroupKinds" yaml:"allGroupKinds"`
	// excluded_group_kinds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gke_backup_restore_plan#excluded_group_kinds GkeBackupRestorePlan#excluded_group_kinds}
	ExcludedGroupKinds interface{} `field:"optional" json:"excludedGroupKinds" yaml:"excludedGroupKinds"`
	// If True, no cluster-scoped resources will be restored. Mutually exclusive to any other field in 'clusterResourceRestoreScope'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gke_backup_restore_plan#no_group_kinds GkeBackupRestorePlan#no_group_kinds}
	NoGroupKinds interface{} `field:"optional" json:"noGroupKinds" yaml:"noGroupKinds"`
	// selected_group_kinds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gke_backup_restore_plan#selected_group_kinds GkeBackupRestorePlan#selected_group_kinds}
	SelectedGroupKinds interface{} `field:"optional" json:"selectedGroupKinds" yaml:"selectedGroupKinds"`
}

