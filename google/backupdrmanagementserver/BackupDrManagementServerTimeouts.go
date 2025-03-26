// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package backupdrmanagementserver


type BackupDrManagementServerTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/backup_dr_management_server#create BackupDrManagementServer#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/backup_dr_management_server#delete BackupDrManagementServer#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

