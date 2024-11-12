// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alloydbcluster


type AlloydbClusterRestoreBackupSource struct {
	// The name of the backup that this cluster is restored from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.0/docs/resources/alloydb_cluster#backup_name AlloydbCluster#backup_name}
	BackupName *string `field:"required" json:"backupName" yaml:"backupName"`
}

