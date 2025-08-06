// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package memorystoreinstance


type MemorystoreInstanceMaintenancePolicy struct {
	// weekly_maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/memorystore_instance#weekly_maintenance_window MemorystoreInstance#weekly_maintenance_window}
	WeeklyMaintenanceWindow interface{} `field:"optional" json:"weeklyMaintenanceWindow" yaml:"weeklyMaintenanceWindow"`
}

