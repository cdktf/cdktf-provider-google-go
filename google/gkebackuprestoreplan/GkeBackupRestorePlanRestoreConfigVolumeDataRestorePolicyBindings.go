// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkebackuprestoreplan


type GkeBackupRestorePlanRestoreConfigVolumeDataRestorePolicyBindings struct {
	// Specifies the mechanism to be used to restore this volume data.
	//
	// See https://cloud.google.com/kubernetes-engine/docs/add-on/backup-for-gke/reference/rest/v1/RestoreConfig#VolumeDataRestorePolicy
	// for more information on each policy option. Possible values: ["RESTORE_VOLUME_DATA_FROM_BACKUP", "REUSE_VOLUME_HANDLE_FROM_BACKUP", "NO_VOLUME_DATA_RESTORATION"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/gke_backup_restore_plan#policy GkeBackupRestorePlan#policy}
	Policy *string `field:"required" json:"policy" yaml:"policy"`
	// The volume type, as determined by the PVC's bound PV, to apply the policy to. Possible values: ["GCE_PERSISTENT_DISK"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.8.0/docs/resources/gke_backup_restore_plan#volume_type GkeBackupRestorePlan#volume_type}
	VolumeType *string `field:"required" json:"volumeType" yaml:"volumeType"`
}

