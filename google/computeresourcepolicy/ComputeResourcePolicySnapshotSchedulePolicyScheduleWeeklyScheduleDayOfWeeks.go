// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeresourcepolicy


type ComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklyScheduleDayOfWeeks struct {
	// The day of the week to create the snapshot.
	//
	// e.g. MONDAY Possible values: ["MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/compute_resource_policy#day ComputeResourcePolicy#day}
	Day *string `field:"required" json:"day" yaml:"day"`
	// Time within the window to start the operations.
	//
	// It must be in format "HH:MM", where HH : [00-23] and MM : [00-00] GMT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.2/docs/resources/compute_resource_policy#start_time ComputeResourcePolicy#start_time}
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
}

