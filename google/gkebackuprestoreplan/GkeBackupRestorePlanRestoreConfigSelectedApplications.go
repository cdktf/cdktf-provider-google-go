// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkebackuprestoreplan


type GkeBackupRestorePlanRestoreConfigSelectedApplications struct {
	// namespaced_names block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/gke_backup_restore_plan#namespaced_names GkeBackupRestorePlan#namespaced_names}
	NamespacedNames interface{} `field:"required" json:"namespacedNames" yaml:"namespacedNames"`
}

