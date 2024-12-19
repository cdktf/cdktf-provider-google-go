// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkebackupbackupplan


type GkeBackupBackupPlanBackupConfig struct {
	// If True, include all namespaced resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/gke_backup_backup_plan#all_namespaces GkeBackupBackupPlan#all_namespaces}
	AllNamespaces interface{} `field:"optional" json:"allNamespaces" yaml:"allNamespaces"`
	// encryption_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/gke_backup_backup_plan#encryption_key GkeBackupBackupPlan#encryption_key}
	EncryptionKey *GkeBackupBackupPlanBackupConfigEncryptionKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// This flag specifies whether Kubernetes Secret resources should be included when they fall into the scope of Backups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/gke_backup_backup_plan#include_secrets GkeBackupBackupPlan#include_secrets}
	IncludeSecrets interface{} `field:"optional" json:"includeSecrets" yaml:"includeSecrets"`
	// This flag specifies whether volume data should be backed up when PVCs are included in the scope of a Backup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/gke_backup_backup_plan#include_volume_data GkeBackupBackupPlan#include_volume_data}
	IncludeVolumeData interface{} `field:"optional" json:"includeVolumeData" yaml:"includeVolumeData"`
	// This flag specifies whether Backups will not fail when Backup for GKE detects Kubernetes configuration that is non-standard or requires additional setup to restore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/gke_backup_backup_plan#permissive_mode GkeBackupBackupPlan#permissive_mode}
	PermissiveMode interface{} `field:"optional" json:"permissiveMode" yaml:"permissiveMode"`
	// selected_applications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/gke_backup_backup_plan#selected_applications GkeBackupBackupPlan#selected_applications}
	SelectedApplications *GkeBackupBackupPlanBackupConfigSelectedApplications `field:"optional" json:"selectedApplications" yaml:"selectedApplications"`
	// selected_namespaces block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.14.1/docs/resources/gke_backup_backup_plan#selected_namespaces GkeBackupBackupPlan#selected_namespaces}
	SelectedNamespaces *GkeBackupBackupPlanBackupConfigSelectedNamespaces `field:"optional" json:"selectedNamespaces" yaml:"selectedNamespaces"`
}

