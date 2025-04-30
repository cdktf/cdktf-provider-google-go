// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package backupdrmanagementserver


type BackupDrManagementServerNetworks struct {
	// Network with format 'projects/{{project_id}}/global/networks/{{network_id}}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/backup_dr_management_server#network BackupDrManagementServer#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// Type of Network peeringMode Default value: "PRIVATE_SERVICE_ACCESS" Possible values: ["PRIVATE_SERVICE_ACCESS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/backup_dr_management_server#peering_mode BackupDrManagementServer#peering_mode}
	PeeringMode *string `field:"optional" json:"peeringMode" yaml:"peeringMode"`
}

