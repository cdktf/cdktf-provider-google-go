// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databasemigrationserviceconnectionprofile


type DatabaseMigrationServiceConnectionProfileAlloydbSettingsPrimaryInstanceSettingsMachineConfig struct {
	// The number of CPU's in the VM instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.80.0/docs/resources/database_migration_service_connection_profile#cpu_count DatabaseMigrationServiceConnectionProfile#cpu_count}
	CpuCount *float64 `field:"required" json:"cpuCount" yaml:"cpuCount"`
}

