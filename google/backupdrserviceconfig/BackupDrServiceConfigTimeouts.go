// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package backupdrserviceconfig


type BackupDrServiceConfigTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/backup_dr_service_config#create BackupDrServiceConfig#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/backup_dr_service_config#delete BackupDrServiceConfig#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

