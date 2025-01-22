// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeresourcepolicy


type ComputeResourcePolicySnapshotSchedulePolicyScheduleWeeklySchedule struct {
	// day_of_weeks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/compute_resource_policy#day_of_weeks ComputeResourcePolicy#day_of_weeks}
	DayOfWeeks interface{} `field:"required" json:"dayOfWeeks" yaml:"dayOfWeeks"`
}

