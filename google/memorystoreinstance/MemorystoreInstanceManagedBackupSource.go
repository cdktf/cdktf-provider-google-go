// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package memorystoreinstance


type MemorystoreInstanceManagedBackupSource struct {
	// Example: //memorystore.googleapis.com/projects/{project}/locations/{location}/backups/{backupId}. In this case, it assumes the backup is under memorystore.googleapis.com.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/memorystore_instance#backup MemorystoreInstance#backup}
	Backup *string `field:"required" json:"backup" yaml:"backup"`
}

