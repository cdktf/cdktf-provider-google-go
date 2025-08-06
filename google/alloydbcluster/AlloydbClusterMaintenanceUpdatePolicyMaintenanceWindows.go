// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alloydbcluster


type AlloydbClusterMaintenanceUpdatePolicyMaintenanceWindows struct {
	// Preferred day of the week for maintenance, e.g. MONDAY, TUESDAY, etc. Possible values: ["MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/alloydb_cluster#day AlloydbCluster#day}
	Day *string `field:"required" json:"day" yaml:"day"`
	// start_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/alloydb_cluster#start_time AlloydbCluster#start_time}
	StartTime *AlloydbClusterMaintenanceUpdatePolicyMaintenanceWindowsStartTime `field:"required" json:"startTime" yaml:"startTime"`
}

