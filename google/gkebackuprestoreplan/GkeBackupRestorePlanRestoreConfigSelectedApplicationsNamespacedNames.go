// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkebackuprestoreplan


type GkeBackupRestorePlanRestoreConfigSelectedApplicationsNamespacedNames struct {
	// The name of a Kubernetes Resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/gke_backup_restore_plan#name GkeBackupRestorePlan#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The namespace of a Kubernetes Resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/gke_backup_restore_plan#namespace GkeBackupRestorePlan#namespace}
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
}

